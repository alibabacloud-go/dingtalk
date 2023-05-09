// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package doc_2_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DentryModel struct {
	ContentType     *string                     `json:"contentType,omitempty" xml:"contentType,omitempty"`
	CreatedTime     *int64                      `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Creator         *DentryModelCreator         `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	DentryId        *string                     `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	DentryType      *string                     `json:"dentryType,omitempty" xml:"dentryType,omitempty"`
	DentryUuid      *string                     `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	DocKey          *string                     `json:"docKey,omitempty" xml:"docKey,omitempty"`
	Extension       *string                     `json:"extension,omitempty" xml:"extension,omitempty"`
	HasChildren     *bool                       `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	LinkSourceInfo  *LinkSourceInfo             `json:"linkSourceInfo,omitempty" xml:"linkSourceInfo,omitempty"`
	Name            *string                     `json:"name,omitempty" xml:"name,omitempty"`
	Path            *string                     `json:"path,omitempty" xml:"path,omitempty"`
	Space           *SpaceModel                 `json:"space,omitempty" xml:"space,omitempty"`
	SpaceId         *string                     `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	StatisticalInfo *DentryModelStatisticalInfo `json:"statisticalInfo,omitempty" xml:"statisticalInfo,omitempty" type:"Struct"`
	UpdatedTime     *int64                      `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	Updater         *DentryModelUpdater         `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	Url             *string                     `json:"url,omitempty" xml:"url,omitempty"`
	VisitorInfo     *DentryModelVisitorInfo     `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
}

func (s DentryModel) String() string {
	return tea.Prettify(s)
}

func (s DentryModel) GoString() string {
	return s.String()
}

func (s *DentryModel) SetContentType(v string) *DentryModel {
	s.ContentType = &v
	return s
}

func (s *DentryModel) SetCreatedTime(v int64) *DentryModel {
	s.CreatedTime = &v
	return s
}

func (s *DentryModel) SetCreator(v *DentryModelCreator) *DentryModel {
	s.Creator = v
	return s
}

func (s *DentryModel) SetDentryId(v string) *DentryModel {
	s.DentryId = &v
	return s
}

func (s *DentryModel) SetDentryType(v string) *DentryModel {
	s.DentryType = &v
	return s
}

func (s *DentryModel) SetDentryUuid(v string) *DentryModel {
	s.DentryUuid = &v
	return s
}

func (s *DentryModel) SetDocKey(v string) *DentryModel {
	s.DocKey = &v
	return s
}

func (s *DentryModel) SetExtension(v string) *DentryModel {
	s.Extension = &v
	return s
}

func (s *DentryModel) SetHasChildren(v bool) *DentryModel {
	s.HasChildren = &v
	return s
}

func (s *DentryModel) SetLinkSourceInfo(v *LinkSourceInfo) *DentryModel {
	s.LinkSourceInfo = v
	return s
}

func (s *DentryModel) SetName(v string) *DentryModel {
	s.Name = &v
	return s
}

func (s *DentryModel) SetPath(v string) *DentryModel {
	s.Path = &v
	return s
}

func (s *DentryModel) SetSpace(v *SpaceModel) *DentryModel {
	s.Space = v
	return s
}

func (s *DentryModel) SetSpaceId(v string) *DentryModel {
	s.SpaceId = &v
	return s
}

func (s *DentryModel) SetStatisticalInfo(v *DentryModelStatisticalInfo) *DentryModel {
	s.StatisticalInfo = v
	return s
}

func (s *DentryModel) SetUpdatedTime(v int64) *DentryModel {
	s.UpdatedTime = &v
	return s
}

func (s *DentryModel) SetUpdater(v *DentryModelUpdater) *DentryModel {
	s.Updater = v
	return s
}

func (s *DentryModel) SetUrl(v string) *DentryModel {
	s.Url = &v
	return s
}

func (s *DentryModel) SetVisitorInfo(v *DentryModelVisitorInfo) *DentryModel {
	s.VisitorInfo = v
	return s
}

type DentryModelCreator struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DentryModelCreator) String() string {
	return tea.Prettify(s)
}

func (s DentryModelCreator) GoString() string {
	return s.String()
}

func (s *DentryModelCreator) SetName(v string) *DentryModelCreator {
	s.Name = &v
	return s
}

func (s *DentryModelCreator) SetUnionId(v string) *DentryModelCreator {
	s.UnionId = &v
	return s
}

type DentryModelStatisticalInfo struct {
	WordCount *int64 `json:"wordCount,omitempty" xml:"wordCount,omitempty"`
}

func (s DentryModelStatisticalInfo) String() string {
	return tea.Prettify(s)
}

func (s DentryModelStatisticalInfo) GoString() string {
	return s.String()
}

func (s *DentryModelStatisticalInfo) SetWordCount(v int64) *DentryModelStatisticalInfo {
	s.WordCount = &v
	return s
}

type DentryModelUpdater struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DentryModelUpdater) String() string {
	return tea.Prettify(s)
}

func (s DentryModelUpdater) GoString() string {
	return s.String()
}

func (s *DentryModelUpdater) SetName(v string) *DentryModelUpdater {
	s.Name = &v
	return s
}

func (s *DentryModelUpdater) SetUnionId(v string) *DentryModelUpdater {
	s.UnionId = &v
	return s
}

type DentryModelVisitorInfo struct {
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	RoleCode      *string   `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	SpaceActions  []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
}

func (s DentryModelVisitorInfo) String() string {
	return tea.Prettify(s)
}

func (s DentryModelVisitorInfo) GoString() string {
	return s.String()
}

func (s *DentryModelVisitorInfo) SetDentryActions(v []*string) *DentryModelVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *DentryModelVisitorInfo) SetRoleCode(v string) *DentryModelVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *DentryModelVisitorInfo) SetSpaceActions(v []*string) *DentryModelVisitorInfo {
	s.SpaceActions = v
	return s
}

type DentryVO struct {
	ContentType    *string              `json:"contentType,omitempty" xml:"contentType,omitempty"`
	CreatedTime    *int64               `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Creator        *DentryVOCreator     `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	DentryId       *string              `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	DentryType     *string              `json:"dentryType,omitempty" xml:"dentryType,omitempty"`
	DentryUuid     *string              `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	DocKey         *string              `json:"docKey,omitempty" xml:"docKey,omitempty"`
	Extension      *string              `json:"extension,omitempty" xml:"extension,omitempty"`
	HasChildren    *bool                `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	LinkSourceInfo *LinkSourceInfo      `json:"linkSourceInfo,omitempty" xml:"linkSourceInfo,omitempty"`
	Name           *string              `json:"name,omitempty" xml:"name,omitempty"`
	Path           *string              `json:"path,omitempty" xml:"path,omitempty"`
	Space          *SpaceModel          `json:"space,omitempty" xml:"space,omitempty"`
	SpaceId        *string              `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	UpdatedTime    *int64               `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	Updater        *DentryVOUpdater     `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	Url            *string              `json:"url,omitempty" xml:"url,omitempty"`
	VisitorInfo    *DentryVOVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
}

func (s DentryVO) String() string {
	return tea.Prettify(s)
}

func (s DentryVO) GoString() string {
	return s.String()
}

func (s *DentryVO) SetContentType(v string) *DentryVO {
	s.ContentType = &v
	return s
}

func (s *DentryVO) SetCreatedTime(v int64) *DentryVO {
	s.CreatedTime = &v
	return s
}

func (s *DentryVO) SetCreator(v *DentryVOCreator) *DentryVO {
	s.Creator = v
	return s
}

func (s *DentryVO) SetDentryId(v string) *DentryVO {
	s.DentryId = &v
	return s
}

func (s *DentryVO) SetDentryType(v string) *DentryVO {
	s.DentryType = &v
	return s
}

func (s *DentryVO) SetDentryUuid(v string) *DentryVO {
	s.DentryUuid = &v
	return s
}

func (s *DentryVO) SetDocKey(v string) *DentryVO {
	s.DocKey = &v
	return s
}

func (s *DentryVO) SetExtension(v string) *DentryVO {
	s.Extension = &v
	return s
}

func (s *DentryVO) SetHasChildren(v bool) *DentryVO {
	s.HasChildren = &v
	return s
}

func (s *DentryVO) SetLinkSourceInfo(v *LinkSourceInfo) *DentryVO {
	s.LinkSourceInfo = v
	return s
}

func (s *DentryVO) SetName(v string) *DentryVO {
	s.Name = &v
	return s
}

func (s *DentryVO) SetPath(v string) *DentryVO {
	s.Path = &v
	return s
}

func (s *DentryVO) SetSpace(v *SpaceModel) *DentryVO {
	s.Space = v
	return s
}

func (s *DentryVO) SetSpaceId(v string) *DentryVO {
	s.SpaceId = &v
	return s
}

func (s *DentryVO) SetUpdatedTime(v int64) *DentryVO {
	s.UpdatedTime = &v
	return s
}

func (s *DentryVO) SetUpdater(v *DentryVOUpdater) *DentryVO {
	s.Updater = v
	return s
}

func (s *DentryVO) SetUrl(v string) *DentryVO {
	s.Url = &v
	return s
}

func (s *DentryVO) SetVisitorInfo(v *DentryVOVisitorInfo) *DentryVO {
	s.VisitorInfo = v
	return s
}

type DentryVOCreator struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DentryVOCreator) String() string {
	return tea.Prettify(s)
}

func (s DentryVOCreator) GoString() string {
	return s.String()
}

func (s *DentryVOCreator) SetName(v string) *DentryVOCreator {
	s.Name = &v
	return s
}

func (s *DentryVOCreator) SetUnionId(v string) *DentryVOCreator {
	s.UnionId = &v
	return s
}

type DentryVOUpdater struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DentryVOUpdater) String() string {
	return tea.Prettify(s)
}

func (s DentryVOUpdater) GoString() string {
	return s.String()
}

func (s *DentryVOUpdater) SetName(v string) *DentryVOUpdater {
	s.Name = &v
	return s
}

func (s *DentryVOUpdater) SetUnionId(v string) *DentryVOUpdater {
	s.UnionId = &v
	return s
}

type DentryVOVisitorInfo struct {
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	RoleCode      *string   `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	SpaceActions  []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
}

func (s DentryVOVisitorInfo) String() string {
	return tea.Prettify(s)
}

func (s DentryVOVisitorInfo) GoString() string {
	return s.String()
}

func (s *DentryVOVisitorInfo) SetDentryActions(v []*string) *DentryVOVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *DentryVOVisitorInfo) SetRoleCode(v string) *DentryVOVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *DentryVOVisitorInfo) SetSpaceActions(v []*string) *DentryVOVisitorInfo {
	s.SpaceActions = v
	return s
}

type LinkSourceInfo struct {
	Extension *string                `json:"extension,omitempty" xml:"extension,omitempty"`
	IconUrl   *LinkSourceInfoIconUrl `json:"iconUrl,omitempty" xml:"iconUrl,omitempty" type:"Struct"`
	Id        *string                `json:"id,omitempty" xml:"id,omitempty"`
	LinkType  *int64                 `json:"linkType,omitempty" xml:"linkType,omitempty"`
	SpaceId   *string                `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s LinkSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s LinkSourceInfo) GoString() string {
	return s.String()
}

func (s *LinkSourceInfo) SetExtension(v string) *LinkSourceInfo {
	s.Extension = &v
	return s
}

func (s *LinkSourceInfo) SetIconUrl(v *LinkSourceInfoIconUrl) *LinkSourceInfo {
	s.IconUrl = v
	return s
}

func (s *LinkSourceInfo) SetId(v string) *LinkSourceInfo {
	s.Id = &v
	return s
}

func (s *LinkSourceInfo) SetLinkType(v int64) *LinkSourceInfo {
	s.LinkType = &v
	return s
}

func (s *LinkSourceInfo) SetSpaceId(v string) *LinkSourceInfo {
	s.SpaceId = &v
	return s
}

type LinkSourceInfoIconUrl struct {
	Line  *string `json:"line,omitempty" xml:"line,omitempty"`
	Small *string `json:"small,omitempty" xml:"small,omitempty"`
}

func (s LinkSourceInfoIconUrl) String() string {
	return tea.Prettify(s)
}

func (s LinkSourceInfoIconUrl) GoString() string {
	return s.String()
}

func (s *LinkSourceInfoIconUrl) SetLine(v string) *LinkSourceInfoIconUrl {
	s.Line = &v
	return s
}

func (s *LinkSourceInfoIconUrl) SetSmall(v string) *LinkSourceInfoIconUrl {
	s.Small = &v
	return s
}

type OpenActionModel struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s OpenActionModel) String() string {
	return tea.Prettify(s)
}

func (s OpenActionModel) GoString() string {
	return s.String()
}

func (s *OpenActionModel) SetName(v string) *OpenActionModel {
	s.Name = &v
	return s
}

func (s *OpenActionModel) SetTimestamp(v int64) *OpenActionModel {
	s.Timestamp = &v
	return s
}

type SpaceModel struct {
	Cover       *string                `json:"cover,omitempty" xml:"cover,omitempty"`
	Description *string                `json:"description,omitempty" xml:"description,omitempty"`
	HdIconVO    *SpaceModelHdIconVO    `json:"hdIconVO,omitempty" xml:"hdIconVO,omitempty" type:"Struct"`
	IconVO      *SpaceModelIconVO      `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	Id          *string                `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string                `json:"name,omitempty" xml:"name,omitempty"`
	Owner       *SpaceModelOwner       `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	RecentList  []*DentryModel         `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
	Type        *int32                 `json:"type,omitempty" xml:"type,omitempty"`
	Url         *string                `json:"url,omitempty" xml:"url,omitempty"`
	VisitorInfo *SpaceModelVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
}

func (s SpaceModel) String() string {
	return tea.Prettify(s)
}

func (s SpaceModel) GoString() string {
	return s.String()
}

func (s *SpaceModel) SetCover(v string) *SpaceModel {
	s.Cover = &v
	return s
}

func (s *SpaceModel) SetDescription(v string) *SpaceModel {
	s.Description = &v
	return s
}

func (s *SpaceModel) SetHdIconVO(v *SpaceModelHdIconVO) *SpaceModel {
	s.HdIconVO = v
	return s
}

func (s *SpaceModel) SetIconVO(v *SpaceModelIconVO) *SpaceModel {
	s.IconVO = v
	return s
}

func (s *SpaceModel) SetId(v string) *SpaceModel {
	s.Id = &v
	return s
}

func (s *SpaceModel) SetName(v string) *SpaceModel {
	s.Name = &v
	return s
}

func (s *SpaceModel) SetOwner(v *SpaceModelOwner) *SpaceModel {
	s.Owner = v
	return s
}

func (s *SpaceModel) SetRecentList(v []*DentryModel) *SpaceModel {
	s.RecentList = v
	return s
}

func (s *SpaceModel) SetType(v int32) *SpaceModel {
	s.Type = &v
	return s
}

func (s *SpaceModel) SetUrl(v string) *SpaceModel {
	s.Url = &v
	return s
}

func (s *SpaceModel) SetVisitorInfo(v *SpaceModelVisitorInfo) *SpaceModel {
	s.VisitorInfo = v
	return s
}

type SpaceModelHdIconVO struct {
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SpaceModelHdIconVO) String() string {
	return tea.Prettify(s)
}

func (s SpaceModelHdIconVO) GoString() string {
	return s.String()
}

func (s *SpaceModelHdIconVO) SetIcon(v string) *SpaceModelHdIconVO {
	s.Icon = &v
	return s
}

func (s *SpaceModelHdIconVO) SetType(v string) *SpaceModelHdIconVO {
	s.Type = &v
	return s
}

type SpaceModelIconVO struct {
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SpaceModelIconVO) String() string {
	return tea.Prettify(s)
}

func (s SpaceModelIconVO) GoString() string {
	return s.String()
}

func (s *SpaceModelIconVO) SetIcon(v string) *SpaceModelIconVO {
	s.Icon = &v
	return s
}

func (s *SpaceModelIconVO) SetType(v string) *SpaceModelIconVO {
	s.Type = &v
	return s
}

type SpaceModelOwner struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SpaceModelOwner) String() string {
	return tea.Prettify(s)
}

func (s SpaceModelOwner) GoString() string {
	return s.String()
}

func (s *SpaceModelOwner) SetName(v string) *SpaceModelOwner {
	s.Name = &v
	return s
}

func (s *SpaceModelOwner) SetUnionId(v string) *SpaceModelOwner {
	s.UnionId = &v
	return s
}

type SpaceModelVisitorInfo struct {
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	RoleCode      *string   `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	SpaceActions  []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
}

func (s SpaceModelVisitorInfo) String() string {
	return tea.Prettify(s)
}

func (s SpaceModelVisitorInfo) GoString() string {
	return s.String()
}

func (s *SpaceModelVisitorInfo) SetDentryActions(v []*string) *SpaceModelVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *SpaceModelVisitorInfo) SetRoleCode(v string) *SpaceModelVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *SpaceModelVisitorInfo) SetSpaceActions(v []*string) *SpaceModelVisitorInfo {
	s.SpaceActions = v
	return s
}

type SpaceVO struct {
	Cover       *string             `json:"cover,omitempty" xml:"cover,omitempty"`
	Description *string             `json:"description,omitempty" xml:"description,omitempty"`
	IconVO      *SpaceVOIconVO      `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	Id          *string             `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string             `json:"name,omitempty" xml:"name,omitempty"`
	Owner       *SpaceVOOwner       `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	Type        *int32              `json:"type,omitempty" xml:"type,omitempty"`
	Url         *string             `json:"url,omitempty" xml:"url,omitempty"`
	VisitorInfo *SpaceVOVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
}

func (s SpaceVO) String() string {
	return tea.Prettify(s)
}

func (s SpaceVO) GoString() string {
	return s.String()
}

func (s *SpaceVO) SetCover(v string) *SpaceVO {
	s.Cover = &v
	return s
}

func (s *SpaceVO) SetDescription(v string) *SpaceVO {
	s.Description = &v
	return s
}

func (s *SpaceVO) SetIconVO(v *SpaceVOIconVO) *SpaceVO {
	s.IconVO = v
	return s
}

func (s *SpaceVO) SetId(v string) *SpaceVO {
	s.Id = &v
	return s
}

func (s *SpaceVO) SetName(v string) *SpaceVO {
	s.Name = &v
	return s
}

func (s *SpaceVO) SetOwner(v *SpaceVOOwner) *SpaceVO {
	s.Owner = v
	return s
}

func (s *SpaceVO) SetType(v int32) *SpaceVO {
	s.Type = &v
	return s
}

func (s *SpaceVO) SetUrl(v string) *SpaceVO {
	s.Url = &v
	return s
}

func (s *SpaceVO) SetVisitorInfo(v *SpaceVOVisitorInfo) *SpaceVO {
	s.VisitorInfo = v
	return s
}

type SpaceVOIconVO struct {
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SpaceVOIconVO) String() string {
	return tea.Prettify(s)
}

func (s SpaceVOIconVO) GoString() string {
	return s.String()
}

func (s *SpaceVOIconVO) SetIcon(v string) *SpaceVOIconVO {
	s.Icon = &v
	return s
}

func (s *SpaceVOIconVO) SetType(v string) *SpaceVOIconVO {
	s.Type = &v
	return s
}

type SpaceVOOwner struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SpaceVOOwner) String() string {
	return tea.Prettify(s)
}

func (s SpaceVOOwner) GoString() string {
	return s.String()
}

func (s *SpaceVOOwner) SetName(v string) *SpaceVOOwner {
	s.Name = &v
	return s
}

func (s *SpaceVOOwner) SetUnionId(v string) *SpaceVOOwner {
	s.UnionId = &v
	return s
}

type SpaceVOVisitorInfo struct {
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	RoleCode      *string   `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	SpaceActions  []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
}

func (s SpaceVOVisitorInfo) String() string {
	return tea.Prettify(s)
}

func (s SpaceVOVisitorInfo) GoString() string {
	return s.String()
}

func (s *SpaceVOVisitorInfo) SetDentryActions(v []*string) *SpaceVOVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *SpaceVOVisitorInfo) SetRoleCode(v string) *SpaceVOVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *SpaceVOVisitorInfo) SetSpaceActions(v []*string) *SpaceVOVisitorInfo {
	s.SpaceActions = v
	return s
}

type TeamModel struct {
	Cover           *string                   `json:"cover,omitempty" xml:"cover,omitempty"`
	CreatedTime     *int64                    `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Creator         *TeamModelCreator         `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	Description     *string                   `json:"description,omitempty" xml:"description,omitempty"`
	Icon            *string                   `json:"icon,omitempty" xml:"icon,omitempty"`
	Id              *string                   `json:"id,omitempty" xml:"id,omitempty"`
	Name            *string                   `json:"name,omitempty" xml:"name,omitempty"`
	RelatedDeptInfo *TeamModelRelatedDeptInfo `json:"relatedDeptInfo,omitempty" xml:"relatedDeptInfo,omitempty" type:"Struct"`
	Status          *int32                    `json:"status,omitempty" xml:"status,omitempty"`
	Type            *int32                    `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedTime     *int64                    `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	Updater         *TeamModelUpdater         `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	Url             *string                   `json:"url,omitempty" xml:"url,omitempty"`
	VisitInfo       *TeamModelVisitInfo       `json:"visitInfo,omitempty" xml:"visitInfo,omitempty" type:"Struct"`
}

func (s TeamModel) String() string {
	return tea.Prettify(s)
}

func (s TeamModel) GoString() string {
	return s.String()
}

func (s *TeamModel) SetCover(v string) *TeamModel {
	s.Cover = &v
	return s
}

func (s *TeamModel) SetCreatedTime(v int64) *TeamModel {
	s.CreatedTime = &v
	return s
}

func (s *TeamModel) SetCreator(v *TeamModelCreator) *TeamModel {
	s.Creator = v
	return s
}

func (s *TeamModel) SetDescription(v string) *TeamModel {
	s.Description = &v
	return s
}

func (s *TeamModel) SetIcon(v string) *TeamModel {
	s.Icon = &v
	return s
}

func (s *TeamModel) SetId(v string) *TeamModel {
	s.Id = &v
	return s
}

func (s *TeamModel) SetName(v string) *TeamModel {
	s.Name = &v
	return s
}

func (s *TeamModel) SetRelatedDeptInfo(v *TeamModelRelatedDeptInfo) *TeamModel {
	s.RelatedDeptInfo = v
	return s
}

func (s *TeamModel) SetStatus(v int32) *TeamModel {
	s.Status = &v
	return s
}

func (s *TeamModel) SetType(v int32) *TeamModel {
	s.Type = &v
	return s
}

func (s *TeamModel) SetUpdatedTime(v int64) *TeamModel {
	s.UpdatedTime = &v
	return s
}

func (s *TeamModel) SetUpdater(v *TeamModelUpdater) *TeamModel {
	s.Updater = v
	return s
}

func (s *TeamModel) SetUrl(v string) *TeamModel {
	s.Url = &v
	return s
}

func (s *TeamModel) SetVisitInfo(v *TeamModelVisitInfo) *TeamModel {
	s.VisitInfo = v
	return s
}

type TeamModelCreator struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s TeamModelCreator) String() string {
	return tea.Prettify(s)
}

func (s TeamModelCreator) GoString() string {
	return s.String()
}

func (s *TeamModelCreator) SetName(v string) *TeamModelCreator {
	s.Name = &v
	return s
}

func (s *TeamModelCreator) SetUnionId(v string) *TeamModelCreator {
	s.UnionId = &v
	return s
}

type TeamModelRelatedDeptInfo struct {
	DeptId   *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
}

func (s TeamModelRelatedDeptInfo) String() string {
	return tea.Prettify(s)
}

func (s TeamModelRelatedDeptInfo) GoString() string {
	return s.String()
}

func (s *TeamModelRelatedDeptInfo) SetDeptId(v string) *TeamModelRelatedDeptInfo {
	s.DeptId = &v
	return s
}

func (s *TeamModelRelatedDeptInfo) SetDeptName(v string) *TeamModelRelatedDeptInfo {
	s.DeptName = &v
	return s
}

type TeamModelUpdater struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s TeamModelUpdater) String() string {
	return tea.Prettify(s)
}

func (s TeamModelUpdater) GoString() string {
	return s.String()
}

func (s *TeamModelUpdater) SetName(v string) *TeamModelUpdater {
	s.Name = &v
	return s
}

func (s *TeamModelUpdater) SetUnionId(v string) *TeamModelUpdater {
	s.UnionId = &v
	return s
}

type TeamModelVisitInfo struct {
	JoinTime *string `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s TeamModelVisitInfo) String() string {
	return tea.Prettify(s)
}

func (s TeamModelVisitInfo) GoString() string {
	return s.String()
}

func (s *TeamModelVisitInfo) SetJoinTime(v string) *TeamModelVisitInfo {
	s.JoinTime = &v
	return s
}

func (s *TeamModelVisitInfo) SetRoleCode(v string) *TeamModelVisitInfo {
	s.RoleCode = &v
	return s
}

type TeamVO struct {
	Cover           *string                `json:"cover,omitempty" xml:"cover,omitempty"`
	CreatedTime     *int64                 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Creator         *TeamVOCreator         `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	Description     *string                `json:"description,omitempty" xml:"description,omitempty"`
	Icon            *string                `json:"icon,omitempty" xml:"icon,omitempty"`
	Id              *string                `json:"id,omitempty" xml:"id,omitempty"`
	Name            *string                `json:"name,omitempty" xml:"name,omitempty"`
	RelatedDeptInfo *TeamVORelatedDeptInfo `json:"relatedDeptInfo,omitempty" xml:"relatedDeptInfo,omitempty" type:"Struct"`
	ShareScopeInfo  *TeamVOShareScopeInfo  `json:"shareScopeInfo,omitempty" xml:"shareScopeInfo,omitempty" type:"Struct"`
	Status          *int32                 `json:"status,omitempty" xml:"status,omitempty"`
	Type            *int32                 `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedTime     *int64                 `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	Updater         *TeamVOUpdater         `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	Url             *string                `json:"url,omitempty" xml:"url,omitempty"`
	VisitInfo       *TeamVOVisitInfo       `json:"visitInfo,omitempty" xml:"visitInfo,omitempty" type:"Struct"`
}

func (s TeamVO) String() string {
	return tea.Prettify(s)
}

func (s TeamVO) GoString() string {
	return s.String()
}

func (s *TeamVO) SetCover(v string) *TeamVO {
	s.Cover = &v
	return s
}

func (s *TeamVO) SetCreatedTime(v int64) *TeamVO {
	s.CreatedTime = &v
	return s
}

func (s *TeamVO) SetCreator(v *TeamVOCreator) *TeamVO {
	s.Creator = v
	return s
}

func (s *TeamVO) SetDescription(v string) *TeamVO {
	s.Description = &v
	return s
}

func (s *TeamVO) SetIcon(v string) *TeamVO {
	s.Icon = &v
	return s
}

func (s *TeamVO) SetId(v string) *TeamVO {
	s.Id = &v
	return s
}

func (s *TeamVO) SetName(v string) *TeamVO {
	s.Name = &v
	return s
}

func (s *TeamVO) SetRelatedDeptInfo(v *TeamVORelatedDeptInfo) *TeamVO {
	s.RelatedDeptInfo = v
	return s
}

func (s *TeamVO) SetShareScopeInfo(v *TeamVOShareScopeInfo) *TeamVO {
	s.ShareScopeInfo = v
	return s
}

func (s *TeamVO) SetStatus(v int32) *TeamVO {
	s.Status = &v
	return s
}

func (s *TeamVO) SetType(v int32) *TeamVO {
	s.Type = &v
	return s
}

func (s *TeamVO) SetUpdatedTime(v int64) *TeamVO {
	s.UpdatedTime = &v
	return s
}

func (s *TeamVO) SetUpdater(v *TeamVOUpdater) *TeamVO {
	s.Updater = v
	return s
}

func (s *TeamVO) SetUrl(v string) *TeamVO {
	s.Url = &v
	return s
}

func (s *TeamVO) SetVisitInfo(v *TeamVOVisitInfo) *TeamVO {
	s.VisitInfo = v
	return s
}

type TeamVOCreator struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s TeamVOCreator) String() string {
	return tea.Prettify(s)
}

func (s TeamVOCreator) GoString() string {
	return s.String()
}

func (s *TeamVOCreator) SetName(v string) *TeamVOCreator {
	s.Name = &v
	return s
}

func (s *TeamVOCreator) SetUnionId(v string) *TeamVOCreator {
	s.UnionId = &v
	return s
}

type TeamVORelatedDeptInfo struct {
	DeptId   *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
}

func (s TeamVORelatedDeptInfo) String() string {
	return tea.Prettify(s)
}

func (s TeamVORelatedDeptInfo) GoString() string {
	return s.String()
}

func (s *TeamVORelatedDeptInfo) SetDeptId(v string) *TeamVORelatedDeptInfo {
	s.DeptId = &v
	return s
}

func (s *TeamVORelatedDeptInfo) SetDeptName(v string) *TeamVORelatedDeptInfo {
	s.DeptName = &v
	return s
}

type TeamVOShareScopeInfo struct {
	RoleId *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	Scope  *int32  `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s TeamVOShareScopeInfo) String() string {
	return tea.Prettify(s)
}

func (s TeamVOShareScopeInfo) GoString() string {
	return s.String()
}

func (s *TeamVOShareScopeInfo) SetRoleId(v string) *TeamVOShareScopeInfo {
	s.RoleId = &v
	return s
}

func (s *TeamVOShareScopeInfo) SetScope(v int32) *TeamVOShareScopeInfo {
	s.Scope = &v
	return s
}

type TeamVOUpdater struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s TeamVOUpdater) String() string {
	return tea.Prettify(s)
}

func (s TeamVOUpdater) GoString() string {
	return s.String()
}

func (s *TeamVOUpdater) SetName(v string) *TeamVOUpdater {
	s.Name = &v
	return s
}

func (s *TeamVOUpdater) SetUnionId(v string) *TeamVOUpdater {
	s.UnionId = &v
	return s
}

type TeamVOVisitInfo struct {
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s TeamVOVisitInfo) String() string {
	return tea.Prettify(s)
}

func (s TeamVOVisitInfo) GoString() string {
	return s.String()
}

func (s *TeamVOVisitInfo) SetRoleCode(v string) *TeamVOVisitInfo {
	s.RoleCode = &v
	return s
}

type BatchDeleteRecentsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchDeleteRecentsHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRecentsHeaders) GoString() string {
	return s.String()
}

func (s *BatchDeleteRecentsHeaders) SetCommonHeaders(v map[string]*string) *BatchDeleteRecentsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchDeleteRecentsHeaders) SetXAcsDingtalkAccessToken(v string) *BatchDeleteRecentsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchDeleteRecentsRequest struct {
	DentryUuids []*string `json:"dentryUuids,omitempty" xml:"dentryUuids,omitempty" type:"Repeated"`
	OperatorId  *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s BatchDeleteRecentsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRecentsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteRecentsRequest) SetDentryUuids(v []*string) *BatchDeleteRecentsRequest {
	s.DentryUuids = v
	return s
}

func (s *BatchDeleteRecentsRequest) SetOperatorId(v string) *BatchDeleteRecentsRequest {
	s.OperatorId = &v
	return s
}

type BatchDeleteRecentsResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchDeleteRecentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRecentsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteRecentsResponseBody) SetSuccess(v bool) *BatchDeleteRecentsResponseBody {
	s.Success = &v
	return s
}

type BatchDeleteRecentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchDeleteRecentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteRecentsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRecentsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteRecentsResponse) SetHeaders(v map[string]*string) *BatchDeleteRecentsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteRecentsResponse) SetStatusCode(v int32) *BatchDeleteRecentsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteRecentsResponse) SetBody(v *BatchDeleteRecentsResponseBody) *BatchDeleteRecentsResponse {
	s.Body = v
	return s
}

type CopyDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CopyDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryHeaders) GoString() string {
	return s.String()
}

func (s *CopyDentryHeaders) SetCommonHeaders(v map[string]*string) *CopyDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyDentryHeaders) SetXAcsDingtalkAccessToken(v string) *CopyDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CopyDentryRequest struct {
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	OperatorId       *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	TargetSpaceId    *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	ToNextDentryId   *string `json:"toNextDentryId,omitempty" xml:"toNextDentryId,omitempty"`
	ToParentDentryId *string `json:"toParentDentryId,omitempty" xml:"toParentDentryId,omitempty"`
	ToPrevDentryId   *string `json:"toPrevDentryId,omitempty" xml:"toPrevDentryId,omitempty"`
}

func (s CopyDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryRequest) GoString() string {
	return s.String()
}

func (s *CopyDentryRequest) SetName(v string) *CopyDentryRequest {
	s.Name = &v
	return s
}

func (s *CopyDentryRequest) SetOperatorId(v string) *CopyDentryRequest {
	s.OperatorId = &v
	return s
}

func (s *CopyDentryRequest) SetTargetSpaceId(v string) *CopyDentryRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *CopyDentryRequest) SetToNextDentryId(v string) *CopyDentryRequest {
	s.ToNextDentryId = &v
	return s
}

func (s *CopyDentryRequest) SetToParentDentryId(v string) *CopyDentryRequest {
	s.ToParentDentryId = &v
	return s
}

func (s *CopyDentryRequest) SetToPrevDentryId(v string) *CopyDentryRequest {
	s.ToPrevDentryId = &v
	return s
}

type CopyDentryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DentryVO          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CopyDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryResponse) GoString() string {
	return s.String()
}

func (s *CopyDentryResponse) SetHeaders(v map[string]*string) *CopyDentryResponse {
	s.Headers = v
	return s
}

func (s *CopyDentryResponse) SetStatusCode(v int32) *CopyDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDentryResponse) SetBody(v *DentryVO) *CopyDentryResponse {
	s.Body = v
	return s
}

type CreateDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDentryHeaders) GoString() string {
	return s.String()
}

func (s *CreateDentryHeaders) SetCommonHeaders(v map[string]*string) *CreateDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDentryHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDentryRequest struct {
	DentryType     *string `json:"dentryType,omitempty" xml:"dentryType,omitempty"`
	DocumentType   *int64  `json:"documentType,omitempty" xml:"documentType,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OperatorId     *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	ParentDentryId *string `json:"parentDentryId,omitempty" xml:"parentDentryId,omitempty"`
}

func (s CreateDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDentryRequest) GoString() string {
	return s.String()
}

func (s *CreateDentryRequest) SetDentryType(v string) *CreateDentryRequest {
	s.DentryType = &v
	return s
}

func (s *CreateDentryRequest) SetDocumentType(v int64) *CreateDentryRequest {
	s.DocumentType = &v
	return s
}

func (s *CreateDentryRequest) SetName(v string) *CreateDentryRequest {
	s.Name = &v
	return s
}

func (s *CreateDentryRequest) SetOperatorId(v string) *CreateDentryRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateDentryRequest) SetParentDentryId(v string) *CreateDentryRequest {
	s.ParentDentryId = &v
	return s
}

type CreateDentryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DentryVO          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDentryResponse) GoString() string {
	return s.String()
}

func (s *CreateDentryResponse) SetHeaders(v map[string]*string) *CreateDentryResponse {
	s.Headers = v
	return s
}

func (s *CreateDentryResponse) SetStatusCode(v int32) *CreateDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDentryResponse) SetBody(v *DentryVO) *CreateDentryResponse {
	s.Body = v
	return s
}

type CreateSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceHeaders) GoString() string {
	return s.String()
}

func (s *CreateSpaceHeaders) SetCommonHeaders(v map[string]*string) *CreateSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSpaceRequest struct {
	Description *string                       `json:"description,omitempty" xml:"description,omitempty"`
	Icon        *string                       `json:"icon,omitempty" xml:"icon,omitempty"`
	Name        *string                       `json:"name,omitempty" xml:"name,omitempty"`
	OperatorId  *string                       `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	SectionId   *string                       `json:"sectionId,omitempty" xml:"sectionId,omitempty"`
	ShareScope  *CreateSpaceRequestShareScope `json:"shareScope,omitempty" xml:"shareScope,omitempty" type:"Struct"`
	TeamId      *string                       `json:"teamId,omitempty" xml:"teamId,omitempty"`
}

func (s CreateSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateSpaceRequest) SetDescription(v string) *CreateSpaceRequest {
	s.Description = &v
	return s
}

func (s *CreateSpaceRequest) SetIcon(v string) *CreateSpaceRequest {
	s.Icon = &v
	return s
}

func (s *CreateSpaceRequest) SetName(v string) *CreateSpaceRequest {
	s.Name = &v
	return s
}

func (s *CreateSpaceRequest) SetOperatorId(v string) *CreateSpaceRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateSpaceRequest) SetSectionId(v string) *CreateSpaceRequest {
	s.SectionId = &v
	return s
}

func (s *CreateSpaceRequest) SetShareScope(v *CreateSpaceRequestShareScope) *CreateSpaceRequest {
	s.ShareScope = v
	return s
}

func (s *CreateSpaceRequest) SetTeamId(v string) *CreateSpaceRequest {
	s.TeamId = &v
	return s
}

type CreateSpaceRequestShareScope struct {
	Scope *int32 `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s CreateSpaceRequestShareScope) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceRequestShareScope) GoString() string {
	return s.String()
}

func (s *CreateSpaceRequestShareScope) SetScope(v int32) *CreateSpaceRequestShareScope {
	s.Scope = &v
	return s
}

type CreateSpaceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SpaceVO           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreateSpaceResponse) SetHeaders(v map[string]*string) *CreateSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreateSpaceResponse) SetStatusCode(v int32) *CreateSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSpaceResponse) SetBody(v *SpaceVO) *CreateSpaceResponse {
	s.Body = v
	return s
}

type CreateTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTeamHeaders) GoString() string {
	return s.String()
}

func (s *CreateTeamHeaders) SetCommonHeaders(v map[string]*string) *CreateTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTeamHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTeamRequest struct {
	Cover       *string                     `json:"cover,omitempty" xml:"cover,omitempty"`
	Description *string                     `json:"description,omitempty" xml:"description,omitempty"`
	Icon        *string                     `json:"icon,omitempty" xml:"icon,omitempty"`
	Members     []*CreateTeamRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	Name        *string                     `json:"name,omitempty" xml:"name,omitempty"`
	OperatorId  *string                     `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	TeamType    *int32                      `json:"teamType,omitempty" xml:"teamType,omitempty"`
}

func (s CreateTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTeamRequest) GoString() string {
	return s.String()
}

func (s *CreateTeamRequest) SetCover(v string) *CreateTeamRequest {
	s.Cover = &v
	return s
}

func (s *CreateTeamRequest) SetDescription(v string) *CreateTeamRequest {
	s.Description = &v
	return s
}

func (s *CreateTeamRequest) SetIcon(v string) *CreateTeamRequest {
	s.Icon = &v
	return s
}

func (s *CreateTeamRequest) SetMembers(v []*CreateTeamRequestMembers) *CreateTeamRequest {
	s.Members = v
	return s
}

func (s *CreateTeamRequest) SetName(v string) *CreateTeamRequest {
	s.Name = &v
	return s
}

func (s *CreateTeamRequest) SetOperatorId(v string) *CreateTeamRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateTeamRequest) SetTeamType(v int32) *CreateTeamRequest {
	s.TeamType = &v
	return s
}

type CreateTeamRequestMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *int32  `json:"memberType,omitempty" xml:"memberType,omitempty"`
	RoleCode   *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s CreateTeamRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s CreateTeamRequestMembers) GoString() string {
	return s.String()
}

func (s *CreateTeamRequestMembers) SetMemberId(v string) *CreateTeamRequestMembers {
	s.MemberId = &v
	return s
}

func (s *CreateTeamRequestMembers) SetMemberType(v int32) *CreateTeamRequestMembers {
	s.MemberType = &v
	return s
}

func (s *CreateTeamRequestMembers) SetRoleCode(v string) *CreateTeamRequestMembers {
	s.RoleCode = &v
	return s
}

type CreateTeamResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TeamVO            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTeamResponse) GoString() string {
	return s.String()
}

func (s *CreateTeamResponse) SetHeaders(v map[string]*string) *CreateTeamResponse {
	s.Headers = v
	return s
}

func (s *CreateTeamResponse) SetStatusCode(v int32) *CreateTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTeamResponse) SetBody(v *TeamVO) *CreateTeamResponse {
	s.Body = v
	return s
}

type DeleteTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeamHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTeamHeaders) SetCommonHeaders(v map[string]*string) *DeleteTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTeamHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteTeamRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeamRequest) GoString() string {
	return s.String()
}

func (s *DeleteTeamRequest) SetOperatorId(v string) *DeleteTeamRequest {
	s.OperatorId = &v
	return s
}

type DeleteTeamResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TeamVO            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeamResponse) GoString() string {
	return s.String()
}

func (s *DeleteTeamResponse) SetHeaders(v map[string]*string) *DeleteTeamResponse {
	s.Headers = v
	return s
}

func (s *DeleteTeamResponse) SetStatusCode(v int32) *DeleteTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTeamResponse) SetBody(v *TeamVO) *DeleteTeamResponse {
	s.Body = v
	return s
}

type GetSchemaHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSchemaHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSchemaHeaders) GoString() string {
	return s.String()
}

func (s *GetSchemaHeaders) SetCommonHeaders(v map[string]*string) *GetSchemaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSchemaHeaders) SetXAcsDingtalkAccessToken(v string) *GetSchemaHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSchemaRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetSchemaRequest) SetOperatorId(v string) *GetSchemaRequest {
	s.OperatorId = &v
	return s
}

type GetSchemaResponseBody struct {
	Revision *int32  `json:"revision,omitempty" xml:"revision,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetSchemaResponseBody) SetRevision(v int32) *GetSchemaResponseBody {
	s.Revision = &v
	return s
}

func (s *GetSchemaResponseBody) SetValue(v string) *GetSchemaResponseBody {
	s.Value = &v
	return s
}

type GetSchemaResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetSchemaResponse) SetHeaders(v map[string]*string) *GetSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetSchemaResponse) SetStatusCode(v int32) *GetSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSchemaResponse) SetBody(v *GetSchemaResponseBody) *GetSchemaResponse {
	s.Body = v
	return s
}

type GetSpaceDirectoriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSpaceDirectoriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceDirectoriesHeaders) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesHeaders) SetCommonHeaders(v map[string]*string) *GetSpaceDirectoriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSpaceDirectoriesHeaders) SetXAcsDingtalkAccessToken(v string) *GetSpaceDirectoriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSpaceDirectoriesRequest struct {
	DentryId   *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetSpaceDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesRequest) SetDentryId(v string) *GetSpaceDirectoriesRequest {
	s.DentryId = &v
	return s
}

func (s *GetSpaceDirectoriesRequest) SetMaxResults(v int32) *GetSpaceDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSpaceDirectoriesRequest) SetNextToken(v string) *GetSpaceDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *GetSpaceDirectoriesRequest) SetOperatorId(v string) *GetSpaceDirectoriesRequest {
	s.OperatorId = &v
	return s
}

type GetSpaceDirectoriesResponseBody struct {
	Children  []*DentryModel `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	HasMore   *bool          `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *string        `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetSpaceDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBody) SetChildren(v []*DentryModel) *GetSpaceDirectoriesResponseBody {
	s.Children = v
	return s
}

func (s *GetSpaceDirectoriesResponseBody) SetHasMore(v bool) *GetSpaceDirectoriesResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBody) SetNextToken(v string) *GetSpaceDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

type GetSpaceDirectoriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSpaceDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSpaceDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponse) SetHeaders(v map[string]*string) *GetSpaceDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceDirectoriesResponse) SetStatusCode(v int32) *GetSpaceDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpaceDirectoriesResponse) SetBody(v *GetSpaceDirectoriesResponseBody) *GetSpaceDirectoriesResponse {
	s.Body = v
	return s
}

type GetStarInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetStarInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetStarInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetStarInfoHeaders) SetCommonHeaders(v map[string]*string) *GetStarInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetStarInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetStarInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetStarInfoRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetStarInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStarInfoRequest) GoString() string {
	return s.String()
}

func (s *GetStarInfoRequest) SetOperatorId(v string) *GetStarInfoRequest {
	s.OperatorId = &v
	return s
}

type GetStarInfoResponseBody struct {
	Starred *bool `json:"starred,omitempty" xml:"starred,omitempty"`
}

func (s GetStarInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStarInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetStarInfoResponseBody) SetStarred(v bool) *GetStarInfoResponseBody {
	s.Starred = &v
	return s
}

type GetStarInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStarInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStarInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStarInfoResponse) GoString() string {
	return s.String()
}

func (s *GetStarInfoResponse) SetHeaders(v map[string]*string) *GetStarInfoResponse {
	s.Headers = v
	return s
}

func (s *GetStarInfoResponse) SetStatusCode(v int32) *GetStarInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStarInfoResponse) SetBody(v *GetStarInfoResponseBody) *GetStarInfoResponse {
	s.Body = v
	return s
}

type GetTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTeamHeaders) GoString() string {
	return s.String()
}

func (s *GetTeamHeaders) SetCommonHeaders(v map[string]*string) *GetTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTeamHeaders) SetXAcsDingtalkAccessToken(v string) *GetTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTeamRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTeamRequest) GoString() string {
	return s.String()
}

func (s *GetTeamRequest) SetOperatorId(v string) *GetTeamRequest {
	s.OperatorId = &v
	return s
}

type GetTeamResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TeamVO            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTeamResponse) GoString() string {
	return s.String()
}

func (s *GetTeamResponse) SetHeaders(v map[string]*string) *GetTeamResponse {
	s.Headers = v
	return s
}

func (s *GetTeamResponse) SetStatusCode(v int32) *GetTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTeamResponse) SetBody(v *TeamVO) *GetTeamResponse {
	s.Body = v
	return s
}

type GetTotalNumberOfDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTotalNumberOfDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTotalNumberOfDentriesHeaders) GoString() string {
	return s.String()
}

func (s *GetTotalNumberOfDentriesHeaders) SetCommonHeaders(v map[string]*string) *GetTotalNumberOfDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTotalNumberOfDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *GetTotalNumberOfDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTotalNumberOfDentriesRequest struct {
	IncludeFolder *bool   `json:"includeFolder,omitempty" xml:"includeFolder,omitempty"`
	OperatorId    *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	SpaceTypes    *string `json:"spaceTypes,omitempty" xml:"spaceTypes,omitempty"`
}

func (s GetTotalNumberOfDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTotalNumberOfDentriesRequest) GoString() string {
	return s.String()
}

func (s *GetTotalNumberOfDentriesRequest) SetIncludeFolder(v bool) *GetTotalNumberOfDentriesRequest {
	s.IncludeFolder = &v
	return s
}

func (s *GetTotalNumberOfDentriesRequest) SetOperatorId(v string) *GetTotalNumberOfDentriesRequest {
	s.OperatorId = &v
	return s
}

func (s *GetTotalNumberOfDentriesRequest) SetSpaceTypes(v string) *GetTotalNumberOfDentriesRequest {
	s.SpaceTypes = &v
	return s
}

type GetTotalNumberOfDentriesResponseBody struct {
	DentriesCount *string `json:"dentriesCount,omitempty" xml:"dentriesCount,omitempty"`
}

func (s GetTotalNumberOfDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTotalNumberOfDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTotalNumberOfDentriesResponseBody) SetDentriesCount(v string) *GetTotalNumberOfDentriesResponseBody {
	s.DentriesCount = &v
	return s
}

type GetTotalNumberOfDentriesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTotalNumberOfDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTotalNumberOfDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTotalNumberOfDentriesResponse) GoString() string {
	return s.String()
}

func (s *GetTotalNumberOfDentriesResponse) SetHeaders(v map[string]*string) *GetTotalNumberOfDentriesResponse {
	s.Headers = v
	return s
}

func (s *GetTotalNumberOfDentriesResponse) SetStatusCode(v int32) *GetTotalNumberOfDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTotalNumberOfDentriesResponse) SetBody(v *GetTotalNumberOfDentriesResponseBody) *GetTotalNumberOfDentriesResponse {
	s.Body = v
	return s
}

type GetTotalNumberOfSpacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTotalNumberOfSpacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTotalNumberOfSpacesHeaders) GoString() string {
	return s.String()
}

func (s *GetTotalNumberOfSpacesHeaders) SetCommonHeaders(v map[string]*string) *GetTotalNumberOfSpacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTotalNumberOfSpacesHeaders) SetXAcsDingtalkAccessToken(v string) *GetTotalNumberOfSpacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTotalNumberOfSpacesRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetTotalNumberOfSpacesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTotalNumberOfSpacesRequest) GoString() string {
	return s.String()
}

func (s *GetTotalNumberOfSpacesRequest) SetOperatorId(v string) *GetTotalNumberOfSpacesRequest {
	s.OperatorId = &v
	return s
}

type GetTotalNumberOfSpacesResponseBody struct {
	SpacesCount *string `json:"spacesCount,omitempty" xml:"spacesCount,omitempty"`
}

func (s GetTotalNumberOfSpacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTotalNumberOfSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTotalNumberOfSpacesResponseBody) SetSpacesCount(v string) *GetTotalNumberOfSpacesResponseBody {
	s.SpacesCount = &v
	return s
}

type GetTotalNumberOfSpacesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTotalNumberOfSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTotalNumberOfSpacesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTotalNumberOfSpacesResponse) GoString() string {
	return s.String()
}

func (s *GetTotalNumberOfSpacesResponse) SetHeaders(v map[string]*string) *GetTotalNumberOfSpacesResponse {
	s.Headers = v
	return s
}

func (s *GetTotalNumberOfSpacesResponse) SetStatusCode(v int32) *GetTotalNumberOfSpacesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTotalNumberOfSpacesResponse) SetBody(v *GetTotalNumberOfSpacesResponseBody) *GetTotalNumberOfSpacesResponse {
	s.Body = v
	return s
}

type GetUserInfoByOpenTokenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserInfoByOpenTokenHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoByOpenTokenHeaders) GoString() string {
	return s.String()
}

func (s *GetUserInfoByOpenTokenHeaders) SetCommonHeaders(v map[string]*string) *GetUserInfoByOpenTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserInfoByOpenTokenHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserInfoByOpenTokenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserInfoByOpenTokenRequest struct {
	DocKey    *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	OpenToken *string `json:"openToken,omitempty" xml:"openToken,omitempty"`
}

func (s GetUserInfoByOpenTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoByOpenTokenRequest) GoString() string {
	return s.String()
}

func (s *GetUserInfoByOpenTokenRequest) SetDocKey(v string) *GetUserInfoByOpenTokenRequest {
	s.DocKey = &v
	return s
}

func (s *GetUserInfoByOpenTokenRequest) SetOpenToken(v string) *GetUserInfoByOpenTokenRequest {
	s.OpenToken = &v
	return s
}

type GetUserInfoByOpenTokenResponseBody struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserInfoByOpenTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoByOpenTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInfoByOpenTokenResponseBody) SetUnionId(v string) *GetUserInfoByOpenTokenResponseBody {
	s.UnionId = &v
	return s
}

func (s *GetUserInfoByOpenTokenResponseBody) SetUserId(v string) *GetUserInfoByOpenTokenResponseBody {
	s.UserId = &v
	return s
}

type GetUserInfoByOpenTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserInfoByOpenTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserInfoByOpenTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoByOpenTokenResponse) GoString() string {
	return s.String()
}

func (s *GetUserInfoByOpenTokenResponse) SetHeaders(v map[string]*string) *GetUserInfoByOpenTokenResponse {
	s.Headers = v
	return s
}

func (s *GetUserInfoByOpenTokenResponse) SetStatusCode(v int32) *GetUserInfoByOpenTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserInfoByOpenTokenResponse) SetBody(v *GetUserInfoByOpenTokenResponseBody) *GetUserInfoByOpenTokenResponse {
	s.Body = v
	return s
}

type ListFeedsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListFeedsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListFeedsHeaders) GoString() string {
	return s.String()
}

func (s *ListFeedsHeaders) SetCommonHeaders(v map[string]*string) *ListFeedsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFeedsHeaders) SetXAcsDingtalkAccessToken(v string) *ListFeedsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListFeedsRequest struct {
	ExcludeFile *bool   `json:"excludeFile,omitempty" xml:"excludeFile,omitempty"`
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperatorId  *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ListFeedsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeedsRequest) GoString() string {
	return s.String()
}

func (s *ListFeedsRequest) SetExcludeFile(v bool) *ListFeedsRequest {
	s.ExcludeFile = &v
	return s
}

func (s *ListFeedsRequest) SetMaxResults(v int32) *ListFeedsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFeedsRequest) SetNextToken(v string) *ListFeedsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFeedsRequest) SetOperatorId(v string) *ListFeedsRequest {
	s.OperatorId = &v
	return s
}

type ListFeedsResponseBody struct {
	HasMore   *bool                         `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Items     []*ListFeedsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextToken *string                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFeedsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeedsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeedsResponseBody) SetHasMore(v bool) *ListFeedsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListFeedsResponseBody) SetItems(v []*ListFeedsResponseBodyItems) *ListFeedsResponseBody {
	s.Items = v
	return s
}

func (s *ListFeedsResponseBody) SetNextToken(v string) *ListFeedsResponseBody {
	s.NextToken = &v
	return s
}

type ListFeedsResponseBodyItems struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Time    *int64  `json:"time,omitempty" xml:"time,omitempty"`
	Type    *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListFeedsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListFeedsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListFeedsResponseBodyItems) SetContent(v string) *ListFeedsResponseBodyItems {
	s.Content = &v
	return s
}

func (s *ListFeedsResponseBodyItems) SetTime(v int64) *ListFeedsResponseBodyItems {
	s.Time = &v
	return s
}

func (s *ListFeedsResponseBodyItems) SetType(v int32) *ListFeedsResponseBodyItems {
	s.Type = &v
	return s
}

type ListFeedsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFeedsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFeedsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeedsResponse) GoString() string {
	return s.String()
}

func (s *ListFeedsResponse) SetHeaders(v map[string]*string) *ListFeedsResponse {
	s.Headers = v
	return s
}

func (s *ListFeedsResponse) SetStatusCode(v int32) *ListFeedsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeedsResponse) SetBody(v *ListFeedsResponseBody) *ListFeedsResponse {
	s.Body = v
	return s
}

type ListHotDocsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListHotDocsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListHotDocsHeaders) GoString() string {
	return s.String()
}

func (s *ListHotDocsHeaders) SetCommonHeaders(v map[string]*string) *ListHotDocsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotDocsHeaders) SetXAcsDingtalkAccessToken(v string) *ListHotDocsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListHotDocsRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ListHotDocsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotDocsRequest) GoString() string {
	return s.String()
}

func (s *ListHotDocsRequest) SetOperatorId(v string) *ListHotDocsRequest {
	s.OperatorId = &v
	return s
}

type ListHotDocsResponseBody struct {
	Items []*DentryModel `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListHotDocsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotDocsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotDocsResponseBody) SetItems(v []*DentryModel) *ListHotDocsResponseBody {
	s.Items = v
	return s
}

type ListHotDocsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHotDocsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotDocsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotDocsResponse) GoString() string {
	return s.String()
}

func (s *ListHotDocsResponse) SetHeaders(v map[string]*string) *ListHotDocsResponse {
	s.Headers = v
	return s
}

func (s *ListHotDocsResponse) SetStatusCode(v int32) *ListHotDocsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotDocsResponse) SetBody(v *ListHotDocsResponseBody) *ListHotDocsResponse {
	s.Body = v
	return s
}

type ListPinSpacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListPinSpacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesHeaders) GoString() string {
	return s.String()
}

func (s *ListPinSpacesHeaders) SetCommonHeaders(v map[string]*string) *ListPinSpacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListPinSpacesHeaders) SetXAcsDingtalkAccessToken(v string) *ListPinSpacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListPinSpacesRequest struct {
	Option     *ListPinSpacesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	OperatorId *string                     `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ListPinSpacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesRequest) GoString() string {
	return s.String()
}

func (s *ListPinSpacesRequest) SetOption(v *ListPinSpacesRequestOption) *ListPinSpacesRequest {
	s.Option = v
	return s
}

func (s *ListPinSpacesRequest) SetOperatorId(v string) *ListPinSpacesRequest {
	s.OperatorId = &v
	return s
}

type ListPinSpacesRequestOption struct {
	MaxResults              *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken               *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	WithSpaceCreatorInfo    *bool   `json:"withSpaceCreatorInfo,omitempty" xml:"withSpaceCreatorInfo,omitempty"`
	WithSpaceModifierInfo   *bool   `json:"withSpaceModifierInfo,omitempty" xml:"withSpaceModifierInfo,omitempty"`
	WithSpacePermissionRole *bool   `json:"withSpacePermissionRole,omitempty" xml:"withSpacePermissionRole,omitempty"`
	WithTeamDetail          *bool   `json:"withTeamDetail,omitempty" xml:"withTeamDetail,omitempty"`
}

func (s ListPinSpacesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesRequestOption) GoString() string {
	return s.String()
}

func (s *ListPinSpacesRequestOption) SetMaxResults(v int32) *ListPinSpacesRequestOption {
	s.MaxResults = &v
	return s
}

func (s *ListPinSpacesRequestOption) SetNextToken(v string) *ListPinSpacesRequestOption {
	s.NextToken = &v
	return s
}

func (s *ListPinSpacesRequestOption) SetWithSpaceCreatorInfo(v bool) *ListPinSpacesRequestOption {
	s.WithSpaceCreatorInfo = &v
	return s
}

func (s *ListPinSpacesRequestOption) SetWithSpaceModifierInfo(v bool) *ListPinSpacesRequestOption {
	s.WithSpaceModifierInfo = &v
	return s
}

func (s *ListPinSpacesRequestOption) SetWithSpacePermissionRole(v bool) *ListPinSpacesRequestOption {
	s.WithSpacePermissionRole = &v
	return s
}

func (s *ListPinSpacesRequestOption) SetWithTeamDetail(v bool) *ListPinSpacesRequestOption {
	s.WithTeamDetail = &v
	return s
}

type ListPinSpacesResponseBody struct {
	NextToken   *string                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResultItems []*ListPinSpacesResponseBodyResultItems `json:"resultItems,omitempty" xml:"resultItems,omitempty" type:"Repeated"`
}

func (s ListPinSpacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPinSpacesResponseBody) SetNextToken(v string) *ListPinSpacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPinSpacesResponseBody) SetResultItems(v []*ListPinSpacesResponseBodyResultItems) *ListPinSpacesResponseBody {
	s.ResultItems = v
	return s
}

type ListPinSpacesResponseBodyResultItems struct {
	CreateTime          *string                                        `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ModifiedTime        *string                                        `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	SpaceInfo           *ListPinSpacesResponseBodyResultItemsSpaceInfo `json:"spaceInfo,omitempty" xml:"spaceInfo,omitempty" type:"Struct"`
	SpacePermissionRole *string                                        `json:"spacePermissionRole,omitempty" xml:"spacePermissionRole,omitempty"`
	TeamInfo            *ListPinSpacesResponseBodyResultItemsTeamInfo  `json:"teamInfo,omitempty" xml:"teamInfo,omitempty" type:"Struct"`
}

func (s ListPinSpacesResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *ListPinSpacesResponseBodyResultItems) SetCreateTime(v string) *ListPinSpacesResponseBodyResultItems {
	s.CreateTime = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItems) SetModifiedTime(v string) *ListPinSpacesResponseBodyResultItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItems) SetSpaceInfo(v *ListPinSpacesResponseBodyResultItemsSpaceInfo) *ListPinSpacesResponseBodyResultItems {
	s.SpaceInfo = v
	return s
}

func (s *ListPinSpacesResponseBodyResultItems) SetSpacePermissionRole(v string) *ListPinSpacesResponseBodyResultItems {
	s.SpacePermissionRole = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItems) SetTeamInfo(v *ListPinSpacesResponseBodyResultItemsTeamInfo) *ListPinSpacesResponseBodyResultItems {
	s.TeamInfo = v
	return s
}

type ListPinSpacesResponseBodyResultItemsSpaceInfo struct {
	Cover        *string                                                `json:"cover,omitempty" xml:"cover,omitempty"`
	CreateTime   *string                                                `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator      *ListPinSpacesResponseBodyResultItemsSpaceInfoCreator  `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	Description  *string                                                `json:"description,omitempty" xml:"description,omitempty"`
	IconVO       *ListPinSpacesResponseBodyResultItemsSpaceInfoIconVO   `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	Id           *string                                                `json:"id,omitempty" xml:"id,omitempty"`
	MobileUrl    *string                                                `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	ModifiedTime *string                                                `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Modifier     *ListPinSpacesResponseBodyResultItemsSpaceInfoModifier `json:"modifier,omitempty" xml:"modifier,omitempty" type:"Struct"`
	Name         *string                                                `json:"name,omitempty" xml:"name,omitempty"`
	PcUrl        *string                                                `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s ListPinSpacesResponseBodyResultItemsSpaceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesResponseBodyResultItemsSpaceInfo) GoString() string {
	return s.String()
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetCover(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.Cover = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetCreateTime(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetCreator(v *ListPinSpacesResponseBodyResultItemsSpaceInfoCreator) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.Creator = v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetDescription(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.Description = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetIconVO(v *ListPinSpacesResponseBodyResultItemsSpaceInfoIconVO) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.IconVO = v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetId(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.Id = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetMobileUrl(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.MobileUrl = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetModifiedTime(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.ModifiedTime = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetModifier(v *ListPinSpacesResponseBodyResultItemsSpaceInfoModifier) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.Modifier = v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetName(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.Name = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfo) SetPcUrl(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfo {
	s.PcUrl = &v
	return s
}

type ListPinSpacesResponseBodyResultItemsSpaceInfoCreator struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListPinSpacesResponseBodyResultItemsSpaceInfoCreator) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesResponseBodyResultItemsSpaceInfoCreator) GoString() string {
	return s.String()
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfoCreator) SetName(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfoCreator {
	s.Name = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfoCreator) SetUserId(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfoCreator {
	s.UserId = &v
	return s
}

type ListPinSpacesResponseBodyResultItemsSpaceInfoIconVO struct {
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPinSpacesResponseBodyResultItemsSpaceInfoIconVO) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesResponseBodyResultItemsSpaceInfoIconVO) GoString() string {
	return s.String()
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfoIconVO) SetIcon(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfoIconVO {
	s.Icon = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfoIconVO) SetType(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfoIconVO {
	s.Type = &v
	return s
}

type ListPinSpacesResponseBodyResultItemsSpaceInfoModifier struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListPinSpacesResponseBodyResultItemsSpaceInfoModifier) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesResponseBodyResultItemsSpaceInfoModifier) GoString() string {
	return s.String()
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfoModifier) SetName(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfoModifier {
	s.Name = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsSpaceInfoModifier) SetUserId(v string) *ListPinSpacesResponseBodyResultItemsSpaceInfoModifier {
	s.UserId = &v
	return s
}

type ListPinSpacesResponseBodyResultItemsTeamInfo struct {
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPinSpacesResponseBodyResultItemsTeamInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesResponseBodyResultItemsTeamInfo) GoString() string {
	return s.String()
}

func (s *ListPinSpacesResponseBodyResultItemsTeamInfo) SetId(v string) *ListPinSpacesResponseBodyResultItemsTeamInfo {
	s.Id = &v
	return s
}

func (s *ListPinSpacesResponseBodyResultItemsTeamInfo) SetName(v string) *ListPinSpacesResponseBodyResultItemsTeamInfo {
	s.Name = &v
	return s
}

type ListPinSpacesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPinSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPinSpacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPinSpacesResponse) GoString() string {
	return s.String()
}

func (s *ListPinSpacesResponse) SetHeaders(v map[string]*string) *ListPinSpacesResponse {
	s.Headers = v
	return s
}

func (s *ListPinSpacesResponse) SetStatusCode(v int32) *ListPinSpacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPinSpacesResponse) SetBody(v *ListPinSpacesResponseBody) *ListPinSpacesResponse {
	s.Body = v
	return s
}

type ListRelatedSpaceTeamsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListRelatedSpaceTeamsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListRelatedSpaceTeamsHeaders) GoString() string {
	return s.String()
}

func (s *ListRelatedSpaceTeamsHeaders) SetCommonHeaders(v map[string]*string) *ListRelatedSpaceTeamsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRelatedSpaceTeamsHeaders) SetXAcsDingtalkAccessToken(v string) *ListRelatedSpaceTeamsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListRelatedSpaceTeamsRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Type       *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRelatedSpaceTeamsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRelatedSpaceTeamsRequest) GoString() string {
	return s.String()
}

func (s *ListRelatedSpaceTeamsRequest) SetOperatorId(v string) *ListRelatedSpaceTeamsRequest {
	s.OperatorId = &v
	return s
}

func (s *ListRelatedSpaceTeamsRequest) SetType(v int32) *ListRelatedSpaceTeamsRequest {
	s.Type = &v
	return s
}

type ListRelatedSpaceTeamsResponseBody struct {
	Items []*TeamModel `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListRelatedSpaceTeamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRelatedSpaceTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRelatedSpaceTeamsResponseBody) SetItems(v []*TeamModel) *ListRelatedSpaceTeamsResponseBody {
	s.Items = v
	return s
}

type ListRelatedSpaceTeamsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRelatedSpaceTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRelatedSpaceTeamsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRelatedSpaceTeamsResponse) GoString() string {
	return s.String()
}

func (s *ListRelatedSpaceTeamsResponse) SetHeaders(v map[string]*string) *ListRelatedSpaceTeamsResponse {
	s.Headers = v
	return s
}

func (s *ListRelatedSpaceTeamsResponse) SetStatusCode(v int32) *ListRelatedSpaceTeamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRelatedSpaceTeamsResponse) SetBody(v *ListRelatedSpaceTeamsResponseBody) *ListRelatedSpaceTeamsResponse {
	s.Body = v
	return s
}

type ListRelatedTeamsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListRelatedTeamsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListRelatedTeamsHeaders) GoString() string {
	return s.String()
}

func (s *ListRelatedTeamsHeaders) SetCommonHeaders(v map[string]*string) *ListRelatedTeamsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRelatedTeamsHeaders) SetXAcsDingtalkAccessToken(v string) *ListRelatedTeamsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListRelatedTeamsRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Type       *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRelatedTeamsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRelatedTeamsRequest) GoString() string {
	return s.String()
}

func (s *ListRelatedTeamsRequest) SetMaxResults(v int32) *ListRelatedTeamsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRelatedTeamsRequest) SetNextToken(v string) *ListRelatedTeamsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRelatedTeamsRequest) SetOperatorId(v string) *ListRelatedTeamsRequest {
	s.OperatorId = &v
	return s
}

func (s *ListRelatedTeamsRequest) SetType(v int32) *ListRelatedTeamsRequest {
	s.Type = &v
	return s
}

type ListRelatedTeamsResponseBody struct {
	HasMore   *bool        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Items     []*TeamModel `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextToken *string      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListRelatedTeamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRelatedTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRelatedTeamsResponseBody) SetHasMore(v bool) *ListRelatedTeamsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListRelatedTeamsResponseBody) SetItems(v []*TeamModel) *ListRelatedTeamsResponseBody {
	s.Items = v
	return s
}

func (s *ListRelatedTeamsResponseBody) SetNextToken(v string) *ListRelatedTeamsResponseBody {
	s.NextToken = &v
	return s
}

type ListRelatedTeamsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRelatedTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRelatedTeamsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRelatedTeamsResponse) GoString() string {
	return s.String()
}

func (s *ListRelatedTeamsResponse) SetHeaders(v map[string]*string) *ListRelatedTeamsResponse {
	s.Headers = v
	return s
}

func (s *ListRelatedTeamsResponse) SetStatusCode(v int32) *ListRelatedTeamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRelatedTeamsResponse) SetBody(v *ListRelatedTeamsResponseBody) *ListRelatedTeamsResponse {
	s.Body = v
	return s
}

type ListSpaceSectionsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSpaceSectionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceSectionsHeaders) GoString() string {
	return s.String()
}

func (s *ListSpaceSectionsHeaders) SetCommonHeaders(v map[string]*string) *ListSpaceSectionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSpaceSectionsHeaders) SetXAcsDingtalkAccessToken(v string) *ListSpaceSectionsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSpaceSectionsRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ListSpaceSectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceSectionsRequest) GoString() string {
	return s.String()
}

func (s *ListSpaceSectionsRequest) SetOperatorId(v string) *ListSpaceSectionsRequest {
	s.OperatorId = &v
	return s
}

type ListSpaceSectionsResponseBody struct {
	Items []*ListSpaceSectionsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListSpaceSectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceSectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpaceSectionsResponseBody) SetItems(v []*ListSpaceSectionsResponseBodyItems) *ListSpaceSectionsResponseBody {
	s.Items = v
	return s
}

type ListSpaceSectionsResponseBodyItems struct {
	DisplayType *string       `json:"displayType,omitempty" xml:"displayType,omitempty"`
	Id          *string       `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string       `json:"name,omitempty" xml:"name,omitempty"`
	SpaceNum    *int32        `json:"spaceNum,omitempty" xml:"spaceNum,omitempty"`
	Spaces      []*SpaceModel `json:"spaces,omitempty" xml:"spaces,omitempty" type:"Repeated"`
}

func (s ListSpaceSectionsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceSectionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSpaceSectionsResponseBodyItems) SetDisplayType(v string) *ListSpaceSectionsResponseBodyItems {
	s.DisplayType = &v
	return s
}

func (s *ListSpaceSectionsResponseBodyItems) SetId(v string) *ListSpaceSectionsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListSpaceSectionsResponseBodyItems) SetName(v string) *ListSpaceSectionsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListSpaceSectionsResponseBodyItems) SetSpaceNum(v int32) *ListSpaceSectionsResponseBodyItems {
	s.SpaceNum = &v
	return s
}

func (s *ListSpaceSectionsResponseBodyItems) SetSpaces(v []*SpaceModel) *ListSpaceSectionsResponseBodyItems {
	s.Spaces = v
	return s
}

type ListSpaceSectionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSpaceSectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSpaceSectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceSectionsResponse) GoString() string {
	return s.String()
}

func (s *ListSpaceSectionsResponse) SetHeaders(v map[string]*string) *ListSpaceSectionsResponse {
	s.Headers = v
	return s
}

func (s *ListSpaceSectionsResponse) SetStatusCode(v int32) *ListSpaceSectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSpaceSectionsResponse) SetBody(v *ListSpaceSectionsResponseBody) *ListSpaceSectionsResponse {
	s.Body = v
	return s
}

type ListStarsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListStarsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListStarsHeaders) GoString() string {
	return s.String()
}

func (s *ListStarsHeaders) SetCommonHeaders(v map[string]*string) *ListStarsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListStarsHeaders) SetXAcsDingtalkAccessToken(v string) *ListStarsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListStarsRequest struct {
	Option     *ListStarsRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	OperatorId *string                 `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ListStarsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStarsRequest) GoString() string {
	return s.String()
}

func (s *ListStarsRequest) SetOption(v *ListStarsRequestOption) *ListStarsRequest {
	s.Option = v
	return s
}

func (s *ListStarsRequest) SetOperatorId(v string) *ListStarsRequest {
	s.OperatorId = &v
	return s
}

type ListStarsRequestOption struct {
	FilterDocTypes           []*string `json:"filterDocTypes,omitempty" xml:"filterDocTypes,omitempty" type:"Repeated"`
	MaxResults               *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken                *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Order                    *string   `json:"order,omitempty" xml:"order,omitempty"`
	OrderBy                  *string   `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	WithDentryCreatorInfo    *bool     `json:"withDentryCreatorInfo,omitempty" xml:"withDentryCreatorInfo,omitempty"`
	WithDentryModifierInfo   *bool     `json:"withDentryModifierInfo,omitempty" xml:"withDentryModifierInfo,omitempty"`
	WithDentryPermissionRole *bool     `json:"withDentryPermissionRole,omitempty" xml:"withDentryPermissionRole,omitempty"`
	WithSpaceDetail          *bool     `json:"withSpaceDetail,omitempty" xml:"withSpaceDetail,omitempty"`
	WithSpacePermissionRole  *bool     `json:"withSpacePermissionRole,omitempty" xml:"withSpacePermissionRole,omitempty"`
	WithTeamDetail           *bool     `json:"withTeamDetail,omitempty" xml:"withTeamDetail,omitempty"`
}

func (s ListStarsRequestOption) String() string {
	return tea.Prettify(s)
}

func (s ListStarsRequestOption) GoString() string {
	return s.String()
}

func (s *ListStarsRequestOption) SetFilterDocTypes(v []*string) *ListStarsRequestOption {
	s.FilterDocTypes = v
	return s
}

func (s *ListStarsRequestOption) SetMaxResults(v int32) *ListStarsRequestOption {
	s.MaxResults = &v
	return s
}

func (s *ListStarsRequestOption) SetNextToken(v string) *ListStarsRequestOption {
	s.NextToken = &v
	return s
}

func (s *ListStarsRequestOption) SetOrder(v string) *ListStarsRequestOption {
	s.Order = &v
	return s
}

func (s *ListStarsRequestOption) SetOrderBy(v string) *ListStarsRequestOption {
	s.OrderBy = &v
	return s
}

func (s *ListStarsRequestOption) SetWithDentryCreatorInfo(v bool) *ListStarsRequestOption {
	s.WithDentryCreatorInfo = &v
	return s
}

func (s *ListStarsRequestOption) SetWithDentryModifierInfo(v bool) *ListStarsRequestOption {
	s.WithDentryModifierInfo = &v
	return s
}

func (s *ListStarsRequestOption) SetWithDentryPermissionRole(v bool) *ListStarsRequestOption {
	s.WithDentryPermissionRole = &v
	return s
}

func (s *ListStarsRequestOption) SetWithSpaceDetail(v bool) *ListStarsRequestOption {
	s.WithSpaceDetail = &v
	return s
}

func (s *ListStarsRequestOption) SetWithSpacePermissionRole(v bool) *ListStarsRequestOption {
	s.WithSpacePermissionRole = &v
	return s
}

func (s *ListStarsRequestOption) SetWithTeamDetail(v bool) *ListStarsRequestOption {
	s.WithTeamDetail = &v
	return s
}

type ListStarsResponseBody struct {
	NextToken *string                          `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StarList  []*ListStarsResponseBodyStarList `json:"starList,omitempty" xml:"starList,omitempty" type:"Repeated"`
}

func (s ListStarsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStarsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStarsResponseBody) SetNextToken(v string) *ListStarsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListStarsResponseBody) SetStarList(v []*ListStarsResponseBodyStarList) *ListStarsResponseBody {
	s.StarList = v
	return s
}

type ListStarsResponseBodyStarList struct {
	CreateTime           *string                                  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DentryInfo           *ListStarsResponseBodyStarListDentryInfo `json:"dentryInfo,omitempty" xml:"dentryInfo,omitempty" type:"Struct"`
	DentryPermissionRole *string                                  `json:"dentryPermissionRole,omitempty" xml:"dentryPermissionRole,omitempty"`
	Id                   *string                                  `json:"id,omitempty" xml:"id,omitempty"`
	IsDeleted            *bool                                    `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	ModifiedTime         *string                                  `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	SpaceInfo            *ListStarsResponseBodyStarListSpaceInfo  `json:"spaceInfo,omitempty" xml:"spaceInfo,omitempty" type:"Struct"`
	SpacePermissionRole  *string                                  `json:"spacePermissionRole,omitempty" xml:"spacePermissionRole,omitempty"`
	StarType             *string                                  `json:"starType,omitempty" xml:"starType,omitempty"`
	TeamInfo             *ListStarsResponseBodyStarListTeamInfo   `json:"teamInfo,omitempty" xml:"teamInfo,omitempty" type:"Struct"`
}

func (s ListStarsResponseBodyStarList) String() string {
	return tea.Prettify(s)
}

func (s ListStarsResponseBodyStarList) GoString() string {
	return s.String()
}

func (s *ListStarsResponseBodyStarList) SetCreateTime(v string) *ListStarsResponseBodyStarList {
	s.CreateTime = &v
	return s
}

func (s *ListStarsResponseBodyStarList) SetDentryInfo(v *ListStarsResponseBodyStarListDentryInfo) *ListStarsResponseBodyStarList {
	s.DentryInfo = v
	return s
}

func (s *ListStarsResponseBodyStarList) SetDentryPermissionRole(v string) *ListStarsResponseBodyStarList {
	s.DentryPermissionRole = &v
	return s
}

func (s *ListStarsResponseBodyStarList) SetId(v string) *ListStarsResponseBodyStarList {
	s.Id = &v
	return s
}

func (s *ListStarsResponseBodyStarList) SetIsDeleted(v bool) *ListStarsResponseBodyStarList {
	s.IsDeleted = &v
	return s
}

func (s *ListStarsResponseBodyStarList) SetModifiedTime(v string) *ListStarsResponseBodyStarList {
	s.ModifiedTime = &v
	return s
}

func (s *ListStarsResponseBodyStarList) SetSpaceInfo(v *ListStarsResponseBodyStarListSpaceInfo) *ListStarsResponseBodyStarList {
	s.SpaceInfo = v
	return s
}

func (s *ListStarsResponseBodyStarList) SetSpacePermissionRole(v string) *ListStarsResponseBodyStarList {
	s.SpacePermissionRole = &v
	return s
}

func (s *ListStarsResponseBodyStarList) SetStarType(v string) *ListStarsResponseBodyStarList {
	s.StarType = &v
	return s
}

func (s *ListStarsResponseBodyStarList) SetTeamInfo(v *ListStarsResponseBodyStarListTeamInfo) *ListStarsResponseBodyStarList {
	s.TeamInfo = v
	return s
}

type ListStarsResponseBodyStarListDentryInfo struct {
	CreateTime   *string                                          `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator      *ListStarsResponseBodyStarListDentryInfoCreator  `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	Extension    *string                                          `json:"extension,omitempty" xml:"extension,omitempty"`
	Id           *string                                          `json:"id,omitempty" xml:"id,omitempty"`
	MobileUrl    *string                                          `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	ModifiedTime *string                                          `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Modifier     *ListStarsResponseBodyStarListDentryInfoModifier `json:"modifier,omitempty" xml:"modifier,omitempty" type:"Struct"`
	Name         *string                                          `json:"name,omitempty" xml:"name,omitempty"`
	PcUrl        *string                                          `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	SpaceId      *string                                          `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status       *string                                          `json:"status,omitempty" xml:"status,omitempty"`
	Type         *string                                          `json:"type,omitempty" xml:"type,omitempty"`
	Uuid         *string                                          `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s ListStarsResponseBodyStarListDentryInfo) String() string {
	return tea.Prettify(s)
}

func (s ListStarsResponseBodyStarListDentryInfo) GoString() string {
	return s.String()
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetCreateTime(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.CreateTime = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetCreator(v *ListStarsResponseBodyStarListDentryInfoCreator) *ListStarsResponseBodyStarListDentryInfo {
	s.Creator = v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetExtension(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.Extension = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetId(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.Id = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetMobileUrl(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.MobileUrl = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetModifiedTime(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.ModifiedTime = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetModifier(v *ListStarsResponseBodyStarListDentryInfoModifier) *ListStarsResponseBodyStarListDentryInfo {
	s.Modifier = v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetName(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.Name = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetPcUrl(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.PcUrl = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetSpaceId(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.SpaceId = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetStatus(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.Status = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetType(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.Type = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfo) SetUuid(v string) *ListStarsResponseBodyStarListDentryInfo {
	s.Uuid = &v
	return s
}

type ListStarsResponseBodyStarListDentryInfoCreator struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListStarsResponseBodyStarListDentryInfoCreator) String() string {
	return tea.Prettify(s)
}

func (s ListStarsResponseBodyStarListDentryInfoCreator) GoString() string {
	return s.String()
}

func (s *ListStarsResponseBodyStarListDentryInfoCreator) SetName(v string) *ListStarsResponseBodyStarListDentryInfoCreator {
	s.Name = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfoCreator) SetUserId(v string) *ListStarsResponseBodyStarListDentryInfoCreator {
	s.UserId = &v
	return s
}

type ListStarsResponseBodyStarListDentryInfoModifier struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListStarsResponseBodyStarListDentryInfoModifier) String() string {
	return tea.Prettify(s)
}

func (s ListStarsResponseBodyStarListDentryInfoModifier) GoString() string {
	return s.String()
}

func (s *ListStarsResponseBodyStarListDentryInfoModifier) SetName(v string) *ListStarsResponseBodyStarListDentryInfoModifier {
	s.Name = &v
	return s
}

func (s *ListStarsResponseBodyStarListDentryInfoModifier) SetUserId(v string) *ListStarsResponseBodyStarListDentryInfoModifier {
	s.UserId = &v
	return s
}

type ListStarsResponseBodyStarListSpaceInfo struct {
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListStarsResponseBodyStarListSpaceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListStarsResponseBodyStarListSpaceInfo) GoString() string {
	return s.String()
}

func (s *ListStarsResponseBodyStarListSpaceInfo) SetId(v string) *ListStarsResponseBodyStarListSpaceInfo {
	s.Id = &v
	return s
}

func (s *ListStarsResponseBodyStarListSpaceInfo) SetName(v string) *ListStarsResponseBodyStarListSpaceInfo {
	s.Name = &v
	return s
}

type ListStarsResponseBodyStarListTeamInfo struct {
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListStarsResponseBodyStarListTeamInfo) String() string {
	return tea.Prettify(s)
}

func (s ListStarsResponseBodyStarListTeamInfo) GoString() string {
	return s.String()
}

func (s *ListStarsResponseBodyStarListTeamInfo) SetId(v string) *ListStarsResponseBodyStarListTeamInfo {
	s.Id = &v
	return s
}

func (s *ListStarsResponseBodyStarListTeamInfo) SetName(v string) *ListStarsResponseBodyStarListTeamInfo {
	s.Name = &v
	return s
}

type ListStarsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStarsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStarsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStarsResponse) GoString() string {
	return s.String()
}

func (s *ListStarsResponse) SetHeaders(v map[string]*string) *ListStarsResponse {
	s.Headers = v
	return s
}

func (s *ListStarsResponse) SetStatusCode(v int32) *ListStarsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStarsResponse) SetBody(v *ListStarsResponseBody) *ListStarsResponse {
	s.Body = v
	return s
}

type ListTeamMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListTeamMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTeamMembersHeaders) GoString() string {
	return s.String()
}

func (s *ListTeamMembersHeaders) SetCommonHeaders(v map[string]*string) *ListTeamMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTeamMembersHeaders) SetXAcsDingtalkAccessToken(v string) *ListTeamMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListTeamMembersRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ListTeamMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTeamMembersRequest) GoString() string {
	return s.String()
}

func (s *ListTeamMembersRequest) SetOperatorId(v string) *ListTeamMembersRequest {
	s.OperatorId = &v
	return s
}

type ListTeamMembersResponseBody struct {
	Members  []*ListTeamMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	TeamName *string                               `json:"teamName,omitempty" xml:"teamName,omitempty"`
}

func (s ListTeamMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTeamMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListTeamMembersResponseBody) SetMembers(v []*ListTeamMembersResponseBodyMembers) *ListTeamMembersResponseBody {
	s.Members = v
	return s
}

func (s *ListTeamMembersResponseBody) SetTeamName(v string) *ListTeamMembersResponseBody {
	s.TeamName = &v
	return s
}

type ListTeamMembersResponseBodyMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *int32  `json:"memberType,omitempty" xml:"memberType,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	RoleCode   *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s ListTeamMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s ListTeamMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListTeamMembersResponseBodyMembers) SetMemberId(v string) *ListTeamMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *ListTeamMembersResponseBodyMembers) SetMemberType(v int32) *ListTeamMembersResponseBodyMembers {
	s.MemberType = &v
	return s
}

func (s *ListTeamMembersResponseBodyMembers) SetName(v string) *ListTeamMembersResponseBodyMembers {
	s.Name = &v
	return s
}

func (s *ListTeamMembersResponseBodyMembers) SetRoleCode(v string) *ListTeamMembersResponseBodyMembers {
	s.RoleCode = &v
	return s
}

type ListTeamMembersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTeamMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTeamMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTeamMembersResponse) GoString() string {
	return s.String()
}

func (s *ListTeamMembersResponse) SetHeaders(v map[string]*string) *ListTeamMembersResponse {
	s.Headers = v
	return s
}

func (s *ListTeamMembersResponse) SetStatusCode(v int32) *ListTeamMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTeamMembersResponse) SetBody(v *ListTeamMembersResponseBody) *ListTeamMembersResponse {
	s.Body = v
	return s
}

type MarkStarHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MarkStarHeaders) String() string {
	return tea.Prettify(s)
}

func (s MarkStarHeaders) GoString() string {
	return s.String()
}

func (s *MarkStarHeaders) SetCommonHeaders(v map[string]*string) *MarkStarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MarkStarHeaders) SetXAcsDingtalkAccessToken(v string) *MarkStarHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MarkStarRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s MarkStarRequest) String() string {
	return tea.Prettify(s)
}

func (s MarkStarRequest) GoString() string {
	return s.String()
}

func (s *MarkStarRequest) SetOperatorId(v string) *MarkStarRequest {
	s.OperatorId = &v
	return s
}

type MarkStarResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MarkStarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MarkStarResponseBody) GoString() string {
	return s.String()
}

func (s *MarkStarResponseBody) SetSuccess(v bool) *MarkStarResponseBody {
	s.Success = &v
	return s
}

type MarkStarResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MarkStarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MarkStarResponse) String() string {
	return tea.Prettify(s)
}

func (s MarkStarResponse) GoString() string {
	return s.String()
}

func (s *MarkStarResponse) SetHeaders(v map[string]*string) *MarkStarResponse {
	s.Headers = v
	return s
}

func (s *MarkStarResponse) SetStatusCode(v int32) *MarkStarResponse {
	s.StatusCode = &v
	return s
}

func (s *MarkStarResponse) SetBody(v *MarkStarResponseBody) *MarkStarResponse {
	s.Body = v
	return s
}

type MoveDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MoveDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryHeaders) GoString() string {
	return s.String()
}

func (s *MoveDentryHeaders) SetCommonHeaders(v map[string]*string) *MoveDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MoveDentryHeaders) SetXAcsDingtalkAccessToken(v string) *MoveDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MoveDentryRequest struct {
	OperatorId       *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	TargetSpaceId    *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	ToNextDentryId   *string `json:"toNextDentryId,omitempty" xml:"toNextDentryId,omitempty"`
	ToParentDentryId *string `json:"toParentDentryId,omitempty" xml:"toParentDentryId,omitempty"`
	ToPrevDentryId   *string `json:"toPrevDentryId,omitempty" xml:"toPrevDentryId,omitempty"`
}

func (s MoveDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryRequest) GoString() string {
	return s.String()
}

func (s *MoveDentryRequest) SetOperatorId(v string) *MoveDentryRequest {
	s.OperatorId = &v
	return s
}

func (s *MoveDentryRequest) SetTargetSpaceId(v string) *MoveDentryRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *MoveDentryRequest) SetToNextDentryId(v string) *MoveDentryRequest {
	s.ToNextDentryId = &v
	return s
}

func (s *MoveDentryRequest) SetToParentDentryId(v string) *MoveDentryRequest {
	s.ToParentDentryId = &v
	return s
}

func (s *MoveDentryRequest) SetToPrevDentryId(v string) *MoveDentryRequest {
	s.ToPrevDentryId = &v
	return s
}

type MoveDentryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DentryVO          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryResponse) GoString() string {
	return s.String()
}

func (s *MoveDentryResponse) SetHeaders(v map[string]*string) *MoveDentryResponse {
	s.Headers = v
	return s
}

func (s *MoveDentryResponse) SetStatusCode(v int32) *MoveDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveDentryResponse) SetBody(v *DentryVO) *MoveDentryResponse {
	s.Body = v
	return s
}

type PinSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PinSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s PinSpaceHeaders) GoString() string {
	return s.String()
}

func (s *PinSpaceHeaders) SetCommonHeaders(v map[string]*string) *PinSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PinSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *PinSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PinSpaceRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s PinSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s PinSpaceRequest) GoString() string {
	return s.String()
}

func (s *PinSpaceRequest) SetOperatorId(v string) *PinSpaceRequest {
	s.OperatorId = &v
	return s
}

type PinSpaceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PinSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PinSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *PinSpaceResponseBody) SetSuccess(v bool) *PinSpaceResponseBody {
	s.Success = &v
	return s
}

type PinSpaceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PinSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PinSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s PinSpaceResponse) GoString() string {
	return s.String()
}

func (s *PinSpaceResponse) SetHeaders(v map[string]*string) *PinSpaceResponse {
	s.Headers = v
	return s
}

func (s *PinSpaceResponse) SetStatusCode(v int32) *PinSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *PinSpaceResponse) SetBody(v *PinSpaceResponseBody) *PinSpaceResponse {
	s.Body = v
	return s
}

type QueryDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDentryHeaders) GoString() string {
	return s.String()
}

func (s *QueryDentryHeaders) SetCommonHeaders(v map[string]*string) *QueryDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDentryHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDentryRequest struct {
	IncludeSpace *bool   `json:"includeSpace,omitempty" xml:"includeSpace,omitempty"`
	OperatorId   *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s QueryDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDentryRequest) GoString() string {
	return s.String()
}

func (s *QueryDentryRequest) SetIncludeSpace(v bool) *QueryDentryRequest {
	s.IncludeSpace = &v
	return s
}

func (s *QueryDentryRequest) SetOperatorId(v string) *QueryDentryRequest {
	s.OperatorId = &v
	return s
}

type QueryDentryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DentryVO          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDentryResponse) GoString() string {
	return s.String()
}

func (s *QueryDentryResponse) SetHeaders(v map[string]*string) *QueryDentryResponse {
	s.Headers = v
	return s
}

func (s *QueryDentryResponse) SetStatusCode(v int32) *QueryDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDentryResponse) SetBody(v *DentryVO) *QueryDentryResponse {
	s.Body = v
	return s
}

type QueryItemByUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryItemByUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryItemByUrlHeaders) GoString() string {
	return s.String()
}

func (s *QueryItemByUrlHeaders) SetCommonHeaders(v map[string]*string) *QueryItemByUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryItemByUrlHeaders) SetXAcsDingtalkAccessToken(v string) *QueryItemByUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryItemByUrlRequest struct {
	OperatorId          *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Url                 *string `json:"url,omitempty" xml:"url,omitempty"`
	WithStatisticalInfo *bool   `json:"withStatisticalInfo,omitempty" xml:"withStatisticalInfo,omitempty"`
}

func (s QueryItemByUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemByUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryItemByUrlRequest) SetOperatorId(v string) *QueryItemByUrlRequest {
	s.OperatorId = &v
	return s
}

func (s *QueryItemByUrlRequest) SetUrl(v string) *QueryItemByUrlRequest {
	s.Url = &v
	return s
}

func (s *QueryItemByUrlRequest) SetWithStatisticalInfo(v bool) *QueryItemByUrlRequest {
	s.WithStatisticalInfo = &v
	return s
}

type QueryItemByUrlResponseBody struct {
	BizType      *string                          `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Dentry       *DentryModel                     `json:"dentry,omitempty" xml:"dentry,omitempty"`
	ResourceType *string                          `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Space        *QueryItemByUrlResponseBodySpace `json:"space,omitempty" xml:"space,omitempty" type:"Struct"`
}

func (s QueryItemByUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryItemByUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryItemByUrlResponseBody) SetBizType(v string) *QueryItemByUrlResponseBody {
	s.BizType = &v
	return s
}

func (s *QueryItemByUrlResponseBody) SetDentry(v *DentryModel) *QueryItemByUrlResponseBody {
	s.Dentry = v
	return s
}

func (s *QueryItemByUrlResponseBody) SetResourceType(v string) *QueryItemByUrlResponseBody {
	s.ResourceType = &v
	return s
}

func (s *QueryItemByUrlResponseBody) SetSpace(v *QueryItemByUrlResponseBodySpace) *QueryItemByUrlResponseBody {
	s.Space = v
	return s
}

type QueryItemByUrlResponseBodySpace struct {
	Description *string                               `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string                               `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string                               `json:"name,omitempty" xml:"name,omitempty"`
	Owner       *QueryItemByUrlResponseBodySpaceOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	Type        *int32                                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryItemByUrlResponseBodySpace) String() string {
	return tea.Prettify(s)
}

func (s QueryItemByUrlResponseBodySpace) GoString() string {
	return s.String()
}

func (s *QueryItemByUrlResponseBodySpace) SetDescription(v string) *QueryItemByUrlResponseBodySpace {
	s.Description = &v
	return s
}

func (s *QueryItemByUrlResponseBodySpace) SetId(v string) *QueryItemByUrlResponseBodySpace {
	s.Id = &v
	return s
}

func (s *QueryItemByUrlResponseBodySpace) SetName(v string) *QueryItemByUrlResponseBodySpace {
	s.Name = &v
	return s
}

func (s *QueryItemByUrlResponseBodySpace) SetOwner(v *QueryItemByUrlResponseBodySpaceOwner) *QueryItemByUrlResponseBodySpace {
	s.Owner = v
	return s
}

func (s *QueryItemByUrlResponseBodySpace) SetType(v int32) *QueryItemByUrlResponseBodySpace {
	s.Type = &v
	return s
}

type QueryItemByUrlResponseBodySpaceOwner struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryItemByUrlResponseBodySpaceOwner) String() string {
	return tea.Prettify(s)
}

func (s QueryItemByUrlResponseBodySpaceOwner) GoString() string {
	return s.String()
}

func (s *QueryItemByUrlResponseBodySpaceOwner) SetName(v string) *QueryItemByUrlResponseBodySpaceOwner {
	s.Name = &v
	return s
}

func (s *QueryItemByUrlResponseBodySpaceOwner) SetUnionId(v string) *QueryItemByUrlResponseBodySpaceOwner {
	s.UnionId = &v
	return s
}

type QueryItemByUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryItemByUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryItemByUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemByUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryItemByUrlResponse) SetHeaders(v map[string]*string) *QueryItemByUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryItemByUrlResponse) SetStatusCode(v int32) *QueryItemByUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryItemByUrlResponse) SetBody(v *QueryItemByUrlResponseBody) *QueryItemByUrlResponse {
	s.Body = v
	return s
}

type QueryMineSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMineSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMineSpaceHeaders) GoString() string {
	return s.String()
}

func (s *QueryMineSpaceHeaders) SetCommonHeaders(v map[string]*string) *QueryMineSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMineSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMineSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMineSpaceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SpaceVO           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMineSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMineSpaceResponse) GoString() string {
	return s.String()
}

func (s *QueryMineSpaceResponse) SetHeaders(v map[string]*string) *QueryMineSpaceResponse {
	s.Headers = v
	return s
}

func (s *QueryMineSpaceResponse) SetStatusCode(v int32) *QueryMineSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMineSpaceResponse) SetBody(v *SpaceVO) *QueryMineSpaceResponse {
	s.Body = v
	return s
}

type QueryRecentListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRecentListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentListHeaders) GoString() string {
	return s.String()
}

func (s *QueryRecentListHeaders) SetCommonHeaders(v map[string]*string) *QueryRecentListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRecentListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRecentListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRecentListRequest struct {
	CreatorType *int32  `json:"creatorType,omitempty" xml:"creatorType,omitempty"`
	FileType    *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperatorId  *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	RecentType  *int32  `json:"recentType,omitempty" xml:"recentType,omitempty"`
}

func (s QueryRecentListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentListRequest) GoString() string {
	return s.String()
}

func (s *QueryRecentListRequest) SetCreatorType(v int32) *QueryRecentListRequest {
	s.CreatorType = &v
	return s
}

func (s *QueryRecentListRequest) SetFileType(v int32) *QueryRecentListRequest {
	s.FileType = &v
	return s
}

func (s *QueryRecentListRequest) SetMaxResults(v int32) *QueryRecentListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryRecentListRequest) SetNextToken(v string) *QueryRecentListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryRecentListRequest) SetOperatorId(v string) *QueryRecentListRequest {
	s.OperatorId = &v
	return s
}

func (s *QueryRecentListRequest) SetRecentType(v int32) *QueryRecentListRequest {
	s.RecentType = &v
	return s
}

type QueryRecentListResponseBody struct {
	HasMore    *bool                                    `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken  *string                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RecentList []*QueryRecentListResponseBodyRecentList `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
}

func (s QueryRecentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecentListResponseBody) SetHasMore(v bool) *QueryRecentListResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryRecentListResponseBody) SetNextToken(v string) *QueryRecentListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryRecentListResponseBody) SetRecentList(v []*QueryRecentListResponseBodyRecentList) *QueryRecentListResponseBody {
	s.RecentList = v
	return s
}

type QueryRecentListResponseBodyRecentList struct {
	Deleted    *bool                                      `json:"deleted,omitempty" xml:"deleted,omitempty"`
	Dentry     *DentryModel                               `json:"dentry,omitempty" xml:"dentry,omitempty"`
	RecentTime *int64                                     `json:"recentTime,omitempty" xml:"recentTime,omitempty"`
	Team       *QueryRecentListResponseBodyRecentListTeam `json:"team,omitempty" xml:"team,omitempty" type:"Struct"`
}

func (s QueryRecentListResponseBodyRecentList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentListResponseBodyRecentList) GoString() string {
	return s.String()
}

func (s *QueryRecentListResponseBodyRecentList) SetDeleted(v bool) *QueryRecentListResponseBodyRecentList {
	s.Deleted = &v
	return s
}

func (s *QueryRecentListResponseBodyRecentList) SetDentry(v *DentryModel) *QueryRecentListResponseBodyRecentList {
	s.Dentry = v
	return s
}

func (s *QueryRecentListResponseBodyRecentList) SetRecentTime(v int64) *QueryRecentListResponseBodyRecentList {
	s.RecentTime = &v
	return s
}

func (s *QueryRecentListResponseBodyRecentList) SetTeam(v *QueryRecentListResponseBodyRecentListTeam) *QueryRecentListResponseBodyRecentList {
	s.Team = v
	return s
}

type QueryRecentListResponseBodyRecentListTeam struct {
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryRecentListResponseBodyRecentListTeam) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentListResponseBodyRecentListTeam) GoString() string {
	return s.String()
}

func (s *QueryRecentListResponseBodyRecentListTeam) SetId(v string) *QueryRecentListResponseBodyRecentListTeam {
	s.Id = &v
	return s
}

func (s *QueryRecentListResponseBodyRecentListTeam) SetName(v string) *QueryRecentListResponseBodyRecentListTeam {
	s.Name = &v
	return s
}

type QueryRecentListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRecentListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRecentListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentListResponse) GoString() string {
	return s.String()
}

func (s *QueryRecentListResponse) SetHeaders(v map[string]*string) *QueryRecentListResponse {
	s.Headers = v
	return s
}

func (s *QueryRecentListResponse) SetStatusCode(v int32) *QueryRecentListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecentListResponse) SetBody(v *QueryRecentListResponseBody) *QueryRecentListResponse {
	s.Body = v
	return s
}

type QuerySpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySpaceHeaders) GoString() string {
	return s.String()
}

func (s *QuerySpaceHeaders) SetCommonHeaders(v map[string]*string) *QuerySpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySpaceHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySpaceRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s QuerySpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySpaceRequest) GoString() string {
	return s.String()
}

func (s *QuerySpaceRequest) SetOperatorId(v string) *QuerySpaceRequest {
	s.OperatorId = &v
	return s
}

type QuerySpaceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SpaceVO           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySpaceResponse) GoString() string {
	return s.String()
}

func (s *QuerySpaceResponse) SetHeaders(v map[string]*string) *QuerySpaceResponse {
	s.Headers = v
	return s
}

func (s *QuerySpaceResponse) SetStatusCode(v int32) *QuerySpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySpaceResponse) SetBody(v *SpaceVO) *QuerySpaceResponse {
	s.Body = v
	return s
}

type RelatedSpacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RelatedSpacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s RelatedSpacesHeaders) GoString() string {
	return s.String()
}

func (s *RelatedSpacesHeaders) SetCommonHeaders(v map[string]*string) *RelatedSpacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RelatedSpacesHeaders) SetXAcsDingtalkAccessToken(v string) *RelatedSpacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RelatedSpacesRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	TeamId     *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
}

func (s RelatedSpacesRequest) String() string {
	return tea.Prettify(s)
}

func (s RelatedSpacesRequest) GoString() string {
	return s.String()
}

func (s *RelatedSpacesRequest) SetMaxResults(v int32) *RelatedSpacesRequest {
	s.MaxResults = &v
	return s
}

func (s *RelatedSpacesRequest) SetNextToken(v string) *RelatedSpacesRequest {
	s.NextToken = &v
	return s
}

func (s *RelatedSpacesRequest) SetOperatorId(v string) *RelatedSpacesRequest {
	s.OperatorId = &v
	return s
}

func (s *RelatedSpacesRequest) SetTeamId(v string) *RelatedSpacesRequest {
	s.TeamId = &v
	return s
}

type RelatedSpacesResponseBody struct {
	HasMore   *bool                             `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Items     []*RelatedSpacesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextToken *string                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s RelatedSpacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RelatedSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *RelatedSpacesResponseBody) SetHasMore(v bool) *RelatedSpacesResponseBody {
	s.HasMore = &v
	return s
}

func (s *RelatedSpacesResponseBody) SetItems(v []*RelatedSpacesResponseBodyItems) *RelatedSpacesResponseBody {
	s.Items = v
	return s
}

func (s *RelatedSpacesResponseBody) SetNextToken(v string) *RelatedSpacesResponseBody {
	s.NextToken = &v
	return s
}

type RelatedSpacesResponseBodyItems struct {
	Cover       *string                                    `json:"cover,omitempty" xml:"cover,omitempty"`
	Description *string                                    `json:"description,omitempty" xml:"description,omitempty"`
	HdIconVO    *RelatedSpacesResponseBodyItemsHdIconVO    `json:"hdIconVO,omitempty" xml:"hdIconVO,omitempty" type:"Struct"`
	IconVO      *RelatedSpacesResponseBodyItemsIconVO      `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	Id          *string                                    `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	Owner       *RelatedSpacesResponseBodyItemsOwner       `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	RecentList  []*DentryModel                             `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
	Type        *int32                                     `json:"type,omitempty" xml:"type,omitempty"`
	Url         *string                                    `json:"url,omitempty" xml:"url,omitempty"`
	VisitorInfo *RelatedSpacesResponseBodyItemsVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
}

func (s RelatedSpacesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s RelatedSpacesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *RelatedSpacesResponseBodyItems) SetCover(v string) *RelatedSpacesResponseBodyItems {
	s.Cover = &v
	return s
}

func (s *RelatedSpacesResponseBodyItems) SetDescription(v string) *RelatedSpacesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *RelatedSpacesResponseBodyItems) SetHdIconVO(v *RelatedSpacesResponseBodyItemsHdIconVO) *RelatedSpacesResponseBodyItems {
	s.HdIconVO = v
	return s
}

func (s *RelatedSpacesResponseBodyItems) SetIconVO(v *RelatedSpacesResponseBodyItemsIconVO) *RelatedSpacesResponseBodyItems {
	s.IconVO = v
	return s
}

func (s *RelatedSpacesResponseBodyItems) SetId(v string) *RelatedSpacesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *RelatedSpacesResponseBodyItems) SetName(v string) *RelatedSpacesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *RelatedSpacesResponseBodyItems) SetOwner(v *RelatedSpacesResponseBodyItemsOwner) *RelatedSpacesResponseBodyItems {
	s.Owner = v
	return s
}

func (s *RelatedSpacesResponseBodyItems) SetRecentList(v []*DentryModel) *RelatedSpacesResponseBodyItems {
	s.RecentList = v
	return s
}

func (s *RelatedSpacesResponseBodyItems) SetType(v int32) *RelatedSpacesResponseBodyItems {
	s.Type = &v
	return s
}

func (s *RelatedSpacesResponseBodyItems) SetUrl(v string) *RelatedSpacesResponseBodyItems {
	s.Url = &v
	return s
}

func (s *RelatedSpacesResponseBodyItems) SetVisitorInfo(v *RelatedSpacesResponseBodyItemsVisitorInfo) *RelatedSpacesResponseBodyItems {
	s.VisitorInfo = v
	return s
}

type RelatedSpacesResponseBodyItemsHdIconVO struct {
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RelatedSpacesResponseBodyItemsHdIconVO) String() string {
	return tea.Prettify(s)
}

func (s RelatedSpacesResponseBodyItemsHdIconVO) GoString() string {
	return s.String()
}

func (s *RelatedSpacesResponseBodyItemsHdIconVO) SetIcon(v string) *RelatedSpacesResponseBodyItemsHdIconVO {
	s.Icon = &v
	return s
}

func (s *RelatedSpacesResponseBodyItemsHdIconVO) SetType(v string) *RelatedSpacesResponseBodyItemsHdIconVO {
	s.Type = &v
	return s
}

type RelatedSpacesResponseBodyItemsIconVO struct {
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RelatedSpacesResponseBodyItemsIconVO) String() string {
	return tea.Prettify(s)
}

func (s RelatedSpacesResponseBodyItemsIconVO) GoString() string {
	return s.String()
}

func (s *RelatedSpacesResponseBodyItemsIconVO) SetIcon(v string) *RelatedSpacesResponseBodyItemsIconVO {
	s.Icon = &v
	return s
}

func (s *RelatedSpacesResponseBodyItemsIconVO) SetType(v string) *RelatedSpacesResponseBodyItemsIconVO {
	s.Type = &v
	return s
}

type RelatedSpacesResponseBodyItemsOwner struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RelatedSpacesResponseBodyItemsOwner) String() string {
	return tea.Prettify(s)
}

func (s RelatedSpacesResponseBodyItemsOwner) GoString() string {
	return s.String()
}

func (s *RelatedSpacesResponseBodyItemsOwner) SetName(v string) *RelatedSpacesResponseBodyItemsOwner {
	s.Name = &v
	return s
}

func (s *RelatedSpacesResponseBodyItemsOwner) SetUnionId(v string) *RelatedSpacesResponseBodyItemsOwner {
	s.UnionId = &v
	return s
}

type RelatedSpacesResponseBodyItemsVisitorInfo struct {
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	Pinned        *bool     `json:"pinned,omitempty" xml:"pinned,omitempty"`
	RoleCode      *string   `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	SpaceActions  []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
}

func (s RelatedSpacesResponseBodyItemsVisitorInfo) String() string {
	return tea.Prettify(s)
}

func (s RelatedSpacesResponseBodyItemsVisitorInfo) GoString() string {
	return s.String()
}

func (s *RelatedSpacesResponseBodyItemsVisitorInfo) SetDentryActions(v []*string) *RelatedSpacesResponseBodyItemsVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *RelatedSpacesResponseBodyItemsVisitorInfo) SetPinned(v bool) *RelatedSpacesResponseBodyItemsVisitorInfo {
	s.Pinned = &v
	return s
}

func (s *RelatedSpacesResponseBodyItemsVisitorInfo) SetRoleCode(v string) *RelatedSpacesResponseBodyItemsVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *RelatedSpacesResponseBodyItemsVisitorInfo) SetSpaceActions(v []*string) *RelatedSpacesResponseBodyItemsVisitorInfo {
	s.SpaceActions = v
	return s
}

type RelatedSpacesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RelatedSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RelatedSpacesResponse) String() string {
	return tea.Prettify(s)
}

func (s RelatedSpacesResponse) GoString() string {
	return s.String()
}

func (s *RelatedSpacesResponse) SetHeaders(v map[string]*string) *RelatedSpacesResponse {
	s.Headers = v
	return s
}

func (s *RelatedSpacesResponse) SetStatusCode(v int32) *RelatedSpacesResponse {
	s.StatusCode = &v
	return s
}

func (s *RelatedSpacesResponse) SetBody(v *RelatedSpacesResponseBody) *RelatedSpacesResponse {
	s.Body = v
	return s
}

type RemoveTeamMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveTeamMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveTeamMembersHeaders) GoString() string {
	return s.String()
}

func (s *RemoveTeamMembersHeaders) SetCommonHeaders(v map[string]*string) *RemoveTeamMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveTeamMembersHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveTeamMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveTeamMembersRequest struct {
	Members    []*RemoveTeamMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	Notify     *bool                              `json:"notify,omitempty" xml:"notify,omitempty"`
	OperatorId *string                            `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s RemoveTeamMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTeamMembersRequest) GoString() string {
	return s.String()
}

func (s *RemoveTeamMembersRequest) SetMembers(v []*RemoveTeamMembersRequestMembers) *RemoveTeamMembersRequest {
	s.Members = v
	return s
}

func (s *RemoveTeamMembersRequest) SetNotify(v bool) *RemoveTeamMembersRequest {
	s.Notify = &v
	return s
}

func (s *RemoveTeamMembersRequest) SetOperatorId(v string) *RemoveTeamMembersRequest {
	s.OperatorId = &v
	return s
}

type RemoveTeamMembersRequestMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *int32  `json:"memberType,omitempty" xml:"memberType,omitempty"`
	RoleCode   *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s RemoveTeamMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s RemoveTeamMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *RemoveTeamMembersRequestMembers) SetMemberId(v string) *RemoveTeamMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *RemoveTeamMembersRequestMembers) SetMemberType(v int32) *RemoveTeamMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *RemoveTeamMembersRequestMembers) SetRoleCode(v string) *RemoveTeamMembersRequestMembers {
	s.RoleCode = &v
	return s
}

type RemoveTeamMembersResponseBody struct {
	NotInOrgMembers []*RemoveTeamMembersResponseBodyNotInOrgMembers `json:"notInOrgMembers,omitempty" xml:"notInOrgMembers,omitempty" type:"Repeated"`
	SaveSuccess     *bool                                           `json:"saveSuccess,omitempty" xml:"saveSuccess,omitempty"`
}

func (s RemoveTeamMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTeamMembersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTeamMembersResponseBody) SetNotInOrgMembers(v []*RemoveTeamMembersResponseBodyNotInOrgMembers) *RemoveTeamMembersResponseBody {
	s.NotInOrgMembers = v
	return s
}

func (s *RemoveTeamMembersResponseBody) SetSaveSuccess(v bool) *RemoveTeamMembersResponseBody {
	s.SaveSuccess = &v
	return s
}

type RemoveTeamMembersResponseBodyNotInOrgMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *int32  `json:"memberType,omitempty" xml:"memberType,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	RoleCode   *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s RemoveTeamMembersResponseBodyNotInOrgMembers) String() string {
	return tea.Prettify(s)
}

func (s RemoveTeamMembersResponseBodyNotInOrgMembers) GoString() string {
	return s.String()
}

func (s *RemoveTeamMembersResponseBodyNotInOrgMembers) SetMemberId(v string) *RemoveTeamMembersResponseBodyNotInOrgMembers {
	s.MemberId = &v
	return s
}

func (s *RemoveTeamMembersResponseBodyNotInOrgMembers) SetMemberType(v int32) *RemoveTeamMembersResponseBodyNotInOrgMembers {
	s.MemberType = &v
	return s
}

func (s *RemoveTeamMembersResponseBodyNotInOrgMembers) SetName(v string) *RemoveTeamMembersResponseBodyNotInOrgMembers {
	s.Name = &v
	return s
}

func (s *RemoveTeamMembersResponseBodyNotInOrgMembers) SetRoleCode(v string) *RemoveTeamMembersResponseBodyNotInOrgMembers {
	s.RoleCode = &v
	return s
}

type RemoveTeamMembersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveTeamMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTeamMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTeamMembersResponse) GoString() string {
	return s.String()
}

func (s *RemoveTeamMembersResponse) SetHeaders(v map[string]*string) *RemoveTeamMembersResponse {
	s.Headers = v
	return s
}

func (s *RemoveTeamMembersResponse) SetStatusCode(v int32) *RemoveTeamMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTeamMembersResponse) SetBody(v *RemoveTeamMembersResponseBody) *RemoveTeamMembersResponse {
	s.Body = v
	return s
}

type RenameDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RenameDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s RenameDentryHeaders) GoString() string {
	return s.String()
}

func (s *RenameDentryHeaders) SetCommonHeaders(v map[string]*string) *RenameDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RenameDentryHeaders) SetXAcsDingtalkAccessToken(v string) *RenameDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RenameDentryRequest struct {
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s RenameDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameDentryRequest) GoString() string {
	return s.String()
}

func (s *RenameDentryRequest) SetName(v string) *RenameDentryRequest {
	s.Name = &v
	return s
}

func (s *RenameDentryRequest) SetOperatorId(v string) *RenameDentryRequest {
	s.OperatorId = &v
	return s
}

type RenameDentryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DentryVO          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameDentryResponse) GoString() string {
	return s.String()
}

func (s *RenameDentryResponse) SetHeaders(v map[string]*string) *RenameDentryResponse {
	s.Headers = v
	return s
}

func (s *RenameDentryResponse) SetStatusCode(v int32) *RenameDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameDentryResponse) SetBody(v *DentryVO) *RenameDentryResponse {
	s.Body = v
	return s
}

type SaveTeamMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveTeamMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveTeamMembersHeaders) GoString() string {
	return s.String()
}

func (s *SaveTeamMembersHeaders) SetCommonHeaders(v map[string]*string) *SaveTeamMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveTeamMembersHeaders) SetXAcsDingtalkAccessToken(v string) *SaveTeamMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveTeamMembersRequest struct {
	Members    []*SaveTeamMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	Notify     *bool                            `json:"notify,omitempty" xml:"notify,omitempty"`
	OperatorId *string                          `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s SaveTeamMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTeamMembersRequest) GoString() string {
	return s.String()
}

func (s *SaveTeamMembersRequest) SetMembers(v []*SaveTeamMembersRequestMembers) *SaveTeamMembersRequest {
	s.Members = v
	return s
}

func (s *SaveTeamMembersRequest) SetNotify(v bool) *SaveTeamMembersRequest {
	s.Notify = &v
	return s
}

func (s *SaveTeamMembersRequest) SetOperatorId(v string) *SaveTeamMembersRequest {
	s.OperatorId = &v
	return s
}

type SaveTeamMembersRequestMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *int32  `json:"memberType,omitempty" xml:"memberType,omitempty"`
	RoleCode   *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s SaveTeamMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s SaveTeamMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *SaveTeamMembersRequestMembers) SetMemberId(v string) *SaveTeamMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *SaveTeamMembersRequestMembers) SetMemberType(v int32) *SaveTeamMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *SaveTeamMembersRequestMembers) SetRoleCode(v string) *SaveTeamMembersRequestMembers {
	s.RoleCode = &v
	return s
}

type SaveTeamMembersResponseBody struct {
	NotInOrgMembers []*SaveTeamMembersResponseBodyNotInOrgMembers `json:"notInOrgMembers,omitempty" xml:"notInOrgMembers,omitempty" type:"Repeated"`
	SaveSuccess     *bool                                         `json:"saveSuccess,omitempty" xml:"saveSuccess,omitempty"`
}

func (s SaveTeamMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTeamMembersResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTeamMembersResponseBody) SetNotInOrgMembers(v []*SaveTeamMembersResponseBodyNotInOrgMembers) *SaveTeamMembersResponseBody {
	s.NotInOrgMembers = v
	return s
}

func (s *SaveTeamMembersResponseBody) SetSaveSuccess(v bool) *SaveTeamMembersResponseBody {
	s.SaveSuccess = &v
	return s
}

type SaveTeamMembersResponseBodyNotInOrgMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberType *int32  `json:"memberType,omitempty" xml:"memberType,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	RoleCode   *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s SaveTeamMembersResponseBodyNotInOrgMembers) String() string {
	return tea.Prettify(s)
}

func (s SaveTeamMembersResponseBodyNotInOrgMembers) GoString() string {
	return s.String()
}

func (s *SaveTeamMembersResponseBodyNotInOrgMembers) SetMemberId(v string) *SaveTeamMembersResponseBodyNotInOrgMembers {
	s.MemberId = &v
	return s
}

func (s *SaveTeamMembersResponseBodyNotInOrgMembers) SetMemberType(v int32) *SaveTeamMembersResponseBodyNotInOrgMembers {
	s.MemberType = &v
	return s
}

func (s *SaveTeamMembersResponseBodyNotInOrgMembers) SetName(v string) *SaveTeamMembersResponseBodyNotInOrgMembers {
	s.Name = &v
	return s
}

func (s *SaveTeamMembersResponseBodyNotInOrgMembers) SetRoleCode(v string) *SaveTeamMembersResponseBodyNotInOrgMembers {
	s.RoleCode = &v
	return s
}

type SaveTeamMembersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveTeamMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTeamMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTeamMembersResponse) GoString() string {
	return s.String()
}

func (s *SaveTeamMembersResponse) SetHeaders(v map[string]*string) *SaveTeamMembersResponse {
	s.Headers = v
	return s
}

func (s *SaveTeamMembersResponse) SetStatusCode(v int32) *SaveTeamMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTeamMembersResponse) SetBody(v *SaveTeamMembersResponseBody) *SaveTeamMembersResponse {
	s.Body = v
	return s
}

type SearchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchHeaders) GoString() string {
	return s.String()
}

func (s *SearchHeaders) SetCommonHeaders(v map[string]*string) *SearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchHeaders) SetXAcsDingtalkAccessToken(v string) *SearchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchRequest struct {
	DentryRequest *SearchRequestDentryRequest `json:"dentryRequest,omitempty" xml:"dentryRequest,omitempty" type:"Struct"`
	Keyword       *string                     `json:"keyword,omitempty" xml:"keyword,omitempty"`
	OperatorId    *string                     `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	SpaceRequest  *SearchRequestSpaceRequest  `json:"spaceRequest,omitempty" xml:"spaceRequest,omitempty" type:"Struct"`
}

func (s SearchRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRequest) GoString() string {
	return s.String()
}

func (s *SearchRequest) SetDentryRequest(v *SearchRequestDentryRequest) *SearchRequest {
	s.DentryRequest = v
	return s
}

func (s *SearchRequest) SetKeyword(v string) *SearchRequest {
	s.Keyword = &v
	return s
}

func (s *SearchRequest) SetOperatorId(v string) *SearchRequest {
	s.OperatorId = &v
	return s
}

func (s *SearchRequest) SetSpaceRequest(v *SearchRequestSpaceRequest) *SearchRequest {
	s.SpaceRequest = v
	return s
}

type SearchRequestDentryRequest struct {
	MaxResults     *int32                                    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *string                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SearchField    *int32                                    `json:"searchField,omitempty" xml:"searchField,omitempty"`
	SearchFileType *int32                                    `json:"searchFileType,omitempty" xml:"searchFileType,omitempty"`
	SpaceId        *string                                   `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	SummaryLength  *int32                                    `json:"summaryLength,omitempty" xml:"summaryLength,omitempty"`
	VisitTimeRange *SearchRequestDentryRequestVisitTimeRange `json:"visitTimeRange,omitempty" xml:"visitTimeRange,omitempty" type:"Struct"`
}

func (s SearchRequestDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRequestDentryRequest) GoString() string {
	return s.String()
}

func (s *SearchRequestDentryRequest) SetMaxResults(v int32) *SearchRequestDentryRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchRequestDentryRequest) SetNextToken(v string) *SearchRequestDentryRequest {
	s.NextToken = &v
	return s
}

func (s *SearchRequestDentryRequest) SetSearchField(v int32) *SearchRequestDentryRequest {
	s.SearchField = &v
	return s
}

func (s *SearchRequestDentryRequest) SetSearchFileType(v int32) *SearchRequestDentryRequest {
	s.SearchFileType = &v
	return s
}

func (s *SearchRequestDentryRequest) SetSpaceId(v string) *SearchRequestDentryRequest {
	s.SpaceId = &v
	return s
}

func (s *SearchRequestDentryRequest) SetSummaryLength(v int32) *SearchRequestDentryRequest {
	s.SummaryLength = &v
	return s
}

func (s *SearchRequestDentryRequest) SetVisitTimeRange(v *SearchRequestDentryRequestVisitTimeRange) *SearchRequestDentryRequest {
	s.VisitTimeRange = v
	return s
}

type SearchRequestDentryRequestVisitTimeRange struct {
	End   *int64 `json:"end,omitempty" xml:"end,omitempty"`
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s SearchRequestDentryRequestVisitTimeRange) String() string {
	return tea.Prettify(s)
}

func (s SearchRequestDentryRequestVisitTimeRange) GoString() string {
	return s.String()
}

func (s *SearchRequestDentryRequestVisitTimeRange) SetEnd(v int64) *SearchRequestDentryRequestVisitTimeRange {
	s.End = &v
	return s
}

func (s *SearchRequestDentryRequestVisitTimeRange) SetStart(v int64) *SearchRequestDentryRequestVisitTimeRange {
	s.Start = &v
	return s
}

type SearchRequestSpaceRequest struct {
	MaxResults   *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken    *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	WithTeamInfo *bool   `json:"withTeamInfo,omitempty" xml:"withTeamInfo,omitempty"`
}

func (s SearchRequestSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRequestSpaceRequest) GoString() string {
	return s.String()
}

func (s *SearchRequestSpaceRequest) SetMaxResults(v int32) *SearchRequestSpaceRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchRequestSpaceRequest) SetNextToken(v string) *SearchRequestSpaceRequest {
	s.NextToken = &v
	return s
}

func (s *SearchRequestSpaceRequest) SetWithTeamInfo(v bool) *SearchRequestSpaceRequest {
	s.WithTeamInfo = &v
	return s
}

type SearchResponseBody struct {
	DentryResult *SearchResponseBodyDentryResult `json:"dentryResult,omitempty" xml:"dentryResult,omitempty" type:"Struct"`
	SpaceResult  *SearchResponseBodySpaceResult  `json:"spaceResult,omitempty" xml:"spaceResult,omitempty" type:"Struct"`
}

func (s SearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBody) GoString() string {
	return s.String()
}

func (s *SearchResponseBody) SetDentryResult(v *SearchResponseBodyDentryResult) *SearchResponseBody {
	s.DentryResult = v
	return s
}

func (s *SearchResponseBody) SetSpaceResult(v *SearchResponseBodySpaceResult) *SearchResponseBody {
	s.SpaceResult = v
	return s
}

type SearchResponseBodyDentryResult struct {
	HasMore   *bool                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Items     []*SearchResponseBodyDentryResultItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextToken *string                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s SearchResponseBodyDentryResult) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyDentryResult) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDentryResult) SetHasMore(v bool) *SearchResponseBodyDentryResult {
	s.HasMore = &v
	return s
}

func (s *SearchResponseBodyDentryResult) SetItems(v []*SearchResponseBodyDentryResultItems) *SearchResponseBodyDentryResult {
	s.Items = v
	return s
}

func (s *SearchResponseBodyDentryResult) SetNextToken(v string) *SearchResponseBodyDentryResult {
	s.NextToken = &v
	return s
}

type SearchResponseBodyDentryResultItems struct {
	Content        *string          `json:"content,omitempty" xml:"content,omitempty"`
	Creation       *OpenActionModel `json:"creation,omitempty" xml:"creation,omitempty"`
	DentryId       *string          `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	DentryUuid     *string          `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	Extension      *string          `json:"extension,omitempty" xml:"extension,omitempty"`
	IconUrl        *string          `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	LastEdition    *OpenActionModel `json:"lastEdition,omitempty" xml:"lastEdition,omitempty"`
	Name           *string          `json:"name,omitempty" xml:"name,omitempty"`
	OriginName     *string          `json:"originName,omitempty" xml:"originName,omitempty"`
	Path           *string          `json:"path,omitempty" xml:"path,omitempty"`
	SceneType      *int32           `json:"sceneType,omitempty" xml:"sceneType,omitempty"`
	SearchFileType *int32           `json:"searchFileType,omitempty" xml:"searchFileType,omitempty"`
	SpaceId        *string          `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	ThumbnailUrl   *string          `json:"thumbnailUrl,omitempty" xml:"thumbnailUrl,omitempty"`
	Url            *string          `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SearchResponseBodyDentryResultItems) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyDentryResultItems) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDentryResultItems) SetContent(v string) *SearchResponseBodyDentryResultItems {
	s.Content = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetCreation(v *OpenActionModel) *SearchResponseBodyDentryResultItems {
	s.Creation = v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetDentryId(v string) *SearchResponseBodyDentryResultItems {
	s.DentryId = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetDentryUuid(v string) *SearchResponseBodyDentryResultItems {
	s.DentryUuid = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetExtension(v string) *SearchResponseBodyDentryResultItems {
	s.Extension = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetIconUrl(v string) *SearchResponseBodyDentryResultItems {
	s.IconUrl = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetLastEdition(v *OpenActionModel) *SearchResponseBodyDentryResultItems {
	s.LastEdition = v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetName(v string) *SearchResponseBodyDentryResultItems {
	s.Name = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetOriginName(v string) *SearchResponseBodyDentryResultItems {
	s.OriginName = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetPath(v string) *SearchResponseBodyDentryResultItems {
	s.Path = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetSceneType(v int32) *SearchResponseBodyDentryResultItems {
	s.SceneType = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetSearchFileType(v int32) *SearchResponseBodyDentryResultItems {
	s.SearchFileType = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetSpaceId(v string) *SearchResponseBodyDentryResultItems {
	s.SpaceId = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetThumbnailUrl(v string) *SearchResponseBodyDentryResultItems {
	s.ThumbnailUrl = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetUrl(v string) *SearchResponseBodyDentryResultItems {
	s.Url = &v
	return s
}

type SearchResponseBodySpaceResult struct {
	HasMore   *bool                                 `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Items     []*SearchResponseBodySpaceResultItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextToken *string                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s SearchResponseBodySpaceResult) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodySpaceResult) GoString() string {
	return s.String()
}

func (s *SearchResponseBodySpaceResult) SetHasMore(v bool) *SearchResponseBodySpaceResult {
	s.HasMore = &v
	return s
}

func (s *SearchResponseBodySpaceResult) SetItems(v []*SearchResponseBodySpaceResultItems) *SearchResponseBodySpaceResult {
	s.Items = v
	return s
}

func (s *SearchResponseBodySpaceResult) SetNextToken(v string) *SearchResponseBodySpaceResult {
	s.NextToken = &v
	return s
}

type SearchResponseBodySpaceResultItems struct {
	IconVO            *SearchResponseBodySpaceResultItemsIconVO            `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	Name              *string                                              `json:"name,omitempty" xml:"name,omitempty"`
	OriginName        *string                                              `json:"originName,omitempty" xml:"originName,omitempty"`
	SpaceId           *string                                              `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	TeamVO            *SearchResponseBodySpaceResultItemsTeamVO            `json:"teamVO,omitempty" xml:"teamVO,omitempty" type:"Struct"`
	Url               *string                                              `json:"url,omitempty" xml:"url,omitempty"`
	UserLastOperation *SearchResponseBodySpaceResultItemsUserLastOperation `json:"userLastOperation,omitempty" xml:"userLastOperation,omitempty" type:"Struct"`
}

func (s SearchResponseBodySpaceResultItems) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodySpaceResultItems) GoString() string {
	return s.String()
}

func (s *SearchResponseBodySpaceResultItems) SetIconVO(v *SearchResponseBodySpaceResultItemsIconVO) *SearchResponseBodySpaceResultItems {
	s.IconVO = v
	return s
}

func (s *SearchResponseBodySpaceResultItems) SetName(v string) *SearchResponseBodySpaceResultItems {
	s.Name = &v
	return s
}

func (s *SearchResponseBodySpaceResultItems) SetOriginName(v string) *SearchResponseBodySpaceResultItems {
	s.OriginName = &v
	return s
}

func (s *SearchResponseBodySpaceResultItems) SetSpaceId(v string) *SearchResponseBodySpaceResultItems {
	s.SpaceId = &v
	return s
}

func (s *SearchResponseBodySpaceResultItems) SetTeamVO(v *SearchResponseBodySpaceResultItemsTeamVO) *SearchResponseBodySpaceResultItems {
	s.TeamVO = v
	return s
}

func (s *SearchResponseBodySpaceResultItems) SetUrl(v string) *SearchResponseBodySpaceResultItems {
	s.Url = &v
	return s
}

func (s *SearchResponseBodySpaceResultItems) SetUserLastOperation(v *SearchResponseBodySpaceResultItemsUserLastOperation) *SearchResponseBodySpaceResultItems {
	s.UserLastOperation = v
	return s
}

type SearchResponseBodySpaceResultItemsIconVO struct {
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchResponseBodySpaceResultItemsIconVO) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodySpaceResultItemsIconVO) GoString() string {
	return s.String()
}

func (s *SearchResponseBodySpaceResultItemsIconVO) SetIcon(v string) *SearchResponseBodySpaceResultItemsIconVO {
	s.Icon = &v
	return s
}

func (s *SearchResponseBodySpaceResultItemsIconVO) SetType(v string) *SearchResponseBodySpaceResultItemsIconVO {
	s.Type = &v
	return s
}

type SearchResponseBodySpaceResultItemsTeamVO struct {
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s SearchResponseBodySpaceResultItemsTeamVO) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodySpaceResultItemsTeamVO) GoString() string {
	return s.String()
}

func (s *SearchResponseBodySpaceResultItemsTeamVO) SetId(v string) *SearchResponseBodySpaceResultItemsTeamVO {
	s.Id = &v
	return s
}

func (s *SearchResponseBodySpaceResultItemsTeamVO) SetName(v string) *SearchResponseBodySpaceResultItemsTeamVO {
	s.Name = &v
	return s
}

type SearchResponseBodySpaceResultItemsUserLastOperation struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Time *int64  `json:"time,omitempty" xml:"time,omitempty"`
}

func (s SearchResponseBodySpaceResultItemsUserLastOperation) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodySpaceResultItemsUserLastOperation) GoString() string {
	return s.String()
}

func (s *SearchResponseBodySpaceResultItemsUserLastOperation) SetName(v string) *SearchResponseBodySpaceResultItemsUserLastOperation {
	s.Name = &v
	return s
}

func (s *SearchResponseBodySpaceResultItemsUserLastOperation) SetTime(v int64) *SearchResponseBodySpaceResultItemsUserLastOperation {
	s.Time = &v
	return s
}

type SearchResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchResponse) GoString() string {
	return s.String()
}

func (s *SearchResponse) SetHeaders(v map[string]*string) *SearchResponse {
	s.Headers = v
	return s
}

func (s *SearchResponse) SetStatusCode(v int32) *SearchResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchResponse) SetBody(v *SearchResponseBody) *SearchResponse {
	s.Body = v
	return s
}

type UnmarkStarHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UnmarkStarHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnmarkStarHeaders) GoString() string {
	return s.String()
}

func (s *UnmarkStarHeaders) SetCommonHeaders(v map[string]*string) *UnmarkStarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnmarkStarHeaders) SetXAcsDingtalkAccessToken(v string) *UnmarkStarHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UnmarkStarRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UnmarkStarRequest) String() string {
	return tea.Prettify(s)
}

func (s UnmarkStarRequest) GoString() string {
	return s.String()
}

func (s *UnmarkStarRequest) SetOperatorId(v string) *UnmarkStarRequest {
	s.OperatorId = &v
	return s
}

type UnmarkStarResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UnmarkStarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnmarkStarResponseBody) GoString() string {
	return s.String()
}

func (s *UnmarkStarResponseBody) SetSuccess(v bool) *UnmarkStarResponseBody {
	s.Success = &v
	return s
}

type UnmarkStarResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnmarkStarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnmarkStarResponse) String() string {
	return tea.Prettify(s)
}

func (s UnmarkStarResponse) GoString() string {
	return s.String()
}

func (s *UnmarkStarResponse) SetHeaders(v map[string]*string) *UnmarkStarResponse {
	s.Headers = v
	return s
}

func (s *UnmarkStarResponse) SetStatusCode(v int32) *UnmarkStarResponse {
	s.StatusCode = &v
	return s
}

func (s *UnmarkStarResponse) SetBody(v *UnmarkStarResponseBody) *UnmarkStarResponse {
	s.Body = v
	return s
}

type UnpinSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UnpinSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnpinSpaceHeaders) GoString() string {
	return s.String()
}

func (s *UnpinSpaceHeaders) SetCommonHeaders(v map[string]*string) *UnpinSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnpinSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *UnpinSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UnpinSpaceRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UnpinSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnpinSpaceRequest) GoString() string {
	return s.String()
}

func (s *UnpinSpaceRequest) SetOperatorId(v string) *UnpinSpaceRequest {
	s.OperatorId = &v
	return s
}

type UnpinSpaceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UnpinSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnpinSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *UnpinSpaceResponseBody) SetSuccess(v bool) *UnpinSpaceResponseBody {
	s.Success = &v
	return s
}

type UnpinSpaceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnpinSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnpinSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnpinSpaceResponse) GoString() string {
	return s.String()
}

func (s *UnpinSpaceResponse) SetHeaders(v map[string]*string) *UnpinSpaceResponse {
	s.Headers = v
	return s
}

func (s *UnpinSpaceResponse) SetStatusCode(v int32) *UnpinSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnpinSpaceResponse) SetBody(v *UnpinSpaceResponseBody) *UnpinSpaceResponse {
	s.Body = v
	return s
}

type UpdateTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTeamHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTeamHeaders) SetCommonHeaders(v map[string]*string) *UpdateTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTeamHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTeamRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OperatorId  *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTeamRequest) GoString() string {
	return s.String()
}

func (s *UpdateTeamRequest) SetDescription(v string) *UpdateTeamRequest {
	s.Description = &v
	return s
}

func (s *UpdateTeamRequest) SetName(v string) *UpdateTeamRequest {
	s.Name = &v
	return s
}

func (s *UpdateTeamRequest) SetOperatorId(v string) *UpdateTeamRequest {
	s.OperatorId = &v
	return s
}

type UpdateTeamResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TeamVO            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTeamResponse) GoString() string {
	return s.String()
}

func (s *UpdateTeamResponse) SetHeaders(v map[string]*string) *UpdateTeamResponse {
	s.Headers = v
	return s
}

func (s *UpdateTeamResponse) SetStatusCode(v int32) *UpdateTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTeamResponse) SetBody(v *TeamVO) *UpdateTeamResponse {
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

func (client *Client) BatchDeleteRecentsWithOptions(request *BatchDeleteRecentsRequest, headers *BatchDeleteRecentsHeaders, runtime *util.RuntimeOptions) (_result *BatchDeleteRecentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryUuids)) {
		body["dentryUuids"] = request.DentryUuids
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
		Action:      tea.String("BatchDeleteRecents"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/recentRecords/batchRemove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteRecentsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteRecents(request *BatchDeleteRecentsRequest) (_result *BatchDeleteRecentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchDeleteRecentsHeaders{}
	_result = &BatchDeleteRecentsResponse{}
	_body, _err := client.BatchDeleteRecentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyDentryWithOptions(spaceId *string, dentryId *string, request *CopyDentryRequest, headers *CopyDentryHeaders, runtime *util.RuntimeOptions) (_result *CopyDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSpaceId)) {
		body["targetSpaceId"] = request.TargetSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.ToNextDentryId)) {
		body["toNextDentryId"] = request.ToNextDentryId
	}

	if !tea.BoolValue(util.IsUnset(request.ToParentDentryId)) {
		body["toParentDentryId"] = request.ToParentDentryId
	}

	if !tea.BoolValue(util.IsUnset(request.ToPrevDentryId)) {
		body["toPrevDentryId"] = request.ToPrevDentryId
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
		Action:      tea.String("CopyDentry"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/copy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyDentry(spaceId *string, dentryId *string, request *CopyDentryRequest) (_result *CopyDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CopyDentryHeaders{}
	_result = &CopyDentryResponse{}
	_body, _err := client.CopyDentryWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDentryWithOptions(spaceId *string, request *CreateDentryRequest, headers *CreateDentryHeaders, runtime *util.RuntimeOptions) (_result *CreateDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryType)) {
		body["dentryType"] = request.DentryType
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentType)) {
		body["documentType"] = request.DocumentType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDentryId)) {
		body["parentDentryId"] = request.ParentDentryId
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
		Action:      tea.String("CreateDentry"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/" + tea.StringValue(spaceId) + "/dentries"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDentry(spaceId *string, request *CreateDentryRequest) (_result *CreateDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDentryHeaders{}
	_result = &CreateDentryResponse{}
	_body, _err := client.CreateDentryWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSpaceWithOptions(request *CreateSpaceRequest, headers *CreateSpaceHeaders, runtime *util.RuntimeOptions) (_result *CreateSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SectionId)) {
		body["sectionId"] = request.SectionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareScope)) {
		body["shareScope"] = request.ShareScope
	}

	if !tea.BoolValue(util.IsUnset(request.TeamId)) {
		body["teamId"] = request.TeamId
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
		Action:      tea.String("CreateSpace"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSpace(request *CreateSpaceRequest) (_result *CreateSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSpaceHeaders{}
	_result = &CreateSpaceResponse{}
	_body, _err := client.CreateSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTeamWithOptions(request *CreateTeamRequest, headers *CreateTeamHeaders, runtime *util.RuntimeOptions) (_result *CreateTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cover)) {
		body["cover"] = request.Cover
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TeamType)) {
		body["teamType"] = request.TeamType
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
		Action:      tea.String("CreateTeam"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTeam(request *CreateTeamRequest) (_result *CreateTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTeamHeaders{}
	_result = &CreateTeamResponse{}
	_body, _err := client.CreateTeamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTeamWithOptions(teamId *string, request *DeleteTeamRequest, headers *DeleteTeamHeaders, runtime *util.RuntimeOptions) (_result *DeleteTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("DeleteTeam"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/" + tea.StringValue(teamId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTeam(teamId *string, request *DeleteTeamRequest) (_result *DeleteTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTeamHeaders{}
	_result = &DeleteTeamResponse{}
	_body, _err := client.DeleteTeamWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSchemaWithOptions(teamId *string, request *GetSchemaRequest, headers *GetSchemaHeaders, runtime *util.RuntimeOptions) (_result *GetSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("GetSchema"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/" + tea.StringValue(teamId) + "/schemas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSchemaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSchema(teamId *string, request *GetSchemaRequest) (_result *GetSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSchemaHeaders{}
	_result = &GetSchemaResponse{}
	_body, _err := client.GetSchemaWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpaceDirectoriesWithOptions(spaceId *string, request *GetSpaceDirectoriesRequest, headers *GetSpaceDirectoriesHeaders, runtime *util.RuntimeOptions) (_result *GetSpaceDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryId)) {
		query["dentryId"] = request.DentryId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("GetSpaceDirectories"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/" + tea.StringValue(spaceId) + "/directories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpaceDirectoriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpaceDirectories(spaceId *string, request *GetSpaceDirectoriesRequest) (_result *GetSpaceDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSpaceDirectoriesHeaders{}
	_result = &GetSpaceDirectoriesResponse{}
	_body, _err := client.GetSpaceDirectoriesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStarInfoWithOptions(dentryUuid *string, request *GetStarInfoRequest, headers *GetStarInfoHeaders, runtime *util.RuntimeOptions) (_result *GetStarInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("GetStarInfo"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/" + tea.StringValue(dentryUuid) + "/starInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStarInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStarInfo(dentryUuid *string, request *GetStarInfoRequest) (_result *GetStarInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetStarInfoHeaders{}
	_result = &GetStarInfoResponse{}
	_body, _err := client.GetStarInfoWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTeamWithOptions(teamId *string, request *GetTeamRequest, headers *GetTeamHeaders, runtime *util.RuntimeOptions) (_result *GetTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("GetTeam"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/" + tea.StringValue(teamId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTeam(teamId *string, request *GetTeamRequest) (_result *GetTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTeamHeaders{}
	_result = &GetTeamResponse{}
	_body, _err := client.GetTeamWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTotalNumberOfDentriesWithOptions(request *GetTotalNumberOfDentriesRequest, headers *GetTotalNumberOfDentriesHeaders, runtime *util.RuntimeOptions) (_result *GetTotalNumberOfDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeFolder)) {
		query["includeFolder"] = request.IncludeFolder
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceTypes)) {
		query["spaceTypes"] = request.SpaceTypes
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
		Action:      tea.String("GetTotalNumberOfDentries"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/statistics/dentryCounts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTotalNumberOfDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTotalNumberOfDentries(request *GetTotalNumberOfDentriesRequest) (_result *GetTotalNumberOfDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTotalNumberOfDentriesHeaders{}
	_result = &GetTotalNumberOfDentriesResponse{}
	_body, _err := client.GetTotalNumberOfDentriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTotalNumberOfSpacesWithOptions(request *GetTotalNumberOfSpacesRequest, headers *GetTotalNumberOfSpacesHeaders, runtime *util.RuntimeOptions) (_result *GetTotalNumberOfSpacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("GetTotalNumberOfSpaces"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/statistics/spaceCounts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTotalNumberOfSpacesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTotalNumberOfSpaces(request *GetTotalNumberOfSpacesRequest) (_result *GetTotalNumberOfSpacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTotalNumberOfSpacesHeaders{}
	_result = &GetTotalNumberOfSpacesResponse{}
	_body, _err := client.GetTotalNumberOfSpacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserInfoByOpenTokenWithOptions(request *GetUserInfoByOpenTokenRequest, headers *GetUserInfoByOpenTokenHeaders, runtime *util.RuntimeOptions) (_result *GetUserInfoByOpenTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocKey)) {
		query["docKey"] = request.DocKey
	}

	if !tea.BoolValue(util.IsUnset(request.OpenToken)) {
		query["openToken"] = request.OpenToken
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
		Action:      tea.String("GetUserInfoByOpenToken"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/userInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserInfoByOpenTokenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserInfoByOpenToken(request *GetUserInfoByOpenTokenRequest) (_result *GetUserInfoByOpenTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserInfoByOpenTokenHeaders{}
	_result = &GetUserInfoByOpenTokenResponse{}
	_body, _err := client.GetUserInfoByOpenTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFeedsWithOptions(teamId *string, request *ListFeedsRequest, headers *ListFeedsHeaders, runtime *util.RuntimeOptions) (_result *ListFeedsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeFile)) {
		query["excludeFile"] = request.ExcludeFile
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("ListFeeds"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/" + tea.StringValue(teamId) + "/feeds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFeedsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFeeds(teamId *string, request *ListFeedsRequest) (_result *ListFeedsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListFeedsHeaders{}
	_result = &ListFeedsResponse{}
	_body, _err := client.ListFeedsWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotDocsWithOptions(teamId *string, request *ListHotDocsRequest, headers *ListHotDocsHeaders, runtime *util.RuntimeOptions) (_result *ListHotDocsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("ListHotDocs"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/" + tea.StringValue(teamId) + "/hotDocs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotDocsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotDocs(teamId *string, request *ListHotDocsRequest) (_result *ListHotDocsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListHotDocsHeaders{}
	_result = &ListHotDocsResponse{}
	_body, _err := client.ListHotDocsWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPinSpacesWithOptions(request *ListPinSpacesRequest, headers *ListPinSpacesHeaders, runtime *util.RuntimeOptions) (_result *ListPinSpacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("ListPinSpaces"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/pinLists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPinSpacesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPinSpaces(request *ListPinSpacesRequest) (_result *ListPinSpacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListPinSpacesHeaders{}
	_result = &ListPinSpacesResponse{}
	_body, _err := client.ListPinSpacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRelatedSpaceTeamsWithOptions(request *ListRelatedSpaceTeamsRequest, headers *ListRelatedSpaceTeamsHeaders, runtime *util.RuntimeOptions) (_result *ListRelatedSpaceTeamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("ListRelatedSpaceTeams"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/relations/spaceTeams"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRelatedSpaceTeamsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRelatedSpaceTeams(request *ListRelatedSpaceTeamsRequest) (_result *ListRelatedSpaceTeamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRelatedSpaceTeamsHeaders{}
	_result = &ListRelatedSpaceTeamsResponse{}
	_body, _err := client.ListRelatedSpaceTeamsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRelatedTeamsWithOptions(request *ListRelatedTeamsRequest, headers *ListRelatedTeamsHeaders, runtime *util.RuntimeOptions) (_result *ListRelatedTeamsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("ListRelatedTeams"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRelatedTeamsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRelatedTeams(request *ListRelatedTeamsRequest) (_result *ListRelatedTeamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRelatedTeamsHeaders{}
	_result = &ListRelatedTeamsResponse{}
	_body, _err := client.ListRelatedTeamsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSpaceSectionsWithOptions(teamId *string, request *ListSpaceSectionsRequest, headers *ListSpaceSectionsHeaders, runtime *util.RuntimeOptions) (_result *ListSpaceSectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("ListSpaceSections"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/" + tea.StringValue(teamId) + "/spaceSections"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSpaceSectionsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSpaceSections(teamId *string, request *ListSpaceSectionsRequest) (_result *ListSpaceSectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSpaceSectionsHeaders{}
	_result = &ListSpaceSectionsResponse{}
	_body, _err := client.ListSpaceSectionsWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStarsWithOptions(request *ListStarsRequest, headers *ListStarsHeaders, runtime *util.RuntimeOptions) (_result *ListStarsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("ListStars"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/starLists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStarsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStars(request *ListStarsRequest) (_result *ListStarsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListStarsHeaders{}
	_result = &ListStarsResponse{}
	_body, _err := client.ListStarsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTeamMembersWithOptions(teamId *string, request *ListTeamMembersRequest, headers *ListTeamMembersHeaders, runtime *util.RuntimeOptions) (_result *ListTeamMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("ListTeamMembers"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/" + tea.StringValue(teamId) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTeamMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTeamMembers(teamId *string, request *ListTeamMembersRequest) (_result *ListTeamMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTeamMembersHeaders{}
	_result = &ListTeamMembersResponse{}
	_body, _err := client.ListTeamMembersWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MarkStarWithOptions(dentryUuid *string, request *MarkStarRequest, headers *MarkStarHeaders, runtime *util.RuntimeOptions) (_result *MarkStarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("MarkStar"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/" + tea.StringValue(dentryUuid) + "/stars/mark"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MarkStarResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MarkStar(dentryUuid *string, request *MarkStarRequest) (_result *MarkStarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MarkStarHeaders{}
	_result = &MarkStarResponse{}
	_body, _err := client.MarkStarWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveDentryWithOptions(spaceId *string, dentryId *string, request *MoveDentryRequest, headers *MoveDentryHeaders, runtime *util.RuntimeOptions) (_result *MoveDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSpaceId)) {
		body["targetSpaceId"] = request.TargetSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.ToNextDentryId)) {
		body["toNextDentryId"] = request.ToNextDentryId
	}

	if !tea.BoolValue(util.IsUnset(request.ToParentDentryId)) {
		body["toParentDentryId"] = request.ToParentDentryId
	}

	if !tea.BoolValue(util.IsUnset(request.ToPrevDentryId)) {
		body["toPrevDentryId"] = request.ToPrevDentryId
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
		Action:      tea.String("MoveDentry"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/move"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveDentry(spaceId *string, dentryId *string, request *MoveDentryRequest) (_result *MoveDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MoveDentryHeaders{}
	_result = &MoveDentryResponse{}
	_body, _err := client.MoveDentryWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PinSpaceWithOptions(spaceId *string, request *PinSpaceRequest, headers *PinSpaceHeaders, runtime *util.RuntimeOptions) (_result *PinSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("PinSpace"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/" + tea.StringValue(spaceId) + "/pin"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PinSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PinSpace(spaceId *string, request *PinSpaceRequest) (_result *PinSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PinSpaceHeaders{}
	_result = &PinSpaceResponse{}
	_body, _err := client.PinSpaceWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDentryWithOptions(spaceId *string, dentryId *string, request *QueryDentryRequest, headers *QueryDentryHeaders, runtime *util.RuntimeOptions) (_result *QueryDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeSpace)) {
		query["includeSpace"] = request.IncludeSpace
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("QueryDentry"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDentry(spaceId *string, dentryId *string, request *QueryDentryRequest) (_result *QueryDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDentryHeaders{}
	_result = &QueryDentryResponse{}
	_body, _err := client.QueryDentryWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryItemByUrlWithOptions(request *QueryItemByUrlRequest, headers *QueryItemByUrlHeaders, runtime *util.RuntimeOptions) (_result *QueryItemByUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.WithStatisticalInfo)) {
		query["withStatisticalInfo"] = request.WithStatisticalInfo
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
		Action:      tea.String("QueryItemByUrl"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/items"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryItemByUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryItemByUrl(request *QueryItemByUrlRequest) (_result *QueryItemByUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryItemByUrlHeaders{}
	_result = &QueryItemByUrlResponse{}
	_body, _err := client.QueryItemByUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMineSpaceWithOptions(unionId *string, headers *QueryMineSpaceHeaders, runtime *util.RuntimeOptions) (_result *QueryMineSpaceResponse, _err error) {
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
		Action:      tea.String("QueryMineSpace"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/users/" + tea.StringValue(unionId) + "/mine"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMineSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMineSpace(unionId *string) (_result *QueryMineSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMineSpaceHeaders{}
	_result = &QueryMineSpaceResponse{}
	_body, _err := client.QueryMineSpaceWithOptions(unionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecentListWithOptions(request *QueryRecentListRequest, headers *QueryRecentListHeaders, runtime *util.RuntimeOptions) (_result *QueryRecentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorType)) {
		query["creatorType"] = request.CreatorType
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["fileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.RecentType)) {
		query["recentType"] = request.RecentType
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
		Action:      tea.String("QueryRecentList"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/docs/recent"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecentListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecentList(request *QueryRecentListRequest) (_result *QueryRecentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRecentListHeaders{}
	_result = &QueryRecentListResponse{}
	_body, _err := client.QueryRecentListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySpaceWithOptions(spaceId *string, request *QuerySpaceRequest, headers *QuerySpaceHeaders, runtime *util.RuntimeOptions) (_result *QuerySpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("QuerySpace"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/" + tea.StringValue(spaceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySpace(spaceId *string, request *QuerySpaceRequest) (_result *QuerySpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySpaceHeaders{}
	_result = &QuerySpaceResponse{}
	_body, _err := client.QuerySpaceWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RelatedSpacesWithOptions(request *RelatedSpacesRequest, headers *RelatedSpacesHeaders, runtime *util.RuntimeOptions) (_result *RelatedSpacesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TeamId)) {
		query["teamId"] = request.TeamId
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
		Action:      tea.String("RelatedSpaces"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/relations/spaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RelatedSpacesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RelatedSpaces(request *RelatedSpacesRequest) (_result *RelatedSpacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RelatedSpacesHeaders{}
	_result = &RelatedSpacesResponse{}
	_body, _err := client.RelatedSpacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveTeamMembersWithOptions(teamId *string, request *RemoveTeamMembersRequest, headers *RemoveTeamMembersHeaders, runtime *util.RuntimeOptions) (_result *RemoveTeamMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.Notify)) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("RemoveTeamMembers"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/" + tea.StringValue(teamId) + "/members/remove"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveTeamMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveTeamMembers(teamId *string, request *RemoveTeamMembersRequest) (_result *RemoveTeamMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveTeamMembersHeaders{}
	_result = &RemoveTeamMembersResponse{}
	_body, _err := client.RemoveTeamMembersWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameDentryWithOptions(spaceId *string, dentryId *string, request *RenameDentryRequest, headers *RenameDentryHeaders, runtime *util.RuntimeOptions) (_result *RenameDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("RenameDentry"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/rename"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameDentry(spaceId *string, dentryId *string, request *RenameDentryRequest) (_result *RenameDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RenameDentryHeaders{}
	_result = &RenameDentryResponse{}
	_body, _err := client.RenameDentryWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTeamMembersWithOptions(teamId *string, request *SaveTeamMembersRequest, headers *SaveTeamMembersHeaders, runtime *util.RuntimeOptions) (_result *SaveTeamMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.Notify)) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("SaveTeamMembers"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/" + tea.StringValue(teamId) + "/members"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveTeamMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTeamMembers(teamId *string, request *SaveTeamMembersRequest) (_result *SaveTeamMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveTeamMembersHeaders{}
	_result = &SaveTeamMembersResponse{}
	_body, _err := client.SaveTeamMembersWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchWithOptions(request *SearchRequest, headers *SearchHeaders, runtime *util.RuntimeOptions) (_result *SearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryRequest)) {
		body["dentryRequest"] = request.DentryRequest
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceRequest)) {
		body["spaceRequest"] = request.SpaceRequest
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
		Action:      tea.String("Search"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Search(request *SearchRequest) (_result *SearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchHeaders{}
	_result = &SearchResponse{}
	_body, _err := client.SearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnmarkStarWithOptions(dentryUuid *string, request *UnmarkStarRequest, headers *UnmarkStarHeaders, runtime *util.RuntimeOptions) (_result *UnmarkStarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("UnmarkStar"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/" + tea.StringValue(dentryUuid) + "/stars/unmark"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UnmarkStarResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnmarkStar(dentryUuid *string, request *UnmarkStarRequest) (_result *UnmarkStarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnmarkStarHeaders{}
	_result = &UnmarkStarResponse{}
	_body, _err := client.UnmarkStarWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnpinSpaceWithOptions(spaceId *string, request *UnpinSpaceRequest, headers *UnpinSpaceHeaders, runtime *util.RuntimeOptions) (_result *UnpinSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("UnpinSpace"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/spaces/" + tea.StringValue(spaceId) + "/unpin"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UnpinSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnpinSpace(spaceId *string, request *UnpinSpaceRequest) (_result *UnpinSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnpinSpaceHeaders{}
	_result = &UnpinSpaceResponse{}
	_body, _err := client.UnpinSpaceWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTeamWithOptions(teamId *string, request *UpdateTeamRequest, headers *UpdateTeamHeaders, runtime *util.RuntimeOptions) (_result *UpdateTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("UpdateTeam"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/" + tea.StringValue(teamId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTeam(teamId *string, request *UpdateTeamRequest) (_result *UpdateTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTeamHeaders{}
	_result = &UpdateTeamResponse{}
	_body, _err := client.UpdateTeamWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
