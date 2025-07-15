// This file is auto-generated, don't edit it. Thanks.
package crm_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AbandonCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AbandonCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s AbandonCustomerHeaders) GoString() string {
	return s.String()
}

func (s *AbandonCustomerHeaders) SetCommonHeaders(v map[string]*string) *AbandonCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AbandonCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *AbandonCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AbandonCustomerRequest struct {
	CustomTrackDesc *string `json:"customTrackDesc,omitempty" xml:"customTrackDesc,omitempty"`
	// This parameter is required.
	InstanceIdList []*string `json:"instanceIdList,omitempty" xml:"instanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123123123
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	OptType        *string `json:"optType,omitempty" xml:"optType,omitempty"`
}

func (s AbandonCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s AbandonCustomerRequest) GoString() string {
	return s.String()
}

func (s *AbandonCustomerRequest) SetCustomTrackDesc(v string) *AbandonCustomerRequest {
	s.CustomTrackDesc = &v
	return s
}

func (s *AbandonCustomerRequest) SetInstanceIdList(v []*string) *AbandonCustomerRequest {
	s.InstanceIdList = v
	return s
}

func (s *AbandonCustomerRequest) SetOperatorUserId(v string) *AbandonCustomerRequest {
	s.OperatorUserId = &v
	return s
}

func (s *AbandonCustomerRequest) SetOptType(v string) *AbandonCustomerRequest {
	s.OptType = &v
	return s
}

type AbandonCustomerResponseBody struct {
	InstanceIdList []*string `json:"instanceIdList,omitempty" xml:"instanceIdList,omitempty" type:"Repeated"`
}

func (s AbandonCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbandonCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *AbandonCustomerResponseBody) SetInstanceIdList(v []*string) *AbandonCustomerResponseBody {
	s.InstanceIdList = v
	return s
}

type AbandonCustomerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbandonCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbandonCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s AbandonCustomerResponse) GoString() string {
	return s.String()
}

func (s *AbandonCustomerResponse) SetHeaders(v map[string]*string) *AbandonCustomerResponse {
	s.Headers = v
	return s
}

func (s *AbandonCustomerResponse) SetStatusCode(v int32) *AbandonCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *AbandonCustomerResponse) SetBody(v *AbandonCustomerResponseBody) *AbandonCustomerResponse {
	s.Body = v
	return s
}

type AddCrmPersonalCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddCrmPersonalCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCrmPersonalCustomerHeaders) GoString() string {
	return s.String()
}

func (s *AddCrmPersonalCustomerHeaders) SetCommonHeaders(v map[string]*string) *AddCrmPersonalCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCrmPersonalCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *AddCrmPersonalCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddCrmPersonalCustomerRequest struct {
	// example:
	//
	// publicDraw
	Action      *string `json:"action,omitempty" xml:"action,omitempty"`
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// This parameter is required.
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// This parameter is required.
	Data       map[string]interface{}                   `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData map[string]interface{}                   `json:"extendData,omitempty" xml:"extendData,omitempty"`
	Permission *AddCrmPersonalCustomerRequestPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// example:
	//
	// crm_customer_personal
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// example:
	//
	// false
	SkipDuplicateCheck *bool `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
}

func (s AddCrmPersonalCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCrmPersonalCustomerRequest) GoString() string {
	return s.String()
}

func (s *AddCrmPersonalCustomerRequest) SetAction(v string) *AddCrmPersonalCustomerRequest {
	s.Action = &v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetCreatorNick(v string) *AddCrmPersonalCustomerRequest {
	s.CreatorNick = &v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetCreatorUserId(v string) *AddCrmPersonalCustomerRequest {
	s.CreatorUserId = &v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetData(v map[string]interface{}) *AddCrmPersonalCustomerRequest {
	s.Data = v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetExtendData(v map[string]interface{}) *AddCrmPersonalCustomerRequest {
	s.ExtendData = v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetPermission(v *AddCrmPersonalCustomerRequestPermission) *AddCrmPersonalCustomerRequest {
	s.Permission = v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetRelationType(v string) *AddCrmPersonalCustomerRequest {
	s.RelationType = &v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetSkipDuplicateCheck(v bool) *AddCrmPersonalCustomerRequest {
	s.SkipDuplicateCheck = &v
	return s
}

type AddCrmPersonalCustomerRequestPermission struct {
	OwnerStaffIds       []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s AddCrmPersonalCustomerRequestPermission) String() string {
	return tea.Prettify(s)
}

func (s AddCrmPersonalCustomerRequestPermission) GoString() string {
	return s.String()
}

func (s *AddCrmPersonalCustomerRequestPermission) SetOwnerStaffIds(v []*string) *AddCrmPersonalCustomerRequestPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *AddCrmPersonalCustomerRequestPermission) SetParticipantStaffIds(v []*string) *AddCrmPersonalCustomerRequestPermission {
	s.ParticipantStaffIds = v
	return s
}

type AddCrmPersonalCustomerResponseBody struct {
	// This parameter is required.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s AddCrmPersonalCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCrmPersonalCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *AddCrmPersonalCustomerResponseBody) SetInstanceId(v string) *AddCrmPersonalCustomerResponseBody {
	s.InstanceId = &v
	return s
}

type AddCrmPersonalCustomerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCrmPersonalCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCrmPersonalCustomerResponse) GoString() string {
	return s.String()
}

func (s *AddCrmPersonalCustomerResponse) SetHeaders(v map[string]*string) *AddCrmPersonalCustomerResponse {
	s.Headers = v
	return s
}

func (s *AddCrmPersonalCustomerResponse) SetStatusCode(v int32) *AddCrmPersonalCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCrmPersonalCustomerResponse) SetBody(v *AddCrmPersonalCustomerResponseBody) *AddCrmPersonalCustomerResponse {
	s.Body = v
	return s
}

type AddCustomerTrackHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddCustomerTrackHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCustomerTrackHeaders) GoString() string {
	return s.String()
}

func (s *AddCustomerTrackHeaders) SetCommonHeaders(v map[string]*string) *AddCustomerTrackHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCustomerTrackHeaders) SetXAcsDingtalkAccessToken(v string) *AddCustomerTrackHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddCustomerTrackRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [华佗](http://******)创建了合同：**今日合同**
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fb037d68-c1bd-4be2-8c3b-6739261d1332
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// {"bizId":"1"}
	ExtraBizInfo *string `json:"extraBizInfo,omitempty" xml:"extraBizInfo,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// fb037d68-c1bd-4be2-8c3b-6739261d1332-1
	IdempotentKey *string `json:"idempotentKey,omitempty" xml:"idempotentKey,omitempty"`
	// example:
	//
	// [华佗](http://******)创建了合同：**今日合同**
	MaskedContent *string `json:"maskedContent,omitempty" xml:"maskedContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager1120
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crm_customer
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 创建合同
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 212
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AddCustomerTrackRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomerTrackRequest) GoString() string {
	return s.String()
}

func (s *AddCustomerTrackRequest) SetContent(v string) *AddCustomerTrackRequest {
	s.Content = &v
	return s
}

func (s *AddCustomerTrackRequest) SetCustomerId(v string) *AddCustomerTrackRequest {
	s.CustomerId = &v
	return s
}

func (s *AddCustomerTrackRequest) SetExtraBizInfo(v string) *AddCustomerTrackRequest {
	s.ExtraBizInfo = &v
	return s
}

func (s *AddCustomerTrackRequest) SetIdempotentKey(v string) *AddCustomerTrackRequest {
	s.IdempotentKey = &v
	return s
}

func (s *AddCustomerTrackRequest) SetMaskedContent(v string) *AddCustomerTrackRequest {
	s.MaskedContent = &v
	return s
}

func (s *AddCustomerTrackRequest) SetOperatorUserId(v string) *AddCustomerTrackRequest {
	s.OperatorUserId = &v
	return s
}

func (s *AddCustomerTrackRequest) SetRelationType(v string) *AddCustomerTrackRequest {
	s.RelationType = &v
	return s
}

func (s *AddCustomerTrackRequest) SetTitle(v string) *AddCustomerTrackRequest {
	s.Title = &v
	return s
}

func (s *AddCustomerTrackRequest) SetType(v int32) *AddCustomerTrackRequest {
	s.Type = &v
	return s
}

type AddCustomerTrackResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddCustomerTrackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCustomerTrackResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomerTrackResponseBody) SetSuccess(v bool) *AddCustomerTrackResponseBody {
	s.Success = &v
	return s
}

type AddCustomerTrackResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomerTrackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomerTrackResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCustomerTrackResponse) GoString() string {
	return s.String()
}

func (s *AddCustomerTrackResponse) SetHeaders(v map[string]*string) *AddCustomerTrackResponse {
	s.Headers = v
	return s
}

func (s *AddCustomerTrackResponse) SetStatusCode(v int32) *AddCustomerTrackResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomerTrackResponse) SetBody(v *AddCustomerTrackResponseBody) *AddCustomerTrackResponse {
	s.Body = v
	return s
}

type AddLeadsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddLeadsHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddLeadsHeaders) GoString() string {
	return s.String()
}

func (s *AddLeadsHeaders) SetCommonHeaders(v map[string]*string) *AddLeadsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddLeadsHeaders) SetXAcsDingtalkAccessToken(v string) *AddLeadsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddLeadsRequest struct {
	// example:
	//
	// 1669360918000
	AssignTimestamp *int64 `json:"assignTimestamp,omitempty" xml:"assignTimestamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager1234
	AssignUserId *string `json:"assignUserId,omitempty" xml:"assignUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager1234
	AssignedUserId *string `json:"assignedUserId,omitempty" xml:"assignedUserId,omitempty"`
	// This parameter is required.
	Leads []*AddLeadsRequestLeads `json:"leads,omitempty" xml:"leads,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// t123123123
	OutTaskId *string `json:"outTaskId,omitempty" xml:"outTaskId,omitempty"`
}

func (s AddLeadsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLeadsRequest) GoString() string {
	return s.String()
}

func (s *AddLeadsRequest) SetAssignTimestamp(v int64) *AddLeadsRequest {
	s.AssignTimestamp = &v
	return s
}

func (s *AddLeadsRequest) SetAssignUserId(v string) *AddLeadsRequest {
	s.AssignUserId = &v
	return s
}

func (s *AddLeadsRequest) SetAssignedUserId(v string) *AddLeadsRequest {
	s.AssignedUserId = &v
	return s
}

func (s *AddLeadsRequest) SetLeads(v []*AddLeadsRequestLeads) *AddLeadsRequest {
	s.Leads = v
	return s
}

func (s *AddLeadsRequest) SetOutTaskId(v string) *AddLeadsRequest {
	s.OutTaskId = &v
	return s
}

type AddLeadsRequestLeads struct {
	// This parameter is required.
	//
	// example:
	//
	// XX公司
	LeadsName *string `json:"leadsName,omitempty" xml:"leadsName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fasd123125
	OutLeadsId *string `json:"outLeadsId,omitempty" xml:"outLeadsId,omitempty"`
}

func (s AddLeadsRequestLeads) String() string {
	return tea.Prettify(s)
}

func (s AddLeadsRequestLeads) GoString() string {
	return s.String()
}

func (s *AddLeadsRequestLeads) SetLeadsName(v string) *AddLeadsRequestLeads {
	s.LeadsName = &v
	return s
}

func (s *AddLeadsRequestLeads) SetOutLeadsId(v string) *AddLeadsRequestLeads {
	s.OutLeadsId = &v
	return s
}

type AddLeadsResponseBody struct {
	OutTaskId *string `json:"outTaskId,omitempty" xml:"outTaskId,omitempty"`
}

func (s AddLeadsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLeadsResponseBody) GoString() string {
	return s.String()
}

func (s *AddLeadsResponseBody) SetOutTaskId(v string) *AddLeadsResponseBody {
	s.OutTaskId = &v
	return s
}

type AddLeadsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLeadsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLeadsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLeadsResponse) GoString() string {
	return s.String()
}

func (s *AddLeadsResponse) SetHeaders(v map[string]*string) *AddLeadsResponse {
	s.Headers = v
	return s
}

func (s *AddLeadsResponse) SetStatusCode(v int32) *AddLeadsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLeadsResponse) SetBody(v *AddLeadsResponseBody) *AddLeadsResponse {
	s.Body = v
	return s
}

type AddMetaModelFieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddMetaModelFieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddMetaModelFieldHeaders) GoString() string {
	return s.String()
}

func (s *AddMetaModelFieldHeaders) SetCommonHeaders(v map[string]*string) *AddMetaModelFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddMetaModelFieldHeaders) SetXAcsDingtalkAccessToken(v string) *AddMetaModelFieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddMetaModelFieldRequest struct {
	// This parameter is required.
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	FieldDTOList []*AddMetaModelFieldRequestFieldDTOList `json:"fieldDTOList,omitempty" xml:"fieldDTOList,omitempty" type:"Repeated"`
	// This parameter is required.
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	Tenant *string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

func (s AddMetaModelFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMetaModelFieldRequest) GoString() string {
	return s.String()
}

func (s *AddMetaModelFieldRequest) SetBizType(v string) *AddMetaModelFieldRequest {
	s.BizType = &v
	return s
}

func (s *AddMetaModelFieldRequest) SetFieldDTOList(v []*AddMetaModelFieldRequestFieldDTOList) *AddMetaModelFieldRequest {
	s.FieldDTOList = v
	return s
}

func (s *AddMetaModelFieldRequest) SetOperatorUserId(v string) *AddMetaModelFieldRequest {
	s.OperatorUserId = &v
	return s
}

func (s *AddMetaModelFieldRequest) SetTenant(v string) *AddMetaModelFieldRequest {
	s.Tenant = &v
	return s
}

type AddMetaModelFieldRequestFieldDTOList struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *AddMetaModelFieldRequestFieldDTOListProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s AddMetaModelFieldRequestFieldDTOList) String() string {
	return tea.Prettify(s)
}

func (s AddMetaModelFieldRequestFieldDTOList) GoString() string {
	return s.String()
}

func (s *AddMetaModelFieldRequestFieldDTOList) SetComponentName(v string) *AddMetaModelFieldRequestFieldDTOList {
	s.ComponentName = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOList) SetProps(v *AddMetaModelFieldRequestFieldDTOListProps) *AddMetaModelFieldRequestFieldDTOList {
	s.Props = v
	return s
}

type AddMetaModelFieldRequestFieldDTOListProps struct {
	Align    *string `json:"align,omitempty" xml:"align,omitempty"`
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice   *int64  `json:"choice,omitempty" xml:"choice,omitempty"`
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	Disabled *bool   `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration *bool   `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	FieldId   *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format    *string `json:"format,omitempty" xml:"format,omitempty"`
	Invisible *bool   `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label               *string                                             `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze *bool                                               `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Link                *string                                             `json:"link,omitempty" xml:"link,omitempty"`
	NeedDetail          *string                                             `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	NotPrint            *string                                             `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	NotUpper            *string                                             `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options             []*AddMetaModelFieldRequestFieldDTOListPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable           *bool                                               `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder         *string                                             `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Required               *bool   `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool   `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Sortable               *bool   `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Unit                   *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s AddMetaModelFieldRequestFieldDTOListProps) String() string {
	return tea.Prettify(s)
}

func (s AddMetaModelFieldRequestFieldDTOListProps) GoString() string {
	return s.String()
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetAlign(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Align = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetBizAlias(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.BizAlias = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetChoice(v int64) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Choice = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetContent(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Content = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetDisabled(v bool) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Disabled = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetDuration(v bool) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Duration = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetFieldId(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.FieldId = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetFormat(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Format = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetInvisible(v bool) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Invisible = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetLabel(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Label = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetLabelEditableFreeze(v bool) *AddMetaModelFieldRequestFieldDTOListProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetLink(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Link = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetNeedDetail(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.NeedDetail = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetNotPrint(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.NotPrint = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetNotUpper(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.NotUpper = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetOptions(v []*AddMetaModelFieldRequestFieldDTOListPropsOptions) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Options = v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetPayEnable(v bool) *AddMetaModelFieldRequestFieldDTOListProps {
	s.PayEnable = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetPlaceholder(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Placeholder = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetRequired(v bool) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Required = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetRequiredEditableFreeze(v bool) *AddMetaModelFieldRequestFieldDTOListProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetSortable(v bool) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Sortable = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListProps) SetUnit(v string) *AddMetaModelFieldRequestFieldDTOListProps {
	s.Unit = &v
	return s
}

type AddMetaModelFieldRequestFieldDTOListPropsOptions struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddMetaModelFieldRequestFieldDTOListPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s AddMetaModelFieldRequestFieldDTOListPropsOptions) GoString() string {
	return s.String()
}

func (s *AddMetaModelFieldRequestFieldDTOListPropsOptions) SetKey(v string) *AddMetaModelFieldRequestFieldDTOListPropsOptions {
	s.Key = &v
	return s
}

func (s *AddMetaModelFieldRequestFieldDTOListPropsOptions) SetValue(v string) *AddMetaModelFieldRequestFieldDTOListPropsOptions {
	s.Value = &v
	return s
}

type AddMetaModelFieldResponseBody struct {
	// This parameter is required.
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
}

func (s AddMetaModelFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMetaModelFieldResponseBody) GoString() string {
	return s.String()
}

func (s *AddMetaModelFieldResponseBody) SetBizType(v string) *AddMetaModelFieldResponseBody {
	s.BizType = &v
	return s
}

type AddMetaModelFieldResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMetaModelFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMetaModelFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMetaModelFieldResponse) GoString() string {
	return s.String()
}

func (s *AddMetaModelFieldResponse) SetHeaders(v map[string]*string) *AddMetaModelFieldResponse {
	s.Headers = v
	return s
}

func (s *AddMetaModelFieldResponse) SetStatusCode(v int32) *AddMetaModelFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMetaModelFieldResponse) SetBody(v *AddMetaModelFieldResponseBody) *AddMetaModelFieldResponse {
	s.Body = v
	return s
}

type AddRelationMetaFieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddRelationMetaFieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddRelationMetaFieldHeaders) GoString() string {
	return s.String()
}

func (s *AddRelationMetaFieldHeaders) SetCommonHeaders(v map[string]*string) *AddRelationMetaFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddRelationMetaFieldHeaders) SetXAcsDingtalkAccessToken(v string) *AddRelationMetaFieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddRelationMetaFieldRequest struct {
	// This parameter is required.
	FieldDTOList []*AddRelationMetaFieldRequestFieldDTOList `json:"fieldDTOList,omitempty" xml:"fieldDTOList,omitempty" type:"Repeated"`
	// This parameter is required.
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// This parameter is required.
	Tenant *string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

func (s AddRelationMetaFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRelationMetaFieldRequest) GoString() string {
	return s.String()
}

func (s *AddRelationMetaFieldRequest) SetFieldDTOList(v []*AddRelationMetaFieldRequestFieldDTOList) *AddRelationMetaFieldRequest {
	s.FieldDTOList = v
	return s
}

func (s *AddRelationMetaFieldRequest) SetOperatorUserId(v string) *AddRelationMetaFieldRequest {
	s.OperatorUserId = &v
	return s
}

func (s *AddRelationMetaFieldRequest) SetRelationType(v string) *AddRelationMetaFieldRequest {
	s.RelationType = &v
	return s
}

func (s *AddRelationMetaFieldRequest) SetTenant(v string) *AddRelationMetaFieldRequest {
	s.Tenant = &v
	return s
}

type AddRelationMetaFieldRequestFieldDTOList struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *AddRelationMetaFieldRequestFieldDTOListProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s AddRelationMetaFieldRequestFieldDTOList) String() string {
	return tea.Prettify(s)
}

func (s AddRelationMetaFieldRequestFieldDTOList) GoString() string {
	return s.String()
}

func (s *AddRelationMetaFieldRequestFieldDTOList) SetComponentName(v string) *AddRelationMetaFieldRequestFieldDTOList {
	s.ComponentName = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOList) SetProps(v *AddRelationMetaFieldRequestFieldDTOListProps) *AddRelationMetaFieldRequestFieldDTOList {
	s.Props = v
	return s
}

type AddRelationMetaFieldRequestFieldDTOListProps struct {
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias  *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice    *int64  `json:"choice,omitempty" xml:"choice,omitempty"`
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	Disabled  *bool   `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration  *bool   `json:"duration,omitempty" xml:"duration,omitempty"`
	FieldId   *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format    *string `json:"format,omitempty" xml:"format,omitempty"`
	Invisible *bool   `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label               *string                                                `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze *bool                                                  `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Link                *string                                                `json:"link,omitempty" xml:"link,omitempty"`
	NeedDetail          *string                                                `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	NotPrint            *string                                                `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	NotUpper            *string                                                `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options             []*AddRelationMetaFieldRequestFieldDTOListPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable           *bool                                                  `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder         *string                                                `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Required               *bool   `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool   `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Sortable               *bool   `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Unit                   *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s AddRelationMetaFieldRequestFieldDTOListProps) String() string {
	return tea.Prettify(s)
}

func (s AddRelationMetaFieldRequestFieldDTOListProps) GoString() string {
	return s.String()
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetAlign(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Align = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetBizAlias(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.BizAlias = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetChoice(v int64) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Choice = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetContent(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Content = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetDisabled(v bool) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Disabled = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetDuration(v bool) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Duration = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetFieldId(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.FieldId = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetFormat(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Format = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetInvisible(v bool) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Invisible = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetLabel(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Label = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetLabelEditableFreeze(v bool) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetLink(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Link = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetNeedDetail(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.NeedDetail = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetNotPrint(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.NotPrint = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetNotUpper(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.NotUpper = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetOptions(v []*AddRelationMetaFieldRequestFieldDTOListPropsOptions) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Options = v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetPayEnable(v bool) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.PayEnable = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetPlaceholder(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Placeholder = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetRequired(v bool) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Required = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetRequiredEditableFreeze(v bool) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetSortable(v bool) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Sortable = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListProps) SetUnit(v string) *AddRelationMetaFieldRequestFieldDTOListProps {
	s.Unit = &v
	return s
}

type AddRelationMetaFieldRequestFieldDTOListPropsOptions struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddRelationMetaFieldRequestFieldDTOListPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s AddRelationMetaFieldRequestFieldDTOListPropsOptions) GoString() string {
	return s.String()
}

func (s *AddRelationMetaFieldRequestFieldDTOListPropsOptions) SetKey(v string) *AddRelationMetaFieldRequestFieldDTOListPropsOptions {
	s.Key = &v
	return s
}

func (s *AddRelationMetaFieldRequestFieldDTOListPropsOptions) SetValue(v string) *AddRelationMetaFieldRequestFieldDTOListPropsOptions {
	s.Value = &v
	return s
}

type AddRelationMetaFieldResponseBody struct {
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s AddRelationMetaFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRelationMetaFieldResponseBody) GoString() string {
	return s.String()
}

func (s *AddRelationMetaFieldResponseBody) SetRelationType(v string) *AddRelationMetaFieldResponseBody {
	s.RelationType = &v
	return s
}

type AddRelationMetaFieldResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRelationMetaFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRelationMetaFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRelationMetaFieldResponse) GoString() string {
	return s.String()
}

func (s *AddRelationMetaFieldResponse) SetHeaders(v map[string]*string) *AddRelationMetaFieldResponse {
	s.Headers = v
	return s
}

func (s *AddRelationMetaFieldResponse) SetStatusCode(v int32) *AddRelationMetaFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRelationMetaFieldResponse) SetBody(v *AddRelationMetaFieldResponseBody) *AddRelationMetaFieldResponse {
	s.Body = v
	return s
}

type AppendCustomerDataAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AppendCustomerDataAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppendCustomerDataAuthHeaders) GoString() string {
	return s.String()
}

func (s *AppendCustomerDataAuthHeaders) SetCommonHeaders(v map[string]*string) *AppendCustomerDataAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppendCustomerDataAuthHeaders) SetXAcsDingtalkAccessToken(v string) *AppendCustomerDataAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AppendCustomerDataAuthRequest struct {
	// This parameter is required.
	CustomerIds []*string `json:"customerIds,omitempty" xml:"customerIds,omitempty" type:"Repeated"`
	// This parameter is required.
	DataAuthUserIds []*string `json:"dataAuthUserIds,omitempty" xml:"dataAuthUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// PROC-98187D45-EFC0-4FC4-887E-45BD24209D69
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// staffId2
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	// example:
	//
	// crm_customer
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// owner
	RoleType *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s AppendCustomerDataAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s AppendCustomerDataAuthRequest) GoString() string {
	return s.String()
}

func (s *AppendCustomerDataAuthRequest) SetCustomerIds(v []*string) *AppendCustomerDataAuthRequest {
	s.CustomerIds = v
	return s
}

func (s *AppendCustomerDataAuthRequest) SetDataAuthUserIds(v []*string) *AppendCustomerDataAuthRequest {
	s.DataAuthUserIds = v
	return s
}

func (s *AppendCustomerDataAuthRequest) SetFormCode(v string) *AppendCustomerDataAuthRequest {
	s.FormCode = &v
	return s
}

func (s *AppendCustomerDataAuthRequest) SetOperateUserId(v string) *AppendCustomerDataAuthRequest {
	s.OperateUserId = &v
	return s
}

func (s *AppendCustomerDataAuthRequest) SetRelationType(v string) *AppendCustomerDataAuthRequest {
	s.RelationType = &v
	return s
}

func (s *AppendCustomerDataAuthRequest) SetRoleType(v string) *AppendCustomerDataAuthRequest {
	s.RoleType = &v
	return s
}

type AppendCustomerDataAuthResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AppendCustomerDataAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppendCustomerDataAuthResponseBody) GoString() string {
	return s.String()
}

func (s *AppendCustomerDataAuthResponseBody) SetResult(v bool) *AppendCustomerDataAuthResponseBody {
	s.Result = &v
	return s
}

type AppendCustomerDataAuthResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppendCustomerDataAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppendCustomerDataAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s AppendCustomerDataAuthResponse) GoString() string {
	return s.String()
}

func (s *AppendCustomerDataAuthResponse) SetHeaders(v map[string]*string) *AppendCustomerDataAuthResponse {
	s.Headers = v
	return s
}

func (s *AppendCustomerDataAuthResponse) SetStatusCode(v int32) *AppendCustomerDataAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *AppendCustomerDataAuthResponse) SetBody(v *AppendCustomerDataAuthResponseBody) *AppendCustomerDataAuthResponse {
	s.Body = v
	return s
}

type BatchAddContactsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchAddContactsHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchAddContactsHeaders) GoString() string {
	return s.String()
}

func (s *BatchAddContactsHeaders) SetCommonHeaders(v map[string]*string) *BatchAddContactsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchAddContactsHeaders) SetXAcsDingtalkAccessToken(v string) *BatchAddContactsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchAddContactsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manager021a
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	RelationList []*BatchAddContactsRequestRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
}

func (s BatchAddContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddContactsRequest) GoString() string {
	return s.String()
}

func (s *BatchAddContactsRequest) SetOperatorUserId(v string) *BatchAddContactsRequest {
	s.OperatorUserId = &v
	return s
}

func (s *BatchAddContactsRequest) SetRelationList(v []*BatchAddContactsRequestRelationList) *BatchAddContactsRequest {
	s.RelationList = v
	return s
}

type BatchAddContactsRequestRelationList struct {
	// This parameter is required.
	BizDataList []*BatchAddContactsRequestRelationListBizDataList `json:"bizDataList,omitempty" xml:"bizDataList,omitempty" type:"Repeated"`
	// if can be null:
	// true
	BizExtMap map[string]*string `json:"bizExtMap,omitempty" xml:"bizExtMap,omitempty"`
	// example:
	//
	// 1343434dd
	SourceDataId *string `json:"sourceDataId,omitempty" xml:"sourceDataId,omitempty"`
}

func (s BatchAddContactsRequestRelationList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddContactsRequestRelationList) GoString() string {
	return s.String()
}

func (s *BatchAddContactsRequestRelationList) SetBizDataList(v []*BatchAddContactsRequestRelationListBizDataList) *BatchAddContactsRequestRelationList {
	s.BizDataList = v
	return s
}

func (s *BatchAddContactsRequestRelationList) SetBizExtMap(v map[string]*string) *BatchAddContactsRequestRelationList {
	s.BizExtMap = v
	return s
}

func (s *BatchAddContactsRequestRelationList) SetSourceDataId(v string) *BatchAddContactsRequestRelationList {
	s.SourceDataId = &v
	return s
}

type BatchAddContactsRequestRelationListBizDataList struct {
	// example:
	//
	// {}
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField_71U51A
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX有限公司
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s BatchAddContactsRequestRelationListBizDataList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddContactsRequestRelationListBizDataList) GoString() string {
	return s.String()
}

func (s *BatchAddContactsRequestRelationListBizDataList) SetExtendValue(v string) *BatchAddContactsRequestRelationListBizDataList {
	s.ExtendValue = &v
	return s
}

func (s *BatchAddContactsRequestRelationListBizDataList) SetKey(v string) *BatchAddContactsRequestRelationListBizDataList {
	s.Key = &v
	return s
}

func (s *BatchAddContactsRequestRelationListBizDataList) SetValue(v string) *BatchAddContactsRequestRelationListBizDataList {
	s.Value = &v
	return s
}

type BatchAddContactsResponseBody struct {
	// example:
	//
	// true
	Results []*BatchAddContactsResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s BatchAddContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddContactsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddContactsResponseBody) SetResults(v []*BatchAddContactsResponseBodyResults) *BatchAddContactsResponseBody {
	s.Results = v
	return s
}

type BatchAddContactsResponseBodyResults struct {
	// example:
	//
	// 1002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 查重失败
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// gads1ag-sfgasdfxcvxb
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchAddContactsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchAddContactsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchAddContactsResponseBodyResults) SetErrorCode(v string) *BatchAddContactsResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *BatchAddContactsResponseBodyResults) SetErrorMsg(v string) *BatchAddContactsResponseBodyResults {
	s.ErrorMsg = &v
	return s
}

func (s *BatchAddContactsResponseBodyResults) SetRelationId(v string) *BatchAddContactsResponseBodyResults {
	s.RelationId = &v
	return s
}

func (s *BatchAddContactsResponseBodyResults) SetSuccess(v bool) *BatchAddContactsResponseBodyResults {
	s.Success = &v
	return s
}

type BatchAddContactsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddContactsResponse) GoString() string {
	return s.String()
}

func (s *BatchAddContactsResponse) SetHeaders(v map[string]*string) *BatchAddContactsResponse {
	s.Headers = v
	return s
}

func (s *BatchAddContactsResponse) SetStatusCode(v int32) *BatchAddContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddContactsResponse) SetBody(v *BatchAddContactsResponseBody) *BatchAddContactsResponse {
	s.Body = v
	return s
}

type BatchAddFollowRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchAddFollowRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFollowRecordsHeaders) GoString() string {
	return s.String()
}

func (s *BatchAddFollowRecordsHeaders) SetCommonHeaders(v map[string]*string) *BatchAddFollowRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchAddFollowRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *BatchAddFollowRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchAddFollowRecordsRequest struct {
	// This parameter is required.
	InstanceList []*BatchAddFollowRecordsRequestInstanceList `json:"instanceList,omitempty" xml:"instanceList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// manager021a
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s BatchAddFollowRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFollowRecordsRequest) GoString() string {
	return s.String()
}

func (s *BatchAddFollowRecordsRequest) SetInstanceList(v []*BatchAddFollowRecordsRequestInstanceList) *BatchAddFollowRecordsRequest {
	s.InstanceList = v
	return s
}

func (s *BatchAddFollowRecordsRequest) SetOperatorUserId(v string) *BatchAddFollowRecordsRequest {
	s.OperatorUserId = &v
	return s
}

type BatchAddFollowRecordsRequestInstanceList struct {
	BizExtMap map[string]*string `json:"bizExtMap,omitempty" xml:"bizExtMap,omitempty"`
	// This parameter is required.
	DataArray []*BatchAddFollowRecordsRequestInstanceListDataArray `json:"dataArray,omitempty" xml:"dataArray,omitempty" type:"Repeated"`
}

func (s BatchAddFollowRecordsRequestInstanceList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFollowRecordsRequestInstanceList) GoString() string {
	return s.String()
}

func (s *BatchAddFollowRecordsRequestInstanceList) SetBizExtMap(v map[string]*string) *BatchAddFollowRecordsRequestInstanceList {
	s.BizExtMap = v
	return s
}

func (s *BatchAddFollowRecordsRequestInstanceList) SetDataArray(v []*BatchAddFollowRecordsRequestInstanceListDataArray) *BatchAddFollowRecordsRequestInstanceList {
	s.DataArray = v
	return s
}

type BatchAddFollowRecordsRequestInstanceListDataArray struct {
	// example:
	//
	// {}
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField_71U51A
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX有限公司
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s BatchAddFollowRecordsRequestInstanceListDataArray) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFollowRecordsRequestInstanceListDataArray) GoString() string {
	return s.String()
}

func (s *BatchAddFollowRecordsRequestInstanceListDataArray) SetExtendValue(v string) *BatchAddFollowRecordsRequestInstanceListDataArray {
	s.ExtendValue = &v
	return s
}

func (s *BatchAddFollowRecordsRequestInstanceListDataArray) SetKey(v string) *BatchAddFollowRecordsRequestInstanceListDataArray {
	s.Key = &v
	return s
}

func (s *BatchAddFollowRecordsRequestInstanceListDataArray) SetValue(v string) *BatchAddFollowRecordsRequestInstanceListDataArray {
	s.Value = &v
	return s
}

type BatchAddFollowRecordsResponseBody struct {
	// example:
	//
	// true
	Results []*BatchAddFollowRecordsResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s BatchAddFollowRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFollowRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddFollowRecordsResponseBody) SetResults(v []*BatchAddFollowRecordsResponseBodyResults) *BatchAddFollowRecordsResponseBody {
	s.Results = v
	return s
}

type BatchAddFollowRecordsResponseBodyResults struct {
	// example:
	//
	// 1002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 查重失败
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// yU9TbER1TDazjPq1rRCzwg04841675924041
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchAddFollowRecordsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFollowRecordsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchAddFollowRecordsResponseBodyResults) SetErrorCode(v string) *BatchAddFollowRecordsResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *BatchAddFollowRecordsResponseBodyResults) SetErrorMsg(v string) *BatchAddFollowRecordsResponseBodyResults {
	s.ErrorMsg = &v
	return s
}

func (s *BatchAddFollowRecordsResponseBodyResults) SetInstanceId(v string) *BatchAddFollowRecordsResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *BatchAddFollowRecordsResponseBodyResults) SetSuccess(v bool) *BatchAddFollowRecordsResponseBodyResults {
	s.Success = &v
	return s
}

type BatchAddFollowRecordsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddFollowRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddFollowRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFollowRecordsResponse) GoString() string {
	return s.String()
}

func (s *BatchAddFollowRecordsResponse) SetHeaders(v map[string]*string) *BatchAddFollowRecordsResponse {
	s.Headers = v
	return s
}

func (s *BatchAddFollowRecordsResponse) SetStatusCode(v int32) *BatchAddFollowRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddFollowRecordsResponse) SetBody(v *BatchAddFollowRecordsResponseBody) *BatchAddFollowRecordsResponse {
	s.Body = v
	return s
}

type BatchAddRelationDatasHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchAddRelationDatasHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchAddRelationDatasHeaders) GoString() string {
	return s.String()
}

func (s *BatchAddRelationDatasHeaders) SetCommonHeaders(v map[string]*string) *BatchAddRelationDatasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchAddRelationDatasHeaders) SetXAcsDingtalkAccessToken(v string) *BatchAddRelationDatasHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchAddRelationDatasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manager021a
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	RelationList []*BatchAddRelationDatasRequestRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// crm_customer
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// example:
	//
	// false
	SkipDuplicateCheck *bool `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
}

func (s BatchAddRelationDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddRelationDatasRequest) GoString() string {
	return s.String()
}

func (s *BatchAddRelationDatasRequest) SetOperatorUserId(v string) *BatchAddRelationDatasRequest {
	s.OperatorUserId = &v
	return s
}

func (s *BatchAddRelationDatasRequest) SetRelationList(v []*BatchAddRelationDatasRequestRelationList) *BatchAddRelationDatasRequest {
	s.RelationList = v
	return s
}

func (s *BatchAddRelationDatasRequest) SetRelationType(v string) *BatchAddRelationDatasRequest {
	s.RelationType = &v
	return s
}

func (s *BatchAddRelationDatasRequest) SetSkipDuplicateCheck(v bool) *BatchAddRelationDatasRequest {
	s.SkipDuplicateCheck = &v
	return s
}

type BatchAddRelationDatasRequestRelationList struct {
	// This parameter is required.
	BizDataList []*BatchAddRelationDatasRequestRelationListBizDataList `json:"bizDataList,omitempty" xml:"bizDataList,omitempty" type:"Repeated"`
	// if can be null:
	// true
	BizExtMap             map[string]*string                                             `json:"bizExtMap,omitempty" xml:"bizExtMap,omitempty"`
	RelationPermissionDTO *BatchAddRelationDatasRequestRelationListRelationPermissionDTO `json:"relationPermissionDTO,omitempty" xml:"relationPermissionDTO,omitempty" type:"Struct"`
	// example:
	//
	// ddsf3234234
	SourceDataId *string `json:"sourceDataId,omitempty" xml:"sourceDataId,omitempty"`
}

func (s BatchAddRelationDatasRequestRelationList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddRelationDatasRequestRelationList) GoString() string {
	return s.String()
}

func (s *BatchAddRelationDatasRequestRelationList) SetBizDataList(v []*BatchAddRelationDatasRequestRelationListBizDataList) *BatchAddRelationDatasRequestRelationList {
	s.BizDataList = v
	return s
}

func (s *BatchAddRelationDatasRequestRelationList) SetBizExtMap(v map[string]*string) *BatchAddRelationDatasRequestRelationList {
	s.BizExtMap = v
	return s
}

func (s *BatchAddRelationDatasRequestRelationList) SetRelationPermissionDTO(v *BatchAddRelationDatasRequestRelationListRelationPermissionDTO) *BatchAddRelationDatasRequestRelationList {
	s.RelationPermissionDTO = v
	return s
}

func (s *BatchAddRelationDatasRequestRelationList) SetSourceDataId(v string) *BatchAddRelationDatasRequestRelationList {
	s.SourceDataId = &v
	return s
}

type BatchAddRelationDatasRequestRelationListBizDataList struct {
	// example:
	//
	// {}
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField_71U51A
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX有限公司
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s BatchAddRelationDatasRequestRelationListBizDataList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddRelationDatasRequestRelationListBizDataList) GoString() string {
	return s.String()
}

func (s *BatchAddRelationDatasRequestRelationListBizDataList) SetExtendValue(v string) *BatchAddRelationDatasRequestRelationListBizDataList {
	s.ExtendValue = &v
	return s
}

func (s *BatchAddRelationDatasRequestRelationListBizDataList) SetKey(v string) *BatchAddRelationDatasRequestRelationListBizDataList {
	s.Key = &v
	return s
}

func (s *BatchAddRelationDatasRequestRelationListBizDataList) SetValue(v string) *BatchAddRelationDatasRequestRelationListBizDataList {
	s.Value = &v
	return s
}

type BatchAddRelationDatasRequestRelationListRelationPermissionDTO struct {
	ParticipantUserIds []*string `json:"participantUserIds,omitempty" xml:"participantUserIds,omitempty" type:"Repeated"`
	PrincipalUserIds   []*string `json:"principalUserIds,omitempty" xml:"principalUserIds,omitempty" type:"Repeated"`
}

func (s BatchAddRelationDatasRequestRelationListRelationPermissionDTO) String() string {
	return tea.Prettify(s)
}

func (s BatchAddRelationDatasRequestRelationListRelationPermissionDTO) GoString() string {
	return s.String()
}

func (s *BatchAddRelationDatasRequestRelationListRelationPermissionDTO) SetParticipantUserIds(v []*string) *BatchAddRelationDatasRequestRelationListRelationPermissionDTO {
	s.ParticipantUserIds = v
	return s
}

func (s *BatchAddRelationDatasRequestRelationListRelationPermissionDTO) SetPrincipalUserIds(v []*string) *BatchAddRelationDatasRequestRelationListRelationPermissionDTO {
	s.PrincipalUserIds = v
	return s
}

type BatchAddRelationDatasResponseBody struct {
	// example:
	//
	// true
	Results []*BatchAddRelationDatasResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s BatchAddRelationDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddRelationDatasResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddRelationDatasResponseBody) SetResults(v []*BatchAddRelationDatasResponseBodyResults) *BatchAddRelationDatasResponseBody {
	s.Results = v
	return s
}

type BatchAddRelationDatasResponseBodyResults struct {
	DuplicatedRelationIds []*string `json:"duplicatedRelationIds,omitempty" xml:"duplicatedRelationIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 查重失败
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// gads1ag-sfgasdfxcvxb
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchAddRelationDatasResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchAddRelationDatasResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchAddRelationDatasResponseBodyResults) SetDuplicatedRelationIds(v []*string) *BatchAddRelationDatasResponseBodyResults {
	s.DuplicatedRelationIds = v
	return s
}

func (s *BatchAddRelationDatasResponseBodyResults) SetErrorCode(v string) *BatchAddRelationDatasResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *BatchAddRelationDatasResponseBodyResults) SetErrorMsg(v string) *BatchAddRelationDatasResponseBodyResults {
	s.ErrorMsg = &v
	return s
}

func (s *BatchAddRelationDatasResponseBodyResults) SetRelationId(v string) *BatchAddRelationDatasResponseBodyResults {
	s.RelationId = &v
	return s
}

func (s *BatchAddRelationDatasResponseBodyResults) SetSuccess(v bool) *BatchAddRelationDatasResponseBodyResults {
	s.Success = &v
	return s
}

type BatchAddRelationDatasResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddRelationDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddRelationDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddRelationDatasResponse) GoString() string {
	return s.String()
}

func (s *BatchAddRelationDatasResponse) SetHeaders(v map[string]*string) *BatchAddRelationDatasResponse {
	s.Headers = v
	return s
}

func (s *BatchAddRelationDatasResponse) SetStatusCode(v int32) *BatchAddRelationDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddRelationDatasResponse) SetBody(v *BatchAddRelationDatasResponseBody) *BatchAddRelationDatasResponse {
	s.Body = v
	return s
}

type BatchCreateClueDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchCreateClueDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateClueDataHeaders) GoString() string {
	return s.String()
}

func (s *BatchCreateClueDataHeaders) SetCommonHeaders(v map[string]*string) *BatchCreateClueDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchCreateClueDataHeaders) SetXAcsDingtalkAccessToken(v string) *BatchCreateClueDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchCreateClueDataRequest struct {
	DataList    []*BatchCreateClueDataRequestDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	PrivateSeas *bool                                 `json:"privateSeas,omitempty" xml:"privateSeas,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d124
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchCreateClueDataRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateClueDataRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateClueDataRequest) SetDataList(v []*BatchCreateClueDataRequestDataList) *BatchCreateClueDataRequest {
	s.DataList = v
	return s
}

func (s *BatchCreateClueDataRequest) SetPrivateSeas(v bool) *BatchCreateClueDataRequest {
	s.PrivateSeas = &v
	return s
}

func (s *BatchCreateClueDataRequest) SetUserId(v string) *BatchCreateClueDataRequest {
	s.UserId = &v
	return s
}

type BatchCreateClueDataRequestDataList struct {
	// This parameter is required.
	ContactList []*BatchCreateClueDataRequestDataListContactList `json:"contactList,omitempty" xml:"contactList,omitempty" type:"Repeated"`
	Enterprise  *BatchCreateClueDataRequestDataListEnterprise    `json:"enterprise,omitempty" xml:"enterprise,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 钉钉中国
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// demo
	SourceType *string                                        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	TagIdList  []*BatchCreateClueDataRequestDataListTagIdList `json:"tagIdList,omitempty" xml:"tagIdList,omitempty" type:"Repeated"`
}

func (s BatchCreateClueDataRequestDataList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateClueDataRequestDataList) GoString() string {
	return s.String()
}

func (s *BatchCreateClueDataRequestDataList) SetContactList(v []*BatchCreateClueDataRequestDataListContactList) *BatchCreateClueDataRequestDataList {
	s.ContactList = v
	return s
}

func (s *BatchCreateClueDataRequestDataList) SetEnterprise(v *BatchCreateClueDataRequestDataListEnterprise) *BatchCreateClueDataRequestDataList {
	s.Enterprise = v
	return s
}

func (s *BatchCreateClueDataRequestDataList) SetName(v string) *BatchCreateClueDataRequestDataList {
	s.Name = &v
	return s
}

func (s *BatchCreateClueDataRequestDataList) SetSourceId(v string) *BatchCreateClueDataRequestDataList {
	s.SourceId = &v
	return s
}

func (s *BatchCreateClueDataRequestDataList) SetSourceType(v string) *BatchCreateClueDataRequestDataList {
	s.SourceType = &v
	return s
}

func (s *BatchCreateClueDataRequestDataList) SetTagIdList(v []*BatchCreateClueDataRequestDataListTagIdList) *BatchCreateClueDataRequestDataList {
	s.TagIdList = v
	return s
}

type BatchCreateClueDataRequestDataListContactList struct {
	// This parameter is required.
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// This parameter is required.
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone    *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Qq       *string `json:"qq,omitempty" xml:"qq,omitempty"`
	WangWang *string `json:"wangWang,omitempty" xml:"wangWang,omitempty"`
	WeChat   *string `json:"weChat,omitempty" xml:"weChat,omitempty"`
}

func (s BatchCreateClueDataRequestDataListContactList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateClueDataRequestDataListContactList) GoString() string {
	return s.String()
}

func (s *BatchCreateClueDataRequestDataListContactList) SetMobile(v string) *BatchCreateClueDataRequestDataListContactList {
	s.Mobile = &v
	return s
}

func (s *BatchCreateClueDataRequestDataListContactList) SetName(v string) *BatchCreateClueDataRequestDataListContactList {
	s.Name = &v
	return s
}

func (s *BatchCreateClueDataRequestDataListContactList) SetPhone(v string) *BatchCreateClueDataRequestDataListContactList {
	s.Phone = &v
	return s
}

func (s *BatchCreateClueDataRequestDataListContactList) SetQq(v string) *BatchCreateClueDataRequestDataListContactList {
	s.Qq = &v
	return s
}

func (s *BatchCreateClueDataRequestDataListContactList) SetWangWang(v string) *BatchCreateClueDataRequestDataListContactList {
	s.WangWang = &v
	return s
}

func (s *BatchCreateClueDataRequestDataListContactList) SetWeChat(v string) *BatchCreateClueDataRequestDataListContactList {
	s.WeChat = &v
	return s
}

type BatchCreateClueDataRequestDataListEnterprise struct {
	Address      *string `json:"address,omitempty" xml:"address,omitempty"`
	IndustryCode *string `json:"industryCode,omitempty" xml:"industryCode,omitempty"`
	// This parameter is required.
	Name                    *string `json:"name,omitempty" xml:"name,omitempty"`
	Region                  *string `json:"region,omitempty" xml:"region,omitempty"`
	UnifiedSocialCreditCode *string `json:"unifiedSocialCreditCode,omitempty" xml:"unifiedSocialCreditCode,omitempty"`
}

func (s BatchCreateClueDataRequestDataListEnterprise) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateClueDataRequestDataListEnterprise) GoString() string {
	return s.String()
}

func (s *BatchCreateClueDataRequestDataListEnterprise) SetAddress(v string) *BatchCreateClueDataRequestDataListEnterprise {
	s.Address = &v
	return s
}

func (s *BatchCreateClueDataRequestDataListEnterprise) SetIndustryCode(v string) *BatchCreateClueDataRequestDataListEnterprise {
	s.IndustryCode = &v
	return s
}

func (s *BatchCreateClueDataRequestDataListEnterprise) SetName(v string) *BatchCreateClueDataRequestDataListEnterprise {
	s.Name = &v
	return s
}

func (s *BatchCreateClueDataRequestDataListEnterprise) SetRegion(v string) *BatchCreateClueDataRequestDataListEnterprise {
	s.Region = &v
	return s
}

func (s *BatchCreateClueDataRequestDataListEnterprise) SetUnifiedSocialCreditCode(v string) *BatchCreateClueDataRequestDataListEnterprise {
	s.UnifiedSocialCreditCode = &v
	return s
}

type BatchCreateClueDataRequestDataListTagIdList struct {
	// This parameter is required.
	TagId   *string `json:"tagId,omitempty" xml:"tagId,omitempty"`
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s BatchCreateClueDataRequestDataListTagIdList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateClueDataRequestDataListTagIdList) GoString() string {
	return s.String()
}

func (s *BatchCreateClueDataRequestDataListTagIdList) SetTagId(v string) *BatchCreateClueDataRequestDataListTagIdList {
	s.TagId = &v
	return s
}

func (s *BatchCreateClueDataRequestDataListTagIdList) SetTagName(v string) *BatchCreateClueDataRequestDataListTagIdList {
	s.TagName = &v
	return s
}

type BatchCreateClueDataResponseBody struct {
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*BatchCreateClueDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s BatchCreateClueDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateClueDataResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateClueDataResponseBody) SetRequestId(v string) *BatchCreateClueDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateClueDataResponseBody) SetResult(v []*BatchCreateClueDataResponseBodyResult) *BatchCreateClueDataResponseBody {
	s.Result = v
	return s
}

type BatchCreateClueDataResponseBodyResult struct {
	ClueId     *string `json:"clueId,omitempty" xml:"clueId,omitempty"`
	DataId     *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	ResultCode *string `json:"resultCode,omitempty" xml:"resultCode,omitempty"`
}

func (s BatchCreateClueDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateClueDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchCreateClueDataResponseBodyResult) SetClueId(v string) *BatchCreateClueDataResponseBodyResult {
	s.ClueId = &v
	return s
}

func (s *BatchCreateClueDataResponseBodyResult) SetDataId(v string) *BatchCreateClueDataResponseBodyResult {
	s.DataId = &v
	return s
}

func (s *BatchCreateClueDataResponseBodyResult) SetResultCode(v string) *BatchCreateClueDataResponseBodyResult {
	s.ResultCode = &v
	return s
}

type BatchCreateClueDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateClueDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateClueDataResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateClueDataResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateClueDataResponse) SetHeaders(v map[string]*string) *BatchCreateClueDataResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateClueDataResponse) SetStatusCode(v int32) *BatchCreateClueDataResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateClueDataResponse) SetBody(v *BatchCreateClueDataResponseBody) *BatchCreateClueDataResponse {
	s.Body = v
	return s
}

type BatchRemoveFollowRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchRemoveFollowRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchRemoveFollowRecordsHeaders) GoString() string {
	return s.String()
}

func (s *BatchRemoveFollowRecordsHeaders) SetCommonHeaders(v map[string]*string) *BatchRemoveFollowRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRemoveFollowRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *BatchRemoveFollowRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchRemoveFollowRecordsRequest struct {
	// This parameter is required.
	InstanceIds []*string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// manager021a
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s BatchRemoveFollowRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRemoveFollowRecordsRequest) GoString() string {
	return s.String()
}

func (s *BatchRemoveFollowRecordsRequest) SetInstanceIds(v []*string) *BatchRemoveFollowRecordsRequest {
	s.InstanceIds = v
	return s
}

func (s *BatchRemoveFollowRecordsRequest) SetOperatorUserId(v string) *BatchRemoveFollowRecordsRequest {
	s.OperatorUserId = &v
	return s
}

type BatchRemoveFollowRecordsResponseBody struct {
	// example:
	//
	// true
	Results []*BatchRemoveFollowRecordsResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s BatchRemoveFollowRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRemoveFollowRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRemoveFollowRecordsResponseBody) SetResults(v []*BatchRemoveFollowRecordsResponseBodyResults) *BatchRemoveFollowRecordsResponseBody {
	s.Results = v
	return s
}

type BatchRemoveFollowRecordsResponseBodyResults struct {
	// example:
	//
	// 1002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 查重失败
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// yU9TbER1TDazjPq1rRCzwg04841675924041
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchRemoveFollowRecordsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchRemoveFollowRecordsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchRemoveFollowRecordsResponseBodyResults) SetErrorCode(v string) *BatchRemoveFollowRecordsResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *BatchRemoveFollowRecordsResponseBodyResults) SetErrorMsg(v string) *BatchRemoveFollowRecordsResponseBodyResults {
	s.ErrorMsg = &v
	return s
}

func (s *BatchRemoveFollowRecordsResponseBodyResults) SetInstanceId(v string) *BatchRemoveFollowRecordsResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *BatchRemoveFollowRecordsResponseBodyResults) SetSuccess(v bool) *BatchRemoveFollowRecordsResponseBodyResults {
	s.Success = &v
	return s
}

type BatchRemoveFollowRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRemoveFollowRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchRemoveFollowRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRemoveFollowRecordsResponse) GoString() string {
	return s.String()
}

func (s *BatchRemoveFollowRecordsResponse) SetHeaders(v map[string]*string) *BatchRemoveFollowRecordsResponse {
	s.Headers = v
	return s
}

func (s *BatchRemoveFollowRecordsResponse) SetStatusCode(v int32) *BatchRemoveFollowRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRemoveFollowRecordsResponse) SetBody(v *BatchRemoveFollowRecordsResponseBody) *BatchRemoveFollowRecordsResponse {
	s.Body = v
	return s
}

type BatchSendOfficialAccountOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *BatchSendOfficialAccountOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *BatchSendOfficialAccountOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// if can be null:
	// true
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	Detail *BatchSendOfficialAccountOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
}

func (s BatchSendOfficialAccountOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequest) SetAccountId(v string) *BatchSendOfficialAccountOTOMessageRequest {
	s.AccountId = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequest) SetBizId(v string) *BatchSendOfficialAccountOTOMessageRequest {
	s.BizId = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequest) SetDetail(v *BatchSendOfficialAccountOTOMessageRequestDetail) *BatchSendOfficialAccountOTOMessageRequest {
	s.Detail = v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetail struct {
	// if can be null:
	// false
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	// This parameter is required.
	MessageBody *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// text
	MsgType    *string   `json:"msgType,omitempty" xml:"msgType,omitempty"`
	SendToAll  *bool     `json:"sendToAll,omitempty" xml:"sendToAll,omitempty"`
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetail) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetBizRequestId(v string) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.BizRequestId = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetMessageBody(v *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.MessageBody = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetMsgType(v string) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.MsgType = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetSendToAll(v bool) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.SendToAll = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetUserIdList(v []*string) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.UserIdList = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetUuid(v string) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.Uuid = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBody struct {
	ActionCard *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard `json:"actionCard,omitempty" xml:"actionCard,omitempty" type:"Struct"`
	Link       *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink       `json:"link,omitempty" xml:"link,omitempty" type:"Struct"`
	Markdown   *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown   `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	Text       *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText       `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) SetActionCard(v *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.ActionCard = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) SetLink(v *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Link = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) SetMarkdown(v *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Markdown = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) SetText(v *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Text = v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard struct {
	ButtonList        []*BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList `json:"buttonList,omitempty" xml:"buttonList,omitempty" type:"Repeated"`
	ButtonOrientation *string                                                                           `json:"buttonOrientation,omitempty" xml:"buttonOrientation,omitempty"`
	Markdown          *string                                                                           `json:"markdown,omitempty" xml:"markdown,omitempty"`
	SingleTitle       *string                                                                           `json:"singleTitle,omitempty" xml:"singleTitle,omitempty"`
	SingleUrl         *string                                                                           `json:"singleUrl,omitempty" xml:"singleUrl,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetButtonList(v []*BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.ButtonList = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetButtonOrientation(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.ButtonOrientation = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetMarkdown(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.Markdown = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetSingleTitle(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.SingleTitle = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetSingleUrl(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.SingleUrl = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetTitle(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.Title = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList struct {
	// This parameter is required.
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) SetActionUrl(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList {
	s.ActionUrl = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) SetTitle(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList {
	s.Title = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink struct {
	// This parameter is required.
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	// This parameter is required.
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetMessageUrl(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.MessageUrl = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetPicUrl(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.PicUrl = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetText(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.Text = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetTitle(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.Title = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown struct {
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) SetText(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown {
	s.Text = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) SetTitle(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown {
	s.Title = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText) SetContent(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText {
	s.Content = &v
	return s
}

type BatchSendOfficialAccountOTOMessageResponseBody struct {
	// example:
	//
	// acs1234
	RequestId *string                                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *BatchSendOfficialAccountOTOMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s BatchSendOfficialAccountOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageResponseBody) SetRequestId(v string) *BatchSendOfficialAccountOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageResponseBody) SetResult(v *BatchSendOfficialAccountOTOMessageResponseBodyResult) *BatchSendOfficialAccountOTOMessageResponseBody {
	s.Result = v
	return s
}

type BatchSendOfficialAccountOTOMessageResponseBodyResult struct {
	// This parameter is required.
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageResponseBodyResult) SetOpenPushId(v string) *BatchSendOfficialAccountOTOMessageResponseBodyResult {
	s.OpenPushId = &v
	return s
}

type BatchSendOfficialAccountOTOMessageResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSendOfficialAccountOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageResponse) SetHeaders(v map[string]*string) *BatchSendOfficialAccountOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageResponse) SetStatusCode(v int32) *BatchSendOfficialAccountOTOMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageResponse) SetBody(v *BatchSendOfficialAccountOTOMessageResponseBody) *BatchSendOfficialAccountOTOMessageResponse {
	s.Body = v
	return s
}

type BatchUpdateContactsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchUpdateContactsHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateContactsHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateContactsHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateContactsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateContactsHeaders) SetXAcsDingtalkAccessToken(v string) *BatchUpdateContactsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchUpdateContactsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manager021a
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	RelationList []*BatchUpdateContactsRequestRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
}

func (s BatchUpdateContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateContactsRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateContactsRequest) SetOperatorUserId(v string) *BatchUpdateContactsRequest {
	s.OperatorUserId = &v
	return s
}

func (s *BatchUpdateContactsRequest) SetRelationList(v []*BatchUpdateContactsRequestRelationList) *BatchUpdateContactsRequest {
	s.RelationList = v
	return s
}

type BatchUpdateContactsRequestRelationList struct {
	BizDataList []*BatchUpdateContactsRequestRelationListBizDataList `json:"bizDataList,omitempty" xml:"bizDataList,omitempty" type:"Repeated"`
	// if can be null:
	// true
	BizExtMap map[string]*string `json:"bizExtMap,omitempty" xml:"bizExtMap,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fasdg8i814-0afsd
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
}

func (s BatchUpdateContactsRequestRelationList) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateContactsRequestRelationList) GoString() string {
	return s.String()
}

func (s *BatchUpdateContactsRequestRelationList) SetBizDataList(v []*BatchUpdateContactsRequestRelationListBizDataList) *BatchUpdateContactsRequestRelationList {
	s.BizDataList = v
	return s
}

func (s *BatchUpdateContactsRequestRelationList) SetBizExtMap(v map[string]*string) *BatchUpdateContactsRequestRelationList {
	s.BizExtMap = v
	return s
}

func (s *BatchUpdateContactsRequestRelationList) SetRelationId(v string) *BatchUpdateContactsRequestRelationList {
	s.RelationId = &v
	return s
}

type BatchUpdateContactsRequestRelationListBizDataList struct {
	// example:
	//
	// {}
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField_71U51A
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX有限公司
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s BatchUpdateContactsRequestRelationListBizDataList) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateContactsRequestRelationListBizDataList) GoString() string {
	return s.String()
}

func (s *BatchUpdateContactsRequestRelationListBizDataList) SetExtendValue(v string) *BatchUpdateContactsRequestRelationListBizDataList {
	s.ExtendValue = &v
	return s
}

func (s *BatchUpdateContactsRequestRelationListBizDataList) SetKey(v string) *BatchUpdateContactsRequestRelationListBizDataList {
	s.Key = &v
	return s
}

func (s *BatchUpdateContactsRequestRelationListBizDataList) SetValue(v string) *BatchUpdateContactsRequestRelationListBizDataList {
	s.Value = &v
	return s
}

type BatchUpdateContactsResponseBody struct {
	// example:
	//
	// true
	Results []*BatchUpdateContactsResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s BatchUpdateContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateContactsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateContactsResponseBody) SetResults(v []*BatchUpdateContactsResponseBodyResults) *BatchUpdateContactsResponseBody {
	s.Results = v
	return s
}

type BatchUpdateContactsResponseBodyResults struct {
	// example:
	//
	// 1002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 查重失败
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// gads1ag-sfgasdfxcvxb
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchUpdateContactsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateContactsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUpdateContactsResponseBodyResults) SetErrorCode(v string) *BatchUpdateContactsResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *BatchUpdateContactsResponseBodyResults) SetErrorMsg(v string) *BatchUpdateContactsResponseBodyResults {
	s.ErrorMsg = &v
	return s
}

func (s *BatchUpdateContactsResponseBodyResults) SetRelationId(v string) *BatchUpdateContactsResponseBodyResults {
	s.RelationId = &v
	return s
}

func (s *BatchUpdateContactsResponseBodyResults) SetSuccess(v bool) *BatchUpdateContactsResponseBodyResults {
	s.Success = &v
	return s
}

type BatchUpdateContactsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateContactsResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateContactsResponse) SetHeaders(v map[string]*string) *BatchUpdateContactsResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateContactsResponse) SetStatusCode(v int32) *BatchUpdateContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateContactsResponse) SetBody(v *BatchUpdateContactsResponseBody) *BatchUpdateContactsResponse {
	s.Body = v
	return s
}

type BatchUpdateFollowRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchUpdateFollowRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFollowRecordsHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateFollowRecordsHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateFollowRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateFollowRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *BatchUpdateFollowRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchUpdateFollowRecordsRequest struct {
	// This parameter is required.
	InstanceList []*BatchUpdateFollowRecordsRequestInstanceList `json:"instanceList,omitempty" xml:"instanceList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// manager021a
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s BatchUpdateFollowRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFollowRecordsRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFollowRecordsRequest) SetInstanceList(v []*BatchUpdateFollowRecordsRequestInstanceList) *BatchUpdateFollowRecordsRequest {
	s.InstanceList = v
	return s
}

func (s *BatchUpdateFollowRecordsRequest) SetOperatorUserId(v string) *BatchUpdateFollowRecordsRequest {
	s.OperatorUserId = &v
	return s
}

type BatchUpdateFollowRecordsRequestInstanceList struct {
	// This parameter is required.
	DataArray []*BatchUpdateFollowRecordsRequestInstanceListDataArray `json:"dataArray,omitempty" xml:"dataArray,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// yU9TbER1TDazjPq1rRCzwg04841675924041
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s BatchUpdateFollowRecordsRequestInstanceList) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFollowRecordsRequestInstanceList) GoString() string {
	return s.String()
}

func (s *BatchUpdateFollowRecordsRequestInstanceList) SetDataArray(v []*BatchUpdateFollowRecordsRequestInstanceListDataArray) *BatchUpdateFollowRecordsRequestInstanceList {
	s.DataArray = v
	return s
}

func (s *BatchUpdateFollowRecordsRequestInstanceList) SetInstanceId(v string) *BatchUpdateFollowRecordsRequestInstanceList {
	s.InstanceId = &v
	return s
}

type BatchUpdateFollowRecordsRequestInstanceListDataArray struct {
	// example:
	//
	// {}
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField_71U51A
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX有限公司
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s BatchUpdateFollowRecordsRequestInstanceListDataArray) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFollowRecordsRequestInstanceListDataArray) GoString() string {
	return s.String()
}

func (s *BatchUpdateFollowRecordsRequestInstanceListDataArray) SetExtendValue(v string) *BatchUpdateFollowRecordsRequestInstanceListDataArray {
	s.ExtendValue = &v
	return s
}

func (s *BatchUpdateFollowRecordsRequestInstanceListDataArray) SetKey(v string) *BatchUpdateFollowRecordsRequestInstanceListDataArray {
	s.Key = &v
	return s
}

func (s *BatchUpdateFollowRecordsRequestInstanceListDataArray) SetValue(v string) *BatchUpdateFollowRecordsRequestInstanceListDataArray {
	s.Value = &v
	return s
}

type BatchUpdateFollowRecordsResponseBody struct {
	// example:
	//
	// true
	Results []*BatchUpdateFollowRecordsResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s BatchUpdateFollowRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFollowRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateFollowRecordsResponseBody) SetResults(v []*BatchUpdateFollowRecordsResponseBodyResults) *BatchUpdateFollowRecordsResponseBody {
	s.Results = v
	return s
}

type BatchUpdateFollowRecordsResponseBodyResults struct {
	// example:
	//
	// 1002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 查重失败
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// yU9TbER1TDazjPq1rRCzwg04841675924041
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchUpdateFollowRecordsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFollowRecordsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUpdateFollowRecordsResponseBodyResults) SetErrorCode(v string) *BatchUpdateFollowRecordsResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *BatchUpdateFollowRecordsResponseBodyResults) SetErrorMsg(v string) *BatchUpdateFollowRecordsResponseBodyResults {
	s.ErrorMsg = &v
	return s
}

func (s *BatchUpdateFollowRecordsResponseBodyResults) SetInstanceId(v string) *BatchUpdateFollowRecordsResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *BatchUpdateFollowRecordsResponseBodyResults) SetSuccess(v bool) *BatchUpdateFollowRecordsResponseBodyResults {
	s.Success = &v
	return s
}

type BatchUpdateFollowRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateFollowRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateFollowRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFollowRecordsResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateFollowRecordsResponse) SetHeaders(v map[string]*string) *BatchUpdateFollowRecordsResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateFollowRecordsResponse) SetStatusCode(v int32) *BatchUpdateFollowRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateFollowRecordsResponse) SetBody(v *BatchUpdateFollowRecordsResponseBody) *BatchUpdateFollowRecordsResponse {
	s.Body = v
	return s
}

type BatchUpdateRelationDatasHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchUpdateRelationDatasHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateRelationDatasHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateRelationDatasHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateRelationDatasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateRelationDatasHeaders) SetXAcsDingtalkAccessToken(v string) *BatchUpdateRelationDatasHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchUpdateRelationDatasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manager021a
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	RelationList []*BatchUpdateRelationDatasRequestRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// crm_customer
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// example:
	//
	// false
	SkipDuplicateCheck *bool `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
}

func (s BatchUpdateRelationDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateRelationDatasRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateRelationDatasRequest) SetOperatorUserId(v string) *BatchUpdateRelationDatasRequest {
	s.OperatorUserId = &v
	return s
}

func (s *BatchUpdateRelationDatasRequest) SetRelationList(v []*BatchUpdateRelationDatasRequestRelationList) *BatchUpdateRelationDatasRequest {
	s.RelationList = v
	return s
}

func (s *BatchUpdateRelationDatasRequest) SetRelationType(v string) *BatchUpdateRelationDatasRequest {
	s.RelationType = &v
	return s
}

func (s *BatchUpdateRelationDatasRequest) SetSkipDuplicateCheck(v bool) *BatchUpdateRelationDatasRequest {
	s.SkipDuplicateCheck = &v
	return s
}

type BatchUpdateRelationDatasRequestRelationList struct {
	BizDataList []*BatchUpdateRelationDatasRequestRelationListBizDataList `json:"bizDataList,omitempty" xml:"bizDataList,omitempty" type:"Repeated"`
	// if can be null:
	// true
	BizExtMap map[string]*string `json:"bizExtMap,omitempty" xml:"bizExtMap,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fasdg8i814-0afsd
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
}

func (s BatchUpdateRelationDatasRequestRelationList) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateRelationDatasRequestRelationList) GoString() string {
	return s.String()
}

func (s *BatchUpdateRelationDatasRequestRelationList) SetBizDataList(v []*BatchUpdateRelationDatasRequestRelationListBizDataList) *BatchUpdateRelationDatasRequestRelationList {
	s.BizDataList = v
	return s
}

func (s *BatchUpdateRelationDatasRequestRelationList) SetBizExtMap(v map[string]*string) *BatchUpdateRelationDatasRequestRelationList {
	s.BizExtMap = v
	return s
}

func (s *BatchUpdateRelationDatasRequestRelationList) SetRelationId(v string) *BatchUpdateRelationDatasRequestRelationList {
	s.RelationId = &v
	return s
}

type BatchUpdateRelationDatasRequestRelationListBizDataList struct {
	// example:
	//
	// {}
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField_71U51A
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX有限公司
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s BatchUpdateRelationDatasRequestRelationListBizDataList) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateRelationDatasRequestRelationListBizDataList) GoString() string {
	return s.String()
}

func (s *BatchUpdateRelationDatasRequestRelationListBizDataList) SetExtendValue(v string) *BatchUpdateRelationDatasRequestRelationListBizDataList {
	s.ExtendValue = &v
	return s
}

func (s *BatchUpdateRelationDatasRequestRelationListBizDataList) SetKey(v string) *BatchUpdateRelationDatasRequestRelationListBizDataList {
	s.Key = &v
	return s
}

func (s *BatchUpdateRelationDatasRequestRelationListBizDataList) SetValue(v string) *BatchUpdateRelationDatasRequestRelationListBizDataList {
	s.Value = &v
	return s
}

type BatchUpdateRelationDatasResponseBody struct {
	// example:
	//
	// true
	Results []*BatchUpdateRelationDatasResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s BatchUpdateRelationDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateRelationDatasResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateRelationDatasResponseBody) SetResults(v []*BatchUpdateRelationDatasResponseBodyResults) *BatchUpdateRelationDatasResponseBody {
	s.Results = v
	return s
}

type BatchUpdateRelationDatasResponseBodyResults struct {
	DuplicatedRelationIds []*string `json:"duplicatedRelationIds,omitempty" xml:"duplicatedRelationIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 查重失败
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// gads1ag-sfgasdfxcvxb
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchUpdateRelationDatasResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateRelationDatasResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUpdateRelationDatasResponseBodyResults) SetDuplicatedRelationIds(v []*string) *BatchUpdateRelationDatasResponseBodyResults {
	s.DuplicatedRelationIds = v
	return s
}

func (s *BatchUpdateRelationDatasResponseBodyResults) SetErrorCode(v string) *BatchUpdateRelationDatasResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *BatchUpdateRelationDatasResponseBodyResults) SetErrorMsg(v string) *BatchUpdateRelationDatasResponseBodyResults {
	s.ErrorMsg = &v
	return s
}

func (s *BatchUpdateRelationDatasResponseBodyResults) SetRelationId(v string) *BatchUpdateRelationDatasResponseBodyResults {
	s.RelationId = &v
	return s
}

func (s *BatchUpdateRelationDatasResponseBodyResults) SetSuccess(v bool) *BatchUpdateRelationDatasResponseBodyResults {
	s.Success = &v
	return s
}

type BatchUpdateRelationDatasResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateRelationDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateRelationDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateRelationDatasResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateRelationDatasResponse) SetHeaders(v map[string]*string) *BatchUpdateRelationDatasResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateRelationDatasResponse) SetStatusCode(v int32) *BatchUpdateRelationDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateRelationDatasResponse) SetBody(v *BatchUpdateRelationDatasResponseBody) *BatchUpdateRelationDatasResponse {
	s.Body = v
	return s
}

type ConsumeBenefitInventoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConsumeBenefitInventoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConsumeBenefitInventoryHeaders) GoString() string {
	return s.String()
}

func (s *ConsumeBenefitInventoryHeaders) SetCommonHeaders(v map[string]*string) *ConsumeBenefitInventoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConsumeBenefitInventoryHeaders) SetXAcsDingtalkAccessToken(v string) *ConsumeBenefitInventoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConsumeBenefitInventoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// B_ACCOUNT_NUMBER
	BenefitCode *string `json:"benefitCode,omitempty" xml:"benefitCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bizId
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	ConsumeQuota *int64 `json:"consumeQuota,omitempty" xml:"consumeQuota,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// optStaffId
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
}

func (s ConsumeBenefitInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsumeBenefitInventoryRequest) GoString() string {
	return s.String()
}

func (s *ConsumeBenefitInventoryRequest) SetBenefitCode(v string) *ConsumeBenefitInventoryRequest {
	s.BenefitCode = &v
	return s
}

func (s *ConsumeBenefitInventoryRequest) SetBizRequestId(v string) *ConsumeBenefitInventoryRequest {
	s.BizRequestId = &v
	return s
}

func (s *ConsumeBenefitInventoryRequest) SetConsumeQuota(v int64) *ConsumeBenefitInventoryRequest {
	s.ConsumeQuota = &v
	return s
}

func (s *ConsumeBenefitInventoryRequest) SetOptUserId(v string) *ConsumeBenefitInventoryRequest {
	s.OptUserId = &v
	return s
}

type ConsumeBenefitInventoryResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ConsumeBenefitInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConsumeBenefitInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *ConsumeBenefitInventoryResponseBody) SetResult(v bool) *ConsumeBenefitInventoryResponseBody {
	s.Result = &v
	return s
}

type ConsumeBenefitInventoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConsumeBenefitInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsumeBenefitInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsumeBenefitInventoryResponse) GoString() string {
	return s.String()
}

func (s *ConsumeBenefitInventoryResponse) SetHeaders(v map[string]*string) *ConsumeBenefitInventoryResponse {
	s.Headers = v
	return s
}

func (s *ConsumeBenefitInventoryResponse) SetStatusCode(v int32) *ConsumeBenefitInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsumeBenefitInventoryResponse) SetBody(v *ConsumeBenefitInventoryResponseBody) *ConsumeBenefitInventoryResponse {
	s.Body = v
	return s
}

type CreateCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerHeaders) GoString() string {
	return s.String()
}

func (s *CreateCustomerHeaders) SetCommonHeaders(v map[string]*string) *CreateCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCustomerRequest struct {
	Contacts []*CreateCustomerRequestContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// manager123
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// This parameter is required.
	Data       map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// crm_customer
	ObjectType *string                          `json:"objectType,omitempty" xml:"objectType,omitempty"`
	Permission *CreateCustomerRequestPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// This parameter is required.
	SaveOption *CreateCustomerRequestSaveOption `json:"saveOption,omitempty" xml:"saveOption,omitempty" type:"Struct"`
}

func (s CreateCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequest) SetContacts(v []*CreateCustomerRequestContacts) *CreateCustomerRequest {
	s.Contacts = v
	return s
}

func (s *CreateCustomerRequest) SetCreatorUserId(v string) *CreateCustomerRequest {
	s.CreatorUserId = &v
	return s
}

func (s *CreateCustomerRequest) SetData(v map[string]interface{}) *CreateCustomerRequest {
	s.Data = v
	return s
}

func (s *CreateCustomerRequest) SetExtendData(v map[string]interface{}) *CreateCustomerRequest {
	s.ExtendData = v
	return s
}

func (s *CreateCustomerRequest) SetInstanceId(v string) *CreateCustomerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomerRequest) SetObjectType(v string) *CreateCustomerRequest {
	s.ObjectType = &v
	return s
}

func (s *CreateCustomerRequest) SetPermission(v *CreateCustomerRequestPermission) *CreateCustomerRequest {
	s.Permission = v
	return s
}

func (s *CreateCustomerRequest) SetSaveOption(v *CreateCustomerRequestSaveOption) *CreateCustomerRequest {
	s.SaveOption = v
	return s
}

type CreateCustomerRequestContacts struct {
	// This parameter is required.
	Data       map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
}

func (s CreateCustomerRequestContacts) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequestContacts) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequestContacts) SetData(v map[string]interface{}) *CreateCustomerRequestContacts {
	s.Data = v
	return s
}

func (s *CreateCustomerRequestContacts) SetExtendData(v map[string]interface{}) *CreateCustomerRequestContacts {
	s.ExtendData = v
	return s
}

type CreateCustomerRequestPermission struct {
	OwnerStaffIds       []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s CreateCustomerRequestPermission) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequestPermission) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequestPermission) SetOwnerStaffIds(v []*string) *CreateCustomerRequestPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *CreateCustomerRequestPermission) SetParticipantStaffIds(v []*string) *CreateCustomerRequestPermission {
	s.ParticipantStaffIds = v
	return s
}

type CreateCustomerRequestSaveOption struct {
	// example:
	//
	// APPEND_CONTACT_FORCE
	CustomerExistedPolicy *string `json:"customerExistedPolicy,omitempty" xml:"customerExistedPolicy,omitempty"`
	// example:
	//
	// false
	SkipDuplicateCheck *bool `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
	// example:
	//
	// 0
	SubscribePolicy *int64 `json:"subscribePolicy,omitempty" xml:"subscribePolicy,omitempty"`
	// example:
	//
	// true
	ThrowExceptionWhileSavingContactFailed *bool `json:"throwExceptionWhileSavingContactFailed,omitempty" xml:"throwExceptionWhileSavingContactFailed,omitempty"`
}

func (s CreateCustomerRequestSaveOption) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequestSaveOption) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequestSaveOption) SetCustomerExistedPolicy(v string) *CreateCustomerRequestSaveOption {
	s.CustomerExistedPolicy = &v
	return s
}

func (s *CreateCustomerRequestSaveOption) SetSkipDuplicateCheck(v bool) *CreateCustomerRequestSaveOption {
	s.SkipDuplicateCheck = &v
	return s
}

func (s *CreateCustomerRequestSaveOption) SetSubscribePolicy(v int64) *CreateCustomerRequestSaveOption {
	s.SubscribePolicy = &v
	return s
}

func (s *CreateCustomerRequestSaveOption) SetThrowExceptionWhileSavingContactFailed(v bool) *CreateCustomerRequestSaveOption {
	s.ThrowExceptionWhileSavingContactFailed = &v
	return s
}

type CreateCustomerResponseBody struct {
	Contacts []*CreateCustomerResponseBodyContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// customer_id
	CustomerInstanceId *string `json:"customerInstanceId,omitempty" xml:"customerInstanceId,omitempty"`
	// example:
	//
	// crm_customer
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s CreateCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponseBody) SetContacts(v []*CreateCustomerResponseBodyContacts) *CreateCustomerResponseBody {
	s.Contacts = v
	return s
}

func (s *CreateCustomerResponseBody) SetCustomerInstanceId(v string) *CreateCustomerResponseBody {
	s.CustomerInstanceId = &v
	return s
}

func (s *CreateCustomerResponseBody) SetObjectType(v string) *CreateCustomerResponseBody {
	s.ObjectType = &v
	return s
}

type CreateCustomerResponseBodyContacts struct {
	// example:
	//
	// contact_id
	ContactInstanceId *string `json:"contactInstanceId,omitempty" xml:"contactInstanceId,omitempty"`
}

func (s CreateCustomerResponseBodyContacts) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponseBodyContacts) SetContactInstanceId(v string) *CreateCustomerResponseBodyContacts {
	s.ContactInstanceId = &v
	return s
}

type CreateCustomerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponse) SetHeaders(v map[string]*string) *CreateCustomerResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerResponse) SetStatusCode(v int32) *CreateCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomerResponse) SetBody(v *CreateCustomerResponseBody) *CreateCustomerResponse {
	s.Body = v
	return s
}

type CreateGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// a,b,c
	MemberUserIds *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc123
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) SetMemberUserIds(v string) *CreateGroupRequest {
	s.MemberUserIds = &v
	return s
}

func (s *CreateGroupRequest) SetOwnerUserId(v string) *CreateGroupRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateGroupRequest) SetRelationType(v string) *CreateGroupRequest {
	s.RelationType = &v
	return s
}

type CreateGroupResponseBody struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetOpenConversationId(v string) *CreateGroupResponseBody {
	s.OpenConversationId = &v
	return s
}

type CreateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateGroupSetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupSetHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupSetHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupSetHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupSetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupSetHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupSetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupSetRequest struct {
	// This parameter is required.
	CreatorUserId  *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	ManagerUserIds *string `json:"managerUserIds,omitempty" xml:"managerUserIds,omitempty"`
	MemberQuota    *int32  `json:"memberQuota,omitempty" xml:"memberQuota,omitempty"`
	// This parameter is required.
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Notice      *string `json:"notice,omitempty" xml:"notice,omitempty"`
	NoticeToped *int32  `json:"noticeToped,omitempty" xml:"noticeToped,omitempty"`
	// This parameter is required.
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	TemplateId   *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Welcome      *string `json:"welcome,omitempty" xml:"welcome,omitempty"`
}

func (s CreateGroupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupSetRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupSetRequest) SetCreatorUserId(v string) *CreateGroupSetRequest {
	s.CreatorUserId = &v
	return s
}

func (s *CreateGroupSetRequest) SetManagerUserIds(v string) *CreateGroupSetRequest {
	s.ManagerUserIds = &v
	return s
}

func (s *CreateGroupSetRequest) SetMemberQuota(v int32) *CreateGroupSetRequest {
	s.MemberQuota = &v
	return s
}

func (s *CreateGroupSetRequest) SetName(v string) *CreateGroupSetRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupSetRequest) SetNotice(v string) *CreateGroupSetRequest {
	s.Notice = &v
	return s
}

func (s *CreateGroupSetRequest) SetNoticeToped(v int32) *CreateGroupSetRequest {
	s.NoticeToped = &v
	return s
}

func (s *CreateGroupSetRequest) SetOwnerUserId(v string) *CreateGroupSetRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateGroupSetRequest) SetRelationType(v string) *CreateGroupSetRequest {
	s.RelationType = &v
	return s
}

func (s *CreateGroupSetRequest) SetTemplateId(v string) *CreateGroupSetRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateGroupSetRequest) SetWelcome(v string) *CreateGroupSetRequest {
	s.Welcome = &v
	return s
}

type CreateGroupSetResponseBody struct {
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InviteLink  *string `json:"inviteLink,omitempty" xml:"inviteLink,omitempty"`
	// This parameter is required.
	LastOpenConversationId *string `json:"lastOpenConversationId,omitempty" xml:"lastOpenConversationId,omitempty"`
	// This parameter is required.
	Manager        []*CreateGroupSetResponseBodyManager `json:"manager,omitempty" xml:"manager,omitempty" type:"Repeated"`
	ManagerUserIds *string                              `json:"managerUserIds,omitempty" xml:"managerUserIds,omitempty"`
	MemberCount    *int64                               `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	MemberQuota    *int64                               `json:"memberQuota,omitempty" xml:"memberQuota,omitempty"`
	Name           *string                              `json:"name,omitempty" xml:"name,omitempty"`
	Notice         *string                              `json:"notice,omitempty" xml:"notice,omitempty"`
	NoticeToped    *int32                               `json:"noticeToped,omitempty" xml:"noticeToped,omitempty"`
	OpenGroupSetId *string                              `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	Owner        *CreateGroupSetResponseBodyOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	OwnerUserId  *string                          `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	RelationType *string                          `json:"relationType,omitempty" xml:"relationType,omitempty"`
	TemplateId   *string                          `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CreateGroupSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupSetResponseBody) SetGmtCreate(v string) *CreateGroupSetResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetGmtModified(v string) *CreateGroupSetResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetInviteLink(v string) *CreateGroupSetResponseBody {
	s.InviteLink = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetLastOpenConversationId(v string) *CreateGroupSetResponseBody {
	s.LastOpenConversationId = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetManager(v []*CreateGroupSetResponseBodyManager) *CreateGroupSetResponseBody {
	s.Manager = v
	return s
}

func (s *CreateGroupSetResponseBody) SetManagerUserIds(v string) *CreateGroupSetResponseBody {
	s.ManagerUserIds = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetMemberCount(v int64) *CreateGroupSetResponseBody {
	s.MemberCount = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetMemberQuota(v int64) *CreateGroupSetResponseBody {
	s.MemberQuota = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetName(v string) *CreateGroupSetResponseBody {
	s.Name = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetNotice(v string) *CreateGroupSetResponseBody {
	s.Notice = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetNoticeToped(v int32) *CreateGroupSetResponseBody {
	s.NoticeToped = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetOpenGroupSetId(v string) *CreateGroupSetResponseBody {
	s.OpenGroupSetId = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetOwner(v *CreateGroupSetResponseBodyOwner) *CreateGroupSetResponseBody {
	s.Owner = v
	return s
}

func (s *CreateGroupSetResponseBody) SetOwnerUserId(v string) *CreateGroupSetResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetRelationType(v string) *CreateGroupSetResponseBody {
	s.RelationType = &v
	return s
}

func (s *CreateGroupSetResponseBody) SetTemplateId(v string) *CreateGroupSetResponseBody {
	s.TemplateId = &v
	return s
}

type CreateGroupSetResponseBodyManager struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateGroupSetResponseBodyManager) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupSetResponseBodyManager) GoString() string {
	return s.String()
}

func (s *CreateGroupSetResponseBodyManager) SetName(v string) *CreateGroupSetResponseBodyManager {
	s.Name = &v
	return s
}

func (s *CreateGroupSetResponseBodyManager) SetUserId(v string) *CreateGroupSetResponseBodyManager {
	s.UserId = &v
	return s
}

type CreateGroupSetResponseBodyOwner struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateGroupSetResponseBodyOwner) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupSetResponseBodyOwner) GoString() string {
	return s.String()
}

func (s *CreateGroupSetResponseBodyOwner) SetName(v string) *CreateGroupSetResponseBodyOwner {
	s.Name = &v
	return s
}

func (s *CreateGroupSetResponseBodyOwner) SetUserId(v string) *CreateGroupSetResponseBodyOwner {
	s.UserId = &v
	return s
}

type CreateGroupSetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupSetResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupSetResponse) SetHeaders(v map[string]*string) *CreateGroupSetResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupSetResponse) SetStatusCode(v int32) *CreateGroupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupSetResponse) SetBody(v *CreateGroupSetResponseBody) *CreateGroupSetResponse {
	s.Body = v
	return s
}

type CreateRelationMetaHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateRelationMetaHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateRelationMetaHeaders) GoString() string {
	return s.String()
}

func (s *CreateRelationMetaHeaders) SetCommonHeaders(v map[string]*string) *CreateRelationMetaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateRelationMetaHeaders) SetXAcsDingtalkAccessToken(v string) *CreateRelationMetaHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateRelationMetaRequest struct {
	// This parameter is required.
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	RelationMetaDTO *CreateRelationMetaRequestRelationMetaDTO `json:"relationMetaDTO,omitempty" xml:"relationMetaDTO,omitempty" type:"Struct"`
	// This parameter is required.
	Tenant *string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

func (s CreateRelationMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRelationMetaRequest) GoString() string {
	return s.String()
}

func (s *CreateRelationMetaRequest) SetOperatorUserId(v string) *CreateRelationMetaRequest {
	s.OperatorUserId = &v
	return s
}

func (s *CreateRelationMetaRequest) SetRelationMetaDTO(v *CreateRelationMetaRequestRelationMetaDTO) *CreateRelationMetaRequest {
	s.RelationMetaDTO = v
	return s
}

func (s *CreateRelationMetaRequest) SetTenant(v string) *CreateRelationMetaRequest {
	s.Tenant = &v
	return s
}

type CreateRelationMetaRequestRelationMetaDTO struct {
	// This parameter is required.
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// This parameter is required.
	Items []*CreateRelationMetaRequestRelationMetaDTOItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s CreateRelationMetaRequestRelationMetaDTO) String() string {
	return tea.Prettify(s)
}

func (s CreateRelationMetaRequestRelationMetaDTO) GoString() string {
	return s.String()
}

func (s *CreateRelationMetaRequestRelationMetaDTO) SetDesc(v string) *CreateRelationMetaRequestRelationMetaDTO {
	s.Desc = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTO) SetItems(v []*CreateRelationMetaRequestRelationMetaDTOItems) *CreateRelationMetaRequestRelationMetaDTO {
	s.Items = v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTO) SetName(v string) *CreateRelationMetaRequestRelationMetaDTO {
	s.Name = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTO) SetRelationType(v string) *CreateRelationMetaRequestRelationMetaDTO {
	s.RelationType = &v
	return s
}

type CreateRelationMetaRequestRelationMetaDTOItems struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *CreateRelationMetaRequestRelationMetaDTOItemsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s CreateRelationMetaRequestRelationMetaDTOItems) String() string {
	return tea.Prettify(s)
}

func (s CreateRelationMetaRequestRelationMetaDTOItems) GoString() string {
	return s.String()
}

func (s *CreateRelationMetaRequestRelationMetaDTOItems) SetComponentName(v string) *CreateRelationMetaRequestRelationMetaDTOItems {
	s.ComponentName = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItems) SetProps(v *CreateRelationMetaRequestRelationMetaDTOItemsProps) *CreateRelationMetaRequestRelationMetaDTOItems {
	s.Props = v
	return s
}

type CreateRelationMetaRequestRelationMetaDTOItemsProps struct {
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias  *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice    *int64  `json:"choice,omitempty" xml:"choice,omitempty"`
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	Disabled  *bool   `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration  *bool   `json:"duration,omitempty" xml:"duration,omitempty"`
	FieldId   *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format    *string `json:"format,omitempty" xml:"format,omitempty"`
	Invisible *bool   `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label               *string                                                      `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze *bool                                                        `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Link                *string                                                      `json:"link,omitempty" xml:"link,omitempty"`
	NeedDetail          *string                                                      `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	NotPrint            *string                                                      `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	NotUpper            *string                                                      `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options             []*CreateRelationMetaRequestRelationMetaDTOItemsPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable           *bool                                                        `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder         *string                                                      `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Required               *bool   `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool   `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Sortable               *bool   `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Unit                   *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s CreateRelationMetaRequestRelationMetaDTOItemsProps) String() string {
	return tea.Prettify(s)
}

func (s CreateRelationMetaRequestRelationMetaDTOItemsProps) GoString() string {
	return s.String()
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetAlign(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Align = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetBizAlias(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.BizAlias = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetChoice(v int64) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Choice = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetContent(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Content = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetDisabled(v bool) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Disabled = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetDuration(v bool) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Duration = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetFieldId(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.FieldId = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetFormat(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Format = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetInvisible(v bool) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Invisible = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetLabel(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Label = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetLabelEditableFreeze(v bool) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetLink(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Link = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetNeedDetail(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.NeedDetail = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetNotPrint(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.NotPrint = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetNotUpper(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.NotUpper = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetOptions(v []*CreateRelationMetaRequestRelationMetaDTOItemsPropsOptions) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Options = v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetPayEnable(v bool) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.PayEnable = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetPlaceholder(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Placeholder = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetRequired(v bool) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Required = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetRequiredEditableFreeze(v bool) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetSortable(v bool) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Sortable = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsProps) SetUnit(v string) *CreateRelationMetaRequestRelationMetaDTOItemsProps {
	s.Unit = &v
	return s
}

type CreateRelationMetaRequestRelationMetaDTOItemsPropsOptions struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateRelationMetaRequestRelationMetaDTOItemsPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateRelationMetaRequestRelationMetaDTOItemsPropsOptions) GoString() string {
	return s.String()
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsPropsOptions) SetKey(v string) *CreateRelationMetaRequestRelationMetaDTOItemsPropsOptions {
	s.Key = &v
	return s
}

func (s *CreateRelationMetaRequestRelationMetaDTOItemsPropsOptions) SetValue(v string) *CreateRelationMetaRequestRelationMetaDTOItemsPropsOptions {
	s.Value = &v
	return s
}

type CreateRelationMetaResponseBody struct {
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s CreateRelationMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRelationMetaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRelationMetaResponseBody) SetRelationType(v string) *CreateRelationMetaResponseBody {
	s.RelationType = &v
	return s
}

type CreateRelationMetaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRelationMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRelationMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRelationMetaResponse) GoString() string {
	return s.String()
}

func (s *CreateRelationMetaResponse) SetHeaders(v map[string]*string) *CreateRelationMetaResponse {
	s.Headers = v
	return s
}

func (s *CreateRelationMetaResponse) SetStatusCode(v int32) *CreateRelationMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRelationMetaResponse) SetBody(v *CreateRelationMetaResponseBody) *CreateRelationMetaResponse {
	s.Body = v
	return s
}

type DeleteCrmCustomObjectDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteCrmCustomObjectDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmCustomObjectDataHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCrmCustomObjectDataHeaders) SetCommonHeaders(v map[string]*string) *DeleteCrmCustomObjectDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCrmCustomObjectDataHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteCrmCustomObjectDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteCrmCustomObjectDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PROC_xx
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s DeleteCrmCustomObjectDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmCustomObjectDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteCrmCustomObjectDataRequest) SetFormCode(v string) *DeleteCrmCustomObjectDataRequest {
	s.FormCode = &v
	return s
}

type DeleteCrmCustomObjectDataResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// INST_xx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteCrmCustomObjectDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmCustomObjectDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCrmCustomObjectDataResponseBody) SetInstanceId(v string) *DeleteCrmCustomObjectDataResponseBody {
	s.InstanceId = &v
	return s
}

type DeleteCrmCustomObjectDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCrmCustomObjectDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCrmCustomObjectDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmCustomObjectDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteCrmCustomObjectDataResponse) SetHeaders(v map[string]*string) *DeleteCrmCustomObjectDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteCrmCustomObjectDataResponse) SetStatusCode(v int32) *DeleteCrmCustomObjectDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCrmCustomObjectDataResponse) SetBody(v *DeleteCrmCustomObjectDataResponseBody) *DeleteCrmCustomObjectDataResponse {
	s.Body = v
	return s
}

type DeleteCrmFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteCrmFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCrmFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *DeleteCrmFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCrmFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteCrmFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteCrmFormInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manager123
	CurrentOperatorUserId *string `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-123
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DeleteCrmFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteCrmFormInstanceRequest) SetCurrentOperatorUserId(v string) *DeleteCrmFormInstanceRequest {
	s.CurrentOperatorUserId = &v
	return s
}

func (s *DeleteCrmFormInstanceRequest) SetName(v string) *DeleteCrmFormInstanceRequest {
	s.Name = &v
	return s
}

type DeleteCrmFormInstanceResponseBody struct {
	// example:
	//
	// intanceId1
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteCrmFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCrmFormInstanceResponseBody) SetInstanceId(v string) *DeleteCrmFormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type DeleteCrmFormInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCrmFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCrmFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCrmFormInstanceResponse) SetHeaders(v map[string]*string) *DeleteCrmFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteCrmFormInstanceResponse) SetStatusCode(v int32) *DeleteCrmFormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCrmFormInstanceResponse) SetBody(v *DeleteCrmFormInstanceResponseBody) *DeleteCrmFormInstanceResponse {
	s.Body = v
	return s
}

type DeleteCrmPersonalCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteCrmPersonalCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmPersonalCustomerHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCrmPersonalCustomerHeaders) SetCommonHeaders(v map[string]*string) *DeleteCrmPersonalCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCrmPersonalCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteCrmPersonalCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteCrmPersonalCustomerRequest struct {
	// This parameter is required.
	CurrentOperatorUserId *string `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
	RelationType          *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s DeleteCrmPersonalCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmPersonalCustomerRequest) GoString() string {
	return s.String()
}

func (s *DeleteCrmPersonalCustomerRequest) SetCurrentOperatorUserId(v string) *DeleteCrmPersonalCustomerRequest {
	s.CurrentOperatorUserId = &v
	return s
}

func (s *DeleteCrmPersonalCustomerRequest) SetRelationType(v string) *DeleteCrmPersonalCustomerRequest {
	s.RelationType = &v
	return s
}

type DeleteCrmPersonalCustomerResponseBody struct {
	// This parameter is required.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteCrmPersonalCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmPersonalCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCrmPersonalCustomerResponseBody) SetInstanceId(v string) *DeleteCrmPersonalCustomerResponseBody {
	s.InstanceId = &v
	return s
}

type DeleteCrmPersonalCustomerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCrmPersonalCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmPersonalCustomerResponse) GoString() string {
	return s.String()
}

func (s *DeleteCrmPersonalCustomerResponse) SetHeaders(v map[string]*string) *DeleteCrmPersonalCustomerResponse {
	s.Headers = v
	return s
}

func (s *DeleteCrmPersonalCustomerResponse) SetStatusCode(v int32) *DeleteCrmPersonalCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCrmPersonalCustomerResponse) SetBody(v *DeleteCrmPersonalCustomerResponseBody) *DeleteCrmPersonalCustomerResponse {
	s.Body = v
	return s
}

type DeleteLeadsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteLeadsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteLeadsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteLeadsHeaders) SetCommonHeaders(v map[string]*string) *DeleteLeadsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteLeadsHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteLeadsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteLeadsRequest struct {
	// This parameter is required.
	OutLeadsIds []*string `json:"outLeadsIds,omitempty" xml:"outLeadsIds,omitempty" type:"Repeated"`
}

func (s DeleteLeadsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLeadsRequest) GoString() string {
	return s.String()
}

func (s *DeleteLeadsRequest) SetOutLeadsIds(v []*string) *DeleteLeadsRequest {
	s.OutLeadsIds = v
	return s
}

type DeleteLeadsResponseBody struct {
	OutLeadsIds []*string `json:"outLeadsIds,omitempty" xml:"outLeadsIds,omitempty" type:"Repeated"`
}

func (s DeleteLeadsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLeadsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLeadsResponseBody) SetOutLeadsIds(v []*string) *DeleteLeadsResponseBody {
	s.OutLeadsIds = v
	return s
}

type DeleteLeadsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLeadsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLeadsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLeadsResponse) GoString() string {
	return s.String()
}

func (s *DeleteLeadsResponse) SetHeaders(v map[string]*string) *DeleteLeadsResponse {
	s.Headers = v
	return s
}

func (s *DeleteLeadsResponse) SetStatusCode(v int32) *DeleteLeadsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLeadsResponse) SetBody(v *DeleteLeadsResponseBody) *DeleteLeadsResponse {
	s.Body = v
	return s
}

type DeleteRelationMetaFieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteRelationMetaFieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteRelationMetaFieldHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRelationMetaFieldHeaders) SetCommonHeaders(v map[string]*string) *DeleteRelationMetaFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRelationMetaFieldHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteRelationMetaFieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteRelationMetaFieldRequest struct {
	// This parameter is required.
	FieldIdList []*string `json:"fieldIdList,omitempty" xml:"fieldIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// This parameter is required.
	Tenant *string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

func (s DeleteRelationMetaFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRelationMetaFieldRequest) GoString() string {
	return s.String()
}

func (s *DeleteRelationMetaFieldRequest) SetFieldIdList(v []*string) *DeleteRelationMetaFieldRequest {
	s.FieldIdList = v
	return s
}

func (s *DeleteRelationMetaFieldRequest) SetOperatorUserId(v string) *DeleteRelationMetaFieldRequest {
	s.OperatorUserId = &v
	return s
}

func (s *DeleteRelationMetaFieldRequest) SetRelationType(v string) *DeleteRelationMetaFieldRequest {
	s.RelationType = &v
	return s
}

func (s *DeleteRelationMetaFieldRequest) SetTenant(v string) *DeleteRelationMetaFieldRequest {
	s.Tenant = &v
	return s
}

type DeleteRelationMetaFieldResponseBody struct {
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s DeleteRelationMetaFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRelationMetaFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRelationMetaFieldResponseBody) SetRelationType(v string) *DeleteRelationMetaFieldResponseBody {
	s.RelationType = &v
	return s
}

type DeleteRelationMetaFieldResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRelationMetaFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRelationMetaFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRelationMetaFieldResponse) GoString() string {
	return s.String()
}

func (s *DeleteRelationMetaFieldResponse) SetHeaders(v map[string]*string) *DeleteRelationMetaFieldResponse {
	s.Headers = v
	return s
}

func (s *DeleteRelationMetaFieldResponse) SetStatusCode(v int32) *DeleteRelationMetaFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRelationMetaFieldResponse) SetBody(v *DeleteRelationMetaFieldResponseBody) *DeleteRelationMetaFieldResponse {
	s.Body = v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaHeaders) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaHeaders) SetCommonHeaders(v map[string]*string) *DescribeCrmPersonalCustomerObjectMetaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaHeaders) SetXAcsDingtalkAccessToken(v string) *DescribeCrmPersonalCustomerObjectMetaHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaRequest struct {
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaRequest) SetRelationType(v string) *DescribeCrmPersonalCustomerObjectMetaRequest {
	s.RelationType = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBody struct {
	// example:
	//
	// PROC-XXXX
	Code       *string                                                    `json:"code,omitempty" xml:"code,omitempty"`
	Customized *bool                                                      `json:"customized,omitempty" xml:"customized,omitempty"`
	Fields     []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Name       *string                                                    `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// PUBLISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBody) SetCode(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBody) SetCustomized(v bool) *DescribeCrmPersonalCustomerObjectMetaResponseBody {
	s.Customized = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBody) SetFields(v []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) *DescribeCrmPersonalCustomerObjectMetaResponseBody {
	s.Fields = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBody) SetName(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBody) SetStatus(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBody {
	s.Status = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBodyFields struct {
	Customized          *bool                                                                         `json:"customized,omitempty" xml:"customized,omitempty"`
	Format              *string                                                                       `json:"format,omitempty" xml:"format,omitempty"`
	Label               *string                                                                       `json:"label,omitempty" xml:"label,omitempty"`
	Name                *string                                                                       `json:"name,omitempty" xml:"name,omitempty"`
	Nillable            *bool                                                                         `json:"nillable,omitempty" xml:"nillable,omitempty"`
	Quote               *bool                                                                         `json:"quote,omitempty" xml:"quote,omitempty"`
	ReferenceFields     []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields     `json:"referenceFields,omitempty" xml:"referenceFields,omitempty" type:"Repeated"`
	ReferenceTo         *string                                                                       `json:"referenceTo,omitempty" xml:"referenceTo,omitempty"`
	RollUpSummaryFields []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields `json:"rollUpSummaryFields,omitempty" xml:"rollUpSummaryFields,omitempty" type:"Repeated"`
	SelectOptions       []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions       `json:"selectOptions,omitempty" xml:"selectOptions,omitempty" type:"Repeated"`
	Type                *string                                                                       `json:"type,omitempty" xml:"type,omitempty"`
	Unit                *string                                                                       `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetCustomized(v bool) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Customized = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetFormat(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Format = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetLabel(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Label = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetName(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Name = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetNillable(v bool) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Nillable = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetQuote(v bool) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Quote = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetReferenceFields(v []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.ReferenceFields = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetReferenceTo(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.ReferenceTo = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetRollUpSummaryFields(v []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.RollUpSummaryFields = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetSelectOptions(v []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.SelectOptions = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetType(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Type = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetUnit(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Unit = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields struct {
	Format        *string                                                                                `json:"format,omitempty" xml:"format,omitempty"`
	Label         *string                                                                                `json:"label,omitempty" xml:"label,omitempty"`
	Name          *string                                                                                `json:"name,omitempty" xml:"name,omitempty"`
	Nillable      *bool                                                                                  `json:"nillable,omitempty" xml:"nillable,omitempty"`
	SelectOptions []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions `json:"selectOptions,omitempty" xml:"selectOptions,omitempty" type:"Repeated"`
	Type          *string                                                                                `json:"type,omitempty" xml:"type,omitempty"`
	Unit          *string                                                                                `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetFormat(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Format = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetLabel(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Label = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetName(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Name = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetNillable(v bool) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Nillable = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetSelectOptions(v []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.SelectOptions = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetType(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Type = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetUnit(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Unit = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions) SetKey(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions {
	s.Key = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions) SetValue(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions {
	s.Value = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields struct {
	// This parameter is required.
	Aggregator *string `json:"aggregator,omitempty" xml:"aggregator,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields) SetAggregator(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields {
	s.Aggregator = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields) SetName(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields {
	s.Name = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions) SetKey(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions {
	s.Key = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions) SetValue(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions {
	s.Value = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCrmPersonalCustomerObjectMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponse) SetHeaders(v map[string]*string) *DescribeCrmPersonalCustomerObjectMetaResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponse) SetStatusCode(v int32) *DescribeCrmPersonalCustomerObjectMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponse) SetBody(v *DescribeCrmPersonalCustomerObjectMetaResponseBody) *DescribeCrmPersonalCustomerObjectMetaResponse {
	s.Body = v
	return s
}

type DescribeMetaModelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DescribeMetaModelHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelHeaders) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelHeaders) SetCommonHeaders(v map[string]*string) *DescribeMetaModelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DescribeMetaModelHeaders) SetXAcsDingtalkAccessToken(v string) *DescribeMetaModelHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DescribeMetaModelRequest struct {
	// This parameter is required.
	BizTypes []*string `json:"bizTypes,omitempty" xml:"bizTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	Tenant *string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

func (s DescribeMetaModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelRequest) SetBizTypes(v []*string) *DescribeMetaModelRequest {
	s.BizTypes = v
	return s
}

func (s *DescribeMetaModelRequest) SetOperatorUserId(v string) *DescribeMetaModelRequest {
	s.OperatorUserId = &v
	return s
}

func (s *DescribeMetaModelRequest) SetTenant(v string) *DescribeMetaModelRequest {
	s.Tenant = &v
	return s
}

type DescribeMetaModelResponseBody struct {
	// This parameter is required.
	MetaModelDTOList []*DescribeMetaModelResponseBodyMetaModelDTOList `json:"metaModelDTOList,omitempty" xml:"metaModelDTOList,omitempty" type:"Repeated"`
}

func (s DescribeMetaModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBody) SetMetaModelDTOList(v []*DescribeMetaModelResponseBodyMetaModelDTOList) *DescribeMetaModelResponseBody {
	s.MetaModelDTOList = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOList struct {
	// This parameter is required.
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 企业客户表
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// This parameter is required.
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	Items []*DescribeMetaModelResponseBodyMetaModelDTOListItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 企业客户
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	RelationMetaCode *string `json:"relationMetaCode,omitempty" xml:"relationMetaCode,omitempty"`
	// This parameter is required.
	RelationMetaStatus *string `json:"relationMetaStatus,omitempty" xml:"relationMetaStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crm_customer
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOList) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOList) SetCreatorUserId(v string) *DescribeMetaModelResponseBodyMetaModelDTOList {
	s.CreatorUserId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOList) SetDesc(v string) *DescribeMetaModelResponseBodyMetaModelDTOList {
	s.Desc = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOList) SetGmtCreate(v string) *DescribeMetaModelResponseBodyMetaModelDTOList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOList) SetGmtModified(v string) *DescribeMetaModelResponseBodyMetaModelDTOList {
	s.GmtModified = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOList) SetItems(v []*DescribeMetaModelResponseBodyMetaModelDTOListItems) *DescribeMetaModelResponseBodyMetaModelDTOList {
	s.Items = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOList) SetName(v string) *DescribeMetaModelResponseBodyMetaModelDTOList {
	s.Name = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOList) SetRelationMetaCode(v string) *DescribeMetaModelResponseBodyMetaModelDTOList {
	s.RelationMetaCode = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOList) SetRelationMetaStatus(v string) *DescribeMetaModelResponseBodyMetaModelDTOList {
	s.RelationMetaStatus = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOList) SetRelationType(v string) *DescribeMetaModelResponseBodyMetaModelDTOList {
	s.RelationType = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItems struct {
	// This parameter is required.
	Children []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItems) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItems) SetChildren(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildren) *DescribeMetaModelResponseBodyMetaModelDTOListItems {
	s.Children = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItems) SetComponentName(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItems {
	s.ComponentName = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItems) SetProps(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) *DescribeMetaModelResponseBodyMetaModelDTOListItems {
	s.Props = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildren struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildren) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildren) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildren) SetComponentName(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildren {
	s.ComponentName = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildren) SetProps(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildren {
	s.Props = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps struct {
	// This parameter is required.
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// This parameter is required.
	Align              *string                                                                              `json:"align,omitempty" xml:"align,omitempty"`
	AvailableTemplates []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsAvailableTemplates `json:"availableTemplates,omitempty" xml:"availableTemplates,omitempty" type:"Repeated"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content    *string                                                                    `json:"content,omitempty" xml:"content,omitempty"`
	DataSource *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// example:
	//
	// 标签字段 颜色属性，格式：#0089FF
	DefaultColor *string `json:"defaultColor,omitempty" xml:"defaultColor,omitempty"`
	// This parameter is required.
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Fields []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// This parameter is required.
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：支持多选，false：单选
	Multiple *bool `json:"multiple,omitempty" xml:"multiple,omitempty"`
	// This parameter is required.
	NotPrint *string `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	// This parameter is required.
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Quote *int64 `json:"quote,omitempty" xml:"quote,omitempty"`
	// This parameter is required.
	Ratio *int64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// This parameter is required.
	RelateSource []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource `json:"relateSource,omitempty" xml:"relateSource,omitempty" type:"Repeated"`
	// This parameter is required.
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	Rule []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRule `json:"rule,omitempty" xml:"rule,omitempty" type:"Repeated"`
	// This parameter is required.
	Sortable *bool `json:"sortable,omitempty" xml:"sortable,omitempty"`
	// This parameter is required.
	Spread *bool `json:"spread,omitempty" xml:"spread,omitempty"`
	// This parameter is required.
	StatField []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	TableViewMode *string `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// This parameter is required.
	Watermark *bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetActionName(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.ActionName = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetAlign(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Align = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetAvailableTemplates(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsAvailableTemplates) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.AvailableTemplates = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetBizAlias(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetChoice(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Choice = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetContent(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Content = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetDataSource(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.DataSource = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetDefaultColor(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.DefaultColor = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetDisabled(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Disabled = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetDuration(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Duration = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetDurationLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.DurationLabel = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetFields(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFields) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Fields = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetFormat(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Format = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetFormula(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Formula = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetInvisible(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Invisible = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetLabelEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetLimit(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Limit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetLink(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Link = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetMode(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Mode = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetMultiple(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Multiple = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetNotPrint(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.NotPrint = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetNotUpper(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetOptions(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Options = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetPayEnable(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetPlaceholder(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetQuote(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Quote = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetRatio(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Ratio = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetRelateSource(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.RelateSource = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetRequired(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Required = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetRequiredEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetRule(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRule) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Rule = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetSortable(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Sortable = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetSpread(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Spread = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetStatField(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.StatField = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetTableViewMode(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.TableViewMode = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetVerticalPrint(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.VerticalPrint = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps) SetWatermark(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenProps {
	s.Watermark = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsAvailableTemplates struct {
	// example:
	//
	// 审批模板id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 审批模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsAvailableTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsAvailableTemplates) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsAvailableTemplates) SetId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsAvailableTemplates {
	s.Id = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsAvailableTemplates) SetName(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsAvailableTemplates {
	s.Name = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource struct {
	// This parameter is required.
	Params *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// This parameter is required.
	Target *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource) SetParams(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParams) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource {
	s.Params = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource) SetTarget(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource {
	s.Target = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource) SetType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSource {
	s.Type = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParams struct {
	// This parameter is required.
	Filters []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParams) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParams) SetFilters(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParams {
	s.Filters = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters) SetFilterType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters {
	s.FilterType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters {
	s.Value = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters) SetValueType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceParamsFilters {
	s.ValueType = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget struct {
	// This parameter is required.
	AppType *int64 `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget) SetAppType(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget) SetAppUuid(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget) SetBizType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget) SetFormCode(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsDataSourceTarget {
	s.FormCode = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFields struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	RelateProps *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFields) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFields) SetComponentName(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFields {
	s.ComponentName = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFields) SetRelateProps(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFields {
	s.RelateProps = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps struct {
	// This parameter is required.
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// This parameter is required.
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Ratio *int64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// This parameter is required.
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	Spread *bool `json:"spread,omitempty" xml:"spread,omitempty"`
	// This parameter is required.
	StatField []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// This parameter is required.
	Watermark *bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetAlign(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Align = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetBizAlias(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetChoice(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Choice = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetContent(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Content = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetDisabled(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Disabled = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetDuration(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Duration = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetDurationLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.DurationLabel = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetFormat(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Format = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetFormula(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Formula = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetInvisible(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Invisible = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetLabelEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetLimit(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Limit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetLink(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Link = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetMode(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Mode = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetNotUpper(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetOptions(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Options = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetPayEnable(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetPlaceholder(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetRatio(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Ratio = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetRequired(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Required = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetRequiredEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetSpread(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Spread = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetStatField(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.StatField = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetVerticalPrint(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.VerticalPrint = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps) SetWatermark(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelateProps {
	s.Watermark = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions struct {
	// This parameter is required.
	Extension *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions) SetExtension(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions {
	s.Extension = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions) SetKey(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptions {
	s.Value = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension struct {
	// This parameter is required.
	EditFreeze *bool `json:"editFreeze,omitempty" xml:"editFreeze,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension) SetEditFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension {
	s.EditFreeze = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField) SetUpper(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsFieldsRelatePropsStatField {
	s.Upper = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions struct {
	// This parameter is required.
	Extension *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions) SetExtension(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptionsExtension) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions {
	s.Extension = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions) SetKey(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptions {
	s.Value = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptionsExtension struct {
	// This parameter is required.
	EditFreeze *bool `json:"editFreeze,omitempty" xml:"editFreeze,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptionsExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptionsExtension) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptionsExtension) SetEditFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsOptionsExtension {
	s.EditFreeze = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource struct {
	// This parameter is required.
	BizType    *string                                                                                `json:"bizType,omitempty" xml:"bizType,omitempty"`
	DataSource *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// This parameter is required.
	Fields []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource) SetBizType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource {
	s.BizType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource) SetDataSource(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource {
	s.DataSource = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource) SetFields(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFields) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSource {
	s.Fields = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource struct {
	// This parameter is required.
	Params *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// This parameter is required.
	Target *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource) SetParams(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParams) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource {
	s.Params = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource) SetTarget(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource {
	s.Target = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource) SetType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSource {
	s.Type = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParams struct {
	// This parameter is required.
	Filters []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParams) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParams) SetFilters(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParams {
	s.Filters = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) SetFilterType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters {
	s.FilterType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters {
	s.Value = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) SetValueType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters {
	s.ValueType = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget struct {
	// This parameter is required.
	AppType *int64 `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget) SetAppType(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget) SetAppUuid(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget) SetBizType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget) SetFormCode(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceDataSourceTarget {
	s.FormCode = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFields struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	RelateProps *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFields) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFields) SetComponentName(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFields {
	s.ComponentName = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFields) SetRelateProps(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFields {
	s.RelateProps = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps struct {
	// This parameter is required.
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1：多选，0：单选
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：自动
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DDDateField和DDDateRangeField
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：隐藏
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	Multi *int64 `json:"multi,omitempty" xml:"multi,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1:不需要大写, 空或者0:需要大写
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true：是
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Quote *int64 `json:"quote,omitempty" xml:"quote,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	StatField []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：纵向，false：横向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetAlign(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Align = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetBizAlias(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetChoice(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Choice = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetContent(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Content = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetDisabled(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Disabled = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetDuration(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Duration = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetFormat(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Format = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetFormula(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Formula = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetInvisible(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Invisible = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetLabelEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetLink(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Link = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetMulti(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Multi = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetNotUpper(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetOptions(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Options = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetPayEnable(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetPlaceholder(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetQuote(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Quote = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetRequired(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Required = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetRequiredEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetStatField(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.StatField = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetVerticalPrint(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.VerticalPrint = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions struct {
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions) SetKey(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions {
	s.Value = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) SetUpper(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField {
	s.Upper = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRule struct {
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRule) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRule) SetType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRule {
	s.Type = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRule) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsRule {
	s.Value = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField) SetUpper(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsChildrenPropsStatField {
	s.Upper = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsProps struct {
	// This parameter is required.
	//
	// example:
	//
	// 明细动作名称
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// top|middle|bottom
	Align              *string                                                                      `json:"align,omitempty" xml:"align,omitempty"`
	AvailableTemplates []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsAvailableTemplates `json:"availableTemplates,omitempty" xml:"availableTemplates,omitempty" type:"Repeated"`
	BehaviorLinkage    []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkage    `json:"behaviorLinkage,omitempty" xml:"behaviorLinkage,omitempty" type:"Repeated"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1：多选，0：单选
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content    *string                                                            `json:"content,omitempty" xml:"content,omitempty"`
	DataSource *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// example:
	//
	// 标签字段 颜色属性，格式：#0089FF
	DefaultColor *string `json:"defaultColor,omitempty" xml:"defaultColor,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：自动
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 日期区间控件，自动计算时长的标题
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Fields []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// DDDateField和DDDateRangeField
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：隐藏
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：不可修改
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 评分组件限制
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 电话控件模式 phone：仅手机，phone_tel： 手机和固话，tel：仅固话
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// This parameter is required.
	Multi *int64 `json:"multi,omitempty" xml:"multi,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：支持多选，false：单选
	Multiple *bool `json:"multiple,omitempty" xml:"multiple,omitempty"`
	// This parameter is required.
	NeedDetail *string `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1：不打印，0：打印
	NotPrint *string `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1:不需要大写, 空或者0:需要大写
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true：是
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 关联表单 1：引用，0：拷贝
	Quote *int64 `json:"quote,omitempty" xml:"quote,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 文本控件、选项控件等限制文本字数ratio
	Ratio *int64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// This parameter is required.
	RelateSource []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource `json:"relateSource,omitempty" xml:"relateSource,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true：必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：不可修改
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	Rule []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRule `json:"rule,omitempty" xml:"rule,omitempty" type:"Repeated"`
	// This parameter is required.
	Sortable *bool `json:"sortable,omitempty" xml:"sortable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 选项控件spread
	Spread *bool `json:"spread,omitempty" xml:"spread,omitempty"`
	// This parameter is required.
	StatField []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 明细填写方式 table：表格，list：列表
	TableViewMode *string `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：纵向，false：横向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 是否水印照片 true：是，false：否
	Watermark *bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetActionName(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.ActionName = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetAlign(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Align = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetAvailableTemplates(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsAvailableTemplates) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.AvailableTemplates = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetBehaviorLinkage(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkage) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.BehaviorLinkage = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetBizAlias(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetChoice(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Choice = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetContent(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Content = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetDataSource(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.DataSource = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetDefaultColor(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.DefaultColor = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetDisabled(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Disabled = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetDuration(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Duration = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetDurationLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.DurationLabel = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetFields(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFields) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Fields = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetFormat(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Format = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetFormula(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Formula = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetInvisible(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Invisible = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetLabelEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetLimit(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Limit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetLink(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Link = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetMode(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Mode = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetMulti(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Multi = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetMultiple(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Multiple = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetNeedDetail(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.NeedDetail = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetNotPrint(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.NotPrint = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetNotUpper(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetOptions(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Options = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetPayEnable(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetPlaceholder(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetQuote(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Quote = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetRatio(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Ratio = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetRelateSource(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.RelateSource = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetRequired(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Required = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetRequiredEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetRule(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRule) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Rule = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetSortable(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Sortable = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetSpread(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Spread = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetStatField(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.StatField = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetTableViewMode(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.TableViewMode = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetVerticalPrint(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.VerticalPrint = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps) SetWatermark(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsProps {
	s.Watermark = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsAvailableTemplates struct {
	// example:
	//
	// 审批模板id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 审批模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsAvailableTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsAvailableTemplates) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsAvailableTemplates) SetId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsAvailableTemplates {
	s.Id = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsAvailableTemplates) SetName(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsAvailableTemplates {
	s.Name = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkage struct {
	// example:
	//
	// option_0
	OptionKey *string                                                                          `json:"optionKey,omitempty" xml:"optionKey,omitempty"`
	Targets   []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkageTargets `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkage) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkage) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkage) SetOptionKey(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkage {
	s.OptionKey = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkage) SetTargets(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkageTargets) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkage {
	s.Targets = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkageTargets struct {
	// example:
	//
	// NORMAL
	Behavior *string `json:"behavior,omitempty" xml:"behavior,omitempty"`
	// example:
	//
	// TextField_1LTIYOR4XIF40
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkageTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkageTargets) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkageTargets) SetBehavior(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkageTargets {
	s.Behavior = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkageTargets) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsBehaviorLinkageTargets {
	s.FieldId = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource struct {
	// This parameter is required.
	Params *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	Target *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	Type   *string                                                                  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource) SetParams(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParams) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource {
	s.Params = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource) SetTarget(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource {
	s.Target = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource) SetType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSource {
	s.Type = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParams struct {
	// This parameter is required.
	Filters []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParams) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParams) SetFilters(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParams {
	s.Filters = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters) SetFilterType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters {
	s.FilterType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters {
	s.Value = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters) SetValueType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceParamsFilters {
	s.ValueType = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget struct {
	// example:
	//
	// 0：流程表单，1：纯表单
	AppType  *int64  `json:"appType,omitempty" xml:"appType,omitempty"`
	AppUuid  *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	BizType  *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget) SetAppType(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget) SetAppUuid(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget) SetBizType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget) SetFormCode(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsDataSourceTarget {
	s.FormCode = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFields struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	RelateProps *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFields) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFields) SetComponentName(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFields {
	s.ComponentName = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFields) SetRelateProps(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFields {
	s.RelateProps = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps struct {
	// This parameter is required.
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1：多选，0：单选
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：自动
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DDDateField和DDDateRangeField
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：隐藏
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1:不需要大写, 空或者0:需要大写
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true：是
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Ratio *int64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	Spread *bool `json:"spread,omitempty" xml:"spread,omitempty"`
	// This parameter is required.
	StatField []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：纵向，false：横向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// This parameter is required.
	Watermark *bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetAlign(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Align = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetBizAlias(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetChoice(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Choice = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetContent(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Content = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetDisabled(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Disabled = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetDuration(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Duration = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetDurationLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.DurationLabel = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetFormat(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Format = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetFormula(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Formula = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetInvisible(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Invisible = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetLabelEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetLimit(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Limit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetLink(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Link = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetMode(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Mode = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetNotUpper(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetOptions(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Options = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetPayEnable(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetPlaceholder(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetRatio(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Ratio = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetRequired(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Required = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetRequiredEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetSpread(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Spread = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetStatField(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.StatField = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetVerticalPrint(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.VerticalPrint = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps) SetWatermark(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelateProps {
	s.Watermark = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions struct {
	// This parameter is required.
	Extension *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions) SetExtension(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptionsExtension) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions {
	s.Extension = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions) SetKey(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptions {
	s.Value = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptionsExtension struct {
	// This parameter is required.
	EditFreeze *bool `json:"editFreeze,omitempty" xml:"editFreeze,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptionsExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptionsExtension) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptionsExtension) SetEditFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsOptionsExtension {
	s.EditFreeze = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField) SetUpper(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsFieldsRelatePropsStatField {
	s.Upper = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions struct {
	// This parameter is required.
	Extension *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	Warn *bool `json:"warn,omitempty" xml:"warn,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions) SetExtension(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptionsExtension) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions {
	s.Extension = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions) SetKey(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions {
	s.Value = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions) SetWarn(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptions {
	s.Warn = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptionsExtension struct {
	// This parameter is required.
	EditFreeze *bool `json:"editFreeze,omitempty" xml:"editFreeze,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptionsExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptionsExtension) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptionsExtension) SetEditFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsOptionsExtension {
	s.EditFreeze = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource struct {
	// This parameter is required.
	BizType    *string                                                                        `json:"bizType,omitempty" xml:"bizType,omitempty"`
	DataSource *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// This parameter is required.
	Fields []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource) SetBizType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource {
	s.BizType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource) SetDataSource(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource {
	s.DataSource = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource) SetFields(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFields) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSource {
	s.Fields = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource struct {
	// This parameter is required.
	Params *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// This parameter is required.
	Target *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource) SetParams(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParams) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource {
	s.Params = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource) SetTarget(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource {
	s.Target = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource) SetType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSource {
	s.Type = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParams struct {
	// This parameter is required.
	Filters []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParams) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParams) SetFilters(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParams {
	s.Filters = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters) SetFilterType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters {
	s.FilterType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters {
	s.Value = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters) SetValueType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceParamsFilters {
	s.ValueType = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget struct {
	// This parameter is required.
	AppType *int64 `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget) SetAppType(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget) SetAppUuid(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget) SetBizType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget) SetFormCode(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceDataSourceTarget {
	s.FormCode = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFields struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	RelateProps *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFields) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFields) SetComponentName(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFields {
	s.ComponentName = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFields) SetRelateProps(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFields {
	s.RelateProps = v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps struct {
	// This parameter is required.
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1：多选，0：单选
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：自动
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DDDateField和DDDateRangeField
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：隐藏
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	Multi *int64 `json:"multi,omitempty" xml:"multi,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1:不需要大写, 空或者0:需要大写
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true：是
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Quote *int64 `json:"quote,omitempty" xml:"quote,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	StatField []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：纵向，false：横向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetAlign(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Align = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetBizAlias(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetChoice(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Choice = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetContent(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Content = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetDisabled(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Disabled = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetDuration(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Duration = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetFormat(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Format = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetFormula(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Formula = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetInvisible(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Invisible = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetLabelEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetLink(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Link = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetMulti(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Multi = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetNotUpper(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetOptions(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Options = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetPayEnable(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetPlaceholder(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetQuote(v int64) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Quote = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetRequired(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Required = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetRequiredEditableFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetStatField(v []*DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.StatField = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps) SetVerticalPrint(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.VerticalPrint = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions struct {
	// This parameter is required.
	Extension *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) SetExtension(v *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions {
	s.Extension = v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) SetKey(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptions {
	s.Value = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension struct {
	// This parameter is required.
	EditFreeze *bool `json:"editFreeze,omitempty" xml:"editFreeze,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension) SetEditFreeze(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension {
	s.EditFreeze = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) SetUpper(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRelateSourceFieldsRelatePropsStatField {
	s.Upper = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRule struct {
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRule) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRule) SetType(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRule {
	s.Type = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRule) SetValue(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsRule {
	s.Value = &v
	return s
}

type DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField) SetFieldId(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField) SetLabel(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField) SetUnit(v string) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField) SetUpper(v bool) *DescribeMetaModelResponseBodyMetaModelDTOListItemsPropsStatField {
	s.Upper = &v
	return s
}

type DescribeMetaModelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetaModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetaModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaModelResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetaModelResponse) SetHeaders(v map[string]*string) *DescribeMetaModelResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetaModelResponse) SetStatusCode(v int32) *DescribeMetaModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetaModelResponse) SetBody(v *DescribeMetaModelResponseBody) *DescribeMetaModelResponse {
	s.Body = v
	return s
}

type DescribeRelationMetaHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DescribeRelationMetaHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaHeaders) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaHeaders) SetCommonHeaders(v map[string]*string) *DescribeRelationMetaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DescribeRelationMetaHeaders) SetXAcsDingtalkAccessToken(v string) *DescribeRelationMetaHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DescribeRelationMetaRequest struct {
	// This parameter is required.
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	RelationTypes []*string `json:"relationTypes,omitempty" xml:"relationTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	Tenant *string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

func (s DescribeRelationMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaRequest) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaRequest) SetOperatorUserId(v string) *DescribeRelationMetaRequest {
	s.OperatorUserId = &v
	return s
}

func (s *DescribeRelationMetaRequest) SetRelationTypes(v []*string) *DescribeRelationMetaRequest {
	s.RelationTypes = v
	return s
}

func (s *DescribeRelationMetaRequest) SetTenant(v string) *DescribeRelationMetaRequest {
	s.Tenant = &v
	return s
}

type DescribeRelationMetaResponseBody struct {
	// This parameter is required.
	RelationMetaDTOList []*DescribeRelationMetaResponseBodyRelationMetaDTOList `json:"relationMetaDTOList,omitempty" xml:"relationMetaDTOList,omitempty" type:"Repeated"`
}

func (s DescribeRelationMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBody) SetRelationMetaDTOList(v []*DescribeRelationMetaResponseBodyRelationMetaDTOList) *DescribeRelationMetaResponseBody {
	s.RelationMetaDTOList = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOList struct {
	// This parameter is required.
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 企业客户表
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// This parameter is required.
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	Items []*DescribeRelationMetaResponseBodyRelationMetaDTOListItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 企业客户
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	RelationMetaCode *string `json:"relationMetaCode,omitempty" xml:"relationMetaCode,omitempty"`
	// This parameter is required.
	RelationMetaStatus *string `json:"relationMetaStatus,omitempty" xml:"relationMetaStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crm_customer
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOList) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOList) SetCreatorUserId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOList {
	s.CreatorUserId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOList) SetDesc(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOList {
	s.Desc = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOList) SetGmtCreate(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOList) SetGmtModified(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOList {
	s.GmtModified = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOList) SetItems(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItems) *DescribeRelationMetaResponseBodyRelationMetaDTOList {
	s.Items = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOList) SetName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOList {
	s.Name = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOList) SetRelationMetaCode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOList {
	s.RelationMetaCode = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOList) SetRelationMetaStatus(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOList {
	s.RelationMetaStatus = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOList) SetRelationType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOList {
	s.RelationType = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItems struct {
	// This parameter is required.
	Children []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItems) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItems) SetChildren(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildren) *DescribeRelationMetaResponseBodyRelationMetaDTOListItems {
	s.Children = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItems) SetComponentName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItems {
	s.ComponentName = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItems) SetProps(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) *DescribeRelationMetaResponseBodyRelationMetaDTOListItems {
	s.Props = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildren struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildren) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildren) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildren) SetComponentName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildren {
	s.ComponentName = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildren) SetProps(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildren {
	s.Props = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps struct {
	// This parameter is required.
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// This parameter is required.
	Align              *string                                                                                    `json:"align,omitempty" xml:"align,omitempty"`
	AvailableTemplates []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsAvailableTemplates `json:"availableTemplates,omitempty" xml:"availableTemplates,omitempty" type:"Repeated"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content    *string                                                                          `json:"content,omitempty" xml:"content,omitempty"`
	DataSource *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// example:
	//
	// 标签字段 颜色属性，格式：#0089FF
	DefaultColor *string `json:"defaultColor,omitempty" xml:"defaultColor,omitempty"`
	// This parameter is required.
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Fields []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// This parameter is required.
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：支持多选，false：单选
	Multiple *bool `json:"multiple,omitempty" xml:"multiple,omitempty"`
	// This parameter is required.
	NotPrint *string `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	// This parameter is required.
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Quote *int64 `json:"quote,omitempty" xml:"quote,omitempty"`
	// This parameter is required.
	Ratio *int64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// This parameter is required.
	RelateSource []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource `json:"relateSource,omitempty" xml:"relateSource,omitempty" type:"Repeated"`
	// This parameter is required.
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	Rule []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRule `json:"rule,omitempty" xml:"rule,omitempty" type:"Repeated"`
	// This parameter is required.
	Sortable *bool `json:"sortable,omitempty" xml:"sortable,omitempty"`
	// This parameter is required.
	Spread *bool `json:"spread,omitempty" xml:"spread,omitempty"`
	// This parameter is required.
	StatField []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	TableViewMode *string `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// This parameter is required.
	Watermark *bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetActionName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.ActionName = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetAlign(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Align = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetAvailableTemplates(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsAvailableTemplates) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.AvailableTemplates = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetBizAlias(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetChoice(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Choice = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetContent(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Content = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetDataSource(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.DataSource = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetDefaultColor(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.DefaultColor = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetDisabled(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Disabled = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetDuration(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Duration = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetDurationLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.DurationLabel = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetFields(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFields) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Fields = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetFormat(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Format = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetFormula(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Formula = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetInvisible(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Invisible = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetLabelEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetLimit(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Limit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetLink(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Link = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetMode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Mode = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetMultiple(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Multiple = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetNotPrint(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.NotPrint = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetNotUpper(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetOptions(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Options = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetPayEnable(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetPlaceholder(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetQuote(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Quote = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetRatio(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Ratio = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetRelateSource(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.RelateSource = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetRequired(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Required = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetRequiredEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetRule(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRule) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Rule = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetSortable(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Sortable = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetSpread(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Spread = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetStatField(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.StatField = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetTableViewMode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.TableViewMode = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetVerticalPrint(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.VerticalPrint = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps) SetWatermark(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps {
	s.Watermark = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsAvailableTemplates struct {
	// example:
	//
	// 审批模板id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 审批模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsAvailableTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsAvailableTemplates) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsAvailableTemplates) SetId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsAvailableTemplates {
	s.Id = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsAvailableTemplates) SetName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsAvailableTemplates {
	s.Name = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource struct {
	// This parameter is required.
	Params *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// This parameter is required.
	Target *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource) SetParams(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParams) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource {
	s.Params = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource) SetTarget(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource {
	s.Target = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource) SetType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource {
	s.Type = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParams struct {
	// This parameter is required.
	Filters []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParams) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParams) SetFilters(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParams {
	s.Filters = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters) SetFilterType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters {
	s.FilterType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters {
	s.Value = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters) SetValueType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParamsFilters {
	s.ValueType = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget struct {
	// This parameter is required.
	AppType *int64 `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget) SetAppType(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget) SetAppUuid(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget) SetBizType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget) SetFormCode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget {
	s.FormCode = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFields struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	RelateProps *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFields) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFields) SetComponentName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFields {
	s.ComponentName = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFields) SetRelateProps(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFields {
	s.RelateProps = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps struct {
	// This parameter is required.
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// This parameter is required.
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Ratio *int64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// This parameter is required.
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	Spread *bool `json:"spread,omitempty" xml:"spread,omitempty"`
	// This parameter is required.
	StatField []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// This parameter is required.
	Watermark *bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetAlign(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Align = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetBizAlias(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetChoice(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Choice = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetContent(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Content = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetDisabled(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Disabled = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetDuration(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Duration = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetDurationLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.DurationLabel = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetFormat(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Format = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetFormula(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Formula = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetInvisible(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Invisible = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetLabelEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetLimit(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Limit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetLink(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Link = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetMode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Mode = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetNotUpper(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetOptions(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Options = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetPayEnable(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetPlaceholder(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetRatio(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Ratio = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetRequired(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Required = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetRequiredEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetSpread(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Spread = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetStatField(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.StatField = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetVerticalPrint(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.VerticalPrint = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps) SetWatermark(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps {
	s.Watermark = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions struct {
	// This parameter is required.
	Extension *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions) SetExtension(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions {
	s.Extension = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions) SetKey(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions {
	s.Value = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension struct {
	// This parameter is required.
	EditFreeze *bool `json:"editFreeze,omitempty" xml:"editFreeze,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension) SetEditFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension {
	s.EditFreeze = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField) SetUpper(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField {
	s.Upper = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions struct {
	// This parameter is required.
	Extension *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions) SetExtension(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptionsExtension) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions {
	s.Extension = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions) SetKey(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions {
	s.Value = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptionsExtension struct {
	// This parameter is required.
	EditFreeze *bool `json:"editFreeze,omitempty" xml:"editFreeze,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptionsExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptionsExtension) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptionsExtension) SetEditFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptionsExtension {
	s.EditFreeze = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource struct {
	// This parameter is required.
	BizType    *string                                                                                      `json:"bizType,omitempty" xml:"bizType,omitempty"`
	DataSource *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// This parameter is required.
	Fields []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource) SetBizType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource {
	s.BizType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource) SetDataSource(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource {
	s.DataSource = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource) SetFields(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFields) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource {
	s.Fields = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource struct {
	// This parameter is required.
	Params *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// This parameter is required.
	Target *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource) SetParams(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParams) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource {
	s.Params = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource) SetTarget(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource {
	s.Target = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource) SetType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource {
	s.Type = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParams struct {
	// This parameter is required.
	Filters []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParams) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParams) SetFilters(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParams {
	s.Filters = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) SetFilterType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters {
	s.FilterType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters {
	s.Value = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters) SetValueType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParamsFilters {
	s.ValueType = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget struct {
	// This parameter is required.
	AppType *int64 `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget) SetAppType(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget) SetAppUuid(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget) SetBizType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget) SetFormCode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget {
	s.FormCode = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFields struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	RelateProps *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFields) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFields) SetComponentName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFields {
	s.ComponentName = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFields) SetRelateProps(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFields {
	s.RelateProps = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps struct {
	// This parameter is required.
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1：多选，0：单选
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：自动
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DDDateField和DDDateRangeField
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：隐藏
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	Multi *int64 `json:"multi,omitempty" xml:"multi,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1:不需要大写, 空或者0:需要大写
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true：是
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Quote *int64 `json:"quote,omitempty" xml:"quote,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	StatField []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：纵向，false：横向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetAlign(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Align = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetBizAlias(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetChoice(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Choice = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetContent(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Content = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetDisabled(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Disabled = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetDuration(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Duration = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetFormat(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Format = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetFormula(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Formula = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetInvisible(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Invisible = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetLabelEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetLink(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Link = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetMulti(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Multi = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetNotUpper(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetOptions(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Options = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetPayEnable(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetPlaceholder(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetQuote(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Quote = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetRequired(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Required = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetRequiredEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetStatField(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.StatField = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps) SetVerticalPrint(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps {
	s.VerticalPrint = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions struct {
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions) SetKey(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions {
	s.Value = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField) SetUpper(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField {
	s.Upper = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRule struct {
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRule) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRule) SetType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRule {
	s.Type = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRule) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRule {
	s.Value = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField) SetUpper(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField {
	s.Upper = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps struct {
	// This parameter is required.
	//
	// example:
	//
	// 明细动作名称
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// top|middle|bottom
	Align              *string                                                                            `json:"align,omitempty" xml:"align,omitempty"`
	AvailableTemplates []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsAvailableTemplates `json:"availableTemplates,omitempty" xml:"availableTemplates,omitempty" type:"Repeated"`
	BehaviorLinkage    []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkage    `json:"behaviorLinkage,omitempty" xml:"behaviorLinkage,omitempty" type:"Repeated"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1：多选，0：单选
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content    *string                                                                  `json:"content,omitempty" xml:"content,omitempty"`
	DataSource *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// example:
	//
	// 标签字段 颜色属性，格式：#0089FF
	DefaultColor *string `json:"defaultColor,omitempty" xml:"defaultColor,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：自动
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 日期区间控件，自动计算时长的标题
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Fields []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// DDDateField和DDDateRangeField
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：隐藏
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：不可修改
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 评分组件限制
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 电话控件模式 phone：仅手机，phone_tel： 手机和固话，tel：仅固话
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// This parameter is required.
	Multi *int64 `json:"multi,omitempty" xml:"multi,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：支持多选，false：单选
	Multiple *bool `json:"multiple,omitempty" xml:"multiple,omitempty"`
	// This parameter is required.
	NeedDetail *string `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1：不打印，0：打印
	NotPrint *string `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1:不需要大写, 空或者0:需要大写
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true：是
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 关联表单 1：引用，0：拷贝
	Quote *int64 `json:"quote,omitempty" xml:"quote,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 文本控件、选项控件等限制文本字数ratio
	Ratio *int64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// This parameter is required.
	RelateSource []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource `json:"relateSource,omitempty" xml:"relateSource,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true：必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：不可修改
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	Rule []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRule `json:"rule,omitempty" xml:"rule,omitempty" type:"Repeated"`
	// This parameter is required.
	Sortable *bool `json:"sortable,omitempty" xml:"sortable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 选项控件spread
	Spread *bool `json:"spread,omitempty" xml:"spread,omitempty"`
	// This parameter is required.
	StatField []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 明细填写方式 table：表格，list：列表
	TableViewMode *string `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：纵向，false：横向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 是否水印照片 true：是，false：否
	Watermark *bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetActionName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.ActionName = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetAlign(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Align = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetAvailableTemplates(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsAvailableTemplates) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.AvailableTemplates = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetBehaviorLinkage(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkage) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.BehaviorLinkage = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetBizAlias(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetChoice(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Choice = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetContent(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Content = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetDataSource(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.DataSource = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetDefaultColor(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.DefaultColor = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetDisabled(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Disabled = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetDuration(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Duration = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetDurationLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.DurationLabel = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetFields(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFields) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Fields = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetFormat(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Format = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetFormula(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Formula = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetInvisible(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Invisible = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetLabelEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetLimit(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Limit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetLink(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Link = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetMode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Mode = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetMulti(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Multi = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetMultiple(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Multiple = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetNeedDetail(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.NeedDetail = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetNotPrint(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.NotPrint = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetNotUpper(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetOptions(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Options = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetPayEnable(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetPlaceholder(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetQuote(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Quote = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetRatio(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Ratio = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetRelateSource(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.RelateSource = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetRequired(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Required = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetRequiredEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetRule(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRule) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Rule = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetSortable(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Sortable = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetSpread(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Spread = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetStatField(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.StatField = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetTableViewMode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.TableViewMode = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetVerticalPrint(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.VerticalPrint = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps) SetWatermark(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps {
	s.Watermark = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsAvailableTemplates struct {
	// example:
	//
	// 审批模板id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 审批模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsAvailableTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsAvailableTemplates) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsAvailableTemplates) SetId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsAvailableTemplates {
	s.Id = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsAvailableTemplates) SetName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsAvailableTemplates {
	s.Name = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkage struct {
	// example:
	//
	// option_0
	OptionKey *string                                                                                `json:"optionKey,omitempty" xml:"optionKey,omitempty"`
	Targets   []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkageTargets `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkage) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkage) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkage) SetOptionKey(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkage {
	s.OptionKey = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkage) SetTargets(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkageTargets) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkage {
	s.Targets = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkageTargets struct {
	// example:
	//
	// NORMAL
	Behavior *string `json:"behavior,omitempty" xml:"behavior,omitempty"`
	// example:
	//
	// TextField_1LTIYOR4XIF40
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkageTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkageTargets) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkageTargets) SetBehavior(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkageTargets {
	s.Behavior = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkageTargets) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsBehaviorLinkageTargets {
	s.FieldId = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource struct {
	// This parameter is required.
	Params *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	Target *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	Type   *string                                                                        `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource) SetParams(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParams) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource {
	s.Params = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource) SetTarget(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource {
	s.Target = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource) SetType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource {
	s.Type = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParams struct {
	// This parameter is required.
	Filters []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParams) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParams) SetFilters(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParams {
	s.Filters = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters) SetFilterType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters {
	s.FilterType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters {
	s.Value = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters) SetValueType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceParamsFilters {
	s.ValueType = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget struct {
	// example:
	//
	// 0：流程表单，1：纯表单
	AppType  *int64  `json:"appType,omitempty" xml:"appType,omitempty"`
	AppUuid  *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	BizType  *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget) SetAppType(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget) SetAppUuid(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget) SetBizType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget) SetFormCode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSourceTarget {
	s.FormCode = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFields struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	RelateProps *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFields) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFields) SetComponentName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFields {
	s.ComponentName = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFields) SetRelateProps(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFields {
	s.RelateProps = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps struct {
	// This parameter is required.
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1：多选，0：单选
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：自动
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DDDateField和DDDateRangeField
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：隐藏
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1:不需要大写, 空或者0:需要大写
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true：是
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Ratio *int64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	Spread *bool `json:"spread,omitempty" xml:"spread,omitempty"`
	// This parameter is required.
	StatField []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：纵向，false：横向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// This parameter is required.
	Watermark *bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetAlign(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Align = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetBizAlias(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetChoice(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Choice = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetContent(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Content = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetDisabled(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Disabled = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetDuration(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Duration = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetDurationLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.DurationLabel = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetFormat(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Format = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetFormula(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Formula = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetInvisible(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Invisible = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetLabelEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetLimit(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Limit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetLink(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Link = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetMode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Mode = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetNotUpper(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetOptions(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Options = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetPayEnable(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetPlaceholder(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetRatio(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Ratio = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetRequired(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Required = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetRequiredEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetSpread(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Spread = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetStatField(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.StatField = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetVerticalPrint(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.VerticalPrint = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps) SetWatermark(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps {
	s.Watermark = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions struct {
	// This parameter is required.
	Extension *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions) SetExtension(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptionsExtension) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions {
	s.Extension = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions) SetKey(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions {
	s.Value = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptionsExtension struct {
	// This parameter is required.
	EditFreeze *bool `json:"editFreeze,omitempty" xml:"editFreeze,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptionsExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptionsExtension) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptionsExtension) SetEditFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptionsExtension {
	s.EditFreeze = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField) SetUpper(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField {
	s.Upper = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions struct {
	// This parameter is required.
	Extension *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	Warn *bool `json:"warn,omitempty" xml:"warn,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions) SetExtension(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptionsExtension) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions {
	s.Extension = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions) SetKey(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions {
	s.Value = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions) SetWarn(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions {
	s.Warn = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptionsExtension struct {
	// This parameter is required.
	EditFreeze *bool `json:"editFreeze,omitempty" xml:"editFreeze,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptionsExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptionsExtension) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptionsExtension) SetEditFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptionsExtension {
	s.EditFreeze = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource struct {
	// This parameter is required.
	BizType    *string                                                                              `json:"bizType,omitempty" xml:"bizType,omitempty"`
	DataSource *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// This parameter is required.
	Fields []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource) SetBizType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource {
	s.BizType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource) SetDataSource(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource {
	s.DataSource = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource) SetFields(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFields) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource {
	s.Fields = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource struct {
	// This parameter is required.
	Params *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// This parameter is required.
	Target *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource) SetParams(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParams) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource {
	s.Params = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource) SetTarget(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource {
	s.Target = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource) SetType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource {
	s.Type = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParams struct {
	// This parameter is required.
	Filters []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParams) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParams) SetFilters(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParams {
	s.Filters = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters) SetFilterType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters {
	s.FilterType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters {
	s.Value = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters) SetValueType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParamsFilters {
	s.ValueType = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget struct {
	// This parameter is required.
	AppType *int64 `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget) SetAppType(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget) SetAppUuid(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget) SetBizType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget) SetFormCode(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget {
	s.FormCode = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFields struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	RelateProps *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFields) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFields) SetComponentName(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFields {
	s.ComponentName = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFields) SetRelateProps(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFields {
	s.RelateProps = v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps struct {
	// This parameter is required.
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1：多选，0：单选
	Choice *int64 `json:"choice,omitempty" xml:"choice,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：自动
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DDDateField和DDDateRangeField
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：隐藏
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// This parameter is required.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	Multi *int64 `json:"multi,omitempty" xml:"multi,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1:不需要大写, 空或者0:需要大写
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// This parameter is required.
	Options []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true：是
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// This parameter is required.
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Quote *int64 `json:"quote,omitempty" xml:"quote,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// This parameter is required.
	StatField []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true：纵向，false：横向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetAlign(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Align = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetBizAlias(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.BizAlias = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetChoice(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Choice = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetContent(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Content = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetDisabled(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Disabled = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetDuration(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Duration = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetFormat(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Format = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetFormula(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Formula = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetInvisible(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Invisible = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetLabelEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetLink(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Link = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetMulti(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Multi = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetNotUpper(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.NotUpper = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetOptions(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Options = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetPayEnable(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.PayEnable = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetPlaceholder(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Placeholder = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetQuote(v int64) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Quote = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetRequired(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Required = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetRequiredEditableFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetStatField(v []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.StatField = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps) SetVerticalPrint(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps {
	s.VerticalPrint = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions struct {
	// This parameter is required.
	Extension *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) SetExtension(v *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions {
	s.Extension = v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) SetKey(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions {
	s.Key = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions {
	s.Value = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension struct {
	// This parameter is required.
	EditFreeze *bool `json:"editFreeze,omitempty" xml:"editFreeze,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension) SetEditFreeze(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension {
	s.EditFreeze = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField) SetUpper(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField {
	s.Upper = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRule struct {
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRule) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRule) SetType(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRule {
	s.Type = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRule) SetValue(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRule {
	s.Value = &v
	return s
}

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField struct {
	// This parameter is required.
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// This parameter is required.
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField) SetFieldId(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField {
	s.FieldId = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField) SetLabel(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField {
	s.Label = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField) SetUnit(v string) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField {
	s.Unit = &v
	return s
}

func (s *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField) SetUpper(v bool) *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField {
	s.Upper = &v
	return s
}

type DescribeRelationMetaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRelationMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRelationMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelationMetaResponse) GoString() string {
	return s.String()
}

func (s *DescribeRelationMetaResponse) SetHeaders(v map[string]*string) *DescribeRelationMetaResponse {
	s.Headers = v
	return s
}

func (s *DescribeRelationMetaResponse) SetStatusCode(v int32) *DescribeRelationMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRelationMetaResponse) SetBody(v *DescribeRelationMetaResponseBody) *DescribeRelationMetaResponse {
	s.Body = v
	return s
}

type FindTargetRelatedFollowRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FindTargetRelatedFollowRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s FindTargetRelatedFollowRecordsHeaders) GoString() string {
	return s.String()
}

func (s *FindTargetRelatedFollowRecordsHeaders) SetCommonHeaders(v map[string]*string) *FindTargetRelatedFollowRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FindTargetRelatedFollowRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *FindTargetRelatedFollowRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FindTargetRelatedFollowRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// customerId
	FollowTargetDataId *string `json:"followTargetDataId,omitempty" xml:"followTargetDataId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer
	FollowTargetType *string `json:"followTargetType,omitempty" xml:"followTargetType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s FindTargetRelatedFollowRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindTargetRelatedFollowRecordsRequest) GoString() string {
	return s.String()
}

func (s *FindTargetRelatedFollowRecordsRequest) SetFollowTargetDataId(v string) *FindTargetRelatedFollowRecordsRequest {
	s.FollowTargetDataId = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsRequest) SetFollowTargetType(v string) *FindTargetRelatedFollowRecordsRequest {
	s.FollowTargetType = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsRequest) SetMaxResults(v int64) *FindTargetRelatedFollowRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsRequest) SetNextToken(v string) *FindTargetRelatedFollowRecordsRequest {
	s.NextToken = &v
	return s
}

type FindTargetRelatedFollowRecordsResponseBody struct {
	Result *FindTargetRelatedFollowRecordsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s FindTargetRelatedFollowRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindTargetRelatedFollowRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *FindTargetRelatedFollowRecordsResponseBody) SetResult(v *FindTargetRelatedFollowRecordsResponseBodyResult) *FindTargetRelatedFollowRecordsResponseBody {
	s.Result = v
	return s
}

type FindTargetRelatedFollowRecordsResponseBodyResult struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 1000
	NextToken  *string                                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResultList []*FindTargetRelatedFollowRecordsResponseBodyResultResultList `json:"resultList,omitempty" xml:"resultList,omitempty" type:"Repeated"`
}

func (s FindTargetRelatedFollowRecordsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FindTargetRelatedFollowRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResult) SetHasMore(v bool) *FindTargetRelatedFollowRecordsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResult) SetNextToken(v string) *FindTargetRelatedFollowRecordsResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResult) SetResultList(v []*FindTargetRelatedFollowRecordsResponseBodyResultResultList) *FindTargetRelatedFollowRecordsResponseBodyResult {
	s.ResultList = v
	return s
}

type FindTargetRelatedFollowRecordsResponseBodyResultResultList struct {
	// example:
	//
	// manager7617
	CreatorUserId *string                                                                    `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	FollowContent []*FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent `json:"followContent,omitempty" xml:"followContent,omitempty" type:"Repeated"`
	// example:
	//
	// customerId
	FollowTargetDataId *string `json:"followTargetDataId,omitempty" xml:"followTargetDataId,omitempty"`
	// example:
	//
	// customer
	FollowTargetType *string `json:"followTargetType,omitempty" xml:"followTargetType,omitempty"`
	// example:
	//
	// 1712629910168
	GmtCreateMilliseconds *string `json:"gmtCreateMilliseconds,omitempty" xml:"gmtCreateMilliseconds,omitempty"`
	// example:
	//
	// 1712629910168
	GmtModifiedMilliseconds *string `json:"gmtModifiedMilliseconds,omitempty" xml:"gmtModifiedMilliseconds,omitempty"`
	// example:
	//
	// manager7617
	ModifierUserId *string `json:"modifierUserId,omitempty" xml:"modifierUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// _aFFogIuRrWlL3hLdvbb5w09951712629910
	RecordInstId *string `json:"recordInstId,omitempty" xml:"recordInstId,omitempty"`
}

func (s FindTargetRelatedFollowRecordsResponseBodyResultResultList) String() string {
	return tea.Prettify(s)
}

func (s FindTargetRelatedFollowRecordsResponseBodyResultResultList) GoString() string {
	return s.String()
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultList) SetCreatorUserId(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultList {
	s.CreatorUserId = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultList) SetFollowContent(v []*FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent) *FindTargetRelatedFollowRecordsResponseBodyResultResultList {
	s.FollowContent = v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultList) SetFollowTargetDataId(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultList {
	s.FollowTargetDataId = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultList) SetFollowTargetType(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultList {
	s.FollowTargetType = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultList) SetGmtCreateMilliseconds(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultList {
	s.GmtCreateMilliseconds = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultList) SetGmtModifiedMilliseconds(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultList {
	s.GmtModifiedMilliseconds = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultList) SetModifierUserId(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultList {
	s.ModifierUserId = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultList) SetRecordInstId(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultList {
	s.RecordInstId = &v
	return s
}

type FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent struct {
	// example:
	//
	// follow_record_related_content
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// example:
	//
	// 扩展value
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// example:
	//
	// TextareaField-K2U5UJAF
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 重点跟进
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent) String() string {
	return tea.Prettify(s)
}

func (s FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent) GoString() string {
	return s.String()
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent) SetBizAlias(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent {
	s.BizAlias = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent) SetExtendValue(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent {
	s.ExtendValue = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent) SetKey(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent {
	s.Key = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent) SetValue(v string) *FindTargetRelatedFollowRecordsResponseBodyResultResultListFollowContent {
	s.Value = &v
	return s
}

type FindTargetRelatedFollowRecordsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindTargetRelatedFollowRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindTargetRelatedFollowRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s FindTargetRelatedFollowRecordsResponse) GoString() string {
	return s.String()
}

func (s *FindTargetRelatedFollowRecordsResponse) SetHeaders(v map[string]*string) *FindTargetRelatedFollowRecordsResponse {
	s.Headers = v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponse) SetStatusCode(v int32) *FindTargetRelatedFollowRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *FindTargetRelatedFollowRecordsResponse) SetBody(v *FindTargetRelatedFollowRecordsResponseBody) *FindTargetRelatedFollowRecordsResponse {
	s.Body = v
	return s
}

type GetAllCustomerRecyclesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAllCustomerRecyclesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAllCustomerRecyclesHeaders) GoString() string {
	return s.String()
}

func (s *GetAllCustomerRecyclesHeaders) SetCommonHeaders(v map[string]*string) *GetAllCustomerRecyclesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllCustomerRecyclesHeaders) SetXAcsDingtalkAccessToken(v string) *GetAllCustomerRecyclesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAllCustomerRecyclesRequest struct {
	// example:
	//
	// 10
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetAllCustomerRecyclesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllCustomerRecyclesRequest) GoString() string {
	return s.String()
}

func (s *GetAllCustomerRecyclesRequest) SetMaxResults(v int32) *GetAllCustomerRecyclesRequest {
	s.MaxResults = &v
	return s
}

func (s *GetAllCustomerRecyclesRequest) SetNextToken(v string) *GetAllCustomerRecyclesRequest {
	s.NextToken = &v
	return s
}

type GetAllCustomerRecyclesResponseBody struct {
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
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// true
	ResultList []*GetAllCustomerRecyclesResponseBodyResultList `json:"resultList,omitempty" xml:"resultList,omitempty" type:"Repeated"`
}

func (s GetAllCustomerRecyclesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllCustomerRecyclesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllCustomerRecyclesResponseBody) SetHasMore(v bool) *GetAllCustomerRecyclesResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetAllCustomerRecyclesResponseBody) SetNextToken(v string) *GetAllCustomerRecyclesResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetAllCustomerRecyclesResponseBody) SetResultList(v []*GetAllCustomerRecyclesResponseBodyResultList) *GetAllCustomerRecyclesResponseBody {
	s.ResultList = v
	return s
}

type GetAllCustomerRecyclesResponseBodyResultList struct {
	// This parameter is required.
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-03-24T09:30Z
	FollowUpActionTime *string `json:"followUpActionTime,omitempty" xml:"followUpActionTime,omitempty"`
	IsDeleted          *bool   `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// example:
	//
	// 2022-03-24T09:30Z
	NotifyTime    *string `json:"notifyTime,omitempty" xml:"notifyTime,omitempty"`
	RecycleRuleId *int64  `json:"recycleRuleId,omitempty" xml:"recycleRuleId,omitempty"`
	// example:
	//
	// 2022-03-24T09:30Z
	RecycleTime *string `json:"recycleTime,omitempty" xml:"recycleTime,omitempty"`
}

func (s GetAllCustomerRecyclesResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s GetAllCustomerRecyclesResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *GetAllCustomerRecyclesResponseBodyResultList) SetCustomerId(v string) *GetAllCustomerRecyclesResponseBodyResultList {
	s.CustomerId = &v
	return s
}

func (s *GetAllCustomerRecyclesResponseBodyResultList) SetFollowUpActionTime(v string) *GetAllCustomerRecyclesResponseBodyResultList {
	s.FollowUpActionTime = &v
	return s
}

func (s *GetAllCustomerRecyclesResponseBodyResultList) SetIsDeleted(v bool) *GetAllCustomerRecyclesResponseBodyResultList {
	s.IsDeleted = &v
	return s
}

func (s *GetAllCustomerRecyclesResponseBodyResultList) SetNotifyTime(v string) *GetAllCustomerRecyclesResponseBodyResultList {
	s.NotifyTime = &v
	return s
}

func (s *GetAllCustomerRecyclesResponseBodyResultList) SetRecycleRuleId(v int64) *GetAllCustomerRecyclesResponseBodyResultList {
	s.RecycleRuleId = &v
	return s
}

func (s *GetAllCustomerRecyclesResponseBodyResultList) SetRecycleTime(v string) *GetAllCustomerRecyclesResponseBodyResultList {
	s.RecycleTime = &v
	return s
}

type GetAllCustomerRecyclesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllCustomerRecyclesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllCustomerRecyclesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllCustomerRecyclesResponse) GoString() string {
	return s.String()
}

func (s *GetAllCustomerRecyclesResponse) SetHeaders(v map[string]*string) *GetAllCustomerRecyclesResponse {
	s.Headers = v
	return s
}

func (s *GetAllCustomerRecyclesResponse) SetStatusCode(v int32) *GetAllCustomerRecyclesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllCustomerRecyclesResponse) SetBody(v *GetAllCustomerRecyclesResponseBody) *GetAllCustomerRecyclesResponse {
	s.Body = v
	return s
}

type GetContactsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetContactsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetContactsHeaders) GoString() string {
	return s.String()
}

func (s *GetContactsHeaders) SetCommonHeaders(v map[string]*string) *GetContactsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetContactsHeaders) SetXAcsDingtalkAccessToken(v string) *GetContactsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetContactsRequest struct {
	// example:
	//
	// user01
	CurrentOperatorUserId *string `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
	// This parameter is required.
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// crm_contact
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// example:
	//
	// dingxxx
	ProviderCorpId *string `json:"providerCorpId,omitempty" xml:"providerCorpId,omitempty"`
	// example:
	//
	// {"queryGroupList":[{"logicType":"AND","queryObjectList":[{"fieldId":"contact_phone","value":"18000000000"},{"fieldId":"contact_related_customer","value":"INST-XXX"}]}]}
	QueryDsl *string `json:"queryDsl,omitempty" xml:"queryDsl,omitempty"`
}

func (s GetContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetContactsRequest) GoString() string {
	return s.String()
}

func (s *GetContactsRequest) SetCurrentOperatorUserId(v string) *GetContactsRequest {
	s.CurrentOperatorUserId = &v
	return s
}

func (s *GetContactsRequest) SetMaxResults(v int64) *GetContactsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetContactsRequest) SetNextToken(v string) *GetContactsRequest {
	s.NextToken = &v
	return s
}

func (s *GetContactsRequest) SetObjectType(v string) *GetContactsRequest {
	s.ObjectType = &v
	return s
}

func (s *GetContactsRequest) SetProviderCorpId(v string) *GetContactsRequest {
	s.ProviderCorpId = &v
	return s
}

func (s *GetContactsRequest) SetQueryDsl(v string) *GetContactsRequest {
	s.QueryDsl = &v
	return s
}

type GetContactsResponseBody struct {
	Result *GetContactsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContactsResponseBody) GoString() string {
	return s.String()
}

func (s *GetContactsResponseBody) SetResult(v *GetContactsResponseBodyResult) *GetContactsResponseBody {
	s.Result = v
	return s
}

type GetContactsResponseBodyResult struct {
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 100
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Values    []*GetContactsResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetContactsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetContactsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetContactsResponseBodyResult) SetHasMore(v bool) *GetContactsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetContactsResponseBodyResult) SetMaxResults(v int64) *GetContactsResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *GetContactsResponseBodyResult) SetNextToken(v string) *GetContactsResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *GetContactsResponseBodyResult) SetValues(v []*GetContactsResponseBodyResultValues) *GetContactsResponseBodyResult {
	s.Values = v
	return s
}

type GetContactsResponseBodyResultValues struct {
	// example:
	//
	// user01
	CreatorUserId *string                `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Data          map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData    map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// example:
	//
	// 2023-11-25 15:33:12
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 20123-12-25 15:33:12
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// INST_XX
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// crm_contact
	ObjectType *string                                        `json:"objectType,omitempty" xml:"objectType,omitempty"`
	Permission *GetContactsResponseBodyResultValuesPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
}

func (s GetContactsResponseBodyResultValues) String() string {
	return tea.Prettify(s)
}

func (s GetContactsResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *GetContactsResponseBodyResultValues) SetCreatorUserId(v string) *GetContactsResponseBodyResultValues {
	s.CreatorUserId = &v
	return s
}

func (s *GetContactsResponseBodyResultValues) SetData(v map[string]interface{}) *GetContactsResponseBodyResultValues {
	s.Data = v
	return s
}

func (s *GetContactsResponseBodyResultValues) SetExtendData(v map[string]interface{}) *GetContactsResponseBodyResultValues {
	s.ExtendData = v
	return s
}

func (s *GetContactsResponseBodyResultValues) SetGmtCreate(v string) *GetContactsResponseBodyResultValues {
	s.GmtCreate = &v
	return s
}

func (s *GetContactsResponseBodyResultValues) SetGmtModified(v string) *GetContactsResponseBodyResultValues {
	s.GmtModified = &v
	return s
}

func (s *GetContactsResponseBodyResultValues) SetInstanceId(v string) *GetContactsResponseBodyResultValues {
	s.InstanceId = &v
	return s
}

func (s *GetContactsResponseBodyResultValues) SetObjectType(v string) *GetContactsResponseBodyResultValues {
	s.ObjectType = &v
	return s
}

func (s *GetContactsResponseBodyResultValues) SetPermission(v *GetContactsResponseBodyResultValuesPermission) *GetContactsResponseBodyResultValues {
	s.Permission = v
	return s
}

type GetContactsResponseBodyResultValuesPermission struct {
	OwnerUserIds       []*string `json:"ownerUserIds,omitempty" xml:"ownerUserIds,omitempty" type:"Repeated"`
	ParticipantUserIds []*string `json:"participantUserIds,omitempty" xml:"participantUserIds,omitempty" type:"Repeated"`
}

func (s GetContactsResponseBodyResultValuesPermission) String() string {
	return tea.Prettify(s)
}

func (s GetContactsResponseBodyResultValuesPermission) GoString() string {
	return s.String()
}

func (s *GetContactsResponseBodyResultValuesPermission) SetOwnerUserIds(v []*string) *GetContactsResponseBodyResultValuesPermission {
	s.OwnerUserIds = v
	return s
}

func (s *GetContactsResponseBodyResultValuesPermission) SetParticipantUserIds(v []*string) *GetContactsResponseBodyResultValuesPermission {
	s.ParticipantUserIds = v
	return s
}

type GetContactsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContactsResponse) GoString() string {
	return s.String()
}

func (s *GetContactsResponse) SetHeaders(v map[string]*string) *GetContactsResponse {
	s.Headers = v
	return s
}

func (s *GetContactsResponse) SetStatusCode(v int32) *GetContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContactsResponse) SetBody(v *GetContactsResponseBody) *GetContactsResponse {
	s.Body = v
	return s
}

type GetCrmGroupChatHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCrmGroupChatHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatHeaders) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatHeaders) SetCommonHeaders(v map[string]*string) *GetCrmGroupChatHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCrmGroupChatHeaders) SetXAcsDingtalkAccessToken(v string) *GetCrmGroupChatHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCrmGroupChatResponseBody struct {
	ChatId    *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	GmtCreate *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// https://static/xx.com/xx.jpg
	IconUrl            *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	MemberCount        *int32  `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenGroupSetId     *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	OwnerUserId        *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	OwnerUserName      *string `json:"ownerUserName,omitempty" xml:"ownerUserName,omitempty"`
}

func (s GetCrmGroupChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatResponseBody) SetChatId(v string) *GetCrmGroupChatResponseBody {
	s.ChatId = &v
	return s
}

func (s *GetCrmGroupChatResponseBody) SetGmtCreate(v int64) *GetCrmGroupChatResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetCrmGroupChatResponseBody) SetIconUrl(v string) *GetCrmGroupChatResponseBody {
	s.IconUrl = &v
	return s
}

func (s *GetCrmGroupChatResponseBody) SetMemberCount(v int32) *GetCrmGroupChatResponseBody {
	s.MemberCount = &v
	return s
}

func (s *GetCrmGroupChatResponseBody) SetName(v string) *GetCrmGroupChatResponseBody {
	s.Name = &v
	return s
}

func (s *GetCrmGroupChatResponseBody) SetOpenConversationId(v string) *GetCrmGroupChatResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GetCrmGroupChatResponseBody) SetOpenGroupSetId(v string) *GetCrmGroupChatResponseBody {
	s.OpenGroupSetId = &v
	return s
}

func (s *GetCrmGroupChatResponseBody) SetOwnerUserId(v string) *GetCrmGroupChatResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *GetCrmGroupChatResponseBody) SetOwnerUserName(v string) *GetCrmGroupChatResponseBody {
	s.OwnerUserName = &v
	return s
}

type GetCrmGroupChatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrmGroupChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrmGroupChatResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatResponse) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatResponse) SetHeaders(v map[string]*string) *GetCrmGroupChatResponse {
	s.Headers = v
	return s
}

func (s *GetCrmGroupChatResponse) SetStatusCode(v int32) *GetCrmGroupChatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrmGroupChatResponse) SetBody(v *GetCrmGroupChatResponseBody) *GetCrmGroupChatResponse {
	s.Body = v
	return s
}

type GetCrmGroupChatMultiHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCrmGroupChatMultiHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatMultiHeaders) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatMultiHeaders) SetCommonHeaders(v map[string]*string) *GetCrmGroupChatMultiHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCrmGroupChatMultiHeaders) SetXAcsDingtalkAccessToken(v string) *GetCrmGroupChatMultiHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCrmGroupChatMultiRequest struct {
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
}

func (s GetCrmGroupChatMultiRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatMultiRequest) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatMultiRequest) SetOpenConversationIds(v []*string) *GetCrmGroupChatMultiRequest {
	s.OpenConversationIds = v
	return s
}

type GetCrmGroupChatMultiResponseBody struct {
	Result []*GetCrmGroupChatMultiResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetCrmGroupChatMultiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatMultiResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatMultiResponseBody) SetResult(v []*GetCrmGroupChatMultiResponseBodyResult) *GetCrmGroupChatMultiResponseBody {
	s.Result = v
	return s
}

type GetCrmGroupChatMultiResponseBodyResult struct {
	// example:
	//
	// 1642078998377
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// https://static/xx.com/xx.jpg
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// example:
	//
	// 12
	MemberCount *int32 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// example:
	//
	// 营销1群
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// xx==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// xxa==
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// example:
	//
	// axaf12
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// example:
	//
	// XX
	OwnerUserName *string `json:"ownerUserName,omitempty" xml:"ownerUserName,omitempty"`
}

func (s GetCrmGroupChatMultiResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatMultiResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatMultiResponseBodyResult) SetGmtCreate(v int64) *GetCrmGroupChatMultiResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *GetCrmGroupChatMultiResponseBodyResult) SetIconUrl(v string) *GetCrmGroupChatMultiResponseBodyResult {
	s.IconUrl = &v
	return s
}

func (s *GetCrmGroupChatMultiResponseBodyResult) SetMemberCount(v int32) *GetCrmGroupChatMultiResponseBodyResult {
	s.MemberCount = &v
	return s
}

func (s *GetCrmGroupChatMultiResponseBodyResult) SetName(v string) *GetCrmGroupChatMultiResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetCrmGroupChatMultiResponseBodyResult) SetOpenConversationId(v string) *GetCrmGroupChatMultiResponseBodyResult {
	s.OpenConversationId = &v
	return s
}

func (s *GetCrmGroupChatMultiResponseBodyResult) SetOpenGroupSetId(v string) *GetCrmGroupChatMultiResponseBodyResult {
	s.OpenGroupSetId = &v
	return s
}

func (s *GetCrmGroupChatMultiResponseBodyResult) SetOwnerUserId(v string) *GetCrmGroupChatMultiResponseBodyResult {
	s.OwnerUserId = &v
	return s
}

func (s *GetCrmGroupChatMultiResponseBodyResult) SetOwnerUserName(v string) *GetCrmGroupChatMultiResponseBodyResult {
	s.OwnerUserName = &v
	return s
}

type GetCrmGroupChatMultiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrmGroupChatMultiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrmGroupChatMultiResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatMultiResponse) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatMultiResponse) SetHeaders(v map[string]*string) *GetCrmGroupChatMultiResponse {
	s.Headers = v
	return s
}

func (s *GetCrmGroupChatMultiResponse) SetStatusCode(v int32) *GetCrmGroupChatMultiResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrmGroupChatMultiResponse) SetBody(v *GetCrmGroupChatMultiResponseBody) *GetCrmGroupChatMultiResponse {
	s.Body = v
	return s
}

type GetCrmGroupChatSingleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCrmGroupChatSingleHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatSingleHeaders) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatSingleHeaders) SetCommonHeaders(v map[string]*string) *GetCrmGroupChatSingleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCrmGroupChatSingleHeaders) SetXAcsDingtalkAccessToken(v string) *GetCrmGroupChatSingleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCrmGroupChatSingleRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetCrmGroupChatSingleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatSingleRequest) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatSingleRequest) SetOpenConversationId(v string) *GetCrmGroupChatSingleRequest {
	s.OpenConversationId = &v
	return s
}

type GetCrmGroupChatSingleResponseBody struct {
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// https://static/xx.com/xx.jpg
	IconUrl            *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	MemberCount        *int32  `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenGroupSetId     *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	OwnerUserId        *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	OwnerUserName      *string `json:"ownerUserName,omitempty" xml:"ownerUserName,omitempty"`
}

func (s GetCrmGroupChatSingleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatSingleResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatSingleResponseBody) SetGmtCreate(v int64) *GetCrmGroupChatSingleResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetCrmGroupChatSingleResponseBody) SetIconUrl(v string) *GetCrmGroupChatSingleResponseBody {
	s.IconUrl = &v
	return s
}

func (s *GetCrmGroupChatSingleResponseBody) SetMemberCount(v int32) *GetCrmGroupChatSingleResponseBody {
	s.MemberCount = &v
	return s
}

func (s *GetCrmGroupChatSingleResponseBody) SetName(v string) *GetCrmGroupChatSingleResponseBody {
	s.Name = &v
	return s
}

func (s *GetCrmGroupChatSingleResponseBody) SetOpenConversationId(v string) *GetCrmGroupChatSingleResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GetCrmGroupChatSingleResponseBody) SetOpenGroupSetId(v string) *GetCrmGroupChatSingleResponseBody {
	s.OpenGroupSetId = &v
	return s
}

func (s *GetCrmGroupChatSingleResponseBody) SetOwnerUserId(v string) *GetCrmGroupChatSingleResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *GetCrmGroupChatSingleResponseBody) SetOwnerUserName(v string) *GetCrmGroupChatSingleResponseBody {
	s.OwnerUserName = &v
	return s
}

type GetCrmGroupChatSingleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrmGroupChatSingleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrmGroupChatSingleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCrmGroupChatSingleResponse) GoString() string {
	return s.String()
}

func (s *GetCrmGroupChatSingleResponse) SetHeaders(v map[string]*string) *GetCrmGroupChatSingleResponse {
	s.Headers = v
	return s
}

func (s *GetCrmGroupChatSingleResponse) SetStatusCode(v int32) *GetCrmGroupChatSingleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrmGroupChatSingleResponse) SetBody(v *GetCrmGroupChatSingleResponseBody) *GetCrmGroupChatSingleResponse {
	s.Body = v
	return s
}

type GetCrmRolePermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCrmRolePermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCrmRolePermissionHeaders) GoString() string {
	return s.String()
}

func (s *GetCrmRolePermissionHeaders) SetCommonHeaders(v map[string]*string) *GetCrmRolePermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCrmRolePermissionHeaders) SetXAcsDingtalkAccessToken(v string) *GetCrmRolePermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCrmRolePermissionRequest struct {
	// example:
	//
	// crm_customer
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// PROC-9EC85C45-E404-4E26-9300-E67455F0FF8F
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
}

func (s GetCrmRolePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCrmRolePermissionRequest) GoString() string {
	return s.String()
}

func (s *GetCrmRolePermissionRequest) SetBizType(v string) *GetCrmRolePermissionRequest {
	s.BizType = &v
	return s
}

func (s *GetCrmRolePermissionRequest) SetResourceId(v string) *GetCrmRolePermissionRequest {
	s.ResourceId = &v
	return s
}

type GetCrmRolePermissionResponseBody struct {
	// This parameter is required.
	Permissions []*GetCrmRolePermissionResponseBodyPermissions `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetCrmRolePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCrmRolePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrmRolePermissionResponseBody) SetPermissions(v []*GetCrmRolePermissionResponseBodyPermissions) *GetCrmRolePermissionResponseBody {
	s.Permissions = v
	return s
}

type GetCrmRolePermissionResponseBodyPermissions struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	DefaultRole *bool `json:"defaultRole,omitempty" xml:"defaultRole,omitempty"`
	// This parameter is required.
	FieldScopes []*GetCrmRolePermissionResponseBodyPermissionsFieldScopes `json:"fieldScopes,omitempty" xml:"fieldScopes,omitempty" type:"Repeated"`
	// This parameter is required.
	ManagingScopeList []*GetCrmRolePermissionResponseBodyPermissionsManagingScopeList `json:"managingScopeList,omitempty" xml:"managingScopeList,omitempty" type:"Repeated"`
	// This parameter is required.
	OperateScopes []*GetCrmRolePermissionResponseBodyPermissionsOperateScopes `json:"operateScopes,omitempty" xml:"operateScopes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-478E50CA-856C-4C08-B806-E664D4CEC8C4
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12821738
	RoleId *float64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// This parameter is required.
	RoleMemberList []*GetCrmRolePermissionResponseBodyPermissionsRoleMemberList `json:"roleMemberList,omitempty" xml:"roleMemberList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 销售权限组
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s GetCrmRolePermissionResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s GetCrmRolePermissionResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *GetCrmRolePermissionResponseBodyPermissions) SetDefaultRole(v bool) *GetCrmRolePermissionResponseBodyPermissions {
	s.DefaultRole = &v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissions) SetFieldScopes(v []*GetCrmRolePermissionResponseBodyPermissionsFieldScopes) *GetCrmRolePermissionResponseBodyPermissions {
	s.FieldScopes = v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissions) SetManagingScopeList(v []*GetCrmRolePermissionResponseBodyPermissionsManagingScopeList) *GetCrmRolePermissionResponseBodyPermissions {
	s.ManagingScopeList = v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissions) SetOperateScopes(v []*GetCrmRolePermissionResponseBodyPermissionsOperateScopes) *GetCrmRolePermissionResponseBodyPermissions {
	s.OperateScopes = v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissions) SetResourceId(v string) *GetCrmRolePermissionResponseBodyPermissions {
	s.ResourceId = &v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissions) SetRoleId(v float64) *GetCrmRolePermissionResponseBodyPermissions {
	s.RoleId = &v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissions) SetRoleMemberList(v []*GetCrmRolePermissionResponseBodyPermissionsRoleMemberList) *GetCrmRolePermissionResponseBodyPermissions {
	s.RoleMemberList = v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissions) SetRoleName(v string) *GetCrmRolePermissionResponseBodyPermissions {
	s.RoleName = &v
	return s
}

type GetCrmRolePermissionResponseBodyPermissionsFieldScopes struct {
	// This parameter is required.
	FieldActions []*string `json:"fieldActions,omitempty" xml:"fieldActions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// "DDSelectField-KI5S975E"
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
}

func (s GetCrmRolePermissionResponseBodyPermissionsFieldScopes) String() string {
	return tea.Prettify(s)
}

func (s GetCrmRolePermissionResponseBodyPermissionsFieldScopes) GoString() string {
	return s.String()
}

func (s *GetCrmRolePermissionResponseBodyPermissionsFieldScopes) SetFieldActions(v []*string) *GetCrmRolePermissionResponseBodyPermissionsFieldScopes {
	s.FieldActions = v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissionsFieldScopes) SetFieldId(v string) *GetCrmRolePermissionResponseBodyPermissionsFieldScopes {
	s.FieldId = &v
	return s
}

type GetCrmRolePermissionResponseBodyPermissionsManagingScopeList struct {
	// This parameter is required.
	Ext *GetCrmRolePermissionResponseBodyPermissionsManagingScopeListExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// false 如果是主管，要做逻辑的单独处理。比如如果设置了管理下属或当前部门，只管理他是主管的部门
	Manager *bool `json:"manager,omitempty" xml:"manager,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10_own 自己负责的 15_all_org 全公司 20_selfDept 同层级 30_selfSubDept 下属的 40_customized 自定义的
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetCrmRolePermissionResponseBodyPermissionsManagingScopeList) String() string {
	return tea.Prettify(s)
}

func (s GetCrmRolePermissionResponseBodyPermissionsManagingScopeList) GoString() string {
	return s.String()
}

func (s *GetCrmRolePermissionResponseBodyPermissionsManagingScopeList) SetExt(v *GetCrmRolePermissionResponseBodyPermissionsManagingScopeListExt) *GetCrmRolePermissionResponseBodyPermissionsManagingScopeList {
	s.Ext = v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissionsManagingScopeList) SetManager(v bool) *GetCrmRolePermissionResponseBodyPermissionsManagingScopeList {
	s.Manager = &v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissionsManagingScopeList) SetType(v string) *GetCrmRolePermissionResponseBodyPermissionsManagingScopeList {
	s.Type = &v
	return s
}

type GetCrmRolePermissionResponseBodyPermissionsManagingScopeListExt struct {
	// This parameter is required.
	DeptIdList []*float64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s GetCrmRolePermissionResponseBodyPermissionsManagingScopeListExt) String() string {
	return tea.Prettify(s)
}

func (s GetCrmRolePermissionResponseBodyPermissionsManagingScopeListExt) GoString() string {
	return s.String()
}

func (s *GetCrmRolePermissionResponseBodyPermissionsManagingScopeListExt) SetDeptIdList(v []*float64) *GetCrmRolePermissionResponseBodyPermissionsManagingScopeListExt {
	s.DeptIdList = v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissionsManagingScopeListExt) SetUserIdList(v []*string) *GetCrmRolePermissionResponseBodyPermissionsManagingScopeListExt {
	s.UserIdList = v
	return s
}

type GetCrmRolePermissionResponseBodyPermissionsOperateScopes struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasAuth *bool `json:"hasAuth,omitempty" xml:"hasAuth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 	- 操作类型      	- 发起：OPERATE_CREATE      	- 查看：OPERATE_VIEW      	- 编辑：OPERATE_EDIT      	- 删除：OPERATE_DELETE      	- 打印：OPERATE_PRINT      	- 分配：ASSIGN      	- 转交：TRANS      	- 导入：IMPORT      	- 导出：EXPORT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetCrmRolePermissionResponseBodyPermissionsOperateScopes) String() string {
	return tea.Prettify(s)
}

func (s GetCrmRolePermissionResponseBodyPermissionsOperateScopes) GoString() string {
	return s.String()
}

func (s *GetCrmRolePermissionResponseBodyPermissionsOperateScopes) SetHasAuth(v bool) *GetCrmRolePermissionResponseBodyPermissionsOperateScopes {
	s.HasAuth = &v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissionsOperateScopes) SetType(v string) *GetCrmRolePermissionResponseBodyPermissionsOperateScopes {
	s.Type = &v
	return s
}

type GetCrmRolePermissionResponseBodyPermissionsRoleMemberList struct {
	// This parameter is required.
	//
	// example:
	//
	// 可以是员工 uid，可以是部门 ID 等，根据 type 确定
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user：组织成员   dept：部门   tag：标签  org：组织     org_res_admin：组织管理员
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// manager1234
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetCrmRolePermissionResponseBodyPermissionsRoleMemberList) String() string {
	return tea.Prettify(s)
}

func (s GetCrmRolePermissionResponseBodyPermissionsRoleMemberList) GoString() string {
	return s.String()
}

func (s *GetCrmRolePermissionResponseBodyPermissionsRoleMemberList) SetMemberId(v string) *GetCrmRolePermissionResponseBodyPermissionsRoleMemberList {
	s.MemberId = &v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissionsRoleMemberList) SetName(v string) *GetCrmRolePermissionResponseBodyPermissionsRoleMemberList {
	s.Name = &v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissionsRoleMemberList) SetType(v string) *GetCrmRolePermissionResponseBodyPermissionsRoleMemberList {
	s.Type = &v
	return s
}

func (s *GetCrmRolePermissionResponseBodyPermissionsRoleMemberList) SetUserId(v string) *GetCrmRolePermissionResponseBodyPermissionsRoleMemberList {
	s.UserId = &v
	return s
}

type GetCrmRolePermissionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrmRolePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrmRolePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCrmRolePermissionResponse) GoString() string {
	return s.String()
}

func (s *GetCrmRolePermissionResponse) SetHeaders(v map[string]*string) *GetCrmRolePermissionResponse {
	s.Headers = v
	return s
}

func (s *GetCrmRolePermissionResponse) SetStatusCode(v int32) *GetCrmRolePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrmRolePermissionResponse) SetBody(v *GetCrmRolePermissionResponseBody) *GetCrmRolePermissionResponse {
	s.Body = v
	return s
}

type GetCustomerTracksByRelationIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCustomerTracksByRelationIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerTracksByRelationIdHeaders) GoString() string {
	return s.String()
}

func (s *GetCustomerTracksByRelationIdHeaders) SetCommonHeaders(v map[string]*string) *GetCustomerTracksByRelationIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCustomerTracksByRelationIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetCustomerTracksByRelationIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCustomerTracksByRelationIdRequest struct {
	// example:
	//
	// 10
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fasd-afsd1-21312-faaa
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// example:
	//
	// 0
	TypeGroup *int32 `json:"typeGroup,omitempty" xml:"typeGroup,omitempty"`
}

func (s GetCustomerTracksByRelationIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerTracksByRelationIdRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerTracksByRelationIdRequest) SetMaxResults(v int32) *GetCustomerTracksByRelationIdRequest {
	s.MaxResults = &v
	return s
}

func (s *GetCustomerTracksByRelationIdRequest) SetNextToken(v string) *GetCustomerTracksByRelationIdRequest {
	s.NextToken = &v
	return s
}

func (s *GetCustomerTracksByRelationIdRequest) SetRelationId(v string) *GetCustomerTracksByRelationIdRequest {
	s.RelationId = &v
	return s
}

func (s *GetCustomerTracksByRelationIdRequest) SetTypeGroup(v int32) *GetCustomerTracksByRelationIdRequest {
	s.TypeGroup = &v
	return s
}

type GetCustomerTracksByRelationIdResponseBody struct {
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
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// true
	ResultList []*GetCustomerTracksByRelationIdResponseBodyResultList `json:"resultList,omitempty" xml:"resultList,omitempty" type:"Repeated"`
}

func (s GetCustomerTracksByRelationIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerTracksByRelationIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerTracksByRelationIdResponseBody) SetHasMore(v bool) *GetCustomerTracksByRelationIdResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBody) SetNextToken(v string) *GetCustomerTracksByRelationIdResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBody) SetResultList(v []*GetCustomerTracksByRelationIdResponseBodyResultList) *GetCustomerTracksByRelationIdResponseBody {
	s.ResultList = v
	return s
}

type GetCustomerTracksByRelationIdResponseBodyResultList struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 华佗
	CreatorName *string            `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	Detail      map[string]*string `json:"detail,omitempty" xml:"detail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// markdown
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-03-24T09:30Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// dadf134234
	Id      *string                                                     `json:"id,omitempty" xml:"id,omitempty"`
	IsvInfo *GetCustomerTracksByRelationIdResponseBodyResultListIsvInfo `json:"isvInfo,omitempty" xml:"isvInfo,omitempty" type:"Struct"`
	Title   *string                                                     `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 201
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TypeGroup *int32 `json:"typeGroup,omitempty" xml:"typeGroup,omitempty"`
}

func (s GetCustomerTracksByRelationIdResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerTracksByRelationIdResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultList) SetContent(v string) *GetCustomerTracksByRelationIdResponseBodyResultList {
	s.Content = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultList) SetCreatorName(v string) *GetCustomerTracksByRelationIdResponseBodyResultList {
	s.CreatorName = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultList) SetDetail(v map[string]*string) *GetCustomerTracksByRelationIdResponseBodyResultList {
	s.Detail = v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultList) SetFormat(v string) *GetCustomerTracksByRelationIdResponseBodyResultList {
	s.Format = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultList) SetGmtCreate(v string) *GetCustomerTracksByRelationIdResponseBodyResultList {
	s.GmtCreate = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultList) SetId(v string) *GetCustomerTracksByRelationIdResponseBodyResultList {
	s.Id = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultList) SetIsvInfo(v *GetCustomerTracksByRelationIdResponseBodyResultListIsvInfo) *GetCustomerTracksByRelationIdResponseBodyResultList {
	s.IsvInfo = v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultList) SetTitle(v string) *GetCustomerTracksByRelationIdResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultList) SetType(v int32) *GetCustomerTracksByRelationIdResponseBodyResultList {
	s.Type = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultList) SetTypeGroup(v int32) *GetCustomerTracksByRelationIdResponseBodyResultList {
	s.TypeGroup = &v
	return s
}

type GetCustomerTracksByRelationIdResponseBodyResultListIsvInfo struct {
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
}

func (s GetCustomerTracksByRelationIdResponseBodyResultListIsvInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerTracksByRelationIdResponseBodyResultListIsvInfo) GoString() string {
	return s.String()
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultListIsvInfo) SetAppName(v string) *GetCustomerTracksByRelationIdResponseBodyResultListIsvInfo {
	s.AppName = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponseBodyResultListIsvInfo) SetOrgName(v string) *GetCustomerTracksByRelationIdResponseBodyResultListIsvInfo {
	s.OrgName = &v
	return s
}

type GetCustomerTracksByRelationIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerTracksByRelationIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerTracksByRelationIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerTracksByRelationIdResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerTracksByRelationIdResponse) SetHeaders(v map[string]*string) *GetCustomerTracksByRelationIdResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerTracksByRelationIdResponse) SetStatusCode(v int32) *GetCustomerTracksByRelationIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerTracksByRelationIdResponse) SetBody(v *GetCustomerTracksByRelationIdResponseBody) *GetCustomerTracksByRelationIdResponse {
	s.Body = v
	return s
}

type GetGroupSetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetGroupSetHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGroupSetHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupSetHeaders) SetCommonHeaders(v map[string]*string) *GetGroupSetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupSetHeaders) SetXAcsDingtalkAccessToken(v string) *GetGroupSetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetGroupSetRequest struct {
	// This parameter is required.
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
}

func (s GetGroupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupSetRequest) GoString() string {
	return s.String()
}

func (s *GetGroupSetRequest) SetOpenGroupSetId(v string) *GetGroupSetRequest {
	s.OpenGroupSetId = &v
	return s
}

type GetGroupSetResponseBody struct {
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 5
	GroupChatCount         *int32  `json:"groupChatCount,omitempty" xml:"groupChatCount,omitempty"`
	InviteLink             *string `json:"inviteLink,omitempty" xml:"inviteLink,omitempty"`
	LastOpenConversationId *string `json:"lastOpenConversationId,omitempty" xml:"lastOpenConversationId,omitempty"`
	// This parameter is required.
	Manager        []*GetGroupSetResponseBodyManager `json:"manager,omitempty" xml:"manager,omitempty" type:"Repeated"`
	ManagerUserIds *string                           `json:"managerUserIds,omitempty" xml:"managerUserIds,omitempty"`
	MemberCount    *int32                            `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	MemberQuota    *int32                            `json:"memberQuota,omitempty" xml:"memberQuota,omitempty"`
	Name           *string                           `json:"name,omitempty" xml:"name,omitempty"`
	Notice         *string                           `json:"notice,omitempty" xml:"notice,omitempty"`
	NoticeToped    *int32                            `json:"noticeToped,omitempty" xml:"noticeToped,omitempty"`
	OpenGroupSetId *string                           `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	Owner        *GetGroupSetResponseBodyOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	OwnerUserId  *string                       `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	RelationType *string                       `json:"relationType,omitempty" xml:"relationType,omitempty"`
	TemplateId   *string                       `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s GetGroupSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupSetResponseBody) SetGmtCreate(v string) *GetGroupSetResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetGroupSetResponseBody) SetGmtModified(v string) *GetGroupSetResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetGroupSetResponseBody) SetGroupChatCount(v int32) *GetGroupSetResponseBody {
	s.GroupChatCount = &v
	return s
}

func (s *GetGroupSetResponseBody) SetInviteLink(v string) *GetGroupSetResponseBody {
	s.InviteLink = &v
	return s
}

func (s *GetGroupSetResponseBody) SetLastOpenConversationId(v string) *GetGroupSetResponseBody {
	s.LastOpenConversationId = &v
	return s
}

func (s *GetGroupSetResponseBody) SetManager(v []*GetGroupSetResponseBodyManager) *GetGroupSetResponseBody {
	s.Manager = v
	return s
}

func (s *GetGroupSetResponseBody) SetManagerUserIds(v string) *GetGroupSetResponseBody {
	s.ManagerUserIds = &v
	return s
}

func (s *GetGroupSetResponseBody) SetMemberCount(v int32) *GetGroupSetResponseBody {
	s.MemberCount = &v
	return s
}

func (s *GetGroupSetResponseBody) SetMemberQuota(v int32) *GetGroupSetResponseBody {
	s.MemberQuota = &v
	return s
}

func (s *GetGroupSetResponseBody) SetName(v string) *GetGroupSetResponseBody {
	s.Name = &v
	return s
}

func (s *GetGroupSetResponseBody) SetNotice(v string) *GetGroupSetResponseBody {
	s.Notice = &v
	return s
}

func (s *GetGroupSetResponseBody) SetNoticeToped(v int32) *GetGroupSetResponseBody {
	s.NoticeToped = &v
	return s
}

func (s *GetGroupSetResponseBody) SetOpenGroupSetId(v string) *GetGroupSetResponseBody {
	s.OpenGroupSetId = &v
	return s
}

func (s *GetGroupSetResponseBody) SetOwner(v *GetGroupSetResponseBodyOwner) *GetGroupSetResponseBody {
	s.Owner = v
	return s
}

func (s *GetGroupSetResponseBody) SetOwnerUserId(v string) *GetGroupSetResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *GetGroupSetResponseBody) SetRelationType(v string) *GetGroupSetResponseBody {
	s.RelationType = &v
	return s
}

func (s *GetGroupSetResponseBody) SetTemplateId(v string) *GetGroupSetResponseBody {
	s.TemplateId = &v
	return s
}

type GetGroupSetResponseBodyManager struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetGroupSetResponseBodyManager) String() string {
	return tea.Prettify(s)
}

func (s GetGroupSetResponseBodyManager) GoString() string {
	return s.String()
}

func (s *GetGroupSetResponseBodyManager) SetName(v string) *GetGroupSetResponseBodyManager {
	s.Name = &v
	return s
}

func (s *GetGroupSetResponseBodyManager) SetUserId(v string) *GetGroupSetResponseBodyManager {
	s.UserId = &v
	return s
}

type GetGroupSetResponseBodyOwner struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetGroupSetResponseBodyOwner) String() string {
	return tea.Prettify(s)
}

func (s GetGroupSetResponseBodyOwner) GoString() string {
	return s.String()
}

func (s *GetGroupSetResponseBodyOwner) SetName(v string) *GetGroupSetResponseBodyOwner {
	s.Name = &v
	return s
}

func (s *GetGroupSetResponseBodyOwner) SetUserId(v string) *GetGroupSetResponseBodyOwner {
	s.UserId = &v
	return s
}

type GetGroupSetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupSetResponse) GoString() string {
	return s.String()
}

func (s *GetGroupSetResponse) SetHeaders(v map[string]*string) *GetGroupSetResponse {
	s.Headers = v
	return s
}

func (s *GetGroupSetResponse) SetStatusCode(v int32) *GetGroupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupSetResponse) SetBody(v *GetGroupSetResponseBody) *GetGroupSetResponse {
	s.Body = v
	return s
}

type GetInAppPurchaseGoodsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInAppPurchaseGoodsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInAppPurchaseGoodsHeaders) GoString() string {
	return s.String()
}

func (s *GetInAppPurchaseGoodsHeaders) SetCommonHeaders(v map[string]*string) *GetInAppPurchaseGoodsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInAppPurchaseGoodsHeaders) SetXAcsDingtalkAccessToken(v string) *GetInAppPurchaseGoodsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInAppPurchaseGoodsRequest struct {
	// example:
	//
	// uhdhjsabdfhjb
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInAppPurchaseGoodsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInAppPurchaseGoodsRequest) GoString() string {
	return s.String()
}

func (s *GetInAppPurchaseGoodsRequest) SetUserId(v string) *GetInAppPurchaseGoodsRequest {
	s.UserId = &v
	return s
}

type GetInAppPurchaseGoodsResponseBody struct {
	Result *GetInAppPurchaseGoodsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetInAppPurchaseGoodsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInAppPurchaseGoodsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInAppPurchaseGoodsResponseBody) SetResult(v *GetInAppPurchaseGoodsResponseBodyResult) *GetInAppPurchaseGoodsResponseBody {
	s.Result = v
	return s
}

type GetInAppPurchaseGoodsResponseBodyResult struct {
	// example:
	//
	// free
	OrderVersion      *string                                                     `json:"orderVersion,omitempty" xml:"orderVersion,omitempty"`
	PurchaseGoodsList []*GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList `json:"purchaseGoodsList,omitempty" xml:"purchaseGoodsList,omitempty" type:"Repeated"`
}

func (s GetInAppPurchaseGoodsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetInAppPurchaseGoodsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetInAppPurchaseGoodsResponseBodyResult) SetOrderVersion(v string) *GetInAppPurchaseGoodsResponseBodyResult {
	s.OrderVersion = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResult) SetPurchaseGoodsList(v []*GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) *GetInAppPurchaseGoodsResponseBodyResult {
	s.PurchaseGoodsList = v
	return s
}

type GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList struct {
	ApplicableVersion []*string `json:"applicableVersion,omitempty" xml:"applicableVersion,omitempty" type:"Repeated"`
	ApplyNum          *int64    `json:"applyNum,omitempty" xml:"applyNum,omitempty"`
	BelongIndustry    []*string `json:"belongIndustry,omitempty" xml:"belongIndustry,omitempty" type:"Repeated"`
	// example:
	//
	// psi
	GoodsId *string `json:"goodsId,omitempty" xml:"goodsId,omitempty"`
	// example:
	//
	// app_function
	GoodsType *string `json:"goodsType,omitempty" xml:"goodsType,omitempty"`
	// example:
	//
	// https://tungee-ali-crm.oss-cn-hangzhou.aliyuncs.com/app-center/icon/进销存.png
	Icon              *string                                                                    `json:"icon,omitempty" xml:"icon,omitempty"`
	MainOperationInfo *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo `json:"mainOperationInfo,omitempty" xml:"mainOperationInfo,omitempty" type:"Struct"`
	Media             []*GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMedia           `json:"media,omitempty" xml:"media,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	Price            *string                                                                   `json:"price,omitempty" xml:"price,omitempty"`
	SubOperationInfo *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo `json:"subOperationInfo,omitempty" xml:"subOperationInfo,omitempty" type:"Struct"`
	// example:
	//
	// 通过进销存管理，连接企业人、财、物，一站式解决进销存仓库管理难题。让货品成本有据可依，避免盲目采购；合理控制库存，防止滞销/脱销；通过往来对账确保资金安全。
	SubTitle *string   `json:"subTitle,omitempty" xml:"subTitle,omitempty"`
	Tag      []*string `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// example:
	//
	// 进销存
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) String() string {
	return tea.Prettify(s)
}

func (s GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) GoString() string {
	return s.String()
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetApplicableVersion(v []*string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.ApplicableVersion = v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetApplyNum(v int64) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.ApplyNum = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetBelongIndustry(v []*string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.BelongIndustry = v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetGoodsId(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.GoodsId = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetGoodsType(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.GoodsType = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetIcon(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.Icon = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetMainOperationInfo(v *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.MainOperationInfo = v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetMedia(v []*GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMedia) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.Media = v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetPrice(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.Price = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetSubOperationInfo(v *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.SubOperationInfo = v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetSubTitle(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.SubTitle = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetTag(v []*string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.Tag = v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList) SetTitle(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsList {
	s.Title = &v
	return s
}

type GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo struct {
	// example:
	//
	// GOODS-002
	GoodsCode *string `json:"goodsCode,omitempty" xml:"goodsCode,omitempty"`
	// example:
	//
	// https://yyy
	OriginalUrl *string `json:"originalUrl,omitempty" xml:"originalUrl,omitempty"`
	// example:
	//
	// https://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo) GoString() string {
	return s.String()
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo) SetGoodsCode(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo {
	s.GoodsCode = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo) SetOriginalUrl(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo {
	s.OriginalUrl = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo) SetUrl(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMainOperationInfo {
	s.Url = &v
	return s
}

type GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMedia struct {
	// example:
	//
	// image
	MediaType *string `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
	// example:
	//
	// https://tungee-ali-crm.oss-cn-hangzhou.aliyuncs.com/app-center/banner/进销存封面.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMedia) String() string {
	return tea.Prettify(s)
}

func (s GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMedia) GoString() string {
	return s.String()
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMedia) SetMediaType(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMedia {
	s.MediaType = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMedia) SetUrl(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListMedia {
	s.Url = &v
	return s
}

type GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo struct {
	// example:
	//
	// GOODS-2120
	GoodsCode *string `json:"goodsCode,omitempty" xml:"goodsCode,omitempty"`
	// example:
	//
	// https://yyy
	OriginalUrl *string `json:"originalUrl,omitempty" xml:"originalUrl,omitempty"`
	// example:
	//
	// https://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo) GoString() string {
	return s.String()
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo) SetGoodsCode(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo {
	s.GoodsCode = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo) SetOriginalUrl(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo {
	s.OriginalUrl = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo) SetUrl(v string) *GetInAppPurchaseGoodsResponseBodyResultPurchaseGoodsListSubOperationInfo {
	s.Url = &v
	return s
}

type GetInAppPurchaseGoodsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInAppPurchaseGoodsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInAppPurchaseGoodsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInAppPurchaseGoodsResponse) GoString() string {
	return s.String()
}

func (s *GetInAppPurchaseGoodsResponse) SetHeaders(v map[string]*string) *GetInAppPurchaseGoodsResponse {
	s.Headers = v
	return s
}

func (s *GetInAppPurchaseGoodsResponse) SetStatusCode(v int32) *GetInAppPurchaseGoodsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInAppPurchaseGoodsResponse) SetBody(v *GetInAppPurchaseGoodsResponseBody) *GetInAppPurchaseGoodsResponse {
	s.Body = v
	return s
}

type GetNavigationCatalogHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetNavigationCatalogHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetNavigationCatalogHeaders) GoString() string {
	return s.String()
}

func (s *GetNavigationCatalogHeaders) SetCommonHeaders(v map[string]*string) *GetNavigationCatalogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNavigationCatalogHeaders) SetXAcsDingtalkAccessToken(v string) *GetNavigationCatalogHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetNavigationCatalogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6360a371-4ffa-464b-a935-39817c3ccbe8
	BizTraceId *string `json:"bizTraceId,omitempty" xml:"bizTraceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sale
	Module *string `json:"module,omitempty" xml:"module,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16044739461008808747
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s GetNavigationCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNavigationCatalogRequest) GoString() string {
	return s.String()
}

func (s *GetNavigationCatalogRequest) SetBizTraceId(v string) *GetNavigationCatalogRequest {
	s.BizTraceId = &v
	return s
}

func (s *GetNavigationCatalogRequest) SetModule(v string) *GetNavigationCatalogRequest {
	s.Module = &v
	return s
}

func (s *GetNavigationCatalogRequest) SetOperatorUserId(v string) *GetNavigationCatalogRequest {
	s.OperatorUserId = &v
	return s
}

type GetNavigationCatalogResponseBody struct {
	Result *GetNavigationCatalogResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetNavigationCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNavigationCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *GetNavigationCatalogResponseBody) SetResult(v *GetNavigationCatalogResponseBodyResult) *GetNavigationCatalogResponseBody {
	s.Result = v
	return s
}

type GetNavigationCatalogResponseBodyResult struct {
	BizTraceId *string                                             `json:"bizTraceId,omitempty" xml:"bizTraceId,omitempty"`
	Module     *string                                             `json:"module,omitempty" xml:"module,omitempty"`
	NavCatalog []*GetNavigationCatalogResponseBodyResultNavCatalog `json:"navCatalog,omitempty" xml:"navCatalog,omitempty" type:"Repeated"`
}

func (s GetNavigationCatalogResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetNavigationCatalogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetNavigationCatalogResponseBodyResult) SetBizTraceId(v string) *GetNavigationCatalogResponseBodyResult {
	s.BizTraceId = &v
	return s
}

func (s *GetNavigationCatalogResponseBodyResult) SetModule(v string) *GetNavigationCatalogResponseBodyResult {
	s.Module = &v
	return s
}

func (s *GetNavigationCatalogResponseBodyResult) SetNavCatalog(v []*GetNavigationCatalogResponseBodyResultNavCatalog) *GetNavigationCatalogResponseBodyResult {
	s.NavCatalog = v
	return s
}

type GetNavigationCatalogResponseBodyResultNavCatalog struct {
	Children interface{} `json:"children,omitempty" xml:"children,omitempty"`
	NavCode  *string     `json:"navCode,omitempty" xml:"navCode,omitempty"`
	NavId    *string     `json:"navId,omitempty" xml:"navId,omitempty"`
	NavName  *string     `json:"navName,omitempty" xml:"navName,omitempty"`
	NavType  *string     `json:"navType,omitempty" xml:"navType,omitempty"`
}

func (s GetNavigationCatalogResponseBodyResultNavCatalog) String() string {
	return tea.Prettify(s)
}

func (s GetNavigationCatalogResponseBodyResultNavCatalog) GoString() string {
	return s.String()
}

func (s *GetNavigationCatalogResponseBodyResultNavCatalog) SetChildren(v interface{}) *GetNavigationCatalogResponseBodyResultNavCatalog {
	s.Children = v
	return s
}

func (s *GetNavigationCatalogResponseBodyResultNavCatalog) SetNavCode(v string) *GetNavigationCatalogResponseBodyResultNavCatalog {
	s.NavCode = &v
	return s
}

func (s *GetNavigationCatalogResponseBodyResultNavCatalog) SetNavId(v string) *GetNavigationCatalogResponseBodyResultNavCatalog {
	s.NavId = &v
	return s
}

func (s *GetNavigationCatalogResponseBodyResultNavCatalog) SetNavName(v string) *GetNavigationCatalogResponseBodyResultNavCatalog {
	s.NavName = &v
	return s
}

func (s *GetNavigationCatalogResponseBodyResultNavCatalog) SetNavType(v string) *GetNavigationCatalogResponseBodyResultNavCatalog {
	s.NavType = &v
	return s
}

type GetNavigationCatalogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNavigationCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNavigationCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNavigationCatalogResponse) GoString() string {
	return s.String()
}

func (s *GetNavigationCatalogResponse) SetHeaders(v map[string]*string) *GetNavigationCatalogResponse {
	s.Headers = v
	return s
}

func (s *GetNavigationCatalogResponse) SetStatusCode(v int32) *GetNavigationCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNavigationCatalogResponse) SetBody(v *GetNavigationCatalogResponseBody) *GetNavigationCatalogResponse {
	s.Body = v
	return s
}

type GetObjectDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetObjectDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetObjectDataHeaders) GoString() string {
	return s.String()
}

func (s *GetObjectDataHeaders) SetCommonHeaders(v map[string]*string) *GetObjectDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetObjectDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetObjectDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetObjectDataRequest struct {
	// example:
	//
	// ding_userid
	CurrentOperatorUserId *string `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-EF199CCA-8AB6-482A-AE10-85EDE5E391D9
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// {"queryGroupList":[{"logicType":"AND","queryObjectList":[{"fieldId":"contact_phone","value":"18000000000"},{"fieldId":"contact_related_customer","value":"INST-XXX"}]}]}
	QueryDsl *string `json:"queryDsl,omitempty" xml:"queryDsl,omitempty"`
}

func (s GetObjectDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetObjectDataRequest) GoString() string {
	return s.String()
}

func (s *GetObjectDataRequest) SetCurrentOperatorUserId(v string) *GetObjectDataRequest {
	s.CurrentOperatorUserId = &v
	return s
}

func (s *GetObjectDataRequest) SetMaxResults(v int64) *GetObjectDataRequest {
	s.MaxResults = &v
	return s
}

func (s *GetObjectDataRequest) SetName(v string) *GetObjectDataRequest {
	s.Name = &v
	return s
}

func (s *GetObjectDataRequest) SetNextToken(v string) *GetObjectDataRequest {
	s.NextToken = &v
	return s
}

func (s *GetObjectDataRequest) SetQueryDsl(v string) *GetObjectDataRequest {
	s.QueryDsl = &v
	return s
}

type GetObjectDataResponseBody struct {
	Result *GetObjectDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetObjectDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetObjectDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectDataResponseBody) SetResult(v *GetObjectDataResponseBodyResult) *GetObjectDataResponseBody {
	s.Result = v
	return s
}

type GetObjectDataResponseBodyResult struct {
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 100
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Values    []*GetObjectDataResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetObjectDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetObjectDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetObjectDataResponseBodyResult) SetHasMore(v bool) *GetObjectDataResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetObjectDataResponseBodyResult) SetMaxResults(v int64) *GetObjectDataResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *GetObjectDataResponseBodyResult) SetNextToken(v string) *GetObjectDataResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *GetObjectDataResponseBodyResult) SetValues(v []*GetObjectDataResponseBodyResultValues) *GetObjectDataResponseBodyResult {
	s.Values = v
	return s
}

type GetObjectDataResponseBodyResultValues struct {
	// example:
	//
	// 张xx
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// example:
	//
	// user01
	CreatorUserId *string                `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Data          map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData    map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// example:
	//
	// 2023-11-25 15:33:12
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2023-12-25 15:33:12
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// INST_XX
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// crm_contact
	ObjectType *string                                          `json:"objectType,omitempty" xml:"objectType,omitempty"`
	Permission *GetObjectDataResponseBodyResultValuesPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// example:
	//
	// COMPLETE
	ProcInstStatus *string `json:"procInstStatus,omitempty" xml:"procInstStatus,omitempty"`
	// example:
	//
	// agree
	ProcOutResult *string `json:"procOutResult,omitempty" xml:"procOutResult,omitempty"`
}

func (s GetObjectDataResponseBodyResultValues) String() string {
	return tea.Prettify(s)
}

func (s GetObjectDataResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *GetObjectDataResponseBodyResultValues) SetCreatorNick(v string) *GetObjectDataResponseBodyResultValues {
	s.CreatorNick = &v
	return s
}

func (s *GetObjectDataResponseBodyResultValues) SetCreatorUserId(v string) *GetObjectDataResponseBodyResultValues {
	s.CreatorUserId = &v
	return s
}

func (s *GetObjectDataResponseBodyResultValues) SetData(v map[string]interface{}) *GetObjectDataResponseBodyResultValues {
	s.Data = v
	return s
}

func (s *GetObjectDataResponseBodyResultValues) SetExtendData(v map[string]interface{}) *GetObjectDataResponseBodyResultValues {
	s.ExtendData = v
	return s
}

func (s *GetObjectDataResponseBodyResultValues) SetGmtCreate(v string) *GetObjectDataResponseBodyResultValues {
	s.GmtCreate = &v
	return s
}

func (s *GetObjectDataResponseBodyResultValues) SetGmtModified(v string) *GetObjectDataResponseBodyResultValues {
	s.GmtModified = &v
	return s
}

func (s *GetObjectDataResponseBodyResultValues) SetInstanceId(v string) *GetObjectDataResponseBodyResultValues {
	s.InstanceId = &v
	return s
}

func (s *GetObjectDataResponseBodyResultValues) SetObjectType(v string) *GetObjectDataResponseBodyResultValues {
	s.ObjectType = &v
	return s
}

func (s *GetObjectDataResponseBodyResultValues) SetPermission(v *GetObjectDataResponseBodyResultValuesPermission) *GetObjectDataResponseBodyResultValues {
	s.Permission = v
	return s
}

func (s *GetObjectDataResponseBodyResultValues) SetProcInstStatus(v string) *GetObjectDataResponseBodyResultValues {
	s.ProcInstStatus = &v
	return s
}

func (s *GetObjectDataResponseBodyResultValues) SetProcOutResult(v string) *GetObjectDataResponseBodyResultValues {
	s.ProcOutResult = &v
	return s
}

type GetObjectDataResponseBodyResultValuesPermission struct {
	OwnerUserIds       []*string `json:"ownerUserIds,omitempty" xml:"ownerUserIds,omitempty" type:"Repeated"`
	ParticipantUserIds []*string `json:"participantUserIds,omitempty" xml:"participantUserIds,omitempty" type:"Repeated"`
}

func (s GetObjectDataResponseBodyResultValuesPermission) String() string {
	return tea.Prettify(s)
}

func (s GetObjectDataResponseBodyResultValuesPermission) GoString() string {
	return s.String()
}

func (s *GetObjectDataResponseBodyResultValuesPermission) SetOwnerUserIds(v []*string) *GetObjectDataResponseBodyResultValuesPermission {
	s.OwnerUserIds = v
	return s
}

func (s *GetObjectDataResponseBodyResultValuesPermission) SetParticipantUserIds(v []*string) *GetObjectDataResponseBodyResultValuesPermission {
	s.ParticipantUserIds = v
	return s
}

type GetObjectDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetObjectDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetObjectDataResponse) GoString() string {
	return s.String()
}

func (s *GetObjectDataResponse) SetHeaders(v map[string]*string) *GetObjectDataResponse {
	s.Headers = v
	return s
}

func (s *GetObjectDataResponse) SetStatusCode(v int32) *GetObjectDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectDataResponse) SetBody(v *GetObjectDataResponseBody) *GetObjectDataResponse {
	s.Body = v
	return s
}

type GetOfficialAccountContactInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOfficialAccountContactInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactInfoHeaders) SetCommonHeaders(v map[string]*string) *GetOfficialAccountContactInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOfficialAccountContactInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetOfficialAccountContactInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOfficialAccountContactInfoResponseBody struct {
	AuthItems []*string `json:"authItems,omitempty" xml:"authItems,omitempty" type:"Repeated"`
	// example:
	//
	// 阿里巴巴钉钉
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
	// example:
	//
	// 18812341234
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// example:
	//
	// +86
	StateCode *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
	// example:
	//
	// unionId
	UnionId   *string   `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserInfos []*string `json:"userInfos,omitempty" xml:"userInfos,omitempty" type:"Repeated"`
}

func (s GetOfficialAccountContactInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactInfoResponseBody) SetAuthItems(v []*string) *GetOfficialAccountContactInfoResponseBody {
	s.AuthItems = v
	return s
}

func (s *GetOfficialAccountContactInfoResponseBody) SetCorpName(v string) *GetOfficialAccountContactInfoResponseBody {
	s.CorpName = &v
	return s
}

func (s *GetOfficialAccountContactInfoResponseBody) SetMobile(v string) *GetOfficialAccountContactInfoResponseBody {
	s.Mobile = &v
	return s
}

func (s *GetOfficialAccountContactInfoResponseBody) SetStateCode(v string) *GetOfficialAccountContactInfoResponseBody {
	s.StateCode = &v
	return s
}

func (s *GetOfficialAccountContactInfoResponseBody) SetUnionId(v string) *GetOfficialAccountContactInfoResponseBody {
	s.UnionId = &v
	return s
}

func (s *GetOfficialAccountContactInfoResponseBody) SetUserInfos(v []*string) *GetOfficialAccountContactInfoResponseBody {
	s.UserInfos = v
	return s
}

type GetOfficialAccountContactInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOfficialAccountContactInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOfficialAccountContactInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactInfoResponse) SetHeaders(v map[string]*string) *GetOfficialAccountContactInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOfficialAccountContactInfoResponse) SetStatusCode(v int32) *GetOfficialAccountContactInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOfficialAccountContactInfoResponse) SetBody(v *GetOfficialAccountContactInfoResponseBody) *GetOfficialAccountContactInfoResponse {
	s.Body = v
	return s
}

type GetOfficialAccountContactsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOfficialAccountContactsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsHeaders) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsHeaders) SetCommonHeaders(v map[string]*string) *GetOfficialAccountContactsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOfficialAccountContactsHeaders) SetXAcsDingtalkAccessToken(v string) *GetOfficialAccountContactsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOfficialAccountContactsRequest struct {
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
	// 123567
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetOfficialAccountContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsRequest) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsRequest) SetMaxResults(v int64) *GetOfficialAccountContactsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetOfficialAccountContactsRequest) SetNextToken(v string) *GetOfficialAccountContactsRequest {
	s.NextToken = &v
	return s
}

type GetOfficialAccountContactsResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 10010
	NextToken *string                                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Values    []*GetOfficialAccountContactsResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetOfficialAccountContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsResponseBody) SetMaxResults(v int64) *GetOfficialAccountContactsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBody) SetNextToken(v string) *GetOfficialAccountContactsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBody) SetValues(v []*GetOfficialAccountContactsResponseBodyValues) *GetOfficialAccountContactsResponseBody {
	s.Values = v
	return s
}

type GetOfficialAccountContactsResponseBodyValues struct {
	Contacts []*GetOfficialAccountContactsResponseBodyValuesContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// example:
	//
	// user_id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetOfficialAccountContactsResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsResponseBodyValues) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsResponseBodyValues) SetContacts(v []*GetOfficialAccountContactsResponseBodyValuesContacts) *GetOfficialAccountContactsResponseBodyValues {
	s.Contacts = v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValues) SetUserId(v string) *GetOfficialAccountContactsResponseBodyValues {
	s.UserId = &v
	return s
}

type GetOfficialAccountContactsResponseBodyValuesContacts struct {
	// example:
	//
	// 2019-12-25 15:33:12
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 张三
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// example:
	//
	// ding_userid
	CreatorUserId *string                `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Data          map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData    map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// example:
	//
	// instance_id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 2019-12-25 15:33:12
	ModifyTime *string                                                         `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	Permission *GetOfficialAccountContactsResponseBodyValuesContactsPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
}

func (s GetOfficialAccountContactsResponseBodyValuesContacts) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsResponseBodyValuesContacts) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetCreateTime(v string) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.CreateTime = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetCreatorNick(v string) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.CreatorNick = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetCreatorUserId(v string) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.CreatorUserId = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetData(v map[string]interface{}) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.Data = v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetExtendData(v map[string]interface{}) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.ExtendData = v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetInstanceId(v string) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.InstanceId = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetModifyTime(v string) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.ModifyTime = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetPermission(v *GetOfficialAccountContactsResponseBodyValuesContactsPermission) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.Permission = v
	return s
}

type GetOfficialAccountContactsResponseBodyValuesContactsPermission struct {
	OwnerStaffIds       []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s GetOfficialAccountContactsResponseBodyValuesContactsPermission) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsResponseBodyValuesContactsPermission) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsResponseBodyValuesContactsPermission) SetOwnerStaffIds(v []*string) *GetOfficialAccountContactsResponseBodyValuesContactsPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContactsPermission) SetParticipantStaffIds(v []*string) *GetOfficialAccountContactsResponseBodyValuesContactsPermission {
	s.ParticipantStaffIds = v
	return s
}

type GetOfficialAccountContactsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOfficialAccountContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOfficialAccountContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsResponse) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsResponse) SetHeaders(v map[string]*string) *GetOfficialAccountContactsResponse {
	s.Headers = v
	return s
}

func (s *GetOfficialAccountContactsResponse) SetStatusCode(v int32) *GetOfficialAccountContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOfficialAccountContactsResponse) SetBody(v *GetOfficialAccountContactsResponseBody) *GetOfficialAccountContactsResponse {
	s.Body = v
	return s
}

type GetOfficialAccountOTOMessageResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOfficialAccountOTOMessageResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountOTOMessageResultHeaders) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountOTOMessageResultHeaders) SetCommonHeaders(v map[string]*string) *GetOfficialAccountOTOMessageResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOfficialAccountOTOMessageResultHeaders) SetXAcsDingtalkAccessToken(v string) *GetOfficialAccountOTOMessageResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOfficialAccountOTOMessageResultRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s GetOfficialAccountOTOMessageResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountOTOMessageResultRequest) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountOTOMessageResultRequest) SetAccountId(v string) *GetOfficialAccountOTOMessageResultRequest {
	s.AccountId = &v
	return s
}

func (s *GetOfficialAccountOTOMessageResultRequest) SetOpenPushId(v string) *GetOfficialAccountOTOMessageResultRequest {
	s.OpenPushId = &v
	return s
}

type GetOfficialAccountOTOMessageResultResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Result *GetOfficialAccountOTOMessageResultResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetOfficialAccountOTOMessageResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountOTOMessageResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountOTOMessageResultResponseBody) SetRequestId(v string) *GetOfficialAccountOTOMessageResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOfficialAccountOTOMessageResultResponseBody) SetResult(v *GetOfficialAccountOTOMessageResultResponseBodyResult) *GetOfficialAccountOTOMessageResultResponseBody {
	s.Result = v
	return s
}

type GetOfficialAccountOTOMessageResultResponseBodyResult struct {
	// This parameter is required.
	ReadUserIdList []*string `json:"readUserIdList,omitempty" xml:"readUserIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetOfficialAccountOTOMessageResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountOTOMessageResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountOTOMessageResultResponseBodyResult) SetReadUserIdList(v []*string) *GetOfficialAccountOTOMessageResultResponseBodyResult {
	s.ReadUserIdList = v
	return s
}

func (s *GetOfficialAccountOTOMessageResultResponseBodyResult) SetStatus(v int64) *GetOfficialAccountOTOMessageResultResponseBodyResult {
	s.Status = &v
	return s
}

type GetOfficialAccountOTOMessageResultResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOfficialAccountOTOMessageResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOfficialAccountOTOMessageResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountOTOMessageResultResponse) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountOTOMessageResultResponse) SetHeaders(v map[string]*string) *GetOfficialAccountOTOMessageResultResponse {
	s.Headers = v
	return s
}

func (s *GetOfficialAccountOTOMessageResultResponse) SetStatusCode(v int32) *GetOfficialAccountOTOMessageResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOfficialAccountOTOMessageResultResponse) SetBody(v *GetOfficialAccountOTOMessageResultResponseBody) *GetOfficialAccountOTOMessageResultResponse {
	s.Body = v
	return s
}

type GetRelatedViewTabDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRelatedViewTabDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabDataHeaders) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabDataHeaders) SetCommonHeaders(v map[string]*string) *GetRelatedViewTabDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRelatedViewTabDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetRelatedViewTabDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRelatedViewTabDataRequest struct {
	// example:
	//
	// PROC-62829702-A377-42A9-9CB3-E1C691A0CEDB
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// OpenDataField_OV2K4SOW2ZGG
	RelatedField *string `json:"relatedField,omitempty" xml:"relatedField,omitempty"`
	// example:
	//
	// u_dxcugzT0aPQvcK2PIkzQ00841721291058
	RelatedInstId *string `json:"relatedInstId,omitempty" xml:"relatedInstId,omitempty"`
	// example:
	//
	// manager6034
	ViewUserId *string `json:"viewUserId,omitempty" xml:"viewUserId,omitempty"`
}

func (s GetRelatedViewTabDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabDataRequest) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabDataRequest) SetFormCode(v string) *GetRelatedViewTabDataRequest {
	s.FormCode = &v
	return s
}

func (s *GetRelatedViewTabDataRequest) SetMaxResults(v int32) *GetRelatedViewTabDataRequest {
	s.MaxResults = &v
	return s
}

func (s *GetRelatedViewTabDataRequest) SetNextToken(v int64) *GetRelatedViewTabDataRequest {
	s.NextToken = &v
	return s
}

func (s *GetRelatedViewTabDataRequest) SetRelatedField(v string) *GetRelatedViewTabDataRequest {
	s.RelatedField = &v
	return s
}

func (s *GetRelatedViewTabDataRequest) SetRelatedInstId(v string) *GetRelatedViewTabDataRequest {
	s.RelatedInstId = &v
	return s
}

func (s *GetRelatedViewTabDataRequest) SetViewUserId(v string) *GetRelatedViewTabDataRequest {
	s.ViewUserId = &v
	return s
}

type GetRelatedViewTabDataResponseBody struct {
	Result *GetRelatedViewTabDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetRelatedViewTabDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabDataResponseBody) SetResult(v *GetRelatedViewTabDataResponseBodyResult) *GetRelatedViewTabDataResponseBody {
	s.Result = v
	return s
}

type GetRelatedViewTabDataResponseBodyResult struct {
	Page *GetRelatedViewTabDataResponseBodyResultPage `json:"page,omitempty" xml:"page,omitempty" type:"Struct"`
}

func (s GetRelatedViewTabDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabDataResponseBodyResult) SetPage(v *GetRelatedViewTabDataResponseBodyResultPage) *GetRelatedViewTabDataResponseBodyResult {
	s.Page = v
	return s
}

type GetRelatedViewTabDataResponseBodyResultPage struct {
	HasMore *bool                                              `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*GetRelatedViewTabDataResponseBodyResultPageList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 5
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetRelatedViewTabDataResponseBodyResultPage) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabDataResponseBodyResultPage) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabDataResponseBodyResultPage) SetHasMore(v bool) *GetRelatedViewTabDataResponseBodyResultPage {
	s.HasMore = &v
	return s
}

func (s *GetRelatedViewTabDataResponseBodyResultPage) SetList(v []*GetRelatedViewTabDataResponseBodyResultPageList) *GetRelatedViewTabDataResponseBodyResultPage {
	s.List = v
	return s
}

func (s *GetRelatedViewTabDataResponseBodyResultPage) SetNextToken(v int64) *GetRelatedViewTabDataResponseBodyResultPage {
	s.NextToken = &v
	return s
}

func (s *GetRelatedViewTabDataResponseBodyResultPage) SetTotalCount(v int64) *GetRelatedViewTabDataResponseBodyResultPage {
	s.TotalCount = &v
	return s
}

type GetRelatedViewTabDataResponseBodyResultPageList struct {
	// example:
	//
	// 西游四人组:孙悟空
	AbstractMessage *string `json:"abstractMessage,omitempty" xml:"abstractMessage,omitempty"`
	CreateTime      *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 王凯提交的楚衣的流程表单2
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetRelatedViewTabDataResponseBodyResultPageList) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabDataResponseBodyResultPageList) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabDataResponseBodyResultPageList) SetAbstractMessage(v string) *GetRelatedViewTabDataResponseBodyResultPageList {
	s.AbstractMessage = &v
	return s
}

func (s *GetRelatedViewTabDataResponseBodyResultPageList) SetCreateTime(v int64) *GetRelatedViewTabDataResponseBodyResultPageList {
	s.CreateTime = &v
	return s
}

func (s *GetRelatedViewTabDataResponseBodyResultPageList) SetTitle(v string) *GetRelatedViewTabDataResponseBodyResultPageList {
	s.Title = &v
	return s
}

type GetRelatedViewTabDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRelatedViewTabDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRelatedViewTabDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabDataResponse) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabDataResponse) SetHeaders(v map[string]*string) *GetRelatedViewTabDataResponse {
	s.Headers = v
	return s
}

func (s *GetRelatedViewTabDataResponse) SetStatusCode(v int32) *GetRelatedViewTabDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRelatedViewTabDataResponse) SetBody(v *GetRelatedViewTabDataResponseBody) *GetRelatedViewTabDataResponse {
	s.Body = v
	return s
}

type GetRelatedViewTabMetaHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRelatedViewTabMetaHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabMetaHeaders) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabMetaHeaders) SetCommonHeaders(v map[string]*string) *GetRelatedViewTabMetaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRelatedViewTabMetaHeaders) SetXAcsDingtalkAccessToken(v string) *GetRelatedViewTabMetaHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRelatedViewTabMetaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PROC-2DB0FF86-CE29-41FF-B0FE-BFDE5BD9A0C1
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	ViewUserId *string `json:"viewUserId,omitempty" xml:"viewUserId,omitempty"`
}

func (s GetRelatedViewTabMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabMetaRequest) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabMetaRequest) SetFormCode(v string) *GetRelatedViewTabMetaRequest {
	s.FormCode = &v
	return s
}

func (s *GetRelatedViewTabMetaRequest) SetViewUserId(v string) *GetRelatedViewTabMetaRequest {
	s.ViewUserId = &v
	return s
}

type GetRelatedViewTabMetaResponseBody struct {
	Results []*GetRelatedViewTabMetaResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s GetRelatedViewTabMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabMetaResponseBody) SetResults(v []*GetRelatedViewTabMetaResponseBodyResults) *GetRelatedViewTabMetaResponseBody {
	s.Results = v
	return s
}

type GetRelatedViewTabMetaResponseBodyResults struct {
	// example:
	//
	// PROC-4EFE895D-A291-4A65-9FD6-99431604DF67
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// example:
	//
	// OpenDataField_K99RPMMRGJ40
	RelateComponentId *string `json:"relateComponentId,omitempty" xml:"relateComponentId,omitempty"`
	// example:
	//
	// 212
	TabTitle *string `json:"tabTitle,omitempty" xml:"tabTitle,omitempty"`
}

func (s GetRelatedViewTabMetaResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabMetaResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabMetaResponseBodyResults) SetFormCode(v string) *GetRelatedViewTabMetaResponseBodyResults {
	s.FormCode = &v
	return s
}

func (s *GetRelatedViewTabMetaResponseBodyResults) SetRelateComponentId(v string) *GetRelatedViewTabMetaResponseBodyResults {
	s.RelateComponentId = &v
	return s
}

func (s *GetRelatedViewTabMetaResponseBodyResults) SetTabTitle(v string) *GetRelatedViewTabMetaResponseBodyResults {
	s.TabTitle = &v
	return s
}

type GetRelatedViewTabMetaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRelatedViewTabMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRelatedViewTabMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedViewTabMetaResponse) GoString() string {
	return s.String()
}

func (s *GetRelatedViewTabMetaResponse) SetHeaders(v map[string]*string) *GetRelatedViewTabMetaResponse {
	s.Headers = v
	return s
}

func (s *GetRelatedViewTabMetaResponse) SetStatusCode(v int32) *GetRelatedViewTabMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRelatedViewTabMetaResponse) SetBody(v *GetRelatedViewTabMetaResponseBody) *GetRelatedViewTabMetaResponse {
	s.Body = v
	return s
}

type GetRelationUkSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRelationUkSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingHeaders) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingHeaders) SetCommonHeaders(v map[string]*string) *GetRelationUkSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRelationUkSettingHeaders) SetXAcsDingtalkAccessToken(v string) *GetRelationUkSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRelationUkSettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// crm_customer
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s GetRelationUkSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingRequest) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingRequest) SetRelationType(v string) *GetRelationUkSettingRequest {
	s.RelationType = &v
	return s
}

type GetRelationUkSettingResponseBody struct {
	// example:
	//
	// true
	Result []*GetRelationUkSettingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetRelationUkSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingResponseBody) SetResult(v []*GetRelationUkSettingResponseBodyResult) *GetRelationUkSettingResponseBody {
	s.Result = v
	return s
}

type GetRelationUkSettingResponseBodyResult struct {
	// example:
	//
	// customer_name
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField_U2K5A122
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
}

func (s GetRelationUkSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingResponseBodyResult) SetBizAlias(v string) *GetRelationUkSettingResponseBodyResult {
	s.BizAlias = &v
	return s
}

func (s *GetRelationUkSettingResponseBodyResult) SetFieldId(v string) *GetRelationUkSettingResponseBodyResult {
	s.FieldId = &v
	return s
}

type GetRelationUkSettingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRelationUkSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRelationUkSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingResponse) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingResponse) SetHeaders(v map[string]*string) *GetRelationUkSettingResponse {
	s.Headers = v
	return s
}

func (s *GetRelationUkSettingResponse) SetStatusCode(v int32) *GetRelationUkSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRelationUkSettingResponse) SetBody(v *GetRelationUkSettingResponseBody) *GetRelationUkSettingResponse {
	s.Body = v
	return s
}

type JoinGroupSetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s JoinGroupSetHeaders) String() string {
	return tea.Prettify(s)
}

func (s JoinGroupSetHeaders) GoString() string {
	return s.String()
}

func (s *JoinGroupSetHeaders) SetCommonHeaders(v map[string]*string) *JoinGroupSetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *JoinGroupSetHeaders) SetXAcsDingtalkAccessToken(v string) *JoinGroupSetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type JoinGroupSetRequest struct {
	BizDataList []*JoinGroupSetRequestBizDataList `json:"bizDataList,omitempty" xml:"bizDataList,omitempty" type:"Repeated"`
	// This parameter is required.
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc123
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc123
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s JoinGroupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinGroupSetRequest) GoString() string {
	return s.String()
}

func (s *JoinGroupSetRequest) SetBizDataList(v []*JoinGroupSetRequestBizDataList) *JoinGroupSetRequest {
	s.BizDataList = v
	return s
}

func (s *JoinGroupSetRequest) SetCorpId(v string) *JoinGroupSetRequest {
	s.CorpId = &v
	return s
}

func (s *JoinGroupSetRequest) SetOpenGroupSetId(v string) *JoinGroupSetRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *JoinGroupSetRequest) SetUnionId(v string) *JoinGroupSetRequest {
	s.UnionId = &v
	return s
}

type JoinGroupSetRequestBizDataList struct {
	// example:
	//
	// {}
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// example:
	//
	// customer_name
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// abc123
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s JoinGroupSetRequestBizDataList) String() string {
	return tea.Prettify(s)
}

func (s JoinGroupSetRequestBizDataList) GoString() string {
	return s.String()
}

func (s *JoinGroupSetRequestBizDataList) SetExtendValue(v string) *JoinGroupSetRequestBizDataList {
	s.ExtendValue = &v
	return s
}

func (s *JoinGroupSetRequestBizDataList) SetKey(v string) *JoinGroupSetRequestBizDataList {
	s.Key = &v
	return s
}

func (s *JoinGroupSetRequestBizDataList) SetValue(v string) *JoinGroupSetRequestBizDataList {
	s.Value = &v
	return s
}

type JoinGroupSetResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// abc123
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc123
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s JoinGroupSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinGroupSetResponseBody) GoString() string {
	return s.String()
}

func (s *JoinGroupSetResponseBody) SetChatId(v string) *JoinGroupSetResponseBody {
	s.ChatId = &v
	return s
}

func (s *JoinGroupSetResponseBody) SetOpenConversationId(v string) *JoinGroupSetResponseBody {
	s.OpenConversationId = &v
	return s
}

type JoinGroupSetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinGroupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinGroupSetResponse) GoString() string {
	return s.String()
}

func (s *JoinGroupSetResponse) SetHeaders(v map[string]*string) *JoinGroupSetResponse {
	s.Headers = v
	return s
}

func (s *JoinGroupSetResponse) SetStatusCode(v int32) *JoinGroupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinGroupSetResponse) SetBody(v *JoinGroupSetResponseBody) *JoinGroupSetResponse {
	s.Body = v
	return s
}

type ListAvailableBenefitHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAvailableBenefitHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBenefitHeaders) GoString() string {
	return s.String()
}

func (s *ListAvailableBenefitHeaders) SetCommonHeaders(v map[string]*string) *ListAvailableBenefitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAvailableBenefitHeaders) SetXAcsDingtalkAccessToken(v string) *ListAvailableBenefitHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAvailableBenefitRequest struct {
	// This parameter is required.
	BenefitCodeList []*string `json:"benefitCodeList,omitempty" xml:"benefitCodeList,omitempty" type:"Repeated"`
}

func (s ListAvailableBenefitRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBenefitRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableBenefitRequest) SetBenefitCodeList(v []*string) *ListAvailableBenefitRequest {
	s.BenefitCodeList = v
	return s
}

type ListAvailableBenefitResponseBody struct {
	// This parameter is required.
	Result []*ListAvailableBenefitResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAvailableBenefitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBenefitResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableBenefitResponseBody) SetResult(v []*ListAvailableBenefitResponseBodyResult) *ListAvailableBenefitResponseBody {
	s.Result = v
	return s
}

type ListAvailableBenefitResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// B_ACCOUNT_NUMBER
	BenefitCode *string `json:"benefitCode,omitempty" xml:"benefitCode,omitempty"`
	// example:
	//
	// 1718696461851
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 200
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// example:
	//
	// 1718696461851
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 10
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	// example:
	//
	// tryout
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 试用版
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty"`
}

func (s ListAvailableBenefitResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBenefitResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAvailableBenefitResponseBodyResult) SetBenefitCode(v string) *ListAvailableBenefitResponseBodyResult {
	s.BenefitCode = &v
	return s
}

func (s *ListAvailableBenefitResponseBodyResult) SetEndTime(v int64) *ListAvailableBenefitResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *ListAvailableBenefitResponseBodyResult) SetQuota(v int64) *ListAvailableBenefitResponseBodyResult {
	s.Quota = &v
	return s
}

func (s *ListAvailableBenefitResponseBodyResult) SetStartTime(v int64) *ListAvailableBenefitResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *ListAvailableBenefitResponseBodyResult) SetUsedQuota(v int64) *ListAvailableBenefitResponseBodyResult {
	s.UsedQuota = &v
	return s
}

func (s *ListAvailableBenefitResponseBodyResult) SetVersion(v string) *ListAvailableBenefitResponseBodyResult {
	s.Version = &v
	return s
}

func (s *ListAvailableBenefitResponseBodyResult) SetVersionName(v string) *ListAvailableBenefitResponseBodyResult {
	s.VersionName = &v
	return s
}

type ListAvailableBenefitResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableBenefitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableBenefitResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBenefitResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableBenefitResponse) SetHeaders(v map[string]*string) *ListAvailableBenefitResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableBenefitResponse) SetStatusCode(v int32) *ListAvailableBenefitResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableBenefitResponse) SetBody(v *ListAvailableBenefitResponseBody) *ListAvailableBenefitResponse {
	s.Body = v
	return s
}

type ListBenefitLicenseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListBenefitLicenseHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListBenefitLicenseHeaders) GoString() string {
	return s.String()
}

func (s *ListBenefitLicenseHeaders) SetCommonHeaders(v map[string]*string) *ListBenefitLicenseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListBenefitLicenseHeaders) SetXAcsDingtalkAccessToken(v string) *ListBenefitLicenseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListBenefitLicenseRequest struct {
	// This parameter is required.
	Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s ListBenefitLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBenefitLicenseRequest) GoString() string {
	return s.String()
}

func (s *ListBenefitLicenseRequest) SetDomains(v []*string) *ListBenefitLicenseRequest {
	s.Domains = v
	return s
}

type ListBenefitLicenseResponseBody struct {
	// This parameter is required.
	Result []*ListBenefitLicenseResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListBenefitLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBenefitLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *ListBenefitLicenseResponseBody) SetResult(v []*ListBenefitLicenseResponseBodyResult) *ListBenefitLicenseResponseBody {
	s.Result = v
	return s
}

type ListBenefitLicenseResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// B_ACCOUNT_NUMBER
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// This parameter is required.
	Licenses []*ListBenefitLicenseResponseBodyResultLicenses `json:"licenses,omitempty" xml:"licenses,omitempty" type:"Repeated"`
}

func (s ListBenefitLicenseResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListBenefitLicenseResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListBenefitLicenseResponseBodyResult) SetDomain(v string) *ListBenefitLicenseResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListBenefitLicenseResponseBodyResult) SetLicenses(v []*ListBenefitLicenseResponseBodyResultLicenses) *ListBenefitLicenseResponseBodyResult {
	s.Licenses = v
	return s
}

type ListBenefitLicenseResponseBodyResultLicenses struct {
	// This parameter is required.
	//
	// example:
	//
	// staffId
	LicenseUserId *string `json:"licenseUserId,omitempty" xml:"licenseUserId,omitempty"`
}

func (s ListBenefitLicenseResponseBodyResultLicenses) String() string {
	return tea.Prettify(s)
}

func (s ListBenefitLicenseResponseBodyResultLicenses) GoString() string {
	return s.String()
}

func (s *ListBenefitLicenseResponseBodyResultLicenses) SetLicenseUserId(v string) *ListBenefitLicenseResponseBodyResultLicenses {
	s.LicenseUserId = &v
	return s
}

type ListBenefitLicenseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBenefitLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBenefitLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBenefitLicenseResponse) GoString() string {
	return s.String()
}

func (s *ListBenefitLicenseResponse) SetHeaders(v map[string]*string) *ListBenefitLicenseResponse {
	s.Headers = v
	return s
}

func (s *ListBenefitLicenseResponse) SetStatusCode(v int32) *ListBenefitLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBenefitLicenseResponse) SetBody(v *ListBenefitLicenseResponseBody) *ListBenefitLicenseResponse {
	s.Body = v
	return s
}

type ListClueTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListClueTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListClueTagHeaders) GoString() string {
	return s.String()
}

func (s *ListClueTagHeaders) SetCommonHeaders(v map[string]*string) *ListClueTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListClueTagHeaders) SetXAcsDingtalkAccessToken(v string) *ListClueTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListClueTagResponseBody struct {
	Result []*ListClueTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListClueTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClueTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListClueTagResponseBody) SetResult(v []*ListClueTagResponseBodyResult) *ListClueTagResponseBody {
	s.Result = v
	return s
}

type ListClueTagResponseBodyResult struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	TagId *string `json:"tagId,omitempty" xml:"tagId,omitempty"`
}

func (s ListClueTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListClueTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClueTagResponseBodyResult) SetName(v string) *ListClueTagResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListClueTagResponseBodyResult) SetTagId(v string) *ListClueTagResponseBodyResult {
	s.TagId = &v
	return s
}

type ListClueTagResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClueTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClueTagResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClueTagResponse) GoString() string {
	return s.String()
}

func (s *ListClueTagResponse) SetHeaders(v map[string]*string) *ListClueTagResponse {
	s.Headers = v
	return s
}

func (s *ListClueTagResponse) SetStatusCode(v int32) *ListClueTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClueTagResponse) SetBody(v *ListClueTagResponseBody) *ListClueTagResponse {
	s.Body = v
	return s
}

type ListCrmPersonalCustomersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListCrmPersonalCustomersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersHeaders) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersHeaders) SetCommonHeaders(v map[string]*string) *ListCrmPersonalCustomersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCrmPersonalCustomersHeaders) SetXAcsDingtalkAccessToken(v string) *ListCrmPersonalCustomersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListCrmPersonalCustomersRequest struct {
	// This parameter is required.
	Body                  []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	CurrentOperatorUserId *string   `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
	RelationType          *string   `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s ListCrmPersonalCustomersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersRequest) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersRequest) SetBody(v []*string) *ListCrmPersonalCustomersRequest {
	s.Body = v
	return s
}

func (s *ListCrmPersonalCustomersRequest) SetCurrentOperatorUserId(v string) *ListCrmPersonalCustomersRequest {
	s.CurrentOperatorUserId = &v
	return s
}

func (s *ListCrmPersonalCustomersRequest) SetRelationType(v string) *ListCrmPersonalCustomersRequest {
	s.RelationType = &v
	return s
}

type ListCrmPersonalCustomersResponseBody struct {
	// This parameter is required.
	Result []*ListCrmPersonalCustomersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListCrmPersonalCustomersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersResponseBody) SetResult(v []*ListCrmPersonalCustomersResponseBodyResult) *ListCrmPersonalCustomersResponseBody {
	s.Result = v
	return s
}

type ListCrmPersonalCustomersResponseBodyResult struct {
	// This parameter is required.
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// This parameter is required.
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// This parameter is required.
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// This parameter is required.
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// This parameter is required.
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// This parameter is required.
	Permission *ListCrmPersonalCustomersResponseBodyResultPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// This parameter is required.
	ProcInstStatus *string `json:"procInstStatus,omitempty" xml:"procInstStatus,omitempty"`
	// This parameter is required.
	ProcOutResult *string `json:"procOutResult,omitempty" xml:"procOutResult,omitempty"`
}

func (s ListCrmPersonalCustomersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetAppUuid(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.AppUuid = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetCreatorNick(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.CreatorNick = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetCreatorUserId(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetData(v map[string]interface{}) *ListCrmPersonalCustomersResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetExtendData(v map[string]interface{}) *ListCrmPersonalCustomersResponseBodyResult {
	s.ExtendData = v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetFormCode(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.FormCode = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetGmtCreate(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetGmtModified(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetInstanceId(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetObjectType(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.ObjectType = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetPermission(v *ListCrmPersonalCustomersResponseBodyResultPermission) *ListCrmPersonalCustomersResponseBodyResult {
	s.Permission = v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetProcInstStatus(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.ProcInstStatus = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetProcOutResult(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.ProcOutResult = &v
	return s
}

type ListCrmPersonalCustomersResponseBodyResultPermission struct {
	// This parameter is required.
	OwnerStaffIds []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	// This parameter is required.
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s ListCrmPersonalCustomersResponseBodyResultPermission) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersResponseBodyResultPermission) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersResponseBodyResultPermission) SetOwnerStaffIds(v []*string) *ListCrmPersonalCustomersResponseBodyResultPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResultPermission) SetParticipantStaffIds(v []*string) *ListCrmPersonalCustomersResponseBodyResultPermission {
	s.ParticipantStaffIds = v
	return s
}

type ListCrmPersonalCustomersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrmPersonalCustomersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrmPersonalCustomersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersResponse) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersResponse) SetHeaders(v map[string]*string) *ListCrmPersonalCustomersResponse {
	s.Headers = v
	return s
}

func (s *ListCrmPersonalCustomersResponse) SetStatusCode(v int32) *ListCrmPersonalCustomersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrmPersonalCustomersResponse) SetBody(v *ListCrmPersonalCustomersResponseBody) *ListCrmPersonalCustomersResponse {
	s.Body = v
	return s
}

type ListGroupSetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListGroupSetHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSetHeaders) GoString() string {
	return s.String()
}

func (s *ListGroupSetHeaders) SetCommonHeaders(v map[string]*string) *ListGroupSetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListGroupSetHeaders) SetXAcsDingtalkAccessToken(v string) *ListGroupSetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListGroupSetRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	QueryDsl   *string `json:"queryDsl,omitempty" xml:"queryDsl,omitempty"`
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s ListGroupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSetRequest) GoString() string {
	return s.String()
}

func (s *ListGroupSetRequest) SetMaxResults(v int32) *ListGroupSetRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupSetRequest) SetNextToken(v string) *ListGroupSetRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupSetRequest) SetQueryDsl(v string) *ListGroupSetRequest {
	s.QueryDsl = &v
	return s
}

func (s *ListGroupSetRequest) SetRelationType(v string) *ListGroupSetRequest {
	s.RelationType = &v
	return s
}

type ListGroupSetResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// fasfasd
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	ResultList []*ListGroupSetResponseBodyResultList `json:"resultList,omitempty" xml:"resultList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListGroupSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSetResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupSetResponseBody) SetHasMore(v bool) *ListGroupSetResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListGroupSetResponseBody) SetNextToken(v string) *ListGroupSetResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupSetResponseBody) SetResultList(v []*ListGroupSetResponseBodyResultList) *ListGroupSetResponseBody {
	s.ResultList = v
	return s
}

func (s *ListGroupSetResponseBody) SetTotalCount(v int32) *ListGroupSetResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupSetResponseBodyResultList struct {
	// example:
	//
	// 2021-12-23T13:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2021-12-23T13:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 10
	GroupChatCount *int32 `json:"groupChatCount,omitempty" xml:"groupChatCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123agsg
	LastOpenConversationId *string `json:"lastOpenConversationId,omitempty" xml:"lastOpenConversationId,omitempty"`
	// This parameter is required.
	Manager []*ListGroupSetResponseBodyResultListManager `json:"manager,omitempty" xml:"manager,omitempty" type:"Repeated"`
	// example:
	//
	// afsd12,afsd13
	ManagerUserIds *string `json:"managerUserIds,omitempty" xml:"managerUserIds,omitempty"`
	// example:
	//
	// 10
	MemberCount *int32 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// example:
	//
	// 100
	MemberQuota *int32 `json:"memberQuota,omitempty" xml:"memberQuota,omitempty"`
	// example:
	//
	// 营销群
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 群公告
	Notice *string `json:"notice,omitempty" xml:"notice,omitempty"`
	// example:
	//
	// 0
	NoticeToped *int32 `json:"noticeToped,omitempty" xml:"noticeToped,omitempty"`
	// example:
	//
	// adfads
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	Owner *ListGroupSetResponseBodyResultListOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// example:
	//
	// afsd12
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// example:
	//
	// crm_customer_personal
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// example:
	//
	// sfasgsab
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s ListGroupSetResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSetResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *ListGroupSetResponseBodyResultList) SetGmtCreate(v string) *ListGroupSetResponseBodyResultList {
	s.GmtCreate = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetGmtModified(v string) *ListGroupSetResponseBodyResultList {
	s.GmtModified = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetGroupChatCount(v int32) *ListGroupSetResponseBodyResultList {
	s.GroupChatCount = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetLastOpenConversationId(v string) *ListGroupSetResponseBodyResultList {
	s.LastOpenConversationId = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetManager(v []*ListGroupSetResponseBodyResultListManager) *ListGroupSetResponseBodyResultList {
	s.Manager = v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetManagerUserIds(v string) *ListGroupSetResponseBodyResultList {
	s.ManagerUserIds = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetMemberCount(v int32) *ListGroupSetResponseBodyResultList {
	s.MemberCount = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetMemberQuota(v int32) *ListGroupSetResponseBodyResultList {
	s.MemberQuota = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetName(v string) *ListGroupSetResponseBodyResultList {
	s.Name = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetNotice(v string) *ListGroupSetResponseBodyResultList {
	s.Notice = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetNoticeToped(v int32) *ListGroupSetResponseBodyResultList {
	s.NoticeToped = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetOpenGroupSetId(v string) *ListGroupSetResponseBodyResultList {
	s.OpenGroupSetId = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetOwner(v *ListGroupSetResponseBodyResultListOwner) *ListGroupSetResponseBodyResultList {
	s.Owner = v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetOwnerUserId(v string) *ListGroupSetResponseBodyResultList {
	s.OwnerUserId = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetRelationType(v string) *ListGroupSetResponseBodyResultList {
	s.RelationType = &v
	return s
}

func (s *ListGroupSetResponseBodyResultList) SetTemplateId(v string) *ListGroupSetResponseBodyResultList {
	s.TemplateId = &v
	return s
}

type ListGroupSetResponseBodyResultListManager struct {
	// example:
	//
	// XX
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// afs1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListGroupSetResponseBodyResultListManager) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSetResponseBodyResultListManager) GoString() string {
	return s.String()
}

func (s *ListGroupSetResponseBodyResultListManager) SetName(v string) *ListGroupSetResponseBodyResultListManager {
	s.Name = &v
	return s
}

func (s *ListGroupSetResponseBodyResultListManager) SetUserId(v string) *ListGroupSetResponseBodyResultListManager {
	s.UserId = &v
	return s
}

type ListGroupSetResponseBodyResultListOwner struct {
	// example:
	//
	// XX
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// afsd12
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListGroupSetResponseBodyResultListOwner) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSetResponseBodyResultListOwner) GoString() string {
	return s.String()
}

func (s *ListGroupSetResponseBodyResultListOwner) SetName(v string) *ListGroupSetResponseBodyResultListOwner {
	s.Name = &v
	return s
}

func (s *ListGroupSetResponseBodyResultListOwner) SetUserId(v string) *ListGroupSetResponseBodyResultListOwner {
	s.UserId = &v
	return s
}

type ListGroupSetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSetResponse) GoString() string {
	return s.String()
}

func (s *ListGroupSetResponse) SetHeaders(v map[string]*string) *ListGroupSetResponse {
	s.Headers = v
	return s
}

func (s *ListGroupSetResponse) SetStatusCode(v int32) *ListGroupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupSetResponse) SetBody(v *ListGroupSetResponseBody) *ListGroupSetResponse {
	s.Body = v
	return s
}

type OverrideUpdateCustomerDataAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OverrideUpdateCustomerDataAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s OverrideUpdateCustomerDataAuthHeaders) GoString() string {
	return s.String()
}

func (s *OverrideUpdateCustomerDataAuthHeaders) SetCommonHeaders(v map[string]*string) *OverrideUpdateCustomerDataAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OverrideUpdateCustomerDataAuthHeaders) SetXAcsDingtalkAccessToken(v string) *OverrideUpdateCustomerDataAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OverrideUpdateCustomerDataAuthRequest struct {
	// This parameter is required.
	CustomerIds []*string `json:"customerIds,omitempty" xml:"customerIds,omitempty" type:"Repeated"`
	// This parameter is required.
	DataAuthUserIds []*string `json:"dataAuthUserIds,omitempty" xml:"dataAuthUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// PROC-98187D45-EFC0-4FC4-887E-45BD24209D69
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// staffId2
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	// example:
	//
	// crm_customer
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// owner
	RoleType *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s OverrideUpdateCustomerDataAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s OverrideUpdateCustomerDataAuthRequest) GoString() string {
	return s.String()
}

func (s *OverrideUpdateCustomerDataAuthRequest) SetCustomerIds(v []*string) *OverrideUpdateCustomerDataAuthRequest {
	s.CustomerIds = v
	return s
}

func (s *OverrideUpdateCustomerDataAuthRequest) SetDataAuthUserIds(v []*string) *OverrideUpdateCustomerDataAuthRequest {
	s.DataAuthUserIds = v
	return s
}

func (s *OverrideUpdateCustomerDataAuthRequest) SetFormCode(v string) *OverrideUpdateCustomerDataAuthRequest {
	s.FormCode = &v
	return s
}

func (s *OverrideUpdateCustomerDataAuthRequest) SetOperateUserId(v string) *OverrideUpdateCustomerDataAuthRequest {
	s.OperateUserId = &v
	return s
}

func (s *OverrideUpdateCustomerDataAuthRequest) SetRelationType(v string) *OverrideUpdateCustomerDataAuthRequest {
	s.RelationType = &v
	return s
}

func (s *OverrideUpdateCustomerDataAuthRequest) SetRoleType(v string) *OverrideUpdateCustomerDataAuthRequest {
	s.RoleType = &v
	return s
}

type OverrideUpdateCustomerDataAuthResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s OverrideUpdateCustomerDataAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OverrideUpdateCustomerDataAuthResponseBody) GoString() string {
	return s.String()
}

func (s *OverrideUpdateCustomerDataAuthResponseBody) SetResult(v bool) *OverrideUpdateCustomerDataAuthResponseBody {
	s.Result = &v
	return s
}

type OverrideUpdateCustomerDataAuthResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OverrideUpdateCustomerDataAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OverrideUpdateCustomerDataAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s OverrideUpdateCustomerDataAuthResponse) GoString() string {
	return s.String()
}

func (s *OverrideUpdateCustomerDataAuthResponse) SetHeaders(v map[string]*string) *OverrideUpdateCustomerDataAuthResponse {
	s.Headers = v
	return s
}

func (s *OverrideUpdateCustomerDataAuthResponse) SetStatusCode(v int32) *OverrideUpdateCustomerDataAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *OverrideUpdateCustomerDataAuthResponse) SetBody(v *OverrideUpdateCustomerDataAuthResponseBody) *OverrideUpdateCustomerDataAuthResponse {
	s.Body = v
	return s
}

type QueryAllCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerHeaders) SetCommonHeaders(v map[string]*string) *QueryAllCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllCustomerRequest struct {
	// example:
	//
	// 100
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 100010
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// crm_customer
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// ding_userid
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s QueryAllCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerRequest) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerRequest) SetMaxResults(v int64) *QueryAllCustomerRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllCustomerRequest) SetNextToken(v string) *QueryAllCustomerRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllCustomerRequest) SetObjectType(v string) *QueryAllCustomerRequest {
	s.ObjectType = &v
	return s
}

func (s *QueryAllCustomerRequest) SetOperatorUserId(v string) *QueryAllCustomerRequest {
	s.OperatorUserId = &v
	return s
}

type QueryAllCustomerResponseBody struct {
	Result *QueryAllCustomerResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryAllCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBody) SetResult(v *QueryAllCustomerResponseBodyResult) *QueryAllCustomerResponseBody {
	s.Result = v
	return s
}

type QueryAllCustomerResponseBodyResult struct {
	// example:
	//
	// 100
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 100
	NextToken *string                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Values    []*QueryAllCustomerResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryAllCustomerResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBodyResult) SetMaxResults(v int64) *QueryAllCustomerResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResult) SetNextToken(v string) *QueryAllCustomerResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResult) SetValues(v []*QueryAllCustomerResponseBodyResultValues) *QueryAllCustomerResponseBodyResult {
	s.Values = v
	return s
}

type QueryAllCustomerResponseBodyResultValues struct {
	// example:
	//
	// 2019-12-25 15:33:12
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 张三
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// example:
	//
	// ding_userid
	CreatorUserId *string                `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Data          map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData    map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// example:
	//
	// instance_id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 2019-12-25 15:33:12
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// example:
	//
	// crm_customer
	ObjectType *string                                             `json:"objectType,omitempty" xml:"objectType,omitempty"`
	Permission *QueryAllCustomerResponseBodyResultValuesPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// example:
	//
	// COMPLATE
	ProcessInstanceStatus *string `json:"processInstanceStatus,omitempty" xml:"processInstanceStatus,omitempty"`
	// example:
	//
	// agree
	ProcessOutResult *string `json:"processOutResult,omitempty" xml:"processOutResult,omitempty"`
}

func (s QueryAllCustomerResponseBodyResultValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBodyResultValues) SetCreateTime(v string) *QueryAllCustomerResponseBodyResultValues {
	s.CreateTime = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetCreatorNick(v string) *QueryAllCustomerResponseBodyResultValues {
	s.CreatorNick = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetCreatorUserId(v string) *QueryAllCustomerResponseBodyResultValues {
	s.CreatorUserId = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetData(v map[string]interface{}) *QueryAllCustomerResponseBodyResultValues {
	s.Data = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetExtendData(v map[string]interface{}) *QueryAllCustomerResponseBodyResultValues {
	s.ExtendData = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetInstanceId(v string) *QueryAllCustomerResponseBodyResultValues {
	s.InstanceId = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetModifyTime(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ModifyTime = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetObjectType(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ObjectType = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetPermission(v *QueryAllCustomerResponseBodyResultValuesPermission) *QueryAllCustomerResponseBodyResultValues {
	s.Permission = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetProcessInstanceStatus(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ProcessInstanceStatus = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetProcessOutResult(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ProcessOutResult = &v
	return s
}

type QueryAllCustomerResponseBodyResultValuesPermission struct {
	OwnerStaffIds       []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s QueryAllCustomerResponseBodyResultValuesPermission) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBodyResultValuesPermission) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBodyResultValuesPermission) SetOwnerStaffIds(v []*string) *QueryAllCustomerResponseBodyResultValuesPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValuesPermission) SetParticipantStaffIds(v []*string) *QueryAllCustomerResponseBodyResultValuesPermission {
	s.ParticipantStaffIds = v
	return s
}

type QueryAllCustomerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAllCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAllCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponse) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponse) SetHeaders(v map[string]*string) *QueryAllCustomerResponse {
	s.Headers = v
	return s
}

func (s *QueryAllCustomerResponse) SetStatusCode(v int32) *QueryAllCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllCustomerResponse) SetBody(v *QueryAllCustomerResponseBody) *QueryAllCustomerResponse {
	s.Body = v
	return s
}

type QueryAllTracksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllTracksHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTracksHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllTracksHeaders) SetCommonHeaders(v map[string]*string) *QueryAllTracksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllTracksHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllTracksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllTracksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 10000
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// asc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
}

func (s QueryAllTracksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTracksRequest) GoString() string {
	return s.String()
}

func (s *QueryAllTracksRequest) SetMaxResults(v int32) *QueryAllTracksRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllTracksRequest) SetNextToken(v string) *QueryAllTracksRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllTracksRequest) SetOrder(v string) *QueryAllTracksRequest {
	s.Order = &v
	return s
}

type QueryAllTracksResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 10001
	NextToken *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Values    []*QueryAllTracksResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryAllTracksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTracksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllTracksResponseBody) SetHasMore(v bool) *QueryAllTracksResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryAllTracksResponseBody) SetMaxResults(v int32) *QueryAllTracksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryAllTracksResponseBody) SetNextToken(v string) *QueryAllTracksResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryAllTracksResponseBody) SetValues(v []*QueryAllTracksResponseBodyValues) *QueryAllTracksResponseBody {
	s.Values = v
	return s
}

type QueryAllTracksResponseBodyValues struct {
	// example:
	//
	// 1234
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// example:
	//
	// manager1234
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// customer_id
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// example:
	//
	// 1237126786127
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asjkdh189127836
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 4
	SubType *int32 `json:"subType,omitempty" xml:"subType,omitempty"`
	// example:
	//
	// 80
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryAllTracksResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTracksResponseBodyValues) GoString() string {
	return s.String()
}

func (s *QueryAllTracksResponseBodyValues) SetBizId(v string) *QueryAllTracksResponseBodyValues {
	s.BizId = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetCreator(v string) *QueryAllTracksResponseBodyValues {
	s.Creator = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetCustomerId(v string) *QueryAllTracksResponseBodyValues {
	s.CustomerId = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetGmtCreate(v int64) *QueryAllTracksResponseBodyValues {
	s.GmtCreate = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetId(v string) *QueryAllTracksResponseBodyValues {
	s.Id = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetSubType(v int32) *QueryAllTracksResponseBodyValues {
	s.SubType = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetType(v int32) *QueryAllTracksResponseBodyValues {
	s.Type = &v
	return s
}

type QueryAllTracksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAllTracksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAllTracksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTracksResponse) GoString() string {
	return s.String()
}

func (s *QueryAllTracksResponse) SetHeaders(v map[string]*string) *QueryAllTracksResponse {
	s.Headers = v
	return s
}

func (s *QueryAllTracksResponse) SetStatusCode(v int32) *QueryAllTracksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllTracksResponse) SetBody(v *QueryAllTracksResponseBody) *QueryAllTracksResponse {
	s.Body = v
	return s
}

type QueryAppManagerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAppManagerHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAppManagerHeaders) GoString() string {
	return s.String()
}

func (s *QueryAppManagerHeaders) SetCommonHeaders(v map[string]*string) *QueryAppManagerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAppManagerHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAppManagerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAppManagerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 34234dfdfddd
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s QueryAppManagerRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppManagerRequest) GoString() string {
	return s.String()
}

func (s *QueryAppManagerRequest) SetOperatorUserId(v string) *QueryAppManagerRequest {
	s.OperatorUserId = &v
	return s
}

type QueryAppManagerResponseBody struct {
	Result []*QueryAppManagerResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryAppManagerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAppManagerResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppManagerResponseBody) SetResult(v []*QueryAppManagerResponseBodyResult) *QueryAppManagerResponseBody {
	s.Result = v
	return s
}

type QueryAppManagerResponseBodyResult struct {
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryAppManagerResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAppManagerResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAppManagerResponseBodyResult) SetAvatarUrl(v string) *QueryAppManagerResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *QueryAppManagerResponseBodyResult) SetName(v string) *QueryAppManagerResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryAppManagerResponseBodyResult) SetUserId(v string) *QueryAppManagerResponseBodyResult {
	s.UserId = &v
	return s
}

type QueryAppManagerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppManagerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppManagerResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAppManagerResponse) GoString() string {
	return s.String()
}

func (s *QueryAppManagerResponse) SetHeaders(v map[string]*string) *QueryAppManagerResponse {
	s.Headers = v
	return s
}

func (s *QueryAppManagerResponse) SetStatusCode(v int32) *QueryAppManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppManagerResponse) SetBody(v *QueryAppManagerResponseBody) *QueryAppManagerResponse {
	s.Body = v
	return s
}

type QueryBenefitInventoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBenefitInventoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBenefitInventoryHeaders) GoString() string {
	return s.String()
}

func (s *QueryBenefitInventoryHeaders) SetCommonHeaders(v map[string]*string) *QueryBenefitInventoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBenefitInventoryHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBenefitInventoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBenefitInventoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// B_CUSTOMER_CAPACITY
	BenefitCode *string `json:"benefitCode,omitempty" xml:"benefitCode,omitempty"`
}

func (s QueryBenefitInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBenefitInventoryRequest) GoString() string {
	return s.String()
}

func (s *QueryBenefitInventoryRequest) SetBenefitCode(v string) *QueryBenefitInventoryRequest {
	s.BenefitCode = &v
	return s
}

type QueryBenefitInventoryResponseBody struct {
	Result *QueryBenefitInventoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryBenefitInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBenefitInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBenefitInventoryResponseBody) SetResult(v *QueryBenefitInventoryResponseBodyResult) *QueryBenefitInventoryResponseBody {
	s.Result = v
	return s
}

type QueryBenefitInventoryResponseBodyResult struct {
	// example:
	//
	// 2000
	TotalQuota *int64 `json:"totalQuota,omitempty" xml:"totalQuota,omitempty"`
	// example:
	//
	// 10
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s QueryBenefitInventoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryBenefitInventoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryBenefitInventoryResponseBodyResult) SetTotalQuota(v int64) *QueryBenefitInventoryResponseBodyResult {
	s.TotalQuota = &v
	return s
}

func (s *QueryBenefitInventoryResponseBodyResult) SetUsedQuota(v int64) *QueryBenefitInventoryResponseBodyResult {
	s.UsedQuota = &v
	return s
}

type QueryBenefitInventoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBenefitInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBenefitInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBenefitInventoryResponse) GoString() string {
	return s.String()
}

func (s *QueryBenefitInventoryResponse) SetHeaders(v map[string]*string) *QueryBenefitInventoryResponse {
	s.Headers = v
	return s
}

func (s *QueryBenefitInventoryResponse) SetStatusCode(v int32) *QueryBenefitInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBenefitInventoryResponse) SetBody(v *QueryBenefitInventoryResponseBody) *QueryBenefitInventoryResponse {
	s.Body = v
	return s
}

type QueryClueFollowStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryClueFollowStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryClueFollowStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryClueFollowStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryClueFollowStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryClueFollowStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryClueFollowStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryClueFollowStatusRequest struct {
	ClueId *string `json:"clueId,omitempty" xml:"clueId,omitempty"`
}

func (s QueryClueFollowStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClueFollowStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryClueFollowStatusRequest) SetClueId(v string) *QueryClueFollowStatusRequest {
	s.ClueId = &v
	return s
}

type QueryClueFollowStatusResponseBody struct {
	Result []*QueryClueFollowStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryClueFollowStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClueFollowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClueFollowStatusResponseBody) SetResult(v []*QueryClueFollowStatusResponseBodyResult) *QueryClueFollowStatusResponseBody {
	s.Result = v
	return s
}

type QueryClueFollowStatusResponseBodyResult struct {
	ClueId *string `json:"clueId,omitempty" xml:"clueId,omitempty"`
	Scope  *string `json:"scope,omitempty" xml:"scope,omitempty"`
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryClueFollowStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryClueFollowStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryClueFollowStatusResponseBodyResult) SetClueId(v string) *QueryClueFollowStatusResponseBodyResult {
	s.ClueId = &v
	return s
}

func (s *QueryClueFollowStatusResponseBodyResult) SetScope(v string) *QueryClueFollowStatusResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *QueryClueFollowStatusResponseBodyResult) SetStatus(v string) *QueryClueFollowStatusResponseBodyResult {
	s.Status = &v
	return s
}

type QueryClueFollowStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryClueFollowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryClueFollowStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClueFollowStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryClueFollowStatusResponse) SetHeaders(v map[string]*string) *QueryClueFollowStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryClueFollowStatusResponse) SetStatusCode(v int32) *QueryClueFollowStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryClueFollowStatusResponse) SetBody(v *QueryClueFollowStatusResponseBody) *QueryClueFollowStatusResponse {
	s.Body = v
	return s
}

type QueryCrmGroupChatsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCrmGroupChatsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmGroupChatsHeaders) GoString() string {
	return s.String()
}

func (s *QueryCrmGroupChatsHeaders) SetCommonHeaders(v map[string]*string) *QueryCrmGroupChatsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCrmGroupChatsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCrmGroupChatsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCrmGroupChatsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	QueryDsl   *string `json:"queryDsl,omitempty" xml:"queryDsl,omitempty"`
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s QueryCrmGroupChatsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmGroupChatsRequest) GoString() string {
	return s.String()
}

func (s *QueryCrmGroupChatsRequest) SetMaxResults(v int32) *QueryCrmGroupChatsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCrmGroupChatsRequest) SetNextToken(v string) *QueryCrmGroupChatsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCrmGroupChatsRequest) SetQueryDsl(v string) *QueryCrmGroupChatsRequest {
	s.QueryDsl = &v
	return s
}

func (s *QueryCrmGroupChatsRequest) SetRelationType(v string) *QueryCrmGroupChatsRequest {
	s.RelationType = &v
	return s
}

type QueryCrmGroupChatsResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// agds12
	NextToken  *string                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResultList []*QueryCrmGroupChatsResponseBodyResultList `json:"resultList,omitempty" xml:"resultList,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryCrmGroupChatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmGroupChatsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCrmGroupChatsResponseBody) SetHasMore(v bool) *QueryCrmGroupChatsResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCrmGroupChatsResponseBody) SetNextToken(v string) *QueryCrmGroupChatsResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCrmGroupChatsResponseBody) SetResultList(v []*QueryCrmGroupChatsResponseBodyResultList) *QueryCrmGroupChatsResponseBody {
	s.ResultList = v
	return s
}

func (s *QueryCrmGroupChatsResponseBody) SetTotalCount(v int32) *QueryCrmGroupChatsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryCrmGroupChatsResponseBodyResultList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1640239655539
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	MemberCount *int32 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 营销1群
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// afsad21
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// afsdba23
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// afds12
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XX
	OwnerUserName *string `json:"ownerUserName,omitempty" xml:"ownerUserName,omitempty"`
}

func (s QueryCrmGroupChatsResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmGroupChatsResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *QueryCrmGroupChatsResponseBodyResultList) SetGmtCreate(v int64) *QueryCrmGroupChatsResponseBodyResultList {
	s.GmtCreate = &v
	return s
}

func (s *QueryCrmGroupChatsResponseBodyResultList) SetMemberCount(v int32) *QueryCrmGroupChatsResponseBodyResultList {
	s.MemberCount = &v
	return s
}

func (s *QueryCrmGroupChatsResponseBodyResultList) SetName(v string) *QueryCrmGroupChatsResponseBodyResultList {
	s.Name = &v
	return s
}

func (s *QueryCrmGroupChatsResponseBodyResultList) SetOpenConversationId(v string) *QueryCrmGroupChatsResponseBodyResultList {
	s.OpenConversationId = &v
	return s
}

func (s *QueryCrmGroupChatsResponseBodyResultList) SetOpenGroupSetId(v string) *QueryCrmGroupChatsResponseBodyResultList {
	s.OpenGroupSetId = &v
	return s
}

func (s *QueryCrmGroupChatsResponseBodyResultList) SetOwnerUserId(v string) *QueryCrmGroupChatsResponseBodyResultList {
	s.OwnerUserId = &v
	return s
}

func (s *QueryCrmGroupChatsResponseBodyResultList) SetOwnerUserName(v string) *QueryCrmGroupChatsResponseBodyResultList {
	s.OwnerUserName = &v
	return s
}

type QueryCrmGroupChatsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCrmGroupChatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCrmGroupChatsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmGroupChatsResponse) GoString() string {
	return s.String()
}

func (s *QueryCrmGroupChatsResponse) SetHeaders(v map[string]*string) *QueryCrmGroupChatsResponse {
	s.Headers = v
	return s
}

func (s *QueryCrmGroupChatsResponse) SetStatusCode(v int32) *QueryCrmGroupChatsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCrmGroupChatsResponse) SetBody(v *QueryCrmGroupChatsResponseBody) *QueryCrmGroupChatsResponse {
	s.Body = v
	return s
}

type QueryCrmPersonalCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCrmPersonalCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerHeaders) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerHeaders) SetCommonHeaders(v map[string]*string) *QueryCrmPersonalCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCrmPersonalCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCrmPersonalCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCrmPersonalCustomerRequest struct {
	CurrentOperatorUserId *string `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
	// This parameter is required.
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	QueryDsl   *string `json:"queryDsl,omitempty" xml:"queryDsl,omitempty"`
	// if can be null:
	// false
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s QueryCrmPersonalCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerRequest) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerRequest) SetCurrentOperatorUserId(v string) *QueryCrmPersonalCustomerRequest {
	s.CurrentOperatorUserId = &v
	return s
}

func (s *QueryCrmPersonalCustomerRequest) SetMaxResults(v int32) *QueryCrmPersonalCustomerRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCrmPersonalCustomerRequest) SetNextToken(v string) *QueryCrmPersonalCustomerRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCrmPersonalCustomerRequest) SetQueryDsl(v string) *QueryCrmPersonalCustomerRequest {
	s.QueryDsl = &v
	return s
}

func (s *QueryCrmPersonalCustomerRequest) SetRelationType(v string) *QueryCrmPersonalCustomerRequest {
	s.RelationType = &v
	return s
}

type QueryCrmPersonalCustomerResponseBody struct {
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int32                                        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Values     []*QueryCrmPersonalCustomerResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryCrmPersonalCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerResponseBody) SetHasMore(v bool) *QueryCrmPersonalCustomerResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBody) SetMaxResults(v int32) *QueryCrmPersonalCustomerResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBody) SetNextToken(v string) *QueryCrmPersonalCustomerResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBody) SetTotalCount(v int32) *QueryCrmPersonalCustomerResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBody) SetValues(v []*QueryCrmPersonalCustomerResponseBodyValues) *QueryCrmPersonalCustomerResponseBody {
	s.Values = v
	return s
}

type QueryCrmPersonalCustomerResponseBodyValues struct {
	// This parameter is required.
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// This parameter is required.
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// This parameter is required.
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// This parameter is required.
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// This parameter is required.
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// This parameter is required.
	Permission *QueryCrmPersonalCustomerResponseBodyValuesPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// This parameter is required.
	ProcInstStatus *string `json:"procInstStatus,omitempty" xml:"procInstStatus,omitempty"`
	// This parameter is required.
	ProcOutResult *string `json:"procOutResult,omitempty" xml:"procOutResult,omitempty"`
}

func (s QueryCrmPersonalCustomerResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerResponseBodyValues) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetCreatorNick(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.CreatorNick = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetCreatorUserId(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.CreatorUserId = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetData(v map[string]interface{}) *QueryCrmPersonalCustomerResponseBodyValues {
	s.Data = v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetExtendData(v map[string]interface{}) *QueryCrmPersonalCustomerResponseBodyValues {
	s.ExtendData = v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetGmtCreate(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.GmtCreate = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetGmtModified(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.GmtModified = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetInstanceId(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.InstanceId = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetObjectType(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.ObjectType = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetPermission(v *QueryCrmPersonalCustomerResponseBodyValuesPermission) *QueryCrmPersonalCustomerResponseBodyValues {
	s.Permission = v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetProcInstStatus(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.ProcInstStatus = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetProcOutResult(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.ProcOutResult = &v
	return s
}

type QueryCrmPersonalCustomerResponseBodyValuesPermission struct {
	// This parameter is required.
	OwnerStaffIds []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	// This parameter is required.
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s QueryCrmPersonalCustomerResponseBodyValuesPermission) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerResponseBodyValuesPermission) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerResponseBodyValuesPermission) SetOwnerStaffIds(v []*string) *QueryCrmPersonalCustomerResponseBodyValuesPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValuesPermission) SetParticipantStaffIds(v []*string) *QueryCrmPersonalCustomerResponseBodyValuesPermission {
	s.ParticipantStaffIds = v
	return s
}

type QueryCrmPersonalCustomerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCrmPersonalCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerResponse) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerResponse) SetHeaders(v map[string]*string) *QueryCrmPersonalCustomerResponse {
	s.Headers = v
	return s
}

func (s *QueryCrmPersonalCustomerResponse) SetStatusCode(v int32) *QueryCrmPersonalCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponse) SetBody(v *QueryCrmPersonalCustomerResponseBody) *QueryCrmPersonalCustomerResponse {
	s.Body = v
	return s
}

type QueryCustomerBizTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCustomerBizTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerBizTypeHeaders) GoString() string {
	return s.String()
}

func (s *QueryCustomerBizTypeHeaders) SetCommonHeaders(v map[string]*string) *QueryCustomerBizTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCustomerBizTypeHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCustomerBizTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCustomerBizTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1343dfd1233
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s QueryCustomerBizTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerBizTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomerBizTypeRequest) SetOperatorUserId(v string) *QueryCustomerBizTypeRequest {
	s.OperatorUserId = &v
	return s
}

type QueryCustomerBizTypeResponseBody struct {
	Result *QueryCustomerBizTypeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryCustomerBizTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomerBizTypeResponseBody) SetResult(v *QueryCustomerBizTypeResponseBodyResult) *QueryCustomerBizTypeResponseBody {
	s.Result = v
	return s
}

type QueryCustomerBizTypeResponseBodyResult struct {
	// example:
	//
	// crm_customer
	CustomerBizType *string `json:"customerBizType,omitempty" xml:"customerBizType,omitempty"`
}

func (s QueryCustomerBizTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerBizTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCustomerBizTypeResponseBodyResult) SetCustomerBizType(v string) *QueryCustomerBizTypeResponseBodyResult {
	s.CustomerBizType = &v
	return s
}

type QueryCustomerBizTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomerBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomerBizTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerBizTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomerBizTypeResponse) SetHeaders(v map[string]*string) *QueryCustomerBizTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomerBizTypeResponse) SetStatusCode(v int32) *QueryCustomerBizTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomerBizTypeResponse) SetBody(v *QueryCustomerBizTypeResponseBody) *QueryCustomerBizTypeResponse {
	s.Body = v
	return s
}

type QueryGlobalInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGlobalInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGlobalInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryGlobalInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryGlobalInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGlobalInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGlobalInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGlobalInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 301227837938
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryGlobalInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGlobalInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryGlobalInfoRequest) SetUserId(v string) *QueryGlobalInfoRequest {
	s.UserId = &v
	return s
}

type QueryGlobalInfoResponseBody struct {
	Result *QueryGlobalInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryGlobalInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGlobalInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGlobalInfoResponseBody) SetResult(v *QueryGlobalInfoResponseBodyResult) *QueryGlobalInfoResponseBody {
	s.Result = v
	return s
}

type QueryGlobalInfoResponseBodyResult struct {
	OemEnable   *bool `json:"oemEnable,omitempty" xml:"oemEnable,omitempty"`
	T2t3Coexist *bool `json:"t2t3Coexist,omitempty" xml:"t2t3Coexist,omitempty"`
}

func (s QueryGlobalInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryGlobalInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryGlobalInfoResponseBodyResult) SetOemEnable(v bool) *QueryGlobalInfoResponseBodyResult {
	s.OemEnable = &v
	return s
}

func (s *QueryGlobalInfoResponseBodyResult) SetT2t3Coexist(v bool) *QueryGlobalInfoResponseBodyResult {
	s.T2t3Coexist = &v
	return s
}

type QueryGlobalInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGlobalInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGlobalInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGlobalInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryGlobalInfoResponse) SetHeaders(v map[string]*string) *QueryGlobalInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryGlobalInfoResponse) SetStatusCode(v int32) *QueryGlobalInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGlobalInfoResponse) SetBody(v *QueryGlobalInfoResponseBody) *QueryGlobalInfoResponse {
	s.Body = v
	return s
}

type QueryHasAppPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHasAppPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHasAppPermissionHeaders) GoString() string {
	return s.String()
}

func (s *QueryHasAppPermissionHeaders) SetCommonHeaders(v map[string]*string) *QueryHasAppPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHasAppPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHasAppPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHasAppPermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 34234dfdfddd
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s QueryHasAppPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHasAppPermissionRequest) GoString() string {
	return s.String()
}

func (s *QueryHasAppPermissionRequest) SetOperatorUserId(v string) *QueryHasAppPermissionRequest {
	s.OperatorUserId = &v
	return s
}

type QueryHasAppPermissionResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s QueryHasAppPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHasAppPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHasAppPermissionResponseBody) SetResult(v bool) *QueryHasAppPermissionResponseBody {
	s.Result = &v
	return s
}

type QueryHasAppPermissionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHasAppPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHasAppPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHasAppPermissionResponse) GoString() string {
	return s.String()
}

func (s *QueryHasAppPermissionResponse) SetHeaders(v map[string]*string) *QueryHasAppPermissionResponse {
	s.Headers = v
	return s
}

func (s *QueryHasAppPermissionResponse) SetStatusCode(v int32) *QueryHasAppPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHasAppPermissionResponse) SetBody(v *QueryHasAppPermissionResponseBody) *QueryHasAppPermissionResponse {
	s.Body = v
	return s
}

type QueryOfficialAccountUserBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOfficialAccountUserBasicInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialAccountUserBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryOfficialAccountUserBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryOfficialAccountUserBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOfficialAccountUserBasicInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOfficialAccountUserBasicInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOfficialAccountUserBasicInfoRequest struct {
	// This parameter is required.
	BindingToken *string `json:"bindingToken,omitempty" xml:"bindingToken,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryOfficialAccountUserBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialAccountUserBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryOfficialAccountUserBasicInfoRequest) SetBindingToken(v string) *QueryOfficialAccountUserBasicInfoRequest {
	s.BindingToken = &v
	return s
}

func (s *QueryOfficialAccountUserBasicInfoRequest) SetUnionId(v string) *QueryOfficialAccountUserBasicInfoRequest {
	s.UnionId = &v
	return s
}

type QueryOfficialAccountUserBasicInfoResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Result *QueryOfficialAccountUserBasicInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryOfficialAccountUserBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialAccountUserBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOfficialAccountUserBasicInfoResponseBody) SetRequestId(v string) *QueryOfficialAccountUserBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOfficialAccountUserBasicInfoResponseBody) SetResult(v *QueryOfficialAccountUserBasicInfoResponseBodyResult) *QueryOfficialAccountUserBasicInfoResponseBody {
	s.Result = v
	return s
}

type QueryOfficialAccountUserBasicInfoResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// FOLLOWED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryOfficialAccountUserBasicInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialAccountUserBasicInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOfficialAccountUserBasicInfoResponseBodyResult) SetStatus(v string) *QueryOfficialAccountUserBasicInfoResponseBodyResult {
	s.Status = &v
	return s
}

type QueryOfficialAccountUserBasicInfoResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOfficialAccountUserBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOfficialAccountUserBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialAccountUserBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryOfficialAccountUserBasicInfoResponse) SetHeaders(v map[string]*string) *QueryOfficialAccountUserBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryOfficialAccountUserBasicInfoResponse) SetStatusCode(v int32) *QueryOfficialAccountUserBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOfficialAccountUserBasicInfoResponse) SetBody(v *QueryOfficialAccountUserBasicInfoResponseBody) *QueryOfficialAccountUserBasicInfoResponse {
	s.Body = v
	return s
}

type QueryRelationDatasByTargetIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRelationDatasByTargetIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationDatasByTargetIdHeaders) GoString() string {
	return s.String()
}

func (s *QueryRelationDatasByTargetIdHeaders) SetCommonHeaders(v map[string]*string) *QueryRelationDatasByTargetIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRelationDatasByTargetIdHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRelationDatasByTargetIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRelationDatasByTargetIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc123
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s QueryRelationDatasByTargetIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationDatasByTargetIdRequest) GoString() string {
	return s.String()
}

func (s *QueryRelationDatasByTargetIdRequest) SetRelationType(v string) *QueryRelationDatasByTargetIdRequest {
	s.RelationType = &v
	return s
}

type QueryRelationDatasByTargetIdResponseBody struct {
	// This parameter is required.
	Relations []*QueryRelationDatasByTargetIdResponseBodyRelations `json:"relations,omitempty" xml:"relations,omitempty" type:"Repeated"`
}

func (s QueryRelationDatasByTargetIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationDatasByTargetIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRelationDatasByTargetIdResponseBody) SetRelations(v []*QueryRelationDatasByTargetIdResponseBodyRelations) *QueryRelationDatasByTargetIdResponseBody {
	s.Relations = v
	return s
}

type QueryRelationDatasByTargetIdResponseBodyRelations struct {
	// This parameter is required.
	BizDataList []*QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList `json:"bizDataList,omitempty" xml:"bizDataList,omitempty" type:"Repeated"`
	// This parameter is required.
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// abc123
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc123
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s QueryRelationDatasByTargetIdResponseBodyRelations) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationDatasByTargetIdResponseBodyRelations) GoString() string {
	return s.String()
}

func (s *QueryRelationDatasByTargetIdResponseBodyRelations) SetBizDataList(v []*QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList) *QueryRelationDatasByTargetIdResponseBodyRelations {
	s.BizDataList = v
	return s
}

func (s *QueryRelationDatasByTargetIdResponseBodyRelations) SetOpenConversationIds(v []*string) *QueryRelationDatasByTargetIdResponseBodyRelations {
	s.OpenConversationIds = v
	return s
}

func (s *QueryRelationDatasByTargetIdResponseBodyRelations) SetRelationId(v string) *QueryRelationDatasByTargetIdResponseBodyRelations {
	s.RelationId = &v
	return s
}

func (s *QueryRelationDatasByTargetIdResponseBodyRelations) SetRelationType(v string) *QueryRelationDatasByTargetIdResponseBodyRelations {
	s.RelationType = &v
	return s
}

type QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList struct {
	// This parameter is required.
	//
	// example:
	//
	// {}
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer_name
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc123
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList) GoString() string {
	return s.String()
}

func (s *QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList) SetExtendValue(v string) *QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList {
	s.ExtendValue = &v
	return s
}

func (s *QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList) SetKey(v string) *QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList {
	s.Key = &v
	return s
}

func (s *QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList) SetValue(v string) *QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList {
	s.Value = &v
	return s
}

type QueryRelationDatasByTargetIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRelationDatasByTargetIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRelationDatasByTargetIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationDatasByTargetIdResponse) GoString() string {
	return s.String()
}

func (s *QueryRelationDatasByTargetIdResponse) SetHeaders(v map[string]*string) *QueryRelationDatasByTargetIdResponse {
	s.Headers = v
	return s
}

func (s *QueryRelationDatasByTargetIdResponse) SetStatusCode(v int32) *QueryRelationDatasByTargetIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRelationDatasByTargetIdResponse) SetBody(v *QueryRelationDatasByTargetIdResponseBody) *QueryRelationDatasByTargetIdResponse {
	s.Body = v
	return s
}

type RecallOfficialAccountOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RecallOfficialAccountOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s RecallOfficialAccountOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *RecallOfficialAccountOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *RecallOfficialAccountOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecallOfficialAccountOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *RecallOfficialAccountOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RecallOfficialAccountOTOMessageRequest struct {
	// example:
	//
	// ding123
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SWXXX
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s RecallOfficialAccountOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallOfficialAccountOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *RecallOfficialAccountOTOMessageRequest) SetAccountId(v string) *RecallOfficialAccountOTOMessageRequest {
	s.AccountId = &v
	return s
}

func (s *RecallOfficialAccountOTOMessageRequest) SetOpenPushId(v string) *RecallOfficialAccountOTOMessageRequest {
	s.OpenPushId = &v
	return s
}

type RecallOfficialAccountOTOMessageResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RecallOfficialAccountOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecallOfficialAccountOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *RecallOfficialAccountOTOMessageResponseBody) SetRequestId(v string) *RecallOfficialAccountOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

type RecallOfficialAccountOTOMessageResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecallOfficialAccountOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecallOfficialAccountOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallOfficialAccountOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *RecallOfficialAccountOTOMessageResponse) SetHeaders(v map[string]*string) *RecallOfficialAccountOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *RecallOfficialAccountOTOMessageResponse) SetStatusCode(v int32) *RecallOfficialAccountOTOMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecallOfficialAccountOTOMessageResponse) SetBody(v *RecallOfficialAccountOTOMessageResponseBody) *RecallOfficialAccountOTOMessageResponse {
	s.Body = v
	return s
}

type SaveBenefitLicenseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveBenefitLicenseHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveBenefitLicenseHeaders) GoString() string {
	return s.String()
}

func (s *SaveBenefitLicenseHeaders) SetCommonHeaders(v map[string]*string) *SaveBenefitLicenseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveBenefitLicenseHeaders) SetXAcsDingtalkAccessToken(v string) *SaveBenefitLicenseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveBenefitLicenseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// B_ACCOUNT_NUMBER
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// This parameter is required.
	Licenses []*SaveBenefitLicenseRequestLicenses `json:"licenses,omitempty" xml:"licenses,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// staffId
	SaveUserId *string `json:"saveUserId,omitempty" xml:"saveUserId,omitempty"`
}

func (s SaveBenefitLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBenefitLicenseRequest) GoString() string {
	return s.String()
}

func (s *SaveBenefitLicenseRequest) SetDomain(v string) *SaveBenefitLicenseRequest {
	s.Domain = &v
	return s
}

func (s *SaveBenefitLicenseRequest) SetLicenses(v []*SaveBenefitLicenseRequestLicenses) *SaveBenefitLicenseRequest {
	s.Licenses = v
	return s
}

func (s *SaveBenefitLicenseRequest) SetSaveUserId(v string) *SaveBenefitLicenseRequest {
	s.SaveUserId = &v
	return s
}

type SaveBenefitLicenseRequestLicenses struct {
	// example:
	//
	// licenseStaffId
	LicenseUserId *string `json:"licenseUserId,omitempty" xml:"licenseUserId,omitempty"`
}

func (s SaveBenefitLicenseRequestLicenses) String() string {
	return tea.Prettify(s)
}

func (s SaveBenefitLicenseRequestLicenses) GoString() string {
	return s.String()
}

func (s *SaveBenefitLicenseRequestLicenses) SetLicenseUserId(v string) *SaveBenefitLicenseRequestLicenses {
	s.LicenseUserId = &v
	return s
}

type SaveBenefitLicenseResponseBody struct {
	Result *SaveBenefitLicenseResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SaveBenefitLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBenefitLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBenefitLicenseResponseBody) SetResult(v *SaveBenefitLicenseResponseBodyResult) *SaveBenefitLicenseResponseBody {
	s.Result = v
	return s
}

type SaveBenefitLicenseResponseBodyResult struct {
	// example:
	//
	// B_ACCOUNT_NUMBER
	Domain   *string                                         `json:"domain,omitempty" xml:"domain,omitempty"`
	Licenses []*SaveBenefitLicenseResponseBodyResultLicenses `json:"licenses,omitempty" xml:"licenses,omitempty" type:"Repeated"`
}

func (s SaveBenefitLicenseResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SaveBenefitLicenseResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SaveBenefitLicenseResponseBodyResult) SetDomain(v string) *SaveBenefitLicenseResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *SaveBenefitLicenseResponseBodyResult) SetLicenses(v []*SaveBenefitLicenseResponseBodyResultLicenses) *SaveBenefitLicenseResponseBodyResult {
	s.Licenses = v
	return s
}

type SaveBenefitLicenseResponseBodyResultLicenses struct {
	// example:
	//
	// license账号信息
	LicenseUserId *string `json:"licenseUserId,omitempty" xml:"licenseUserId,omitempty"`
}

func (s SaveBenefitLicenseResponseBodyResultLicenses) String() string {
	return tea.Prettify(s)
}

func (s SaveBenefitLicenseResponseBodyResultLicenses) GoString() string {
	return s.String()
}

func (s *SaveBenefitLicenseResponseBodyResultLicenses) SetLicenseUserId(v string) *SaveBenefitLicenseResponseBodyResultLicenses {
	s.LicenseUserId = &v
	return s
}

type SaveBenefitLicenseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBenefitLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBenefitLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBenefitLicenseResponse) GoString() string {
	return s.String()
}

func (s *SaveBenefitLicenseResponse) SetHeaders(v map[string]*string) *SaveBenefitLicenseResponse {
	s.Headers = v
	return s
}

func (s *SaveBenefitLicenseResponse) SetStatusCode(v int32) *SaveBenefitLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBenefitLicenseResponse) SetBody(v *SaveBenefitLicenseResponseBody) *SaveBenefitLicenseResponse {
	s.Body = v
	return s
}

type SendOfficialAccountOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendOfficialAccountOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *SendOfficialAccountOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendOfficialAccountOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendOfficialAccountOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendOfficialAccountOTOMessageRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	BizId     *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	Detail *SendOfficialAccountOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
}

func (s SendOfficialAccountOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequest) SetAccountId(v string) *SendOfficialAccountOTOMessageRequest {
	s.AccountId = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequest) SetBizId(v string) *SendOfficialAccountOTOMessageRequest {
	s.BizId = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequest) SetDetail(v *SendOfficialAccountOTOMessageRequestDetail) *SendOfficialAccountOTOMessageRequest {
	s.Detail = v
	return s
}

type SendOfficialAccountOTOMessageRequestDetail struct {
	// This parameter is required.
	MessageBody *SendOfficialAccountOTOMessageRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// text
	MsgType *string `json:"msgType,omitempty" xml:"msgType,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetail) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetail) SetMessageBody(v *SendOfficialAccountOTOMessageRequestDetailMessageBody) *SendOfficialAccountOTOMessageRequestDetail {
	s.MessageBody = v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetail) SetMsgType(v string) *SendOfficialAccountOTOMessageRequestDetail {
	s.MsgType = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetail) SetUnionId(v string) *SendOfficialAccountOTOMessageRequestDetail {
	s.UnionId = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetail) SetUserId(v string) *SendOfficialAccountOTOMessageRequestDetail {
	s.UserId = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetail) SetUuid(v string) *SendOfficialAccountOTOMessageRequestDetail {
	s.Uuid = &v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBody struct {
	ActionCard *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard `json:"actionCard,omitempty" xml:"actionCard,omitempty" type:"Struct"`
	Image      *SendOfficialAccountOTOMessageRequestDetailMessageBodyImage      `json:"image,omitempty" xml:"image,omitempty" type:"Struct"`
	Link       *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink       `json:"link,omitempty" xml:"link,omitempty" type:"Struct"`
	Markdown   *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown   `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	Text       *SendOfficialAccountOTOMessageRequestDetailMessageBodyText       `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBody) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBody) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBody) SetActionCard(v *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) *SendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.ActionCard = v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBody) SetImage(v *SendOfficialAccountOTOMessageRequestDetailMessageBodyImage) *SendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Image = v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBody) SetLink(v *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) *SendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Link = v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBody) SetMarkdown(v *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) *SendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Markdown = v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBody) SetText(v *SendOfficialAccountOTOMessageRequestDetailMessageBodyText) *SendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Text = v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard struct {
	ButtonList        []*SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList `json:"buttonList,omitempty" xml:"buttonList,omitempty" type:"Repeated"`
	ButtonOrientation *string                                                                      `json:"buttonOrientation,omitempty" xml:"buttonOrientation,omitempty"`
	Markdown          *string                                                                      `json:"markdown,omitempty" xml:"markdown,omitempty"`
	SingleTitle       *string                                                                      `json:"singleTitle,omitempty" xml:"singleTitle,omitempty"`
	SingleUrl         *string                                                                      `json:"singleUrl,omitempty" xml:"singleUrl,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetButtonList(v []*SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.ButtonList = v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetButtonOrientation(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.ButtonOrientation = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetMarkdown(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.Markdown = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetSingleTitle(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.SingleTitle = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetSingleUrl(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.SingleUrl = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetTitle(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.Title = &v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList struct {
	// This parameter is required.
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) SetActionUrl(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList {
	s.ActionUrl = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) SetTitle(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList {
	s.Title = &v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyImage struct {
	// example:
	//
	// @lALPBbCc1XuaP_rNAljNAl
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyImage) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyImage) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyImage) SetMediaId(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyImage {
	s.MediaId = &v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyLink struct {
	// This parameter is required.
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	// This parameter is required.
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetMessageUrl(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.MessageUrl = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetPicUrl(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.PicUrl = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetText(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.Text = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetTitle(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.Title = &v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown struct {
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) SetText(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown {
	s.Text = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) SetTitle(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown {
	s.Title = &v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyText struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyText) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyText) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyText) SetContent(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyText {
	s.Content = &v
	return s
}

type SendOfficialAccountOTOMessageResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Result *SendOfficialAccountOTOMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SendOfficialAccountOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageResponseBody) SetRequestId(v string) *SendOfficialAccountOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendOfficialAccountOTOMessageResponseBody) SetResult(v *SendOfficialAccountOTOMessageResponseBodyResult) *SendOfficialAccountOTOMessageResponseBody {
	s.Result = v
	return s
}

type SendOfficialAccountOTOMessageResponseBodyResult struct {
	// This parameter is required.
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s SendOfficialAccountOTOMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageResponseBodyResult) SetOpenPushId(v string) *SendOfficialAccountOTOMessageResponseBodyResult {
	s.OpenPushId = &v
	return s
}

type SendOfficialAccountOTOMessageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendOfficialAccountOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendOfficialAccountOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageResponse) SetHeaders(v map[string]*string) *SendOfficialAccountOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *SendOfficialAccountOTOMessageResponse) SetStatusCode(v int32) *SendOfficialAccountOTOMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendOfficialAccountOTOMessageResponse) SetBody(v *SendOfficialAccountOTOMessageResponseBody) *SendOfficialAccountOTOMessageResponse {
	s.Body = v
	return s
}

type SendOfficialAccountSNSMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendOfficialAccountSNSMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageHeaders) SetCommonHeaders(v map[string]*string) *SendOfficialAccountSNSMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendOfficialAccountSNSMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendOfficialAccountSNSMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendOfficialAccountSNSMessageRequest struct {
	// This parameter is required.
	BindingToken *string `json:"bindingToken,omitempty" xml:"bindingToken,omitempty"`
	BizId        *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	Detail *SendOfficialAccountSNSMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
}

func (s SendOfficialAccountSNSMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageRequest) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageRequest) SetBindingToken(v string) *SendOfficialAccountSNSMessageRequest {
	s.BindingToken = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequest) SetBizId(v string) *SendOfficialAccountSNSMessageRequest {
	s.BizId = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequest) SetDetail(v *SendOfficialAccountSNSMessageRequestDetail) *SendOfficialAccountSNSMessageRequest {
	s.Detail = v
	return s
}

type SendOfficialAccountSNSMessageRequestDetail struct {
	// This parameter is required.
	MessageBody *SendOfficialAccountSNSMessageRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// text
	MsgType *string `json:"msgType,omitempty" xml:"msgType,omitempty"`
	// This parameter is required.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s SendOfficialAccountSNSMessageRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageRequestDetail) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageRequestDetail) SetMessageBody(v *SendOfficialAccountSNSMessageRequestDetailMessageBody) *SendOfficialAccountSNSMessageRequestDetail {
	s.MessageBody = v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetail) SetMsgType(v string) *SendOfficialAccountSNSMessageRequestDetail {
	s.MsgType = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetail) SetUuid(v string) *SendOfficialAccountSNSMessageRequestDetail {
	s.Uuid = &v
	return s
}

type SendOfficialAccountSNSMessageRequestDetailMessageBody struct {
	ActionCard *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard `json:"actionCard,omitempty" xml:"actionCard,omitempty" type:"Struct"`
	Link       *SendOfficialAccountSNSMessageRequestDetailMessageBodyLink       `json:"link,omitempty" xml:"link,omitempty" type:"Struct"`
	Markdown   *SendOfficialAccountSNSMessageRequestDetailMessageBodyMarkdown   `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	Text       *SendOfficialAccountSNSMessageRequestDetailMessageBodyText       `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBody) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBody) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBody) SetActionCard(v *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard) *SendOfficialAccountSNSMessageRequestDetailMessageBody {
	s.ActionCard = v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBody) SetLink(v *SendOfficialAccountSNSMessageRequestDetailMessageBodyLink) *SendOfficialAccountSNSMessageRequestDetailMessageBody {
	s.Link = v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBody) SetMarkdown(v *SendOfficialAccountSNSMessageRequestDetailMessageBodyMarkdown) *SendOfficialAccountSNSMessageRequestDetailMessageBody {
	s.Markdown = v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBody) SetText(v *SendOfficialAccountSNSMessageRequestDetailMessageBodyText) *SendOfficialAccountSNSMessageRequestDetailMessageBody {
	s.Text = v
	return s
}

type SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard struct {
	ButtonList        []*SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCardButtonList `json:"buttonList,omitempty" xml:"buttonList,omitempty" type:"Repeated"`
	ButtonOrientation *string                                                                      `json:"buttonOrientation,omitempty" xml:"buttonOrientation,omitempty"`
	Markdown          *string                                                                      `json:"markdown,omitempty" xml:"markdown,omitempty"`
	SingleTitle       *string                                                                      `json:"singleTitle,omitempty" xml:"singleTitle,omitempty"`
	SingleUrl         *string                                                                      `json:"singleUrl,omitempty" xml:"singleUrl,omitempty"`
	Title             *string                                                                      `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard) SetButtonList(v []*SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCardButtonList) *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard {
	s.ButtonList = v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard) SetButtonOrientation(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard {
	s.ButtonOrientation = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard) SetMarkdown(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard {
	s.Markdown = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard) SetSingleTitle(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard {
	s.SingleTitle = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard) SetSingleUrl(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard {
	s.SingleUrl = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard) SetTitle(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCard {
	s.Title = &v
	return s
}

type SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCardButtonList struct {
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCardButtonList) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCardButtonList) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCardButtonList) SetActionUrl(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCardButtonList {
	s.ActionUrl = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCardButtonList) SetTitle(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyActionCardButtonList {
	s.Title = &v
	return s
}

type SendOfficialAccountSNSMessageRequestDetailMessageBodyLink struct {
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	PicUrl     *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	Text       *string `json:"text,omitempty" xml:"text,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBodyLink) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBodyLink) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyLink) SetMessageUrl(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyLink {
	s.MessageUrl = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyLink) SetPicUrl(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyLink {
	s.PicUrl = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyLink) SetText(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyLink {
	s.Text = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyLink) SetTitle(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyLink {
	s.Title = &v
	return s
}

type SendOfficialAccountSNSMessageRequestDetailMessageBodyMarkdown struct {
	Text  *string `json:"text,omitempty" xml:"text,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBodyMarkdown) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBodyMarkdown) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyMarkdown) SetText(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyMarkdown {
	s.Text = &v
	return s
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyMarkdown) SetTitle(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyMarkdown {
	s.Title = &v
	return s
}

type SendOfficialAccountSNSMessageRequestDetailMessageBodyText struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBodyText) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageRequestDetailMessageBodyText) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageRequestDetailMessageBodyText) SetContent(v string) *SendOfficialAccountSNSMessageRequestDetailMessageBodyText {
	s.Content = &v
	return s
}

type SendOfficialAccountSNSMessageResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Result *SendOfficialAccountSNSMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SendOfficialAccountSNSMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageResponseBody) SetRequestId(v string) *SendOfficialAccountSNSMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendOfficialAccountSNSMessageResponseBody) SetResult(v *SendOfficialAccountSNSMessageResponseBodyResult) *SendOfficialAccountSNSMessageResponseBody {
	s.Result = v
	return s
}

type SendOfficialAccountSNSMessageResponseBodyResult struct {
	// This parameter is required.
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s SendOfficialAccountSNSMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageResponseBodyResult) SetOpenPushId(v string) *SendOfficialAccountSNSMessageResponseBodyResult {
	s.OpenPushId = &v
	return s
}

type SendOfficialAccountSNSMessageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendOfficialAccountSNSMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendOfficialAccountSNSMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountSNSMessageResponse) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountSNSMessageResponse) SetHeaders(v map[string]*string) *SendOfficialAccountSNSMessageResponse {
	s.Headers = v
	return s
}

func (s *SendOfficialAccountSNSMessageResponse) SetStatusCode(v int32) *SendOfficialAccountSNSMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendOfficialAccountSNSMessageResponse) SetBody(v *SendOfficialAccountSNSMessageResponseBody) *SendOfficialAccountSNSMessageResponse {
	s.Body = v
	return s
}

type ServiceWindowMessageBatchPushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ServiceWindowMessageBatchPushHeaders) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushHeaders) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushHeaders) SetCommonHeaders(v map[string]*string) *ServiceWindowMessageBatchPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ServiceWindowMessageBatchPushHeaders) SetXAcsDingtalkAccessToken(v string) *ServiceWindowMessageBatchPushHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ServiceWindowMessageBatchPushRequest struct {
	// if can be null:
	// true
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	Detail *ServiceWindowMessageBatchPushRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
}

func (s ServiceWindowMessageBatchPushRequest) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequest) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequest) SetBizId(v string) *ServiceWindowMessageBatchPushRequest {
	s.BizId = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequest) SetDetail(v *ServiceWindowMessageBatchPushRequestDetail) *ServiceWindowMessageBatchPushRequest {
	s.Detail = v
	return s
}

type ServiceWindowMessageBatchPushRequestDetail struct {
	// if can be null:
	// false
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	// This parameter is required.
	MessageBody *ServiceWindowMessageBatchPushRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// if can be null:
	// true
	MsgType   *string `json:"msgType,omitempty" xml:"msgType,omitempty"`
	SendToAll *bool   `json:"sendToAll,omitempty" xml:"sendToAll,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// if can be null:
	// true
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetail) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetBizRequestId(v string) *ServiceWindowMessageBatchPushRequestDetail {
	s.BizRequestId = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetMessageBody(v *ServiceWindowMessageBatchPushRequestDetailMessageBody) *ServiceWindowMessageBatchPushRequestDetail {
	s.MessageBody = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetMsgType(v string) *ServiceWindowMessageBatchPushRequestDetail {
	s.MsgType = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetSendToAll(v bool) *ServiceWindowMessageBatchPushRequestDetail {
	s.SendToAll = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetUserIdList(v []*string) *ServiceWindowMessageBatchPushRequestDetail {
	s.UserIdList = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetUuid(v string) *ServiceWindowMessageBatchPushRequestDetail {
	s.Uuid = &v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBody struct {
	ActionCard *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard `json:"actionCard,omitempty" xml:"actionCard,omitempty" type:"Struct"`
	Link       *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink       `json:"link,omitempty" xml:"link,omitempty" type:"Struct"`
	Markdown   *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown   `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	Text       *ServiceWindowMessageBatchPushRequestDetailMessageBodyText       `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBody) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBody) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBody) SetActionCard(v *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) *ServiceWindowMessageBatchPushRequestDetailMessageBody {
	s.ActionCard = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBody) SetLink(v *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) *ServiceWindowMessageBatchPushRequestDetailMessageBody {
	s.Link = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBody) SetMarkdown(v *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown) *ServiceWindowMessageBatchPushRequestDetailMessageBody {
	s.Markdown = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBody) SetText(v *ServiceWindowMessageBatchPushRequestDetailMessageBodyText) *ServiceWindowMessageBatchPushRequestDetailMessageBody {
	s.Text = v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard struct {
	ButtonList        []*ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList `json:"buttonList,omitempty" xml:"buttonList,omitempty" type:"Repeated"`
	ButtonOrientation *string                                                                      `json:"buttonOrientation,omitempty" xml:"buttonOrientation,omitempty"`
	Markdown          *string                                                                      `json:"markdown,omitempty" xml:"markdown,omitempty"`
	SingleTitle       *string                                                                      `json:"singleTitle,omitempty" xml:"singleTitle,omitempty"`
	SingleUrl         *string                                                                      `json:"singleUrl,omitempty" xml:"singleUrl,omitempty"`
	Title             *string                                                                      `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetButtonList(v []*ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.ButtonList = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetButtonOrientation(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.ButtonOrientation = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetMarkdown(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.Markdown = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetSingleTitle(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.SingleTitle = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetSingleUrl(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.SingleUrl = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetTitle(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.Title = &v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList struct {
	// This parameter is required.
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList) SetActionUrl(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList {
	s.ActionUrl = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList) SetTitle(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList {
	s.Title = &v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBodyLink struct {
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	PicUrl     *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	Text       *string `json:"text,omitempty" xml:"text,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) SetMessageUrl(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink {
	s.MessageUrl = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) SetPicUrl(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink {
	s.PicUrl = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) SetText(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink {
	s.Text = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) SetTitle(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink {
	s.Title = &v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown struct {
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown) SetText(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown {
	s.Text = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown) SetTitle(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown {
	s.Title = &v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBodyText struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyText) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyText) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyText) SetContent(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyText {
	s.Content = &v
	return s
}

type ServiceWindowMessageBatchPushResponseBody struct {
	RequestId *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ServiceWindowMessageBatchPushResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ServiceWindowMessageBatchPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushResponseBody) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushResponseBody) SetRequestId(v string) *ServiceWindowMessageBatchPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *ServiceWindowMessageBatchPushResponseBody) SetResult(v *ServiceWindowMessageBatchPushResponseBodyResult) *ServiceWindowMessageBatchPushResponseBody {
	s.Result = v
	return s
}

type ServiceWindowMessageBatchPushResponseBodyResult struct {
	// This parameter is required.
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s ServiceWindowMessageBatchPushResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushResponseBodyResult) SetOpenPushId(v string) *ServiceWindowMessageBatchPushResponseBodyResult {
	s.OpenPushId = &v
	return s
}

type ServiceWindowMessageBatchPushResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ServiceWindowMessageBatchPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ServiceWindowMessageBatchPushResponse) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushResponse) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushResponse) SetHeaders(v map[string]*string) *ServiceWindowMessageBatchPushResponse {
	s.Headers = v
	return s
}

func (s *ServiceWindowMessageBatchPushResponse) SetStatusCode(v int32) *ServiceWindowMessageBatchPushResponse {
	s.StatusCode = &v
	return s
}

func (s *ServiceWindowMessageBatchPushResponse) SetBody(v *ServiceWindowMessageBatchPushResponseBody) *ServiceWindowMessageBatchPushResponse {
	s.Body = v
	return s
}

type SetUserVersionToFreeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetUserVersionToFreeHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetUserVersionToFreeHeaders) GoString() string {
	return s.String()
}

func (s *SetUserVersionToFreeHeaders) SetCommonHeaders(v map[string]*string) *SetUserVersionToFreeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetUserVersionToFreeHeaders) SetXAcsDingtalkAccessToken(v string) *SetUserVersionToFreeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetUserVersionToFreeRequest struct {
	// example:
	//
	// 012829186736-1115677667
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// example:
	//
	// other
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SetUserVersionToFreeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserVersionToFreeRequest) GoString() string {
	return s.String()
}

func (s *SetUserVersionToFreeRequest) SetOperatorUserId(v string) *SetUserVersionToFreeRequest {
	s.OperatorUserId = &v
	return s
}

func (s *SetUserVersionToFreeRequest) SetVersion(v string) *SetUserVersionToFreeRequest {
	s.Version = &v
	return s
}

type SetUserVersionToFreeResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SetUserVersionToFreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetUserVersionToFreeResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserVersionToFreeResponseBody) SetResult(v bool) *SetUserVersionToFreeResponseBody {
	s.Result = &v
	return s
}

type SetUserVersionToFreeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetUserVersionToFreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetUserVersionToFreeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserVersionToFreeResponse) GoString() string {
	return s.String()
}

func (s *SetUserVersionToFreeResponse) SetHeaders(v map[string]*string) *SetUserVersionToFreeResponse {
	s.Headers = v
	return s
}

func (s *SetUserVersionToFreeResponse) SetStatusCode(v int32) *SetUserVersionToFreeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserVersionToFreeResponse) SetBody(v *SetUserVersionToFreeResponseBody) *SetUserVersionToFreeResponse {
	s.Body = v
	return s
}

type TwoPhaseCommitInventoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TwoPhaseCommitInventoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s TwoPhaseCommitInventoryHeaders) GoString() string {
	return s.String()
}

func (s *TwoPhaseCommitInventoryHeaders) SetCommonHeaders(v map[string]*string) *TwoPhaseCommitInventoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TwoPhaseCommitInventoryHeaders) SetXAcsDingtalkAccessToken(v string) *TwoPhaseCommitInventoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TwoPhaseCommitInventoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// B_ACCOUNT_NUMBER
	BenefitCode *string `json:"benefitCode,omitempty" xml:"benefitCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bizId
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	ExecuteResult *bool `json:"executeResult,omitempty" xml:"executeResult,omitempty"`
	// example:
	//
	// 10
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
}

func (s TwoPhaseCommitInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s TwoPhaseCommitInventoryRequest) GoString() string {
	return s.String()
}

func (s *TwoPhaseCommitInventoryRequest) SetBenefitCode(v string) *TwoPhaseCommitInventoryRequest {
	s.BenefitCode = &v
	return s
}

func (s *TwoPhaseCommitInventoryRequest) SetBizRequestId(v string) *TwoPhaseCommitInventoryRequest {
	s.BizRequestId = &v
	return s
}

func (s *TwoPhaseCommitInventoryRequest) SetExecuteResult(v bool) *TwoPhaseCommitInventoryRequest {
	s.ExecuteResult = &v
	return s
}

func (s *TwoPhaseCommitInventoryRequest) SetQuota(v int64) *TwoPhaseCommitInventoryRequest {
	s.Quota = &v
	return s
}

type TwoPhaseCommitInventoryResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s TwoPhaseCommitInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TwoPhaseCommitInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *TwoPhaseCommitInventoryResponseBody) SetResult(v bool) *TwoPhaseCommitInventoryResponseBody {
	s.Result = &v
	return s
}

type TwoPhaseCommitInventoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TwoPhaseCommitInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TwoPhaseCommitInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s TwoPhaseCommitInventoryResponse) GoString() string {
	return s.String()
}

func (s *TwoPhaseCommitInventoryResponse) SetHeaders(v map[string]*string) *TwoPhaseCommitInventoryResponse {
	s.Headers = v
	return s
}

func (s *TwoPhaseCommitInventoryResponse) SetStatusCode(v int32) *TwoPhaseCommitInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *TwoPhaseCommitInventoryResponse) SetBody(v *TwoPhaseCommitInventoryResponseBody) *TwoPhaseCommitInventoryResponse {
	s.Body = v
	return s
}

type UpdateCrmPersonalCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCrmPersonalCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrmPersonalCustomerHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCrmPersonalCustomerHeaders) SetCommonHeaders(v map[string]*string) *UpdateCrmPersonalCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCrmPersonalCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCrmPersonalCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCrmPersonalCustomerRequest struct {
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// This parameter is required.
	Data       map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ModifierNick *string `json:"modifierNick,omitempty" xml:"modifierNick,omitempty"`
	// This parameter is required.
	ModifierUserId *string                                     `json:"modifierUserId,omitempty" xml:"modifierUserId,omitempty"`
	Permission     *UpdateCrmPersonalCustomerRequestPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	RelationType   *string                                     `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// example:
	//
	// false
	SkipDuplicateCheck *bool `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
}

func (s UpdateCrmPersonalCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrmPersonalCustomerRequest) GoString() string {
	return s.String()
}

func (s *UpdateCrmPersonalCustomerRequest) SetAction(v string) *UpdateCrmPersonalCustomerRequest {
	s.Action = &v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetData(v map[string]interface{}) *UpdateCrmPersonalCustomerRequest {
	s.Data = v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetExtendData(v map[string]interface{}) *UpdateCrmPersonalCustomerRequest {
	s.ExtendData = v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetInstanceId(v string) *UpdateCrmPersonalCustomerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetModifierNick(v string) *UpdateCrmPersonalCustomerRequest {
	s.ModifierNick = &v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetModifierUserId(v string) *UpdateCrmPersonalCustomerRequest {
	s.ModifierUserId = &v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetPermission(v *UpdateCrmPersonalCustomerRequestPermission) *UpdateCrmPersonalCustomerRequest {
	s.Permission = v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetRelationType(v string) *UpdateCrmPersonalCustomerRequest {
	s.RelationType = &v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetSkipDuplicateCheck(v bool) *UpdateCrmPersonalCustomerRequest {
	s.SkipDuplicateCheck = &v
	return s
}

type UpdateCrmPersonalCustomerRequestPermission struct {
	OwnerStaffIds       []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s UpdateCrmPersonalCustomerRequestPermission) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrmPersonalCustomerRequestPermission) GoString() string {
	return s.String()
}

func (s *UpdateCrmPersonalCustomerRequestPermission) SetOwnerStaffIds(v []*string) *UpdateCrmPersonalCustomerRequestPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *UpdateCrmPersonalCustomerRequestPermission) SetParticipantStaffIds(v []*string) *UpdateCrmPersonalCustomerRequestPermission {
	s.ParticipantStaffIds = v
	return s
}

type UpdateCrmPersonalCustomerResponseBody struct {
	// This parameter is required.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s UpdateCrmPersonalCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrmPersonalCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCrmPersonalCustomerResponseBody) SetInstanceId(v string) *UpdateCrmPersonalCustomerResponseBody {
	s.InstanceId = &v
	return s
}

type UpdateCrmPersonalCustomerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCrmPersonalCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrmPersonalCustomerResponse) GoString() string {
	return s.String()
}

func (s *UpdateCrmPersonalCustomerResponse) SetHeaders(v map[string]*string) *UpdateCrmPersonalCustomerResponse {
	s.Headers = v
	return s
}

func (s *UpdateCrmPersonalCustomerResponse) SetStatusCode(v int32) *UpdateCrmPersonalCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCrmPersonalCustomerResponse) SetBody(v *UpdateCrmPersonalCustomerResponseBody) *UpdateCrmPersonalCustomerResponse {
	s.Body = v
	return s
}

type UpdateCustomerBizTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCustomerBizTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomerBizTypeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCustomerBizTypeHeaders) SetCommonHeaders(v map[string]*string) *UpdateCustomerBizTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCustomerBizTypeHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCustomerBizTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCustomerBizTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// crm_customer
	CustomerBizType *string `json:"customerBizType,omitempty" xml:"customerBizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 34234234ddddad
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s UpdateCustomerBizTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomerBizTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomerBizTypeRequest) SetCustomerBizType(v string) *UpdateCustomerBizTypeRequest {
	s.CustomerBizType = &v
	return s
}

func (s *UpdateCustomerBizTypeRequest) SetOperatorUserId(v string) *UpdateCustomerBizTypeRequest {
	s.OperatorUserId = &v
	return s
}

type UpdateCustomerBizTypeResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateCustomerBizTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomerBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomerBizTypeResponseBody) SetResult(v bool) *UpdateCustomerBizTypeResponseBody {
	s.Result = &v
	return s
}

type UpdateCustomerBizTypeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomerBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomerBizTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomerBizTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomerBizTypeResponse) SetHeaders(v map[string]*string) *UpdateCustomerBizTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomerBizTypeResponse) SetStatusCode(v int32) *UpdateCustomerBizTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomerBizTypeResponse) SetBody(v *UpdateCustomerBizTypeResponseBody) *UpdateCustomerBizTypeResponse {
	s.Body = v
	return s
}

type UpdateGroupSetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupSetHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSetHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupSetHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupSetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupSetHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupSetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupSetRequest struct {
	ManagerUserIds *string `json:"managerUserIds,omitempty" xml:"managerUserIds,omitempty"`
	MemberQuota    *int32  `json:"memberQuota,omitempty" xml:"memberQuota,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Notice         *string `json:"notice,omitempty" xml:"notice,omitempty"`
	NoticeToped    *int32  `json:"noticeToped,omitempty" xml:"noticeToped,omitempty"`
	// This parameter is required.
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	OwnerUserId    *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	TemplateId     *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Welcome        *string `json:"welcome,omitempty" xml:"welcome,omitempty"`
}

func (s UpdateGroupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupSetRequest) SetManagerUserIds(v string) *UpdateGroupSetRequest {
	s.ManagerUserIds = &v
	return s
}

func (s *UpdateGroupSetRequest) SetMemberQuota(v int32) *UpdateGroupSetRequest {
	s.MemberQuota = &v
	return s
}

func (s *UpdateGroupSetRequest) SetName(v string) *UpdateGroupSetRequest {
	s.Name = &v
	return s
}

func (s *UpdateGroupSetRequest) SetNotice(v string) *UpdateGroupSetRequest {
	s.Notice = &v
	return s
}

func (s *UpdateGroupSetRequest) SetNoticeToped(v int32) *UpdateGroupSetRequest {
	s.NoticeToped = &v
	return s
}

func (s *UpdateGroupSetRequest) SetOpenGroupSetId(v string) *UpdateGroupSetRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *UpdateGroupSetRequest) SetOwnerUserId(v string) *UpdateGroupSetRequest {
	s.OwnerUserId = &v
	return s
}

func (s *UpdateGroupSetRequest) SetTemplateId(v string) *UpdateGroupSetRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateGroupSetRequest) SetWelcome(v string) *UpdateGroupSetRequest {
	s.Welcome = &v
	return s
}

type UpdateGroupSetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *bool              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupSetResponse) SetHeaders(v map[string]*string) *UpdateGroupSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupSetResponse) SetStatusCode(v int32) *UpdateGroupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupSetResponse) SetBody(v bool) *UpdateGroupSetResponse {
	s.Body = &v
	return s
}

type UpdateMenuDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMenuDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMenuDataHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMenuDataHeaders) SetCommonHeaders(v map[string]*string) *UpdateMenuDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMenuDataHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMenuDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMenuDataRequest struct {
	Attr map[string]interface{} `json:"attr,omitempty" xml:"attr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 89ez9DwVWQVgkhwndJNt1ZY
	BizTraceId *string `json:"bizTraceId,omitempty" xml:"bizTraceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sale
	Module *string `json:"module,omitempty" xml:"module,omitempty"`
	// This parameter is required.
	NavData *UpdateMenuDataRequestNavData `json:"navData,omitempty" xml:"navData,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// add
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16044739461008808646
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s UpdateMenuDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMenuDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateMenuDataRequest) SetAttr(v map[string]interface{}) *UpdateMenuDataRequest {
	s.Attr = v
	return s
}

func (s *UpdateMenuDataRequest) SetBizTraceId(v string) *UpdateMenuDataRequest {
	s.BizTraceId = &v
	return s
}

func (s *UpdateMenuDataRequest) SetModule(v string) *UpdateMenuDataRequest {
	s.Module = &v
	return s
}

func (s *UpdateMenuDataRequest) SetNavData(v *UpdateMenuDataRequestNavData) *UpdateMenuDataRequest {
	s.NavData = v
	return s
}

func (s *UpdateMenuDataRequest) SetOperateType(v string) *UpdateMenuDataRequest {
	s.OperateType = &v
	return s
}

func (s *UpdateMenuDataRequest) SetOperatorUserId(v string) *UpdateMenuDataRequest {
	s.OperatorUserId = &v
	return s
}

type UpdateMenuDataRequestNavData struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DisplayStatus *string `json:"displayStatus,omitempty" xml:"displayStatus,omitempty"`
	// example:
	//
	// icon-biaodan_baogao
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// #CCEEFF
	IconBgColor *string `json:"iconBgColor,omitempty" xml:"iconBgColor,omitempty"`
	// example:
	//
	// #007FFF
	IconColor *string `json:"iconColor,omitempty" xml:"iconColor,omitempty"`
	// example:
	//
	// same_page
	IntegrationProtocol *string `json:"integrationProtocol,omitempty" xml:"integrationProtocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 库存账单
	MobileNavName *string `json:"mobileNavName,omitempty" xml:"mobileNavName,omitempty"`
	// example:
	//
	// https://tj-ali-crm-test.tangees.com/tungee/crm/mobile/?corpid=dinge18b222ec5637b04ffe93478753d9884#/form/PROC-ECC13160-7AFF-4D5B-8E73-E5B98BB9A59B?库存流水&psi_stock_flow&fromPage=home
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lowcode_customize_form
	NavCode    *string                                 `json:"navCode,omitempty" xml:"navCode,omitempty"`
	NavExtInfo *UpdateMenuDataRequestNavDataNavExtInfo `json:"navExtInfo,omitempty" xml:"navExtInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// lowcode_customize_form:PROC-0279E824-ED47-4E75-86F2-11B665F3704D
	NavId *string `json:"navId,omitempty" xml:"navId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 库存流水
	NavName *string `json:"navName,omitempty" xml:"navName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PUBLISHED
	NavStatus *string `json:"navStatus,omitempty" xml:"navStatus,omitempty"`
	// example:
	//
	// item
	NavType *string `json:"navType,omitempty" xml:"navType,omitempty"`
	// example:
	//
	// crm_product
	ParentNavId *string `json:"parentNavId,omitempty" xml:"parentNavId,omitempty"`
	// example:
	//
	// tj
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	SortNum  *int32  `json:"sortNum,omitempty" xml:"sortNum,omitempty"`
	// example:
	//
	// /form/PROC-ECC13160-7AFF-4D5B-8E73-E5B98BB9A59B?bizType=psi_stock_flow&formName=库存流水
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UpdateMenuDataRequestNavData) String() string {
	return tea.Prettify(s)
}

func (s UpdateMenuDataRequestNavData) GoString() string {
	return s.String()
}

func (s *UpdateMenuDataRequestNavData) SetDisplayStatus(v string) *UpdateMenuDataRequestNavData {
	s.DisplayStatus = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetIcon(v string) *UpdateMenuDataRequestNavData {
	s.Icon = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetIconBgColor(v string) *UpdateMenuDataRequestNavData {
	s.IconBgColor = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetIconColor(v string) *UpdateMenuDataRequestNavData {
	s.IconColor = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetIntegrationProtocol(v string) *UpdateMenuDataRequestNavData {
	s.IntegrationProtocol = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetMobileNavName(v string) *UpdateMenuDataRequestNavData {
	s.MobileNavName = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetMobileUrl(v string) *UpdateMenuDataRequestNavData {
	s.MobileUrl = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetNavCode(v string) *UpdateMenuDataRequestNavData {
	s.NavCode = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetNavExtInfo(v *UpdateMenuDataRequestNavDataNavExtInfo) *UpdateMenuDataRequestNavData {
	s.NavExtInfo = v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetNavId(v string) *UpdateMenuDataRequestNavData {
	s.NavId = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetNavName(v string) *UpdateMenuDataRequestNavData {
	s.NavName = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetNavStatus(v string) *UpdateMenuDataRequestNavData {
	s.NavStatus = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetNavType(v string) *UpdateMenuDataRequestNavData {
	s.NavType = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetParentNavId(v string) *UpdateMenuDataRequestNavData {
	s.ParentNavId = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetProvider(v string) *UpdateMenuDataRequestNavData {
	s.Provider = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetSortNum(v int32) *UpdateMenuDataRequestNavData {
	s.SortNum = &v
	return s
}

func (s *UpdateMenuDataRequestNavData) SetUrl(v string) *UpdateMenuDataRequestNavData {
	s.Url = &v
	return s
}

type UpdateMenuDataRequestNavDataNavExtInfo struct {
	// example:
	//
	// oem
	ProductMode *string `json:"productMode,omitempty" xml:"productMode,omitempty"`
	// example:
	//
	// tj
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s UpdateMenuDataRequestNavDataNavExtInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateMenuDataRequestNavDataNavExtInfo) GoString() string {
	return s.String()
}

func (s *UpdateMenuDataRequestNavDataNavExtInfo) SetProductMode(v string) *UpdateMenuDataRequestNavDataNavExtInfo {
	s.ProductMode = &v
	return s
}

func (s *UpdateMenuDataRequestNavDataNavExtInfo) SetProvider(v string) *UpdateMenuDataRequestNavDataNavExtInfo {
	s.Provider = &v
	return s
}

type UpdateMenuDataResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateMenuDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMenuDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMenuDataResponseBody) SetResult(v string) *UpdateMenuDataResponseBody {
	s.Result = &v
	return s
}

type UpdateMenuDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMenuDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMenuDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMenuDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateMenuDataResponse) SetHeaders(v map[string]*string) *UpdateMenuDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateMenuDataResponse) SetStatusCode(v int32) *UpdateMenuDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMenuDataResponse) SetBody(v *UpdateMenuDataResponseBody) *UpdateMenuDataResponse {
	s.Body = v
	return s
}

type UpdateMetaModelFieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMetaModelFieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetaModelFieldHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMetaModelFieldHeaders) SetCommonHeaders(v map[string]*string) *UpdateMetaModelFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMetaModelFieldHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMetaModelFieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMetaModelFieldRequest struct {
	// This parameter is required.
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	FieldDTOList []*UpdateMetaModelFieldRequestFieldDTOList `json:"fieldDTOList,omitempty" xml:"fieldDTOList,omitempty" type:"Repeated"`
	// This parameter is required.
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	Tenant *string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

func (s UpdateMetaModelFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetaModelFieldRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaModelFieldRequest) SetBizType(v string) *UpdateMetaModelFieldRequest {
	s.BizType = &v
	return s
}

func (s *UpdateMetaModelFieldRequest) SetFieldDTOList(v []*UpdateMetaModelFieldRequestFieldDTOList) *UpdateMetaModelFieldRequest {
	s.FieldDTOList = v
	return s
}

func (s *UpdateMetaModelFieldRequest) SetOperatorUserId(v string) *UpdateMetaModelFieldRequest {
	s.OperatorUserId = &v
	return s
}

func (s *UpdateMetaModelFieldRequest) SetTenant(v string) *UpdateMetaModelFieldRequest {
	s.Tenant = &v
	return s
}

type UpdateMetaModelFieldRequestFieldDTOList struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *UpdateMetaModelFieldRequestFieldDTOListProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s UpdateMetaModelFieldRequestFieldDTOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetaModelFieldRequestFieldDTOList) GoString() string {
	return s.String()
}

func (s *UpdateMetaModelFieldRequestFieldDTOList) SetComponentName(v string) *UpdateMetaModelFieldRequestFieldDTOList {
	s.ComponentName = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOList) SetProps(v *UpdateMetaModelFieldRequestFieldDTOListProps) *UpdateMetaModelFieldRequestFieldDTOList {
	s.Props = v
	return s
}

type UpdateMetaModelFieldRequestFieldDTOListProps struct {
	Align    *string `json:"align,omitempty" xml:"align,omitempty"`
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice   *int64  `json:"choice,omitempty" xml:"choice,omitempty"`
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	Disabled *bool   `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration *bool   `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	FieldId   *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format    *string `json:"format,omitempty" xml:"format,omitempty"`
	Invisible *bool   `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label               *string                                                `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze *bool                                                  `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Link                *string                                                `json:"link,omitempty" xml:"link,omitempty"`
	NeedDetail          *string                                                `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	NotPrint            *string                                                `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	NotUpper            *string                                                `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options             []*UpdateMetaModelFieldRequestFieldDTOListPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable           *bool                                                  `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder         *string                                                `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Required               *bool   `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool   `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Sortable               *bool   `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Unit                   *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s UpdateMetaModelFieldRequestFieldDTOListProps) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetaModelFieldRequestFieldDTOListProps) GoString() string {
	return s.String()
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetAlign(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Align = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetBizAlias(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.BizAlias = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetChoice(v int64) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Choice = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetContent(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Content = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetDisabled(v bool) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Disabled = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetDuration(v bool) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Duration = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetFieldId(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.FieldId = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetFormat(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Format = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetInvisible(v bool) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Invisible = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetLabel(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Label = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetLabelEditableFreeze(v bool) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetLink(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Link = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetNeedDetail(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.NeedDetail = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetNotPrint(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.NotPrint = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetNotUpper(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.NotUpper = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetOptions(v []*UpdateMetaModelFieldRequestFieldDTOListPropsOptions) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Options = v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetPayEnable(v bool) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.PayEnable = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetPlaceholder(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Placeholder = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetRequired(v bool) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Required = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetRequiredEditableFreeze(v bool) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetSortable(v bool) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Sortable = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListProps) SetUnit(v string) *UpdateMetaModelFieldRequestFieldDTOListProps {
	s.Unit = &v
	return s
}

type UpdateMetaModelFieldRequestFieldDTOListPropsOptions struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateMetaModelFieldRequestFieldDTOListPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetaModelFieldRequestFieldDTOListPropsOptions) GoString() string {
	return s.String()
}

func (s *UpdateMetaModelFieldRequestFieldDTOListPropsOptions) SetKey(v string) *UpdateMetaModelFieldRequestFieldDTOListPropsOptions {
	s.Key = &v
	return s
}

func (s *UpdateMetaModelFieldRequestFieldDTOListPropsOptions) SetValue(v string) *UpdateMetaModelFieldRequestFieldDTOListPropsOptions {
	s.Value = &v
	return s
}

type UpdateMetaModelFieldResponseBody struct {
	// This parameter is required.
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
}

func (s UpdateMetaModelFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetaModelFieldResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetaModelFieldResponseBody) SetBizType(v string) *UpdateMetaModelFieldResponseBody {
	s.BizType = &v
	return s
}

type UpdateMetaModelFieldResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetaModelFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetaModelFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetaModelFieldResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetaModelFieldResponse) SetHeaders(v map[string]*string) *UpdateMetaModelFieldResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetaModelFieldResponse) SetStatusCode(v int32) *UpdateMetaModelFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetaModelFieldResponse) SetBody(v *UpdateMetaModelFieldResponseBody) *UpdateMetaModelFieldResponse {
	s.Body = v
	return s
}

type UpdateRelationMetaFieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRelationMetaFieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRelationMetaFieldHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRelationMetaFieldHeaders) SetCommonHeaders(v map[string]*string) *UpdateRelationMetaFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRelationMetaFieldHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRelationMetaFieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRelationMetaFieldRequest struct {
	// This parameter is required.
	FieldDTOList []*UpdateRelationMetaFieldRequestFieldDTOList `json:"fieldDTOList,omitempty" xml:"fieldDTOList,omitempty" type:"Repeated"`
	// This parameter is required.
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// This parameter is required.
	Tenant *string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

func (s UpdateRelationMetaFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRelationMetaFieldRequest) GoString() string {
	return s.String()
}

func (s *UpdateRelationMetaFieldRequest) SetFieldDTOList(v []*UpdateRelationMetaFieldRequestFieldDTOList) *UpdateRelationMetaFieldRequest {
	s.FieldDTOList = v
	return s
}

func (s *UpdateRelationMetaFieldRequest) SetOperatorUserId(v string) *UpdateRelationMetaFieldRequest {
	s.OperatorUserId = &v
	return s
}

func (s *UpdateRelationMetaFieldRequest) SetRelationType(v string) *UpdateRelationMetaFieldRequest {
	s.RelationType = &v
	return s
}

func (s *UpdateRelationMetaFieldRequest) SetTenant(v string) *UpdateRelationMetaFieldRequest {
	s.Tenant = &v
	return s
}

type UpdateRelationMetaFieldRequestFieldDTOList struct {
	// This parameter is required.
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *UpdateRelationMetaFieldRequestFieldDTOListProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s UpdateRelationMetaFieldRequestFieldDTOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateRelationMetaFieldRequestFieldDTOList) GoString() string {
	return s.String()
}

func (s *UpdateRelationMetaFieldRequestFieldDTOList) SetComponentName(v string) *UpdateRelationMetaFieldRequestFieldDTOList {
	s.ComponentName = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOList) SetProps(v *UpdateRelationMetaFieldRequestFieldDTOListProps) *UpdateRelationMetaFieldRequestFieldDTOList {
	s.Props = v
	return s
}

type UpdateRelationMetaFieldRequestFieldDTOListProps struct {
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice   *int64  `json:"choice,omitempty" xml:"choice,omitempty"`
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	Disabled *bool   `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration *bool   `json:"duration,omitempty" xml:"duration,omitempty"`
	// This parameter is required.
	FieldId   *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format    *string `json:"format,omitempty" xml:"format,omitempty"`
	Invisible *bool   `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// This parameter is required.
	Label               *string                                                   `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze *bool                                                     `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Link                *string                                                   `json:"link,omitempty" xml:"link,omitempty"`
	NeedDetail          *string                                                   `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	NotPrint            *string                                                   `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	NotUpper            *string                                                   `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options             []*UpdateRelationMetaFieldRequestFieldDTOListPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable           *bool                                                     `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder         *string                                                   `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// This parameter is required.
	Required               *bool   `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool   `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Sortable               *bool   `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Unit                   *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s UpdateRelationMetaFieldRequestFieldDTOListProps) String() string {
	return tea.Prettify(s)
}

func (s UpdateRelationMetaFieldRequestFieldDTOListProps) GoString() string {
	return s.String()
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetAlign(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Align = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetBizAlias(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.BizAlias = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetChoice(v int64) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Choice = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetContent(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Content = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetDisabled(v bool) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Disabled = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetDuration(v bool) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Duration = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetFieldId(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.FieldId = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetFormat(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Format = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetInvisible(v bool) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Invisible = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetLabel(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Label = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetLabelEditableFreeze(v bool) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetLink(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Link = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetNeedDetail(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.NeedDetail = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetNotPrint(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.NotPrint = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetNotUpper(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.NotUpper = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetOptions(v []*UpdateRelationMetaFieldRequestFieldDTOListPropsOptions) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Options = v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetPayEnable(v bool) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.PayEnable = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetPlaceholder(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Placeholder = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetRequired(v bool) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Required = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetRequiredEditableFreeze(v bool) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetSortable(v bool) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Sortable = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListProps) SetUnit(v string) *UpdateRelationMetaFieldRequestFieldDTOListProps {
	s.Unit = &v
	return s
}

type UpdateRelationMetaFieldRequestFieldDTOListPropsOptions struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateRelationMetaFieldRequestFieldDTOListPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateRelationMetaFieldRequestFieldDTOListPropsOptions) GoString() string {
	return s.String()
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListPropsOptions) SetKey(v string) *UpdateRelationMetaFieldRequestFieldDTOListPropsOptions {
	s.Key = &v
	return s
}

func (s *UpdateRelationMetaFieldRequestFieldDTOListPropsOptions) SetValue(v string) *UpdateRelationMetaFieldRequestFieldDTOListPropsOptions {
	s.Value = &v
	return s
}

type UpdateRelationMetaFieldResponseBody struct {
	// This parameter is required.
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s UpdateRelationMetaFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRelationMetaFieldResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRelationMetaFieldResponseBody) SetRelationType(v string) *UpdateRelationMetaFieldResponseBody {
	s.RelationType = &v
	return s
}

type UpdateRelationMetaFieldResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRelationMetaFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRelationMetaFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRelationMetaFieldResponse) GoString() string {
	return s.String()
}

func (s *UpdateRelationMetaFieldResponse) SetHeaders(v map[string]*string) *UpdateRelationMetaFieldResponse {
	s.Headers = v
	return s
}

func (s *UpdateRelationMetaFieldResponse) SetStatusCode(v int32) *UpdateRelationMetaFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRelationMetaFieldResponse) SetBody(v *UpdateRelationMetaFieldResponseBody) *UpdateRelationMetaFieldResponse {
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
// 从私海放弃客户（退回公海）
//
// @param request - AbandonCustomerRequest
//
// @param headers - AbandonCustomerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbandonCustomerResponse
func (client *Client) AbandonCustomerWithOptions(request *AbandonCustomerRequest, headers *AbandonCustomerHeaders, runtime *util.RuntimeOptions) (_result *AbandonCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomTrackDesc)) {
		body["customTrackDesc"] = request.CustomTrackDesc
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		body["instanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OptType)) {
		body["optType"] = request.OptType
	}

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
		Action:      tea.String("AbandonCustomer"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customers/abandon"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AbandonCustomerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从私海放弃客户（退回公海）
//
// @param request - AbandonCustomerRequest
//
// @return AbandonCustomerResponse
func (client *Client) AbandonCustomer(request *AbandonCustomerRequest) (_result *AbandonCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AbandonCustomerHeaders{}
	_result = &AbandonCustomerResponse{}
	_body, _err := client.AbandonCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加crm个人客户（或企业客户）
//
// @param request - AddCrmPersonalCustomerRequest
//
// @param headers - AddCrmPersonalCustomerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCrmPersonalCustomerResponse
func (client *Client) AddCrmPersonalCustomerWithOptions(request *AddCrmPersonalCustomerRequest, headers *AddCrmPersonalCustomerHeaders, runtime *util.RuntimeOptions) (_result *AddCrmPersonalCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorNick)) {
		body["creatorNick"] = request.CreatorNick
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorUserId)) {
		body["creatorUserId"] = request.CreatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(request.Permission)) {
		body["permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.SkipDuplicateCheck)) {
		body["skipDuplicateCheck"] = request.SkipDuplicateCheck
	}

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
		Action:      tea.String("AddCrmPersonalCustomer"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/personalCustomers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCrmPersonalCustomerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加crm个人客户（或企业客户）
//
// @param request - AddCrmPersonalCustomerRequest
//
// @return AddCrmPersonalCustomerResponse
func (client *Client) AddCrmPersonalCustomer(request *AddCrmPersonalCustomerRequest) (_result *AddCrmPersonalCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCrmPersonalCustomerHeaders{}
	_result = &AddCrmPersonalCustomerResponse{}
	_body, _err := client.AddCrmPersonalCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增动态
//
// @param request - AddCustomerTrackRequest
//
// @param headers - AddCustomerTrackHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomerTrackResponse
func (client *Client) AddCustomerTrackWithOptions(request *AddCustomerTrackRequest, headers *AddCustomerTrackHeaders, runtime *util.RuntimeOptions) (_result *AddCustomerTrackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		body["customerId"] = request.CustomerId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraBizInfo)) {
		body["extraBizInfo"] = request.ExtraBizInfo
	}

	if !tea.BoolValue(util.IsUnset(request.IdempotentKey)) {
		body["idempotentKey"] = request.IdempotentKey
	}

	if !tea.BoolValue(util.IsUnset(request.MaskedContent)) {
		body["maskedContent"] = request.MaskedContent
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Action:      tea.String("AddCustomerTrack"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customerTracks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCustomerTrackResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增动态
//
// @param request - AddCustomerTrackRequest
//
// @return AddCustomerTrackResponse
func (client *Client) AddCustomerTrack(request *AddCustomerTrackRequest) (_result *AddCustomerTrackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCustomerTrackHeaders{}
	_result = &AddCustomerTrackResponse{}
	_body, _err := client.AddCustomerTrackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加线索
//
// @param request - AddLeadsRequest
//
// @param headers - AddLeadsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLeadsResponse
func (client *Client) AddLeadsWithOptions(request *AddLeadsRequest, headers *AddLeadsHeaders, runtime *util.RuntimeOptions) (_result *AddLeadsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignTimestamp)) {
		body["assignTimestamp"] = request.AssignTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.AssignUserId)) {
		body["assignUserId"] = request.AssignUserId
	}

	if !tea.BoolValue(util.IsUnset(request.AssignedUserId)) {
		body["assignedUserId"] = request.AssignedUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Leads)) {
		body["leads"] = request.Leads
	}

	if !tea.BoolValue(util.IsUnset(request.OutTaskId)) {
		body["outTaskId"] = request.OutTaskId
	}

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
		Action:      tea.String("AddLeads"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/leads"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddLeadsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加线索
//
// @param request - AddLeadsRequest
//
// @return AddLeadsResponse
func (client *Client) AddLeads(request *AddLeadsRequest) (_result *AddLeadsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddLeadsHeaders{}
	_result = &AddLeadsResponse{}
	_body, _err := client.AddLeadsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模型表结构增加字段
//
// @param request - AddMetaModelFieldRequest
//
// @param headers - AddMetaModelFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMetaModelFieldResponse
func (client *Client) AddMetaModelFieldWithOptions(request *AddMetaModelFieldRequest, headers *AddMetaModelFieldHeaders, runtime *util.RuntimeOptions) (_result *AddMetaModelFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.FieldDTOList)) {
		body["fieldDTOList"] = request.FieldDTOList
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Tenant)) {
		body["tenant"] = request.Tenant
	}

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
		Action:      tea.String("AddMetaModelField"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/metas/models/fields"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMetaModelFieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模型表结构增加字段
//
// @param request - AddMetaModelFieldRequest
//
// @return AddMetaModelFieldResponse
func (client *Client) AddMetaModelField(request *AddMetaModelFieldRequest) (_result *AddMetaModelFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddMetaModelFieldHeaders{}
	_result = &AddMetaModelFieldResponse{}
	_body, _err := client.AddMetaModelFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关系模型表结构增加字段
//
// @param request - AddRelationMetaFieldRequest
//
// @param headers - AddRelationMetaFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRelationMetaFieldResponse
func (client *Client) AddRelationMetaFieldWithOptions(request *AddRelationMetaFieldRequest, headers *AddRelationMetaFieldHeaders, runtime *util.RuntimeOptions) (_result *AddRelationMetaFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FieldDTOList)) {
		body["fieldDTOList"] = request.FieldDTOList
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.Tenant)) {
		body["tenant"] = request.Tenant
	}

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
		Action:      tea.String("AddRelationMetaField"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/relations/metas/fields"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRelationMetaFieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关系模型表结构增加字段
//
// @param request - AddRelationMetaFieldRequest
//
// @return AddRelationMetaFieldResponse
func (client *Client) AddRelationMetaField(request *AddRelationMetaFieldRequest) (_result *AddRelationMetaFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddRelationMetaFieldHeaders{}
	_result = &AddRelationMetaFieldResponse{}
	_body, _err := client.AddRelationMetaFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 追加客户数据权限
//
// @param request - AppendCustomerDataAuthRequest
//
// @param headers - AppendCustomerDataAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppendCustomerDataAuthResponse
func (client *Client) AppendCustomerDataAuthWithOptions(request *AppendCustomerDataAuthRequest, headers *AppendCustomerDataAuthHeaders, runtime *util.RuntimeOptions) (_result *AppendCustomerDataAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerIds)) {
		body["customerIds"] = request.CustomerIds
	}

	if !tea.BoolValue(util.IsUnset(request.DataAuthUserIds)) {
		body["dataAuthUserIds"] = request.DataAuthUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		body["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["roleType"] = request.RoleType
	}

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
		Action:      tea.String("AppendCustomerDataAuth"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customers/dataAuth/append"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AppendCustomerDataAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 追加客户数据权限
//
// @param request - AppendCustomerDataAuthRequest
//
// @return AppendCustomerDataAuthResponse
func (client *Client) AppendCustomerDataAuth(request *AppendCustomerDataAuthRequest) (_result *AppendCustomerDataAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppendCustomerDataAuthHeaders{}
	_result = &AppendCustomerDataAuthResponse{}
	_body, _err := client.AppendCustomerDataAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量新增联系人
//
// @param request - BatchAddContactsRequest
//
// @param headers - BatchAddContactsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddContactsResponse
func (client *Client) BatchAddContactsWithOptions(request *BatchAddContactsRequest, headers *BatchAddContactsHeaders, runtime *util.RuntimeOptions) (_result *BatchAddContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationList)) {
		body["relationList"] = request.RelationList
	}

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
		Action:      tea.String("BatchAddContacts"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/contacts/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddContactsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量新增联系人
//
// @param request - BatchAddContactsRequest
//
// @return BatchAddContactsResponse
func (client *Client) BatchAddContacts(request *BatchAddContactsRequest) (_result *BatchAddContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchAddContactsHeaders{}
	_result = &BatchAddContactsResponse{}
	_body, _err := client.BatchAddContactsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量新增跟进记录
//
// @param request - BatchAddFollowRecordsRequest
//
// @param headers - BatchAddFollowRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddFollowRecordsResponse
func (client *Client) BatchAddFollowRecordsWithOptions(request *BatchAddFollowRecordsRequest, headers *BatchAddFollowRecordsHeaders, runtime *util.RuntimeOptions) (_result *BatchAddFollowRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceList)) {
		body["instanceList"] = request.InstanceList
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

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
		Action:      tea.String("BatchAddFollowRecords"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/followRecords/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddFollowRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量新增跟进记录
//
// @param request - BatchAddFollowRecordsRequest
//
// @return BatchAddFollowRecordsResponse
func (client *Client) BatchAddFollowRecords(request *BatchAddFollowRecordsRequest) (_result *BatchAddFollowRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchAddFollowRecordsHeaders{}
	_result = &BatchAddFollowRecordsResponse{}
	_body, _err := client.BatchAddFollowRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量新增关系数据
//
// @param request - BatchAddRelationDatasRequest
//
// @param headers - BatchAddRelationDatasHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddRelationDatasResponse
func (client *Client) BatchAddRelationDatasWithOptions(request *BatchAddRelationDatasRequest, headers *BatchAddRelationDatasHeaders, runtime *util.RuntimeOptions) (_result *BatchAddRelationDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationList)) {
		body["relationList"] = request.RelationList
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.SkipDuplicateCheck)) {
		body["skipDuplicateCheck"] = request.SkipDuplicateCheck
	}

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
		Action:      tea.String("BatchAddRelationDatas"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/relationDatas/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddRelationDatasResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量新增关系数据
//
// @param request - BatchAddRelationDatasRequest
//
// @return BatchAddRelationDatasResponse
func (client *Client) BatchAddRelationDatas(request *BatchAddRelationDatasRequest) (_result *BatchAddRelationDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchAddRelationDatasHeaders{}
	_result = &BatchAddRelationDatasResponse{}
	_body, _err := client.BatchAddRelationDatasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量创建线索数据
//
// @param request - BatchCreateClueDataRequest
//
// @param headers - BatchCreateClueDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateClueDataResponse
func (client *Client) BatchCreateClueDataWithOptions(request *BatchCreateClueDataRequest, headers *BatchCreateClueDataHeaders, runtime *util.RuntimeOptions) (_result *BatchCreateClueDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataList)) {
		body["dataList"] = request.DataList
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateSeas)) {
		body["privateSeas"] = request.PrivateSeas
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
		Action:      tea.String("BatchCreateClueData"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/clues/datas/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCreateClueDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量创建线索数据
//
// @param request - BatchCreateClueDataRequest
//
// @return BatchCreateClueDataResponse
func (client *Client) BatchCreateClueData(request *BatchCreateClueDataRequest) (_result *BatchCreateClueDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchCreateClueDataHeaders{}
	_result = &BatchCreateClueDataResponse{}
	_body, _err := client.BatchCreateClueDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除跟进记录
//
// @param request - BatchRemoveFollowRecordsRequest
//
// @param headers - BatchRemoveFollowRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchRemoveFollowRecordsResponse
func (client *Client) BatchRemoveFollowRecordsWithOptions(request *BatchRemoveFollowRecordsRequest, headers *BatchRemoveFollowRecordsHeaders, runtime *util.RuntimeOptions) (_result *BatchRemoveFollowRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		body["instanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

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
		Action:      tea.String("BatchRemoveFollowRecords"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/followRecords/batchRemove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchRemoveFollowRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除跟进记录
//
// @param request - BatchRemoveFollowRecordsRequest
//
// @return BatchRemoveFollowRecordsResponse
func (client *Client) BatchRemoveFollowRecords(request *BatchRemoveFollowRecordsRequest) (_result *BatchRemoveFollowRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchRemoveFollowRecordsHeaders{}
	_result = &BatchRemoveFollowRecordsResponse{}
	_body, _err := client.BatchRemoveFollowRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务窗消息群发
//
// @param request - BatchSendOfficialAccountOTOMessageRequest
//
// @param headers - BatchSendOfficialAccountOTOMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSendOfficialAccountOTOMessageResponse
func (client *Client) BatchSendOfficialAccountOTOMessageWithOptions(request *BatchSendOfficialAccountOTOMessageRequest, headers *BatchSendOfficialAccountOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *BatchSendOfficialAccountOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
	}

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
		Action:      tea.String("BatchSendOfficialAccountOTOMessage"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/officialAccounts/oToMessages/batchSend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSendOfficialAccountOTOMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务窗消息群发
//
// @param request - BatchSendOfficialAccountOTOMessageRequest
//
// @return BatchSendOfficialAccountOTOMessageResponse
func (client *Client) BatchSendOfficialAccountOTOMessage(request *BatchSendOfficialAccountOTOMessageRequest) (_result *BatchSendOfficialAccountOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchSendOfficialAccountOTOMessageHeaders{}
	_result = &BatchSendOfficialAccountOTOMessageResponse{}
	_body, _err := client.BatchSendOfficialAccountOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量修改联系人
//
// @param request - BatchUpdateContactsRequest
//
// @param headers - BatchUpdateContactsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateContactsResponse
func (client *Client) BatchUpdateContactsWithOptions(request *BatchUpdateContactsRequest, headers *BatchUpdateContactsHeaders, runtime *util.RuntimeOptions) (_result *BatchUpdateContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationList)) {
		body["relationList"] = request.RelationList
	}

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
		Action:      tea.String("BatchUpdateContacts"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/contacts/batch"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdateContactsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量修改联系人
//
// @param request - BatchUpdateContactsRequest
//
// @return BatchUpdateContactsResponse
func (client *Client) BatchUpdateContacts(request *BatchUpdateContactsRequest) (_result *BatchUpdateContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchUpdateContactsHeaders{}
	_result = &BatchUpdateContactsResponse{}
	_body, _err := client.BatchUpdateContactsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量修改跟进记录
//
// @param request - BatchUpdateFollowRecordsRequest
//
// @param headers - BatchUpdateFollowRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateFollowRecordsResponse
func (client *Client) BatchUpdateFollowRecordsWithOptions(request *BatchUpdateFollowRecordsRequest, headers *BatchUpdateFollowRecordsHeaders, runtime *util.RuntimeOptions) (_result *BatchUpdateFollowRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceList)) {
		body["instanceList"] = request.InstanceList
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

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
		Action:      tea.String("BatchUpdateFollowRecords"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/followRecords/batch"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdateFollowRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量修改跟进记录
//
// @param request - BatchUpdateFollowRecordsRequest
//
// @return BatchUpdateFollowRecordsResponse
func (client *Client) BatchUpdateFollowRecords(request *BatchUpdateFollowRecordsRequest) (_result *BatchUpdateFollowRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchUpdateFollowRecordsHeaders{}
	_result = &BatchUpdateFollowRecordsResponse{}
	_body, _err := client.BatchUpdateFollowRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量修改关系数据
//
// @param request - BatchUpdateRelationDatasRequest
//
// @param headers - BatchUpdateRelationDatasHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateRelationDatasResponse
func (client *Client) BatchUpdateRelationDatasWithOptions(request *BatchUpdateRelationDatasRequest, headers *BatchUpdateRelationDatasHeaders, runtime *util.RuntimeOptions) (_result *BatchUpdateRelationDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationList)) {
		body["relationList"] = request.RelationList
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.SkipDuplicateCheck)) {
		body["skipDuplicateCheck"] = request.SkipDuplicateCheck
	}

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
		Action:      tea.String("BatchUpdateRelationDatas"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/relationDatas/batch"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdateRelationDatasResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量修改关系数据
//
// @param request - BatchUpdateRelationDatasRequest
//
// @return BatchUpdateRelationDatasResponse
func (client *Client) BatchUpdateRelationDatas(request *BatchUpdateRelationDatasRequest) (_result *BatchUpdateRelationDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchUpdateRelationDatasHeaders{}
	_result = &BatchUpdateRelationDatasResponse{}
	_body, _err := client.BatchUpdateRelationDatasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 核销权益库存
//
// @param request - ConsumeBenefitInventoryRequest
//
// @param headers - ConsumeBenefitInventoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsumeBenefitInventoryResponse
func (client *Client) ConsumeBenefitInventoryWithOptions(request *ConsumeBenefitInventoryRequest, headers *ConsumeBenefitInventoryHeaders, runtime *util.RuntimeOptions) (_result *ConsumeBenefitInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitCode)) {
		body["benefitCode"] = request.BenefitCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizRequestId)) {
		body["bizRequestId"] = request.BizRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsumeQuota)) {
		body["consumeQuota"] = request.ConsumeQuota
	}

	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		body["optUserId"] = request.OptUserId
	}

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
		Action:      tea.String("ConsumeBenefitInventory"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/benefitInventories/consume"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConsumeBenefitInventoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 核销权益库存
//
// @param request - ConsumeBenefitInventoryRequest
//
// @return ConsumeBenefitInventoryResponse
func (client *Client) ConsumeBenefitInventory(request *ConsumeBenefitInventoryRequest) (_result *ConsumeBenefitInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConsumeBenefitInventoryHeaders{}
	_result = &ConsumeBenefitInventoryResponse{}
	_body, _err := client.ConsumeBenefitInventoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// CRM客户通讯录数据写入接口，支持客户&联系人数据合并写入
//
// @param request - CreateCustomerRequest
//
// @param headers - CreateCustomerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomerResponse
func (client *Client) CreateCustomerWithOptions(request *CreateCustomerRequest, headers *CreateCustomerHeaders, runtime *util.RuntimeOptions) (_result *CreateCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Contacts)) {
		body["contacts"] = request.Contacts
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorUserId)) {
		body["creatorUserId"] = request.CreatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["objectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.Permission)) {
		body["permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(request.SaveOption)) {
		body["saveOption"] = request.SaveOption
	}

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
		Action:      tea.String("CreateCustomer"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// CRM客户通讯录数据写入接口，支持客户&联系人数据合并写入
//
// @param request - CreateCustomerRequest
//
// @return CreateCustomerResponse
func (client *Client) CreateCustomer(request *CreateCustomerRequest) (_result *CreateCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCustomerHeaders{}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CreateCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建客户群
//
// @param request - CreateGroupRequest
//
// @param headers - CreateGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, headers *CreateGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUserIds)) {
		body["memberUserIds"] = request.MemberUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		body["ownerUserId"] = request.OwnerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

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
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建客户群
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupHeaders{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建群组
//
// @param request - CreateGroupSetRequest
//
// @param headers - CreateGroupSetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupSetResponse
func (client *Client) CreateGroupSetWithOptions(request *CreateGroupSetRequest, headers *CreateGroupSetHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorUserId)) {
		body["creatorUserId"] = request.CreatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerUserIds)) {
		body["managerUserIds"] = request.ManagerUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.MemberQuota)) {
		body["memberQuota"] = request.MemberQuota
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Notice)) {
		body["notice"] = request.Notice
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeToped)) {
		body["noticeToped"] = request.NoticeToped
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		body["ownerUserId"] = request.OwnerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Welcome)) {
		body["welcome"] = request.Welcome
	}

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
		Action:      tea.String("CreateGroupSet"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/groupSets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建群组
//
// @param request - CreateGroupSetRequest
//
// @return CreateGroupSetResponse
func (client *Client) CreateGroupSet(request *CreateGroupSetRequest) (_result *CreateGroupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupSetHeaders{}
	_result = &CreateGroupSetResponse{}
	_body, _err := client.CreateGroupSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建关系模型表结构
//
// @param request - CreateRelationMetaRequest
//
// @param headers - CreateRelationMetaHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRelationMetaResponse
func (client *Client) CreateRelationMetaWithOptions(request *CreateRelationMetaRequest, headers *CreateRelationMetaHeaders, runtime *util.RuntimeOptions) (_result *CreateRelationMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationMetaDTO)) {
		body["relationMetaDTO"] = request.RelationMetaDTO
	}

	if !tea.BoolValue(util.IsUnset(request.Tenant)) {
		body["tenant"] = request.Tenant
	}

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
		Action:      tea.String("CreateRelationMeta"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/relations/metas/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRelationMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建关系模型表结构
//
// @param request - CreateRelationMetaRequest
//
// @return CreateRelationMetaResponse
func (client *Client) CreateRelationMeta(request *CreateRelationMetaRequest) (_result *CreateRelationMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateRelationMetaHeaders{}
	_result = &CreateRelationMetaResponse{}
	_body, _err := client.CreateRelationMetaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除CRM自定义对象数据
//
// @param request - DeleteCrmCustomObjectDataRequest
//
// @param headers - DeleteCrmCustomObjectDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCrmCustomObjectDataResponse
func (client *Client) DeleteCrmCustomObjectDataWithOptions(instanceId *string, request *DeleteCrmCustomObjectDataRequest, headers *DeleteCrmCustomObjectDataHeaders, runtime *util.RuntimeOptions) (_result *DeleteCrmCustomObjectDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
	}

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
		Action:      tea.String("DeleteCrmCustomObjectData"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customObjectDatas/instances/" + tea.StringValue(instanceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCrmCustomObjectDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除CRM自定义对象数据
//
// @param request - DeleteCrmCustomObjectDataRequest
//
// @return DeleteCrmCustomObjectDataResponse
func (client *Client) DeleteCrmCustomObjectData(instanceId *string, request *DeleteCrmCustomObjectDataRequest) (_result *DeleteCrmCustomObjectDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCrmCustomObjectDataHeaders{}
	_result = &DeleteCrmCustomObjectDataResponse{}
	_body, _err := client.DeleteCrmCustomObjectDataWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// crm自定义表单数据删除接口
//
// @param request - DeleteCrmFormInstanceRequest
//
// @param headers - DeleteCrmFormInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCrmFormInstanceResponse
func (client *Client) DeleteCrmFormInstanceWithOptions(instanceId *string, request *DeleteCrmFormInstanceRequest, headers *DeleteCrmFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *DeleteCrmFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentOperatorUserId)) {
		query["currentOperatorUserId"] = request.CurrentOperatorUserId
	}

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
		Action:      tea.String("DeleteCrmFormInstance"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/formInstances/" + tea.StringValue(instanceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCrmFormInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// crm自定义表单数据删除接口
//
// @param request - DeleteCrmFormInstanceRequest
//
// @return DeleteCrmFormInstanceResponse
func (client *Client) DeleteCrmFormInstance(instanceId *string, request *DeleteCrmFormInstanceRequest) (_result *DeleteCrmFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCrmFormInstanceHeaders{}
	_result = &DeleteCrmFormInstanceResponse{}
	_body, _err := client.DeleteCrmFormInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除crm个人客户（或企业客户）
//
// @param request - DeleteCrmPersonalCustomerRequest
//
// @param headers - DeleteCrmPersonalCustomerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCrmPersonalCustomerResponse
func (client *Client) DeleteCrmPersonalCustomerWithOptions(dataId *string, request *DeleteCrmPersonalCustomerRequest, headers *DeleteCrmPersonalCustomerHeaders, runtime *util.RuntimeOptions) (_result *DeleteCrmPersonalCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentOperatorUserId)) {
		query["currentOperatorUserId"] = request.CurrentOperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		query["relationType"] = request.RelationType
	}

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
		Action:      tea.String("DeleteCrmPersonalCustomer"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/personalCustomers/" + tea.StringValue(dataId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCrmPersonalCustomerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除crm个人客户（或企业客户）
//
// @param request - DeleteCrmPersonalCustomerRequest
//
// @return DeleteCrmPersonalCustomerResponse
func (client *Client) DeleteCrmPersonalCustomer(dataId *string, request *DeleteCrmPersonalCustomerRequest) (_result *DeleteCrmPersonalCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCrmPersonalCustomerHeaders{}
	_result = &DeleteCrmPersonalCustomerResponse{}
	_body, _err := client.DeleteCrmPersonalCustomerWithOptions(dataId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除线索
//
// @param request - DeleteLeadsRequest
//
// @param headers - DeleteLeadsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLeadsResponse
func (client *Client) DeleteLeadsWithOptions(request *DeleteLeadsRequest, headers *DeleteLeadsHeaders, runtime *util.RuntimeOptions) (_result *DeleteLeadsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutLeadsIds)) {
		body["outLeadsIds"] = request.OutLeadsIds
	}

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
		Action:      tea.String("DeleteLeads"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/leads/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLeadsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除线索
//
// @param request - DeleteLeadsRequest
//
// @return DeleteLeadsResponse
func (client *Client) DeleteLeads(request *DeleteLeadsRequest) (_result *DeleteLeadsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteLeadsHeaders{}
	_result = &DeleteLeadsResponse{}
	_body, _err := client.DeleteLeadsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关系模型表结构删除字段
//
// @param request - DeleteRelationMetaFieldRequest
//
// @param headers - DeleteRelationMetaFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRelationMetaFieldResponse
func (client *Client) DeleteRelationMetaFieldWithOptions(request *DeleteRelationMetaFieldRequest, headers *DeleteRelationMetaFieldHeaders, runtime *util.RuntimeOptions) (_result *DeleteRelationMetaFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FieldIdList)) {
		body["fieldIdList"] = request.FieldIdList
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.Tenant)) {
		body["tenant"] = request.Tenant
	}

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
		Action:      tea.String("DeleteRelationMetaField"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/relations/metas/fields/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRelationMetaFieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关系模型表结构删除字段
//
// @param request - DeleteRelationMetaFieldRequest
//
// @return DeleteRelationMetaFieldResponse
func (client *Client) DeleteRelationMetaField(request *DeleteRelationMetaFieldRequest) (_result *DeleteRelationMetaFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRelationMetaFieldHeaders{}
	_result = &DeleteRelationMetaFieldResponse{}
	_body, _err := client.DeleteRelationMetaFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取CRM客户对象的元数据描述
//
// @param request - DescribeCrmPersonalCustomerObjectMetaRequest
//
// @param headers - DescribeCrmPersonalCustomerObjectMetaHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCrmPersonalCustomerObjectMetaResponse
func (client *Client) DescribeCrmPersonalCustomerObjectMetaWithOptions(request *DescribeCrmPersonalCustomerObjectMetaRequest, headers *DescribeCrmPersonalCustomerObjectMetaHeaders, runtime *util.RuntimeOptions) (_result *DescribeCrmPersonalCustomerObjectMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		query["relationType"] = request.RelationType
	}

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
		Action:      tea.String("DescribeCrmPersonalCustomerObjectMeta"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/personalCustomers/objectMeta"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCrmPersonalCustomerObjectMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取CRM客户对象的元数据描述
//
// @param request - DescribeCrmPersonalCustomerObjectMetaRequest
//
// @return DescribeCrmPersonalCustomerObjectMetaResponse
func (client *Client) DescribeCrmPersonalCustomerObjectMeta(request *DescribeCrmPersonalCustomerObjectMetaRequest) (_result *DescribeCrmPersonalCustomerObjectMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DescribeCrmPersonalCustomerObjectMetaHeaders{}
	_result = &DescribeCrmPersonalCustomerObjectMetaResponse{}
	_body, _err := client.DescribeCrmPersonalCustomerObjectMetaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询模型表结构
//
// @param request - DescribeMetaModelRequest
//
// @param headers - DescribeMetaModelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetaModelResponse
func (client *Client) DescribeMetaModelWithOptions(request *DescribeMetaModelRequest, headers *DescribeMetaModelHeaders, runtime *util.RuntimeOptions) (_result *DescribeMetaModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		body["bizTypes"] = request.BizTypes
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Tenant)) {
		body["tenant"] = request.Tenant
	}

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
		Action:      tea.String("DescribeMetaModel"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/metas/models/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMetaModelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型表结构
//
// @param request - DescribeMetaModelRequest
//
// @return DescribeMetaModelResponse
func (client *Client) DescribeMetaModel(request *DescribeMetaModelRequest) (_result *DescribeMetaModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DescribeMetaModelHeaders{}
	_result = &DescribeMetaModelResponse{}
	_body, _err := client.DescribeMetaModelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询关系模型表结构
//
// @param request - DescribeRelationMetaRequest
//
// @param headers - DescribeRelationMetaHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRelationMetaResponse
func (client *Client) DescribeRelationMetaWithOptions(request *DescribeRelationMetaRequest, headers *DescribeRelationMetaHeaders, runtime *util.RuntimeOptions) (_result *DescribeRelationMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationTypes)) {
		body["relationTypes"] = request.RelationTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Tenant)) {
		body["tenant"] = request.Tenant
	}

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
		Action:      tea.String("DescribeRelationMeta"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/relations/metas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRelationMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询关系模型表结构
//
// @param request - DescribeRelationMetaRequest
//
// @return DescribeRelationMetaResponse
func (client *Client) DescribeRelationMeta(request *DescribeRelationMetaRequest) (_result *DescribeRelationMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DescribeRelationMetaHeaders{}
	_result = &DescribeRelationMetaResponse{}
	_body, _err := client.DescribeRelationMetaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取关联对象的跟进记录列表
//
// @param request - FindTargetRelatedFollowRecordsRequest
//
// @param headers - FindTargetRelatedFollowRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindTargetRelatedFollowRecordsResponse
func (client *Client) FindTargetRelatedFollowRecordsWithOptions(request *FindTargetRelatedFollowRecordsRequest, headers *FindTargetRelatedFollowRecordsHeaders, runtime *util.RuntimeOptions) (_result *FindTargetRelatedFollowRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FollowTargetDataId)) {
		body["followTargetDataId"] = request.FollowTargetDataId
	}

	if !tea.BoolValue(util.IsUnset(request.FollowTargetType)) {
		body["followTargetType"] = request.FollowTargetType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

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
		Action:      tea.String("FindTargetRelatedFollowRecords"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/targetFollowRecords/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FindTargetRelatedFollowRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取关联对象的跟进记录列表
//
// @param request - FindTargetRelatedFollowRecordsRequest
//
// @return FindTargetRelatedFollowRecordsResponse
func (client *Client) FindTargetRelatedFollowRecords(request *FindTargetRelatedFollowRecordsRequest) (_result *FindTargetRelatedFollowRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FindTargetRelatedFollowRecordsHeaders{}
	_result = &FindTargetRelatedFollowRecordsResponse{}
	_body, _err := client.FindTargetRelatedFollowRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取所有客户的掉保时间数据
//
// @param request - GetAllCustomerRecyclesRequest
//
// @param headers - GetAllCustomerRecyclesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAllCustomerRecyclesResponse
func (client *Client) GetAllCustomerRecyclesWithOptions(request *GetAllCustomerRecyclesRequest, headers *GetAllCustomerRecyclesHeaders, runtime *util.RuntimeOptions) (_result *GetAllCustomerRecyclesResponse, _err error) {
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
		Action:      tea.String("GetAllCustomerRecycles"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customerRecycles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllCustomerRecyclesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取所有客户的掉保时间数据
//
// @param request - GetAllCustomerRecyclesRequest
//
// @return GetAllCustomerRecyclesResponse
func (client *Client) GetAllCustomerRecycles(request *GetAllCustomerRecyclesRequest) (_result *GetAllCustomerRecyclesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAllCustomerRecyclesHeaders{}
	_result = &GetAllCustomerRecyclesResponse{}
	_body, _err := client.GetAllCustomerRecyclesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据指定条件查询联系人数据
//
// @param request - GetContactsRequest
//
// @param headers - GetContactsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContactsResponse
func (client *Client) GetContactsWithOptions(request *GetContactsRequest, headers *GetContactsHeaders, runtime *util.RuntimeOptions) (_result *GetContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentOperatorUserId)) {
		body["currentOperatorUserId"] = request.CurrentOperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["objectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.ProviderCorpId)) {
		body["providerCorpId"] = request.ProviderCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDsl)) {
		body["queryDsl"] = request.QueryDsl
	}

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
		Action:      tea.String("GetContacts"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customObjects/contacts/datas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetContactsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据指定条件查询联系人数据
//
// @param request - GetContactsRequest
//
// @return GetContactsResponse
func (client *Client) GetContacts(request *GetContactsRequest) (_result *GetContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetContactsHeaders{}
	_result = &GetContactsResponse{}
	_body, _err := client.GetContactsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个客户群
//
// @param headers - GetCrmGroupChatHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCrmGroupChatResponse
func (client *Client) GetCrmGroupChatWithOptions(openConversationId *string, headers *GetCrmGroupChatHeaders, runtime *util.RuntimeOptions) (_result *GetCrmGroupChatResponse, _err error) {
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
		Action:      tea.String("GetCrmGroupChat"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/crmGroupChats/" + tea.StringValue(openConversationId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCrmGroupChatResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个客户群
//
// @return GetCrmGroupChatResponse
func (client *Client) GetCrmGroupChat(openConversationId *string) (_result *GetCrmGroupChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCrmGroupChatHeaders{}
	_result = &GetCrmGroupChatResponse{}
	_body, _err := client.GetCrmGroupChatWithOptions(openConversationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取多个客户群
//
// @param request - GetCrmGroupChatMultiRequest
//
// @param headers - GetCrmGroupChatMultiHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCrmGroupChatMultiResponse
func (client *Client) GetCrmGroupChatMultiWithOptions(request *GetCrmGroupChatMultiRequest, headers *GetCrmGroupChatMultiHeaders, runtime *util.RuntimeOptions) (_result *GetCrmGroupChatMultiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationIds)) {
		body["openConversationIds"] = request.OpenConversationIds
	}

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
		Action:      tea.String("GetCrmGroupChatMulti"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/crmGroupChats/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCrmGroupChatMultiResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取多个客户群
//
// @param request - GetCrmGroupChatMultiRequest
//
// @return GetCrmGroupChatMultiResponse
func (client *Client) GetCrmGroupChatMulti(request *GetCrmGroupChatMultiRequest) (_result *GetCrmGroupChatMultiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCrmGroupChatMultiHeaders{}
	_result = &GetCrmGroupChatMultiResponse{}
	_body, _err := client.GetCrmGroupChatMultiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个客户群
//
// @param request - GetCrmGroupChatSingleRequest
//
// @param headers - GetCrmGroupChatSingleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCrmGroupChatSingleResponse
func (client *Client) GetCrmGroupChatSingleWithOptions(request *GetCrmGroupChatSingleRequest, headers *GetCrmGroupChatSingleHeaders, runtime *util.RuntimeOptions) (_result *GetCrmGroupChatSingleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
	}

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
		Action:      tea.String("GetCrmGroupChatSingle"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/crmGroupChats/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCrmGroupChatSingleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个客户群
//
// @param request - GetCrmGroupChatSingleRequest
//
// @return GetCrmGroupChatSingleResponse
func (client *Client) GetCrmGroupChatSingle(request *GetCrmGroupChatSingleRequest) (_result *GetCrmGroupChatSingleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCrmGroupChatSingleHeaders{}
	_result = &GetCrmGroupChatSingleResponse{}
	_body, _err := client.GetCrmGroupChatSingleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取CRM表单权限配置
//
// @param request - GetCrmRolePermissionRequest
//
// @param headers - GetCrmRolePermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCrmRolePermissionResponse
func (client *Client) GetCrmRolePermissionWithOptions(request *GetCrmRolePermissionRequest, headers *GetCrmRolePermissionHeaders, runtime *util.RuntimeOptions) (_result *GetCrmRolePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["resourceId"] = request.ResourceId
	}

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
		Action:      tea.String("GetCrmRolePermission"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/permissions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCrmRolePermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取CRM表单权限配置
//
// @param request - GetCrmRolePermissionRequest
//
// @return GetCrmRolePermissionResponse
func (client *Client) GetCrmRolePermission(request *GetCrmRolePermissionRequest) (_result *GetCrmRolePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCrmRolePermissionHeaders{}
	_result = &GetCrmRolePermissionResponse{}
	_body, _err := client.GetCrmRolePermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取某个客户的客户动态
//
// @param request - GetCustomerTracksByRelationIdRequest
//
// @param headers - GetCustomerTracksByRelationIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerTracksByRelationIdResponse
func (client *Client) GetCustomerTracksByRelationIdWithOptions(request *GetCustomerTracksByRelationIdRequest, headers *GetCustomerTracksByRelationIdHeaders, runtime *util.RuntimeOptions) (_result *GetCustomerTracksByRelationIdResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RelationId)) {
		query["relationId"] = request.RelationId
	}

	if !tea.BoolValue(util.IsUnset(request.TypeGroup)) {
		query["typeGroup"] = request.TypeGroup
	}

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
		Action:      tea.String("GetCustomerTracksByRelationId"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customerTracks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomerTracksByRelationIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取某个客户的客户动态
//
// @param request - GetCustomerTracksByRelationIdRequest
//
// @return GetCustomerTracksByRelationIdResponse
func (client *Client) GetCustomerTracksByRelationId(request *GetCustomerTracksByRelationIdRequest) (_result *GetCustomerTracksByRelationIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCustomerTracksByRelationIdHeaders{}
	_result = &GetCustomerTracksByRelationIdResponse{}
	_body, _err := client.GetCustomerTracksByRelationIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群组
//
// @param request - GetGroupSetRequest
//
// @param headers - GetGroupSetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupSetResponse
func (client *Client) GetGroupSetWithOptions(request *GetGroupSetRequest, headers *GetGroupSetHeaders, runtime *util.RuntimeOptions) (_result *GetGroupSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		query["openGroupSetId"] = request.OpenGroupSetId
	}

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
		Action:      tea.String("GetGroupSet"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/groupSets"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群组
//
// @param request - GetGroupSetRequest
//
// @return GetGroupSetResponse
func (client *Client) GetGroupSet(request *GetGroupSetRequest) (_result *GetGroupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGroupSetHeaders{}
	_result = &GetGroupSetResponse{}
	_body, _err := client.GetGroupSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取内购商品信息
//
// @param request - GetInAppPurchaseGoodsRequest
//
// @param headers - GetInAppPurchaseGoodsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInAppPurchaseGoodsResponse
func (client *Client) GetInAppPurchaseGoodsWithOptions(request *GetInAppPurchaseGoodsRequest, headers *GetInAppPurchaseGoodsHeaders, runtime *util.RuntimeOptions) (_result *GetInAppPurchaseGoodsResponse, _err error) {
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
		Action:      tea.String("GetInAppPurchaseGoods"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/inAppPurchaseGoods/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInAppPurchaseGoodsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取内购商品信息
//
// @param request - GetInAppPurchaseGoodsRequest
//
// @return GetInAppPurchaseGoodsResponse
func (client *Client) GetInAppPurchaseGoods(request *GetInAppPurchaseGoodsRequest) (_result *GetInAppPurchaseGoodsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInAppPurchaseGoodsHeaders{}
	_result = &GetInAppPurchaseGoodsResponse{}
	_body, _err := client.GetInAppPurchaseGoodsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取自定义导航挂靠节点结构
//
// @param request - GetNavigationCatalogRequest
//
// @param headers - GetNavigationCatalogHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNavigationCatalogResponse
func (client *Client) GetNavigationCatalogWithOptions(request *GetNavigationCatalogRequest, headers *GetNavigationCatalogHeaders, runtime *util.RuntimeOptions) (_result *GetNavigationCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTraceId)) {
		query["bizTraceId"] = request.BizTraceId
	}

	if !tea.BoolValue(util.IsUnset(request.Module)) {
		query["module"] = request.Module
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		query["operatorUserId"] = request.OperatorUserId
	}

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
		Action:      tea.String("GetNavigationCatalog"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/navigations/catalogs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNavigationCatalogResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义导航挂靠节点结构
//
// @param request - GetNavigationCatalogRequest
//
// @return GetNavigationCatalogResponse
func (client *Client) GetNavigationCatalog(request *GetNavigationCatalogRequest) (_result *GetNavigationCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetNavigationCatalogHeaders{}
	_result = &GetNavigationCatalogResponse{}
	_body, _err := client.GetNavigationCatalogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据指定条件查询自定义对象数据
//
// @param request - GetObjectDataRequest
//
// @param headers - GetObjectDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetObjectDataResponse
func (client *Client) GetObjectDataWithOptions(request *GetObjectDataRequest, headers *GetObjectDataHeaders, runtime *util.RuntimeOptions) (_result *GetObjectDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentOperatorUserId)) {
		body["currentOperatorUserId"] = request.CurrentOperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDsl)) {
		body["queryDsl"] = request.QueryDsl
	}

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
		Action:      tea.String("GetObjectData"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customObjects/datas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetObjectDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据指定条件查询自定义对象数据
//
// @param request - GetObjectDataRequest
//
// @return GetObjectDataResponse
func (client *Client) GetObjectData(request *GetObjectDataRequest) (_result *GetObjectDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetObjectDataHeaders{}
	_result = &GetObjectDataResponse{}
	_body, _err := client.GetObjectDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取关注服务窗的联系人信息，包括手机号、主企业等字段，调用前先进行用户授权
//
// @param headers - GetOfficialAccountContactInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOfficialAccountContactInfoResponse
func (client *Client) GetOfficialAccountContactInfoWithOptions(userId *string, headers *GetOfficialAccountContactInfoHeaders, runtime *util.RuntimeOptions) (_result *GetOfficialAccountContactInfoResponse, _err error) {
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
		Action:      tea.String("GetOfficialAccountContactInfo"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/officialAccounts/contacts/" + tea.StringValue(userId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOfficialAccountContactInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取关注服务窗的联系人信息，包括手机号、主企业等字段，调用前先进行用户授权
//
// @return GetOfficialAccountContactInfoResponse
func (client *Client) GetOfficialAccountContactInfo(userId *string) (_result *GetOfficialAccountContactInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOfficialAccountContactInfoHeaders{}
	_result = &GetOfficialAccountContactInfoResponse{}
	_body, _err := client.GetOfficialAccountContactInfoWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取服务窗联系人信息
//
// @param request - GetOfficialAccountContactsRequest
//
// @param headers - GetOfficialAccountContactsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOfficialAccountContactsResponse
func (client *Client) GetOfficialAccountContactsWithOptions(request *GetOfficialAccountContactsRequest, headers *GetOfficialAccountContactsHeaders, runtime *util.RuntimeOptions) (_result *GetOfficialAccountContactsResponse, _err error) {
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
		Action:      tea.String("GetOfficialAccountContacts"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/officialAccounts/contacts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOfficialAccountContactsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取服务窗联系人信息
//
// @param request - GetOfficialAccountContactsRequest
//
// @return GetOfficialAccountContactsResponse
func (client *Client) GetOfficialAccountContacts(request *GetOfficialAccountContactsRequest) (_result *GetOfficialAccountContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOfficialAccountContactsHeaders{}
	_result = &GetOfficialAccountContactsResponse{}
	_body, _err := client.GetOfficialAccountContactsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取服务窗消息发送的结果
//
// @param request - GetOfficialAccountOTOMessageResultRequest
//
// @param headers - GetOfficialAccountOTOMessageResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOfficialAccountOTOMessageResultResponse
func (client *Client) GetOfficialAccountOTOMessageResultWithOptions(request *GetOfficialAccountOTOMessageResultRequest, headers *GetOfficialAccountOTOMessageResultHeaders, runtime *util.RuntimeOptions) (_result *GetOfficialAccountOTOMessageResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPushId)) {
		query["openPushId"] = request.OpenPushId
	}

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
		Action:      tea.String("GetOfficialAccountOTOMessageResult"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/officialAccounts/oToMessages/sendResults"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOfficialAccountOTOMessageResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务窗消息发送的结果
//
// @param request - GetOfficialAccountOTOMessageResultRequest
//
// @return GetOfficialAccountOTOMessageResultResponse
func (client *Client) GetOfficialAccountOTOMessageResult(request *GetOfficialAccountOTOMessageResultRequest) (_result *GetOfficialAccountOTOMessageResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOfficialAccountOTOMessageResultHeaders{}
	_result = &GetOfficialAccountOTOMessageResultResponse{}
	_body, _err := client.GetOfficialAccountOTOMessageResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取某个和oa关联的表单的具体数据
//
// @param request - GetRelatedViewTabDataRequest
//
// @param headers - GetRelatedViewTabDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRelatedViewTabDataResponse
func (client *Client) GetRelatedViewTabDataWithOptions(request *GetRelatedViewTabDataRequest, headers *GetRelatedViewTabDataHeaders, runtime *util.RuntimeOptions) (_result *GetRelatedViewTabDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		body["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedField)) {
		body["relatedField"] = request.RelatedField
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedInstId)) {
		body["relatedInstId"] = request.RelatedInstId
	}

	if !tea.BoolValue(util.IsUnset(request.ViewUserId)) {
		body["viewUserId"] = request.ViewUserId
	}

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
		Action:      tea.String("GetRelatedViewTabData"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/formRelatedTabs/datas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRelatedViewTabDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某个和oa关联的表单的具体数据
//
// @param request - GetRelatedViewTabDataRequest
//
// @return GetRelatedViewTabDataResponse
func (client *Client) GetRelatedViewTabData(request *GetRelatedViewTabDataRequest) (_result *GetRelatedViewTabDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRelatedViewTabDataHeaders{}
	_result = &GetRelatedViewTabDataResponse{}
	_body, _err := client.GetRelatedViewTabDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取和oa关联的表单tab信息
//
// @param request - GetRelatedViewTabMetaRequest
//
// @param headers - GetRelatedViewTabMetaHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRelatedViewTabMetaResponse
func (client *Client) GetRelatedViewTabMetaWithOptions(request *GetRelatedViewTabMetaRequest, headers *GetRelatedViewTabMetaHeaders, runtime *util.RuntimeOptions) (_result *GetRelatedViewTabMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		body["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.ViewUserId)) {
		body["viewUserId"] = request.ViewUserId
	}

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
		Action:      tea.String("GetRelatedViewTabMeta"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/formRelatedTabs/meta/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRelatedViewTabMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取和oa关联的表单tab信息
//
// @param request - GetRelatedViewTabMetaRequest
//
// @return GetRelatedViewTabMetaResponse
func (client *Client) GetRelatedViewTabMeta(request *GetRelatedViewTabMetaRequest) (_result *GetRelatedViewTabMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRelatedViewTabMetaHeaders{}
	_result = &GetRelatedViewTabMetaResponse{}
	_body, _err := client.GetRelatedViewTabMetaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取关系数据查重规则
//
// @param request - GetRelationUkSettingRequest
//
// @param headers - GetRelationUkSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRelationUkSettingResponse
func (client *Client) GetRelationUkSettingWithOptions(request *GetRelationUkSettingRequest, headers *GetRelationUkSettingHeaders, runtime *util.RuntimeOptions) (_result *GetRelationUkSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		query["relationType"] = request.RelationType
	}

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
		Action:      tea.String("GetRelationUkSetting"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/relationUkSettings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRelationUkSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取关系数据查重规则
//
// @param request - GetRelationUkSettingRequest
//
// @return GetRelationUkSettingResponse
func (client *Client) GetRelationUkSetting(request *GetRelationUkSettingRequest) (_result *GetRelationUkSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRelationUkSettingHeaders{}
	_result = &GetRelationUkSettingResponse{}
	_body, _err := client.GetRelationUkSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 加入群组
//
// @param request - JoinGroupSetRequest
//
// @param headers - JoinGroupSetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinGroupSetResponse
func (client *Client) JoinGroupSetWithOptions(request *JoinGroupSetRequest, headers *JoinGroupSetHeaders, runtime *util.RuntimeOptions) (_result *JoinGroupSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizDataList)) {
		body["bizDataList"] = request.BizDataList
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
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
		Action:      tea.String("JoinGroupSet"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/groupSets/join"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinGroupSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 加入群组
//
// @param request - JoinGroupSetRequest
//
// @return JoinGroupSetResponse
func (client *Client) JoinGroupSet(request *JoinGroupSetRequest) (_result *JoinGroupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &JoinGroupSetHeaders{}
	_result = &JoinGroupSetResponse{}
	_body, _err := client.JoinGroupSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询可用权益
//
// @param request - ListAvailableBenefitRequest
//
// @param headers - ListAvailableBenefitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableBenefitResponse
func (client *Client) ListAvailableBenefitWithOptions(request *ListAvailableBenefitRequest, headers *ListAvailableBenefitHeaders, runtime *util.RuntimeOptions) (_result *ListAvailableBenefitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitCodeList)) {
		body["benefitCodeList"] = request.BenefitCodeList
	}

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
		Action:      tea.String("ListAvailableBenefit"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/benefits/lists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableBenefitResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询可用权益
//
// @param request - ListAvailableBenefitRequest
//
// @return ListAvailableBenefitResponse
func (client *Client) ListAvailableBenefit(request *ListAvailableBenefitRequest) (_result *ListAvailableBenefitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAvailableBenefitHeaders{}
	_result = &ListAvailableBenefitResponse{}
	_body, _err := client.ListAvailableBenefitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询license
//
// @param request - ListBenefitLicenseRequest
//
// @param headers - ListBenefitLicenseHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBenefitLicenseResponse
func (client *Client) ListBenefitLicenseWithOptions(request *ListBenefitLicenseRequest, headers *ListBenefitLicenseHeaders, runtime *util.RuntimeOptions) (_result *ListBenefitLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domains)) {
		body["domains"] = request.Domains
	}

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
		Action:      tea.String("ListBenefitLicense"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/benefitLicenses/lists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBenefitLicenseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询license
//
// @param request - ListBenefitLicenseRequest
//
// @return ListBenefitLicenseResponse
func (client *Client) ListBenefitLicense(request *ListBenefitLicenseRequest) (_result *ListBenefitLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListBenefitLicenseHeaders{}
	_result = &ListBenefitLicenseResponse{}
	_body, _err := client.ListBenefitLicenseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取线索标签列表
//
// @param headers - ListClueTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClueTagResponse
func (client *Client) ListClueTagWithOptions(headers *ListClueTagHeaders, runtime *util.RuntimeOptions) (_result *ListClueTagResponse, _err error) {
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
		Action:      tea.String("ListClueTag"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/clues/tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClueTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取线索标签列表
//
// @return ListClueTagResponse
func (client *Client) ListClueTag() (_result *ListClueTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListClueTagHeaders{}
	_result = &ListClueTagResponse{}
	_body, _err := client.ListClueTagWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取crm个人客户
//
// @param request - ListCrmPersonalCustomersRequest
//
// @param headers - ListCrmPersonalCustomersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrmPersonalCustomersResponse
func (client *Client) ListCrmPersonalCustomersWithOptions(request *ListCrmPersonalCustomersRequest, headers *ListCrmPersonalCustomersHeaders, runtime *util.RuntimeOptions) (_result *ListCrmPersonalCustomersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentOperatorUserId)) {
		query["currentOperatorUserId"] = request.CurrentOperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		query["relationType"] = request.RelationType
	}

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
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("ListCrmPersonalCustomers"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/personalCustomers/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCrmPersonalCustomersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取crm个人客户
//
// @param request - ListCrmPersonalCustomersRequest
//
// @return ListCrmPersonalCustomersResponse
func (client *Client) ListCrmPersonalCustomers(request *ListCrmPersonalCustomersRequest) (_result *ListCrmPersonalCustomersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCrmPersonalCustomersHeaders{}
	_result = &ListCrmPersonalCustomersResponse{}
	_body, _err := client.ListCrmPersonalCustomersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群组列表
//
// @param request - ListGroupSetRequest
//
// @param headers - ListGroupSetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupSetResponse
func (client *Client) ListGroupSetWithOptions(request *ListGroupSetRequest, headers *ListGroupSetHeaders, runtime *util.RuntimeOptions) (_result *ListGroupSetResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.QueryDsl)) {
		query["queryDsl"] = request.QueryDsl
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		query["relationType"] = request.RelationType
	}

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
		Action:      tea.String("ListGroupSet"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/groupSets/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群组列表
//
// @param request - ListGroupSetRequest
//
// @return ListGroupSetResponse
func (client *Client) ListGroupSet(request *ListGroupSetRequest) (_result *ListGroupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListGroupSetHeaders{}
	_result = &ListGroupSetResponse{}
	_body, _err := client.ListGroupSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 覆盖更新客户数据权限
//
// @param request - OverrideUpdateCustomerDataAuthRequest
//
// @param headers - OverrideUpdateCustomerDataAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OverrideUpdateCustomerDataAuthResponse
func (client *Client) OverrideUpdateCustomerDataAuthWithOptions(request *OverrideUpdateCustomerDataAuthRequest, headers *OverrideUpdateCustomerDataAuthHeaders, runtime *util.RuntimeOptions) (_result *OverrideUpdateCustomerDataAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerIds)) {
		body["customerIds"] = request.CustomerIds
	}

	if !tea.BoolValue(util.IsUnset(request.DataAuthUserIds)) {
		body["dataAuthUserIds"] = request.DataAuthUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		body["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["roleType"] = request.RoleType
	}

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
		Action:      tea.String("OverrideUpdateCustomerDataAuth"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customers/dataAuth/overrideUpdate"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OverrideUpdateCustomerDataAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 覆盖更新客户数据权限
//
// @param request - OverrideUpdateCustomerDataAuthRequest
//
// @return OverrideUpdateCustomerDataAuthResponse
func (client *Client) OverrideUpdateCustomerDataAuth(request *OverrideUpdateCustomerDataAuthRequest) (_result *OverrideUpdateCustomerDataAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OverrideUpdateCustomerDataAuthHeaders{}
	_result = &OverrideUpdateCustomerDataAuthResponse{}
	_body, _err := client.OverrideUpdateCustomerDataAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取全量客户数据，根据不同的类型可以获取私海个人客户、企业客户，以及公海个人客户、企业客户，最多一次可获取100条数据
//
// @param request - QueryAllCustomerRequest
//
// @param headers - QueryAllCustomerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAllCustomerResponse
func (client *Client) QueryAllCustomerWithOptions(request *QueryAllCustomerRequest, headers *QueryAllCustomerHeaders, runtime *util.RuntimeOptions) (_result *QueryAllCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["objectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

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
		Action:      tea.String("QueryAllCustomer"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customerInstances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllCustomerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取全量客户数据，根据不同的类型可以获取私海个人客户、企业客户，以及公海个人客户、企业客户，最多一次可获取100条数据
//
// @param request - QueryAllCustomerRequest
//
// @return QueryAllCustomerResponse
func (client *Client) QueryAllCustomer(request *QueryAllCustomerRequest) (_result *QueryAllCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllCustomerHeaders{}
	_result = &QueryAllCustomerResponse{}
	_body, _err := client.QueryAllCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询企业客户动态
//
// @param request - QueryAllTracksRequest
//
// @param headers - QueryAllTracksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAllTracksResponse
func (client *Client) QueryAllTracksWithOptions(request *QueryAllTracksRequest, headers *QueryAllTracksHeaders, runtime *util.RuntimeOptions) (_result *QueryAllTracksResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["order"] = request.Order
	}

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
		Action:      tea.String("QueryAllTracks"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/customers/tracks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllTracksResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询企业客户动态
//
// @param request - QueryAllTracksRequest
//
// @return QueryAllTracksResponse
func (client *Client) QueryAllTracks(request *QueryAllTracksRequest) (_result *QueryAllTracksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllTracksHeaders{}
	_result = &QueryAllTracksResponse{}
	_body, _err := client.QueryAllTracksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询客户管理应用管理员
//
// @param request - QueryAppManagerRequest
//
// @param headers - QueryAppManagerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAppManagerResponse
func (client *Client) QueryAppManagerWithOptions(request *QueryAppManagerRequest, headers *QueryAppManagerHeaders, runtime *util.RuntimeOptions) (_result *QueryAppManagerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

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
		Action:      tea.String("QueryAppManager"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/apps/managers/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAppManagerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户管理应用管理员
//
// @param request - QueryAppManagerRequest
//
// @return QueryAppManagerResponse
func (client *Client) QueryAppManager(request *QueryAppManagerRequest) (_result *QueryAppManagerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAppManagerHeaders{}
	_result = &QueryAppManagerResponse{}
	_body, _err := client.QueryAppManagerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询权益库存
//
// @param request - QueryBenefitInventoryRequest
//
// @param headers - QueryBenefitInventoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBenefitInventoryResponse
func (client *Client) QueryBenefitInventoryWithOptions(request *QueryBenefitInventoryRequest, headers *QueryBenefitInventoryHeaders, runtime *util.RuntimeOptions) (_result *QueryBenefitInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitCode)) {
		body["benefitCode"] = request.BenefitCode
	}

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
		Action:      tea.String("QueryBenefitInventory"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/benefitInventories/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBenefitInventoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询权益库存
//
// @param request - QueryBenefitInventoryRequest
//
// @return QueryBenefitInventoryResponse
func (client *Client) QueryBenefitInventory(request *QueryBenefitInventoryRequest) (_result *QueryBenefitInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBenefitInventoryHeaders{}
	_result = &QueryBenefitInventoryResponse{}
	_body, _err := client.QueryBenefitInventoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询线索跟进状态
//
// @param request - QueryClueFollowStatusRequest
//
// @param headers - QueryClueFollowStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryClueFollowStatusResponse
func (client *Client) QueryClueFollowStatusWithOptions(request *QueryClueFollowStatusRequest, headers *QueryClueFollowStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryClueFollowStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClueId)) {
		query["clueId"] = request.ClueId
	}

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
		Action:      tea.String("QueryClueFollowStatus"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/clues/followStatuses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryClueFollowStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询线索跟进状态
//
// @param request - QueryClueFollowStatusRequest
//
// @return QueryClueFollowStatusResponse
func (client *Client) QueryClueFollowStatus(request *QueryClueFollowStatusRequest) (_result *QueryClueFollowStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryClueFollowStatusHeaders{}
	_result = &QueryClueFollowStatusResponse{}
	_body, _err := client.QueryClueFollowStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询客户群
//
// @param request - QueryCrmGroupChatsRequest
//
// @param headers - QueryCrmGroupChatsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCrmGroupChatsResponse
func (client *Client) QueryCrmGroupChatsWithOptions(request *QueryCrmGroupChatsRequest, headers *QueryCrmGroupChatsHeaders, runtime *util.RuntimeOptions) (_result *QueryCrmGroupChatsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.QueryDsl)) {
		query["queryDsl"] = request.QueryDsl
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		query["relationType"] = request.RelationType
	}

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
		Action:      tea.String("QueryCrmGroupChats"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/crmGroupChats"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCrmGroupChatsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户群
//
// @param request - QueryCrmGroupChatsRequest
//
// @return QueryCrmGroupChatsResponse
func (client *Client) QueryCrmGroupChats(request *QueryCrmGroupChatsRequest) (_result *QueryCrmGroupChatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCrmGroupChatsHeaders{}
	_result = &QueryCrmGroupChatsResponse{}
	_body, _err := client.QueryCrmGroupChatsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据指定查询条件批量获取客户数据
//
// @param request - QueryCrmPersonalCustomerRequest
//
// @param headers - QueryCrmPersonalCustomerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCrmPersonalCustomerResponse
func (client *Client) QueryCrmPersonalCustomerWithOptions(request *QueryCrmPersonalCustomerRequest, headers *QueryCrmPersonalCustomerHeaders, runtime *util.RuntimeOptions) (_result *QueryCrmPersonalCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentOperatorUserId)) {
		query["currentOperatorUserId"] = request.CurrentOperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDsl)) {
		query["queryDsl"] = request.QueryDsl
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		query["relationType"] = request.RelationType
	}

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
		Action:      tea.String("QueryCrmPersonalCustomer"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/personalCustomers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCrmPersonalCustomerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据指定查询条件批量获取客户数据
//
// @param request - QueryCrmPersonalCustomerRequest
//
// @return QueryCrmPersonalCustomerResponse
func (client *Client) QueryCrmPersonalCustomer(request *QueryCrmPersonalCustomerRequest) (_result *QueryCrmPersonalCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCrmPersonalCustomerHeaders{}
	_result = &QueryCrmPersonalCustomerResponse{}
	_body, _err := client.QueryCrmPersonalCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询客户模板启用类型
//
// @param request - QueryCustomerBizTypeRequest
//
// @param headers - QueryCustomerBizTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerBizTypeResponse
func (client *Client) QueryCustomerBizTypeWithOptions(request *QueryCustomerBizTypeRequest, headers *QueryCustomerBizTypeHeaders, runtime *util.RuntimeOptions) (_result *QueryCustomerBizTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

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
		Action:      tea.String("QueryCustomerBizType"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/orgSettings/templates/customerBizTypes/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCustomerBizTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户模板启用类型
//
// @param request - QueryCustomerBizTypeRequest
//
// @return QueryCustomerBizTypeResponse
func (client *Client) QueryCustomerBizType(request *QueryCustomerBizTypeRequest) (_result *QueryCustomerBizTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCustomerBizTypeHeaders{}
	_result = &QueryCustomerBizTypeResponse{}
	_body, _err := client.QueryCustomerBizTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 营销服融合三方全局信息
//
// @param request - QueryGlobalInfoRequest
//
// @param headers - QueryGlobalInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGlobalInfoResponse
func (client *Client) QueryGlobalInfoWithOptions(request *QueryGlobalInfoRequest, headers *QueryGlobalInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryGlobalInfoResponse, _err error) {
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
		Action:      tea.String("QueryGlobalInfo"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/globalInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGlobalInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 营销服融合三方全局信息
//
// @param request - QueryGlobalInfoRequest
//
// @return QueryGlobalInfoResponse
func (client *Client) QueryGlobalInfo(request *QueryGlobalInfoRequest) (_result *QueryGlobalInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGlobalInfoHeaders{}
	_result = &QueryGlobalInfoResponse{}
	_body, _err := client.QueryGlobalInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户是否有应用管理员权限
//
// @param request - QueryHasAppPermissionRequest
//
// @param headers - QueryHasAppPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHasAppPermissionResponse
func (client *Client) QueryHasAppPermissionWithOptions(request *QueryHasAppPermissionRequest, headers *QueryHasAppPermissionHeaders, runtime *util.RuntimeOptions) (_result *QueryHasAppPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

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
		Action:      tea.String("QueryHasAppPermission"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/apps/adminPermissions/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHasAppPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户是否有应用管理员权限
//
// @param request - QueryHasAppPermissionRequest
//
// @return QueryHasAppPermissionResponse
func (client *Client) QueryHasAppPermission(request *QueryHasAppPermissionRequest) (_result *QueryHasAppPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHasAppPermissionHeaders{}
	_result = &QueryHasAppPermissionResponse{}
	_body, _err := client.QueryHasAppPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务窗用户基础信息
//
// @param request - QueryOfficialAccountUserBasicInfoRequest
//
// @param headers - QueryOfficialAccountUserBasicInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOfficialAccountUserBasicInfoResponse
func (client *Client) QueryOfficialAccountUserBasicInfoWithOptions(request *QueryOfficialAccountUserBasicInfoRequest, headers *QueryOfficialAccountUserBasicInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryOfficialAccountUserBasicInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingToken)) {
		query["bindingToken"] = request.BindingToken
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
		Action:      tea.String("QueryOfficialAccountUserBasicInfo"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/officialAccounts/basics/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOfficialAccountUserBasicInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务窗用户基础信息
//
// @param request - QueryOfficialAccountUserBasicInfoRequest
//
// @return QueryOfficialAccountUserBasicInfoResponse
func (client *Client) QueryOfficialAccountUserBasicInfo(request *QueryOfficialAccountUserBasicInfoRequest) (_result *QueryOfficialAccountUserBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOfficialAccountUserBasicInfoHeaders{}
	_result = &QueryOfficialAccountUserBasicInfoResponse{}
	_body, _err := client.QueryOfficialAccountUserBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据targetId查询关系数据
//
// @param request - QueryRelationDatasByTargetIdRequest
//
// @param headers - QueryRelationDatasByTargetIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRelationDatasByTargetIdResponse
func (client *Client) QueryRelationDatasByTargetIdWithOptions(targetId *string, request *QueryRelationDatasByTargetIdRequest, headers *QueryRelationDatasByTargetIdHeaders, runtime *util.RuntimeOptions) (_result *QueryRelationDatasByTargetIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		query["relationType"] = request.RelationType
	}

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
		Action:      tea.String("QueryRelationDatasByTargetId"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/relations/datas/targets/" + tea.StringValue(targetId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRelationDatasByTargetIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据targetId查询关系数据
//
// @param request - QueryRelationDatasByTargetIdRequest
//
// @return QueryRelationDatasByTargetIdResponse
func (client *Client) QueryRelationDatasByTargetId(targetId *string, request *QueryRelationDatasByTargetIdRequest) (_result *QueryRelationDatasByTargetIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRelationDatasByTargetIdHeaders{}
	_result = &QueryRelationDatasByTargetIdResponse{}
	_body, _err := client.QueryRelationDatasByTargetIdWithOptions(targetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务窗消息撤回
//
// @param request - RecallOfficialAccountOTOMessageRequest
//
// @param headers - RecallOfficialAccountOTOMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecallOfficialAccountOTOMessageResponse
func (client *Client) RecallOfficialAccountOTOMessageWithOptions(request *RecallOfficialAccountOTOMessageRequest, headers *RecallOfficialAccountOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *RecallOfficialAccountOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPushId)) {
		body["openPushId"] = request.OpenPushId
	}

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
		Action:      tea.String("RecallOfficialAccountOTOMessage"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/officialAccounts/oToMessages/recall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RecallOfficialAccountOTOMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务窗消息撤回
//
// @param request - RecallOfficialAccountOTOMessageRequest
//
// @return RecallOfficialAccountOTOMessageResponse
func (client *Client) RecallOfficialAccountOTOMessage(request *RecallOfficialAccountOTOMessageRequest) (_result *RecallOfficialAccountOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecallOfficialAccountOTOMessageHeaders{}
	_result = &RecallOfficialAccountOTOMessageResponse{}
	_body, _err := client.RecallOfficialAccountOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存license
//
// @param request - SaveBenefitLicenseRequest
//
// @param headers - SaveBenefitLicenseHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBenefitLicenseResponse
func (client *Client) SaveBenefitLicenseWithOptions(request *SaveBenefitLicenseRequest, headers *SaveBenefitLicenseHeaders, runtime *util.RuntimeOptions) (_result *SaveBenefitLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Licenses)) {
		body["licenses"] = request.Licenses
	}

	if !tea.BoolValue(util.IsUnset(request.SaveUserId)) {
		body["saveUserId"] = request.SaveUserId
	}

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
		Action:      tea.String("SaveBenefitLicense"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/benefitLicenses/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBenefitLicenseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存license
//
// @param request - SaveBenefitLicenseRequest
//
// @return SaveBenefitLicenseResponse
func (client *Client) SaveBenefitLicense(request *SaveBenefitLicenseRequest) (_result *SaveBenefitLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveBenefitLicenseHeaders{}
	_result = &SaveBenefitLicenseResponse{}
	_body, _err := client.SaveBenefitLicenseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务窗单发接口，指定消息接收人发送
//
// @param request - SendOfficialAccountOTOMessageRequest
//
// @param headers - SendOfficialAccountOTOMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendOfficialAccountOTOMessageResponse
func (client *Client) SendOfficialAccountOTOMessageWithOptions(request *SendOfficialAccountOTOMessageRequest, headers *SendOfficialAccountOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *SendOfficialAccountOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
	}

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
		Action:      tea.String("SendOfficialAccountOTOMessage"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/officialAccounts/oToMessages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendOfficialAccountOTOMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务窗单发接口，指定消息接收人发送
//
// @param request - SendOfficialAccountOTOMessageRequest
//
// @return SendOfficialAccountOTOMessageResponse
func (client *Client) SendOfficialAccountOTOMessage(request *SendOfficialAccountOTOMessageRequest) (_result *SendOfficialAccountOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendOfficialAccountOTOMessageHeaders{}
	_result = &SendOfficialAccountOTOMessageResponse{}
	_body, _err := client.SendOfficialAccountOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个人应用发送服务窗消息
//
// @param request - SendOfficialAccountSNSMessageRequest
//
// @param headers - SendOfficialAccountSNSMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendOfficialAccountSNSMessageResponse
func (client *Client) SendOfficialAccountSNSMessageWithOptions(request *SendOfficialAccountSNSMessageRequest, headers *SendOfficialAccountSNSMessageHeaders, runtime *util.RuntimeOptions) (_result *SendOfficialAccountSNSMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingToken)) {
		body["bindingToken"] = request.BindingToken
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
	}

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
		Action:      tea.String("SendOfficialAccountSNSMessage"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/officialAccounts/snsMessages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendOfficialAccountSNSMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个人应用发送服务窗消息
//
// @param request - SendOfficialAccountSNSMessageRequest
//
// @return SendOfficialAccountSNSMessageResponse
func (client *Client) SendOfficialAccountSNSMessage(request *SendOfficialAccountSNSMessageRequest) (_result *SendOfficialAccountSNSMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendOfficialAccountSNSMessageHeaders{}
	_result = &SendOfficialAccountSNSMessageResponse{}
	_body, _err := client.SendOfficialAccountSNSMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务窗消息群发
//
// @param request - ServiceWindowMessageBatchPushRequest
//
// @param headers - ServiceWindowMessageBatchPushHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ServiceWindowMessageBatchPushResponse
func (client *Client) ServiceWindowMessageBatchPushWithOptions(request *ServiceWindowMessageBatchPushRequest, headers *ServiceWindowMessageBatchPushHeaders, runtime *util.RuntimeOptions) (_result *ServiceWindowMessageBatchPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
	}

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
		Action:      tea.String("ServiceWindowMessageBatchPush"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/messages/batchSend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ServiceWindowMessageBatchPushResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务窗消息群发
//
// @param request - ServiceWindowMessageBatchPushRequest
//
// @return ServiceWindowMessageBatchPushResponse
func (client *Client) ServiceWindowMessageBatchPush(request *ServiceWindowMessageBatchPushRequest) (_result *ServiceWindowMessageBatchPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ServiceWindowMessageBatchPushHeaders{}
	_result = &ServiceWindowMessageBatchPushResponse{}
	_body, _err := client.ServiceWindowMessageBatchPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置用户版本为免费版
//
// @param request - SetUserVersionToFreeRequest
//
// @param headers - SetUserVersionToFreeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserVersionToFreeResponse
func (client *Client) SetUserVersionToFreeWithOptions(request *SetUserVersionToFreeRequest, headers *SetUserVersionToFreeHeaders, runtime *util.RuntimeOptions) (_result *SetUserVersionToFreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
	}

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
		Action:      tea.String("SetUserVersionToFree"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/versions/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetUserVersionToFreeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置用户版本为免费版
//
// @param request - SetUserVersionToFreeRequest
//
// @return SetUserVersionToFreeResponse
func (client *Client) SetUserVersionToFree(request *SetUserVersionToFreeRequest) (_result *SetUserVersionToFreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetUserVersionToFreeHeaders{}
	_result = &SetUserVersionToFreeResponse{}
	_body, _err := client.SetUserVersionToFreeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 二阶段提交权益库存结果
//
// @param request - TwoPhaseCommitInventoryRequest
//
// @param headers - TwoPhaseCommitInventoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TwoPhaseCommitInventoryResponse
func (client *Client) TwoPhaseCommitInventoryWithOptions(request *TwoPhaseCommitInventoryRequest, headers *TwoPhaseCommitInventoryHeaders, runtime *util.RuntimeOptions) (_result *TwoPhaseCommitInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitCode)) {
		body["benefitCode"] = request.BenefitCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizRequestId)) {
		body["bizRequestId"] = request.BizRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteResult)) {
		body["executeResult"] = request.ExecuteResult
	}

	if !tea.BoolValue(util.IsUnset(request.Quota)) {
		body["quota"] = request.Quota
	}

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
		Action:      tea.String("TwoPhaseCommitInventory"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/benefitInventories/twoPhases/commit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TwoPhaseCommitInventoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 二阶段提交权益库存结果
//
// @param request - TwoPhaseCommitInventoryRequest
//
// @return TwoPhaseCommitInventoryResponse
func (client *Client) TwoPhaseCommitInventory(request *TwoPhaseCommitInventoryRequest) (_result *TwoPhaseCommitInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TwoPhaseCommitInventoryHeaders{}
	_result = &TwoPhaseCommitInventoryResponse{}
	_body, _err := client.TwoPhaseCommitInventoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新crm个人客户（或企业客户）
//
// @param request - UpdateCrmPersonalCustomerRequest
//
// @param headers - UpdateCrmPersonalCustomerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCrmPersonalCustomerResponse
func (client *Client) UpdateCrmPersonalCustomerWithOptions(request *UpdateCrmPersonalCustomerRequest, headers *UpdateCrmPersonalCustomerHeaders, runtime *util.RuntimeOptions) (_result *UpdateCrmPersonalCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifierNick)) {
		body["modifierNick"] = request.ModifierNick
	}

	if !tea.BoolValue(util.IsUnset(request.ModifierUserId)) {
		body["modifierUserId"] = request.ModifierUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Permission)) {
		body["permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.SkipDuplicateCheck)) {
		body["skipDuplicateCheck"] = request.SkipDuplicateCheck
	}

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
		Action:      tea.String("UpdateCrmPersonalCustomer"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/personalCustomers"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCrmPersonalCustomerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新crm个人客户（或企业客户）
//
// @param request - UpdateCrmPersonalCustomerRequest
//
// @return UpdateCrmPersonalCustomerResponse
func (client *Client) UpdateCrmPersonalCustomer(request *UpdateCrmPersonalCustomerRequest) (_result *UpdateCrmPersonalCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCrmPersonalCustomerHeaders{}
	_result = &UpdateCrmPersonalCustomerResponse{}
	_body, _err := client.UpdateCrmPersonalCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新客户模板类型
//
// @param request - UpdateCustomerBizTypeRequest
//
// @param headers - UpdateCustomerBizTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomerBizTypeResponse
func (client *Client) UpdateCustomerBizTypeWithOptions(request *UpdateCustomerBizTypeRequest, headers *UpdateCustomerBizTypeHeaders, runtime *util.RuntimeOptions) (_result *UpdateCustomerBizTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerBizType)) {
		body["customerBizType"] = request.CustomerBizType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

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
		Action:      tea.String("UpdateCustomerBizType"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/orgSettings/templates/customerBizTypes"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomerBizTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新客户模板类型
//
// @param request - UpdateCustomerBizTypeRequest
//
// @return UpdateCustomerBizTypeResponse
func (client *Client) UpdateCustomerBizType(request *UpdateCustomerBizTypeRequest) (_result *UpdateCustomerBizTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCustomerBizTypeHeaders{}
	_result = &UpdateCustomerBizTypeResponse{}
	_body, _err := client.UpdateCustomerBizTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新群组
//
// @param request - UpdateGroupSetRequest
//
// @param headers - UpdateGroupSetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupSetResponse
func (client *Client) UpdateGroupSetWithOptions(request *UpdateGroupSetRequest, headers *UpdateGroupSetHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ManagerUserIds)) {
		body["managerUserIds"] = request.ManagerUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.MemberQuota)) {
		body["memberQuota"] = request.MemberQuota
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Notice)) {
		body["notice"] = request.Notice
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeToped)) {
		body["noticeToped"] = request.NoticeToped
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		body["ownerUserId"] = request.OwnerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Welcome)) {
		body["welcome"] = request.Welcome
	}

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
		Action:      tea.String("UpdateGroupSet"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/groupSets/set"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("boolean"),
	}
	_result = &UpdateGroupSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新群组
//
// @param request - UpdateGroupSetRequest
//
// @return UpdateGroupSetResponse
func (client *Client) UpdateGroupSet(request *UpdateGroupSetRequest) (_result *UpdateGroupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupSetHeaders{}
	_result = &UpdateGroupSetResponse{}
	_body, _err := client.UpdateGroupSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 增量同步导航数据
//
// @param request - UpdateMenuDataRequest
//
// @param headers - UpdateMenuDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMenuDataResponse
func (client *Client) UpdateMenuDataWithOptions(request *UpdateMenuDataRequest, headers *UpdateMenuDataHeaders, runtime *util.RuntimeOptions) (_result *UpdateMenuDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attr)) {
		body["attr"] = request.Attr
	}

	if !tea.BoolValue(util.IsUnset(request.BizTraceId)) {
		body["bizTraceId"] = request.BizTraceId
	}

	if !tea.BoolValue(util.IsUnset(request.Module)) {
		body["module"] = request.Module
	}

	if !tea.BoolValue(util.IsUnset(request.NavData)) {
		body["navData"] = request.NavData
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["operateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

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
		Action:      tea.String("UpdateMenuData"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/navigations/menus/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMenuDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增量同步导航数据
//
// @param request - UpdateMenuDataRequest
//
// @return UpdateMenuDataResponse
func (client *Client) UpdateMenuData(request *UpdateMenuDataRequest) (_result *UpdateMenuDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMenuDataHeaders{}
	_result = &UpdateMenuDataResponse{}
	_body, _err := client.UpdateMenuDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模型表结构更新字段
//
// @param request - UpdateMetaModelFieldRequest
//
// @param headers - UpdateMetaModelFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetaModelFieldResponse
func (client *Client) UpdateMetaModelFieldWithOptions(request *UpdateMetaModelFieldRequest, headers *UpdateMetaModelFieldHeaders, runtime *util.RuntimeOptions) (_result *UpdateMetaModelFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.FieldDTOList)) {
		body["fieldDTOList"] = request.FieldDTOList
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Tenant)) {
		body["tenant"] = request.Tenant
	}

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
		Action:      tea.String("UpdateMetaModelField"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/metas/models/fields"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMetaModelFieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模型表结构更新字段
//
// @param request - UpdateMetaModelFieldRequest
//
// @return UpdateMetaModelFieldResponse
func (client *Client) UpdateMetaModelField(request *UpdateMetaModelFieldRequest) (_result *UpdateMetaModelFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMetaModelFieldHeaders{}
	_result = &UpdateMetaModelFieldResponse{}
	_body, _err := client.UpdateMetaModelFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关系模型表结构更新字段
//
// @param request - UpdateRelationMetaFieldRequest
//
// @param headers - UpdateRelationMetaFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRelationMetaFieldResponse
func (client *Client) UpdateRelationMetaFieldWithOptions(request *UpdateRelationMetaFieldRequest, headers *UpdateRelationMetaFieldHeaders, runtime *util.RuntimeOptions) (_result *UpdateRelationMetaFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FieldDTOList)) {
		body["fieldDTOList"] = request.FieldDTOList
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		body["relationType"] = request.RelationType
	}

	if !tea.BoolValue(util.IsUnset(request.Tenant)) {
		body["tenant"] = request.Tenant
	}

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
		Action:      tea.String("UpdateRelationMetaField"),
		Version:     tea.String("crm_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/crm/relations/metas/fields"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRelationMetaFieldResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关系模型表结构更新字段
//
// @param request - UpdateRelationMetaFieldRequest
//
// @return UpdateRelationMetaFieldResponse
func (client *Client) UpdateRelationMetaField(request *UpdateRelationMetaFieldRequest) (_result *UpdateRelationMetaFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRelationMetaFieldHeaders{}
	_result = &UpdateRelationMetaFieldResponse{}
	_body, _err := client.UpdateRelationMetaFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
