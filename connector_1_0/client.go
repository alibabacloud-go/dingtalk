// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package connector_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateActionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateActionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateActionHeaders) GoString() string {
	return s.String()
}

func (s *CreateActionHeaders) SetCommonHeaders(v map[string]*string) *CreateActionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateActionHeaders) SetXAcsDingtalkAccessToken(v string) *CreateActionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateActionRequest struct {
	ActionInfo     []*CreateActionRequestActionInfo `json:"actionInfo,omitempty" xml:"actionInfo,omitempty" type:"Repeated"`
	IntegratorFlag *string                          `json:"integratorFlag,omitempty" xml:"integratorFlag,omitempty"`
}

func (s CreateActionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateActionRequest) GoString() string {
	return s.String()
}

func (s *CreateActionRequest) SetActionInfo(v []*CreateActionRequestActionInfo) *CreateActionRequest {
	s.ActionInfo = v
	return s
}

func (s *CreateActionRequest) SetIntegratorFlag(v string) *CreateActionRequest {
	s.IntegratorFlag = &v
	return s
}

type CreateActionRequestActionInfo struct {
	ApiPath               *string                                           `json:"apiPath,omitempty" xml:"apiPath,omitempty"`
	Description           *string                                           `json:"description,omitempty" xml:"description,omitempty"`
	DingConnectorId       *string                                           `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	InputMappingConfig    *CreateActionRequestActionInfoInputMappingConfig  `json:"inputMappingConfig,omitempty" xml:"inputMappingConfig,omitempty" type:"Struct"`
	InputSchema           *string                                           `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	IntegratorActionId    *string                                           `json:"integratorActionId,omitempty" xml:"integratorActionId,omitempty"`
	IntegratorConnectorId *string                                           `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	Name                  *string                                           `json:"name,omitempty" xml:"name,omitempty"`
	OutputDataRules       []*CreateActionRequestActionInfoOutputDataRules   `json:"outputDataRules,omitempty" xml:"outputDataRules,omitempty" type:"Repeated"`
	OutputMappingConfig   *CreateActionRequestActionInfoOutputMappingConfig `json:"outputMappingConfig,omitempty" xml:"outputMappingConfig,omitempty" type:"Struct"`
	OutputSchema          *string                                           `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
}

func (s CreateActionRequestActionInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateActionRequestActionInfo) GoString() string {
	return s.String()
}

func (s *CreateActionRequestActionInfo) SetApiPath(v string) *CreateActionRequestActionInfo {
	s.ApiPath = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetDescription(v string) *CreateActionRequestActionInfo {
	s.Description = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetDingConnectorId(v string) *CreateActionRequestActionInfo {
	s.DingConnectorId = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetInputMappingConfig(v *CreateActionRequestActionInfoInputMappingConfig) *CreateActionRequestActionInfo {
	s.InputMappingConfig = v
	return s
}

func (s *CreateActionRequestActionInfo) SetInputSchema(v string) *CreateActionRequestActionInfo {
	s.InputSchema = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetIntegratorActionId(v string) *CreateActionRequestActionInfo {
	s.IntegratorActionId = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetIntegratorConnectorId(v string) *CreateActionRequestActionInfo {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetName(v string) *CreateActionRequestActionInfo {
	s.Name = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetOutputDataRules(v []*CreateActionRequestActionInfoOutputDataRules) *CreateActionRequestActionInfo {
	s.OutputDataRules = v
	return s
}

func (s *CreateActionRequestActionInfo) SetOutputMappingConfig(v *CreateActionRequestActionInfoOutputMappingConfig) *CreateActionRequestActionInfo {
	s.OutputMappingConfig = v
	return s
}

func (s *CreateActionRequestActionInfo) SetOutputSchema(v string) *CreateActionRequestActionInfo {
	s.OutputSchema = &v
	return s
}

type CreateActionRequestActionInfoInputMappingConfig struct {
	CustomSchemaMapping *string `json:"customSchemaMapping,omitempty" xml:"customSchemaMapping,omitempty"`
	Rules               *string `json:"rules,omitempty" xml:"rules,omitempty"`
}

func (s CreateActionRequestActionInfoInputMappingConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateActionRequestActionInfoInputMappingConfig) GoString() string {
	return s.String()
}

func (s *CreateActionRequestActionInfoInputMappingConfig) SetCustomSchemaMapping(v string) *CreateActionRequestActionInfoInputMappingConfig {
	s.CustomSchemaMapping = &v
	return s
}

func (s *CreateActionRequestActionInfoInputMappingConfig) SetRules(v string) *CreateActionRequestActionInfoInputMappingConfig {
	s.Rules = &v
	return s
}

type CreateActionRequestActionInfoOutputDataRules struct {
	ExpectValue  *string `json:"expectValue,omitempty" xml:"expectValue,omitempty"`
	Operate      *string `json:"operate,omitempty" xml:"operate,omitempty"`
	PropertyPath *string `json:"propertyPath,omitempty" xml:"propertyPath,omitempty"`
}

func (s CreateActionRequestActionInfoOutputDataRules) String() string {
	return tea.Prettify(s)
}

func (s CreateActionRequestActionInfoOutputDataRules) GoString() string {
	return s.String()
}

func (s *CreateActionRequestActionInfoOutputDataRules) SetExpectValue(v string) *CreateActionRequestActionInfoOutputDataRules {
	s.ExpectValue = &v
	return s
}

func (s *CreateActionRequestActionInfoOutputDataRules) SetOperate(v string) *CreateActionRequestActionInfoOutputDataRules {
	s.Operate = &v
	return s
}

func (s *CreateActionRequestActionInfoOutputDataRules) SetPropertyPath(v string) *CreateActionRequestActionInfoOutputDataRules {
	s.PropertyPath = &v
	return s
}

type CreateActionRequestActionInfoOutputMappingConfig struct {
	CustomSchemaMapping *string `json:"customSchemaMapping,omitempty" xml:"customSchemaMapping,omitempty"`
	Rules               *string `json:"rules,omitempty" xml:"rules,omitempty"`
}

func (s CreateActionRequestActionInfoOutputMappingConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateActionRequestActionInfoOutputMappingConfig) GoString() string {
	return s.String()
}

func (s *CreateActionRequestActionInfoOutputMappingConfig) SetCustomSchemaMapping(v string) *CreateActionRequestActionInfoOutputMappingConfig {
	s.CustomSchemaMapping = &v
	return s
}

func (s *CreateActionRequestActionInfoOutputMappingConfig) SetRules(v string) *CreateActionRequestActionInfoOutputMappingConfig {
	s.Rules = &v
	return s
}

type CreateActionResponseBody struct {
	Item []*CreateActionResponseBodyItem `json:"item,omitempty" xml:"item,omitempty" type:"Repeated"`
}

func (s CreateActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateActionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateActionResponseBody) SetItem(v []*CreateActionResponseBodyItem) *CreateActionResponseBody {
	s.Item = v
	return s
}

type CreateActionResponseBodyItem struct {
	DingActionId          *string `json:"dingActionId,omitempty" xml:"dingActionId,omitempty"`
	DingConnectorId       *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	IntegratorActionId    *string `json:"integratorActionId,omitempty" xml:"integratorActionId,omitempty"`
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	SubErrCode            *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	SubErrMsg             *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	Success               *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateActionResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s CreateActionResponseBodyItem) GoString() string {
	return s.String()
}

func (s *CreateActionResponseBodyItem) SetDingActionId(v string) *CreateActionResponseBodyItem {
	s.DingActionId = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetDingConnectorId(v string) *CreateActionResponseBodyItem {
	s.DingConnectorId = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetIntegratorActionId(v string) *CreateActionResponseBodyItem {
	s.IntegratorActionId = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetIntegratorConnectorId(v string) *CreateActionResponseBodyItem {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetSubErrCode(v string) *CreateActionResponseBodyItem {
	s.SubErrCode = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetSubErrMsg(v string) *CreateActionResponseBodyItem {
	s.SubErrMsg = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetSuccess(v string) *CreateActionResponseBodyItem {
	s.Success = &v
	return s
}

type CreateActionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateActionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateActionResponse) GoString() string {
	return s.String()
}

func (s *CreateActionResponse) SetHeaders(v map[string]*string) *CreateActionResponse {
	s.Headers = v
	return s
}

func (s *CreateActionResponse) SetStatusCode(v int32) *CreateActionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateActionResponse) SetBody(v *CreateActionResponseBody) *CreateActionResponse {
	s.Body = v
	return s
}

type CreateConnectorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateConnectorHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorHeaders) GoString() string {
	return s.String()
}

func (s *CreateConnectorHeaders) SetCommonHeaders(v map[string]*string) *CreateConnectorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateConnectorHeaders) SetXAcsDingtalkAccessToken(v string) *CreateConnectorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateConnectorRequest struct {
	ConnectorInfo  []*CreateConnectorRequestConnectorInfo `json:"connectorInfo,omitempty" xml:"connectorInfo,omitempty" type:"Repeated"`
	IntegratorFlag *string                                `json:"integratorFlag,omitempty" xml:"integratorFlag,omitempty"`
}

func (s CreateConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorRequest) GoString() string {
	return s.String()
}

func (s *CreateConnectorRequest) SetConnectorInfo(v []*CreateConnectorRequestConnectorInfo) *CreateConnectorRequest {
	s.ConnectorInfo = v
	return s
}

func (s *CreateConnectorRequest) SetIntegratorFlag(v string) *CreateConnectorRequest {
	s.IntegratorFlag = &v
	return s
}

type CreateConnectorRequestConnectorInfo struct {
	ApiDomain             *string `json:"apiDomain,omitempty" xml:"apiDomain,omitempty"`
	ApiSecret             *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	AppId                 *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
	AuthValueEnv          *bool   `json:"authValueEnv,omitempty" xml:"authValueEnv,omitempty"`
	Description           *string `json:"description,omitempty" xml:"description,omitempty"`
	DomainEnv             *bool   `json:"domainEnv,omitempty" xml:"domainEnv,omitempty"`
	IconMediaId           *string `json:"iconMediaId,omitempty" xml:"iconMediaId,omitempty"`
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateConnectorRequestConnectorInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorRequestConnectorInfo) GoString() string {
	return s.String()
}

func (s *CreateConnectorRequestConnectorInfo) SetApiDomain(v string) *CreateConnectorRequestConnectorInfo {
	s.ApiDomain = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetApiSecret(v string) *CreateConnectorRequestConnectorInfo {
	s.ApiSecret = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetAppId(v int64) *CreateConnectorRequestConnectorInfo {
	s.AppId = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetAuthValueEnv(v bool) *CreateConnectorRequestConnectorInfo {
	s.AuthValueEnv = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetDescription(v string) *CreateConnectorRequestConnectorInfo {
	s.Description = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetDomainEnv(v bool) *CreateConnectorRequestConnectorInfo {
	s.DomainEnv = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetIconMediaId(v string) *CreateConnectorRequestConnectorInfo {
	s.IconMediaId = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetIntegratorConnectorId(v string) *CreateConnectorRequestConnectorInfo {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetName(v string) *CreateConnectorRequestConnectorInfo {
	s.Name = &v
	return s
}

type CreateConnectorResponseBody struct {
	Item []*CreateConnectorResponseBodyItem `json:"item,omitempty" xml:"item,omitempty" type:"Repeated"`
}

func (s CreateConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConnectorResponseBody) SetItem(v []*CreateConnectorResponseBodyItem) *CreateConnectorResponseBody {
	s.Item = v
	return s
}

type CreateConnectorResponseBodyItem struct {
	DingConnectorId       *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	SubErrCode            *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	SubErrMsg             *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	Success               *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateConnectorResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorResponseBodyItem) GoString() string {
	return s.String()
}

func (s *CreateConnectorResponseBodyItem) SetDingConnectorId(v string) *CreateConnectorResponseBodyItem {
	s.DingConnectorId = &v
	return s
}

func (s *CreateConnectorResponseBodyItem) SetIntegratorConnectorId(v string) *CreateConnectorResponseBodyItem {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateConnectorResponseBodyItem) SetSubErrCode(v string) *CreateConnectorResponseBodyItem {
	s.SubErrCode = &v
	return s
}

func (s *CreateConnectorResponseBodyItem) SetSubErrMsg(v string) *CreateConnectorResponseBodyItem {
	s.SubErrMsg = &v
	return s
}

func (s *CreateConnectorResponseBodyItem) SetSuccess(v bool) *CreateConnectorResponseBodyItem {
	s.Success = &v
	return s
}

type CreateConnectorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorResponse) GoString() string {
	return s.String()
}

func (s *CreateConnectorResponse) SetHeaders(v map[string]*string) *CreateConnectorResponse {
	s.Headers = v
	return s
}

func (s *CreateConnectorResponse) SetStatusCode(v int32) *CreateConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConnectorResponse) SetBody(v *CreateConnectorResponseBody) *CreateConnectorResponse {
	s.Body = v
	return s
}

type CreateInvocableInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateInvocableInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateInvocableInstanceHeaders) GoString() string {
	return s.String()
}

func (s *CreateInvocableInstanceHeaders) SetCommonHeaders(v map[string]*string) *CreateInvocableInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateInvocableInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateInvocableInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateInvocableInstanceRequest struct {
	ConnectAssetUri *string `json:"connectAssetUri,omitempty" xml:"connectAssetUri,omitempty"`
	InstanceKey     *string `json:"instanceKey,omitempty" xml:"instanceKey,omitempty"`
}

func (s CreateInvocableInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInvocableInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInvocableInstanceRequest) SetConnectAssetUri(v string) *CreateInvocableInstanceRequest {
	s.ConnectAssetUri = &v
	return s
}

func (s *CreateInvocableInstanceRequest) SetInstanceKey(v string) *CreateInvocableInstanceRequest {
	s.InstanceKey = &v
	return s
}

type CreateInvocableInstanceResponseBody struct {
	ConnectAssetUri *string `json:"connectAssetUri,omitempty" xml:"connectAssetUri,omitempty"`
	VersionId       *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s CreateInvocableInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInvocableInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInvocableInstanceResponseBody) SetConnectAssetUri(v string) *CreateInvocableInstanceResponseBody {
	s.ConnectAssetUri = &v
	return s
}

func (s *CreateInvocableInstanceResponseBody) SetVersionId(v string) *CreateInvocableInstanceResponseBody {
	s.VersionId = &v
	return s
}

type CreateInvocableInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInvocableInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInvocableInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInvocableInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInvocableInstanceResponse) SetHeaders(v map[string]*string) *CreateInvocableInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInvocableInstanceResponse) SetStatusCode(v int32) *CreateInvocableInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInvocableInstanceResponse) SetBody(v *CreateInvocableInstanceResponseBody) *CreateInvocableInstanceResponse {
	s.Body = v
	return s
}

type CreateTriggerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTriggerHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerHeaders) GoString() string {
	return s.String()
}

func (s *CreateTriggerHeaders) SetCommonHeaders(v map[string]*string) *CreateTriggerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTriggerHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTriggerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTriggerRequest struct {
	IntegratorFlag *string                            `json:"integratorFlag,omitempty" xml:"integratorFlag,omitempty"`
	TriggerInfo    []*CreateTriggerRequestTriggerInfo `json:"triggerInfo,omitempty" xml:"triggerInfo,omitempty" type:"Repeated"`
}

func (s CreateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequest) SetIntegratorFlag(v string) *CreateTriggerRequest {
	s.IntegratorFlag = &v
	return s
}

func (s *CreateTriggerRequest) SetTriggerInfo(v []*CreateTriggerRequestTriggerInfo) *CreateTriggerRequest {
	s.TriggerInfo = v
	return s
}

type CreateTriggerRequestTriggerInfo struct {
	Description           *string `json:"description,omitempty" xml:"description,omitempty"`
	DingConnectorId       *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	InputSchema           *string `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	IntegratorTriggerId   *string `json:"integratorTriggerId,omitempty" xml:"integratorTriggerId,omitempty"`
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateTriggerRequestTriggerInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequestTriggerInfo) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequestTriggerInfo) SetDescription(v string) *CreateTriggerRequestTriggerInfo {
	s.Description = &v
	return s
}

func (s *CreateTriggerRequestTriggerInfo) SetDingConnectorId(v string) *CreateTriggerRequestTriggerInfo {
	s.DingConnectorId = &v
	return s
}

func (s *CreateTriggerRequestTriggerInfo) SetInputSchema(v string) *CreateTriggerRequestTriggerInfo {
	s.InputSchema = &v
	return s
}

func (s *CreateTriggerRequestTriggerInfo) SetIntegratorConnectorId(v string) *CreateTriggerRequestTriggerInfo {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateTriggerRequestTriggerInfo) SetIntegratorTriggerId(v string) *CreateTriggerRequestTriggerInfo {
	s.IntegratorTriggerId = &v
	return s
}

func (s *CreateTriggerRequestTriggerInfo) SetName(v string) *CreateTriggerRequestTriggerInfo {
	s.Name = &v
	return s
}

type CreateTriggerResponseBody struct {
	Item []*CreateTriggerResponseBodyItem `json:"item,omitempty" xml:"item,omitempty" type:"Repeated"`
}

func (s CreateTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponseBody) SetItem(v []*CreateTriggerResponseBodyItem) *CreateTriggerResponseBody {
	s.Item = v
	return s
}

type CreateTriggerResponseBodyItem struct {
	DingConnectorId       *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	DingTriggerId         *string `json:"dingTriggerId,omitempty" xml:"dingTriggerId,omitempty"`
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	IntegratorTriggerId   *string `json:"integratorTriggerId,omitempty" xml:"integratorTriggerId,omitempty"`
	SubErrCode            *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	SubErrMsg             *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	Success               *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTriggerResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponseBodyItem) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponseBodyItem) SetDingConnectorId(v string) *CreateTriggerResponseBodyItem {
	s.DingConnectorId = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetDingTriggerId(v string) *CreateTriggerResponseBodyItem {
	s.DingTriggerId = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetIntegratorConnectorId(v string) *CreateTriggerResponseBodyItem {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetIntegratorTriggerId(v string) *CreateTriggerResponseBodyItem {
	s.IntegratorTriggerId = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetSubErrCode(v string) *CreateTriggerResponseBodyItem {
	s.SubErrCode = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetSubErrMsg(v string) *CreateTriggerResponseBodyItem {
	s.SubErrMsg = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetSuccess(v bool) *CreateTriggerResponseBodyItem {
	s.Success = &v
	return s
}

type CreateTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponse) SetHeaders(v map[string]*string) *CreateTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateTriggerResponse) SetStatusCode(v int32) *CreateTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTriggerResponse) SetBody(v *CreateTriggerResponseBody) *CreateTriggerResponse {
	s.Body = v
	return s
}

type GetActionDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetActionDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetActionDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetActionDetailHeaders) SetCommonHeaders(v map[string]*string) *GetActionDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetActionDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetActionDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetActionDetailRequest struct {
	ConnectAssetUri *string `json:"connectAssetUri,omitempty" xml:"connectAssetUri,omitempty"`
}

func (s GetActionDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetActionDetailRequest) GoString() string {
	return s.String()
}

func (s *GetActionDetailRequest) SetConnectAssetUri(v string) *GetActionDetailRequest {
	s.ConnectAssetUri = &v
	return s
}

type GetActionDetailResponseBody struct {
	ConnectAssetUri   *string                                       `json:"connectAssetUri,omitempty" xml:"connectAssetUri,omitempty"`
	InputSchema       *string                                       `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	IntegrationConfig *GetActionDetailResponseBodyIntegrationConfig `json:"integrationConfig,omitempty" xml:"integrationConfig,omitempty" type:"Struct"`
	Name              *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	OutputSchema      *string                                       `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
	RefId             *string                                       `json:"refId,omitempty" xml:"refId,omitempty"`
	RefProviderCorpId *string                                       `json:"refProviderCorpId,omitempty" xml:"refProviderCorpId,omitempty"`
	RefType           *string                                       `json:"refType,omitempty" xml:"refType,omitempty"`
}

func (s GetActionDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetActionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetActionDetailResponseBody) SetConnectAssetUri(v string) *GetActionDetailResponseBody {
	s.ConnectAssetUri = &v
	return s
}

func (s *GetActionDetailResponseBody) SetInputSchema(v string) *GetActionDetailResponseBody {
	s.InputSchema = &v
	return s
}

func (s *GetActionDetailResponseBody) SetIntegrationConfig(v *GetActionDetailResponseBodyIntegrationConfig) *GetActionDetailResponseBody {
	s.IntegrationConfig = v
	return s
}

func (s *GetActionDetailResponseBody) SetName(v string) *GetActionDetailResponseBody {
	s.Name = &v
	return s
}

func (s *GetActionDetailResponseBody) SetOutputSchema(v string) *GetActionDetailResponseBody {
	s.OutputSchema = &v
	return s
}

func (s *GetActionDetailResponseBody) SetRefId(v string) *GetActionDetailResponseBody {
	s.RefId = &v
	return s
}

func (s *GetActionDetailResponseBody) SetRefProviderCorpId(v string) *GetActionDetailResponseBody {
	s.RefProviderCorpId = &v
	return s
}

func (s *GetActionDetailResponseBody) SetRefType(v string) *GetActionDetailResponseBody {
	s.RefType = &v
	return s
}

type GetActionDetailResponseBodyIntegrationConfig struct {
	CategoryNames []*GetActionDetailResponseBodyIntegrationConfigCategoryNames `json:"categoryNames,omitempty" xml:"categoryNames,omitempty" type:"Repeated"`
	EntityName    *string                                                      `json:"entityName,omitempty" xml:"entityName,omitempty"`
	Props         []*GetActionDetailResponseBodyIntegrationConfigProps         `json:"props,omitempty" xml:"props,omitempty" type:"Repeated"`
}

func (s GetActionDetailResponseBodyIntegrationConfig) String() string {
	return tea.Prettify(s)
}

func (s GetActionDetailResponseBodyIntegrationConfig) GoString() string {
	return s.String()
}

func (s *GetActionDetailResponseBodyIntegrationConfig) SetCategoryNames(v []*GetActionDetailResponseBodyIntegrationConfigCategoryNames) *GetActionDetailResponseBodyIntegrationConfig {
	s.CategoryNames = v
	return s
}

func (s *GetActionDetailResponseBodyIntegrationConfig) SetEntityName(v string) *GetActionDetailResponseBodyIntegrationConfig {
	s.EntityName = &v
	return s
}

func (s *GetActionDetailResponseBodyIntegrationConfig) SetProps(v []*GetActionDetailResponseBodyIntegrationConfigProps) *GetActionDetailResponseBodyIntegrationConfig {
	s.Props = v
	return s
}

type GetActionDetailResponseBodyIntegrationConfigCategoryNames struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetActionDetailResponseBodyIntegrationConfigCategoryNames) String() string {
	return tea.Prettify(s)
}

func (s GetActionDetailResponseBodyIntegrationConfigCategoryNames) GoString() string {
	return s.String()
}

func (s *GetActionDetailResponseBodyIntegrationConfigCategoryNames) SetValue(v string) *GetActionDetailResponseBodyIntegrationConfigCategoryNames {
	s.Value = &v
	return s
}

type GetActionDetailResponseBodyIntegrationConfigProps struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetActionDetailResponseBodyIntegrationConfigProps) String() string {
	return tea.Prettify(s)
}

func (s GetActionDetailResponseBodyIntegrationConfigProps) GoString() string {
	return s.String()
}

func (s *GetActionDetailResponseBodyIntegrationConfigProps) SetKey(v string) *GetActionDetailResponseBodyIntegrationConfigProps {
	s.Key = &v
	return s
}

func (s *GetActionDetailResponseBodyIntegrationConfigProps) SetValue(v string) *GetActionDetailResponseBodyIntegrationConfigProps {
	s.Value = &v
	return s
}

type GetActionDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetActionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetActionDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetActionDetailResponse) GoString() string {
	return s.String()
}

func (s *GetActionDetailResponse) SetHeaders(v map[string]*string) *GetActionDetailResponse {
	s.Headers = v
	return s
}

func (s *GetActionDetailResponse) SetStatusCode(v int32) *GetActionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetActionDetailResponse) SetBody(v *GetActionDetailResponseBody) *GetActionDetailResponse {
	s.Body = v
	return s
}

type InvokeInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InvokeInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvokeInstanceHeaders) GoString() string {
	return s.String()
}

func (s *InvokeInstanceHeaders) SetCommonHeaders(v map[string]*string) *InvokeInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *InvokeInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InvokeInstanceRequest struct {
	ConnectAssetUri *string `json:"connectAssetUri,omitempty" xml:"connectAssetUri,omitempty"`
	InputJsonString *string `json:"inputJsonString,omitempty" xml:"inputJsonString,omitempty"`
	InstanceKey     *string `json:"instanceKey,omitempty" xml:"instanceKey,omitempty"`
}

func (s InvokeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeInstanceRequest) GoString() string {
	return s.String()
}

func (s *InvokeInstanceRequest) SetConnectAssetUri(v string) *InvokeInstanceRequest {
	s.ConnectAssetUri = &v
	return s
}

func (s *InvokeInstanceRequest) SetInputJsonString(v string) *InvokeInstanceRequest {
	s.InputJsonString = &v
	return s
}

func (s *InvokeInstanceRequest) SetInstanceKey(v string) *InvokeInstanceRequest {
	s.InstanceKey = &v
	return s
}

type InvokeInstanceResponseBody struct {
	Cost         *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	InstanceId   *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OutputJson   *string `json:"outputJson,omitempty" xml:"outputJson,omitempty"`
}

func (s InvokeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeInstanceResponseBody) SetCost(v int64) *InvokeInstanceResponseBody {
	s.Cost = &v
	return s
}

func (s *InvokeInstanceResponseBody) SetErrorCode(v string) *InvokeInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InvokeInstanceResponseBody) SetErrorMessage(v string) *InvokeInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *InvokeInstanceResponseBody) SetInstanceId(v string) *InvokeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *InvokeInstanceResponseBody) SetOutputJson(v string) *InvokeInstanceResponseBody {
	s.OutputJson = &v
	return s
}

type InvokeInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvokeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeInstanceResponse) GoString() string {
	return s.String()
}

func (s *InvokeInstanceResponse) SetHeaders(v map[string]*string) *InvokeInstanceResponse {
	s.Headers = v
	return s
}

func (s *InvokeInstanceResponse) SetStatusCode(v int32) *InvokeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeInstanceResponse) SetBody(v *InvokeInstanceResponseBody) *InvokeInstanceResponse {
	s.Body = v
	return s
}

type PullDataByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PullDataByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPageHeaders) GoString() string {
	return s.String()
}

func (s *PullDataByPageHeaders) SetCommonHeaders(v map[string]*string) *PullDataByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PullDataByPageHeaders) SetXAcsDingtalkAccessToken(v string) *PullDataByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PullDataByPageRequest struct {
	AppId               *string `json:"appId,omitempty" xml:"appId,omitempty"`
	DataModelId         *string `json:"dataModelId,omitempty" xml:"dataModelId,omitempty"`
	DatetimeFilterField *string `json:"datetimeFilterField,omitempty" xml:"datetimeFilterField,omitempty"`
	MaxDatetime         *int64  `json:"maxDatetime,omitempty" xml:"maxDatetime,omitempty"`
	MaxResults          *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	MinDatetime         *int64  `json:"minDatetime,omitempty" xml:"minDatetime,omitempty"`
	NextToken           *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s PullDataByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPageRequest) GoString() string {
	return s.String()
}

func (s *PullDataByPageRequest) SetAppId(v string) *PullDataByPageRequest {
	s.AppId = &v
	return s
}

func (s *PullDataByPageRequest) SetDataModelId(v string) *PullDataByPageRequest {
	s.DataModelId = &v
	return s
}

func (s *PullDataByPageRequest) SetDatetimeFilterField(v string) *PullDataByPageRequest {
	s.DatetimeFilterField = &v
	return s
}

func (s *PullDataByPageRequest) SetMaxDatetime(v int64) *PullDataByPageRequest {
	s.MaxDatetime = &v
	return s
}

func (s *PullDataByPageRequest) SetMaxResults(v int64) *PullDataByPageRequest {
	s.MaxResults = &v
	return s
}

func (s *PullDataByPageRequest) SetMinDatetime(v int64) *PullDataByPageRequest {
	s.MinDatetime = &v
	return s
}

func (s *PullDataByPageRequest) SetNextToken(v string) *PullDataByPageRequest {
	s.NextToken = &v
	return s
}

type PullDataByPageResponseBody struct {
	List       []*PullDataByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	MaxResults *int64                            `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s PullDataByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPageResponseBody) GoString() string {
	return s.String()
}

func (s *PullDataByPageResponseBody) SetList(v []*PullDataByPageResponseBodyList) *PullDataByPageResponseBody {
	s.List = v
	return s
}

func (s *PullDataByPageResponseBody) SetMaxResults(v int64) *PullDataByPageResponseBody {
	s.MaxResults = &v
	return s
}

func (s *PullDataByPageResponseBody) SetNextToken(v string) *PullDataByPageResponseBody {
	s.NextToken = &v
	return s
}

type PullDataByPageResponseBodyList struct {
	DataCreateAppId     *string `json:"dataCreateAppId,omitempty" xml:"dataCreateAppId,omitempty"`
	DataCreateAppType   *string `json:"dataCreateAppType,omitempty" xml:"dataCreateAppType,omitempty"`
	DataGmtCreate       *int64  `json:"dataGmtCreate,omitempty" xml:"dataGmtCreate,omitempty"`
	DataGmtModified     *int64  `json:"dataGmtModified,omitempty" xml:"dataGmtModified,omitempty"`
	DataModifiedAppId   *string `json:"dataModifiedAppId,omitempty" xml:"dataModifiedAppId,omitempty"`
	DataModifiedAppType *string `json:"dataModifiedAppType,omitempty" xml:"dataModifiedAppType,omitempty"`
	JsonData            *string `json:"jsonData,omitempty" xml:"jsonData,omitempty"`
}

func (s PullDataByPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *PullDataByPageResponseBodyList) SetDataCreateAppId(v string) *PullDataByPageResponseBodyList {
	s.DataCreateAppId = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetDataCreateAppType(v string) *PullDataByPageResponseBodyList {
	s.DataCreateAppType = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetDataGmtCreate(v int64) *PullDataByPageResponseBodyList {
	s.DataGmtCreate = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetDataGmtModified(v int64) *PullDataByPageResponseBodyList {
	s.DataGmtModified = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetDataModifiedAppId(v string) *PullDataByPageResponseBodyList {
	s.DataModifiedAppId = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetDataModifiedAppType(v string) *PullDataByPageResponseBodyList {
	s.DataModifiedAppType = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetJsonData(v string) *PullDataByPageResponseBodyList {
	s.JsonData = &v
	return s
}

type PullDataByPageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PullDataByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PullDataByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPageResponse) GoString() string {
	return s.String()
}

func (s *PullDataByPageResponse) SetHeaders(v map[string]*string) *PullDataByPageResponse {
	s.Headers = v
	return s
}

func (s *PullDataByPageResponse) SetStatusCode(v int32) *PullDataByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *PullDataByPageResponse) SetBody(v *PullDataByPageResponseBody) *PullDataByPageResponse {
	s.Body = v
	return s
}

type PullDataByPkHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PullDataByPkHeaders) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPkHeaders) GoString() string {
	return s.String()
}

func (s *PullDataByPkHeaders) SetCommonHeaders(v map[string]*string) *PullDataByPkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PullDataByPkHeaders) SetXAcsDingtalkAccessToken(v string) *PullDataByPkHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PullDataByPkRequest struct {
	AppId      *string `json:"appId,omitempty" xml:"appId,omitempty"`
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
}

func (s PullDataByPkRequest) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPkRequest) GoString() string {
	return s.String()
}

func (s *PullDataByPkRequest) SetAppId(v string) *PullDataByPkRequest {
	s.AppId = &v
	return s
}

func (s *PullDataByPkRequest) SetPrimaryKey(v string) *PullDataByPkRequest {
	s.PrimaryKey = &v
	return s
}

type PullDataByPkResponseBody struct {
	DataCreateAppId     *string `json:"dataCreateAppId,omitempty" xml:"dataCreateAppId,omitempty"`
	DataCreateAppType   *string `json:"dataCreateAppType,omitempty" xml:"dataCreateAppType,omitempty"`
	DataGmtCreate       *int64  `json:"dataGmtCreate,omitempty" xml:"dataGmtCreate,omitempty"`
	DataGmtModified     *int64  `json:"dataGmtModified,omitempty" xml:"dataGmtModified,omitempty"`
	DataModifiedAppId   *string `json:"dataModifiedAppId,omitempty" xml:"dataModifiedAppId,omitempty"`
	DataModifiedAppType *string `json:"dataModifiedAppType,omitempty" xml:"dataModifiedAppType,omitempty"`
	JsonData            *string `json:"jsonData,omitempty" xml:"jsonData,omitempty"`
}

func (s PullDataByPkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPkResponseBody) GoString() string {
	return s.String()
}

func (s *PullDataByPkResponseBody) SetDataCreateAppId(v string) *PullDataByPkResponseBody {
	s.DataCreateAppId = &v
	return s
}

func (s *PullDataByPkResponseBody) SetDataCreateAppType(v string) *PullDataByPkResponseBody {
	s.DataCreateAppType = &v
	return s
}

func (s *PullDataByPkResponseBody) SetDataGmtCreate(v int64) *PullDataByPkResponseBody {
	s.DataGmtCreate = &v
	return s
}

func (s *PullDataByPkResponseBody) SetDataGmtModified(v int64) *PullDataByPkResponseBody {
	s.DataGmtModified = &v
	return s
}

func (s *PullDataByPkResponseBody) SetDataModifiedAppId(v string) *PullDataByPkResponseBody {
	s.DataModifiedAppId = &v
	return s
}

func (s *PullDataByPkResponseBody) SetDataModifiedAppType(v string) *PullDataByPkResponseBody {
	s.DataModifiedAppType = &v
	return s
}

func (s *PullDataByPkResponseBody) SetJsonData(v string) *PullDataByPkResponseBody {
	s.JsonData = &v
	return s
}

type PullDataByPkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PullDataByPkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PullDataByPkResponse) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPkResponse) GoString() string {
	return s.String()
}

func (s *PullDataByPkResponse) SetHeaders(v map[string]*string) *PullDataByPkResponse {
	s.Headers = v
	return s
}

func (s *PullDataByPkResponse) SetStatusCode(v int32) *PullDataByPkResponse {
	s.StatusCode = &v
	return s
}

func (s *PullDataByPkResponse) SetBody(v *PullDataByPkResponseBody) *PullDataByPkResponse {
	s.Body = v
	return s
}

type SearchActionsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchActionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchActionsHeaders) GoString() string {
	return s.String()
}

func (s *SearchActionsHeaders) SetCommonHeaders(v map[string]*string) *SearchActionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchActionsHeaders) SetXAcsDingtalkAccessToken(v string) *SearchActionsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchActionsRequest struct {
	ConnectorId             *string   `json:"connectorId,omitempty" xml:"connectorId,omitempty"`
	ConnectorProviderCorpId *string   `json:"connectorProviderCorpId,omitempty" xml:"connectorProviderCorpId,omitempty"`
	IntegrationTypes        []*string `json:"integrationTypes,omitempty" xml:"integrationTypes,omitempty" type:"Repeated"`
	MaxResults              *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken               *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s SearchActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchActionsRequest) GoString() string {
	return s.String()
}

func (s *SearchActionsRequest) SetConnectorId(v string) *SearchActionsRequest {
	s.ConnectorId = &v
	return s
}

func (s *SearchActionsRequest) SetConnectorProviderCorpId(v string) *SearchActionsRequest {
	s.ConnectorProviderCorpId = &v
	return s
}

func (s *SearchActionsRequest) SetIntegrationTypes(v []*string) *SearchActionsRequest {
	s.IntegrationTypes = v
	return s
}

func (s *SearchActionsRequest) SetMaxResults(v int32) *SearchActionsRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchActionsRequest) SetNextToken(v string) *SearchActionsRequest {
	s.NextToken = &v
	return s
}

type SearchActionsResponseBody struct {
	HasMore    *bool                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*SearchActionsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken  *string                          `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *int64                           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchActionsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchActionsResponseBody) SetHasMore(v bool) *SearchActionsResponseBody {
	s.HasMore = &v
	return s
}

func (s *SearchActionsResponseBody) SetList(v []*SearchActionsResponseBodyList) *SearchActionsResponseBody {
	s.List = v
	return s
}

func (s *SearchActionsResponseBody) SetNextToken(v string) *SearchActionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchActionsResponseBody) SetTotalCount(v int64) *SearchActionsResponseBody {
	s.TotalCount = &v
	return s
}

type SearchActionsResponseBodyList struct {
	AuthorityUrl    *string `json:"authorityUrl,omitempty" xml:"authorityUrl,omitempty"`
	Authorized      *bool   `json:"authorized,omitempty" xml:"authorized,omitempty"`
	ConnectAssetUri *string `json:"connectAssetUri,omitempty" xml:"connectAssetUri,omitempty"`
	ConnectorId     *string `json:"connectorId,omitempty" xml:"connectorId,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon            *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Id              *string `json:"id,omitempty" xml:"id,omitempty"`
	IntegrationType *string `json:"integrationType,omitempty" xml:"integrationType,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	ProviderCorpId  *string `json:"providerCorpId,omitempty" xml:"providerCorpId,omitempty"`
}

func (s SearchActionsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s SearchActionsResponseBodyList) GoString() string {
	return s.String()
}

func (s *SearchActionsResponseBodyList) SetAuthorityUrl(v string) *SearchActionsResponseBodyList {
	s.AuthorityUrl = &v
	return s
}

func (s *SearchActionsResponseBodyList) SetAuthorized(v bool) *SearchActionsResponseBodyList {
	s.Authorized = &v
	return s
}

func (s *SearchActionsResponseBodyList) SetConnectAssetUri(v string) *SearchActionsResponseBodyList {
	s.ConnectAssetUri = &v
	return s
}

func (s *SearchActionsResponseBodyList) SetConnectorId(v string) *SearchActionsResponseBodyList {
	s.ConnectorId = &v
	return s
}

func (s *SearchActionsResponseBodyList) SetDescription(v string) *SearchActionsResponseBodyList {
	s.Description = &v
	return s
}

func (s *SearchActionsResponseBodyList) SetIcon(v string) *SearchActionsResponseBodyList {
	s.Icon = &v
	return s
}

func (s *SearchActionsResponseBodyList) SetId(v string) *SearchActionsResponseBodyList {
	s.Id = &v
	return s
}

func (s *SearchActionsResponseBodyList) SetIntegrationType(v string) *SearchActionsResponseBodyList {
	s.IntegrationType = &v
	return s
}

func (s *SearchActionsResponseBodyList) SetName(v string) *SearchActionsResponseBodyList {
	s.Name = &v
	return s
}

func (s *SearchActionsResponseBodyList) SetProviderCorpId(v string) *SearchActionsResponseBodyList {
	s.ProviderCorpId = &v
	return s
}

type SearchActionsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchActionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchActionsResponse) GoString() string {
	return s.String()
}

func (s *SearchActionsResponse) SetHeaders(v map[string]*string) *SearchActionsResponse {
	s.Headers = v
	return s
}

func (s *SearchActionsResponse) SetStatusCode(v int32) *SearchActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchActionsResponse) SetBody(v *SearchActionsResponseBody) *SearchActionsResponse {
	s.Body = v
	return s
}

type SearchConnectorsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchConnectorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchConnectorsHeaders) GoString() string {
	return s.String()
}

func (s *SearchConnectorsHeaders) SetCommonHeaders(v map[string]*string) *SearchConnectorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchConnectorsHeaders) SetXAcsDingtalkAccessToken(v string) *SearchConnectorsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchConnectorsRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchConnectorsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchConnectorsRequest) GoString() string {
	return s.String()
}

func (s *SearchConnectorsRequest) SetMaxResults(v int32) *SearchConnectorsRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchConnectorsRequest) SetNextToken(v string) *SearchConnectorsRequest {
	s.NextToken = &v
	return s
}

func (s *SearchConnectorsRequest) SetType(v string) *SearchConnectorsRequest {
	s.Type = &v
	return s
}

type SearchConnectorsResponseBody struct {
	HasMore    *bool                               `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*SearchConnectorsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken  *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *string                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchConnectorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchConnectorsResponseBody) SetHasMore(v bool) *SearchConnectorsResponseBody {
	s.HasMore = &v
	return s
}

func (s *SearchConnectorsResponseBody) SetList(v []*SearchConnectorsResponseBodyList) *SearchConnectorsResponseBody {
	s.List = v
	return s
}

func (s *SearchConnectorsResponseBody) SetNextToken(v string) *SearchConnectorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchConnectorsResponseBody) SetTotalCount(v string) *SearchConnectorsResponseBody {
	s.TotalCount = &v
	return s
}

type SearchConnectorsResponseBodyList struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon           *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Id             *string `json:"id,omitempty" xml:"id,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ProviderCorpId *string `json:"providerCorpId,omitempty" xml:"providerCorpId,omitempty"`
}

func (s SearchConnectorsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s SearchConnectorsResponseBodyList) GoString() string {
	return s.String()
}

func (s *SearchConnectorsResponseBodyList) SetDescription(v string) *SearchConnectorsResponseBodyList {
	s.Description = &v
	return s
}

func (s *SearchConnectorsResponseBodyList) SetIcon(v string) *SearchConnectorsResponseBodyList {
	s.Icon = &v
	return s
}

func (s *SearchConnectorsResponseBodyList) SetId(v string) *SearchConnectorsResponseBodyList {
	s.Id = &v
	return s
}

func (s *SearchConnectorsResponseBodyList) SetName(v string) *SearchConnectorsResponseBodyList {
	s.Name = &v
	return s
}

func (s *SearchConnectorsResponseBodyList) SetProviderCorpId(v string) *SearchConnectorsResponseBodyList {
	s.ProviderCorpId = &v
	return s
}

type SearchConnectorsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchConnectorsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchConnectorsResponse) GoString() string {
	return s.String()
}

func (s *SearchConnectorsResponse) SetHeaders(v map[string]*string) *SearchConnectorsResponse {
	s.Headers = v
	return s
}

func (s *SearchConnectorsResponse) SetStatusCode(v int32) *SearchConnectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchConnectorsResponse) SetBody(v *SearchConnectorsResponseBody) *SearchConnectorsResponse {
	s.Body = v
	return s
}

type SyncDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncDataHeaders) GoString() string {
	return s.String()
}

func (s *SyncDataHeaders) SetCommonHeaders(v map[string]*string) *SyncDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncDataHeaders) SetXAcsDingtalkAccessToken(v string) *SyncDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncDataRequest struct {
	AppId           *string                           `json:"appId,omitempty" xml:"appId,omitempty"`
	TriggerDataList []*SyncDataRequestTriggerDataList `json:"triggerDataList,omitempty" xml:"triggerDataList,omitempty" type:"Repeated"`
}

func (s SyncDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncDataRequest) GoString() string {
	return s.String()
}

func (s *SyncDataRequest) SetAppId(v string) *SyncDataRequest {
	s.AppId = &v
	return s
}

func (s *SyncDataRequest) SetTriggerDataList(v []*SyncDataRequestTriggerDataList) *SyncDataRequest {
	s.TriggerDataList = v
	return s
}

type SyncDataRequestTriggerDataList struct {
	Action            *string `json:"action,omitempty" xml:"action,omitempty"`
	CustomTriggerId   *string `json:"customTriggerId,omitempty" xml:"customTriggerId,omitempty"`
	DataGmtCreate     *int64  `json:"dataGmtCreate,omitempty" xml:"dataGmtCreate,omitempty"`
	DataGmtModified   *int64  `json:"dataGmtModified,omitempty" xml:"dataGmtModified,omitempty"`
	IntegrationObject *string `json:"integrationObject,omitempty" xml:"integrationObject,omitempty"`
	JsonData          *string `json:"jsonData,omitempty" xml:"jsonData,omitempty"`
	TriggerCondition  *string `json:"triggerCondition,omitempty" xml:"triggerCondition,omitempty"`
	TriggerId         *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
}

func (s SyncDataRequestTriggerDataList) String() string {
	return tea.Prettify(s)
}

func (s SyncDataRequestTriggerDataList) GoString() string {
	return s.String()
}

func (s *SyncDataRequestTriggerDataList) SetAction(v string) *SyncDataRequestTriggerDataList {
	s.Action = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetCustomTriggerId(v string) *SyncDataRequestTriggerDataList {
	s.CustomTriggerId = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetDataGmtCreate(v int64) *SyncDataRequestTriggerDataList {
	s.DataGmtCreate = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetDataGmtModified(v int64) *SyncDataRequestTriggerDataList {
	s.DataGmtModified = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetIntegrationObject(v string) *SyncDataRequestTriggerDataList {
	s.IntegrationObject = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetJsonData(v string) *SyncDataRequestTriggerDataList {
	s.JsonData = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetTriggerCondition(v string) *SyncDataRequestTriggerDataList {
	s.TriggerCondition = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetTriggerId(v string) *SyncDataRequestTriggerDataList {
	s.TriggerId = &v
	return s
}

type SyncDataResponseBody struct {
	List []*SyncDataResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s SyncDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDataResponseBody) SetList(v []*SyncDataResponseBodyList) *SyncDataResponseBody {
	s.List = v
	return s
}

type SyncDataResponseBodyList struct {
	BizPrimaryKey *string `json:"bizPrimaryKey,omitempty" xml:"bizPrimaryKey,omitempty"`
	SubErrCode    *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	SubErrMsg     *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	Success       *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TriggerId     *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
}

func (s SyncDataResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponseBodyList) GoString() string {
	return s.String()
}

func (s *SyncDataResponseBodyList) SetBizPrimaryKey(v string) *SyncDataResponseBodyList {
	s.BizPrimaryKey = &v
	return s
}

func (s *SyncDataResponseBodyList) SetSubErrCode(v string) *SyncDataResponseBodyList {
	s.SubErrCode = &v
	return s
}

func (s *SyncDataResponseBodyList) SetSubErrMsg(v string) *SyncDataResponseBodyList {
	s.SubErrMsg = &v
	return s
}

func (s *SyncDataResponseBodyList) SetSuccess(v bool) *SyncDataResponseBodyList {
	s.Success = &v
	return s
}

func (s *SyncDataResponseBodyList) SetTriggerId(v string) *SyncDataResponseBodyList {
	s.TriggerId = &v
	return s
}

type SyncDataResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponse) GoString() string {
	return s.String()
}

func (s *SyncDataResponse) SetHeaders(v map[string]*string) *SyncDataResponse {
	s.Headers = v
	return s
}

func (s *SyncDataResponse) SetStatusCode(v int32) *SyncDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDataResponse) SetBody(v *SyncDataResponseBody) *SyncDataResponse {
	s.Body = v
	return s
}

type UpdateActionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateActionHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionHeaders) GoString() string {
	return s.String()
}

func (s *UpdateActionHeaders) SetCommonHeaders(v map[string]*string) *UpdateActionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateActionHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateActionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateActionRequest struct {
	ActionInfo     []*UpdateActionRequestActionInfo `json:"actionInfo,omitempty" xml:"actionInfo,omitempty" type:"Repeated"`
	IntegratorFlag *string                          `json:"integratorFlag,omitempty" xml:"integratorFlag,omitempty"`
}

func (s UpdateActionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionRequest) GoString() string {
	return s.String()
}

func (s *UpdateActionRequest) SetActionInfo(v []*UpdateActionRequestActionInfo) *UpdateActionRequest {
	s.ActionInfo = v
	return s
}

func (s *UpdateActionRequest) SetIntegratorFlag(v string) *UpdateActionRequest {
	s.IntegratorFlag = &v
	return s
}

type UpdateActionRequestActionInfo struct {
	ApiPath               *string                                           `json:"apiPath,omitempty" xml:"apiPath,omitempty"`
	Description           *string                                           `json:"description,omitempty" xml:"description,omitempty"`
	DingActionId          *string                                           `json:"dingActionId,omitempty" xml:"dingActionId,omitempty"`
	DingConnectorId       *string                                           `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	InputMappingConfig    *UpdateActionRequestActionInfoInputMappingConfig  `json:"inputMappingConfig,omitempty" xml:"inputMappingConfig,omitempty" type:"Struct"`
	InputSchema           *string                                           `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	IntegratorActionId    *string                                           `json:"integratorActionId,omitempty" xml:"integratorActionId,omitempty"`
	IntegratorConnectorId *string                                           `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	Name                  *string                                           `json:"name,omitempty" xml:"name,omitempty"`
	OutputDataRules       []*UpdateActionRequestActionInfoOutputDataRules   `json:"outputDataRules,omitempty" xml:"outputDataRules,omitempty" type:"Repeated"`
	OutputMappingConfig   *UpdateActionRequestActionInfoOutputMappingConfig `json:"outputMappingConfig,omitempty" xml:"outputMappingConfig,omitempty" type:"Struct"`
	OutputSchema          *string                                           `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
}

func (s UpdateActionRequestActionInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionRequestActionInfo) GoString() string {
	return s.String()
}

func (s *UpdateActionRequestActionInfo) SetApiPath(v string) *UpdateActionRequestActionInfo {
	s.ApiPath = &v
	return s
}

func (s *UpdateActionRequestActionInfo) SetDescription(v string) *UpdateActionRequestActionInfo {
	s.Description = &v
	return s
}

func (s *UpdateActionRequestActionInfo) SetDingActionId(v string) *UpdateActionRequestActionInfo {
	s.DingActionId = &v
	return s
}

func (s *UpdateActionRequestActionInfo) SetDingConnectorId(v string) *UpdateActionRequestActionInfo {
	s.DingConnectorId = &v
	return s
}

func (s *UpdateActionRequestActionInfo) SetInputMappingConfig(v *UpdateActionRequestActionInfoInputMappingConfig) *UpdateActionRequestActionInfo {
	s.InputMappingConfig = v
	return s
}

func (s *UpdateActionRequestActionInfo) SetInputSchema(v string) *UpdateActionRequestActionInfo {
	s.InputSchema = &v
	return s
}

func (s *UpdateActionRequestActionInfo) SetIntegratorActionId(v string) *UpdateActionRequestActionInfo {
	s.IntegratorActionId = &v
	return s
}

func (s *UpdateActionRequestActionInfo) SetIntegratorConnectorId(v string) *UpdateActionRequestActionInfo {
	s.IntegratorConnectorId = &v
	return s
}

func (s *UpdateActionRequestActionInfo) SetName(v string) *UpdateActionRequestActionInfo {
	s.Name = &v
	return s
}

func (s *UpdateActionRequestActionInfo) SetOutputDataRules(v []*UpdateActionRequestActionInfoOutputDataRules) *UpdateActionRequestActionInfo {
	s.OutputDataRules = v
	return s
}

func (s *UpdateActionRequestActionInfo) SetOutputMappingConfig(v *UpdateActionRequestActionInfoOutputMappingConfig) *UpdateActionRequestActionInfo {
	s.OutputMappingConfig = v
	return s
}

func (s *UpdateActionRequestActionInfo) SetOutputSchema(v string) *UpdateActionRequestActionInfo {
	s.OutputSchema = &v
	return s
}

type UpdateActionRequestActionInfoInputMappingConfig struct {
	CustomSchemaMapping *string `json:"customSchemaMapping,omitempty" xml:"customSchemaMapping,omitempty"`
	Rules               *string `json:"rules,omitempty" xml:"rules,omitempty"`
}

func (s UpdateActionRequestActionInfoInputMappingConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionRequestActionInfoInputMappingConfig) GoString() string {
	return s.String()
}

func (s *UpdateActionRequestActionInfoInputMappingConfig) SetCustomSchemaMapping(v string) *UpdateActionRequestActionInfoInputMappingConfig {
	s.CustomSchemaMapping = &v
	return s
}

func (s *UpdateActionRequestActionInfoInputMappingConfig) SetRules(v string) *UpdateActionRequestActionInfoInputMappingConfig {
	s.Rules = &v
	return s
}

type UpdateActionRequestActionInfoOutputDataRules struct {
	ExpectValue  *string `json:"expectValue,omitempty" xml:"expectValue,omitempty"`
	Operate      *string `json:"operate,omitempty" xml:"operate,omitempty"`
	PropertyPath *string `json:"propertyPath,omitempty" xml:"propertyPath,omitempty"`
}

func (s UpdateActionRequestActionInfoOutputDataRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionRequestActionInfoOutputDataRules) GoString() string {
	return s.String()
}

func (s *UpdateActionRequestActionInfoOutputDataRules) SetExpectValue(v string) *UpdateActionRequestActionInfoOutputDataRules {
	s.ExpectValue = &v
	return s
}

func (s *UpdateActionRequestActionInfoOutputDataRules) SetOperate(v string) *UpdateActionRequestActionInfoOutputDataRules {
	s.Operate = &v
	return s
}

func (s *UpdateActionRequestActionInfoOutputDataRules) SetPropertyPath(v string) *UpdateActionRequestActionInfoOutputDataRules {
	s.PropertyPath = &v
	return s
}

type UpdateActionRequestActionInfoOutputMappingConfig struct {
	CustomSchemaMapping *string `json:"customSchemaMapping,omitempty" xml:"customSchemaMapping,omitempty"`
	Rules               *string `json:"rules,omitempty" xml:"rules,omitempty"`
}

func (s UpdateActionRequestActionInfoOutputMappingConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionRequestActionInfoOutputMappingConfig) GoString() string {
	return s.String()
}

func (s *UpdateActionRequestActionInfoOutputMappingConfig) SetCustomSchemaMapping(v string) *UpdateActionRequestActionInfoOutputMappingConfig {
	s.CustomSchemaMapping = &v
	return s
}

func (s *UpdateActionRequestActionInfoOutputMappingConfig) SetRules(v string) *UpdateActionRequestActionInfoOutputMappingConfig {
	s.Rules = &v
	return s
}

type UpdateActionResponseBody struct {
	Item []*UpdateActionResponseBodyItem `json:"item,omitempty" xml:"item,omitempty" type:"Repeated"`
}

func (s UpdateActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateActionResponseBody) SetItem(v []*UpdateActionResponseBodyItem) *UpdateActionResponseBody {
	s.Item = v
	return s
}

type UpdateActionResponseBodyItem struct {
	DingActionId          *string `json:"dingActionId,omitempty" xml:"dingActionId,omitempty"`
	DingConnectorId       *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	IntegratorActionId    *string `json:"integratorActionId,omitempty" xml:"integratorActionId,omitempty"`
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	SubErrCode            *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	SubErrMsg             *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	Success               *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateActionResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionResponseBodyItem) GoString() string {
	return s.String()
}

func (s *UpdateActionResponseBodyItem) SetDingActionId(v string) *UpdateActionResponseBodyItem {
	s.DingActionId = &v
	return s
}

func (s *UpdateActionResponseBodyItem) SetDingConnectorId(v string) *UpdateActionResponseBodyItem {
	s.DingConnectorId = &v
	return s
}

func (s *UpdateActionResponseBodyItem) SetIntegratorActionId(v string) *UpdateActionResponseBodyItem {
	s.IntegratorActionId = &v
	return s
}

func (s *UpdateActionResponseBodyItem) SetIntegratorConnectorId(v string) *UpdateActionResponseBodyItem {
	s.IntegratorConnectorId = &v
	return s
}

func (s *UpdateActionResponseBodyItem) SetSubErrCode(v string) *UpdateActionResponseBodyItem {
	s.SubErrCode = &v
	return s
}

func (s *UpdateActionResponseBodyItem) SetSubErrMsg(v string) *UpdateActionResponseBodyItem {
	s.SubErrMsg = &v
	return s
}

func (s *UpdateActionResponseBodyItem) SetSuccess(v string) *UpdateActionResponseBodyItem {
	s.Success = &v
	return s
}

type UpdateActionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateActionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionResponse) GoString() string {
	return s.String()
}

func (s *UpdateActionResponse) SetHeaders(v map[string]*string) *UpdateActionResponse {
	s.Headers = v
	return s
}

func (s *UpdateActionResponse) SetStatusCode(v int32) *UpdateActionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateActionResponse) SetBody(v *UpdateActionResponseBody) *UpdateActionResponse {
	s.Body = v
	return s
}

type UpdateConnectorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateConnectorHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectorHeaders) GoString() string {
	return s.String()
}

func (s *UpdateConnectorHeaders) SetCommonHeaders(v map[string]*string) *UpdateConnectorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateConnectorHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateConnectorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateConnectorRequest struct {
	ConnectorInfo  []*UpdateConnectorRequestConnectorInfo `json:"connectorInfo,omitempty" xml:"connectorInfo,omitempty" type:"Repeated"`
	IntegratorFlag *string                                `json:"integratorFlag,omitempty" xml:"integratorFlag,omitempty"`
}

func (s UpdateConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectorRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectorRequest) SetConnectorInfo(v []*UpdateConnectorRequestConnectorInfo) *UpdateConnectorRequest {
	s.ConnectorInfo = v
	return s
}

func (s *UpdateConnectorRequest) SetIntegratorFlag(v string) *UpdateConnectorRequest {
	s.IntegratorFlag = &v
	return s
}

type UpdateConnectorRequestConnectorInfo struct {
	ApiDomain             *string `json:"apiDomain,omitempty" xml:"apiDomain,omitempty"`
	ApiSecret             *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	AppId                 *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
	AuthValueEnv          *bool   `json:"authValueEnv,omitempty" xml:"authValueEnv,omitempty"`
	Description           *string `json:"description,omitempty" xml:"description,omitempty"`
	DingConnectorId       *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	DomainEnv             *bool   `json:"domainEnv,omitempty" xml:"domainEnv,omitempty"`
	IconMediaId           *string `json:"iconMediaId,omitempty" xml:"iconMediaId,omitempty"`
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateConnectorRequestConnectorInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectorRequestConnectorInfo) GoString() string {
	return s.String()
}

func (s *UpdateConnectorRequestConnectorInfo) SetApiDomain(v string) *UpdateConnectorRequestConnectorInfo {
	s.ApiDomain = &v
	return s
}

func (s *UpdateConnectorRequestConnectorInfo) SetApiSecret(v string) *UpdateConnectorRequestConnectorInfo {
	s.ApiSecret = &v
	return s
}

func (s *UpdateConnectorRequestConnectorInfo) SetAppId(v int64) *UpdateConnectorRequestConnectorInfo {
	s.AppId = &v
	return s
}

func (s *UpdateConnectorRequestConnectorInfo) SetAuthValueEnv(v bool) *UpdateConnectorRequestConnectorInfo {
	s.AuthValueEnv = &v
	return s
}

func (s *UpdateConnectorRequestConnectorInfo) SetDescription(v string) *UpdateConnectorRequestConnectorInfo {
	s.Description = &v
	return s
}

func (s *UpdateConnectorRequestConnectorInfo) SetDingConnectorId(v string) *UpdateConnectorRequestConnectorInfo {
	s.DingConnectorId = &v
	return s
}

func (s *UpdateConnectorRequestConnectorInfo) SetDomainEnv(v bool) *UpdateConnectorRequestConnectorInfo {
	s.DomainEnv = &v
	return s
}

func (s *UpdateConnectorRequestConnectorInfo) SetIconMediaId(v string) *UpdateConnectorRequestConnectorInfo {
	s.IconMediaId = &v
	return s
}

func (s *UpdateConnectorRequestConnectorInfo) SetIntegratorConnectorId(v string) *UpdateConnectorRequestConnectorInfo {
	s.IntegratorConnectorId = &v
	return s
}

func (s *UpdateConnectorRequestConnectorInfo) SetName(v string) *UpdateConnectorRequestConnectorInfo {
	s.Name = &v
	return s
}

type UpdateConnectorResponseBody struct {
	Item []*UpdateConnectorResponseBodyItem `json:"item,omitempty" xml:"item,omitempty" type:"Repeated"`
}

func (s UpdateConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponseBody) SetItem(v []*UpdateConnectorResponseBodyItem) *UpdateConnectorResponseBody {
	s.Item = v
	return s
}

type UpdateConnectorResponseBodyItem struct {
	DingConnectorId       *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	SubErrCode            *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	SubErrMsg             *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	Success               *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateConnectorResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectorResponseBodyItem) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponseBodyItem) SetDingConnectorId(v string) *UpdateConnectorResponseBodyItem {
	s.DingConnectorId = &v
	return s
}

func (s *UpdateConnectorResponseBodyItem) SetIntegratorConnectorId(v string) *UpdateConnectorResponseBodyItem {
	s.IntegratorConnectorId = &v
	return s
}

func (s *UpdateConnectorResponseBodyItem) SetSubErrCode(v string) *UpdateConnectorResponseBodyItem {
	s.SubErrCode = &v
	return s
}

func (s *UpdateConnectorResponseBodyItem) SetSubErrMsg(v string) *UpdateConnectorResponseBodyItem {
	s.SubErrMsg = &v
	return s
}

func (s *UpdateConnectorResponseBodyItem) SetSuccess(v bool) *UpdateConnectorResponseBodyItem {
	s.Success = &v
	return s
}

type UpdateConnectorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectorResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponse) SetHeaders(v map[string]*string) *UpdateConnectorResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnectorResponse) SetStatusCode(v int32) *UpdateConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConnectorResponse) SetBody(v *UpdateConnectorResponseBody) *UpdateConnectorResponse {
	s.Body = v
	return s
}

type UpdateTriggerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTriggerHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTriggerHeaders) SetCommonHeaders(v map[string]*string) *UpdateTriggerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTriggerHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTriggerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTriggerRequest struct {
	IntegratorFlag *string                            `json:"integratorFlag,omitempty" xml:"integratorFlag,omitempty"`
	TriggerInfo    []*UpdateTriggerRequestTriggerInfo `json:"triggerInfo,omitempty" xml:"triggerInfo,omitempty" type:"Repeated"`
}

func (s UpdateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequest) SetIntegratorFlag(v string) *UpdateTriggerRequest {
	s.IntegratorFlag = &v
	return s
}

func (s *UpdateTriggerRequest) SetTriggerInfo(v []*UpdateTriggerRequestTriggerInfo) *UpdateTriggerRequest {
	s.TriggerInfo = v
	return s
}

type UpdateTriggerRequestTriggerInfo struct {
	Description           *string `json:"description,omitempty" xml:"description,omitempty"`
	DingConnectorId       *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	DingTriggerId         *string `json:"dingTriggerId,omitempty" xml:"dingTriggerId,omitempty"`
	InputSchema           *string `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	IntegratorTriggerId   *string `json:"integratorTriggerId,omitempty" xml:"integratorTriggerId,omitempty"`
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateTriggerRequestTriggerInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerRequestTriggerInfo) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequestTriggerInfo) SetDescription(v string) *UpdateTriggerRequestTriggerInfo {
	s.Description = &v
	return s
}

func (s *UpdateTriggerRequestTriggerInfo) SetDingConnectorId(v string) *UpdateTriggerRequestTriggerInfo {
	s.DingConnectorId = &v
	return s
}

func (s *UpdateTriggerRequestTriggerInfo) SetDingTriggerId(v string) *UpdateTriggerRequestTriggerInfo {
	s.DingTriggerId = &v
	return s
}

func (s *UpdateTriggerRequestTriggerInfo) SetInputSchema(v string) *UpdateTriggerRequestTriggerInfo {
	s.InputSchema = &v
	return s
}

func (s *UpdateTriggerRequestTriggerInfo) SetIntegratorConnectorId(v string) *UpdateTriggerRequestTriggerInfo {
	s.IntegratorConnectorId = &v
	return s
}

func (s *UpdateTriggerRequestTriggerInfo) SetIntegratorTriggerId(v string) *UpdateTriggerRequestTriggerInfo {
	s.IntegratorTriggerId = &v
	return s
}

func (s *UpdateTriggerRequestTriggerInfo) SetName(v string) *UpdateTriggerRequestTriggerInfo {
	s.Name = &v
	return s
}

type UpdateTriggerResponseBody struct {
	Item []*UpdateTriggerResponseBodyItem `json:"item,omitempty" xml:"item,omitempty" type:"Repeated"`
}

func (s UpdateTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponseBody) SetItem(v []*UpdateTriggerResponseBodyItem) *UpdateTriggerResponseBody {
	s.Item = v
	return s
}

type UpdateTriggerResponseBodyItem struct {
	DingConnectorId       *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	DingTriggerId         *string `json:"dingTriggerId,omitempty" xml:"dingTriggerId,omitempty"`
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	IntegratorTriggerId   *string `json:"integratorTriggerId,omitempty" xml:"integratorTriggerId,omitempty"`
	SubErrCode            *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	SubErrMsg             *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	Success               *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateTriggerResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerResponseBodyItem) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponseBodyItem) SetDingConnectorId(v string) *UpdateTriggerResponseBodyItem {
	s.DingConnectorId = &v
	return s
}

func (s *UpdateTriggerResponseBodyItem) SetDingTriggerId(v string) *UpdateTriggerResponseBodyItem {
	s.DingTriggerId = &v
	return s
}

func (s *UpdateTriggerResponseBodyItem) SetIntegratorConnectorId(v string) *UpdateTriggerResponseBodyItem {
	s.IntegratorConnectorId = &v
	return s
}

func (s *UpdateTriggerResponseBodyItem) SetIntegratorTriggerId(v string) *UpdateTriggerResponseBodyItem {
	s.IntegratorTriggerId = &v
	return s
}

func (s *UpdateTriggerResponseBodyItem) SetSubErrCode(v string) *UpdateTriggerResponseBodyItem {
	s.SubErrCode = &v
	return s
}

func (s *UpdateTriggerResponseBodyItem) SetSubErrMsg(v string) *UpdateTriggerResponseBodyItem {
	s.SubErrMsg = &v
	return s
}

func (s *UpdateTriggerResponseBodyItem) SetSuccess(v bool) *UpdateTriggerResponseBodyItem {
	s.Success = &v
	return s
}

type UpdateTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponse) SetHeaders(v map[string]*string) *UpdateTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTriggerResponse) SetStatusCode(v int32) *UpdateTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTriggerResponse) SetBody(v *UpdateTriggerResponseBody) *UpdateTriggerResponse {
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

func (client *Client) CreateActionWithOptions(request *CreateActionRequest, headers *CreateActionHeaders, runtime *util.RuntimeOptions) (_result *CreateActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionInfo)) {
		body["actionInfo"] = request.ActionInfo
	}

	if !tea.BoolValue(util.IsUnset(request.IntegratorFlag)) {
		body["integratorFlag"] = request.IntegratorFlag
	}

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
		Action:      tea.String("CreateAction"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/actions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateActionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAction(request *CreateActionRequest) (_result *CreateActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateActionHeaders{}
	_result = &CreateActionResponse{}
	_body, _err := client.CreateActionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConnectorWithOptions(request *CreateConnectorRequest, headers *CreateConnectorHeaders, runtime *util.RuntimeOptions) (_result *CreateConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectorInfo)) {
		body["connectorInfo"] = request.ConnectorInfo
	}

	if !tea.BoolValue(util.IsUnset(request.IntegratorFlag)) {
		body["integratorFlag"] = request.IntegratorFlag
	}

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
		Action:      tea.String("CreateConnector"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/connectors"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConnectorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConnector(request *CreateConnectorRequest) (_result *CreateConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateConnectorHeaders{}
	_result = &CreateConnectorResponse{}
	_body, _err := client.CreateConnectorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInvocableInstanceWithOptions(request *CreateInvocableInstanceRequest, headers *CreateInvocableInstanceHeaders, runtime *util.RuntimeOptions) (_result *CreateInvocableInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectAssetUri)) {
		body["connectAssetUri"] = request.ConnectAssetUri
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceKey)) {
		body["instanceKey"] = request.InstanceKey
	}

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
		Action:      tea.String("CreateInvocableInstance"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInvocableInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInvocableInstance(request *CreateInvocableInstanceRequest) (_result *CreateInvocableInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateInvocableInstanceHeaders{}
	_result = &CreateInvocableInstanceResponse{}
	_body, _err := client.CreateInvocableInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTriggerWithOptions(request *CreateTriggerRequest, headers *CreateTriggerHeaders, runtime *util.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IntegratorFlag)) {
		body["integratorFlag"] = request.IntegratorFlag
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerInfo)) {
		body["triggerInfo"] = request.TriggerInfo
	}

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
		Action:      tea.String("CreateTrigger"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/triggers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTriggerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrigger(request *CreateTriggerRequest) (_result *CreateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTriggerHeaders{}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CreateTriggerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetActionDetailWithOptions(request *GetActionDetailRequest, headers *GetActionDetailHeaders, runtime *util.RuntimeOptions) (_result *GetActionDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectAssetUri)) {
		body["connectAssetUri"] = request.ConnectAssetUri
	}

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
		Action:      tea.String("GetActionDetail"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/assets/actions/details/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetActionDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetActionDetail(request *GetActionDetailRequest) (_result *GetActionDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetActionDetailHeaders{}
	_result = &GetActionDetailResponse{}
	_body, _err := client.GetActionDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeInstanceWithOptions(request *InvokeInstanceRequest, headers *InvokeInstanceHeaders, runtime *util.RuntimeOptions) (_result *InvokeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectAssetUri)) {
		body["connectAssetUri"] = request.ConnectAssetUri
	}

	if !tea.BoolValue(util.IsUnset(request.InputJsonString)) {
		body["inputJsonString"] = request.InputJsonString
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceKey)) {
		body["instanceKey"] = request.InstanceKey
	}

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
		Action:      tea.String("InvokeInstance"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/instances/invoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeInstance(request *InvokeInstanceRequest) (_result *InvokeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvokeInstanceHeaders{}
	_result = &InvokeInstanceResponse{}
	_body, _err := client.InvokeInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullDataByPageWithOptions(request *PullDataByPageRequest, headers *PullDataByPageHeaders, runtime *util.RuntimeOptions) (_result *PullDataByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DataModelId)) {
		query["dataModelId"] = request.DataModelId
	}

	if !tea.BoolValue(util.IsUnset(request.DatetimeFilterField)) {
		query["datetimeFilterField"] = request.DatetimeFilterField
	}

	if !tea.BoolValue(util.IsUnset(request.MaxDatetime)) {
		query["maxDatetime"] = request.MaxDatetime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.MinDatetime)) {
		query["minDatetime"] = request.MinDatetime
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
		Action:      tea.String("PullDataByPage"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/data"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PullDataByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PullDataByPage(request *PullDataByPageRequest) (_result *PullDataByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PullDataByPageHeaders{}
	_result = &PullDataByPageResponse{}
	_body, _err := client.PullDataByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullDataByPkWithOptions(dataModelId *string, request *PullDataByPkRequest, headers *PullDataByPkHeaders, runtime *util.RuntimeOptions) (_result *PullDataByPkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryKey)) {
		query["primaryKey"] = request.PrimaryKey
	}

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
		Action:      tea.String("PullDataByPk"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/data/" + tea.StringValue(dataModelId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PullDataByPkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PullDataByPk(dataModelId *string, request *PullDataByPkRequest) (_result *PullDataByPkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PullDataByPkHeaders{}
	_result = &PullDataByPkResponse{}
	_body, _err := client.PullDataByPkWithOptions(dataModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchActionsWithOptions(request *SearchActionsRequest, headers *SearchActionsHeaders, runtime *util.RuntimeOptions) (_result *SearchActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectorId)) {
		body["connectorId"] = request.ConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectorProviderCorpId)) {
		body["connectorProviderCorpId"] = request.ConnectorProviderCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationTypes)) {
		body["integrationTypes"] = request.IntegrationTypes
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
		Action:      tea.String("SearchActions"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/assets/actions/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchActionsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchActions(request *SearchActionsRequest) (_result *SearchActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchActionsHeaders{}
	_result = &SearchActionsResponse{}
	_body, _err := client.SearchActionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchConnectorsWithOptions(request *SearchConnectorsRequest, headers *SearchConnectorsHeaders, runtime *util.RuntimeOptions) (_result *SearchConnectorsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

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
		Action:      tea.String("SearchConnectors"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/assets/connectors"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchConnectorsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchConnectors(request *SearchConnectorsRequest) (_result *SearchConnectorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchConnectorsHeaders{}
	_result = &SearchConnectorsResponse{}
	_body, _err := client.SearchConnectorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncDataWithOptions(request *SyncDataRequest, headers *SyncDataHeaders, runtime *util.RuntimeOptions) (_result *SyncDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerDataList)) {
		body["triggerDataList"] = request.TriggerDataList
	}

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
		Action:      tea.String("SyncData"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/triggers/data/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncData(request *SyncDataRequest) (_result *SyncDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncDataHeaders{}
	_result = &SyncDataResponse{}
	_body, _err := client.SyncDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateActionWithOptions(request *UpdateActionRequest, headers *UpdateActionHeaders, runtime *util.RuntimeOptions) (_result *UpdateActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionInfo)) {
		body["actionInfo"] = request.ActionInfo
	}

	if !tea.BoolValue(util.IsUnset(request.IntegratorFlag)) {
		body["integratorFlag"] = request.IntegratorFlag
	}

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
		Action:      tea.String("UpdateAction"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/actions"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateActionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAction(request *UpdateActionRequest) (_result *UpdateActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateActionHeaders{}
	_result = &UpdateActionResponse{}
	_body, _err := client.UpdateActionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConnectorWithOptions(request *UpdateConnectorRequest, headers *UpdateConnectorHeaders, runtime *util.RuntimeOptions) (_result *UpdateConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectorInfo)) {
		body["connectorInfo"] = request.ConnectorInfo
	}

	if !tea.BoolValue(util.IsUnset(request.IntegratorFlag)) {
		body["integratorFlag"] = request.IntegratorFlag
	}

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
		Action:      tea.String("UpdateConnector"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/connectors"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConnectorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConnector(request *UpdateConnectorRequest) (_result *UpdateConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateConnectorHeaders{}
	_result = &UpdateConnectorResponse{}
	_body, _err := client.UpdateConnectorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTriggerWithOptions(request *UpdateTriggerRequest, headers *UpdateTriggerHeaders, runtime *util.RuntimeOptions) (_result *UpdateTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IntegratorFlag)) {
		body["integratorFlag"] = request.IntegratorFlag
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerInfo)) {
		body["triggerInfo"] = request.TriggerInfo
	}

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
		Action:      tea.String("UpdateTrigger"),
		Version:     tea.String("connector_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/connector/triggers"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTrigger(request *UpdateTriggerRequest) (_result *UpdateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTriggerHeaders{}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.UpdateTriggerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
