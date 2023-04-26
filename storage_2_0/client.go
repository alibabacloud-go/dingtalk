// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package storage_2_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DentryAppPropertiesValue struct {
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Value      *string `json:"value,omitempty" xml:"value,omitempty"`
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s DentryAppPropertiesValue) String() string {
	return tea.Prettify(s)
}

func (s DentryAppPropertiesValue) GoString() string {
	return s.String()
}

func (s *DentryAppPropertiesValue) SetName(v string) *DentryAppPropertiesValue {
	s.Name = &v
	return s
}

func (s *DentryAppPropertiesValue) SetValue(v string) *DentryAppPropertiesValue {
	s.Value = &v
	return s
}

func (s *DentryAppPropertiesValue) SetVisibility(v string) *DentryAppPropertiesValue {
	s.Visibility = &v
	return s
}

type CommitFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CommitFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s CommitFileHeaders) GoString() string {
	return s.String()
}

func (s *CommitFileHeaders) SetCommonHeaders(v map[string]*string) *CommitFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CommitFileHeaders) SetXAcsDingtalkAccessToken(v string) *CommitFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CommitFileRequest struct {
	Name      *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Option    *CommitFileRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	UploadKey *string                  `json:"uploadKey,omitempty" xml:"uploadKey,omitempty"`
	UnionId   *string                  `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CommitFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CommitFileRequest) GoString() string {
	return s.String()
}

func (s *CommitFileRequest) SetName(v string) *CommitFileRequest {
	s.Name = &v
	return s
}

func (s *CommitFileRequest) SetOption(v *CommitFileRequestOption) *CommitFileRequest {
	s.Option = v
	return s
}

func (s *CommitFileRequest) SetUploadKey(v string) *CommitFileRequest {
	s.UploadKey = &v
	return s
}

func (s *CommitFileRequest) SetUnionId(v string) *CommitFileRequest {
	s.UnionId = &v
	return s
}

type CommitFileRequestOption struct {
	AppProperties                        []*CommitFileRequestOptionAppProperties `json:"appProperties,omitempty" xml:"appProperties,omitempty" type:"Repeated"`
	ConflictStrategy                     *string                                 `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
	ConvertToOnlineDoc                   *bool                                   `json:"convertToOnlineDoc,omitempty" xml:"convertToOnlineDoc,omitempty"`
	ConvertToOnlineDocTargetDocumentType *string                                 `json:"convertToOnlineDocTargetDocumentType,omitempty" xml:"convertToOnlineDocTargetDocumentType,omitempty"`
	Size                                 *int64                                  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s CommitFileRequestOption) String() string {
	return tea.Prettify(s)
}

func (s CommitFileRequestOption) GoString() string {
	return s.String()
}

func (s *CommitFileRequestOption) SetAppProperties(v []*CommitFileRequestOptionAppProperties) *CommitFileRequestOption {
	s.AppProperties = v
	return s
}

func (s *CommitFileRequestOption) SetConflictStrategy(v string) *CommitFileRequestOption {
	s.ConflictStrategy = &v
	return s
}

func (s *CommitFileRequestOption) SetConvertToOnlineDoc(v bool) *CommitFileRequestOption {
	s.ConvertToOnlineDoc = &v
	return s
}

func (s *CommitFileRequestOption) SetConvertToOnlineDocTargetDocumentType(v string) *CommitFileRequestOption {
	s.ConvertToOnlineDocTargetDocumentType = &v
	return s
}

func (s *CommitFileRequestOption) SetSize(v int64) *CommitFileRequestOption {
	s.Size = &v
	return s
}

type CommitFileRequestOptionAppProperties struct {
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Value      *string `json:"value,omitempty" xml:"value,omitempty"`
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s CommitFileRequestOptionAppProperties) String() string {
	return tea.Prettify(s)
}

func (s CommitFileRequestOptionAppProperties) GoString() string {
	return s.String()
}

func (s *CommitFileRequestOptionAppProperties) SetName(v string) *CommitFileRequestOptionAppProperties {
	s.Name = &v
	return s
}

func (s *CommitFileRequestOptionAppProperties) SetValue(v string) *CommitFileRequestOptionAppProperties {
	s.Value = &v
	return s
}

func (s *CommitFileRequestOptionAppProperties) SetVisibility(v string) *CommitFileRequestOptionAppProperties {
	s.Visibility = &v
	return s
}

type CommitFileResponseBody struct {
	Dentry *CommitFileResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
}

func (s CommitFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponseBody) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBody) SetDentry(v *CommitFileResponseBodyDentry) *CommitFileResponseBody {
	s.Dentry = v
	return s
}

type CommitFileResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue  `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	Category      *string                                 `json:"category,omitempty" xml:"category,omitempty"`
	CreateTime    *string                                 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId     *string                                 `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension     *string                                 `json:"extension,omitempty" xml:"extension,omitempty"`
	Id            *string                                 `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedTime  *string                                 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId    *string                                 `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name          *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	ParentId      *string                                 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	PartitionType *string                                 `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	Path          *string                                 `json:"path,omitempty" xml:"path,omitempty"`
	Properties    *CommitFileResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	Size          *int64                                  `json:"size,omitempty" xml:"size,omitempty"`
	SpaceId       *string                                 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status        *string                                 `json:"status,omitempty" xml:"status,omitempty"`
	StorageDriver *string                                 `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	Thumbnail     *CommitFileResponseBodyDentryThumbnail  `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
	Type          *string                                 `json:"type,omitempty" xml:"type,omitempty"`
	Uuid          *string                                 `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version       *int64                                  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CommitFileResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *CommitFileResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *CommitFileResponseBodyDentry) SetCategory(v string) *CommitFileResponseBodyDentry {
	s.Category = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetCreateTime(v string) *CommitFileResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetCreatorId(v string) *CommitFileResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetExtension(v string) *CommitFileResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetId(v string) *CommitFileResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetModifiedTime(v string) *CommitFileResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetModifierId(v string) *CommitFileResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetName(v string) *CommitFileResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetParentId(v string) *CommitFileResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetPartitionType(v string) *CommitFileResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetPath(v string) *CommitFileResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetProperties(v *CommitFileResponseBodyDentryProperties) *CommitFileResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *CommitFileResponseBodyDentry) SetSize(v int64) *CommitFileResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetSpaceId(v string) *CommitFileResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetStatus(v string) *CommitFileResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetStorageDriver(v string) *CommitFileResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetThumbnail(v *CommitFileResponseBodyDentryThumbnail) *CommitFileResponseBodyDentry {
	s.Thumbnail = v
	return s
}

func (s *CommitFileResponseBodyDentry) SetType(v string) *CommitFileResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetUuid(v string) *CommitFileResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetVersion(v int64) *CommitFileResponseBodyDentry {
	s.Version = &v
	return s
}

type CommitFileResponseBodyDentryProperties struct {
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s CommitFileResponseBodyDentryProperties) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBodyDentryProperties) SetReadOnly(v bool) *CommitFileResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

type CommitFileResponseBodyDentryThumbnail struct {
	Height *int32  `json:"height,omitempty" xml:"height,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
	Width  *int32  `json:"width,omitempty" xml:"width,omitempty"`
}

func (s CommitFileResponseBodyDentryThumbnail) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponseBodyDentryThumbnail) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBodyDentryThumbnail) SetHeight(v int32) *CommitFileResponseBodyDentryThumbnail {
	s.Height = &v
	return s
}

func (s *CommitFileResponseBodyDentryThumbnail) SetUrl(v string) *CommitFileResponseBodyDentryThumbnail {
	s.Url = &v
	return s
}

func (s *CommitFileResponseBodyDentryThumbnail) SetWidth(v int32) *CommitFileResponseBodyDentryThumbnail {
	s.Width = &v
	return s
}

type CommitFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CommitFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CommitFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponse) GoString() string {
	return s.String()
}

func (s *CommitFileResponse) SetHeaders(v map[string]*string) *CommitFileResponse {
	s.Headers = v
	return s
}

func (s *CommitFileResponse) SetStatusCode(v int32) *CommitFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitFileResponse) SetBody(v *CommitFileResponseBody) *CommitFileResponse {
	s.Body = v
	return s
}

type GetFileUploadInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFileUploadInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFileUploadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileUploadInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetFileUploadInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFileUploadInfoRequest struct {
	Option   *GetFileUploadInfoRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	Protocol *string                         `json:"protocol,omitempty" xml:"protocol,omitempty"`
	UnionId  *string                         `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetFileUploadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequest) SetOption(v *GetFileUploadInfoRequestOption) *GetFileUploadInfoRequest {
	s.Option = v
	return s
}

func (s *GetFileUploadInfoRequest) SetProtocol(v string) *GetFileUploadInfoRequest {
	s.Protocol = &v
	return s
}

func (s *GetFileUploadInfoRequest) SetUnionId(v string) *GetFileUploadInfoRequest {
	s.UnionId = &v
	return s
}

type GetFileUploadInfoRequestOption struct {
	PreCheckParam  *GetFileUploadInfoRequestOptionPreCheckParam `json:"preCheckParam,omitempty" xml:"preCheckParam,omitempty" type:"Struct"`
	PreferIntranet *bool                                        `json:"preferIntranet,omitempty" xml:"preferIntranet,omitempty"`
	PreferRegion   *string                                      `json:"preferRegion,omitempty" xml:"preferRegion,omitempty"`
	StorageDriver  *string                                      `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
}

func (s GetFileUploadInfoRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoRequestOption) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequestOption) SetPreCheckParam(v *GetFileUploadInfoRequestOptionPreCheckParam) *GetFileUploadInfoRequestOption {
	s.PreCheckParam = v
	return s
}

func (s *GetFileUploadInfoRequestOption) SetPreferIntranet(v bool) *GetFileUploadInfoRequestOption {
	s.PreferIntranet = &v
	return s
}

func (s *GetFileUploadInfoRequestOption) SetPreferRegion(v string) *GetFileUploadInfoRequestOption {
	s.PreferRegion = &v
	return s
}

func (s *GetFileUploadInfoRequestOption) SetStorageDriver(v string) *GetFileUploadInfoRequestOption {
	s.StorageDriver = &v
	return s
}

type GetFileUploadInfoRequestOptionPreCheckParam struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Size *int64  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s GetFileUploadInfoRequestOptionPreCheckParam) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoRequestOptionPreCheckParam) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) SetName(v string) *GetFileUploadInfoRequestOptionPreCheckParam {
	s.Name = &v
	return s
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) SetSize(v int64) *GetFileUploadInfoRequestOptionPreCheckParam {
	s.Size = &v
	return s
}

type GetFileUploadInfoResponseBody struct {
	HeaderSignatureInfo *GetFileUploadInfoResponseBodyHeaderSignatureInfo `json:"headerSignatureInfo,omitempty" xml:"headerSignatureInfo,omitempty" type:"Struct"`
	Protocol            *string                                           `json:"protocol,omitempty" xml:"protocol,omitempty"`
	StorageDriver       *string                                           `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	UploadKey           *string                                           `json:"uploadKey,omitempty" xml:"uploadKey,omitempty"`
}

func (s GetFileUploadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponseBody) SetHeaderSignatureInfo(v *GetFileUploadInfoResponseBodyHeaderSignatureInfo) *GetFileUploadInfoResponseBody {
	s.HeaderSignatureInfo = v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetProtocol(v string) *GetFileUploadInfoResponseBody {
	s.Protocol = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetStorageDriver(v string) *GetFileUploadInfoResponseBody {
	s.StorageDriver = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetUploadKey(v string) *GetFileUploadInfoResponseBody {
	s.UploadKey = &v
	return s
}

type GetFileUploadInfoResponseBodyHeaderSignatureInfo struct {
	ExpirationSeconds    *int32             `json:"expirationSeconds,omitempty" xml:"expirationSeconds,omitempty"`
	Headers              map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	InternalResourceUrls []*string          `json:"internalResourceUrls,omitempty" xml:"internalResourceUrls,omitempty" type:"Repeated"`
	Region               *string            `json:"region,omitempty" xml:"region,omitempty"`
	ResourceUrls         []*string          `json:"resourceUrls,omitempty" xml:"resourceUrls,omitempty" type:"Repeated"`
}

func (s GetFileUploadInfoResponseBodyHeaderSignatureInfo) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoResponseBodyHeaderSignatureInfo) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetExpirationSeconds(v int32) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.ExpirationSeconds = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetHeaders(v map[string]*string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.Headers = v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetInternalResourceUrls(v []*string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.InternalResourceUrls = v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetRegion(v string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.Region = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetResourceUrls(v []*string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.ResourceUrls = v
	return s
}

type GetFileUploadInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileUploadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponse) SetHeaders(v map[string]*string) *GetFileUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFileUploadInfoResponse) SetStatusCode(v int32) *GetFileUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileUploadInfoResponse) SetBody(v *GetFileUploadInfoResponseBody) *GetFileUploadInfoResponse {
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

func (client *Client) CommitFileWithOptions(parentDentryUuid *string, request *CommitFileRequest, headers *CommitFileHeaders, runtime *util.RuntimeOptions) (_result *CommitFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.UploadKey)) {
		body["uploadKey"] = request.UploadKey
	}

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
	params := &openapi.Params{
		Action:      tea.String("CommitFile"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/files/" + tea.StringValue(parentDentryUuid) + "/commit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CommitFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CommitFile(parentDentryUuid *string, request *CommitFileRequest) (_result *CommitFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CommitFileHeaders{}
	_result = &CommitFileResponse{}
	_body, _err := client.CommitFileWithOptions(parentDentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileUploadInfoWithOptions(parentDentryUuid *string, request *GetFileUploadInfoRequest, headers *GetFileUploadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["protocol"] = request.Protocol
	}

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
	params := &openapi.Params{
		Action:      tea.String("GetFileUploadInfo"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/files/" + tea.StringValue(parentDentryUuid) + "/uploadInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileUploadInfo(parentDentryUuid *string, request *GetFileUploadInfoRequest) (_result *GetFileUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileUploadInfoHeaders{}
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.GetFileUploadInfoWithOptions(parentDentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
