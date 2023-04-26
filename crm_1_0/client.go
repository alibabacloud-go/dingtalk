// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package crm_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
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
	CustomTrackDesc *string   `json:"customTrackDesc,omitempty" xml:"customTrackDesc,omitempty"`
	InstanceIdList  []*string `json:"instanceIdList,omitempty" xml:"instanceIdList,omitempty" type:"Repeated"`
	OperatorUserId  *string   `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	OptType         *string   `json:"optType,omitempty" xml:"optType,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AbandonCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Action             *string                                  `json:"action,omitempty" xml:"action,omitempty"`
	CreatorNick        *string                                  `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	CreatorUserId      *string                                  `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Data               map[string]interface{}                   `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData         map[string]interface{}                   `json:"extendData,omitempty" xml:"extendData,omitempty"`
	Permission         *AddCrmPersonalCustomerRequestPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	RelationType       *string                                  `json:"relationType,omitempty" xml:"relationType,omitempty"`
	SkipDuplicateCheck *bool                                    `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	ExtraBizInfo   *string `json:"extraBizInfo,omitempty" xml:"extraBizInfo,omitempty"`
	IdempotentKey  *string `json:"idempotentKey,omitempty" xml:"idempotentKey,omitempty"`
	MaskedContent  *string `json:"maskedContent,omitempty" xml:"maskedContent,omitempty"`
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	RelationType   *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
	Type           *int32  `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddCustomerTrackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AssignTimestamp *int64                  `json:"assignTimestamp,omitempty" xml:"assignTimestamp,omitempty"`
	AssignUserId    *string                 `json:"assignUserId,omitempty" xml:"assignUserId,omitempty"`
	AssignedUserId  *string                 `json:"assignedUserId,omitempty" xml:"assignedUserId,omitempty"`
	Leads           []*AddLeadsRequestLeads `json:"leads,omitempty" xml:"leads,omitempty" type:"Repeated"`
	OutTaskId       *string                 `json:"outTaskId,omitempty" xml:"outTaskId,omitempty"`
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
	LeadsName  *string `json:"leadsName,omitempty" xml:"leadsName,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddLeadsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	FieldDTOList   []*AddRelationMetaFieldRequestFieldDTOList `json:"fieldDTOList,omitempty" xml:"fieldDTOList,omitempty" type:"Repeated"`
	OperatorUserId *string                                    `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	RelationType   *string                                    `json:"relationType,omitempty" xml:"relationType,omitempty"`
	Tenant         *string                                    `json:"tenant,omitempty" xml:"tenant,omitempty"`
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
	ComponentName *string                                       `json:"componentName,omitempty" xml:"componentName,omitempty"`
	Props         *AddRelationMetaFieldRequestFieldDTOListProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
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
	Align                  *string                                                `json:"align,omitempty" xml:"align,omitempty"`
	BizAlias               *string                                                `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice                 *int64                                                 `json:"choice,omitempty" xml:"choice,omitempty"`
	Content                *string                                                `json:"content,omitempty" xml:"content,omitempty"`
	Disabled               *bool                                                  `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration               *bool                                                  `json:"duration,omitempty" xml:"duration,omitempty"`
	FieldId                *string                                                `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format                 *string                                                `json:"format,omitempty" xml:"format,omitempty"`
	Invisible              *bool                                                  `json:"invisible,omitempty" xml:"invisible,omitempty"`
	Label                  *string                                                `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze    *bool                                                  `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Link                   *string                                                `json:"link,omitempty" xml:"link,omitempty"`
	NeedDetail             *string                                                `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	NotPrint               *string                                                `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	NotUpper               *string                                                `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options                []*AddRelationMetaFieldRequestFieldDTOListPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable              *bool                                                  `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder            *string                                                `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Required               *bool                                                  `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool                                                  `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Sortable               *bool                                                  `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Unit                   *string                                                `json:"unit,omitempty" xml:"unit,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddRelationMetaFieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OperatorUserId *string                                `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	RelationList   []*BatchAddContactsRequestRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
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
	BizDataList []*BatchAddContactsRequestRelationListBizDataList `json:"bizDataList,omitempty" xml:"bizDataList,omitempty" type:"Repeated"`
	BizExtMap   map[string]*string                                `json:"bizExtMap,omitempty" xml:"bizExtMap,omitempty"`
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

type BatchAddContactsRequestRelationListBizDataList struct {
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
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
	ErrorCode  *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchAddContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceList   []*BatchAddFollowRecordsRequestInstanceList `json:"instanceList,omitempty" xml:"instanceList,omitempty" type:"Repeated"`
	OperatorUserId *string                                     `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
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
	DataArray []*BatchAddFollowRecordsRequestInstanceListDataArray `json:"dataArray,omitempty" xml:"dataArray,omitempty" type:"Repeated"`
}

func (s BatchAddFollowRecordsRequestInstanceList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFollowRecordsRequestInstanceList) GoString() string {
	return s.String()
}

func (s *BatchAddFollowRecordsRequestInstanceList) SetDataArray(v []*BatchAddFollowRecordsRequestInstanceListDataArray) *BatchAddFollowRecordsRequestInstanceList {
	s.DataArray = v
	return s
}

type BatchAddFollowRecordsRequestInstanceListDataArray struct {
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
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
	ErrorCode  *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchAddFollowRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OperatorUserId     *string                                     `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	RelationList       []*BatchAddRelationDatasRequestRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
	RelationType       *string                                     `json:"relationType,omitempty" xml:"relationType,omitempty"`
	SkipDuplicateCheck *bool                                       `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
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
	BizDataList           []*BatchAddRelationDatasRequestRelationListBizDataList         `json:"bizDataList,omitempty" xml:"bizDataList,omitempty" type:"Repeated"`
	BizExtMap             map[string]*string                                             `json:"bizExtMap,omitempty" xml:"bizExtMap,omitempty"`
	RelationPermissionDTO *BatchAddRelationDatasRequestRelationListRelationPermissionDTO `json:"relationPermissionDTO,omitempty" xml:"relationPermissionDTO,omitempty" type:"Struct"`
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

type BatchAddRelationDatasRequestRelationListBizDataList struct {
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
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
	ErrorCode             *string   `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg              *string   `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RelationId            *string   `json:"relationId,omitempty" xml:"relationId,omitempty"`
	Success               *bool     `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchAddRelationDatasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceIds    []*string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	OperatorUserId *string   `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
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
	ErrorCode  *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchRemoveFollowRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId *string                                          `json:"accountId,omitempty" xml:"accountId,omitempty"`
	BizId     *string                                          `json:"bizId,omitempty" xml:"bizId,omitempty"`
	Detail    *BatchSendOfficialAccountOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
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
	BizRequestId *string                                                     `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	MessageBody  *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	MsgType      *string                                                     `json:"msgType,omitempty" xml:"msgType,omitempty"`
	SendToAll    *bool                                                       `json:"sendToAll,omitempty" xml:"sendToAll,omitempty"`
	UserIdList   []*string                                                   `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	Uuid         *string                                                     `json:"uuid,omitempty" xml:"uuid,omitempty"`
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
	Title             *string                                                                           `json:"title,omitempty" xml:"title,omitempty"`
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
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
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
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	PicUrl     *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	Text       *string `json:"text,omitempty" xml:"text,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
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
	Text  *string `json:"text,omitempty" xml:"text,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchSendOfficialAccountOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OperatorUserId *string                                   `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	RelationList   []*BatchUpdateContactsRequestRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
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
	BizExtMap   map[string]*string                                   `json:"bizExtMap,omitempty" xml:"bizExtMap,omitempty"`
	RelationId  *string                                              `json:"relationId,omitempty" xml:"relationId,omitempty"`
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
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
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
	ErrorCode  *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUpdateContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceList   []*BatchUpdateFollowRecordsRequestInstanceList `json:"instanceList,omitempty" xml:"instanceList,omitempty" type:"Repeated"`
	OperatorUserId *string                                        `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
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
	DataArray  []*BatchUpdateFollowRecordsRequestInstanceListDataArray `json:"dataArray,omitempty" xml:"dataArray,omitempty" type:"Repeated"`
	InstanceId *string                                                 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
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
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
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
	ErrorCode  *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUpdateFollowRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OperatorUserId     *string                                        `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	RelationList       []*BatchUpdateRelationDatasRequestRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
	RelationType       *string                                        `json:"relationType,omitempty" xml:"relationType,omitempty"`
	SkipDuplicateCheck *bool                                          `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
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
	BizExtMap   map[string]*string                                        `json:"bizExtMap,omitempty" xml:"bizExtMap,omitempty"`
	RelationId  *string                                                   `json:"relationId,omitempty" xml:"relationId,omitempty"`
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
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
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
	ErrorCode             *string   `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg              *string   `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RelationId            *string   `json:"relationId,omitempty" xml:"relationId,omitempty"`
	Success               *bool     `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUpdateRelationDatasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Contacts      []*CreateCustomerRequestContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	CreatorUserId *string                          `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Data          map[string]interface{}           `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData    map[string]interface{}           `json:"extendData,omitempty" xml:"extendData,omitempty"`
	InstanceId    *string                          `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ObjectType    *string                          `json:"objectType,omitempty" xml:"objectType,omitempty"`
	Permission    *CreateCustomerRequestPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	SaveOption    *CreateCustomerRequestSaveOption `json:"saveOption,omitempty" xml:"saveOption,omitempty" type:"Struct"`
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
	CustomerExistedPolicy                  *string `json:"customerExistedPolicy,omitempty" xml:"customerExistedPolicy,omitempty"`
	SkipDuplicateCheck                     *bool   `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
	SubscribePolicy                        *int64  `json:"subscribePolicy,omitempty" xml:"subscribePolicy,omitempty"`
	ThrowExceptionWhileSavingContactFailed *bool   `json:"throwExceptionWhileSavingContactFailed,omitempty" xml:"throwExceptionWhileSavingContactFailed,omitempty"`
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
	Contacts           []*CreateCustomerResponseBodyContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	CustomerInstanceId *string                               `json:"customerInstanceId,omitempty" xml:"customerInstanceId,omitempty"`
	ObjectType         *string                               `json:"objectType,omitempty" xml:"objectType,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupName     *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	MemberUserIds *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	OwnerUserId   *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	RelationType  *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CreatorUserId  *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	ManagerUserIds *string `json:"managerUserIds,omitempty" xml:"managerUserIds,omitempty"`
	MemberQuota    *int32  `json:"memberQuota,omitempty" xml:"memberQuota,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Notice         *string `json:"notice,omitempty" xml:"notice,omitempty"`
	NoticeToped    *int32  `json:"noticeToped,omitempty" xml:"noticeToped,omitempty"`
	OwnerUserId    *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	RelationType   *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	TemplateId     *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Welcome        *string `json:"welcome,omitempty" xml:"welcome,omitempty"`
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
	GmtCreate              *string                              `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *string                              `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InviteLink             *string                              `json:"inviteLink,omitempty" xml:"inviteLink,omitempty"`
	LastOpenConversationId *string                              `json:"lastOpenConversationId,omitempty" xml:"lastOpenConversationId,omitempty"`
	Manager                []*CreateGroupSetResponseBodyManager `json:"manager,omitempty" xml:"manager,omitempty" type:"Repeated"`
	ManagerUserIds         *string                              `json:"managerUserIds,omitempty" xml:"managerUserIds,omitempty"`
	MemberCount            *int64                               `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	MemberQuota            *int64                               `json:"memberQuota,omitempty" xml:"memberQuota,omitempty"`
	Name                   *string                              `json:"name,omitempty" xml:"name,omitempty"`
	Notice                 *string                              `json:"notice,omitempty" xml:"notice,omitempty"`
	NoticeToped            *int32                               `json:"noticeToped,omitempty" xml:"noticeToped,omitempty"`
	OpenGroupSetId         *string                              `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	Owner                  *CreateGroupSetResponseBodyOwner     `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	OwnerUserId            *string                              `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	RelationType           *string                              `json:"relationType,omitempty" xml:"relationType,omitempty"`
	TemplateId             *string                              `json:"templateId,omitempty" xml:"templateId,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OperatorUserId  *string                                   `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	RelationMetaDTO *CreateRelationMetaRequestRelationMetaDTO `json:"relationMetaDTO,omitempty" xml:"relationMetaDTO,omitempty" type:"Struct"`
	Tenant          *string                                   `json:"tenant,omitempty" xml:"tenant,omitempty"`
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
	Desc         *string                                          `json:"desc,omitempty" xml:"desc,omitempty"`
	Items        []*CreateRelationMetaRequestRelationMetaDTOItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Name         *string                                          `json:"name,omitempty" xml:"name,omitempty"`
	RelationType *string                                          `json:"relationType,omitempty" xml:"relationType,omitempty"`
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
	ComponentName *string                                             `json:"componentName,omitempty" xml:"componentName,omitempty"`
	Props         *CreateRelationMetaRequestRelationMetaDTOItemsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
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
	Align                  *string                                                      `json:"align,omitempty" xml:"align,omitempty"`
	BizAlias               *string                                                      `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice                 *int64                                                       `json:"choice,omitempty" xml:"choice,omitempty"`
	Content                *string                                                      `json:"content,omitempty" xml:"content,omitempty"`
	Disabled               *bool                                                        `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration               *bool                                                        `json:"duration,omitempty" xml:"duration,omitempty"`
	FieldId                *string                                                      `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format                 *string                                                      `json:"format,omitempty" xml:"format,omitempty"`
	Invisible              *bool                                                        `json:"invisible,omitempty" xml:"invisible,omitempty"`
	Label                  *string                                                      `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze    *bool                                                        `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Link                   *string                                                      `json:"link,omitempty" xml:"link,omitempty"`
	NeedDetail             *string                                                      `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	NotPrint               *string                                                      `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	NotUpper               *string                                                      `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options                []*CreateRelationMetaRequestRelationMetaDTOItemsPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable              *bool                                                        `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder            *string                                                      `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Required               *bool                                                        `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool                                                        `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Sortable               *bool                                                        `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Unit                   *string                                                      `json:"unit,omitempty" xml:"unit,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRelationMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCrmCustomObjectDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CurrentOperatorUserId *string `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCrmFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLeadsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	FieldIdList    []*string `json:"fieldIdList,omitempty" xml:"fieldIdList,omitempty" type:"Repeated"`
	OperatorUserId *string   `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	RelationType   *string   `json:"relationType,omitempty" xml:"relationType,omitempty"`
	Tenant         *string   `json:"tenant,omitempty" xml:"tenant,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRelationMetaFieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code       *string                                                    `json:"code,omitempty" xml:"code,omitempty"`
	Customized *bool                                                      `json:"customized,omitempty" xml:"customized,omitempty"`
	Fields     []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Name       *string                                                    `json:"name,omitempty" xml:"name,omitempty"`
	Status     *string                                                    `json:"status,omitempty" xml:"status,omitempty"`
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
	Aggregator *string `json:"aggregator,omitempty" xml:"aggregator,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCrmPersonalCustomerObjectMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OperatorUserId *string   `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	RelationTypes  []*string `json:"relationTypes,omitempty" xml:"relationTypes,omitempty" type:"Repeated"`
	Tenant         *string   `json:"tenant,omitempty" xml:"tenant,omitempty"`
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
	CreatorUserId      *string                                                     `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Desc               *string                                                     `json:"desc,omitempty" xml:"desc,omitempty"`
	GmtCreate          *string                                                     `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *string                                                     `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Items              []*DescribeRelationMetaResponseBodyRelationMetaDTOListItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Name               *string                                                     `json:"name,omitempty" xml:"name,omitempty"`
	RelationMetaCode   *string                                                     `json:"relationMetaCode,omitempty" xml:"relationMetaCode,omitempty"`
	RelationMetaStatus *string                                                     `json:"relationMetaStatus,omitempty" xml:"relationMetaStatus,omitempty"`
	RelationType       *string                                                     `json:"relationType,omitempty" xml:"relationType,omitempty"`
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
	Children      []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	ComponentName *string                                                             `json:"componentName,omitempty" xml:"componentName,omitempty"`
	Props         *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsProps      `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
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
	ComponentName *string                                                                `json:"componentName,omitempty" xml:"componentName,omitempty"`
	Props         *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
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
	ActionName             *string                                                                                    `json:"actionName,omitempty" xml:"actionName,omitempty"`
	Align                  *string                                                                                    `json:"align,omitempty" xml:"align,omitempty"`
	AvailableTemplates     []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsAvailableTemplates `json:"availableTemplates,omitempty" xml:"availableTemplates,omitempty" type:"Repeated"`
	BizAlias               *string                                                                                    `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice                 *int64                                                                                     `json:"choice,omitempty" xml:"choice,omitempty"`
	Content                *string                                                                                    `json:"content,omitempty" xml:"content,omitempty"`
	DataSource             *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSource           `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	DefaultColor           *string                                                                                    `json:"defaultColor,omitempty" xml:"defaultColor,omitempty"`
	Disabled               *bool                                                                                      `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration               *bool                                                                                      `json:"duration,omitempty" xml:"duration,omitempty"`
	DurationLabel          *string                                                                                    `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	FieldId                *string                                                                                    `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Fields                 []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFields             `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Format                 *string                                                                                    `json:"format,omitempty" xml:"format,omitempty"`
	Formula                *string                                                                                    `json:"formula,omitempty" xml:"formula,omitempty"`
	Invisible              *bool                                                                                      `json:"invisible,omitempty" xml:"invisible,omitempty"`
	Label                  *string                                                                                    `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze    *bool                                                                                      `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Limit                  *int64                                                                                     `json:"limit,omitempty" xml:"limit,omitempty"`
	Link                   *string                                                                                    `json:"link,omitempty" xml:"link,omitempty"`
	Mode                   *string                                                                                    `json:"mode,omitempty" xml:"mode,omitempty"`
	Multiple               *bool                                                                                      `json:"multiple,omitempty" xml:"multiple,omitempty"`
	NotPrint               *string                                                                                    `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	NotUpper               *string                                                                                    `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options                []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptions            `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable              *bool                                                                                      `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder            *string                                                                                    `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Quote                  *int64                                                                                     `json:"quote,omitempty" xml:"quote,omitempty"`
	Ratio                  *int64                                                                                     `json:"ratio,omitempty" xml:"ratio,omitempty"`
	RelateSource           []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSource       `json:"relateSource,omitempty" xml:"relateSource,omitempty" type:"Repeated"`
	Required               *bool                                                                                      `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool                                                                                      `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Rule                   []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRule               `json:"rule,omitempty" xml:"rule,omitempty" type:"Repeated"`
	Sortable               *bool                                                                                      `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Spread                 *bool                                                                                      `json:"spread,omitempty" xml:"spread,omitempty"`
	StatField              []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsStatField          `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	TableViewMode          *string                                                                                    `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
	Unit                   *string                                                                                    `json:"unit,omitempty" xml:"unit,omitempty"`
	VerticalPrint          *bool                                                                                      `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	Watermark              *bool                                                                                      `json:"watermark,omitempty" xml:"watermark,omitempty"`
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
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
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
	Params *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	Target *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	Type   *string                                                                                `json:"type,omitempty" xml:"type,omitempty"`
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
	FieldId    *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	Value      *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType  *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
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
	AppType  *int64  `json:"appType,omitempty" xml:"appType,omitempty"`
	AppUuid  *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	BizType  *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
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
	ComponentName *string                                                                                 `json:"componentName,omitempty" xml:"componentName,omitempty"`
	RelateProps   *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
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
	Align                  *string                                                                                            `json:"align,omitempty" xml:"align,omitempty"`
	BizAlias               *string                                                                                            `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice                 *int64                                                                                             `json:"choice,omitempty" xml:"choice,omitempty"`
	Content                *string                                                                                            `json:"content,omitempty" xml:"content,omitempty"`
	Disabled               *bool                                                                                              `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration               *bool                                                                                              `json:"duration,omitempty" xml:"duration,omitempty"`
	DurationLabel          *string                                                                                            `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	FieldId                *string                                                                                            `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format                 *string                                                                                            `json:"format,omitempty" xml:"format,omitempty"`
	Formula                *string                                                                                            `json:"formula,omitempty" xml:"formula,omitempty"`
	Invisible              *bool                                                                                              `json:"invisible,omitempty" xml:"invisible,omitempty"`
	Label                  *string                                                                                            `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze    *bool                                                                                              `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Limit                  *int64                                                                                             `json:"limit,omitempty" xml:"limit,omitempty"`
	Link                   *string                                                                                            `json:"link,omitempty" xml:"link,omitempty"`
	Mode                   *string                                                                                            `json:"mode,omitempty" xml:"mode,omitempty"`
	NotUpper               *string                                                                                            `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options                []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptions   `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable              *bool                                                                                              `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder            *string                                                                                            `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Ratio                  *int64                                                                                             `json:"ratio,omitempty" xml:"ratio,omitempty"`
	Required               *bool                                                                                              `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool                                                                                              `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Spread                 *bool                                                                                              `json:"spread,omitempty" xml:"spread,omitempty"`
	StatField              []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	Unit                   *string                                                                                            `json:"unit,omitempty" xml:"unit,omitempty"`
	VerticalPrint          *bool                                                                                              `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	Watermark              *bool                                                                                              `json:"watermark,omitempty" xml:"watermark,omitempty"`
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
	Extension *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsFieldsRelatePropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	Key       *string                                                                                                 `json:"key,omitempty" xml:"key,omitempty"`
	Value     *string                                                                                                 `json:"value,omitempty" xml:"value,omitempty"`
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
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Unit    *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Upper   *bool   `json:"upper,omitempty" xml:"upper,omitempty"`
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
	Extension *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	Key       *string                                                                                `json:"key,omitempty" xml:"key,omitempty"`
	Value     *string                                                                                `json:"value,omitempty" xml:"value,omitempty"`
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
	BizType    *string                                                                                      `json:"bizType,omitempty" xml:"bizType,omitempty"`
	DataSource *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	Fields     []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFields   `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
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
	Params *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	Target *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	Type   *string                                                                                            `json:"type,omitempty" xml:"type,omitempty"`
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
	FieldId    *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	Value      *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType  *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
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
	AppType  *int64  `json:"appType,omitempty" xml:"appType,omitempty"`
	AppUuid  *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	BizType  *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
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
	ComponentName *string                                                                                             `json:"componentName,omitempty" xml:"componentName,omitempty"`
	RelateProps   *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
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
	Align                  *string                                                                                                        `json:"align,omitempty" xml:"align,omitempty"`
	BizAlias               *string                                                                                                        `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice                 *int64                                                                                                         `json:"choice,omitempty" xml:"choice,omitempty"`
	Content                *string                                                                                                        `json:"content,omitempty" xml:"content,omitempty"`
	Disabled               *bool                                                                                                          `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration               *string                                                                                                        `json:"duration,omitempty" xml:"duration,omitempty"`
	FieldId                *string                                                                                                        `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format                 *string                                                                                                        `json:"format,omitempty" xml:"format,omitempty"`
	Formula                *string                                                                                                        `json:"formula,omitempty" xml:"formula,omitempty"`
	Invisible              *bool                                                                                                          `json:"invisible,omitempty" xml:"invisible,omitempty"`
	Label                  *string                                                                                                        `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze    *bool                                                                                                          `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Link                   *string                                                                                                        `json:"link,omitempty" xml:"link,omitempty"`
	Multi                  *int64                                                                                                         `json:"multi,omitempty" xml:"multi,omitempty"`
	NotUpper               *string                                                                                                        `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options                []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsOptions   `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable              *bool                                                                                                          `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder            *string                                                                                                        `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Quote                  *int64                                                                                                         `json:"quote,omitempty" xml:"quote,omitempty"`
	Required               *bool                                                                                                          `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool                                                                                                          `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	StatField              []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsChildrenPropsRelateSourceFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	Unit                   *string                                                                                                        `json:"unit,omitempty" xml:"unit,omitempty"`
	VerticalPrint          *bool                                                                                                          `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
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
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
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
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Unit    *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Upper   *bool   `json:"upper,omitempty" xml:"upper,omitempty"`
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
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
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
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Unit    *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Upper   *bool   `json:"upper,omitempty" xml:"upper,omitempty"`
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
	ActionName             *string                                                                            `json:"actionName,omitempty" xml:"actionName,omitempty"`
	Align                  *string                                                                            `json:"align,omitempty" xml:"align,omitempty"`
	AvailableTemplates     []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsAvailableTemplates `json:"availableTemplates,omitempty" xml:"availableTemplates,omitempty" type:"Repeated"`
	BizAlias               *string                                                                            `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice                 *int64                                                                             `json:"choice,omitempty" xml:"choice,omitempty"`
	Content                *string                                                                            `json:"content,omitempty" xml:"content,omitempty"`
	DataSource             *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource           `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	DefaultColor           *string                                                                            `json:"defaultColor,omitempty" xml:"defaultColor,omitempty"`
	Disabled               *bool                                                                              `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration               *bool                                                                              `json:"duration,omitempty" xml:"duration,omitempty"`
	DurationLabel          *string                                                                            `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	FieldId                *string                                                                            `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Fields                 []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFields             `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Format                 *string                                                                            `json:"format,omitempty" xml:"format,omitempty"`
	Formula                *string                                                                            `json:"formula,omitempty" xml:"formula,omitempty"`
	Invisible              *bool                                                                              `json:"invisible,omitempty" xml:"invisible,omitempty"`
	Label                  *string                                                                            `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze    *bool                                                                              `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Limit                  *int64                                                                             `json:"limit,omitempty" xml:"limit,omitempty"`
	Link                   *string                                                                            `json:"link,omitempty" xml:"link,omitempty"`
	Mode                   *string                                                                            `json:"mode,omitempty" xml:"mode,omitempty"`
	Multi                  *int64                                                                             `json:"multi,omitempty" xml:"multi,omitempty"`
	Multiple               *bool                                                                              `json:"multiple,omitempty" xml:"multiple,omitempty"`
	NeedDetail             *string                                                                            `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	NotPrint               *string                                                                            `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	NotUpper               *string                                                                            `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options                []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptions            `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable              *bool                                                                              `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder            *string                                                                            `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Quote                  *int64                                                                             `json:"quote,omitempty" xml:"quote,omitempty"`
	Ratio                  *int64                                                                             `json:"ratio,omitempty" xml:"ratio,omitempty"`
	RelateSource           []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSource       `json:"relateSource,omitempty" xml:"relateSource,omitempty" type:"Repeated"`
	Required               *bool                                                                              `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool                                                                              `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Rule                   []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRule               `json:"rule,omitempty" xml:"rule,omitempty" type:"Repeated"`
	Sortable               *bool                                                                              `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Spread                 *bool                                                                              `json:"spread,omitempty" xml:"spread,omitempty"`
	StatField              []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsStatField          `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	TableViewMode          *string                                                                            `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
	Unit                   *string                                                                            `json:"unit,omitempty" xml:"unit,omitempty"`
	VerticalPrint          *bool                                                                              `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	Watermark              *bool                                                                              `json:"watermark,omitempty" xml:"watermark,omitempty"`
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
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
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

type DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsDataSource struct {
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
	FieldId    *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	Value      *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType  *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
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
	ComponentName *string                                                                         `json:"componentName,omitempty" xml:"componentName,omitempty"`
	RelateProps   *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
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
	Align                  *string                                                                                    `json:"align,omitempty" xml:"align,omitempty"`
	BizAlias               *string                                                                                    `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice                 *int64                                                                                     `json:"choice,omitempty" xml:"choice,omitempty"`
	Content                *string                                                                                    `json:"content,omitempty" xml:"content,omitempty"`
	Disabled               *bool                                                                                      `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration               *string                                                                                    `json:"duration,omitempty" xml:"duration,omitempty"`
	DurationLabel          *string                                                                                    `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	FieldId                *string                                                                                    `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format                 *string                                                                                    `json:"format,omitempty" xml:"format,omitempty"`
	Formula                *string                                                                                    `json:"formula,omitempty" xml:"formula,omitempty"`
	Invisible              *bool                                                                                      `json:"invisible,omitempty" xml:"invisible,omitempty"`
	Label                  *string                                                                                    `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze    *bool                                                                                      `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Limit                  *int64                                                                                     `json:"limit,omitempty" xml:"limit,omitempty"`
	Link                   *string                                                                                    `json:"link,omitempty" xml:"link,omitempty"`
	Mode                   *string                                                                                    `json:"mode,omitempty" xml:"mode,omitempty"`
	NotUpper               *string                                                                                    `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options                []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptions   `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable              *bool                                                                                      `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder            *string                                                                                    `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Ratio                  *int64                                                                                     `json:"ratio,omitempty" xml:"ratio,omitempty"`
	Required               *bool                                                                                      `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool                                                                                      `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Spread                 *bool                                                                                      `json:"spread,omitempty" xml:"spread,omitempty"`
	StatField              []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	Unit                   *string                                                                                    `json:"unit,omitempty" xml:"unit,omitempty"`
	VerticalPrint          *bool                                                                                      `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	Watermark              *bool                                                                                      `json:"watermark,omitempty" xml:"watermark,omitempty"`
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
	Extension *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsFieldsRelatePropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	Key       *string                                                                                         `json:"key,omitempty" xml:"key,omitempty"`
	Value     *string                                                                                         `json:"value,omitempty" xml:"value,omitempty"`
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
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Unit    *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Upper   *bool   `json:"upper,omitempty" xml:"upper,omitempty"`
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
	Extension *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	Key       *string                                                                        `json:"key,omitempty" xml:"key,omitempty"`
	Value     *string                                                                        `json:"value,omitempty" xml:"value,omitempty"`
	Warn      *bool                                                                          `json:"warn,omitempty" xml:"warn,omitempty"`
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
	BizType    *string                                                                              `json:"bizType,omitempty" xml:"bizType,omitempty"`
	DataSource *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	Fields     []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFields   `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
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
	Params *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	Target *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	Type   *string                                                                                    `json:"type,omitempty" xml:"type,omitempty"`
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
	FieldId    *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	Value      *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType  *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
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
	AppType  *int64  `json:"appType,omitempty" xml:"appType,omitempty"`
	AppUuid  *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	BizType  *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
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
	ComponentName *string                                                                                     `json:"componentName,omitempty" xml:"componentName,omitempty"`
	RelateProps   *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelateProps `json:"relateProps,omitempty" xml:"relateProps,omitempty" type:"Struct"`
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
	Align                  *string                                                                                                `json:"align,omitempty" xml:"align,omitempty"`
	BizAlias               *string                                                                                                `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice                 *int64                                                                                                 `json:"choice,omitempty" xml:"choice,omitempty"`
	Content                *string                                                                                                `json:"content,omitempty" xml:"content,omitempty"`
	Disabled               *bool                                                                                                  `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration               *string                                                                                                `json:"duration,omitempty" xml:"duration,omitempty"`
	FieldId                *string                                                                                                `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format                 *string                                                                                                `json:"format,omitempty" xml:"format,omitempty"`
	Formula                *string                                                                                                `json:"formula,omitempty" xml:"formula,omitempty"`
	Invisible              *bool                                                                                                  `json:"invisible,omitempty" xml:"invisible,omitempty"`
	Label                  *string                                                                                                `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze    *bool                                                                                                  `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Link                   *string                                                                                                `json:"link,omitempty" xml:"link,omitempty"`
	Multi                  *int64                                                                                                 `json:"multi,omitempty" xml:"multi,omitempty"`
	NotUpper               *string                                                                                                `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options                []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptions   `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable              *bool                                                                                                  `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder            *string                                                                                                `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Quote                  *int64                                                                                                 `json:"quote,omitempty" xml:"quote,omitempty"`
	Required               *bool                                                                                                  `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool                                                                                                  `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	StatField              []*DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	Unit                   *string                                                                                                `json:"unit,omitempty" xml:"unit,omitempty"`
	VerticalPrint          *bool                                                                                                  `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
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
	Extension *DescribeRelationMetaResponseBodyRelationMetaDTOListItemsPropsRelateSourceFieldsRelatePropsOptionsExtension `json:"extension,omitempty" xml:"extension,omitempty" type:"Struct"`
	Key       *string                                                                                                     `json:"key,omitempty" xml:"key,omitempty"`
	Value     *string                                                                                                     `json:"value,omitempty" xml:"value,omitempty"`
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
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Unit    *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Upper   *bool   `json:"upper,omitempty" xml:"upper,omitempty"`
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
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
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
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Unit    *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Upper   *bool   `json:"upper,omitempty" xml:"upper,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRelationMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	HasMore    *bool                                           `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken  *string                                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	CustomerId         *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FollowUpActionTime *string `json:"followUpActionTime,omitempty" xml:"followUpActionTime,omitempty"`
	IsDeleted          *bool   `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	NotifyTime         *string `json:"notifyTime,omitempty" xml:"notifyTime,omitempty"`
	RecycleRuleId      *int64  `json:"recycleRuleId,omitempty" xml:"recycleRuleId,omitempty"`
	RecycleTime        *string `json:"recycleTime,omitempty" xml:"recycleTime,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAllCustomerRecyclesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ChatId             *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	GmtCreate          *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCrmGroupChatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GmtCreate          *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	IconUrl            *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	MemberCount        *int32  `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenGroupSetId     *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	OwnerUserId        *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	OwnerUserName      *string `json:"ownerUserName,omitempty" xml:"ownerUserName,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCrmGroupChatMultiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GmtCreate          *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCrmGroupChatSingleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BizType    *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
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
	DefaultRole       *bool                                                           `json:"defaultRole,omitempty" xml:"defaultRole,omitempty"`
	FieldScopes       []*GetCrmRolePermissionResponseBodyPermissionsFieldScopes       `json:"fieldScopes,omitempty" xml:"fieldScopes,omitempty" type:"Repeated"`
	ManagingScopeList []*GetCrmRolePermissionResponseBodyPermissionsManagingScopeList `json:"managingScopeList,omitempty" xml:"managingScopeList,omitempty" type:"Repeated"`
	OperateScopes     []*GetCrmRolePermissionResponseBodyPermissionsOperateScopes     `json:"operateScopes,omitempty" xml:"operateScopes,omitempty" type:"Repeated"`
	ResourceId        *string                                                         `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	RoleId            *float64                                                        `json:"roleId,omitempty" xml:"roleId,omitempty"`
	RoleMemberList    []*GetCrmRolePermissionResponseBodyPermissionsRoleMemberList    `json:"roleMemberList,omitempty" xml:"roleMemberList,omitempty" type:"Repeated"`
	RoleName          *string                                                         `json:"roleName,omitempty" xml:"roleName,omitempty"`
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
	FieldActions []*string `json:"fieldActions,omitempty" xml:"fieldActions,omitempty" type:"Repeated"`
	FieldId      *string   `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
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
	Ext     *GetCrmRolePermissionResponseBodyPermissionsManagingScopeListExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	Manager *bool                                                            `json:"manager,omitempty" xml:"manager,omitempty"`
	Type    *string                                                          `json:"type,omitempty" xml:"type,omitempty"`
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
	DeptIdList []*float64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	UserIdList []*string  `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
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
	HasAuth *bool   `json:"hasAuth,omitempty" xml:"hasAuth,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
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
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCrmRolePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RelationId *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	TypeGroup  *int32  `json:"typeGroup,omitempty" xml:"typeGroup,omitempty"`
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
	HasMore    *bool                                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken  *string                                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Content     *string                                                     `json:"content,omitempty" xml:"content,omitempty"`
	CreatorName *string                                                     `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	Detail      map[string]*string                                          `json:"detail,omitempty" xml:"detail,omitempty"`
	Format      *string                                                     `json:"format,omitempty" xml:"format,omitempty"`
	GmtCreate   *string                                                     `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	IsvInfo     *GetCustomerTracksByRelationIdResponseBodyResultListIsvInfo `json:"isvInfo,omitempty" xml:"isvInfo,omitempty" type:"Struct"`
	Title       *string                                                     `json:"title,omitempty" xml:"title,omitempty"`
	Type        *int32                                                      `json:"type,omitempty" xml:"type,omitempty"`
	TypeGroup   *int32                                                      `json:"typeGroup,omitempty" xml:"typeGroup,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCustomerTracksByRelationIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GmtCreate              *string                           `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *string                           `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	GroupChatCount         *int32                            `json:"groupChatCount,omitempty" xml:"groupChatCount,omitempty"`
	InviteLink             *string                           `json:"inviteLink,omitempty" xml:"inviteLink,omitempty"`
	LastOpenConversationId *string                           `json:"lastOpenConversationId,omitempty" xml:"lastOpenConversationId,omitempty"`
	Manager                []*GetGroupSetResponseBodyManager `json:"manager,omitempty" xml:"manager,omitempty" type:"Repeated"`
	ManagerUserIds         *string                           `json:"managerUserIds,omitempty" xml:"managerUserIds,omitempty"`
	MemberCount            *int32                            `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	MemberQuota            *int32                            `json:"memberQuota,omitempty" xml:"memberQuota,omitempty"`
	Name                   *string                           `json:"name,omitempty" xml:"name,omitempty"`
	Notice                 *string                           `json:"notice,omitempty" xml:"notice,omitempty"`
	NoticeToped            *int32                            `json:"noticeToped,omitempty" xml:"noticeToped,omitempty"`
	OpenGroupSetId         *string                           `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	Owner                  *GetGroupSetResponseBodyOwner     `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	OwnerUserId            *string                           `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	RelationType           *string                           `json:"relationType,omitempty" xml:"relationType,omitempty"`
	TemplateId             *string                           `json:"templateId,omitempty" xml:"templateId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CorpName  *string   `json:"corpName,omitempty" xml:"corpName,omitempty"`
	Mobile    *string   `json:"mobile,omitempty" xml:"mobile,omitempty"`
	StateCode *string   `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOfficialAccountContactInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	MaxResults *int64                                          `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Values     []*GetOfficialAccountContactsResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
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
	UserId   *string                                                 `json:"userId,omitempty" xml:"userId,omitempty"`
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
	CreateTime    *string                                                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorNick   *string                                                         `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	CreatorUserId *string                                                         `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Data          map[string]interface{}                                          `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData    map[string]interface{}                                          `json:"extendData,omitempty" xml:"extendData,omitempty"`
	InstanceId    *string                                                         `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ModifyTime    *string                                                         `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	Permission    *GetOfficialAccountContactsResponseBodyValuesContactsPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOfficialAccountContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId  *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
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
	RequestId *string                                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetOfficialAccountOTOMessageResultResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
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
	ReadUserIdList []*string `json:"readUserIdList,omitempty" xml:"readUserIdList,omitempty" type:"Repeated"`
	Status         *int64    `json:"status,omitempty" xml:"status,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOfficialAccountOTOMessageResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	FieldId  *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRelationUkSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BizDataList    []*JoinGroupSetRequestBizDataList `json:"bizDataList,omitempty" xml:"bizDataList,omitempty" type:"Repeated"`
	CorpId         *string                           `json:"corpId,omitempty" xml:"corpId,omitempty"`
	OpenGroupSetId *string                           `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	UnionId        *string                           `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
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
	ChatId             *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *JoinGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppUuid        *string                                               `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	CreatorNick    *string                                               `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	CreatorUserId  *string                                               `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Data           map[string]interface{}                                `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData     map[string]interface{}                                `json:"extendData,omitempty" xml:"extendData,omitempty"`
	FormCode       *string                                               `json:"formCode,omitempty" xml:"formCode,omitempty"`
	GmtCreate      *string                                               `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string                                               `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InstanceId     *string                                               `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ObjectType     *string                                               `json:"objectType,omitempty" xml:"objectType,omitempty"`
	Permission     *ListCrmPersonalCustomersResponseBodyResultPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	ProcInstStatus *string                                               `json:"procInstStatus,omitempty" xml:"procInstStatus,omitempty"`
	ProcOutResult  *string                                               `json:"procOutResult,omitempty" xml:"procOutResult,omitempty"`
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
	OwnerStaffIds       []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCrmPersonalCustomersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults   *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken    *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	QueryDsl     *string `json:"queryDsl,omitempty" xml:"queryDsl,omitempty"`
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
	HasMore    *bool                                 `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken  *string                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResultList []*ListGroupSetResponseBodyResultList `json:"resultList,omitempty" xml:"resultList,omitempty" type:"Repeated"`
	TotalCount *int32                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	GmtCreate              *string                                      `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *string                                      `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	GroupChatCount         *int32                                       `json:"groupChatCount,omitempty" xml:"groupChatCount,omitempty"`
	LastOpenConversationId *string                                      `json:"lastOpenConversationId,omitempty" xml:"lastOpenConversationId,omitempty"`
	Manager                []*ListGroupSetResponseBodyResultListManager `json:"manager,omitempty" xml:"manager,omitempty" type:"Repeated"`
	ManagerUserIds         *string                                      `json:"managerUserIds,omitempty" xml:"managerUserIds,omitempty"`
	MemberCount            *int32                                       `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	MemberQuota            *int32                                       `json:"memberQuota,omitempty" xml:"memberQuota,omitempty"`
	Name                   *string                                      `json:"name,omitempty" xml:"name,omitempty"`
	Notice                 *string                                      `json:"notice,omitempty" xml:"notice,omitempty"`
	NoticeToped            *int32                                       `json:"noticeToped,omitempty" xml:"noticeToped,omitempty"`
	OpenGroupSetId         *string                                      `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	Owner                  *ListGroupSetResponseBodyResultListOwner     `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	OwnerUserId            *string                                      `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	RelationType           *string                                      `json:"relationType,omitempty" xml:"relationType,omitempty"`
	TemplateId             *string                                      `json:"templateId,omitempty" xml:"templateId,omitempty"`
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
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults     *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ObjectType     *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
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
	MaxResults *int64                                      `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Values     []*QueryAllCustomerResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
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
	CreateTime            *string                                             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorNick           *string                                             `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	CreatorUserId         *string                                             `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Data                  map[string]interface{}                              `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData            map[string]interface{}                              `json:"extendData,omitempty" xml:"extendData,omitempty"`
	InstanceId            *string                                             `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ModifyTime            *string                                             `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	ObjectType            *string                                             `json:"objectType,omitempty" xml:"objectType,omitempty"`
	Permission            *QueryAllCustomerResponseBodyResultValuesPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	ProcessInstanceStatus *string                                             `json:"processInstanceStatus,omitempty" xml:"processInstanceStatus,omitempty"`
	ProcessOutResult      *string                                             `json:"processOutResult,omitempty" xml:"processOutResult,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAllCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Order      *string `json:"order,omitempty" xml:"order,omitempty"`
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
	HasMore    *bool                               `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	MaxResults *int32                              `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Values     []*QueryAllTracksResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
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
	BizId      *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	Creator    *string `json:"creator,omitempty" xml:"creator,omitempty"`
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	GmtCreate  *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	SubType    *int32  `json:"subType,omitempty" xml:"subType,omitempty"`
	Type       *int32  `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAllTracksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults   *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken    *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	QueryDsl     *string `json:"queryDsl,omitempty" xml:"queryDsl,omitempty"`
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
	HasMore    *bool                                       `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken  *string                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResultList []*QueryCrmGroupChatsResponseBodyResultList `json:"resultList,omitempty" xml:"resultList,omitempty" type:"Repeated"`
	TotalCount *int32                                      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	GmtCreate          *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	MemberCount        *int32  `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenGroupSetId     *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	OwnerUserId        *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	OwnerUserName      *string `json:"ownerUserName,omitempty" xml:"ownerUserName,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCrmGroupChatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults            *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken             *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	QueryDsl              *string `json:"queryDsl,omitempty" xml:"queryDsl,omitempty"`
	RelationType          *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
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
	HasMore    *bool                                         `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	MaxResults *int32                                        `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	CreatorNick    *string                                               `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	CreatorUserId  *string                                               `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Data           map[string]interface{}                                `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData     map[string]interface{}                                `json:"extendData,omitempty" xml:"extendData,omitempty"`
	GmtCreate      *string                                               `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string                                               `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InstanceId     *string                                               `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ObjectType     *string                                               `json:"objectType,omitempty" xml:"objectType,omitempty"`
	Permission     *QueryCrmPersonalCustomerResponseBodyValuesPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	ProcInstStatus *string                                               `json:"procInstStatus,omitempty" xml:"procInstStatus,omitempty"`
	ProcOutResult  *string                                               `json:"procOutResult,omitempty" xml:"procOutResult,omitempty"`
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
	OwnerStaffIds       []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BindingToken *string `json:"bindingToken,omitempty" xml:"bindingToken,omitempty"`
	UnionId      *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	RequestId *string                                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *QueryOfficialAccountUserBasicInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOfficialAccountUserBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BizDataList         []*QueryRelationDatasByTargetIdResponseBodyRelationsBizDataList `json:"bizDataList,omitempty" xml:"bizDataList,omitempty" type:"Repeated"`
	OpenConversationIds []*string                                                       `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	RelationId          *string                                                         `json:"relationId,omitempty" xml:"relationId,omitempty"`
	RelationType        *string                                                         `json:"relationType,omitempty" xml:"relationType,omitempty"`
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
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRelationDatasByTargetIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId  *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecallOfficialAccountOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountId *string                                     `json:"accountId,omitempty" xml:"accountId,omitempty"`
	BizId     *string                                     `json:"bizId,omitempty" xml:"bizId,omitempty"`
	Detail    *SendOfficialAccountOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
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
	MessageBody *SendOfficialAccountOTOMessageRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	MsgType     *string                                                `json:"msgType,omitempty" xml:"msgType,omitempty"`
	UnionId     *string                                                `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId      *string                                                `json:"userId,omitempty" xml:"userId,omitempty"`
	Uuid        *string                                                `json:"uuid,omitempty" xml:"uuid,omitempty"`
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
	Title             *string                                                                      `json:"title,omitempty" xml:"title,omitempty"`
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
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
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
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	PicUrl     *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	Text       *string `json:"text,omitempty" xml:"text,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
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
	Text  *string `json:"text,omitempty" xml:"text,omitempty"`
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
	RequestId *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *SendOfficialAccountOTOMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendOfficialAccountOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BindingToken *string                                     `json:"bindingToken,omitempty" xml:"bindingToken,omitempty"`
	BizId        *string                                     `json:"bizId,omitempty" xml:"bizId,omitempty"`
	Detail       *SendOfficialAccountSNSMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
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
	MessageBody *SendOfficialAccountSNSMessageRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	MsgType     *string                                                `json:"msgType,omitempty" xml:"msgType,omitempty"`
	Uuid        *string                                                `json:"uuid,omitempty" xml:"uuid,omitempty"`
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
	RequestId *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *SendOfficialAccountSNSMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendOfficialAccountSNSMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BizId  *string                                     `json:"bizId,omitempty" xml:"bizId,omitempty"`
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
	BizRequestId *string                                                `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	MessageBody  *ServiceWindowMessageBatchPushRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	MsgType      *string                                                `json:"msgType,omitempty" xml:"msgType,omitempty"`
	SendToAll    *bool                                                  `json:"sendToAll,omitempty" xml:"sendToAll,omitempty"`
	UserIdList   []*string                                              `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	Uuid         *string                                                `json:"uuid,omitempty" xml:"uuid,omitempty"`
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
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
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
	Text  *string `json:"text,omitempty" xml:"text,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ServiceWindowMessageBatchPushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Action             *string                                     `json:"action,omitempty" xml:"action,omitempty"`
	Data               map[string]interface{}                      `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData         map[string]interface{}                      `json:"extendData,omitempty" xml:"extendData,omitempty"`
	InstanceId         *string                                     `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ModifierNick       *string                                     `json:"modifierNick,omitempty" xml:"modifierNick,omitempty"`
	ModifierUserId     *string                                     `json:"modifierUserId,omitempty" xml:"modifierUserId,omitempty"`
	Permission         *UpdateCrmPersonalCustomerRequestPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	RelationType       *string                                     `json:"relationType,omitempty" xml:"relationType,omitempty"`
	SkipDuplicateCheck *bool                                       `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *bool              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	FieldDTOList   []*UpdateRelationMetaFieldRequestFieldDTOList `json:"fieldDTOList,omitempty" xml:"fieldDTOList,omitempty" type:"Repeated"`
	OperatorUserId *string                                       `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	RelationType   *string                                       `json:"relationType,omitempty" xml:"relationType,omitempty"`
	Tenant         *string                                       `json:"tenant,omitempty" xml:"tenant,omitempty"`
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
	ComponentName *string                                          `json:"componentName,omitempty" xml:"componentName,omitempty"`
	Props         *UpdateRelationMetaFieldRequestFieldDTOListProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
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
	Align                  *string                                                   `json:"align,omitempty" xml:"align,omitempty"`
	BizAlias               *string                                                   `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Choice                 *int64                                                    `json:"choice,omitempty" xml:"choice,omitempty"`
	Content                *string                                                   `json:"content,omitempty" xml:"content,omitempty"`
	Disabled               *bool                                                     `json:"disabled,omitempty" xml:"disabled,omitempty"`
	Duration               *bool                                                     `json:"duration,omitempty" xml:"duration,omitempty"`
	FieldId                *string                                                   `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Format                 *string                                                   `json:"format,omitempty" xml:"format,omitempty"`
	Invisible              *bool                                                     `json:"invisible,omitempty" xml:"invisible,omitempty"`
	Label                  *string                                                   `json:"label,omitempty" xml:"label,omitempty"`
	LabelEditableFreeze    *bool                                                     `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	Link                   *string                                                   `json:"link,omitempty" xml:"link,omitempty"`
	NeedDetail             *string                                                   `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
	NotPrint               *string                                                   `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	NotUpper               *string                                                   `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	Options                []*UpdateRelationMetaFieldRequestFieldDTOListPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	PayEnable              *bool                                                     `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	Placeholder            *string                                                   `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Required               *bool                                                     `json:"required,omitempty" xml:"required,omitempty"`
	RequiredEditableFreeze *bool                                                     `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Sortable               *bool                                                     `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Unit                   *string                                                   `json:"unit,omitempty" xml:"unit,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRelationMetaFieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
		ReqBodyType: tea.String("json"),
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
