// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package sns_storage_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ResultItemsDentryAppPropertiesValue struct {
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Value      *string `json:"value,omitempty" xml:"value,omitempty"`
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s ResultItemsDentryAppPropertiesValue) String() string {
	return tea.Prettify(s)
}

func (s ResultItemsDentryAppPropertiesValue) GoString() string {
	return s.String()
}

func (s *ResultItemsDentryAppPropertiesValue) SetName(v string) *ResultItemsDentryAppPropertiesValue {
	s.Name = &v
	return s
}

func (s *ResultItemsDentryAppPropertiesValue) SetValue(v string) *ResultItemsDentryAppPropertiesValue {
	s.Value = &v
	return s
}

func (s *ResultItemsDentryAppPropertiesValue) SetVisibility(v string) *ResultItemsDentryAppPropertiesValue {
	s.Visibility = &v
	return s
}

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

type DentriesAppPropertiesValue struct {
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Value      *string `json:"value,omitempty" xml:"value,omitempty"`
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s DentriesAppPropertiesValue) String() string {
	return tea.Prettify(s)
}

func (s DentriesAppPropertiesValue) GoString() string {
	return s.String()
}

func (s *DentriesAppPropertiesValue) SetName(v string) *DentriesAppPropertiesValue {
	s.Name = &v
	return s
}

func (s *DentriesAppPropertiesValue) SetValue(v string) *DentriesAppPropertiesValue {
	s.Value = &v
	return s
}

func (s *DentriesAppPropertiesValue) SetVisibility(v string) *DentriesAppPropertiesValue {
	s.Visibility = &v
	return s
}

type GetDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesHeaders) GoString() string {
	return s.String()
}

func (s *GetDentriesHeaders) SetCommonHeaders(v map[string]*string) *GetDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *GetDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDentriesRequest struct {
	DentryIds []*string                 `json:"dentryIds,omitempty" xml:"dentryIds,omitempty" type:"Repeated"`
	Option    *GetDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	UnionId   *string                   `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesRequest) GoString() string {
	return s.String()
}

func (s *GetDentriesRequest) SetDentryIds(v []*string) *GetDentriesRequest {
	s.DentryIds = v
	return s
}

func (s *GetDentriesRequest) SetOption(v *GetDentriesRequestOption) *GetDentriesRequest {
	s.Option = v
	return s
}

func (s *GetDentriesRequest) SetUnionId(v string) *GetDentriesRequest {
	s.UnionId = &v
	return s
}

type GetDentriesRequestOption struct {
	AppIdsForAppProperties []*string `json:"appIdsForAppProperties,omitempty" xml:"appIdsForAppProperties,omitempty" type:"Repeated"`
}

func (s GetDentriesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesRequestOption) GoString() string {
	return s.String()
}

func (s *GetDentriesRequestOption) SetAppIdsForAppProperties(v []*string) *GetDentriesRequestOption {
	s.AppIdsForAppProperties = v
	return s
}

type GetDentriesResponseBody struct {
	ResultItems []*GetDentriesResponseBodyResultItems `json:"resultItems,omitempty" xml:"resultItems,omitempty" type:"Repeated"`
}

func (s GetDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDentriesResponseBody) SetResultItems(v []*GetDentriesResponseBodyResultItems) *GetDentriesResponseBody {
	s.ResultItems = v
	return s
}

type GetDentriesResponseBodyResultItems struct {
	Dentry    *GetDentriesResponseBodyResultItemsDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
	DentryId  *string                                   `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	ErrorCode *string                                   `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	SpaceId   *string                                   `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Success   *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDentriesResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *GetDentriesResponseBodyResultItems) SetDentry(v *GetDentriesResponseBodyResultItemsDentry) *GetDentriesResponseBodyResultItems {
	s.Dentry = v
	return s
}

func (s *GetDentriesResponseBodyResultItems) SetDentryId(v string) *GetDentriesResponseBodyResultItems {
	s.DentryId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItems) SetErrorCode(v string) *GetDentriesResponseBodyResultItems {
	s.ErrorCode = &v
	return s
}

func (s *GetDentriesResponseBodyResultItems) SetSpaceId(v string) *GetDentriesResponseBodyResultItems {
	s.SpaceId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItems) SetSuccess(v bool) *GetDentriesResponseBodyResultItems {
	s.Success = &v
	return s
}

type GetDentriesResponseBodyResultItemsDentry struct {
	AppProperties map[string][]*ResultItemsDentryAppPropertiesValue   `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	CreateTime    *string                                             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId     *string                                             `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension     *string                                             `json:"extension,omitempty" xml:"extension,omitempty"`
	Id            *string                                             `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedTime  *string                                             `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId    *string                                             `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name          *string                                             `json:"name,omitempty" xml:"name,omitempty"`
	ParentId      *string                                             `json:"parentId,omitempty" xml:"parentId,omitempty"`
	PartitionType *string                                             `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	Path          *string                                             `json:"path,omitempty" xml:"path,omitempty"`
	Properties    *GetDentriesResponseBodyResultItemsDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	Size          *int64                                              `json:"size,omitempty" xml:"size,omitempty"`
	SpaceId       *string                                             `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status        *string                                             `json:"status,omitempty" xml:"status,omitempty"`
	StorageDriver *string                                             `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	Thumbnail     *GetDentriesResponseBodyResultItemsDentryThumbnail  `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
	Type          *string                                             `json:"type,omitempty" xml:"type,omitempty"`
	Uuid          *string                                             `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version       *int64                                              `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDentriesResponseBodyResultItemsDentry) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponseBodyResultItemsDentry) GoString() string {
	return s.String()
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetAppProperties(v map[string][]*ResultItemsDentryAppPropertiesValue) *GetDentriesResponseBodyResultItemsDentry {
	s.AppProperties = v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetCreateTime(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.CreateTime = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetCreatorId(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.CreatorId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetExtension(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Extension = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetId(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Id = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetModifiedTime(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.ModifiedTime = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetModifierId(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.ModifierId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetName(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Name = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetParentId(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.ParentId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetPartitionType(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.PartitionType = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetPath(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Path = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetProperties(v *GetDentriesResponseBodyResultItemsDentryProperties) *GetDentriesResponseBodyResultItemsDentry {
	s.Properties = v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetSize(v int64) *GetDentriesResponseBodyResultItemsDentry {
	s.Size = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetSpaceId(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.SpaceId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetStatus(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Status = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetStorageDriver(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.StorageDriver = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetThumbnail(v *GetDentriesResponseBodyResultItemsDentryThumbnail) *GetDentriesResponseBodyResultItemsDentry {
	s.Thumbnail = v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetType(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Type = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetUuid(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Uuid = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetVersion(v int64) *GetDentriesResponseBodyResultItemsDentry {
	s.Version = &v
	return s
}

type GetDentriesResponseBodyResultItemsDentryProperties struct {
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s GetDentriesResponseBodyResultItemsDentryProperties) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponseBodyResultItemsDentryProperties) GoString() string {
	return s.String()
}

func (s *GetDentriesResponseBodyResultItemsDentryProperties) SetReadOnly(v bool) *GetDentriesResponseBodyResultItemsDentryProperties {
	s.ReadOnly = &v
	return s
}

type GetDentriesResponseBodyResultItemsDentryThumbnail struct {
	Height *int32  `json:"height,omitempty" xml:"height,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
	Width  *int32  `json:"width,omitempty" xml:"width,omitempty"`
}

func (s GetDentriesResponseBodyResultItemsDentryThumbnail) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponseBodyResultItemsDentryThumbnail) GoString() string {
	return s.String()
}

func (s *GetDentriesResponseBodyResultItemsDentryThumbnail) SetHeight(v int32) *GetDentriesResponseBodyResultItemsDentryThumbnail {
	s.Height = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentryThumbnail) SetUrl(v string) *GetDentriesResponseBodyResultItemsDentryThumbnail {
	s.Url = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentryThumbnail) SetWidth(v int32) *GetDentriesResponseBodyResultItemsDentryThumbnail {
	s.Width = &v
	return s
}

type GetDentriesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponse) GoString() string {
	return s.String()
}

func (s *GetDentriesResponse) SetHeaders(v map[string]*string) *GetDentriesResponse {
	s.Headers = v
	return s
}

func (s *GetDentriesResponse) SetStatusCode(v int32) *GetDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDentriesResponse) SetBody(v *GetDentriesResponseBody) *GetDentriesResponse {
	s.Body = v
	return s
}

type GetDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDentryHeaders) GoString() string {
	return s.String()
}

func (s *GetDentryHeaders) SetCommonHeaders(v map[string]*string) *GetDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDentryHeaders) SetXAcsDingtalkAccessToken(v string) *GetDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDentryRequest struct {
	Option  *GetDentryRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	UnionId *string                 `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDentryRequest) GoString() string {
	return s.String()
}

func (s *GetDentryRequest) SetOption(v *GetDentryRequestOption) *GetDentryRequest {
	s.Option = v
	return s
}

func (s *GetDentryRequest) SetUnionId(v string) *GetDentryRequest {
	s.UnionId = &v
	return s
}

type GetDentryRequestOption struct {
	AppIdsForAppProperties []*string `json:"appIdsForAppProperties,omitempty" xml:"appIdsForAppProperties,omitempty" type:"Repeated"`
	WithThumbnail          *bool     `json:"withThumbnail,omitempty" xml:"withThumbnail,omitempty"`
}

func (s GetDentryRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetDentryRequestOption) GoString() string {
	return s.String()
}

func (s *GetDentryRequestOption) SetAppIdsForAppProperties(v []*string) *GetDentryRequestOption {
	s.AppIdsForAppProperties = v
	return s
}

func (s *GetDentryRequestOption) SetWithThumbnail(v bool) *GetDentryRequestOption {
	s.WithThumbnail = &v
	return s
}

type GetDentryResponseBody struct {
	Dentry *GetDentryResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
}

func (s GetDentryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDentryResponseBody) SetDentry(v *GetDentryResponseBodyDentry) *GetDentryResponseBody {
	s.Dentry = v
	return s
}

type GetDentryResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	CreateTime    *string                                `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId     *string                                `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension     *string                                `json:"extension,omitempty" xml:"extension,omitempty"`
	Id            *string                                `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedTime  *string                                `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId    *string                                `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name          *string                                `json:"name,omitempty" xml:"name,omitempty"`
	ParentId      *string                                `json:"parentId,omitempty" xml:"parentId,omitempty"`
	PartitionType *string                                `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	Path          *string                                `json:"path,omitempty" xml:"path,omitempty"`
	Properties    *GetDentryResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	Size          *int64                                 `json:"size,omitempty" xml:"size,omitempty"`
	SpaceId       *string                                `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status        *string                                `json:"status,omitempty" xml:"status,omitempty"`
	StorageDriver *string                                `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	Thumbnail     *GetDentryResponseBodyDentryThumbnail  `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
	Type          *string                                `json:"type,omitempty" xml:"type,omitempty"`
	Uuid          *string                                `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version       *int64                                 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDentryResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *GetDentryResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *GetDentryResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *GetDentryResponseBodyDentry) SetCreateTime(v string) *GetDentryResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetCreatorId(v string) *GetDentryResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetExtension(v string) *GetDentryResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetId(v string) *GetDentryResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetModifiedTime(v string) *GetDentryResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetModifierId(v string) *GetDentryResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetName(v string) *GetDentryResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetParentId(v string) *GetDentryResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetPartitionType(v string) *GetDentryResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetPath(v string) *GetDentryResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetProperties(v *GetDentryResponseBodyDentryProperties) *GetDentryResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *GetDentryResponseBodyDentry) SetSize(v int64) *GetDentryResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetSpaceId(v string) *GetDentryResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetStatus(v string) *GetDentryResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetStorageDriver(v string) *GetDentryResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetThumbnail(v *GetDentryResponseBodyDentryThumbnail) *GetDentryResponseBodyDentry {
	s.Thumbnail = v
	return s
}

func (s *GetDentryResponseBodyDentry) SetType(v string) *GetDentryResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetUuid(v string) *GetDentryResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetVersion(v int64) *GetDentryResponseBodyDentry {
	s.Version = &v
	return s
}

type GetDentryResponseBodyDentryProperties struct {
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s GetDentryResponseBodyDentryProperties) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *GetDentryResponseBodyDentryProperties) SetReadOnly(v bool) *GetDentryResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

type GetDentryResponseBodyDentryThumbnail struct {
	Height *int32  `json:"height,omitempty" xml:"height,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
	Width  *int32  `json:"width,omitempty" xml:"width,omitempty"`
}

func (s GetDentryResponseBodyDentryThumbnail) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponseBodyDentryThumbnail) GoString() string {
	return s.String()
}

func (s *GetDentryResponseBodyDentryThumbnail) SetHeight(v int32) *GetDentryResponseBodyDentryThumbnail {
	s.Height = &v
	return s
}

func (s *GetDentryResponseBodyDentryThumbnail) SetUrl(v string) *GetDentryResponseBodyDentryThumbnail {
	s.Url = &v
	return s
}

func (s *GetDentryResponseBodyDentryThumbnail) SetWidth(v int32) *GetDentryResponseBodyDentryThumbnail {
	s.Width = &v
	return s
}

type GetDentryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDentryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponse) GoString() string {
	return s.String()
}

func (s *GetDentryResponse) SetHeaders(v map[string]*string) *GetDentryResponse {
	s.Headers = v
	return s
}

func (s *GetDentryResponse) SetStatusCode(v int32) *GetDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDentryResponse) SetBody(v *GetDentryResponseBody) *GetDentryResponse {
	s.Body = v
	return s
}

type GetDentryThumbnailsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDentryThumbnailsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsHeaders) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsHeaders) SetCommonHeaders(v map[string]*string) *GetDentryThumbnailsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDentryThumbnailsHeaders) SetXAcsDingtalkAccessToken(v string) *GetDentryThumbnailsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDentryThumbnailsRequest struct {
	DentryIds []*string `json:"dentryIds,omitempty" xml:"dentryIds,omitempty" type:"Repeated"`
	UnionId   *string   `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetDentryThumbnailsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsRequest) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsRequest) SetDentryIds(v []*string) *GetDentryThumbnailsRequest {
	s.DentryIds = v
	return s
}

func (s *GetDentryThumbnailsRequest) SetUnionId(v string) *GetDentryThumbnailsRequest {
	s.UnionId = &v
	return s
}

type GetDentryThumbnailsResponseBody struct {
	ResultItems []*GetDentryThumbnailsResponseBodyResultItems `json:"resultItems,omitempty" xml:"resultItems,omitempty" type:"Repeated"`
}

func (s GetDentryThumbnailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsResponseBody) SetResultItems(v []*GetDentryThumbnailsResponseBodyResultItems) *GetDentryThumbnailsResponseBody {
	s.ResultItems = v
	return s
}

type GetDentryThumbnailsResponseBodyResultItems struct {
	DentryId  *string                                              `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	ErrorCode *string                                              `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	SpaceId   *string                                              `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Success   *bool                                                `json:"success,omitempty" xml:"success,omitempty"`
	Thumbnail *GetDentryThumbnailsResponseBodyResultItemsThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
}

func (s GetDentryThumbnailsResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsResponseBodyResultItems) SetDentryId(v string) *GetDentryThumbnailsResponseBodyResultItems {
	s.DentryId = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItems) SetErrorCode(v string) *GetDentryThumbnailsResponseBodyResultItems {
	s.ErrorCode = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItems) SetSpaceId(v string) *GetDentryThumbnailsResponseBodyResultItems {
	s.SpaceId = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItems) SetSuccess(v bool) *GetDentryThumbnailsResponseBodyResultItems {
	s.Success = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItems) SetThumbnail(v *GetDentryThumbnailsResponseBodyResultItemsThumbnail) *GetDentryThumbnailsResponseBodyResultItems {
	s.Thumbnail = v
	return s
}

type GetDentryThumbnailsResponseBodyResultItemsThumbnail struct {
	Height *int32  `json:"height,omitempty" xml:"height,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
	Width  *int32  `json:"width,omitempty" xml:"width,omitempty"`
}

func (s GetDentryThumbnailsResponseBodyResultItemsThumbnail) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsResponseBodyResultItemsThumbnail) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsResponseBodyResultItemsThumbnail) SetHeight(v int32) *GetDentryThumbnailsResponseBodyResultItemsThumbnail {
	s.Height = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItemsThumbnail) SetUrl(v string) *GetDentryThumbnailsResponseBodyResultItemsThumbnail {
	s.Url = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItemsThumbnail) SetWidth(v int32) *GetDentryThumbnailsResponseBodyResultItemsThumbnail {
	s.Width = &v
	return s
}

type GetDentryThumbnailsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDentryThumbnailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDentryThumbnailsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsResponse) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsResponse) SetHeaders(v map[string]*string) *GetDentryThumbnailsResponse {
	s.Headers = v
	return s
}

func (s *GetDentryThumbnailsResponse) SetStatusCode(v int32) *GetDentryThumbnailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDentryThumbnailsResponse) SetBody(v *GetDentryThumbnailsResponseBody) *GetDentryThumbnailsResponse {
	s.Body = v
	return s
}

type GetFileDownloadInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFileDownloadInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFileDownloadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileDownloadInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetFileDownloadInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFileDownloadInfoRequest struct {
	Option  *GetFileDownloadInfoRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	UnionId *string                           `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetFileDownloadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoRequest) SetOption(v *GetFileDownloadInfoRequestOption) *GetFileDownloadInfoRequest {
	s.Option = v
	return s
}

func (s *GetFileDownloadInfoRequest) SetUnionId(v string) *GetFileDownloadInfoRequest {
	s.UnionId = &v
	return s
}

type GetFileDownloadInfoRequestOption struct {
	PreferIntranet *bool  `json:"preferIntranet,omitempty" xml:"preferIntranet,omitempty"`
	Version        *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetFileDownloadInfoRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoRequestOption) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoRequestOption) SetPreferIntranet(v bool) *GetFileDownloadInfoRequestOption {
	s.PreferIntranet = &v
	return s
}

func (s *GetFileDownloadInfoRequestOption) SetVersion(v int64) *GetFileDownloadInfoRequestOption {
	s.Version = &v
	return s
}

type GetFileDownloadInfoResponseBody struct {
	HeaderSignatureInfo *GetFileDownloadInfoResponseBodyHeaderSignatureInfo `json:"headerSignatureInfo,omitempty" xml:"headerSignatureInfo,omitempty" type:"Struct"`
	Protocol            *string                                             `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GetFileDownloadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoResponseBody) SetHeaderSignatureInfo(v *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) *GetFileDownloadInfoResponseBody {
	s.HeaderSignatureInfo = v
	return s
}

func (s *GetFileDownloadInfoResponseBody) SetProtocol(v string) *GetFileDownloadInfoResponseBody {
	s.Protocol = &v
	return s
}

type GetFileDownloadInfoResponseBodyHeaderSignatureInfo struct {
	ExpirationSeconds    *int32             `json:"expirationSeconds,omitempty" xml:"expirationSeconds,omitempty"`
	Headers              map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	InternalResourceUrls []*string          `json:"internalResourceUrls,omitempty" xml:"internalResourceUrls,omitempty" type:"Repeated"`
	Region               *string            `json:"region,omitempty" xml:"region,omitempty"`
	ResourceUrls         []*string          `json:"resourceUrls,omitempty" xml:"resourceUrls,omitempty" type:"Repeated"`
}

func (s GetFileDownloadInfoResponseBodyHeaderSignatureInfo) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoResponseBodyHeaderSignatureInfo) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetExpirationSeconds(v int32) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.ExpirationSeconds = &v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetHeaders(v map[string]*string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.Headers = v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetInternalResourceUrls(v []*string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.InternalResourceUrls = v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetRegion(v string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.Region = &v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetResourceUrls(v []*string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.ResourceUrls = v
	return s
}

type GetFileDownloadInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileDownloadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoResponse) SetHeaders(v map[string]*string) *GetFileDownloadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFileDownloadInfoResponse) SetStatusCode(v int32) *GetFileDownloadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileDownloadInfoResponse) SetBody(v *GetFileDownloadInfoResponseBody) *GetFileDownloadInfoResponse {
	s.Body = v
	return s
}

type GetSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceHeaders) GoString() string {
	return s.String()
}

func (s *GetSpaceHeaders) SetCommonHeaders(v map[string]*string) *GetSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *GetSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSpaceRequest struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UnionId            *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceRequest) SetOpenConversationId(v string) *GetSpaceRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetSpaceRequest) SetUnionId(v string) *GetSpaceRequest {
	s.UnionId = &v
	return s
}

type GetSpaceResponseBody struct {
	Space *GetSpaceResponseBodySpace `json:"space,omitempty" xml:"space,omitempty" type:"Struct"`
}

func (s GetSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceResponseBody) SetSpace(v *GetSpaceResponseBodySpace) *GetSpaceResponseBody {
	s.Space = v
	return s
}

type GetSpaceResponseBodySpace struct {
	CorpId       *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CreateTime   *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	SpaceId      *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetSpaceResponseBodySpace) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBodySpace) GoString() string {
	return s.String()
}

func (s *GetSpaceResponseBodySpace) SetCorpId(v string) *GetSpaceResponseBodySpace {
	s.CorpId = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetCreateTime(v string) *GetSpaceResponseBodySpace {
	s.CreateTime = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetModifiedTime(v string) *GetSpaceResponseBodySpace {
	s.ModifiedTime = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetSpaceId(v string) *GetSpaceResponseBodySpace {
	s.SpaceId = &v
	return s
}

type GetSpaceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceResponse) SetHeaders(v map[string]*string) *GetSpaceResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceResponse) SetStatusCode(v int32) *GetSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpaceResponse) SetBody(v *GetSpaceResponseBody) *GetSpaceResponse {
	s.Body = v
	return s
}

type ListAllDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAllDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesHeaders) GoString() string {
	return s.String()
}

func (s *ListAllDentriesHeaders) SetCommonHeaders(v map[string]*string) *ListAllDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAllDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *ListAllDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAllDentriesRequest struct {
	Option  *ListAllDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	UnionId *string                       `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListAllDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesRequest) GoString() string {
	return s.String()
}

func (s *ListAllDentriesRequest) SetOption(v *ListAllDentriesRequestOption) *ListAllDentriesRequest {
	s.Option = v
	return s
}

func (s *ListAllDentriesRequest) SetUnionId(v string) *ListAllDentriesRequest {
	s.UnionId = &v
	return s
}

type ListAllDentriesRequestOption struct {
	MaxResults    *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken     *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Order         *string `json:"order,omitempty" xml:"order,omitempty"`
	WithThumbnail *bool   `json:"withThumbnail,omitempty" xml:"withThumbnail,omitempty"`
}

func (s ListAllDentriesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesRequestOption) GoString() string {
	return s.String()
}

func (s *ListAllDentriesRequestOption) SetMaxResults(v int32) *ListAllDentriesRequestOption {
	s.MaxResults = &v
	return s
}

func (s *ListAllDentriesRequestOption) SetNextToken(v string) *ListAllDentriesRequestOption {
	s.NextToken = &v
	return s
}

func (s *ListAllDentriesRequestOption) SetOrder(v string) *ListAllDentriesRequestOption {
	s.Order = &v
	return s
}

func (s *ListAllDentriesRequestOption) SetWithThumbnail(v bool) *ListAllDentriesRequestOption {
	s.WithThumbnail = &v
	return s
}

type ListAllDentriesResponseBody struct {
	Dentries  []*ListAllDentriesResponseBodyDentries `json:"dentries,omitempty" xml:"dentries,omitempty" type:"Repeated"`
	NextToken *string                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAllDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllDentriesResponseBody) SetDentries(v []*ListAllDentriesResponseBodyDentries) *ListAllDentriesResponseBody {
	s.Dentries = v
	return s
}

func (s *ListAllDentriesResponseBody) SetNextToken(v string) *ListAllDentriesResponseBody {
	s.NextToken = &v
	return s
}

type ListAllDentriesResponseBodyDentries struct {
	AppProperties map[string][]*DentriesAppPropertiesValue       `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	CreateTime    *string                                        `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId     *string                                        `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension     *string                                        `json:"extension,omitempty" xml:"extension,omitempty"`
	Id            *string                                        `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedTime  *string                                        `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId    *string                                        `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name          *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	ParentId      *string                                        `json:"parentId,omitempty" xml:"parentId,omitempty"`
	PartitionType *string                                        `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	Path          *string                                        `json:"path,omitempty" xml:"path,omitempty"`
	Properties    *ListAllDentriesResponseBodyDentriesProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	Size          *int64                                         `json:"size,omitempty" xml:"size,omitempty"`
	SpaceId       *string                                        `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status        *string                                        `json:"status,omitempty" xml:"status,omitempty"`
	StorageDriver *string                                        `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	Thumbnail     *ListAllDentriesResponseBodyDentriesThumbnail  `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
	Type          *string                                        `json:"type,omitempty" xml:"type,omitempty"`
	Uuid          *string                                        `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version       *int64                                         `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListAllDentriesResponseBodyDentries) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesResponseBodyDentries) GoString() string {
	return s.String()
}

func (s *ListAllDentriesResponseBodyDentries) SetAppProperties(v map[string][]*DentriesAppPropertiesValue) *ListAllDentriesResponseBodyDentries {
	s.AppProperties = v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetCreateTime(v string) *ListAllDentriesResponseBodyDentries {
	s.CreateTime = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetCreatorId(v string) *ListAllDentriesResponseBodyDentries {
	s.CreatorId = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetExtension(v string) *ListAllDentriesResponseBodyDentries {
	s.Extension = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetId(v string) *ListAllDentriesResponseBodyDentries {
	s.Id = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetModifiedTime(v string) *ListAllDentriesResponseBodyDentries {
	s.ModifiedTime = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetModifierId(v string) *ListAllDentriesResponseBodyDentries {
	s.ModifierId = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetName(v string) *ListAllDentriesResponseBodyDentries {
	s.Name = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetParentId(v string) *ListAllDentriesResponseBodyDentries {
	s.ParentId = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetPartitionType(v string) *ListAllDentriesResponseBodyDentries {
	s.PartitionType = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetPath(v string) *ListAllDentriesResponseBodyDentries {
	s.Path = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetProperties(v *ListAllDentriesResponseBodyDentriesProperties) *ListAllDentriesResponseBodyDentries {
	s.Properties = v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetSize(v int64) *ListAllDentriesResponseBodyDentries {
	s.Size = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetSpaceId(v string) *ListAllDentriesResponseBodyDentries {
	s.SpaceId = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetStatus(v string) *ListAllDentriesResponseBodyDentries {
	s.Status = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetStorageDriver(v string) *ListAllDentriesResponseBodyDentries {
	s.StorageDriver = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetThumbnail(v *ListAllDentriesResponseBodyDentriesThumbnail) *ListAllDentriesResponseBodyDentries {
	s.Thumbnail = v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetType(v string) *ListAllDentriesResponseBodyDentries {
	s.Type = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetUuid(v string) *ListAllDentriesResponseBodyDentries {
	s.Uuid = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetVersion(v int64) *ListAllDentriesResponseBodyDentries {
	s.Version = &v
	return s
}

type ListAllDentriesResponseBodyDentriesProperties struct {
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s ListAllDentriesResponseBodyDentriesProperties) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesResponseBodyDentriesProperties) GoString() string {
	return s.String()
}

func (s *ListAllDentriesResponseBodyDentriesProperties) SetReadOnly(v bool) *ListAllDentriesResponseBodyDentriesProperties {
	s.ReadOnly = &v
	return s
}

type ListAllDentriesResponseBodyDentriesThumbnail struct {
	Height *int32  `json:"height,omitempty" xml:"height,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
	Width  *int32  `json:"width,omitempty" xml:"width,omitempty"`
}

func (s ListAllDentriesResponseBodyDentriesThumbnail) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesResponseBodyDentriesThumbnail) GoString() string {
	return s.String()
}

func (s *ListAllDentriesResponseBodyDentriesThumbnail) SetHeight(v int32) *ListAllDentriesResponseBodyDentriesThumbnail {
	s.Height = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentriesThumbnail) SetUrl(v string) *ListAllDentriesResponseBodyDentriesThumbnail {
	s.Url = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentriesThumbnail) SetWidth(v int32) *ListAllDentriesResponseBodyDentriesThumbnail {
	s.Width = &v
	return s
}

type ListAllDentriesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAllDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAllDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesResponse) GoString() string {
	return s.String()
}

func (s *ListAllDentriesResponse) SetHeaders(v map[string]*string) *ListAllDentriesResponse {
	s.Headers = v
	return s
}

func (s *ListAllDentriesResponse) SetStatusCode(v int32) *ListAllDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllDentriesResponse) SetBody(v *ListAllDentriesResponseBody) *ListAllDentriesResponse {
	s.Body = v
	return s
}

type ListDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesHeaders) GoString() string {
	return s.String()
}

func (s *ListDentriesHeaders) SetCommonHeaders(v map[string]*string) *ListDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *ListDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListDentriesRequest struct {
	MaxResults    *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken     *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Order         *string `json:"order,omitempty" xml:"order,omitempty"`
	OrderBy       *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	ParentId      *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	UnionId       *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	WithThumbnail *bool   `json:"withThumbnail,omitempty" xml:"withThumbnail,omitempty"`
}

func (s ListDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesRequest) GoString() string {
	return s.String()
}

func (s *ListDentriesRequest) SetMaxResults(v int32) *ListDentriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDentriesRequest) SetNextToken(v string) *ListDentriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDentriesRequest) SetOrder(v string) *ListDentriesRequest {
	s.Order = &v
	return s
}

func (s *ListDentriesRequest) SetOrderBy(v string) *ListDentriesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDentriesRequest) SetParentId(v string) *ListDentriesRequest {
	s.ParentId = &v
	return s
}

func (s *ListDentriesRequest) SetUnionId(v string) *ListDentriesRequest {
	s.UnionId = &v
	return s
}

func (s *ListDentriesRequest) SetWithThumbnail(v bool) *ListDentriesRequest {
	s.WithThumbnail = &v
	return s
}

type ListDentriesResponseBody struct {
	Dentries  []*ListDentriesResponseBodyDentries `json:"dentries,omitempty" xml:"dentries,omitempty" type:"Repeated"`
	NextToken *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBody) SetDentries(v []*ListDentriesResponseBodyDentries) *ListDentriesResponseBody {
	s.Dentries = v
	return s
}

func (s *ListDentriesResponseBody) SetNextToken(v string) *ListDentriesResponseBody {
	s.NextToken = &v
	return s
}

type ListDentriesResponseBodyDentries struct {
	AppProperties map[string][]*DentriesAppPropertiesValue    `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	CreateTime    *string                                     `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId     *string                                     `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension     *string                                     `json:"extension,omitempty" xml:"extension,omitempty"`
	Id            *string                                     `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedTime  *string                                     `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId    *string                                     `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name          *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	ParentId      *string                                     `json:"parentId,omitempty" xml:"parentId,omitempty"`
	PartitionType *string                                     `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	Path          *string                                     `json:"path,omitempty" xml:"path,omitempty"`
	Properties    *ListDentriesResponseBodyDentriesProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	Size          *int64                                      `json:"size,omitempty" xml:"size,omitempty"`
	SpaceId       *string                                     `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status        *string                                     `json:"status,omitempty" xml:"status,omitempty"`
	StorageDriver *string                                     `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	Thumbnail     *ListDentriesResponseBodyDentriesThumbnail  `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
	Type          *string                                     `json:"type,omitempty" xml:"type,omitempty"`
	Uuid          *string                                     `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version       *int64                                      `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListDentriesResponseBodyDentries) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponseBodyDentries) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBodyDentries) SetAppProperties(v map[string][]*DentriesAppPropertiesValue) *ListDentriesResponseBodyDentries {
	s.AppProperties = v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetCreateTime(v string) *ListDentriesResponseBodyDentries {
	s.CreateTime = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetCreatorId(v string) *ListDentriesResponseBodyDentries {
	s.CreatorId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetExtension(v string) *ListDentriesResponseBodyDentries {
	s.Extension = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetId(v string) *ListDentriesResponseBodyDentries {
	s.Id = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetModifiedTime(v string) *ListDentriesResponseBodyDentries {
	s.ModifiedTime = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetModifierId(v string) *ListDentriesResponseBodyDentries {
	s.ModifierId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetName(v string) *ListDentriesResponseBodyDentries {
	s.Name = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetParentId(v string) *ListDentriesResponseBodyDentries {
	s.ParentId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetPartitionType(v string) *ListDentriesResponseBodyDentries {
	s.PartitionType = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetPath(v string) *ListDentriesResponseBodyDentries {
	s.Path = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetProperties(v *ListDentriesResponseBodyDentriesProperties) *ListDentriesResponseBodyDentries {
	s.Properties = v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetSize(v int64) *ListDentriesResponseBodyDentries {
	s.Size = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetSpaceId(v string) *ListDentriesResponseBodyDentries {
	s.SpaceId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetStatus(v string) *ListDentriesResponseBodyDentries {
	s.Status = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetStorageDriver(v string) *ListDentriesResponseBodyDentries {
	s.StorageDriver = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetThumbnail(v *ListDentriesResponseBodyDentriesThumbnail) *ListDentriesResponseBodyDentries {
	s.Thumbnail = v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetType(v string) *ListDentriesResponseBodyDentries {
	s.Type = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetUuid(v string) *ListDentriesResponseBodyDentries {
	s.Uuid = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetVersion(v int64) *ListDentriesResponseBodyDentries {
	s.Version = &v
	return s
}

type ListDentriesResponseBodyDentriesProperties struct {
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s ListDentriesResponseBodyDentriesProperties) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponseBodyDentriesProperties) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBodyDentriesProperties) SetReadOnly(v bool) *ListDentriesResponseBodyDentriesProperties {
	s.ReadOnly = &v
	return s
}

type ListDentriesResponseBodyDentriesThumbnail struct {
	Height *int32  `json:"height,omitempty" xml:"height,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
	Width  *int32  `json:"width,omitempty" xml:"width,omitempty"`
}

func (s ListDentriesResponseBodyDentriesThumbnail) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponseBodyDentriesThumbnail) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBodyDentriesThumbnail) SetHeight(v int32) *ListDentriesResponseBodyDentriesThumbnail {
	s.Height = &v
	return s
}

func (s *ListDentriesResponseBodyDentriesThumbnail) SetUrl(v string) *ListDentriesResponseBodyDentriesThumbnail {
	s.Url = &v
	return s
}

func (s *ListDentriesResponseBodyDentriesThumbnail) SetWidth(v int32) *ListDentriesResponseBodyDentriesThumbnail {
	s.Width = &v
	return s
}

type ListDentriesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponse) GoString() string {
	return s.String()
}

func (s *ListDentriesResponse) SetHeaders(v map[string]*string) *ListDentriesResponse {
	s.Headers = v
	return s
}

func (s *ListDentriesResponse) SetStatusCode(v int32) *ListDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDentriesResponse) SetBody(v *ListDentriesResponseBody) *ListDentriesResponse {
	s.Body = v
	return s
}

type ListExpiredHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListExpiredHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredHeaders) GoString() string {
	return s.String()
}

func (s *ListExpiredHeaders) SetCommonHeaders(v map[string]*string) *ListExpiredHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListExpiredHeaders) SetXAcsDingtalkAccessToken(v string) *ListExpiredHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListExpiredRequest struct {
	OpenConversationId *string                   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Option             *ListExpiredRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	UnionId            *string                   `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListExpiredRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredRequest) GoString() string {
	return s.String()
}

func (s *ListExpiredRequest) SetOpenConversationId(v string) *ListExpiredRequest {
	s.OpenConversationId = &v
	return s
}

func (s *ListExpiredRequest) SetOption(v *ListExpiredRequestOption) *ListExpiredRequest {
	s.Option = v
	return s
}

func (s *ListExpiredRequest) SetUnionId(v string) *ListExpiredRequest {
	s.UnionId = &v
	return s
}

type ListExpiredRequestOption struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListExpiredRequestOption) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredRequestOption) GoString() string {
	return s.String()
}

func (s *ListExpiredRequestOption) SetMaxResults(v int32) *ListExpiredRequestOption {
	s.MaxResults = &v
	return s
}

func (s *ListExpiredRequestOption) SetNextToken(v string) *ListExpiredRequestOption {
	s.NextToken = &v
	return s
}

type ListExpiredResponseBody struct {
	Files     []*ListExpiredResponseBodyFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	NextToken *string                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListExpiredResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredResponseBody) GoString() string {
	return s.String()
}

func (s *ListExpiredResponseBody) SetFiles(v []*ListExpiredResponseBodyFiles) *ListExpiredResponseBody {
	s.Files = v
	return s
}

func (s *ListExpiredResponseBody) SetNextToken(v string) *ListExpiredResponseBody {
	s.NextToken = &v
	return s
}

type ListExpiredResponseBodyFiles struct {
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	CreateTime     *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId      *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension      *string `json:"extension,omitempty" xml:"extension,omitempty"`
	Id             *string `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedTime   *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId     *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentId       *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
	Size           *int64  `json:"size,omitempty" xml:"size,omitempty"`
	SpaceId        *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
	Uuid           *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version        *int64  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListExpiredResponseBodyFiles) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *ListExpiredResponseBodyFiles) SetConversationId(v string) *ListExpiredResponseBodyFiles {
	s.ConversationId = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetCreateTime(v string) *ListExpiredResponseBodyFiles {
	s.CreateTime = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetCreatorId(v string) *ListExpiredResponseBodyFiles {
	s.CreatorId = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetExtension(v string) *ListExpiredResponseBodyFiles {
	s.Extension = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetId(v string) *ListExpiredResponseBodyFiles {
	s.Id = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetModifiedTime(v string) *ListExpiredResponseBodyFiles {
	s.ModifiedTime = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetModifierId(v string) *ListExpiredResponseBodyFiles {
	s.ModifierId = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetName(v string) *ListExpiredResponseBodyFiles {
	s.Name = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetParentId(v string) *ListExpiredResponseBodyFiles {
	s.ParentId = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetPath(v string) *ListExpiredResponseBodyFiles {
	s.Path = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetSize(v int64) *ListExpiredResponseBodyFiles {
	s.Size = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetSpaceId(v string) *ListExpiredResponseBodyFiles {
	s.SpaceId = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetStatus(v string) *ListExpiredResponseBodyFiles {
	s.Status = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetType(v string) *ListExpiredResponseBodyFiles {
	s.Type = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetUuid(v string) *ListExpiredResponseBodyFiles {
	s.Uuid = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetVersion(v int64) *ListExpiredResponseBodyFiles {
	s.Version = &v
	return s
}

type ListExpiredResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListExpiredResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExpiredResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredResponse) GoString() string {
	return s.String()
}

func (s *ListExpiredResponse) SetHeaders(v map[string]*string) *ListExpiredResponse {
	s.Headers = v
	return s
}

func (s *ListExpiredResponse) SetStatusCode(v int32) *ListExpiredResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExpiredResponse) SetBody(v *ListExpiredResponseBody) *ListExpiredResponse {
	s.Body = v
	return s
}

type SubscribeEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SubscribeEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubscribeEventHeaders) GoString() string {
	return s.String()
}

func (s *SubscribeEventHeaders) SetCommonHeaders(v map[string]*string) *SubscribeEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubscribeEventHeaders) SetXAcsDingtalkAccessToken(v string) *SubscribeEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SubscribeEventRequest struct {
	Scope   *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeId *string `json:"scopeId,omitempty" xml:"scopeId,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SubscribeEventRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeEventRequest) GoString() string {
	return s.String()
}

func (s *SubscribeEventRequest) SetScope(v string) *SubscribeEventRequest {
	s.Scope = &v
	return s
}

func (s *SubscribeEventRequest) SetScopeId(v string) *SubscribeEventRequest {
	s.ScopeId = &v
	return s
}

func (s *SubscribeEventRequest) SetUnionId(v string) *SubscribeEventRequest {
	s.UnionId = &v
	return s
}

type SubscribeEventResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubscribeEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscribeEventResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeEventResponseBody) SetSuccess(v bool) *SubscribeEventResponseBody {
	s.Success = &v
	return s
}

type SubscribeEventResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubscribeEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubscribeEventResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeEventResponse) GoString() string {
	return s.String()
}

func (s *SubscribeEventResponse) SetHeaders(v map[string]*string) *SubscribeEventResponse {
	s.Headers = v
	return s
}

func (s *SubscribeEventResponse) SetStatusCode(v int32) *SubscribeEventResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeEventResponse) SetBody(v *SubscribeEventResponseBody) *SubscribeEventResponse {
	s.Body = v
	return s
}

type UnsubscribeEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UnsubscribeEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeEventHeaders) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventHeaders) SetCommonHeaders(v map[string]*string) *UnsubscribeEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnsubscribeEventHeaders) SetXAcsDingtalkAccessToken(v string) *UnsubscribeEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UnsubscribeEventRequest struct {
	Scope   *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeId *string `json:"scopeId,omitempty" xml:"scopeId,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UnsubscribeEventRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeEventRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventRequest) SetScope(v string) *UnsubscribeEventRequest {
	s.Scope = &v
	return s
}

func (s *UnsubscribeEventRequest) SetScopeId(v string) *UnsubscribeEventRequest {
	s.ScopeId = &v
	return s
}

func (s *UnsubscribeEventRequest) SetUnionId(v string) *UnsubscribeEventRequest {
	s.UnionId = &v
	return s
}

type UnsubscribeEventResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UnsubscribeEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeEventResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventResponseBody) SetSuccess(v bool) *UnsubscribeEventResponseBody {
	s.Success = &v
	return s
}

type UnsubscribeEventResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnsubscribeEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnsubscribeEventResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeEventResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventResponse) SetHeaders(v map[string]*string) *UnsubscribeEventResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeEventResponse) SetStatusCode(v int32) *UnsubscribeEventResponse {
	s.StatusCode = &v
	return s
}

func (s *UnsubscribeEventResponse) SetBody(v *UnsubscribeEventResponseBody) *UnsubscribeEventResponse {
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

func (client *Client) GetDentriesWithOptions(spaceId *string, request *GetDentriesRequest, headers *GetDentriesHeaders, runtime *util.RuntimeOptions) (_result *GetDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryIds)) {
		body["dentryIds"] = request.DentryIds
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

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
		Action:      tea.String("GetDentries"),
		Version:     tea.String("snsStorage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/snsStorage/spaces/" + tea.StringValue(spaceId) + "/dentries/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDentries(spaceId *string, request *GetDentriesRequest) (_result *GetDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDentriesHeaders{}
	_result = &GetDentriesResponse{}
	_body, _err := client.GetDentriesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDentryWithOptions(spaceId *string, dentryId *string, request *GetDentryRequest, headers *GetDentryHeaders, runtime *util.RuntimeOptions) (_result *GetDentryResponse, _err error) {
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
		Action:      tea.String("GetDentry"),
		Version:     tea.String("snsStorage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/snsStorage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDentry(spaceId *string, dentryId *string, request *GetDentryRequest) (_result *GetDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDentryHeaders{}
	_result = &GetDentryResponse{}
	_body, _err := client.GetDentryWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDentryThumbnailsWithOptions(spaceId *string, request *GetDentryThumbnailsRequest, headers *GetDentryThumbnailsHeaders, runtime *util.RuntimeOptions) (_result *GetDentryThumbnailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryIds)) {
		body["dentryIds"] = request.DentryIds
	}

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
		Action:      tea.String("GetDentryThumbnails"),
		Version:     tea.String("snsStorage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/snsStorage/spaces/" + tea.StringValue(spaceId) + "/thumbnails/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDentryThumbnailsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDentryThumbnails(spaceId *string, request *GetDentryThumbnailsRequest) (_result *GetDentryThumbnailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDentryThumbnailsHeaders{}
	_result = &GetDentryThumbnailsResponse{}
	_body, _err := client.GetDentryThumbnailsWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileDownloadInfoWithOptions(spaceId *string, dentryId *string, request *GetFileDownloadInfoRequest, headers *GetFileDownloadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileDownloadInfoResponse, _err error) {
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
		Action:      tea.String("GetFileDownloadInfo"),
		Version:     tea.String("snsStorage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/snsStorage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/downloadInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileDownloadInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileDownloadInfo(spaceId *string, dentryId *string, request *GetFileDownloadInfoRequest) (_result *GetFileDownloadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileDownloadInfoHeaders{}
	_result = &GetFileDownloadInfoResponse{}
	_body, _err := client.GetFileDownloadInfoWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpaceWithOptions(request *GetSpaceRequest, headers *GetSpaceHeaders, runtime *util.RuntimeOptions) (_result *GetSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

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
		Action:      tea.String("GetSpace"),
		Version:     tea.String("snsStorage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/snsStorage/conversations/spaces/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpace(request *GetSpaceRequest) (_result *GetSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSpaceHeaders{}
	_result = &GetSpaceResponse{}
	_body, _err := client.GetSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAllDentriesWithOptions(spaceId *string, request *ListAllDentriesRequest, headers *ListAllDentriesHeaders, runtime *util.RuntimeOptions) (_result *ListAllDentriesResponse, _err error) {
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
		Action:      tea.String("ListAllDentries"),
		Version:     tea.String("snsStorage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/snsStorage/spaces/" + tea.StringValue(spaceId) + "/dentries/listAll"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAllDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAllDentries(spaceId *string, request *ListAllDentriesRequest) (_result *ListAllDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAllDentriesHeaders{}
	_result = &ListAllDentriesResponse{}
	_body, _err := client.ListAllDentriesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDentriesWithOptions(spaceId *string, request *ListDentriesRequest, headers *ListDentriesHeaders, runtime *util.RuntimeOptions) (_result *ListDentriesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["parentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.WithThumbnail)) {
		query["withThumbnail"] = request.WithThumbnail
	}

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
		Action:      tea.String("ListDentries"),
		Version:     tea.String("snsStorage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/snsStorage/spaces/" + tea.StringValue(spaceId) + "/dentries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDentries(spaceId *string, request *ListDentriesRequest) (_result *ListDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDentriesHeaders{}
	_result = &ListDentriesResponse{}
	_body, _err := client.ListDentriesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExpiredWithOptions(request *ListExpiredRequest, headers *ListExpiredHeaders, runtime *util.RuntimeOptions) (_result *ListExpiredResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

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
		Action:      tea.String("ListExpired"),
		Version:     tea.String("snsStorage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/snsStorage/conversations/expiredFileLists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExpiredResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExpired(request *ListExpiredRequest) (_result *ListExpiredResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListExpiredHeaders{}
	_result = &ListExpiredResponse{}
	_body, _err := client.ListExpiredWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubscribeEventWithOptions(request *SubscribeEventRequest, headers *SubscribeEventHeaders, runtime *util.RuntimeOptions) (_result *SubscribeEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeId)) {
		body["scopeId"] = request.ScopeId
	}

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
		Action:      tea.String("SubscribeEvent"),
		Version:     tea.String("snsStorage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/snsStorage/events/subscribe"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SubscribeEventResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubscribeEvent(request *SubscribeEventRequest) (_result *SubscribeEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubscribeEventHeaders{}
	_result = &SubscribeEventResponse{}
	_body, _err := client.SubscribeEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnsubscribeEventWithOptions(request *UnsubscribeEventRequest, headers *UnsubscribeEventHeaders, runtime *util.RuntimeOptions) (_result *UnsubscribeEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeId)) {
		body["scopeId"] = request.ScopeId
	}

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
		Action:      tea.String("UnsubscribeEvent"),
		Version:     tea.String("snsStorage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/snsStorage/events/unsubscribe"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UnsubscribeEventResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnsubscribeEvent(request *UnsubscribeEventRequest) (_result *UnsubscribeEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnsubscribeEventHeaders{}
	_result = &UnsubscribeEventResponse{}
	_body, _err := client.UnsubscribeEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
