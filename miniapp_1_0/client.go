// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package miniapp_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateMiniAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateMiniAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateMiniAppHeaders) GoString() string {
	return s.String()
}

func (s *CreateMiniAppHeaders) SetCommonHeaders(v map[string]*string) *CreateMiniAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMiniAppHeaders) SetXAcsDingtalkAccessToken(v string) *CreateMiniAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateMiniAppRequest struct {
	BizId    *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	BizType  *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	BundleId *string `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	Desc     *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Icon     *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateMiniAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMiniAppRequest) GoString() string {
	return s.String()
}

func (s *CreateMiniAppRequest) SetBizId(v string) *CreateMiniAppRequest {
	s.BizId = &v
	return s
}

func (s *CreateMiniAppRequest) SetBizType(v int32) *CreateMiniAppRequest {
	s.BizType = &v
	return s
}

func (s *CreateMiniAppRequest) SetBundleId(v string) *CreateMiniAppRequest {
	s.BundleId = &v
	return s
}

func (s *CreateMiniAppRequest) SetDesc(v string) *CreateMiniAppRequest {
	s.Desc = &v
	return s
}

func (s *CreateMiniAppRequest) SetIcon(v string) *CreateMiniAppRequest {
	s.Icon = &v
	return s
}

func (s *CreateMiniAppRequest) SetName(v string) *CreateMiniAppRequest {
	s.Name = &v
	return s
}

type CreateMiniAppResponseBody struct {
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
}

func (s CreateMiniAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMiniAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMiniAppResponseBody) SetMiniAppId(v string) *CreateMiniAppResponseBody {
	s.MiniAppId = &v
	return s
}

type CreateMiniAppResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMiniAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMiniAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMiniAppResponse) GoString() string {
	return s.String()
}

func (s *CreateMiniAppResponse) SetHeaders(v map[string]*string) *CreateMiniAppResponse {
	s.Headers = v
	return s
}

func (s *CreateMiniAppResponse) SetStatusCode(v int32) *CreateMiniAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMiniAppResponse) SetBody(v *CreateMiniAppResponseBody) *CreateMiniAppResponse {
	s.Body = v
	return s
}

type CreateMiniAppPluginHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateMiniAppPluginHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateMiniAppPluginHeaders) GoString() string {
	return s.String()
}

func (s *CreateMiniAppPluginHeaders) SetCommonHeaders(v map[string]*string) *CreateMiniAppPluginHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMiniAppPluginHeaders) SetXAcsDingtalkAccessToken(v string) *CreateMiniAppPluginHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateMiniAppPluginRequest struct {
	BizId    *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	BizType  *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	BundleId *string `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	Desc     *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Icon     *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateMiniAppPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMiniAppPluginRequest) GoString() string {
	return s.String()
}

func (s *CreateMiniAppPluginRequest) SetBizId(v string) *CreateMiniAppPluginRequest {
	s.BizId = &v
	return s
}

func (s *CreateMiniAppPluginRequest) SetBizType(v int32) *CreateMiniAppPluginRequest {
	s.BizType = &v
	return s
}

func (s *CreateMiniAppPluginRequest) SetBundleId(v string) *CreateMiniAppPluginRequest {
	s.BundleId = &v
	return s
}

func (s *CreateMiniAppPluginRequest) SetDesc(v string) *CreateMiniAppPluginRequest {
	s.Desc = &v
	return s
}

func (s *CreateMiniAppPluginRequest) SetIcon(v string) *CreateMiniAppPluginRequest {
	s.Icon = &v
	return s
}

func (s *CreateMiniAppPluginRequest) SetName(v string) *CreateMiniAppPluginRequest {
	s.Name = &v
	return s
}

type CreateMiniAppPluginResponseBody struct {
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
}

func (s CreateMiniAppPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMiniAppPluginResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMiniAppPluginResponseBody) SetMiniAppId(v string) *CreateMiniAppPluginResponseBody {
	s.MiniAppId = &v
	return s
}

type CreateMiniAppPluginResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMiniAppPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMiniAppPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMiniAppPluginResponse) GoString() string {
	return s.String()
}

func (s *CreateMiniAppPluginResponse) SetHeaders(v map[string]*string) *CreateMiniAppPluginResponse {
	s.Headers = v
	return s
}

func (s *CreateMiniAppPluginResponse) SetStatusCode(v int32) *CreateMiniAppPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMiniAppPluginResponse) SetBody(v *CreateMiniAppPluginResponseBody) *CreateMiniAppPluginResponse {
	s.Body = v
	return s
}

type CreateVersionAcrossBundleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateVersionAcrossBundleHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionAcrossBundleHeaders) GoString() string {
	return s.String()
}

func (s *CreateVersionAcrossBundleHeaders) SetCommonHeaders(v map[string]*string) *CreateVersionAcrossBundleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateVersionAcrossBundleHeaders) SetXAcsDingtalkAccessToken(v string) *CreateVersionAcrossBundleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateVersionAcrossBundleRequest struct {
	BundleId       *string `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	MiniAppId      *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	SourceBundleId *string `json:"sourceBundleId,omitempty" xml:"sourceBundleId,omitempty"`
	SourceVersion  *string `json:"sourceVersion,omitempty" xml:"sourceVersion,omitempty"`
	Version        *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateVersionAcrossBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionAcrossBundleRequest) GoString() string {
	return s.String()
}

func (s *CreateVersionAcrossBundleRequest) SetBundleId(v string) *CreateVersionAcrossBundleRequest {
	s.BundleId = &v
	return s
}

func (s *CreateVersionAcrossBundleRequest) SetMiniAppId(v string) *CreateVersionAcrossBundleRequest {
	s.MiniAppId = &v
	return s
}

func (s *CreateVersionAcrossBundleRequest) SetSourceBundleId(v string) *CreateVersionAcrossBundleRequest {
	s.SourceBundleId = &v
	return s
}

func (s *CreateVersionAcrossBundleRequest) SetSourceVersion(v string) *CreateVersionAcrossBundleRequest {
	s.SourceVersion = &v
	return s
}

func (s *CreateVersionAcrossBundleRequest) SetVersion(v string) *CreateVersionAcrossBundleRequest {
	s.Version = &v
	return s
}

type CreateVersionAcrossBundleResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateVersionAcrossBundleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionAcrossBundleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVersionAcrossBundleResponseBody) SetResult(v string) *CreateVersionAcrossBundleResponseBody {
	s.Result = &v
	return s
}

type CreateVersionAcrossBundleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVersionAcrossBundleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVersionAcrossBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVersionAcrossBundleResponse) GoString() string {
	return s.String()
}

func (s *CreateVersionAcrossBundleResponse) SetHeaders(v map[string]*string) *CreateVersionAcrossBundleResponse {
	s.Headers = v
	return s
}

func (s *CreateVersionAcrossBundleResponse) SetStatusCode(v int32) *CreateVersionAcrossBundleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVersionAcrossBundleResponse) SetBody(v *CreateVersionAcrossBundleResponseBody) *CreateVersionAcrossBundleResponse {
	s.Body = v
	return s
}

type GetMaxVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMaxVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMaxVersionHeaders) GoString() string {
	return s.String()
}

func (s *GetMaxVersionHeaders) SetCommonHeaders(v map[string]*string) *GetMaxVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMaxVersionHeaders) SetXAcsDingtalkAccessToken(v string) *GetMaxVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMaxVersionRequest struct {
	BundleId  *string `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	Version   *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetMaxVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMaxVersionRequest) GoString() string {
	return s.String()
}

func (s *GetMaxVersionRequest) SetBundleId(v string) *GetMaxVersionRequest {
	s.BundleId = &v
	return s
}

func (s *GetMaxVersionRequest) SetMiniAppId(v string) *GetMaxVersionRequest {
	s.MiniAppId = &v
	return s
}

func (s *GetMaxVersionRequest) SetVersion(v string) *GetMaxVersionRequest {
	s.Version = &v
	return s
}

type GetMaxVersionResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetMaxVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMaxVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetMaxVersionResponseBody) SetResult(v string) *GetMaxVersionResponseBody {
	s.Result = &v
	return s
}

type GetMaxVersionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMaxVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMaxVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMaxVersionResponse) GoString() string {
	return s.String()
}

func (s *GetMaxVersionResponse) SetHeaders(v map[string]*string) *GetMaxVersionResponse {
	s.Headers = v
	return s
}

func (s *GetMaxVersionResponse) SetStatusCode(v int32) *GetMaxVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMaxVersionResponse) SetBody(v *GetMaxVersionResponseBody) *GetMaxVersionResponse {
	s.Body = v
	return s
}

type GetMiniAppMetaDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMiniAppMetaDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMiniAppMetaDataHeaders) GoString() string {
	return s.String()
}

func (s *GetMiniAppMetaDataHeaders) SetCommonHeaders(v map[string]*string) *GetMiniAppMetaDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMiniAppMetaDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetMiniAppMetaDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMiniAppMetaDataRequest struct {
	BundleId                  *string                `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	BundleIdTableGmtModified  map[string]interface{} `json:"bundleIdTableGmtModified,omitempty" xml:"bundleIdTableGmtModified,omitempty"`
	FromAppName               *string                `json:"fromAppName,omitempty" xml:"fromAppName,omitempty"`
	MiniAppIdTableGmtModified map[string]interface{} `json:"miniAppIdTableGmtModified,omitempty" xml:"miniAppIdTableGmtModified,omitempty"`
}

func (s GetMiniAppMetaDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMiniAppMetaDataRequest) GoString() string {
	return s.String()
}

func (s *GetMiniAppMetaDataRequest) SetBundleId(v string) *GetMiniAppMetaDataRequest {
	s.BundleId = &v
	return s
}

func (s *GetMiniAppMetaDataRequest) SetBundleIdTableGmtModified(v map[string]interface{}) *GetMiniAppMetaDataRequest {
	s.BundleIdTableGmtModified = v
	return s
}

func (s *GetMiniAppMetaDataRequest) SetFromAppName(v string) *GetMiniAppMetaDataRequest {
	s.FromAppName = &v
	return s
}

func (s *GetMiniAppMetaDataRequest) SetMiniAppIdTableGmtModified(v map[string]interface{}) *GetMiniAppMetaDataRequest {
	s.MiniAppIdTableGmtModified = v
	return s
}

type GetMiniAppMetaDataResponseBody struct {
	DingOpenErrcode *int64                                `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string                               `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *GetMiniAppMetaDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success         *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMiniAppMetaDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMiniAppMetaDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetMiniAppMetaDataResponseBody) SetDingOpenErrcode(v int64) *GetMiniAppMetaDataResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *GetMiniAppMetaDataResponseBody) SetErrorMsg(v string) *GetMiniAppMetaDataResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetMiniAppMetaDataResponseBody) SetResult(v *GetMiniAppMetaDataResponseBodyResult) *GetMiniAppMetaDataResponseBody {
	s.Result = v
	return s
}

func (s *GetMiniAppMetaDataResponseBody) SetSuccess(v bool) *GetMiniAppMetaDataResponseBody {
	s.Success = &v
	return s
}

type GetMiniAppMetaDataResponseBodyResult struct {
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetMiniAppMetaDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMiniAppMetaDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMiniAppMetaDataResponseBodyResult) SetData(v map[string]interface{}) *GetMiniAppMetaDataResponseBodyResult {
	s.Data = v
	return s
}

type GetMiniAppMetaDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMiniAppMetaDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMiniAppMetaDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMiniAppMetaDataResponse) GoString() string {
	return s.String()
}

func (s *GetMiniAppMetaDataResponse) SetHeaders(v map[string]*string) *GetMiniAppMetaDataResponse {
	s.Headers = v
	return s
}

func (s *GetMiniAppMetaDataResponse) SetStatusCode(v int32) *GetMiniAppMetaDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMiniAppMetaDataResponse) SetBody(v *GetMiniAppMetaDataResponseBody) *GetMiniAppMetaDataResponse {
	s.Body = v
	return s
}

type GetSettingByMiniAppIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSettingByMiniAppIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSettingByMiniAppIdHeaders) GoString() string {
	return s.String()
}

func (s *GetSettingByMiniAppIdHeaders) SetCommonHeaders(v map[string]*string) *GetSettingByMiniAppIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSettingByMiniAppIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetSettingByMiniAppIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSettingByMiniAppIdResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetSettingByMiniAppIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSettingByMiniAppIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetSettingByMiniAppIdResponseBody) SetResult(v string) *GetSettingByMiniAppIdResponseBody {
	s.Result = &v
	return s
}

type GetSettingByMiniAppIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSettingByMiniAppIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSettingByMiniAppIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSettingByMiniAppIdResponse) GoString() string {
	return s.String()
}

func (s *GetSettingByMiniAppIdResponse) SetHeaders(v map[string]*string) *GetSettingByMiniAppIdResponse {
	s.Headers = v
	return s
}

func (s *GetSettingByMiniAppIdResponse) SetStatusCode(v int32) *GetSettingByMiniAppIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSettingByMiniAppIdResponse) SetBody(v *GetSettingByMiniAppIdResponseBody) *GetSettingByMiniAppIdResponse {
	s.Body = v
	return s
}

type InvokeHtmlBundleBuildHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InvokeHtmlBundleBuildHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvokeHtmlBundleBuildHeaders) GoString() string {
	return s.String()
}

func (s *InvokeHtmlBundleBuildHeaders) SetCommonHeaders(v map[string]*string) *InvokeHtmlBundleBuildHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeHtmlBundleBuildHeaders) SetXAcsDingtalkAccessToken(v string) *InvokeHtmlBundleBuildHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InvokeHtmlBundleBuildRequest struct {
	BundleId  *string `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	Version   *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s InvokeHtmlBundleBuildRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeHtmlBundleBuildRequest) GoString() string {
	return s.String()
}

func (s *InvokeHtmlBundleBuildRequest) SetBundleId(v string) *InvokeHtmlBundleBuildRequest {
	s.BundleId = &v
	return s
}

func (s *InvokeHtmlBundleBuildRequest) SetMiniAppId(v string) *InvokeHtmlBundleBuildRequest {
	s.MiniAppId = &v
	return s
}

func (s *InvokeHtmlBundleBuildRequest) SetVersion(v string) *InvokeHtmlBundleBuildRequest {
	s.Version = &v
	return s
}

type InvokeHtmlBundleBuildResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s InvokeHtmlBundleBuildResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeHtmlBundleBuildResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeHtmlBundleBuildResponseBody) SetResult(v string) *InvokeHtmlBundleBuildResponseBody {
	s.Result = &v
	return s
}

type InvokeHtmlBundleBuildResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvokeHtmlBundleBuildResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeHtmlBundleBuildResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeHtmlBundleBuildResponse) GoString() string {
	return s.String()
}

func (s *InvokeHtmlBundleBuildResponse) SetHeaders(v map[string]*string) *InvokeHtmlBundleBuildResponse {
	s.Headers = v
	return s
}

func (s *InvokeHtmlBundleBuildResponse) SetStatusCode(v int32) *InvokeHtmlBundleBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeHtmlBundleBuildResponse) SetBody(v *InvokeHtmlBundleBuildResponseBody) *InvokeHtmlBundleBuildResponse {
	s.Body = v
	return s
}

type ListAvaiableVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAvaiableVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAvaiableVersionHeaders) GoString() string {
	return s.String()
}

func (s *ListAvaiableVersionHeaders) SetCommonHeaders(v map[string]*string) *ListAvaiableVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAvaiableVersionHeaders) SetXAcsDingtalkAccessToken(v string) *ListAvaiableVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAvaiableVersionRequest struct {
	BundleId       *string  `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	MiniAppId      *string  `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	PageNum        *int32   `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize       *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	VersionTypeSet []*int32 `json:"versionTypeSet,omitempty" xml:"versionTypeSet,omitempty" type:"Repeated"`
}

func (s ListAvaiableVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvaiableVersionRequest) GoString() string {
	return s.String()
}

func (s *ListAvaiableVersionRequest) SetBundleId(v string) *ListAvaiableVersionRequest {
	s.BundleId = &v
	return s
}

func (s *ListAvaiableVersionRequest) SetMiniAppId(v string) *ListAvaiableVersionRequest {
	s.MiniAppId = &v
	return s
}

func (s *ListAvaiableVersionRequest) SetPageNum(v int32) *ListAvaiableVersionRequest {
	s.PageNum = &v
	return s
}

func (s *ListAvaiableVersionRequest) SetPageSize(v int32) *ListAvaiableVersionRequest {
	s.PageSize = &v
	return s
}

func (s *ListAvaiableVersionRequest) SetVersionTypeSet(v []*int32) *ListAvaiableVersionRequest {
	s.VersionTypeSet = v
	return s
}

type ListAvaiableVersionResponseBody struct {
	Versions []*ListAvaiableVersionResponseBodyVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListAvaiableVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvaiableVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvaiableVersionResponseBody) SetVersions(v []*ListAvaiableVersionResponseBodyVersions) *ListAvaiableVersionResponseBody {
	s.Versions = v
	return s
}

type ListAvaiableVersionResponseBodyVersions struct {
	BuildStatus *int64  `json:"buildStatus,omitempty" xml:"buildStatus,omitempty"`
	H5Bundle    *string `json:"h5Bundle,omitempty" xml:"h5Bundle,omitempty"`
	PackageSize *string `json:"packageSize,omitempty" xml:"packageSize,omitempty"`
	PackageUrl  *string `json:"packageUrl,omitempty" xml:"packageUrl,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListAvaiableVersionResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListAvaiableVersionResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListAvaiableVersionResponseBodyVersions) SetBuildStatus(v int64) *ListAvaiableVersionResponseBodyVersions {
	s.BuildStatus = &v
	return s
}

func (s *ListAvaiableVersionResponseBodyVersions) SetH5Bundle(v string) *ListAvaiableVersionResponseBodyVersions {
	s.H5Bundle = &v
	return s
}

func (s *ListAvaiableVersionResponseBodyVersions) SetPackageSize(v string) *ListAvaiableVersionResponseBodyVersions {
	s.PackageSize = &v
	return s
}

func (s *ListAvaiableVersionResponseBodyVersions) SetPackageUrl(v string) *ListAvaiableVersionResponseBodyVersions {
	s.PackageUrl = &v
	return s
}

func (s *ListAvaiableVersionResponseBodyVersions) SetVersion(v string) *ListAvaiableVersionResponseBodyVersions {
	s.Version = &v
	return s
}

type ListAvaiableVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAvaiableVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvaiableVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvaiableVersionResponse) GoString() string {
	return s.String()
}

func (s *ListAvaiableVersionResponse) SetHeaders(v map[string]*string) *ListAvaiableVersionResponse {
	s.Headers = v
	return s
}

func (s *ListAvaiableVersionResponse) SetStatusCode(v int32) *ListAvaiableVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvaiableVersionResponse) SetBody(v *ListAvaiableVersionResponseBody) *ListAvaiableVersionResponse {
	s.Body = v
	return s
}

type QueryHtmlBundleBuildHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHtmlBundleBuildHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHtmlBundleBuildHeaders) GoString() string {
	return s.String()
}

func (s *QueryHtmlBundleBuildHeaders) SetCommonHeaders(v map[string]*string) *QueryHtmlBundleBuildHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHtmlBundleBuildHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHtmlBundleBuildHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHtmlBundleBuildRequest struct {
	BundleId  *string `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	Version   *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QueryHtmlBundleBuildRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHtmlBundleBuildRequest) GoString() string {
	return s.String()
}

func (s *QueryHtmlBundleBuildRequest) SetBundleId(v string) *QueryHtmlBundleBuildRequest {
	s.BundleId = &v
	return s
}

func (s *QueryHtmlBundleBuildRequest) SetMiniAppId(v string) *QueryHtmlBundleBuildRequest {
	s.MiniAppId = &v
	return s
}

func (s *QueryHtmlBundleBuildRequest) SetVersion(v string) *QueryHtmlBundleBuildRequest {
	s.Version = &v
	return s
}

type QueryHtmlBundleBuildResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s QueryHtmlBundleBuildResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHtmlBundleBuildResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHtmlBundleBuildResponseBody) SetResult(v string) *QueryHtmlBundleBuildResponseBody {
	s.Result = &v
	return s
}

type QueryHtmlBundleBuildResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHtmlBundleBuildResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHtmlBundleBuildResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHtmlBundleBuildResponse) GoString() string {
	return s.String()
}

func (s *QueryHtmlBundleBuildResponse) SetHeaders(v map[string]*string) *QueryHtmlBundleBuildResponse {
	s.Headers = v
	return s
}

func (s *QueryHtmlBundleBuildResponse) SetStatusCode(v int32) *QueryHtmlBundleBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHtmlBundleBuildResponse) SetBody(v *QueryHtmlBundleBuildResponseBody) *QueryHtmlBundleBuildResponse {
	s.Body = v
	return s
}

type RollBackVersionRequest struct {
	BundleId        *string `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	MiniAppId       *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	RollbackVersion *string `json:"rollbackVersion,omitempty" xml:"rollbackVersion,omitempty"`
	TargetVersion   *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s RollBackVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s RollBackVersionRequest) GoString() string {
	return s.String()
}

func (s *RollBackVersionRequest) SetBundleId(v string) *RollBackVersionRequest {
	s.BundleId = &v
	return s
}

func (s *RollBackVersionRequest) SetMiniAppId(v string) *RollBackVersionRequest {
	s.MiniAppId = &v
	return s
}

func (s *RollBackVersionRequest) SetRollbackVersion(v string) *RollBackVersionRequest {
	s.RollbackVersion = &v
	return s
}

func (s *RollBackVersionRequest) SetTargetVersion(v string) *RollBackVersionRequest {
	s.TargetVersion = &v
	return s
}

type RollBackVersionResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RollBackVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollBackVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RollBackVersionResponseBody) SetResult(v string) *RollBackVersionResponseBody {
	s.Result = &v
	return s
}

type RollBackVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RollBackVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollBackVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s RollBackVersionResponse) GoString() string {
	return s.String()
}

func (s *RollBackVersionResponse) SetHeaders(v map[string]*string) *RollBackVersionResponse {
	s.Headers = v
	return s
}

func (s *RollBackVersionResponse) SetStatusCode(v int32) *RollBackVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollBackVersionResponse) SetBody(v *RollBackVersionResponseBody) *RollBackVersionResponse {
	s.Body = v
	return s
}

type SetExtendSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetExtendSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetExtendSettingHeaders) GoString() string {
	return s.String()
}

func (s *SetExtendSettingHeaders) SetCommonHeaders(v map[string]*string) *SetExtendSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetExtendSettingHeaders) SetXAcsDingtalkAccessToken(v string) *SetExtendSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetExtendSettingRequest struct {
	BuildH5Bundle *bool   `json:"buildH5Bundle,omitempty" xml:"buildH5Bundle,omitempty"`
	MiniAppId     *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
}

func (s SetExtendSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s SetExtendSettingRequest) GoString() string {
	return s.String()
}

func (s *SetExtendSettingRequest) SetBuildH5Bundle(v bool) *SetExtendSettingRequest {
	s.BuildH5Bundle = &v
	return s
}

func (s *SetExtendSettingRequest) SetMiniAppId(v string) *SetExtendSettingRequest {
	s.MiniAppId = &v
	return s
}

type SetExtendSettingResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SetExtendSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetExtendSettingResponseBody) GoString() string {
	return s.String()
}

func (s *SetExtendSettingResponseBody) SetResult(v string) *SetExtendSettingResponseBody {
	s.Result = &v
	return s
}

type SetExtendSettingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetExtendSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetExtendSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s SetExtendSettingResponse) GoString() string {
	return s.String()
}

func (s *SetExtendSettingResponse) SetHeaders(v map[string]*string) *SetExtendSettingResponse {
	s.Headers = v
	return s
}

func (s *SetExtendSettingResponse) SetStatusCode(v int32) *SetExtendSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *SetExtendSettingResponse) SetBody(v *SetExtendSettingResponseBody) *SetExtendSettingResponse {
	s.Body = v
	return s
}

type UpdateVersionStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateVersionStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateVersionStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateVersionStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateVersionStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateVersionStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateVersionStatusRequest struct {
	BundleId    *string `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	MiniAppId   *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
	VersionType *int32  `json:"versionType,omitempty" xml:"versionType,omitempty"`
}

func (s UpdateVersionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateVersionStatusRequest) SetBundleId(v string) *UpdateVersionStatusRequest {
	s.BundleId = &v
	return s
}

func (s *UpdateVersionStatusRequest) SetMiniAppId(v string) *UpdateVersionStatusRequest {
	s.MiniAppId = &v
	return s
}

func (s *UpdateVersionStatusRequest) SetVersion(v string) *UpdateVersionStatusRequest {
	s.Version = &v
	return s
}

func (s *UpdateVersionStatusRequest) SetVersionType(v int32) *UpdateVersionStatusRequest {
	s.VersionType = &v
	return s
}

type UpdateVersionStatusResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateVersionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVersionStatusResponseBody) SetResult(v string) *UpdateVersionStatusResponseBody {
	s.Result = &v
	return s
}

type UpdateVersionStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVersionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVersionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVersionStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateVersionStatusResponse) SetHeaders(v map[string]*string) *UpdateVersionStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateVersionStatusResponse) SetStatusCode(v int32) *UpdateVersionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVersionStatusResponse) SetBody(v *UpdateVersionStatusResponseBody) *UpdateVersionStatusResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) CreateMiniAppWithOptions(request *CreateMiniAppRequest, headers *CreateMiniAppHeaders, runtime *util.RuntimeOptions) (_result *CreateMiniAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		body["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

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
		Action:      tea.String("CreateMiniApp"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/apps"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMiniAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMiniApp(request *CreateMiniAppRequest) (_result *CreateMiniAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateMiniAppHeaders{}
	_result = &CreateMiniAppResponse{}
	_body, _err := client.CreateMiniAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMiniAppPluginWithOptions(request *CreateMiniAppPluginRequest, headers *CreateMiniAppPluginHeaders, runtime *util.RuntimeOptions) (_result *CreateMiniAppPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		body["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

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
		Action:      tea.String("CreateMiniAppPlugin"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/plugins"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMiniAppPluginResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMiniAppPlugin(request *CreateMiniAppPluginRequest) (_result *CreateMiniAppPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateMiniAppPluginHeaders{}
	_result = &CreateMiniAppPluginResponse{}
	_body, _err := client.CreateMiniAppPluginWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVersionAcrossBundleWithOptions(request *CreateVersionAcrossBundleRequest, headers *CreateVersionAcrossBundleHeaders, runtime *util.RuntimeOptions) (_result *CreateVersionAcrossBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		body["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceBundleId)) {
		body["sourceBundleId"] = request.SourceBundleId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceVersion)) {
		body["sourceVersion"] = request.SourceVersion
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
		Action:      tea.String("CreateVersionAcrossBundle"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/versions/createAcrossBundle"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVersionAcrossBundleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVersionAcrossBundle(request *CreateVersionAcrossBundleRequest) (_result *CreateVersionAcrossBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateVersionAcrossBundleHeaders{}
	_result = &CreateVersionAcrossBundleResponse{}
	_body, _err := client.CreateVersionAcrossBundleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMaxVersionWithOptions(request *GetMaxVersionRequest, headers *GetMaxVersionHeaders, runtime *util.RuntimeOptions) (_result *GetMaxVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
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
	params := &openapi.Params{
		Action:      tea.String("GetMaxVersion"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/apps/maxVersions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMaxVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMaxVersion(request *GetMaxVersionRequest) (_result *GetMaxVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMaxVersionHeaders{}
	_result = &GetMaxVersionResponse{}
	_body, _err := client.GetMaxVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMiniAppMetaDataWithOptions(request *GetMiniAppMetaDataRequest, headers *GetMiniAppMetaDataHeaders, runtime *util.RuntimeOptions) (_result *GetMiniAppMetaDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		body["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.BundleIdTableGmtModified)) {
		body["bundleIdTableGmtModified"] = request.BundleIdTableGmtModified
	}

	if !tea.BoolValue(util.IsUnset(request.FromAppName)) {
		body["fromAppName"] = request.FromAppName
	}

	if !tea.BoolValue(util.IsUnset(request.MiniAppIdTableGmtModified)) {
		body["miniAppIdTableGmtModified"] = request.MiniAppIdTableGmtModified
	}

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
		Action:      tea.String("GetMiniAppMetaData"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/apps/metadata"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMiniAppMetaDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMiniAppMetaData(request *GetMiniAppMetaDataRequest) (_result *GetMiniAppMetaDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMiniAppMetaDataHeaders{}
	_result = &GetMiniAppMetaDataResponse{}
	_body, _err := client.GetMiniAppMetaDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSettingByMiniAppIdWithOptions(miniAppId *string, headers *GetSettingByMiniAppIdHeaders, runtime *util.RuntimeOptions) (_result *GetSettingByMiniAppIdResponse, _err error) {
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
		Action:      tea.String("GetSettingByMiniAppId"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/apps/settings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSettingByMiniAppIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSettingByMiniAppId(miniAppId *string) (_result *GetSettingByMiniAppIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSettingByMiniAppIdHeaders{}
	_result = &GetSettingByMiniAppIdResponse{}
	_body, _err := client.GetSettingByMiniAppIdWithOptions(miniAppId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeHtmlBundleBuildWithOptions(request *InvokeHtmlBundleBuildRequest, headers *InvokeHtmlBundleBuildHeaders, runtime *util.RuntimeOptions) (_result *InvokeHtmlBundleBuildResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		body["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
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
		Action:      tea.String("InvokeHtmlBundleBuild"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/h5Bundles/build"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeHtmlBundleBuildResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeHtmlBundleBuild(request *InvokeHtmlBundleBuildRequest) (_result *InvokeHtmlBundleBuildResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvokeHtmlBundleBuildHeaders{}
	_result = &InvokeHtmlBundleBuildResponse{}
	_body, _err := client.InvokeHtmlBundleBuildWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvaiableVersionWithOptions(request *ListAvaiableVersionRequest, headers *ListAvaiableVersionHeaders, runtime *util.RuntimeOptions) (_result *ListAvaiableVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		body["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.VersionTypeSet)) {
		body["versionTypeSet"] = request.VersionTypeSet
	}

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
		Action:      tea.String("ListAvaiableVersion"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/apps/versions/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvaiableVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvaiableVersion(request *ListAvaiableVersionRequest) (_result *ListAvaiableVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAvaiableVersionHeaders{}
	_result = &ListAvaiableVersionResponse{}
	_body, _err := client.ListAvaiableVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHtmlBundleBuildWithOptions(request *QueryHtmlBundleBuildRequest, headers *QueryHtmlBundleBuildHeaders, runtime *util.RuntimeOptions) (_result *QueryHtmlBundleBuildResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
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
	params := &openapi.Params{
		Action:      tea.String("QueryHtmlBundleBuild"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/h5Bundles/buildResults"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHtmlBundleBuildResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHtmlBundleBuild(request *QueryHtmlBundleBuildRequest) (_result *QueryHtmlBundleBuildResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHtmlBundleBuildHeaders{}
	_result = &QueryHtmlBundleBuildResponse{}
	_body, _err := client.QueryHtmlBundleBuildWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollBackVersionWithOptions(request *RollBackVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RollBackVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		body["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.RollbackVersion)) {
		body["rollbackVersion"] = request.RollbackVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersion)) {
		body["targetVersion"] = request.TargetVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RollBackVersion"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/versions/rollback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RollBackVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollBackVersion(request *RollBackVersionRequest) (_result *RollBackVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RollBackVersionResponse{}
	_body, _err := client.RollBackVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetExtendSettingWithOptions(request *SetExtendSettingRequest, headers *SetExtendSettingHeaders, runtime *util.RuntimeOptions) (_result *SetExtendSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildH5Bundle)) {
		body["buildH5Bundle"] = request.BuildH5Bundle
	}

	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

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
		Action:      tea.String("SetExtendSetting"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/apps/settings"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetExtendSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetExtendSetting(request *SetExtendSettingRequest) (_result *SetExtendSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetExtendSettingHeaders{}
	_result = &SetExtendSettingResponse{}
	_body, _err := client.SetExtendSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVersionStatusWithOptions(request *UpdateVersionStatusRequest, headers *UpdateVersionStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateVersionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		body["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		body["versionType"] = request.VersionType
	}

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
		Action:      tea.String("UpdateVersionStatus"),
		Version:     tea.String("miniapp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/miniapp/versions/status"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVersionStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVersionStatus(request *UpdateVersionStatusRequest) (_result *UpdateVersionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateVersionStatusHeaders{}
	_result = &UpdateVersionStatusResponse{}
	_body, _err := client.UpdateVersionStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
