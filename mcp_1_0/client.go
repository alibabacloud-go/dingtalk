// This file is auto-generated, don't edit it. Thanks.
package mcp_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateMcpHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ActivateMcpHeaders) String() string {
	return tea.Prettify(s)
}

func (s ActivateMcpHeaders) GoString() string {
	return s.String()
}

func (s *ActivateMcpHeaders) SetCommonHeaders(v map[string]*string) *ActivateMcpHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ActivateMcpHeaders) SetXAcsDingtalkAccessToken(v string) *ActivateMcpHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ActivateMcpRequest struct {
	// This parameter is required.
	McpId  *int64  `json:"mcpId,omitempty" xml:"mcpId,omitempty"`
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ActivateMcpRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateMcpRequest) GoString() string {
	return s.String()
}

func (s *ActivateMcpRequest) SetMcpId(v int64) *ActivateMcpRequest {
	s.McpId = &v
	return s
}

func (s *ActivateMcpRequest) SetSource(v string) *ActivateMcpRequest {
	s.Source = &v
	return s
}

type ActivateMcpResponseBody struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	JsonConfig *string `json:"jsonConfig,omitempty" xml:"jsonConfig,omitempty"`
	Url        *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ActivateMcpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateMcpResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateMcpResponseBody) SetInstanceId(v string) *ActivateMcpResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ActivateMcpResponseBody) SetJsonConfig(v string) *ActivateMcpResponseBody {
	s.JsonConfig = &v
	return s
}

func (s *ActivateMcpResponseBody) SetUrl(v string) *ActivateMcpResponseBody {
	s.Url = &v
	return s
}

type ActivateMcpResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateMcpResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateMcpResponse) GoString() string {
	return s.String()
}

func (s *ActivateMcpResponse) SetHeaders(v map[string]*string) *ActivateMcpResponse {
	s.Headers = v
	return s
}

func (s *ActivateMcpResponse) SetStatusCode(v int32) *ActivateMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateMcpResponse) SetBody(v *ActivateMcpResponseBody) *ActivateMcpResponse {
	s.Body = v
	return s
}

type DeleteMcpHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteMcpHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMcpHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMcpHeaders) SetCommonHeaders(v map[string]*string) *DeleteMcpHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMcpHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteMcpHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteMcpRequest struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteMcpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMcpRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcpRequest) SetInstanceId(v string) *DeleteMcpRequest {
	s.InstanceId = &v
	return s
}

type DeleteMcpResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteMcpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMcpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcpResponseBody) SetSuccess(v bool) *DeleteMcpResponseBody {
	s.Success = &v
	return s
}

type DeleteMcpResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMcpResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcpResponse) SetHeaders(v map[string]*string) *DeleteMcpResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcpResponse) SetStatusCode(v int32) *DeleteMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcpResponse) SetBody(v *DeleteMcpResponseBody) *DeleteMcpResponse {
	s.Body = v
	return s
}

type GetMcpDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMcpDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMcpDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetMcpDetailHeaders) SetCommonHeaders(v map[string]*string) *GetMcpDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMcpDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetMcpDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMcpDetailResponseBody struct {
	Categories   []*GetMcpDetailResponseBodyCategories `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	Charged      *bool                                 `json:"charged,omitempty" xml:"charged,omitempty"`
	Description  *string                               `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl    *string                               `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	Icon         *string                               `json:"icon,omitempty" xml:"icon,omitempty"`
	Introduction *string                               `json:"introduction,omitempty" xml:"introduction,omitempty"`
	McpId        *string                               `json:"mcpId,omitempty" xml:"mcpId,omitempty"`
	Name         *string                               `json:"name,omitempty" xml:"name,omitempty"`
	Tools        []*GetMcpDetailResponseBodyTools      `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s GetMcpDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMcpDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpDetailResponseBody) SetCategories(v []*GetMcpDetailResponseBodyCategories) *GetMcpDetailResponseBody {
	s.Categories = v
	return s
}

func (s *GetMcpDetailResponseBody) SetCharged(v bool) *GetMcpDetailResponseBody {
	s.Charged = &v
	return s
}

func (s *GetMcpDetailResponseBody) SetDescription(v string) *GetMcpDetailResponseBody {
	s.Description = &v
	return s
}

func (s *GetMcpDetailResponseBody) SetDetailUrl(v string) *GetMcpDetailResponseBody {
	s.DetailUrl = &v
	return s
}

func (s *GetMcpDetailResponseBody) SetIcon(v string) *GetMcpDetailResponseBody {
	s.Icon = &v
	return s
}

func (s *GetMcpDetailResponseBody) SetIntroduction(v string) *GetMcpDetailResponseBody {
	s.Introduction = &v
	return s
}

func (s *GetMcpDetailResponseBody) SetMcpId(v string) *GetMcpDetailResponseBody {
	s.McpId = &v
	return s
}

func (s *GetMcpDetailResponseBody) SetName(v string) *GetMcpDetailResponseBody {
	s.Name = &v
	return s
}

func (s *GetMcpDetailResponseBody) SetTools(v []*GetMcpDetailResponseBodyTools) *GetMcpDetailResponseBody {
	s.Tools = v
	return s
}

type GetMcpDetailResponseBodyCategories struct {
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s GetMcpDetailResponseBodyCategories) String() string {
	return tea.Prettify(s)
}

func (s GetMcpDetailResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *GetMcpDetailResponseBodyCategories) SetCode(v string) *GetMcpDetailResponseBodyCategories {
	s.Code = &v
	return s
}

func (s *GetMcpDetailResponseBodyCategories) SetDisplayName(v string) *GetMcpDetailResponseBodyCategories {
	s.DisplayName = &v
	return s
}

type GetMcpDetailResponseBodyTools struct {
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName  *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	InputSchema  *string `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	OutputSchema *string `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
}

func (s GetMcpDetailResponseBodyTools) String() string {
	return tea.Prettify(s)
}

func (s GetMcpDetailResponseBodyTools) GoString() string {
	return s.String()
}

func (s *GetMcpDetailResponseBodyTools) SetDescription(v string) *GetMcpDetailResponseBodyTools {
	s.Description = &v
	return s
}

func (s *GetMcpDetailResponseBodyTools) SetDisplayName(v string) *GetMcpDetailResponseBodyTools {
	s.DisplayName = &v
	return s
}

func (s *GetMcpDetailResponseBodyTools) SetInputSchema(v string) *GetMcpDetailResponseBodyTools {
	s.InputSchema = &v
	return s
}

func (s *GetMcpDetailResponseBodyTools) SetName(v string) *GetMcpDetailResponseBodyTools {
	s.Name = &v
	return s
}

func (s *GetMcpDetailResponseBodyTools) SetOutputSchema(v string) *GetMcpDetailResponseBodyTools {
	s.OutputSchema = &v
	return s
}

type GetMcpDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMcpDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMcpDetailResponse) SetHeaders(v map[string]*string) *GetMcpDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMcpDetailResponse) SetStatusCode(v int32) *GetMcpDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpDetailResponse) SetBody(v *GetMcpDetailResponseBody) *GetMcpDetailResponse {
	s.Body = v
	return s
}

type GetSkillDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSkillDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSkillDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetSkillDetailHeaders) SetCommonHeaders(v map[string]*string) *GetSkillDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSkillDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetSkillDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSkillDetailResponseBody struct {
	Categories   []*GetSkillDetailResponseBodyCategories   `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	Dependencies []*GetSkillDetailResponseBodyDependencies `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Repeated"`
	Description  *string                                   `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl    *string                                   `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	Icon         *string                                   `json:"icon,omitempty" xml:"icon,omitempty"`
	Introduction *string                                   `json:"introduction,omitempty" xml:"introduction,omitempty"`
	Name         *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	PackageUrl   *string                                   `json:"packageUrl,omitempty" xml:"packageUrl,omitempty"`
	SkillId      *string                                   `json:"skillId,omitempty" xml:"skillId,omitempty"`
}

func (s GetSkillDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSkillDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillDetailResponseBody) SetCategories(v []*GetSkillDetailResponseBodyCategories) *GetSkillDetailResponseBody {
	s.Categories = v
	return s
}

func (s *GetSkillDetailResponseBody) SetDependencies(v []*GetSkillDetailResponseBodyDependencies) *GetSkillDetailResponseBody {
	s.Dependencies = v
	return s
}

func (s *GetSkillDetailResponseBody) SetDescription(v string) *GetSkillDetailResponseBody {
	s.Description = &v
	return s
}

func (s *GetSkillDetailResponseBody) SetDetailUrl(v string) *GetSkillDetailResponseBody {
	s.DetailUrl = &v
	return s
}

func (s *GetSkillDetailResponseBody) SetIcon(v string) *GetSkillDetailResponseBody {
	s.Icon = &v
	return s
}

func (s *GetSkillDetailResponseBody) SetIntroduction(v string) *GetSkillDetailResponseBody {
	s.Introduction = &v
	return s
}

func (s *GetSkillDetailResponseBody) SetName(v string) *GetSkillDetailResponseBody {
	s.Name = &v
	return s
}

func (s *GetSkillDetailResponseBody) SetPackageUrl(v string) *GetSkillDetailResponseBody {
	s.PackageUrl = &v
	return s
}

func (s *GetSkillDetailResponseBody) SetSkillId(v string) *GetSkillDetailResponseBody {
	s.SkillId = &v
	return s
}

type GetSkillDetailResponseBodyCategories struct {
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s GetSkillDetailResponseBodyCategories) String() string {
	return tea.Prettify(s)
}

func (s GetSkillDetailResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *GetSkillDetailResponseBodyCategories) SetCode(v string) *GetSkillDetailResponseBodyCategories {
	s.Code = &v
	return s
}

func (s *GetSkillDetailResponseBodyCategories) SetDisplayName(v string) *GetSkillDetailResponseBodyCategories {
	s.DisplayName = &v
	return s
}

type GetSkillDetailResponseBodyDependencies struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	RefId       *string `json:"refId,omitempty" xml:"refId,omitempty"`
	Required    *bool   `json:"required,omitempty" xml:"required,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetSkillDetailResponseBodyDependencies) String() string {
	return tea.Prettify(s)
}

func (s GetSkillDetailResponseBodyDependencies) GoString() string {
	return s.String()
}

func (s *GetSkillDetailResponseBodyDependencies) SetDisplayName(v string) *GetSkillDetailResponseBodyDependencies {
	s.DisplayName = &v
	return s
}

func (s *GetSkillDetailResponseBodyDependencies) SetRefId(v string) *GetSkillDetailResponseBodyDependencies {
	s.RefId = &v
	return s
}

func (s *GetSkillDetailResponseBodyDependencies) SetRequired(v bool) *GetSkillDetailResponseBodyDependencies {
	s.Required = &v
	return s
}

func (s *GetSkillDetailResponseBodyDependencies) SetType(v string) *GetSkillDetailResponseBodyDependencies {
	s.Type = &v
	return s
}

type GetSkillDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSkillDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSkillDetailResponse) SetHeaders(v map[string]*string) *GetSkillDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSkillDetailResponse) SetStatusCode(v int32) *GetSkillDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillDetailResponse) SetBody(v *GetSkillDetailResponseBody) *GetSkillDetailResponse {
	s.Body = v
	return s
}

type ListMarketMcpsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListMarketMcpsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListMarketMcpsHeaders) GoString() string {
	return s.String()
}

func (s *ListMarketMcpsHeaders) SetCommonHeaders(v map[string]*string) *ListMarketMcpsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMarketMcpsHeaders) SetXAcsDingtalkAccessToken(v string) *ListMarketMcpsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListMarketMcpsRequest struct {
	Page     *int32 `json:"page,omitempty" xml:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMarketMcpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMarketMcpsRequest) GoString() string {
	return s.String()
}

func (s *ListMarketMcpsRequest) SetPage(v int32) *ListMarketMcpsRequest {
	s.Page = &v
	return s
}

func (s *ListMarketMcpsRequest) SetPageSize(v int32) *ListMarketMcpsRequest {
	s.PageSize = &v
	return s
}

type ListMarketMcpsResponseBody struct {
	Items    []*ListMarketMcpsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Page     *int32                             `json:"page,omitempty" xml:"page,omitempty"`
	PageSize *int32                             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                             `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMarketMcpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMarketMcpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMarketMcpsResponseBody) SetItems(v []*ListMarketMcpsResponseBodyItems) *ListMarketMcpsResponseBody {
	s.Items = v
	return s
}

func (s *ListMarketMcpsResponseBody) SetPage(v int32) *ListMarketMcpsResponseBody {
	s.Page = &v
	return s
}

func (s *ListMarketMcpsResponseBody) SetPageSize(v int32) *ListMarketMcpsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMarketMcpsResponseBody) SetTotal(v int32) *ListMarketMcpsResponseBody {
	s.Total = &v
	return s
}

type ListMarketMcpsResponseBodyItems struct {
	Categories  []*ListMarketMcpsResponseBodyItemsCategories `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	Charged     *bool                                        `json:"charged,omitempty" xml:"charged,omitempty"`
	Description *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl   *string                                      `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	Icon        *string                                      `json:"icon,omitempty" xml:"icon,omitempty"`
	McpId       *string                                      `json:"mcpId,omitempty" xml:"mcpId,omitempty"`
	Name        *string                                      `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListMarketMcpsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListMarketMcpsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMarketMcpsResponseBodyItems) SetCategories(v []*ListMarketMcpsResponseBodyItemsCategories) *ListMarketMcpsResponseBodyItems {
	s.Categories = v
	return s
}

func (s *ListMarketMcpsResponseBodyItems) SetCharged(v bool) *ListMarketMcpsResponseBodyItems {
	s.Charged = &v
	return s
}

func (s *ListMarketMcpsResponseBodyItems) SetDescription(v string) *ListMarketMcpsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListMarketMcpsResponseBodyItems) SetDetailUrl(v string) *ListMarketMcpsResponseBodyItems {
	s.DetailUrl = &v
	return s
}

func (s *ListMarketMcpsResponseBodyItems) SetIcon(v string) *ListMarketMcpsResponseBodyItems {
	s.Icon = &v
	return s
}

func (s *ListMarketMcpsResponseBodyItems) SetMcpId(v string) *ListMarketMcpsResponseBodyItems {
	s.McpId = &v
	return s
}

func (s *ListMarketMcpsResponseBodyItems) SetName(v string) *ListMarketMcpsResponseBodyItems {
	s.Name = &v
	return s
}

type ListMarketMcpsResponseBodyItemsCategories struct {
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s ListMarketMcpsResponseBodyItemsCategories) String() string {
	return tea.Prettify(s)
}

func (s ListMarketMcpsResponseBodyItemsCategories) GoString() string {
	return s.String()
}

func (s *ListMarketMcpsResponseBodyItemsCategories) SetCode(v string) *ListMarketMcpsResponseBodyItemsCategories {
	s.Code = &v
	return s
}

func (s *ListMarketMcpsResponseBodyItemsCategories) SetDisplayName(v string) *ListMarketMcpsResponseBodyItemsCategories {
	s.DisplayName = &v
	return s
}

type ListMarketMcpsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMarketMcpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMarketMcpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMarketMcpsResponse) GoString() string {
	return s.String()
}

func (s *ListMarketMcpsResponse) SetHeaders(v map[string]*string) *ListMarketMcpsResponse {
	s.Headers = v
	return s
}

func (s *ListMarketMcpsResponse) SetStatusCode(v int32) *ListMarketMcpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMarketMcpsResponse) SetBody(v *ListMarketMcpsResponseBody) *ListMarketMcpsResponse {
	s.Body = v
	return s
}

type ListSkillsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSkillsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSkillsHeaders) GoString() string {
	return s.String()
}

func (s *ListSkillsHeaders) SetCommonHeaders(v map[string]*string) *ListSkillsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSkillsHeaders) SetXAcsDingtalkAccessToken(v string) *ListSkillsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSkillsRequest struct {
	Page     *int32 `json:"page,omitempty" xml:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListSkillsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsRequest) SetPage(v int32) *ListSkillsRequest {
	s.Page = &v
	return s
}

func (s *ListSkillsRequest) SetPageSize(v int32) *ListSkillsRequest {
	s.PageSize = &v
	return s
}

type ListSkillsResponseBody struct {
	Items    []*ListSkillsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Page     *int32                         `json:"page,omitempty" xml:"page,omitempty"`
	PageSize *int32                         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                         `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListSkillsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBody) SetItems(v []*ListSkillsResponseBodyItems) *ListSkillsResponseBody {
	s.Items = v
	return s
}

func (s *ListSkillsResponseBody) SetPage(v int32) *ListSkillsResponseBody {
	s.Page = &v
	return s
}

func (s *ListSkillsResponseBody) SetPageSize(v int32) *ListSkillsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSkillsResponseBody) SetTotal(v int32) *ListSkillsResponseBody {
	s.Total = &v
	return s
}

type ListSkillsResponseBodyItems struct {
	Categories  []*ListSkillsResponseBodyItemsCategories `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	Description *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl   *string                                  `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	Icon        *string                                  `json:"icon,omitempty" xml:"icon,omitempty"`
	Name        *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	SkillId     *string                                  `json:"skillId,omitempty" xml:"skillId,omitempty"`
}

func (s ListSkillsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListSkillsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodyItems) SetCategories(v []*ListSkillsResponseBodyItemsCategories) *ListSkillsResponseBodyItems {
	s.Categories = v
	return s
}

func (s *ListSkillsResponseBodyItems) SetDescription(v string) *ListSkillsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetDetailUrl(v string) *ListSkillsResponseBodyItems {
	s.DetailUrl = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetIcon(v string) *ListSkillsResponseBodyItems {
	s.Icon = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetName(v string) *ListSkillsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetSkillId(v string) *ListSkillsResponseBodyItems {
	s.SkillId = &v
	return s
}

type ListSkillsResponseBodyItemsCategories struct {
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s ListSkillsResponseBodyItemsCategories) String() string {
	return tea.Prettify(s)
}

func (s ListSkillsResponseBodyItemsCategories) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodyItemsCategories) SetCode(v string) *ListSkillsResponseBodyItemsCategories {
	s.Code = &v
	return s
}

func (s *ListSkillsResponseBodyItemsCategories) SetDisplayName(v string) *ListSkillsResponseBodyItemsCategories {
	s.DisplayName = &v
	return s
}

type ListSkillsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSkillsResponse) GoString() string {
	return s.String()
}

func (s *ListSkillsResponse) SetHeaders(v map[string]*string) *ListSkillsResponse {
	s.Headers = v
	return s
}

func (s *ListSkillsResponse) SetStatusCode(v int32) *ListSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillsResponse) SetBody(v *ListSkillsResponseBody) *ListSkillsResponse {
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
// 激活MCP
//
// @param request - ActivateMcpRequest
//
// @param headers - ActivateMcpHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateMcpResponse
func (client *Client) ActivateMcpWithOptions(request *ActivateMcpRequest, headers *ActivateMcpHeaders, runtime *util.RuntimeOptions) (_result *ActivateMcpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.McpId)) {
		body["mcpId"] = request.McpId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

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
		Action:      tea.String("ActivateMcp"),
		Version:     tea.String("mcp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mcp/activate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateMcpResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 激活MCP
//
// @param request - ActivateMcpRequest
//
// @return ActivateMcpResponse
func (client *Client) ActivateMcp(request *ActivateMcpRequest) (_result *ActivateMcpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ActivateMcpHeaders{}
	_result = &ActivateMcpResponse{}
	_body, _err := client.ActivateMcpWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除MCP实例
//
// @param request - DeleteMcpRequest
//
// @param headers - DeleteMcpHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcpResponse
func (client *Client) DeleteMcpWithOptions(request *DeleteMcpRequest, headers *DeleteMcpHeaders, runtime *util.RuntimeOptions) (_result *DeleteMcpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	params := &openapi.Params{
		Action:      tea.String("DeleteMcp"),
		Version:     tea.String("mcp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mcp/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMcpResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除MCP实例
//
// @param request - DeleteMcpRequest
//
// @return DeleteMcpResponse
func (client *Client) DeleteMcp(request *DeleteMcpRequest) (_result *DeleteMcpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMcpHeaders{}
	_result = &DeleteMcpResponse{}
	_body, _err := client.DeleteMcpWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据mcpId获取MCP详情
//
// @param headers - GetMcpDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpDetailResponse
func (client *Client) GetMcpDetailWithOptions(mcpId *string, headers *GetMcpDetailHeaders, runtime *util.RuntimeOptions) (_result *GetMcpDetailResponse, _err error) {
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
		Action:      tea.String("GetMcpDetail"),
		Version:     tea.String("mcp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mcp/mcps/" + tea.StringValue(mcpId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMcpDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据mcpId获取MCP详情
//
// @return GetMcpDetailResponse
func (client *Client) GetMcpDetail(mcpId *string) (_result *GetMcpDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMcpDetailHeaders{}
	_result = &GetMcpDetailResponse{}
	_body, _err := client.GetMcpDetailWithOptions(mcpId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据skillId查询Skill详情及安装包下载链接
//
// @param headers - GetSkillDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillDetailResponse
func (client *Client) GetSkillDetailWithOptions(skillId *string, headers *GetSkillDetailHeaders, runtime *util.RuntimeOptions) (_result *GetSkillDetailResponse, _err error) {
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
		Action:      tea.String("GetSkillDetail"),
		Version:     tea.String("mcp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mcp/skills/" + tea.StringValue(skillId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSkillDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据skillId查询Skill详情及安装包下载链接
//
// @return GetSkillDetailResponse
func (client *Client) GetSkillDetail(skillId *string) (_result *GetSkillDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSkillDetailHeaders{}
	_result = &GetSkillDetailResponse{}
	_body, _err := client.GetSkillDetailWithOptions(skillId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询MCP广场精选MCP列表
//
// @param request - ListMarketMcpsRequest
//
// @param headers - ListMarketMcpsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMarketMcpsResponse
func (client *Client) ListMarketMcpsWithOptions(request *ListMarketMcpsRequest, headers *ListMarketMcpsHeaders, runtime *util.RuntimeOptions) (_result *ListMarketMcpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
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
	params := &openapi.Params{
		Action:      tea.String("ListMarketMcps"),
		Version:     tea.String("mcp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mcp/mcps"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMarketMcpsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询MCP广场精选MCP列表
//
// @param request - ListMarketMcpsRequest
//
// @return ListMarketMcpsResponse
func (client *Client) ListMarketMcps(request *ListMarketMcpsRequest) (_result *ListMarketMcpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListMarketMcpsHeaders{}
	_result = &ListMarketMcpsResponse{}
	_body, _err := client.ListMarketMcpsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询Skill广场精选列表
//
// @param request - ListSkillsRequest
//
// @param headers - ListSkillsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillsResponse
func (client *Client) ListSkillsWithOptions(request *ListSkillsRequest, headers *ListSkillsHeaders, runtime *util.RuntimeOptions) (_result *ListSkillsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
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
	params := &openapi.Params{
		Action:      tea.String("ListSkills"),
		Version:     tea.String("mcp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mcp/skills"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSkillsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询Skill广场精选列表
//
// @param request - ListSkillsRequest
//
// @return ListSkillsResponse
func (client *Client) ListSkills(request *ListSkillsRequest) (_result *ListSkillsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSkillsHeaders{}
	_result = &ListSkillsResponse{}
	_body, _err := client.ListSkillsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
