// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package micro_app_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddAppToWorkBenchGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddAppToWorkBenchGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupHeaders) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupHeaders) SetCommonHeaders(v map[string]*string) *AddAppToWorkBenchGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAppToWorkBenchGroupHeaders) SetXAcsDingtalkAccessToken(v string) *AddAppToWorkBenchGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddAppToWorkBenchGroupRequest struct {
	// 创建人unionId
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// 关联组织corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
	// 工作台分组id
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
}

func (s AddAppToWorkBenchGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupRequest) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupRequest) SetOpUnionId(v string) *AddAppToWorkBenchGroupRequest {
	s.OpUnionId = &v
	return s
}

func (s *AddAppToWorkBenchGroupRequest) SetEcologicalCorpId(v string) *AddAppToWorkBenchGroupRequest {
	s.EcologicalCorpId = &v
	return s
}

func (s *AddAppToWorkBenchGroupRequest) SetComponentId(v string) *AddAppToWorkBenchGroupRequest {
	s.ComponentId = &v
	return s
}

type AddAppToWorkBenchGroupResponseBody struct {
	// 更新结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddAppToWorkBenchGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupResponseBody) SetResult(v bool) *AddAppToWorkBenchGroupResponseBody {
	s.Result = &v
	return s
}

type AddAppToWorkBenchGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAppToWorkBenchGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAppToWorkBenchGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupResponse) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupResponse) SetHeaders(v map[string]*string) *AddAppToWorkBenchGroupResponse {
	s.Headers = v
	return s
}

func (s *AddAppToWorkBenchGroupResponse) SetBody(v *AddAppToWorkBenchGroupResponseBody) *AddAppToWorkBenchGroupResponse {
	s.Body = v
	return s
}

type CreateInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *CreateInnerAppHeaders) SetCommonHeaders(v map[string]*string) *CreateInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *CreateInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateInnerAppRequest struct {
	// 创建人unionId
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// 关联组织corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
	// 应用名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 应用描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 应用图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 应用首页地址
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// 应用PC端地址
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	// 应用管理后台地址
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// 服务器出口ip白名单
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
	// 权限类型
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
}

func (s CreateInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppRequest) GoString() string {
	return s.String()
}

func (s *CreateInnerAppRequest) SetOpUnionId(v string) *CreateInnerAppRequest {
	s.OpUnionId = &v
	return s
}

func (s *CreateInnerAppRequest) SetEcologicalCorpId(v string) *CreateInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

func (s *CreateInnerAppRequest) SetName(v string) *CreateInnerAppRequest {
	s.Name = &v
	return s
}

func (s *CreateInnerAppRequest) SetDesc(v string) *CreateInnerAppRequest {
	s.Desc = &v
	return s
}

func (s *CreateInnerAppRequest) SetIcon(v string) *CreateInnerAppRequest {
	s.Icon = &v
	return s
}

func (s *CreateInnerAppRequest) SetHomepageLink(v string) *CreateInnerAppRequest {
	s.HomepageLink = &v
	return s
}

func (s *CreateInnerAppRequest) SetPcHomepageLink(v string) *CreateInnerAppRequest {
	s.PcHomepageLink = &v
	return s
}

func (s *CreateInnerAppRequest) SetOmpLink(v string) *CreateInnerAppRequest {
	s.OmpLink = &v
	return s
}

func (s *CreateInnerAppRequest) SetIpWhiteList(v []*string) *CreateInnerAppRequest {
	s.IpWhiteList = v
	return s
}

func (s *CreateInnerAppRequest) SetScopeType(v string) *CreateInnerAppRequest {
	s.ScopeType = &v
	return s
}

type CreateInnerAppResponseBody struct {
	// 应用id
	AgentId   *int64  `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AppKey    *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
}

func (s CreateInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInnerAppResponseBody) SetAgentId(v int64) *CreateInnerAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *CreateInnerAppResponseBody) SetAppKey(v string) *CreateInnerAppResponseBody {
	s.AppKey = &v
	return s
}

func (s *CreateInnerAppResponseBody) SetAppSecret(v string) *CreateInnerAppResponseBody {
	s.AppSecret = &v
	return s
}

type CreateInnerAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppResponse) GoString() string {
	return s.String()
}

func (s *CreateInnerAppResponse) SetHeaders(v map[string]*string) *CreateInnerAppResponse {
	s.Headers = v
	return s
}

func (s *CreateInnerAppResponse) SetBody(v *CreateInnerAppResponseBody) *CreateInnerAppResponse {
	s.Body = v
	return s
}

type UpdateInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppHeaders) SetCommonHeaders(v map[string]*string) *UpdateInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInnerAppRequest struct {
	// 创建人unionId
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// 关联组织corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
	// 应用名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 应用描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 应用图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 应用首页地址
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// 应用PC端地址
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	// 应用管理后台地址
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// 服务器出口ip白名单
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
}

func (s UpdateInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppRequest) SetOpUnionId(v string) *UpdateInnerAppRequest {
	s.OpUnionId = &v
	return s
}

func (s *UpdateInnerAppRequest) SetEcologicalCorpId(v string) *UpdateInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

func (s *UpdateInnerAppRequest) SetName(v string) *UpdateInnerAppRequest {
	s.Name = &v
	return s
}

func (s *UpdateInnerAppRequest) SetDesc(v string) *UpdateInnerAppRequest {
	s.Desc = &v
	return s
}

func (s *UpdateInnerAppRequest) SetIcon(v string) *UpdateInnerAppRequest {
	s.Icon = &v
	return s
}

func (s *UpdateInnerAppRequest) SetHomepageLink(v string) *UpdateInnerAppRequest {
	s.HomepageLink = &v
	return s
}

func (s *UpdateInnerAppRequest) SetPcHomepageLink(v string) *UpdateInnerAppRequest {
	s.PcHomepageLink = &v
	return s
}

func (s *UpdateInnerAppRequest) SetOmpLink(v string) *UpdateInnerAppRequest {
	s.OmpLink = &v
	return s
}

func (s *UpdateInnerAppRequest) SetIpWhiteList(v []*string) *UpdateInnerAppRequest {
	s.IpWhiteList = v
	return s
}

type UpdateInnerAppResponseBody struct {
	// 更新结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppResponseBody) SetResult(v bool) *UpdateInnerAppResponseBody {
	s.Result = &v
	return s
}

type UpdateInnerAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppResponse) SetHeaders(v map[string]*string) *UpdateInnerAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateInnerAppResponse) SetBody(v *UpdateInnerAppResponseBody) *UpdateInnerAppResponse {
	s.Body = v
	return s
}

type DeleteInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppHeaders) SetCommonHeaders(v map[string]*string) *DeleteInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteInnerAppRequest struct {
	// 操作人unionId
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// 合作空间corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
}

func (s DeleteInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppRequest) SetOpUnionId(v string) *DeleteInnerAppRequest {
	s.OpUnionId = &v
	return s
}

func (s *DeleteInnerAppRequest) SetEcologicalCorpId(v string) *DeleteInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

type DeleteInnerAppResponseBody struct {
	// 删除结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppResponseBody) SetResult(v bool) *DeleteInnerAppResponseBody {
	s.Result = &v
	return s
}

type DeleteInnerAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppResponse) SetHeaders(v map[string]*string) *DeleteInnerAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteInnerAppResponse) SetBody(v *DeleteInnerAppResponseBody) *DeleteInnerAppResponse {
	s.Body = v
	return s
}

type GetInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *GetInnerAppHeaders) SetCommonHeaders(v map[string]*string) *GetInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *GetInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInnerAppRequest struct {
	// 操作人unionId
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// 关联组织corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
}

func (s GetInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppRequest) GoString() string {
	return s.String()
}

func (s *GetInnerAppRequest) SetOpUnionId(v string) *GetInnerAppRequest {
	s.OpUnionId = &v
	return s
}

func (s *GetInnerAppRequest) SetEcologicalCorpId(v string) *GetInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

type GetInnerAppResponseBody struct {
	// 应用id
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// 应用名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 应用描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 应用图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 应用移动端首页地址
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// 应用PC端首页地址
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	// 应用管理后台地址
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// 应用的appkey
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// 应用的secret
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
	// 服务器出口ip
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
}

func (s GetInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetInnerAppResponseBody) SetAgentId(v int64) *GetInnerAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *GetInnerAppResponseBody) SetName(v string) *GetInnerAppResponseBody {
	s.Name = &v
	return s
}

func (s *GetInnerAppResponseBody) SetDesc(v string) *GetInnerAppResponseBody {
	s.Desc = &v
	return s
}

func (s *GetInnerAppResponseBody) SetIcon(v string) *GetInnerAppResponseBody {
	s.Icon = &v
	return s
}

func (s *GetInnerAppResponseBody) SetHomepageLink(v string) *GetInnerAppResponseBody {
	s.HomepageLink = &v
	return s
}

func (s *GetInnerAppResponseBody) SetPcHomepageLink(v string) *GetInnerAppResponseBody {
	s.PcHomepageLink = &v
	return s
}

func (s *GetInnerAppResponseBody) SetOmpLink(v string) *GetInnerAppResponseBody {
	s.OmpLink = &v
	return s
}

func (s *GetInnerAppResponseBody) SetAppKey(v string) *GetInnerAppResponseBody {
	s.AppKey = &v
	return s
}

func (s *GetInnerAppResponseBody) SetAppSecret(v string) *GetInnerAppResponseBody {
	s.AppSecret = &v
	return s
}

func (s *GetInnerAppResponseBody) SetIpWhiteList(v []*string) *GetInnerAppResponseBody {
	s.IpWhiteList = v
	return s
}

type GetInnerAppResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppResponse) GoString() string {
	return s.String()
}

func (s *GetInnerAppResponse) SetHeaders(v map[string]*string) *GetInnerAppResponse {
	s.Headers = v
	return s
}

func (s *GetInnerAppResponse) SetBody(v *GetInnerAppResponseBody) *GetInnerAppResponse {
	s.Body = v
	return s
}

type ListInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *ListInnerAppHeaders) SetCommonHeaders(v map[string]*string) *ListInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *ListInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListInnerAppRequest struct {
	// 合作空间corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
}

func (s ListInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppRequest) GoString() string {
	return s.String()
}

func (s *ListInnerAppRequest) SetEcologicalCorpId(v string) *ListInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

type ListInnerAppResponseBody struct {
	// 应用列表
	AppList []*ListInnerAppResponseBodyAppList `json:"appList,omitempty" xml:"appList,omitempty" type:"Repeated"`
}

func (s ListInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListInnerAppResponseBody) SetAppList(v []*ListInnerAppResponseBodyAppList) *ListInnerAppResponseBody {
	s.AppList = v
	return s
}

type ListInnerAppResponseBodyAppList struct {
	// 应用id
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// 应用名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 应用描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 应用图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 应用移动端首页地址
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// 应用PC端首页地址
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	// 应用管理后台地址
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
}

func (s ListInnerAppResponseBodyAppList) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *ListInnerAppResponseBodyAppList) SetAgentId(v int64) *ListInnerAppResponseBodyAppList {
	s.AgentId = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetName(v string) *ListInnerAppResponseBodyAppList {
	s.Name = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetDesc(v string) *ListInnerAppResponseBodyAppList {
	s.Desc = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetIcon(v string) *ListInnerAppResponseBodyAppList {
	s.Icon = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetHomepageLink(v string) *ListInnerAppResponseBodyAppList {
	s.HomepageLink = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetPcHomepageLink(v string) *ListInnerAppResponseBodyAppList {
	s.PcHomepageLink = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetOmpLink(v string) *ListInnerAppResponseBodyAppList {
	s.OmpLink = &v
	return s
}

type ListInnerAppResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppResponse) GoString() string {
	return s.String()
}

func (s *ListInnerAppResponse) SetHeaders(v map[string]*string) *ListInnerAppResponse {
	s.Headers = v
	return s
}

func (s *ListInnerAppResponse) SetBody(v *ListInnerAppResponseBody) *ListInnerAppResponse {
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

func (client *Client) AddAppToWorkBenchGroup(agentId *string, request *AddAppToWorkBenchGroupRequest) (_result *AddAppToWorkBenchGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddAppToWorkBenchGroupHeaders{}
	_result = &AddAppToWorkBenchGroupResponse{}
	_body, _err := client.AddAppToWorkBenchGroupWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAppToWorkBenchGroupWithOptions(agentId *string, request *AddAppToWorkBenchGroupRequest, headers *AddAppToWorkBenchGroupHeaders, runtime *util.RuntimeOptions) (_result *AddAppToWorkBenchGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		body["ecologicalCorpId"] = request.EcologicalCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentId)) {
		body["componentId"] = request.ComponentId
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
	_result = &AddAppToWorkBenchGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("AddAppToWorkBenchGroup"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/addToWorkBenchGroup"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInnerApp(request *CreateInnerAppRequest) (_result *CreateInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateInnerAppHeaders{}
	_result = &CreateInnerAppResponse{}
	_body, _err := client.CreateInnerAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInnerAppWithOptions(request *CreateInnerAppRequest, headers *CreateInnerAppHeaders, runtime *util.RuntimeOptions) (_result *CreateInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		body["ecologicalCorpId"] = request.EcologicalCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageLink)) {
		body["homepageLink"] = request.HomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.PcHomepageLink)) {
		body["pcHomepageLink"] = request.PcHomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.OmpLink)) {
		body["ompLink"] = request.OmpLink
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhiteList)) {
		body["ipWhiteList"] = request.IpWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeType)) {
		body["scopeType"] = request.ScopeType
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
	_result = &CreateInnerAppResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInnerApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/microApp/apps"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInnerApp(agentId *string, request *UpdateInnerAppRequest) (_result *UpdateInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInnerAppHeaders{}
	_result = &UpdateInnerAppResponse{}
	_body, _err := client.UpdateInnerAppWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInnerAppWithOptions(agentId *string, request *UpdateInnerAppRequest, headers *UpdateInnerAppHeaders, runtime *util.RuntimeOptions) (_result *UpdateInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		body["ecologicalCorpId"] = request.EcologicalCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageLink)) {
		body["homepageLink"] = request.HomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.PcHomepageLink)) {
		body["pcHomepageLink"] = request.PcHomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.OmpLink)) {
		body["ompLink"] = request.OmpLink
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhiteList)) {
		body["ipWhiteList"] = request.IpWhiteList
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
	_result = &UpdateInnerAppResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInnerApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInnerApp(agentId *string, request *DeleteInnerAppRequest) (_result *DeleteInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteInnerAppHeaders{}
	_result = &DeleteInnerAppResponse{}
	_body, _err := client.DeleteInnerAppWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInnerAppWithOptions(agentId *string, request *DeleteInnerAppRequest, headers *DeleteInnerAppHeaders, runtime *util.RuntimeOptions) (_result *DeleteInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		query["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		query["ecologicalCorpId"] = request.EcologicalCorpId
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
	_result = &DeleteInnerAppResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteInnerApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInnerApp(agentId *string, request *GetInnerAppRequest) (_result *GetInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInnerAppHeaders{}
	_result = &GetInnerAppResponse{}
	_body, _err := client.GetInnerAppWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInnerAppWithOptions(agentId *string, request *GetInnerAppRequest, headers *GetInnerAppHeaders, runtime *util.RuntimeOptions) (_result *GetInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		query["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		query["ecologicalCorpId"] = request.EcologicalCorpId
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
	_result = &GetInnerAppResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInnerApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInnerApp(request *ListInnerAppRequest) (_result *ListInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListInnerAppHeaders{}
	_result = &ListInnerAppResponse{}
	_body, _err := client.ListInnerAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInnerAppWithOptions(request *ListInnerAppRequest, headers *ListInnerAppHeaders, runtime *util.RuntimeOptions) (_result *ListInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		query["ecologicalCorpId"] = request.EcologicalCorpId
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
	_result = &ListInnerAppResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInnerApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/microApp/apps"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
