// This file is auto-generated, don't edit it. Thanks.
package doc_2_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DentryModel struct {
	// example:
	//
	// alidoc
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1663918630284
	CreatedTime *int64              `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Creator     *DentryModelCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// YRBd*****KGDA
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// file
	DentryType *string `json:"dentryType,omitempty" xml:"dentryType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6or0dp8Z****XWa91xzy3
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// example:
	//
	// v1GXn****KqD4
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	// example:
	//
	// adoc
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	HasChildren    *bool           `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	LinkSourceInfo *LinkSourceInfo `json:"linkSourceInfo,omitempty" xml:"linkSourceInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉钉文档
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 测试组织/测试知识库/abc
	Path  *string     `json:"path,omitempty" xml:"path,omitempty"`
	Space *SpaceModel `json:"space,omitempty" xml:"space,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YGv0****0xXAr
	SpaceId         *string                     `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	StatisticalInfo *DentryModelStatisticalInfo `json:"statisticalInfo,omitempty" xml:"statisticalInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1663918630284
	UpdatedTime *int64              `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	Updater     *DentryModelUpdater `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	// example:
	//
	// https://xxx.yy
	Url         *string                 `json:"url,omitempty" xml:"url,omitempty"`
	VisitorInfo *DentryModelVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
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
	// This parameter is required.
	//
	// example:
	//
	// DingTalk
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
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
	// This parameter is required.
	//
	// example:
	//
	// DingTalk
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
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
	// example:
	//
	// 5
	RoleCode     *string   `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	SpaceActions []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
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
	// example:
	//
	// alidoc
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1663918630284
	CreatedTime *int64           `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Creator     *DentryVOCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// YRBd*****KGDA
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// file
	DentryType *string `json:"dentryType,omitempty" xml:"dentryType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6or0dp8Z****XWa91xzy3
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// example:
	//
	// v1GXn****KqD4
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	// example:
	//
	// alidoc
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	HasChildren    *bool           `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	LinkSourceInfo *LinkSourceInfo `json:"linkSourceInfo,omitempty" xml:"linkSourceInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉钉文档
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 测试组织/测试知识库/abc
	Path  *string     `json:"path,omitempty" xml:"path,omitempty"`
	Space *SpaceModel `json:"space,omitempty" xml:"space,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YGv0****0xXAr
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1663918630284
	UpdatedTime *int64           `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	Updater     *DentryVOUpdater `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	// example:
	//
	// https://xxx.yy
	Url         *string              `json:"url,omitempty" xml:"url,omitempty"`
	VisitorInfo *DentryVOVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
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
	// This parameter is required.
	//
	// example:
	//
	// DingTalk
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
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
	// This parameter is required.
	//
	// example:
	//
	// DingTalk
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
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
	// example:
	//
	// 5
	RoleCode     *string   `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	SpaceActions []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
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
	// example:
	//
	// docx
	Extension *string                `json:"extension,omitempty" xml:"extension,omitempty"`
	IconUrl   *LinkSourceInfoIconUrl `json:"iconUrl,omitempty" xml:"iconUrl,omitempty" type:"Struct"`
	// example:
	//
	// abc
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	LinkType *int64 `json:"linkType,omitempty" xml:"linkType,omitempty"`
	// example:
	//
	// def
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
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
	// example:
	//
	// gh
	Line *string `json:"line,omitempty" xml:"line,omitempty"`
	// example:
	//
	// def
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
	// example:
	//
	// sky
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 273829092
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
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
	// example:
	//
	// https://img.abc.yyy
	Cover *string `json:"cover,omitempty" xml:"cover,omitempty"`
	// example:
	//
	// This is some description.
	Description *string             `json:"description,omitempty" xml:"description,omitempty"`
	HdIconVO    *SpaceModelHdIconVO `json:"hdIconVO,omitempty" xml:"hdIconVO,omitempty" type:"Struct"`
	IconVO      *SpaceModelIconVO   `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hello
	Name       *string          `json:"name,omitempty" xml:"name,omitempty"`
	Owner      *SpaceModelOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	RecentList []*DentryModel   `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// https://xx.yy
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
	// This parameter is required.
	//
	// example:
	//
	// https://img.xxx.yyy
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// url
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
	// This parameter is required.
	//
	// example:
	//
	// https://img.xxx.yyy
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// url
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
	// This parameter is required.
	//
	// example:
	//
	// dingtalk
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcd
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
	// example:
	//
	// 3
	RoleCode     *string   `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	SpaceActions []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
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
	Cover       *string        `json:"cover,omitempty" xml:"cover,omitempty"`
	Description *string        `json:"description,omitempty" xml:"description,omitempty"`
	IconVO      *SpaceVOIconVO `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hello
	Name  *string       `json:"name,omitempty" xml:"name,omitempty"`
	Owner *SpaceVOOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// https://xx.yy
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
	// This parameter is required.
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// dingtalk
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcd
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
	// example:
	//
	// https://abc.com
	Cover *string `json:"cover,omitempty" xml:"cover,omitempty"`
	// example:
	//
	// 12340000
	CreatedTime *int64            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Creator     *TeamModelCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// example:
	//
	// 这里是团队描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// https://def.com
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AbcDef
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试团队名称
	Name            *string                   `json:"name,omitempty" xml:"name,omitempty"`
	RelatedDeptInfo *TeamModelRelatedDeptInfo `json:"relatedDeptInfo,omitempty" xml:"relatedDeptInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 34560000
	UpdatedTime *int64            `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	Updater     *TeamModelUpdater `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// https://abc.com
	Url       *string             `json:"url,omitempty" xml:"url,omitempty"`
	VisitInfo *TeamModelVisitInfo `json:"visitInfo,omitempty" xml:"visitInfo,omitempty" type:"Struct"`
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
	// example:
	//
	// 测试
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// abcd
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
	// example:
	//
	// abc
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// 测试部门
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
	// example:
	//
	// 测试
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// abcd
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
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	JoinTime *string `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	// example:
	//
	// 5
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
	// example:
	//
	// https://abc.com
	Cover *string `json:"cover,omitempty" xml:"cover,omitempty"`
	// example:
	//
	// 12340000
	CreatedTime *int64         `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Creator     *TeamVOCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// example:
	//
	// 这里是团队描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// https://def.com
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AbcDef
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试团队名称
	Name            *string                `json:"name,omitempty" xml:"name,omitempty"`
	RelatedDeptInfo *TeamVORelatedDeptInfo `json:"relatedDeptInfo,omitempty" xml:"relatedDeptInfo,omitempty" type:"Struct"`
	ShareScopeInfo  *TeamVOShareScopeInfo  `json:"shareScopeInfo,omitempty" xml:"shareScopeInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 34560000
	UpdatedTime *int64         `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	Updater     *TeamVOUpdater `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// https://abc.com
	Url       *string          `json:"url,omitempty" xml:"url,omitempty"`
	VisitInfo *TeamVOVisitInfo `json:"visitInfo,omitempty" xml:"visitInfo,omitempty" type:"Struct"`
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
	// example:
	//
	// 测试
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// abc
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
	// example:
	//
	// abc
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// 测试部门
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
	// example:
	//
	// 测试
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// abc
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
	// example:
	//
	// 5
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

type MapValue struct {
	TemplateId         *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
	Type               *int32  `json:"type,omitempty" xml:"type,omitempty"`
	CoverDownloadUrl   *string `json:"coverDownloadUrl,omitempty" xml:"coverDownloadUrl,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	AuthorName         *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	CreateTime         *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ModifiedTime       *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	WorkspaceId        *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	WorkspaceName      *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	UsedCount          *int64  `json:"usedCount,omitempty" xml:"usedCount,omitempty"`
	Belong             *string `json:"belong,omitempty" xml:"belong,omitempty"`
	ContentDownloadUrl *string `json:"contentDownloadUrl,omitempty" xml:"contentDownloadUrl,omitempty"`
}

func (s MapValue) String() string {
	return tea.Prettify(s)
}

func (s MapValue) GoString() string {
	return s.String()
}

func (s *MapValue) SetTemplateId(v string) *MapValue {
	s.TemplateId = &v
	return s
}

func (s *MapValue) SetTitle(v string) *MapValue {
	s.Title = &v
	return s
}

func (s *MapValue) SetType(v int32) *MapValue {
	s.Type = &v
	return s
}

func (s *MapValue) SetCoverDownloadUrl(v string) *MapValue {
	s.CoverDownloadUrl = &v
	return s
}

func (s *MapValue) SetDescription(v string) *MapValue {
	s.Description = &v
	return s
}

func (s *MapValue) SetAuthorName(v string) *MapValue {
	s.AuthorName = &v
	return s
}

func (s *MapValue) SetCreateTime(v string) *MapValue {
	s.CreateTime = &v
	return s
}

func (s *MapValue) SetModifiedTime(v string) *MapValue {
	s.ModifiedTime = &v
	return s
}

func (s *MapValue) SetWorkspaceId(v string) *MapValue {
	s.WorkspaceId = &v
	return s
}

func (s *MapValue) SetWorkspaceName(v string) *MapValue {
	s.WorkspaceName = &v
	return s
}

func (s *MapValue) SetUsedCount(v int64) *MapValue {
	s.UsedCount = &v
	return s
}

func (s *MapValue) SetBelong(v string) *MapValue {
	s.Belong = &v
	return s
}

func (s *MapValue) SetContentDownloadUrl(v string) *MapValue {
	s.ContentDownloadUrl = &v
	return s
}

type BatchCreateTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchCreateTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTeamHeaders) GoString() string {
	return s.String()
}

func (s *BatchCreateTeamHeaders) SetCommonHeaders(v map[string]*string) *BatchCreateTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchCreateTeamHeaders) SetXAcsDingtalkAccessToken(v string) *BatchCreateTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchCreateTeamRequest struct {
	// This parameter is required.
	Param *BatchCreateTeamRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s BatchCreateTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTeamRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateTeamRequest) SetParam(v *BatchCreateTeamRequestParam) *BatchCreateTeamRequest {
	s.Param = v
	return s
}

func (s *BatchCreateTeamRequest) SetOperatorId(v string) *BatchCreateTeamRequest {
	s.OperatorId = &v
	return s
}

type BatchCreateTeamRequestParam struct {
	// This parameter is required.
	CreateTeamParamList []*BatchCreateTeamRequestParamCreateTeamParamList `json:"createTeamParamList,omitempty" xml:"createTeamParamList,omitempty" type:"Repeated"`
}

func (s BatchCreateTeamRequestParam) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTeamRequestParam) GoString() string {
	return s.String()
}

func (s *BatchCreateTeamRequestParam) SetCreateTeamParamList(v []*BatchCreateTeamRequestParamCreateTeamParamList) *BatchCreateTeamRequestParam {
	s.CreateTeamParamList = v
	return s
}

type BatchCreateTeamRequestParamCreateTeamParamList struct {
	AdminUnionIdList []*string `json:"adminUnionIdList,omitempty" xml:"adminUnionIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// creator_union_id
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dept_id
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// team_name
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
}

func (s BatchCreateTeamRequestParamCreateTeamParamList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTeamRequestParamCreateTeamParamList) GoString() string {
	return s.String()
}

func (s *BatchCreateTeamRequestParamCreateTeamParamList) SetAdminUnionIdList(v []*string) *BatchCreateTeamRequestParamCreateTeamParamList {
	s.AdminUnionIdList = v
	return s
}

func (s *BatchCreateTeamRequestParamCreateTeamParamList) SetCreatorUnionId(v string) *BatchCreateTeamRequestParamCreateTeamParamList {
	s.CreatorUnionId = &v
	return s
}

func (s *BatchCreateTeamRequestParamCreateTeamParamList) SetDeptId(v string) *BatchCreateTeamRequestParamCreateTeamParamList {
	s.DeptId = &v
	return s
}

func (s *BatchCreateTeamRequestParamCreateTeamParamList) SetTeamName(v string) *BatchCreateTeamRequestParamCreateTeamParamList {
	s.TeamName = &v
	return s
}

type BatchCreateTeamResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchCreateTeamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTeamResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateTeamResponseBody) SetSuccess(v bool) *BatchCreateTeamResponseBody {
	s.Success = &v
	return s
}

type BatchCreateTeamResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateTeamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTeamResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateTeamResponse) SetHeaders(v map[string]*string) *BatchCreateTeamResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateTeamResponse) SetStatusCode(v int32) *BatchCreateTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateTeamResponse) SetBody(v *BatchCreateTeamResponseBody) *BatchCreateTeamResponse {
	s.Body = v
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
	// This parameter is required.
	DentryUuids []*string `json:"dentryUuids,omitempty" xml:"dentryUuids,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
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
	// example:
	//
	// true
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteRecentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CategoriesTemplatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CategoriesTemplatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s CategoriesTemplatesHeaders) GoString() string {
	return s.String()
}

func (s *CategoriesTemplatesHeaders) SetCommonHeaders(v map[string]*string) *CategoriesTemplatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CategoriesTemplatesHeaders) SetXAcsDingtalkAccessToken(v string) *CategoriesTemplatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CategoriesTemplatesRequest struct {
	Option *CategoriesTemplatesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	Param *CategoriesTemplatesRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CategoriesTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s CategoriesTemplatesRequest) GoString() string {
	return s.String()
}

func (s *CategoriesTemplatesRequest) SetOption(v *CategoriesTemplatesRequestOption) *CategoriesTemplatesRequest {
	s.Option = v
	return s
}

func (s *CategoriesTemplatesRequest) SetParam(v *CategoriesTemplatesRequestParam) *CategoriesTemplatesRequest {
	s.Param = v
	return s
}

func (s *CategoriesTemplatesRequest) SetOperatorId(v string) *CategoriesTemplatesRequest {
	s.OperatorId = &v
	return s
}

type CategoriesTemplatesRequestOption struct {
	// example:
	//
	// 1
	CategoryStatus *int32 `json:"categoryStatus,omitempty" xml:"categoryStatus,omitempty"`
	// example:
	//
	// pc
	QueryPlatform *string `json:"queryPlatform,omitempty" xml:"queryPlatform,omitempty"`
	// example:
	//
	// 1
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"templateStatus,omitempty" xml:"templateStatus,omitempty"`
}

func (s CategoriesTemplatesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s CategoriesTemplatesRequestOption) GoString() string {
	return s.String()
}

func (s *CategoriesTemplatesRequestOption) SetCategoryStatus(v int32) *CategoriesTemplatesRequestOption {
	s.CategoryStatus = &v
	return s
}

func (s *CategoriesTemplatesRequestOption) SetQueryPlatform(v string) *CategoriesTemplatesRequestOption {
	s.QueryPlatform = &v
	return s
}

func (s *CategoriesTemplatesRequestOption) SetSize(v int32) *CategoriesTemplatesRequestOption {
	s.Size = &v
	return s
}

func (s *CategoriesTemplatesRequestOption) SetTemplateStatus(v int32) *CategoriesTemplatesRequestOption {
	s.TemplateStatus = &v
	return s
}

type CategoriesTemplatesRequestParam struct {
	// This parameter is required.
	//
	// example:
	//
	// categoryIds
	CategoryIds []*string `json:"categoryIds,omitempty" xml:"categoryIds,omitempty" type:"Repeated"`
}

func (s CategoriesTemplatesRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CategoriesTemplatesRequestParam) GoString() string {
	return s.String()
}

func (s *CategoriesTemplatesRequestParam) SetCategoryIds(v []*string) *CategoriesTemplatesRequestParam {
	s.CategoryIds = v
	return s
}

type CategoriesTemplatesResponseBody struct {
	Map map[string][]*MapValue `json:"map,omitempty" xml:"map,omitempty"`
}

func (s CategoriesTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CategoriesTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *CategoriesTemplatesResponseBody) SetMap(v map[string][]*MapValue) *CategoriesTemplatesResponseBody {
	s.Map = v
	return s
}

type CategoriesTemplatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CategoriesTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CategoriesTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s CategoriesTemplatesResponse) GoString() string {
	return s.String()
}

func (s *CategoriesTemplatesResponse) SetHeaders(v map[string]*string) *CategoriesTemplatesResponse {
	s.Headers = v
	return s
}

func (s *CategoriesTemplatesResponse) SetStatusCode(v int32) *CategoriesTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *CategoriesTemplatesResponse) SetBody(v *CategoriesTemplatesResponseBody) *CategoriesTemplatesResponse {
	s.Body = v
	return s
}

type CategoryTemplatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CategoryTemplatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s CategoryTemplatesHeaders) GoString() string {
	return s.String()
}

func (s *CategoryTemplatesHeaders) SetCommonHeaders(v map[string]*string) *CategoryTemplatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CategoryTemplatesHeaders) SetXAcsDingtalkAccessToken(v string) *CategoryTemplatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CategoryTemplatesRequest struct {
	Option *CategoryTemplatesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	Param *CategoryTemplatesRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CategoryTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s CategoryTemplatesRequest) GoString() string {
	return s.String()
}

func (s *CategoryTemplatesRequest) SetOption(v *CategoryTemplatesRequestOption) *CategoryTemplatesRequest {
	s.Option = v
	return s
}

func (s *CategoryTemplatesRequest) SetParam(v *CategoryTemplatesRequestParam) *CategoryTemplatesRequest {
	s.Param = v
	return s
}

func (s *CategoryTemplatesRequest) SetOperatorId(v string) *CategoryTemplatesRequest {
	s.OperatorId = &v
	return s
}

type CategoryTemplatesRequestOption struct {
	// example:
	//
	// 1
	CategoryStatus *int32 `json:"categoryStatus,omitempty" xml:"categoryStatus,omitempty"`
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"templateStatus,omitempty" xml:"templateStatus,omitempty"`
}

func (s CategoryTemplatesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s CategoryTemplatesRequestOption) GoString() string {
	return s.String()
}

func (s *CategoryTemplatesRequestOption) SetCategoryStatus(v int32) *CategoryTemplatesRequestOption {
	s.CategoryStatus = &v
	return s
}

func (s *CategoryTemplatesRequestOption) SetTemplateStatus(v int32) *CategoryTemplatesRequestOption {
	s.TemplateStatus = &v
	return s
}

type CategoryTemplatesRequestParam struct {
	// This parameter is required.
	//
	// example:
	//
	// categoryId
	CategoryId *string `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
}

func (s CategoryTemplatesRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CategoryTemplatesRequestParam) GoString() string {
	return s.String()
}

func (s *CategoryTemplatesRequestParam) SetCategoryId(v string) *CategoryTemplatesRequestParam {
	s.CategoryId = &v
	return s
}

type CategoryTemplatesResponseBody struct {
	List []*CategoryTemplatesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s CategoryTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CategoryTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *CategoryTemplatesResponseBody) SetList(v []*CategoryTemplatesResponseBodyList) *CategoryTemplatesResponseBody {
	s.List = v
	return s
}

type CategoryTemplatesResponseBodyList struct {
	AuthorName         *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	Belong             *string `json:"belong,omitempty" xml:"belong,omitempty"`
	ContentDownloadUrl *string `json:"contentDownloadUrl,omitempty" xml:"contentDownloadUrl,omitempty"`
	CoverDownloadUrl   *string `json:"coverDownloadUrl,omitempty" xml:"coverDownloadUrl,omitempty"`
	CreateTime         *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	ModifiedTime       *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	TemplateId         *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
	Type               *int32  `json:"type,omitempty" xml:"type,omitempty"`
	UsedCount          *int64  `json:"usedCount,omitempty" xml:"usedCount,omitempty"`
	WorkspaceId        *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	WorkspaceName      *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s CategoryTemplatesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s CategoryTemplatesResponseBodyList) GoString() string {
	return s.String()
}

func (s *CategoryTemplatesResponseBodyList) SetAuthorName(v string) *CategoryTemplatesResponseBodyList {
	s.AuthorName = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetBelong(v string) *CategoryTemplatesResponseBodyList {
	s.Belong = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetContentDownloadUrl(v string) *CategoryTemplatesResponseBodyList {
	s.ContentDownloadUrl = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetCoverDownloadUrl(v string) *CategoryTemplatesResponseBodyList {
	s.CoverDownloadUrl = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetCreateTime(v string) *CategoryTemplatesResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetDescription(v string) *CategoryTemplatesResponseBodyList {
	s.Description = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetModifiedTime(v string) *CategoryTemplatesResponseBodyList {
	s.ModifiedTime = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetTemplateId(v string) *CategoryTemplatesResponseBodyList {
	s.TemplateId = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetTitle(v string) *CategoryTemplatesResponseBodyList {
	s.Title = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetType(v int32) *CategoryTemplatesResponseBodyList {
	s.Type = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetUsedCount(v int64) *CategoryTemplatesResponseBodyList {
	s.UsedCount = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetWorkspaceId(v string) *CategoryTemplatesResponseBodyList {
	s.WorkspaceId = &v
	return s
}

func (s *CategoryTemplatesResponseBodyList) SetWorkspaceName(v string) *CategoryTemplatesResponseBodyList {
	s.WorkspaceName = &v
	return s
}

type CategoryTemplatesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CategoryTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CategoryTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s CategoryTemplatesResponse) GoString() string {
	return s.String()
}

func (s *CategoryTemplatesResponse) SetHeaders(v map[string]*string) *CategoryTemplatesResponse {
	s.Headers = v
	return s
}

func (s *CategoryTemplatesResponse) SetStatusCode(v int32) *CategoryTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *CategoryTemplatesResponse) SetBody(v *CategoryTemplatesResponseBody) *CategoryTemplatesResponse {
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
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DentryVO          `json:"body,omitempty" xml:"body,omitempty"`
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

type CopyDocHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CopyDocHeaders) String() string {
	return tea.Prettify(s)
}

func (s CopyDocHeaders) GoString() string {
	return s.String()
}

func (s *CopyDocHeaders) SetCommonHeaders(v map[string]*string) *CopyDocHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyDocHeaders) SetXAcsDingtalkAccessToken(v string) *CopyDocHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CopyDocRequest struct {
	// This parameter is required.
	Param *CopyDocRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
}

func (s CopyDocRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyDocRequest) GoString() string {
	return s.String()
}

func (s *CopyDocRequest) SetParam(v *CopyDocRequestParam) *CopyDocRequest {
	s.Param = v
	return s
}

type CopyDocRequestParam struct {
	// This parameter is required.
	//
	// example:
	//
	// dentryUuid
	SourceDentryUuid *string `json:"sourceDentryUuid,omitempty" xml:"sourceDentryUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dentryUuid
	TargetParentDentryUuid *string `json:"targetParentDentryUuid,omitempty" xml:"targetParentDentryUuid,omitempty"`
	// example:
	//
	// dentryUuid
	TargetPreDentryUuid *string `json:"targetPreDentryUuid,omitempty" xml:"targetPreDentryUuid,omitempty"`
}

func (s CopyDocRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CopyDocRequestParam) GoString() string {
	return s.String()
}

func (s *CopyDocRequestParam) SetSourceDentryUuid(v string) *CopyDocRequestParam {
	s.SourceDentryUuid = &v
	return s
}

func (s *CopyDocRequestParam) SetTargetParentDentryUuid(v string) *CopyDocRequestParam {
	s.TargetParentDentryUuid = &v
	return s
}

func (s *CopyDocRequestParam) SetTargetPreDentryUuid(v string) *CopyDocRequestParam {
	s.TargetPreDentryUuid = &v
	return s
}

type CopyDocResponseBody struct {
	// example:
	//
	// true
	IsAsync *bool `json:"isAsync,omitempty" xml:"isAsync,omitempty"`
	// example:
	//
	// true
	SyncCopyResult *CopyDocResponseBodySyncCopyResult `json:"syncCopyResult,omitempty" xml:"syncCopyResult,omitempty" type:"Struct"`
	// example:
	//
	// true
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CopyDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyDocResponseBody) GoString() string {
	return s.String()
}

func (s *CopyDocResponseBody) SetIsAsync(v bool) *CopyDocResponseBody {
	s.IsAsync = &v
	return s
}

func (s *CopyDocResponseBody) SetSyncCopyResult(v *CopyDocResponseBodySyncCopyResult) *CopyDocResponseBody {
	s.SyncCopyResult = v
	return s
}

func (s *CopyDocResponseBody) SetTaskId(v string) *CopyDocResponseBody {
	s.TaskId = &v
	return s
}

type CopyDocResponseBodySyncCopyResult struct {
	// example:
	//
	// dentryUuid
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// example:
	//
	// driveDentryId
	DriveDentryId *string `json:"driveDentryId,omitempty" xml:"driveDentryId,omitempty"`
	// example:
	//
	// driveSpaceId
	DriveSpaceId *string `json:"driveSpaceId,omitempty" xml:"driveSpaceId,omitempty"`
	// example:
	//
	// docx
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// spaceInfo
	SpaceInfo *CopyDocResponseBodySyncCopyResultSpaceInfo `json:"spaceInfo,omitempty" xml:"spaceInfo,omitempty" type:"Struct"`
	// example:
	//
	// https://example.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CopyDocResponseBodySyncCopyResult) String() string {
	return tea.Prettify(s)
}

func (s CopyDocResponseBodySyncCopyResult) GoString() string {
	return s.String()
}

func (s *CopyDocResponseBodySyncCopyResult) SetDentryUuid(v string) *CopyDocResponseBodySyncCopyResult {
	s.DentryUuid = &v
	return s
}

func (s *CopyDocResponseBodySyncCopyResult) SetDriveDentryId(v string) *CopyDocResponseBodySyncCopyResult {
	s.DriveDentryId = &v
	return s
}

func (s *CopyDocResponseBodySyncCopyResult) SetDriveSpaceId(v string) *CopyDocResponseBodySyncCopyResult {
	s.DriveSpaceId = &v
	return s
}

func (s *CopyDocResponseBodySyncCopyResult) SetExtension(v string) *CopyDocResponseBodySyncCopyResult {
	s.Extension = &v
	return s
}

func (s *CopyDocResponseBodySyncCopyResult) SetName(v string) *CopyDocResponseBodySyncCopyResult {
	s.Name = &v
	return s
}

func (s *CopyDocResponseBodySyncCopyResult) SetSpaceInfo(v *CopyDocResponseBodySyncCopyResultSpaceInfo) *CopyDocResponseBodySyncCopyResult {
	s.SpaceInfo = v
	return s
}

func (s *CopyDocResponseBodySyncCopyResult) SetUrl(v string) *CopyDocResponseBodySyncCopyResult {
	s.Url = &v
	return s
}

type CopyDocResponseBodySyncCopyResultSpaceInfo struct {
	// example:
	//
	// im
	SceneType *string `json:"sceneType,omitempty" xml:"sceneType,omitempty"`
}

func (s CopyDocResponseBodySyncCopyResultSpaceInfo) String() string {
	return tea.Prettify(s)
}

func (s CopyDocResponseBodySyncCopyResultSpaceInfo) GoString() string {
	return s.String()
}

func (s *CopyDocResponseBodySyncCopyResultSpaceInfo) SetSceneType(v string) *CopyDocResponseBodySyncCopyResultSpaceInfo {
	s.SceneType = &v
	return s
}

type CopyDocResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyDocResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyDocResponse) GoString() string {
	return s.String()
}

func (s *CopyDocResponse) SetHeaders(v map[string]*string) *CopyDocResponse {
	s.Headers = v
	return s
}

func (s *CopyDocResponse) SetStatusCode(v int32) *CopyDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDocResponse) SetBody(v *CopyDocResponseBody) *CopyDocResponse {
	s.Body = v
	return s
}

type CopyWorkspaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CopyWorkspaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CopyWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *CopyWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *CopyWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyWorkspaceHeaders) SetXAcsDingtalkAccessToken(v string) *CopyWorkspaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CopyWorkspaceRequest struct {
	// This parameter is required.
	Param *CopyWorkspaceRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
}

func (s CopyWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CopyWorkspaceRequest) SetParam(v *CopyWorkspaceRequestParam) *CopyWorkspaceRequest {
	s.Param = v
	return s
}

type CopyWorkspaceRequestParam struct {
	// example:
	//
	// 0
	OriginWorkspaceId *string `json:"originWorkspaceId,omitempty" xml:"originWorkspaceId,omitempty"`
}

func (s CopyWorkspaceRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CopyWorkspaceRequestParam) GoString() string {
	return s.String()
}

func (s *CopyWorkspaceRequestParam) SetOriginWorkspaceId(v string) *CopyWorkspaceRequestParam {
	s.OriginWorkspaceId = &v
	return s
}

type CopyWorkspaceResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CopyWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CopyWorkspaceResponseBody) SetSuccess(v bool) *CopyWorkspaceResponseBody {
	s.Success = &v
	return s
}

type CopyWorkspaceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CopyWorkspaceResponse) SetHeaders(v map[string]*string) *CopyWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CopyWorkspaceResponse) SetStatusCode(v int32) *CopyWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyWorkspaceResponse) SetBody(v *CopyWorkspaceResponseBody) *CopyWorkspaceResponse {
	s.Body = v
	return s
}

type CopyWorkspaceAsyncHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CopyWorkspaceAsyncHeaders) String() string {
	return tea.Prettify(s)
}

func (s CopyWorkspaceAsyncHeaders) GoString() string {
	return s.String()
}

func (s *CopyWorkspaceAsyncHeaders) SetCommonHeaders(v map[string]*string) *CopyWorkspaceAsyncHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyWorkspaceAsyncHeaders) SetXAcsDingtalkAccessToken(v string) *CopyWorkspaceAsyncHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CopyWorkspaceAsyncRequest struct {
	// This parameter is required.
	Param *CopyWorkspaceAsyncRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
}

func (s CopyWorkspaceAsyncRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyWorkspaceAsyncRequest) GoString() string {
	return s.String()
}

func (s *CopyWorkspaceAsyncRequest) SetParam(v *CopyWorkspaceAsyncRequestParam) *CopyWorkspaceAsyncRequest {
	s.Param = v
	return s
}

type CopyWorkspaceAsyncRequestParam struct {
	// example:
	//
	// 0
	OriginWorkspaceId *string `json:"originWorkspaceId,omitempty" xml:"originWorkspaceId,omitempty"`
}

func (s CopyWorkspaceAsyncRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CopyWorkspaceAsyncRequestParam) GoString() string {
	return s.String()
}

func (s *CopyWorkspaceAsyncRequestParam) SetOriginWorkspaceId(v string) *CopyWorkspaceAsyncRequestParam {
	s.OriginWorkspaceId = &v
	return s
}

type CopyWorkspaceAsyncResponseBody struct {
	// example:
	//
	// 0
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CopyWorkspaceAsyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyWorkspaceAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *CopyWorkspaceAsyncResponseBody) SetId(v int64) *CopyWorkspaceAsyncResponseBody {
	s.Id = &v
	return s
}

type CopyWorkspaceAsyncResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyWorkspaceAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyWorkspaceAsyncResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyWorkspaceAsyncResponse) GoString() string {
	return s.String()
}

func (s *CopyWorkspaceAsyncResponse) SetHeaders(v map[string]*string) *CopyWorkspaceAsyncResponse {
	s.Headers = v
	return s
}

func (s *CopyWorkspaceAsyncResponse) SetStatusCode(v int32) *CopyWorkspaceAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyWorkspaceAsyncResponse) SetBody(v *CopyWorkspaceAsyncResponseBody) *CopyWorkspaceAsyncResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// file
	DentryType *string `json:"dentryType,omitempty" xml:"dentryType,omitempty"`
	// example:
	//
	// 0
	DocumentType *int64 `json:"documentType,omitempty" xml:"documentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingtalk
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// abcedf
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DentryVO          `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateShortcutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateShortcutHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateShortcutHeaders) GoString() string {
	return s.String()
}

func (s *CreateShortcutHeaders) SetCommonHeaders(v map[string]*string) *CreateShortcutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateShortcutHeaders) SetXAcsDingtalkAccessToken(v string) *CreateShortcutHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateShortcutRequest struct {
	// This parameter is required.
	Param *CreateShortcutRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
}

func (s CreateShortcutRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateShortcutRequest) GoString() string {
	return s.String()
}

func (s *CreateShortcutRequest) SetParam(v *CreateShortcutRequestParam) *CreateShortcutRequest {
	s.Param = v
	return s
}

type CreateShortcutRequestParam struct {
	// This parameter is required.
	//
	// example:
	//
	// dentryUuid
	SourceResourceId *string `json:"sourceResourceId,omitempty" xml:"sourceResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DENTRY
	SourceResourceType *string `json:"sourceResourceType,omitempty" xml:"sourceResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dentryUuid
	TargetResourceId *string `json:"targetResourceId,omitempty" xml:"targetResourceId,omitempty"`
	// example:
	//
	// 123test
	TargetResourceName *string `json:"targetResourceName,omitempty" xml:"targetResourceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DENTRY
	TargetResourceType *string `json:"targetResourceType,omitempty" xml:"targetResourceType,omitempty"`
}

func (s CreateShortcutRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CreateShortcutRequestParam) GoString() string {
	return s.String()
}

func (s *CreateShortcutRequestParam) SetSourceResourceId(v string) *CreateShortcutRequestParam {
	s.SourceResourceId = &v
	return s
}

func (s *CreateShortcutRequestParam) SetSourceResourceType(v string) *CreateShortcutRequestParam {
	s.SourceResourceType = &v
	return s
}

func (s *CreateShortcutRequestParam) SetTargetResourceId(v string) *CreateShortcutRequestParam {
	s.TargetResourceId = &v
	return s
}

func (s *CreateShortcutRequestParam) SetTargetResourceName(v string) *CreateShortcutRequestParam {
	s.TargetResourceName = &v
	return s
}

func (s *CreateShortcutRequestParam) SetTargetResourceType(v string) *CreateShortcutRequestParam {
	s.TargetResourceType = &v
	return s
}

type CreateShortcutResponseBody struct {
	// example:
	//
	// dentry
	OpenDentryInfo *CreateShortcutResponseBodyOpenDentryInfo `json:"openDentryInfo,omitempty" xml:"openDentryInfo,omitempty" type:"Struct"`
}

func (s CreateShortcutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateShortcutResponseBody) GoString() string {
	return s.String()
}

func (s *CreateShortcutResponseBody) SetOpenDentryInfo(v *CreateShortcutResponseBodyOpenDentryInfo) *CreateShortcutResponseBody {
	s.OpenDentryInfo = v
	return s
}

type CreateShortcutResponseBodyOpenDentryInfo struct {
	// example:
	//
	// dentryUuid
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// example:
	//
	// driveDentryId
	DriveDentryId *string `json:"driveDentryId,omitempty" xml:"driveDentryId,omitempty"`
	// example:
	//
	// driveSpaceId
	DriveSpaceId *string `json:"driveSpaceId,omitempty" xml:"driveSpaceId,omitempty"`
	// example:
	//
	// docx
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// spaceInfo
	SpaceInfo *CreateShortcutResponseBodyOpenDentryInfoSpaceInfo `json:"spaceInfo,omitempty" xml:"spaceInfo,omitempty" type:"Struct"`
	// example:
	//
	// https://example.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateShortcutResponseBodyOpenDentryInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateShortcutResponseBodyOpenDentryInfo) GoString() string {
	return s.String()
}

func (s *CreateShortcutResponseBodyOpenDentryInfo) SetDentryUuid(v string) *CreateShortcutResponseBodyOpenDentryInfo {
	s.DentryUuid = &v
	return s
}

func (s *CreateShortcutResponseBodyOpenDentryInfo) SetDriveDentryId(v string) *CreateShortcutResponseBodyOpenDentryInfo {
	s.DriveDentryId = &v
	return s
}

func (s *CreateShortcutResponseBodyOpenDentryInfo) SetDriveSpaceId(v string) *CreateShortcutResponseBodyOpenDentryInfo {
	s.DriveSpaceId = &v
	return s
}

func (s *CreateShortcutResponseBodyOpenDentryInfo) SetExtension(v string) *CreateShortcutResponseBodyOpenDentryInfo {
	s.Extension = &v
	return s
}

func (s *CreateShortcutResponseBodyOpenDentryInfo) SetName(v string) *CreateShortcutResponseBodyOpenDentryInfo {
	s.Name = &v
	return s
}

func (s *CreateShortcutResponseBodyOpenDentryInfo) SetSpaceInfo(v *CreateShortcutResponseBodyOpenDentryInfoSpaceInfo) *CreateShortcutResponseBodyOpenDentryInfo {
	s.SpaceInfo = v
	return s
}

func (s *CreateShortcutResponseBodyOpenDentryInfo) SetUrl(v string) *CreateShortcutResponseBodyOpenDentryInfo {
	s.Url = &v
	return s
}

type CreateShortcutResponseBodyOpenDentryInfoSpaceInfo struct {
	// example:
	//
	// im
	SceneType *string `json:"sceneType,omitempty" xml:"sceneType,omitempty"`
}

func (s CreateShortcutResponseBodyOpenDentryInfoSpaceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateShortcutResponseBodyOpenDentryInfoSpaceInfo) GoString() string {
	return s.String()
}

func (s *CreateShortcutResponseBodyOpenDentryInfoSpaceInfo) SetSceneType(v string) *CreateShortcutResponseBodyOpenDentryInfoSpaceInfo {
	s.SceneType = &v
	return s
}

type CreateShortcutResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateShortcutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateShortcutResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateShortcutResponse) GoString() string {
	return s.String()
}

func (s *CreateShortcutResponse) SetHeaders(v map[string]*string) *CreateShortcutResponse {
	s.Headers = v
	return s
}

func (s *CreateShortcutResponse) SetStatusCode(v int32) *CreateShortcutResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateShortcutResponse) SetBody(v *CreateShortcutResponseBody) *CreateShortcutResponse {
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
	// example:
	//
	// 这里是知识库的简介
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1***.png
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试知识库
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// l6gYG9****mo9Z
	SectionId *string `json:"sectionId,omitempty" xml:"sectionId,omitempty"`
	// This parameter is required.
	ShareScope *CreateSpaceRequestShareScope `json:"shareScope,omitempty" xml:"shareScope,omitempty" type:"Struct"`
	// example:
	//
	// 5YRB***GDAr
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 0
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SpaceVO           `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1***.png
	Cover *string `json:"cover,omitempty" xml:"cover,omitempty"`
	// example:
	//
	// 这是小组的介绍
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1***.png
	Icon    *string                     `json:"icon,omitempty" xml:"icon,omitempty"`
	Members []*CreateTeamRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 测试小组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	TeamType *int32 `json:"teamType,omitempty" xml:"teamType,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	MemberType *int32 `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TeamVO            `json:"body,omitempty" xml:"body,omitempty"`
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

type CrossOrgMigrateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CrossOrgMigrateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CrossOrgMigrateHeaders) GoString() string {
	return s.String()
}

func (s *CrossOrgMigrateHeaders) SetCommonHeaders(v map[string]*string) *CrossOrgMigrateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CrossOrgMigrateHeaders) SetXAcsDingtalkAccessToken(v string) *CrossOrgMigrateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CrossOrgMigrateRequest struct {
	Option *CrossOrgMigrateRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	Param *CrossOrgMigrateRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CrossOrgMigrateRequest) String() string {
	return tea.Prettify(s)
}

func (s CrossOrgMigrateRequest) GoString() string {
	return s.String()
}

func (s *CrossOrgMigrateRequest) SetOption(v *CrossOrgMigrateRequestOption) *CrossOrgMigrateRequest {
	s.Option = v
	return s
}

func (s *CrossOrgMigrateRequest) SetParam(v *CrossOrgMigrateRequestParam) *CrossOrgMigrateRequest {
	s.Param = v
	return s
}

func (s *CrossOrgMigrateRequest) SetOperatorId(v string) *CrossOrgMigrateRequest {
	s.OperatorId = &v
	return s
}

type CrossOrgMigrateRequestOption struct {
	// example:
	//
	// true
	CheckOperatorSourceRole *bool `json:"checkOperatorSourceRole,omitempty" xml:"checkOperatorSourceRole,omitempty"`
	// example:
	//
	// true
	DeleteSource *bool `json:"deleteSource,omitempty" xml:"deleteSource,omitempty"`
	// example:
	//
	// true
	NeedRecycleFailedWorkspaceId *bool `json:"needRecycleFailedWorkspaceId,omitempty" xml:"needRecycleFailedWorkspaceId,omitempty"`
	// example:
	//
	// 1L
	RelateTeamId *int64 `json:"relateTeamId,omitempty" xml:"relateTeamId,omitempty"`
	// example:
	//
	// team_id
	RelateTeamIdStr *string `json:"relateTeamIdStr,omitempty" xml:"relateTeamIdStr,omitempty"`
	// example:
	//
	// true
	RetainOrgGroup *bool `json:"retainOrgGroup,omitempty" xml:"retainOrgGroup,omitempty"`
	// example:
	//
	// true
	SkipRole        *bool     `json:"skipRole,omitempty" xml:"skipRole,omitempty"`
	WorkspaceIdStrs []*string `json:"workspaceIdStrs,omitempty" xml:"workspaceIdStrs,omitempty" type:"Repeated"`
	WorkspaceIds    []*int64  `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty" type:"Repeated"`
}

func (s CrossOrgMigrateRequestOption) String() string {
	return tea.Prettify(s)
}

func (s CrossOrgMigrateRequestOption) GoString() string {
	return s.String()
}

func (s *CrossOrgMigrateRequestOption) SetCheckOperatorSourceRole(v bool) *CrossOrgMigrateRequestOption {
	s.CheckOperatorSourceRole = &v
	return s
}

func (s *CrossOrgMigrateRequestOption) SetDeleteSource(v bool) *CrossOrgMigrateRequestOption {
	s.DeleteSource = &v
	return s
}

func (s *CrossOrgMigrateRequestOption) SetNeedRecycleFailedWorkspaceId(v bool) *CrossOrgMigrateRequestOption {
	s.NeedRecycleFailedWorkspaceId = &v
	return s
}

func (s *CrossOrgMigrateRequestOption) SetRelateTeamId(v int64) *CrossOrgMigrateRequestOption {
	s.RelateTeamId = &v
	return s
}

func (s *CrossOrgMigrateRequestOption) SetRelateTeamIdStr(v string) *CrossOrgMigrateRequestOption {
	s.RelateTeamIdStr = &v
	return s
}

func (s *CrossOrgMigrateRequestOption) SetRetainOrgGroup(v bool) *CrossOrgMigrateRequestOption {
	s.RetainOrgGroup = &v
	return s
}

func (s *CrossOrgMigrateRequestOption) SetSkipRole(v bool) *CrossOrgMigrateRequestOption {
	s.SkipRole = &v
	return s
}

func (s *CrossOrgMigrateRequestOption) SetWorkspaceIdStrs(v []*string) *CrossOrgMigrateRequestOption {
	s.WorkspaceIdStrs = v
	return s
}

func (s *CrossOrgMigrateRequestOption) SetWorkspaceIds(v []*int64) *CrossOrgMigrateRequestOption {
	s.WorkspaceIds = v
	return s
}

type CrossOrgMigrateRequestParam struct {
	// example:
	//
	// corp_id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s CrossOrgMigrateRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CrossOrgMigrateRequestParam) GoString() string {
	return s.String()
}

func (s *CrossOrgMigrateRequestParam) SetCorpId(v string) *CrossOrgMigrateRequestParam {
	s.CorpId = &v
	return s
}

type CrossOrgMigrateResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CrossOrgMigrateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CrossOrgMigrateResponseBody) GoString() string {
	return s.String()
}

func (s *CrossOrgMigrateResponseBody) SetSuccess(v bool) *CrossOrgMigrateResponseBody {
	s.Success = &v
	return s
}

type CrossOrgMigrateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CrossOrgMigrateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CrossOrgMigrateResponse) String() string {
	return tea.Prettify(s)
}

func (s CrossOrgMigrateResponse) GoString() string {
	return s.String()
}

func (s *CrossOrgMigrateResponse) SetHeaders(v map[string]*string) *CrossOrgMigrateResponse {
	s.Headers = v
	return s
}

func (s *CrossOrgMigrateResponse) SetStatusCode(v int32) *CrossOrgMigrateResponse {
	s.StatusCode = &v
	return s
}

func (s *CrossOrgMigrateResponse) SetBody(v *CrossOrgMigrateResponseBody) *CrossOrgMigrateResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TeamVO            `json:"body,omitempty" xml:"body,omitempty"`
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

type DocContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DocContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s DocContentHeaders) GoString() string {
	return s.String()
}

func (s *DocContentHeaders) SetCommonHeaders(v map[string]*string) *DocContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocContentHeaders) SetXAcsDingtalkAccessToken(v string) *DocContentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DocContentRequest struct {
	Option *DocContentRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DocContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DocContentRequest) GoString() string {
	return s.String()
}

func (s *DocContentRequest) SetOption(v *DocContentRequestOption) *DocContentRequest {
	s.Option = v
	return s
}

func (s *DocContentRequest) SetOperatorId(v string) *DocContentRequest {
	s.OperatorId = &v
	return s
}

type DocContentRequestOption struct {
	// example:
	//
	// markdown
	TargetFormat *string `json:"targetFormat,omitempty" xml:"targetFormat,omitempty"`
}

func (s DocContentRequestOption) String() string {
	return tea.Prettify(s)
}

func (s DocContentRequestOption) GoString() string {
	return s.String()
}

func (s *DocContentRequestOption) SetTargetFormat(v string) *DocContentRequestOption {
	s.TargetFormat = &v
	return s
}

type DocContentResponseBody struct {
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DocContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DocContentResponseBody) GoString() string {
	return s.String()
}

func (s *DocContentResponseBody) SetTaskId(v int64) *DocContentResponseBody {
	s.TaskId = &v
	return s
}

type DocContentResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DocContentResponse) GoString() string {
	return s.String()
}

func (s *DocContentResponse) SetHeaders(v map[string]*string) *DocContentResponse {
	s.Headers = v
	return s
}

func (s *DocContentResponse) SetStatusCode(v int32) *DocContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DocContentResponse) SetBody(v *DocContentResponseBody) *DocContentResponse {
	s.Body = v
	return s
}

type DocExportByDelegatedPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DocExportByDelegatedPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DocExportByDelegatedPermissionHeaders) GoString() string {
	return s.String()
}

func (s *DocExportByDelegatedPermissionHeaders) SetCommonHeaders(v map[string]*string) *DocExportByDelegatedPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocExportByDelegatedPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *DocExportByDelegatedPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DocExportByDelegatedPermissionRequest struct {
	// example:
	//
	// false
	GenerateCp *bool `json:"generateCp,omitempty" xml:"generateCp,omitempty"`
	// example:
	//
	// markdown
	TargetFormat *string `json:"targetFormat,omitempty" xml:"targetFormat,omitempty"`
}

func (s DocExportByDelegatedPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DocExportByDelegatedPermissionRequest) GoString() string {
	return s.String()
}

func (s *DocExportByDelegatedPermissionRequest) SetGenerateCp(v bool) *DocExportByDelegatedPermissionRequest {
	s.GenerateCp = &v
	return s
}

func (s *DocExportByDelegatedPermissionRequest) SetTargetFormat(v string) *DocExportByDelegatedPermissionRequest {
	s.TargetFormat = &v
	return s
}

type DocExportByDelegatedPermissionResponseBody struct {
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DocExportByDelegatedPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DocExportByDelegatedPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DocExportByDelegatedPermissionResponseBody) SetTaskId(v int64) *DocExportByDelegatedPermissionResponseBody {
	s.TaskId = &v
	return s
}

type DocExportByDelegatedPermissionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocExportByDelegatedPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocExportByDelegatedPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DocExportByDelegatedPermissionResponse) GoString() string {
	return s.String()
}

func (s *DocExportByDelegatedPermissionResponse) SetHeaders(v map[string]*string) *DocExportByDelegatedPermissionResponse {
	s.Headers = v
	return s
}

func (s *DocExportByDelegatedPermissionResponse) SetStatusCode(v int32) *DocExportByDelegatedPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DocExportByDelegatedPermissionResponse) SetBody(v *DocExportByDelegatedPermissionResponseBody) *DocExportByDelegatedPermissionResponse {
	s.Body = v
	return s
}

type DocUpdateContentWithDelegatedPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DocUpdateContentWithDelegatedPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DocUpdateContentWithDelegatedPermissionHeaders) GoString() string {
	return s.String()
}

func (s *DocUpdateContentWithDelegatedPermissionHeaders) SetCommonHeaders(v map[string]*string) *DocUpdateContentWithDelegatedPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocUpdateContentWithDelegatedPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *DocUpdateContentWithDelegatedPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DocUpdateContentWithDelegatedPermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// data_type
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
}

func (s DocUpdateContentWithDelegatedPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DocUpdateContentWithDelegatedPermissionRequest) GoString() string {
	return s.String()
}

func (s *DocUpdateContentWithDelegatedPermissionRequest) SetContent(v string) *DocUpdateContentWithDelegatedPermissionRequest {
	s.Content = &v
	return s
}

func (s *DocUpdateContentWithDelegatedPermissionRequest) SetDataType(v string) *DocUpdateContentWithDelegatedPermissionRequest {
	s.DataType = &v
	return s
}

type DocUpdateContentWithDelegatedPermissionResponseBody struct {
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DocUpdateContentWithDelegatedPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DocUpdateContentWithDelegatedPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DocUpdateContentWithDelegatedPermissionResponseBody) SetData(v map[string]interface{}) *DocUpdateContentWithDelegatedPermissionResponseBody {
	s.Data = v
	return s
}

type DocUpdateContentWithDelegatedPermissionResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocUpdateContentWithDelegatedPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocUpdateContentWithDelegatedPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DocUpdateContentWithDelegatedPermissionResponse) GoString() string {
	return s.String()
}

func (s *DocUpdateContentWithDelegatedPermissionResponse) SetHeaders(v map[string]*string) *DocUpdateContentWithDelegatedPermissionResponse {
	s.Headers = v
	return s
}

func (s *DocUpdateContentWithDelegatedPermissionResponse) SetStatusCode(v int32) *DocUpdateContentWithDelegatedPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DocUpdateContentWithDelegatedPermissionResponse) SetBody(v *DocUpdateContentWithDelegatedPermissionResponseBody) *DocUpdateContentWithDelegatedPermissionResponse {
	s.Body = v
	return s
}

type ExportDocHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExportDocHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExportDocHeaders) GoString() string {
	return s.String()
}

func (s *ExportDocHeaders) SetCommonHeaders(v map[string]*string) *ExportDocHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExportDocHeaders) SetXAcsDingtalkAccessToken(v string) *ExportDocHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExportDocRequest struct {
	// This parameter is required.
	Param *ExportDocRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
}

func (s ExportDocRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportDocRequest) GoString() string {
	return s.String()
}

func (s *ExportDocRequest) SetParam(v *ExportDocRequestParam) *ExportDocRequest {
	s.Param = v
	return s
}

type ExportDocRequestParam struct {
	// This parameter is required.
	//
	// example:
	//
	// dentryUuid
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingTalksheetToxlsx
	ExportType *string `json:"exportType,omitempty" xml:"exportType,omitempty"`
}

func (s ExportDocRequestParam) String() string {
	return tea.Prettify(s)
}

func (s ExportDocRequestParam) GoString() string {
	return s.String()
}

func (s *ExportDocRequestParam) SetDentryUuid(v string) *ExportDocRequestParam {
	s.DentryUuid = &v
	return s
}

func (s *ExportDocRequestParam) SetExportType(v string) *ExportDocRequestParam {
	s.ExportType = &v
	return s
}

type ExportDocResponseBody struct {
	// example:
	//
	// 12345678
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// init
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ExportDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportDocResponseBody) GoString() string {
	return s.String()
}

func (s *ExportDocResponseBody) SetJobId(v string) *ExportDocResponseBody {
	s.JobId = &v
	return s
}

func (s *ExportDocResponseBody) SetStatus(v string) *ExportDocResponseBody {
	s.Status = &v
	return s
}

type ExportDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportDocResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportDocResponse) GoString() string {
	return s.String()
}

func (s *ExportDocResponse) SetHeaders(v map[string]*string) *ExportDocResponse {
	s.Headers = v
	return s
}

func (s *ExportDocResponse) SetStatusCode(v int32) *ExportDocResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportDocResponse) SetBody(v *ExportDocResponseBody) *ExportDocResponse {
	s.Body = v
	return s
}

type GetDentryIdByUuidHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDentryIdByUuidHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDentryIdByUuidHeaders) GoString() string {
	return s.String()
}

func (s *GetDentryIdByUuidHeaders) SetCommonHeaders(v map[string]*string) *GetDentryIdByUuidHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDentryIdByUuidHeaders) SetXAcsDingtalkAccessToken(v string) *GetDentryIdByUuidHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDentryIdByUuidRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetDentryIdByUuidRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDentryIdByUuidRequest) GoString() string {
	return s.String()
}

func (s *GetDentryIdByUuidRequest) SetOperatorId(v string) *GetDentryIdByUuidRequest {
	s.OperatorId = &v
	return s
}

type GetDentryIdByUuidResponseBody struct {
	DentryId   *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	SpaceId    *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetDentryIdByUuidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDentryIdByUuidResponseBody) GoString() string {
	return s.String()
}

func (s *GetDentryIdByUuidResponseBody) SetDentryId(v string) *GetDentryIdByUuidResponseBody {
	s.DentryId = &v
	return s
}

func (s *GetDentryIdByUuidResponseBody) SetDentryUuid(v string) *GetDentryIdByUuidResponseBody {
	s.DentryUuid = &v
	return s
}

func (s *GetDentryIdByUuidResponseBody) SetSpaceId(v string) *GetDentryIdByUuidResponseBody {
	s.SpaceId = &v
	return s
}

type GetDentryIdByUuidResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDentryIdByUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDentryIdByUuidResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDentryIdByUuidResponse) GoString() string {
	return s.String()
}

func (s *GetDentryIdByUuidResponse) SetHeaders(v map[string]*string) *GetDentryIdByUuidResponse {
	s.Headers = v
	return s
}

func (s *GetDentryIdByUuidResponse) SetStatusCode(v int32) *GetDentryIdByUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDentryIdByUuidResponse) SetBody(v *GetDentryIdByUuidResponseBody) *GetDentryIdByUuidResponse {
	s.Body = v
	return s
}

type GetDocContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDocContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDocContentHeaders) GoString() string {
	return s.String()
}

func (s *GetDocContentHeaders) SetCommonHeaders(v map[string]*string) *GetDocContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDocContentHeaders) SetXAcsDingtalkAccessToken(v string) *GetDocContentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDocContentRequest struct {
	// example:
	//
	// false
	GenerateCp *bool `json:"generateCp,omitempty" xml:"generateCp,omitempty"`
	// example:
	//
	// markdown
	TargetFormat *string `json:"targetFormat,omitempty" xml:"targetFormat,omitempty"`
}

func (s GetDocContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocContentRequest) GoString() string {
	return s.String()
}

func (s *GetDocContentRequest) SetGenerateCp(v bool) *GetDocContentRequest {
	s.GenerateCp = &v
	return s
}

func (s *GetDocContentRequest) SetTargetFormat(v string) *GetDocContentRequest {
	s.TargetFormat = &v
	return s
}

type GetDocContentResponseBody struct {
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetDocContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocContentResponseBody) SetTaskId(v int64) *GetDocContentResponseBody {
	s.TaskId = &v
	return s
}

type GetDocContentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocContentResponse) GoString() string {
	return s.String()
}

func (s *GetDocContentResponse) SetHeaders(v map[string]*string) *GetDocContentResponse {
	s.Headers = v
	return s
}

func (s *GetDocContentResponse) SetStatusCode(v int32) *GetDocContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocContentResponse) SetBody(v *GetDocContentResponseBody) *GetDocContentResponse {
	s.Body = v
	return s
}

type GetDocContentForELMHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDocContentForELMHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDocContentForELMHeaders) GoString() string {
	return s.String()
}

func (s *GetDocContentForELMHeaders) SetCommonHeaders(v map[string]*string) *GetDocContentForELMHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDocContentForELMHeaders) SetXAcsDingtalkAccessToken(v string) *GetDocContentForELMHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDocContentForELMRequest struct {
	// example:
	//
	// markdown
	TargetFormat *string `json:"targetFormat,omitempty" xml:"targetFormat,omitempty"`
}

func (s GetDocContentForELMRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocContentForELMRequest) GoString() string {
	return s.String()
}

func (s *GetDocContentForELMRequest) SetTargetFormat(v string) *GetDocContentForELMRequest {
	s.TargetFormat = &v
	return s
}

type GetDocContentForELMResponseBody struct {
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetDocContentForELMResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocContentForELMResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocContentForELMResponseBody) SetTaskId(v int64) *GetDocContentForELMResponseBody {
	s.TaskId = &v
	return s
}

type GetDocContentForELMResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocContentForELMResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocContentForELMResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocContentForELMResponse) GoString() string {
	return s.String()
}

func (s *GetDocContentForELMResponse) SetHeaders(v map[string]*string) *GetDocContentForELMResponse {
	s.Headers = v
	return s
}

func (s *GetDocContentForELMResponse) SetStatusCode(v int32) *GetDocContentForELMResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocContentForELMResponse) SetBody(v *GetDocContentForELMResponseBody) *GetDocContentForELMResponse {
	s.Body = v
	return s
}

type GetMySpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMySpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMySpaceHeaders) GoString() string {
	return s.String()
}

func (s *GetMySpaceHeaders) SetCommonHeaders(v map[string]*string) *GetMySpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMySpaceHeaders) SetXAcsDingtalkAccessToken(v string) *GetMySpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMySpaceRequest struct {
	// example:
	//
	// true
	IsMySpace *bool `json:"isMySpace,omitempty" xml:"isMySpace,omitempty"`
}

func (s GetMySpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMySpaceRequest) GoString() string {
	return s.String()
}

func (s *GetMySpaceRequest) SetIsMySpace(v bool) *GetMySpaceRequest {
	s.IsMySpace = &v
	return s
}

type GetMySpaceResponseBody struct {
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// example:
	//
	// 1L
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// space_name
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	// example:
	//
	// personal|org|custom|unknown
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// example:
	//
	// 1L
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s GetMySpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMySpaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMySpaceResponseBody) SetCreateTime(v string) *GetMySpaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetMySpaceResponseBody) SetModifyTime(v string) *GetMySpaceResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetMySpaceResponseBody) SetQuota(v int64) *GetMySpaceResponseBody {
	s.Quota = &v
	return s
}

func (s *GetMySpaceResponseBody) SetSpaceId(v string) *GetMySpaceResponseBody {
	s.SpaceId = &v
	return s
}

func (s *GetMySpaceResponseBody) SetSpaceName(v string) *GetMySpaceResponseBody {
	s.SpaceName = &v
	return s
}

func (s *GetMySpaceResponseBody) SetSpaceType(v string) *GetMySpaceResponseBody {
	s.SpaceType = &v
	return s
}

func (s *GetMySpaceResponseBody) SetUsedQuota(v int64) *GetMySpaceResponseBody {
	s.UsedQuota = &v
	return s
}

type GetMySpaceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMySpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMySpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMySpaceResponse) GoString() string {
	return s.String()
}

func (s *GetMySpaceResponse) SetHeaders(v map[string]*string) *GetMySpaceResponse {
	s.Headers = v
	return s
}

func (s *GetMySpaceResponse) SetStatusCode(v int32) *GetMySpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMySpaceResponse) SetBody(v *GetMySpaceResponseBody) *GetMySpaceResponse {
	s.Body = v
	return s
}

type GetOrgOrWebOpenDocContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOrgOrWebOpenDocContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentHeaders) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentHeaders) SetCommonHeaders(v map[string]*string) *GetOrgOrWebOpenDocContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrgOrWebOpenDocContentHeaders) SetXAcsDingtalkAccessToken(v string) *GetOrgOrWebOpenDocContentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOrgOrWebOpenDocContentRequest struct {
	// example:
	//
	// false
	GenerateCp *bool `json:"generateCp,omitempty" xml:"generateCp,omitempty"`
	// example:
	//
	// 0
	ScopeType *int32 `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// example:
	//
	// markdown
	TargetFormat *string `json:"targetFormat,omitempty" xml:"targetFormat,omitempty"`
}

func (s GetOrgOrWebOpenDocContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentRequest) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentRequest) SetGenerateCp(v bool) *GetOrgOrWebOpenDocContentRequest {
	s.GenerateCp = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentRequest) SetScopeType(v int32) *GetOrgOrWebOpenDocContentRequest {
	s.ScopeType = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentRequest) SetTargetFormat(v string) *GetOrgOrWebOpenDocContentRequest {
	s.TargetFormat = &v
	return s
}

type GetOrgOrWebOpenDocContentResponseBody struct {
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetOrgOrWebOpenDocContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentResponseBody) SetTaskId(v int64) *GetOrgOrWebOpenDocContentResponseBody {
	s.TaskId = &v
	return s
}

type GetOrgOrWebOpenDocContentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrgOrWebOpenDocContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrgOrWebOpenDocContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentResponse) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentResponse) SetHeaders(v map[string]*string) *GetOrgOrWebOpenDocContentResponse {
	s.Headers = v
	return s
}

func (s *GetOrgOrWebOpenDocContentResponse) SetStatusCode(v int32) *GetOrgOrWebOpenDocContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentResponse) SetBody(v *GetOrgOrWebOpenDocContentResponseBody) *GetOrgOrWebOpenDocContentResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// abcd
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
	// example:
	//
	// 0
	Revision *int32 `json:"revision,omitempty" xml:"revision,omitempty"`
	// example:
	//
	// "{\"pageVersion\":\"1.0.0\",\"pageSchema\":{\"version\":\"1.0.0\",\"componentsMap\":[],\"componentsTree\":[]}}"
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// This parameter is required.
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpaceDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// union_id
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStarInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetTaskInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTaskInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTaskInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetTaskInfoHeaders) SetCommonHeaders(v map[string]*string) *GetTaskInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTaskInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetTaskInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTaskInfoResponseBody struct {
	// example:
	//
	// 0
	FailCount *int32 `json:"failCount,omitempty" xml:"failCount,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	SuccCount *int32 `json:"succCount,omitempty" xml:"succCount,omitempty"`
	// example:
	//
	// abcd
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskInfoResponseBody) SetFailCount(v int32) *GetTaskInfoResponseBody {
	s.FailCount = &v
	return s
}

func (s *GetTaskInfoResponseBody) SetStatus(v int32) *GetTaskInfoResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskInfoResponseBody) SetSuccCount(v int32) *GetTaskInfoResponseBody {
	s.SuccCount = &v
	return s
}

func (s *GetTaskInfoResponseBody) SetTaskId(v string) *GetTaskInfoResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetTaskInfoResponseBody) SetTotalCount(v int32) *GetTaskInfoResponseBody {
	s.TotalCount = &v
	return s
}

type GetTaskInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTaskInfoResponse) SetHeaders(v map[string]*string) *GetTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTaskInfoResponse) SetStatusCode(v int32) *GetTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskInfoResponse) SetBody(v *GetTaskInfoResponseBody) *GetTaskInfoResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TeamVO            `json:"body,omitempty" xml:"body,omitempty"`
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
	IncludeFolder *bool `json:"includeFolder,omitempty" xml:"includeFolder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcd
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	SpaceTypes *string `json:"spaceTypes,omitempty" xml:"spaceTypes,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTotalNumberOfDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// abcd
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTotalNumberOfSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserInfoByOpenTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetUuidByDentryIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUuidByDentryIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUuidByDentryIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUuidByDentryIdHeaders) SetCommonHeaders(v map[string]*string) *GetUuidByDentryIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUuidByDentryIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetUuidByDentryIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUuidByDentryIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1L
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetUuidByDentryIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUuidByDentryIdRequest) GoString() string {
	return s.String()
}

func (s *GetUuidByDentryIdRequest) SetOperatorId(v string) *GetUuidByDentryIdRequest {
	s.OperatorId = &v
	return s
}

func (s *GetUuidByDentryIdRequest) SetSpaceId(v string) *GetUuidByDentryIdRequest {
	s.SpaceId = &v
	return s
}

type GetUuidByDentryIdResponseBody struct {
	DentryId   *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	SpaceId    *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetUuidByDentryIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUuidByDentryIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUuidByDentryIdResponseBody) SetDentryId(v string) *GetUuidByDentryIdResponseBody {
	s.DentryId = &v
	return s
}

func (s *GetUuidByDentryIdResponseBody) SetDentryUuid(v string) *GetUuidByDentryIdResponseBody {
	s.DentryUuid = &v
	return s
}

func (s *GetUuidByDentryIdResponseBody) SetSpaceId(v string) *GetUuidByDentryIdResponseBody {
	s.SpaceId = &v
	return s
}

type GetUuidByDentryIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUuidByDentryIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUuidByDentryIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUuidByDentryIdResponse) GoString() string {
	return s.String()
}

func (s *GetUuidByDentryIdResponse) SetHeaders(v map[string]*string) *GetUuidByDentryIdResponse {
	s.Headers = v
	return s
}

func (s *GetUuidByDentryIdResponse) SetStatusCode(v int32) *GetUuidByDentryIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUuidByDentryIdResponse) SetBody(v *GetUuidByDentryIdResponseBody) *GetUuidByDentryIdResponse {
	s.Body = v
	return s
}

type GetWorkspacePermissionScopesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetWorkspacePermissionScopesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacePermissionScopesHeaders) GoString() string {
	return s.String()
}

func (s *GetWorkspacePermissionScopesHeaders) SetCommonHeaders(v map[string]*string) *GetWorkspacePermissionScopesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWorkspacePermissionScopesHeaders) SetXAcsDingtalkAccessToken(v string) *GetWorkspacePermissionScopesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetWorkspacePermissionScopesRequest struct {
	// This parameter is required.
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetWorkspacePermissionScopesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacePermissionScopesRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspacePermissionScopesRequest) SetOperatorId(v string) *GetWorkspacePermissionScopesRequest {
	s.OperatorId = &v
	return s
}

type GetWorkspacePermissionScopesResponseBody struct {
	Members []*GetWorkspacePermissionScopesResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s GetWorkspacePermissionScopesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacePermissionScopesResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspacePermissionScopesResponseBody) SetMembers(v []*GetWorkspacePermissionScopesResponseBodyMembers) *GetWorkspacePermissionScopesResponseBody {
	s.Members = v
	return s
}

type GetWorkspacePermissionScopesResponseBodyMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberRole *string `json:"memberRole,omitempty" xml:"memberRole,omitempty"`
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s GetWorkspacePermissionScopesResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacePermissionScopesResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *GetWorkspacePermissionScopesResponseBodyMembers) SetMemberId(v string) *GetWorkspacePermissionScopesResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *GetWorkspacePermissionScopesResponseBodyMembers) SetMemberRole(v string) *GetWorkspacePermissionScopesResponseBodyMembers {
	s.MemberRole = &v
	return s
}

func (s *GetWorkspacePermissionScopesResponseBodyMembers) SetMemberType(v string) *GetWorkspacePermissionScopesResponseBodyMembers {
	s.MemberType = &v
	return s
}

type GetWorkspacePermissionScopesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspacePermissionScopesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspacePermissionScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacePermissionScopesResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspacePermissionScopesResponse) SetHeaders(v map[string]*string) *GetWorkspacePermissionScopesResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspacePermissionScopesResponse) SetStatusCode(v int32) *GetWorkspacePermissionScopesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspacePermissionScopesResponse) SetBody(v *GetWorkspacePermissionScopesResponseBody) *GetWorkspacePermissionScopesResponse {
	s.Body = v
	return s
}

type HandoverTeamWithoutAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HandoverTeamWithoutAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s HandoverTeamWithoutAuthHeaders) GoString() string {
	return s.String()
}

func (s *HandoverTeamWithoutAuthHeaders) SetCommonHeaders(v map[string]*string) *HandoverTeamWithoutAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HandoverTeamWithoutAuthHeaders) SetXAcsDingtalkAccessToken(v string) *HandoverTeamWithoutAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HandoverTeamWithoutAuthRequest struct {
	// This parameter is required.
	Param *HandoverTeamWithoutAuthRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
}

func (s HandoverTeamWithoutAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s HandoverTeamWithoutAuthRequest) GoString() string {
	return s.String()
}

func (s *HandoverTeamWithoutAuthRequest) SetParam(v *HandoverTeamWithoutAuthRequestParam) *HandoverTeamWithoutAuthRequest {
	s.Param = v
	return s
}

type HandoverTeamWithoutAuthRequestParam struct {
	Leave *bool `json:"leave,omitempty" xml:"leave,omitempty"`
	// This parameter is required.
	NewOwner *string `json:"newOwner,omitempty" xml:"newOwner,omitempty"`
	Notify   *bool   `json:"notify,omitempty" xml:"notify,omitempty"`
	// This parameter is required.
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
}

func (s HandoverTeamWithoutAuthRequestParam) String() string {
	return tea.Prettify(s)
}

func (s HandoverTeamWithoutAuthRequestParam) GoString() string {
	return s.String()
}

func (s *HandoverTeamWithoutAuthRequestParam) SetLeave(v bool) *HandoverTeamWithoutAuthRequestParam {
	s.Leave = &v
	return s
}

func (s *HandoverTeamWithoutAuthRequestParam) SetNewOwner(v string) *HandoverTeamWithoutAuthRequestParam {
	s.NewOwner = &v
	return s
}

func (s *HandoverTeamWithoutAuthRequestParam) SetNotify(v bool) *HandoverTeamWithoutAuthRequestParam {
	s.Notify = &v
	return s
}

func (s *HandoverTeamWithoutAuthRequestParam) SetTeamId(v string) *HandoverTeamWithoutAuthRequestParam {
	s.TeamId = &v
	return s
}

type HandoverTeamWithoutAuthResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HandoverTeamWithoutAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandoverTeamWithoutAuthResponseBody) GoString() string {
	return s.String()
}

func (s *HandoverTeamWithoutAuthResponseBody) SetSuccess(v bool) *HandoverTeamWithoutAuthResponseBody {
	s.Success = &v
	return s
}

type HandoverTeamWithoutAuthResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandoverTeamWithoutAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandoverTeamWithoutAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s HandoverTeamWithoutAuthResponse) GoString() string {
	return s.String()
}

func (s *HandoverTeamWithoutAuthResponse) SetHeaders(v map[string]*string) *HandoverTeamWithoutAuthResponse {
	s.Headers = v
	return s
}

func (s *HandoverTeamWithoutAuthResponse) SetStatusCode(v int32) *HandoverTeamWithoutAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *HandoverTeamWithoutAuthResponse) SetBody(v *HandoverTeamWithoutAuthResponseBody) *HandoverTeamWithoutAuthResponse {
	s.Body = v
	return s
}

type HandoveryWorkspaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HandoveryWorkspaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s HandoveryWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *HandoveryWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *HandoveryWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HandoveryWorkspaceHeaders) SetXAcsDingtalkAccessToken(v string) *HandoveryWorkspaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HandoveryWorkspaceRequest struct {
	// This parameter is required.
	Param *HandoveryWorkspaceRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
}

func (s HandoveryWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s HandoveryWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *HandoveryWorkspaceRequest) SetParam(v *HandoveryWorkspaceRequestParam) *HandoveryWorkspaceRequest {
	s.Param = v
	return s
}

type HandoveryWorkspaceRequestParam struct {
	// example:
	//
	// unionId
	Leave *bool `json:"leave,omitempty" xml:"leave,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// unionId
	ReceiverUnionId *string `json:"receiverUnionId,omitempty" xml:"receiverUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root_node_uuid
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
}

func (s HandoveryWorkspaceRequestParam) String() string {
	return tea.Prettify(s)
}

func (s HandoveryWorkspaceRequestParam) GoString() string {
	return s.String()
}

func (s *HandoveryWorkspaceRequestParam) SetLeave(v bool) *HandoveryWorkspaceRequestParam {
	s.Leave = &v
	return s
}

func (s *HandoveryWorkspaceRequestParam) SetReceiverUnionId(v string) *HandoveryWorkspaceRequestParam {
	s.ReceiverUnionId = &v
	return s
}

func (s *HandoveryWorkspaceRequestParam) SetResourceId(v string) *HandoveryWorkspaceRequestParam {
	s.ResourceId = &v
	return s
}

type HandoveryWorkspaceResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s HandoveryWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandoveryWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *HandoveryWorkspaceResponseBody) SetResult(v bool) *HandoveryWorkspaceResponseBody {
	s.Result = &v
	return s
}

type HandoveryWorkspaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandoveryWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandoveryWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s HandoveryWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *HandoveryWorkspaceResponse) SetHeaders(v map[string]*string) *HandoveryWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *HandoveryWorkspaceResponse) SetStatusCode(v int32) *HandoveryWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *HandoveryWorkspaceResponse) SetBody(v *HandoveryWorkspaceResponseBody) *HandoveryWorkspaceResponse {
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
	ExcludeFile *bool `json:"excludeFile,omitempty" xml:"excludeFile,omitempty"`
	// This parameter is required.
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcd
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
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
	HasMore *bool                         `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Items   []*ListFeedsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// abcdef
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// example:
	//
	// "{}"
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 12340000
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeedsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// abcd
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotDocsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Option *ListPinSpacesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
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
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// true
	WithSpaceCreatorInfo *bool `json:"withSpaceCreatorInfo,omitempty" xml:"withSpaceCreatorInfo,omitempty"`
	// example:
	//
	// true
	WithSpaceModifierInfo *bool `json:"withSpaceModifierInfo,omitempty" xml:"withSpaceModifierInfo,omitempty"`
	// example:
	//
	// true
	WithSpacePermissionRole *bool `json:"withSpacePermissionRole,omitempty" xml:"withSpacePermissionRole,omitempty"`
	// example:
	//
	// true
	WithTeamDetail *bool `json:"withTeamDetail,omitempty" xml:"withTeamDetail,omitempty"`
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
	// example:
	//
	// next_token
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
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string                                        `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	SpaceInfo    *ListPinSpacesResponseBodyResultItemsSpaceInfo `json:"spaceInfo,omitempty" xml:"spaceInfo,omitempty" type:"Struct"`
	// example:
	//
	// NO_PERMISSION
	SpacePermissionRole *string                                       `json:"spacePermissionRole,omitempty" xml:"spacePermissionRole,omitempty"`
	TeamInfo            *ListPinSpacesResponseBodyResultItemsTeamInfo `json:"teamInfo,omitempty" xml:"teamInfo,omitempty" type:"Struct"`
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
	// example:
	//
	// space_cover
	Cover *string `json:"cover,omitempty" xml:"cover,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string                                               `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator    *ListPinSpacesResponseBodyResultItemsSpaceInfoCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// example:
	//
	// description
	Description *string                                              `json:"description,omitempty" xml:"description,omitempty"`
	IconVO      *ListPinSpacesResponseBodyResultItemsSpaceInfoIconVO `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	// example:
	//
	// space_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// mobile_url
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string                                                `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Modifier     *ListPinSpacesResponseBodyResultItemsSpaceInfoModifier `json:"modifier,omitempty" xml:"modifier,omitempty" type:"Struct"`
	// example:
	//
	// space_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// pc_url
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
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
	// example:
	//
	// user_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// union_id
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
	// example:
	//
	// icon_url
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// URL
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
	// example:
	//
	// user_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// union_id
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
	// example:
	//
	// team_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// team_name
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPinSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListRecentsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListRecentsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListRecentsHeaders) GoString() string {
	return s.String()
}

func (s *ListRecentsHeaders) SetCommonHeaders(v map[string]*string) *ListRecentsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRecentsHeaders) SetXAcsDingtalkAccessToken(v string) *ListRecentsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListRecentsRequest struct {
	// This parameter is required.
	Param *ListRecentsRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
}

func (s ListRecentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecentsRequest) GoString() string {
	return s.String()
}

func (s *ListRecentsRequest) SetParam(v *ListRecentsRequestParam) *ListRecentsRequest {
	s.Param = v
	return s
}

type ListRecentsRequestParam struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	FileTypes []*int32 `json:"fileTypes,omitempty" xml:"fileTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// nextToken
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	OperateTypes []*int32 `json:"operateTypes,omitempty" xml:"operateTypes,omitempty" type:"Repeated"`
}

func (s ListRecentsRequestParam) String() string {
	return tea.Prettify(s)
}

func (s ListRecentsRequestParam) GoString() string {
	return s.String()
}

func (s *ListRecentsRequestParam) SetFileTypes(v []*int32) *ListRecentsRequestParam {
	s.FileTypes = v
	return s
}

func (s *ListRecentsRequestParam) SetMaxResults(v int32) *ListRecentsRequestParam {
	s.MaxResults = &v
	return s
}

func (s *ListRecentsRequestParam) SetNextToken(v string) *ListRecentsRequestParam {
	s.NextToken = &v
	return s
}

func (s *ListRecentsRequestParam) SetOperateTypes(v []*int32) *ListRecentsRequestParam {
	s.OperateTypes = v
	return s
}

type ListRecentsResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// nextToken
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// recentDentryList
	RecentDentryList []*ListRecentsResponseBodyRecentDentryList `json:"recentDentryList,omitempty" xml:"recentDentryList,omitempty" type:"Repeated"`
}

func (s ListRecentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentsResponseBody) SetHasMore(v bool) *ListRecentsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListRecentsResponseBody) SetNextToken(v string) *ListRecentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecentsResponseBody) SetRecentDentryList(v []*ListRecentsResponseBodyRecentDentryList) *ListRecentsResponseBody {
	s.RecentDentryList = v
	return s
}

type ListRecentsResponseBodyRecentDentryList struct {
	// example:
	//
	// 123
	AccessTime *int64 `json:"accessTime,omitempty" xml:"accessTime,omitempty"`
	// example:
	//
	// true
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// example:
	//
	// https://example.com
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 0
	OperateType *int32 `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// example:
	//
	// resource
	Resource *ListRecentsResponseBodyRecentDentryListResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
}

func (s ListRecentsResponseBodyRecentDentryList) String() string {
	return tea.Prettify(s)
}

func (s ListRecentsResponseBodyRecentDentryList) GoString() string {
	return s.String()
}

func (s *ListRecentsResponseBodyRecentDentryList) SetAccessTime(v int64) *ListRecentsResponseBodyRecentDentryList {
	s.AccessTime = &v
	return s
}

func (s *ListRecentsResponseBodyRecentDentryList) SetDeleted(v bool) *ListRecentsResponseBodyRecentDentryList {
	s.Deleted = &v
	return s
}

func (s *ListRecentsResponseBodyRecentDentryList) SetIcon(v string) *ListRecentsResponseBodyRecentDentryList {
	s.Icon = &v
	return s
}

func (s *ListRecentsResponseBodyRecentDentryList) SetOperateType(v int32) *ListRecentsResponseBodyRecentDentryList {
	s.OperateType = &v
	return s
}

func (s *ListRecentsResponseBodyRecentDentryList) SetResource(v *ListRecentsResponseBodyRecentDentryListResource) *ListRecentsResponseBodyRecentDentryList {
	s.Resource = v
	return s
}

type ListRecentsResponseBodyRecentDentryListResource struct {
	// example:
	//
	// dentryUuid
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// example:
	//
	// driveDentryId
	DriveDentryId *string `json:"driveDentryId,omitempty" xml:"driveDentryId,omitempty"`
	// example:
	//
	// driveSpaceId
	DriveSpaceId *string `json:"driveSpaceId,omitempty" xml:"driveSpaceId,omitempty"`
	// example:
	//
	// docx
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// spaceInfo
	SpaceInfo *ListRecentsResponseBodyRecentDentryListResourceSpaceInfo `json:"spaceInfo,omitempty" xml:"spaceInfo,omitempty" type:"Struct"`
	// example:
	//
	// https://example.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListRecentsResponseBodyRecentDentryListResource) String() string {
	return tea.Prettify(s)
}

func (s ListRecentsResponseBodyRecentDentryListResource) GoString() string {
	return s.String()
}

func (s *ListRecentsResponseBodyRecentDentryListResource) SetDentryUuid(v string) *ListRecentsResponseBodyRecentDentryListResource {
	s.DentryUuid = &v
	return s
}

func (s *ListRecentsResponseBodyRecentDentryListResource) SetDriveDentryId(v string) *ListRecentsResponseBodyRecentDentryListResource {
	s.DriveDentryId = &v
	return s
}

func (s *ListRecentsResponseBodyRecentDentryListResource) SetDriveSpaceId(v string) *ListRecentsResponseBodyRecentDentryListResource {
	s.DriveSpaceId = &v
	return s
}

func (s *ListRecentsResponseBodyRecentDentryListResource) SetExtension(v string) *ListRecentsResponseBodyRecentDentryListResource {
	s.Extension = &v
	return s
}

func (s *ListRecentsResponseBodyRecentDentryListResource) SetName(v string) *ListRecentsResponseBodyRecentDentryListResource {
	s.Name = &v
	return s
}

func (s *ListRecentsResponseBodyRecentDentryListResource) SetSpaceInfo(v *ListRecentsResponseBodyRecentDentryListResourceSpaceInfo) *ListRecentsResponseBodyRecentDentryListResource {
	s.SpaceInfo = v
	return s
}

func (s *ListRecentsResponseBodyRecentDentryListResource) SetUrl(v string) *ListRecentsResponseBodyRecentDentryListResource {
	s.Url = &v
	return s
}

type ListRecentsResponseBodyRecentDentryListResourceSpaceInfo struct {
	// example:
	//
	// im
	SceneType *string `json:"sceneType,omitempty" xml:"sceneType,omitempty"`
}

func (s ListRecentsResponseBodyRecentDentryListResourceSpaceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListRecentsResponseBodyRecentDentryListResourceSpaceInfo) GoString() string {
	return s.String()
}

func (s *ListRecentsResponseBodyRecentDentryListResourceSpaceInfo) SetSceneType(v string) *ListRecentsResponseBodyRecentDentryListResourceSpaceInfo {
	s.SceneType = &v
	return s
}

type ListRecentsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecentsResponse) GoString() string {
	return s.String()
}

func (s *ListRecentsResponse) SetHeaders(v map[string]*string) *ListRecentsResponse {
	s.Headers = v
	return s
}

func (s *ListRecentsResponse) SetStatusCode(v int32) *ListRecentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentsResponse) SetBody(v *ListRecentsResponseBody) *ListRecentsResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// abcd
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRelatedSpaceTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ajYkbc7
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
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
	HasMore *bool        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Items   []*TeamModel `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// cjk72iEakdim
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRelatedTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// abcd
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
	// example:
	//
	// base
	DisplayType *string `json:"displayType,omitempty" xml:"displayType,omitempty"`
	// example:
	//
	// abc
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 测试分组
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	SpaceNum *int32        `json:"spaceNum,omitempty" xml:"spaceNum,omitempty"`
	Spaces   []*SpaceModel `json:"spaces,omitempty" xml:"spaces,omitempty" type:"Repeated"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSpaceSectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Option *ListStarsRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
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
	ContentTypeList []*string `json:"contentTypeList,omitempty" xml:"contentTypeList,omitempty" type:"Repeated"`
	FilterDocTypes  []*string `json:"filterDocTypes,omitempty" xml:"filterDocTypes,omitempty" type:"Repeated"`
	// example:
	//
	// true
	ListV2 *bool `json:"listV2,omitempty" xml:"listV2,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// createTime
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// true
	WithDentryCreatorInfo *bool `json:"withDentryCreatorInfo,omitempty" xml:"withDentryCreatorInfo,omitempty"`
	// example:
	//
	// true
	WithDentryModifierInfo *bool `json:"withDentryModifierInfo,omitempty" xml:"withDentryModifierInfo,omitempty"`
	// example:
	//
	// true
	WithDentryPermissionRole *bool `json:"withDentryPermissionRole,omitempty" xml:"withDentryPermissionRole,omitempty"`
	// example:
	//
	// true
	WithSpaceDetail *bool `json:"withSpaceDetail,omitempty" xml:"withSpaceDetail,omitempty"`
	// example:
	//
	// true
	WithSpacePermissionRole *bool `json:"withSpacePermissionRole,omitempty" xml:"withSpacePermissionRole,omitempty"`
	// example:
	//
	// true
	WithTeamDetail *bool `json:"withTeamDetail,omitempty" xml:"withTeamDetail,omitempty"`
}

func (s ListStarsRequestOption) String() string {
	return tea.Prettify(s)
}

func (s ListStarsRequestOption) GoString() string {
	return s.String()
}

func (s *ListStarsRequestOption) SetContentTypeList(v []*string) *ListStarsRequestOption {
	s.ContentTypeList = v
	return s
}

func (s *ListStarsRequestOption) SetFilterDocTypes(v []*string) *ListStarsRequestOption {
	s.FilterDocTypes = v
	return s
}

func (s *ListStarsRequestOption) SetListV2(v bool) *ListStarsRequestOption {
	s.ListV2 = &v
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
	// example:
	//
	// next_token
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
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string                                  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DentryInfo *ListStarsResponseBodyStarListDentryInfo `json:"dentryInfo,omitempty" xml:"dentryInfo,omitempty" type:"Struct"`
	// example:
	//
	// NO_PERMISSION
	DentryPermissionRole *string `json:"dentryPermissionRole,omitempty" xml:"dentryPermissionRole,omitempty"`
	// example:
	//
	// star_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string                                 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	SpaceInfo    *ListStarsResponseBodyStarListSpaceInfo `json:"spaceInfo,omitempty" xml:"spaceInfo,omitempty" type:"Struct"`
	// example:
	//
	// NO_PERMISSION
	SpacePermissionRole *string                                `json:"spacePermissionRole,omitempty" xml:"spacePermissionRole,omitempty"`
	StarType            *string                                `json:"starType,omitempty" xml:"starType,omitempty"`
	TeamInfo            *ListStarsResponseBodyStarListTeamInfo `json:"teamInfo,omitempty" xml:"teamInfo,omitempty" type:"Struct"`
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
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string                                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator    *ListStarsResponseBodyStarListDentryInfoCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// mobile_url
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string                                          `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Modifier     *ListStarsResponseBodyStarListDentryInfoModifier `json:"modifier,omitempty" xml:"modifier,omitempty" type:"Struct"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// pc_url
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// FILE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
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
	// example:
	//
	// user_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// union_id
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
	// example:
	//
	// user_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// union_id
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
	// example:
	//
	// space_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// space_name
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
	// example:
	//
	// team_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// team_name
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStarsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
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
	Members []*ListTeamMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// example:
	//
	// 测试小组名称
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
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
	// example:
	//
	// YEp3JcM******UIbhwiE
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// example:
	//
	// 2
	MemberType *int32 `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// 小钉
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTeamMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// union_id
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
	// example:
	//
	// true
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MarkStarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DentryVO          `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// union_id
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
	// example:
	//
	// true
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PinSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// false
	IncludeSpace *bool `json:"includeSpace,omitempty" xml:"includeSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcd
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DentryVO          `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryDocContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDocContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDocContentHeaders) GoString() string {
	return s.String()
}

func (s *QueryDocContentHeaders) SetCommonHeaders(v map[string]*string) *QueryDocContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDocContentHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDocContentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDocContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// markdown
	TargetFormat *string `json:"targetFormat,omitempty" xml:"targetFormat,omitempty"`
}

func (s QueryDocContentRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDocContentRequest) GoString() string {
	return s.String()
}

func (s *QueryDocContentRequest) SetOperatorId(v string) *QueryDocContentRequest {
	s.OperatorId = &v
	return s
}

func (s *QueryDocContentRequest) SetTargetFormat(v string) *QueryDocContentRequest {
	s.TargetFormat = &v
	return s
}

type QueryDocContentResponseBody struct {
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryDocContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDocContentResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDocContentResponseBody) SetTaskId(v int64) *QueryDocContentResponseBody {
	s.TaskId = &v
	return s
}

type QueryDocContentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDocContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDocContentResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDocContentResponse) GoString() string {
	return s.String()
}

func (s *QueryDocContentResponse) SetHeaders(v map[string]*string) *QueryDocContentResponse {
	s.Headers = v
	return s
}

func (s *QueryDocContentResponse) SetStatusCode(v int32) *QueryDocContentResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDocContentResponse) SetBody(v *QueryDocContentResponseBody) *QueryDocContentResponse {
	s.Body = v
	return s
}

type QueryGetContentJobHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGetContentJobHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGetContentJobHeaders) GoString() string {
	return s.String()
}

func (s *QueryGetContentJobHeaders) SetCommonHeaders(v map[string]*string) *QueryGetContentJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGetContentJobHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGetContentJobHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGetContentJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// markdown
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryGetContentJobRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGetContentJobRequest) GoString() string {
	return s.String()
}

func (s *QueryGetContentJobRequest) SetOperatorId(v string) *QueryGetContentJobRequest {
	s.OperatorId = &v
	return s
}

func (s *QueryGetContentJobRequest) SetTaskId(v int64) *QueryGetContentJobRequest {
	s.TaskId = &v
	return s
}

type QueryGetContentJobResponseBody struct {
	ContentKey *string `json:"contentKey,omitempty" xml:"contentKey,omitempty"`
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryGetContentJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGetContentJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGetContentJobResponseBody) SetContentKey(v string) *QueryGetContentJobResponseBody {
	s.ContentKey = &v
	return s
}

func (s *QueryGetContentJobResponseBody) SetStatus(v int32) *QueryGetContentJobResponseBody {
	s.Status = &v
	return s
}

type QueryGetContentJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGetContentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGetContentJobResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGetContentJobResponse) GoString() string {
	return s.String()
}

func (s *QueryGetContentJobResponse) SetHeaders(v map[string]*string) *QueryGetContentJobResponse {
	s.Headers = v
	return s
}

func (s *QueryGetContentJobResponse) SetStatusCode(v int32) *QueryGetContentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGetContentJobResponse) SetBody(v *QueryGetContentJobResponseBody) *QueryGetContentJobResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://alidocs.dingtalk.com/i/nodes/m0Xw6OYE4D7VLeaBP***
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
	// example:
	//
	// mainsite
	BizType *string      `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Dentry  *DentryModel `json:"dentry,omitempty" xml:"dentry,omitempty"`
	// example:
	//
	// file
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
	// example:
	//
	// 这是知识库简介
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// YRBG******vJXDAr
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 这是知识库名称
	Name  *string                               `json:"name,omitempty" xml:"name,omitempty"`
	Owner *QueryItemByUrlResponseBodySpaceOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
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
	// example:
	//
	// 小钉
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// SgExXM*******L0tAiEiE
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryItemByUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SpaceVO           `json:"body,omitempty" xml:"body,omitempty"`
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
	CreatorType *int32 `json:"creatorType,omitempty" xml:"creatorType,omitempty"`
	FileType    *int32 `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// This parameter is required.
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
	RecentType *int32 `json:"recentType,omitempty" xml:"recentType,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// abcd
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SpaceVO           `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcd
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// abc
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
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
	// example:
	//
	// https://img.abc.yyy
	Cover *string `json:"cover,omitempty" xml:"cover,omitempty"`
	// example:
	//
	// This is some description.
	Description *string                                 `json:"description,omitempty" xml:"description,omitempty"`
	HdIconVO    *RelatedSpacesResponseBodyItemsHdIconVO `json:"hdIconVO,omitempty" xml:"hdIconVO,omitempty" type:"Struct"`
	IconVO      *RelatedSpacesResponseBodyItemsIconVO   `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hello
	Name       *string                              `json:"name,omitempty" xml:"name,omitempty"`
	Owner      *RelatedSpacesResponseBodyItemsOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	RecentList []*DentryModel                       `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// https://xx.yy
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
	// This parameter is required.
	//
	// example:
	//
	// https://img.xxx.yyy
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// url
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
	// This parameter is required.
	//
	// example:
	//
	// https://img.xxx.yyy
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// url
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
	// This parameter is required.
	//
	// example:
	//
	// dingtalk
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcd
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
	// example:
	//
	// 3
	RoleCode     *string   `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	SpaceActions []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RelatedSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	Members []*RemoveTeamMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	Notify  *bool                              `json:"notify,omitempty" xml:"notify,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	MemberType *int32 `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// 1
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
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
	// example:
	//
	// YEp3JcM******UIbhwiE
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// example:
	//
	// 2
	MemberType *int32 `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// 小钉
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTeamMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DentryVO          `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	Members []*SaveTeamMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	Notify  *bool                            `json:"notify,omitempty" xml:"notify,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	MemberType *int32 `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// 1
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
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
	// example:
	//
	// YEp3JcM******UIbhwiE
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// example:
	//
	// 2
	MemberType *int32 `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// 小钉
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTeamMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 测试搜索关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// This parameter is required.
	OperatorId   *string                    `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	SpaceRequest *SearchRequestSpaceRequest `json:"spaceRequest,omitempty" xml:"spaceRequest,omitempty" type:"Struct"`
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
	CreateTimeRange *SearchRequestDentryRequestCreateTimeRange `json:"createTimeRange,omitempty" xml:"createTimeRange,omitempty" type:"Struct"`
	CreateUsers     []*string                                  `json:"createUsers,omitempty" xml:"createUsers,omitempty" type:"Repeated"`
	Editors         []*string                                  `json:"editors,omitempty" xml:"editors,omitempty" type:"Repeated"`
	// This parameter is required.
	MaxResults     *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SearchField    *int32    `json:"searchField,omitempty" xml:"searchField,omitempty"`
	SearchFileType *int32    `json:"searchFileType,omitempty" xml:"searchFileType,omitempty"`
	SpaceId        *string   `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	SpaceIds       []*string `json:"spaceIds,omitempty" xml:"spaceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 40
	SummaryLength  *int32                                    `json:"summaryLength,omitempty" xml:"summaryLength,omitempty"`
	VisitTimeRange *SearchRequestDentryRequestVisitTimeRange `json:"visitTimeRange,omitempty" xml:"visitTimeRange,omitempty" type:"Struct"`
}

func (s SearchRequestDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRequestDentryRequest) GoString() string {
	return s.String()
}

func (s *SearchRequestDentryRequest) SetCreateTimeRange(v *SearchRequestDentryRequestCreateTimeRange) *SearchRequestDentryRequest {
	s.CreateTimeRange = v
	return s
}

func (s *SearchRequestDentryRequest) SetCreateUsers(v []*string) *SearchRequestDentryRequest {
	s.CreateUsers = v
	return s
}

func (s *SearchRequestDentryRequest) SetEditors(v []*string) *SearchRequestDentryRequest {
	s.Editors = v
	return s
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

func (s *SearchRequestDentryRequest) SetSpaceIds(v []*string) *SearchRequestDentryRequest {
	s.SpaceIds = v
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

type SearchRequestDentryRequestCreateTimeRange struct {
	End   *int64 `json:"end,omitempty" xml:"end,omitempty"`
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s SearchRequestDentryRequestCreateTimeRange) String() string {
	return tea.Prettify(s)
}

func (s SearchRequestDentryRequestCreateTimeRange) GoString() string {
	return s.String()
}

func (s *SearchRequestDentryRequestCreateTimeRange) SetEnd(v int64) *SearchRequestDentryRequestCreateTimeRange {
	s.End = &v
	return s
}

func (s *SearchRequestDentryRequestCreateTimeRange) SetStart(v int64) *SearchRequestDentryRequestCreateTimeRange {
	s.Start = &v
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
	// This parameter is required.
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
	Content     *string          `json:"content,omitempty" xml:"content,omitempty"`
	Creation    *OpenActionModel `json:"creation,omitempty" xml:"creation,omitempty"`
	DentryId    *string          `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	DentryUuid  *string          `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	Extension   *string          `json:"extension,omitempty" xml:"extension,omitempty"`
	IconUrl     *string          `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	LastEdition *OpenActionModel `json:"lastEdition,omitempty" xml:"lastEdition,omitempty"`
	Name        *string          `json:"name,omitempty" xml:"name,omitempty"`
	OriginName  *string          `json:"originName,omitempty" xml:"originName,omitempty"`
	Path        *string          `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 1
	SceneType      *int32  `json:"sceneType,omitempty" xml:"sceneType,omitempty"`
	SearchFileType *int32  `json:"searchFileType,omitempty" xml:"searchFileType,omitempty"`
	SpaceId        *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	ThumbnailUrl   *string `json:"thumbnailUrl,omitempty" xml:"thumbnailUrl,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
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
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type SearchTemplatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchTemplatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchTemplatesHeaders) GoString() string {
	return s.String()
}

func (s *SearchTemplatesHeaders) SetCommonHeaders(v map[string]*string) *SearchTemplatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchTemplatesHeaders) SetXAcsDingtalkAccessToken(v string) *SearchTemplatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchTemplatesRequest struct {
	Option *SearchTemplatesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	Param *SearchTemplatesRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s SearchTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTemplatesRequest) GoString() string {
	return s.String()
}

func (s *SearchTemplatesRequest) SetOption(v *SearchTemplatesRequestOption) *SearchTemplatesRequest {
	s.Option = v
	return s
}

func (s *SearchTemplatesRequest) SetParam(v *SearchTemplatesRequestParam) *SearchTemplatesRequest {
	s.Param = v
	return s
}

func (s *SearchTemplatesRequest) SetOperatorId(v string) *SearchTemplatesRequest {
	s.OperatorId = &v
	return s
}

type SearchTemplatesRequestOption struct {
	ExcludeWorkspaceIds []*string `json:"excludeWorkspaceIds,omitempty" xml:"excludeWorkspaceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// pc
	Platform      *string  `json:"platform,omitempty" xml:"platform,omitempty"`
	TemplateTypes []*int32 `json:"templateTypes,omitempty" xml:"templateTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Version      *int32    `json:"version,omitempty" xml:"version,omitempty"`
	WorkspaceIds []*string `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty" type:"Repeated"`
}

func (s SearchTemplatesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s SearchTemplatesRequestOption) GoString() string {
	return s.String()
}

func (s *SearchTemplatesRequestOption) SetExcludeWorkspaceIds(v []*string) *SearchTemplatesRequestOption {
	s.ExcludeWorkspaceIds = v
	return s
}

func (s *SearchTemplatesRequestOption) SetMaxResults(v int32) *SearchTemplatesRequestOption {
	s.MaxResults = &v
	return s
}

func (s *SearchTemplatesRequestOption) SetNextToken(v string) *SearchTemplatesRequestOption {
	s.NextToken = &v
	return s
}

func (s *SearchTemplatesRequestOption) SetPlatform(v string) *SearchTemplatesRequestOption {
	s.Platform = &v
	return s
}

func (s *SearchTemplatesRequestOption) SetTemplateTypes(v []*int32) *SearchTemplatesRequestOption {
	s.TemplateTypes = v
	return s
}

func (s *SearchTemplatesRequestOption) SetVersion(v int32) *SearchTemplatesRequestOption {
	s.Version = &v
	return s
}

func (s *SearchTemplatesRequestOption) SetWorkspaceIds(v []*string) *SearchTemplatesRequestOption {
	s.WorkspaceIds = v
	return s
}

type SearchTemplatesRequestParam struct {
	// This parameter is required.
	//
	// example:
	//
	// user_template
	Belong *string `json:"belong,omitempty" xml:"belong,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// searchName
	SearchName *string `json:"searchName,omitempty" xml:"searchName,omitempty"`
}

func (s SearchTemplatesRequestParam) String() string {
	return tea.Prettify(s)
}

func (s SearchTemplatesRequestParam) GoString() string {
	return s.String()
}

func (s *SearchTemplatesRequestParam) SetBelong(v string) *SearchTemplatesRequestParam {
	s.Belong = &v
	return s
}

func (s *SearchTemplatesRequestParam) SetSearchName(v string) *SearchTemplatesRequestParam {
	s.SearchName = &v
	return s
}

type SearchTemplatesResponseBody struct {
	// example:
	//
	// next_token
	NextToken    *string                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TemplateList []*SearchTemplatesResponseBodyTemplateList `json:"templateList,omitempty" xml:"templateList,omitempty" type:"Repeated"`
}

func (s SearchTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTemplatesResponseBody) SetNextToken(v string) *SearchTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchTemplatesResponseBody) SetTemplateList(v []*SearchTemplatesResponseBodyTemplateList) *SearchTemplatesResponseBody {
	s.TemplateList = v
	return s
}

type SearchTemplatesResponseBodyTemplateList struct {
	AuthorName         *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	Belong             *string `json:"belong,omitempty" xml:"belong,omitempty"`
	ContentDownloadUrl *string `json:"contentDownloadUrl,omitempty" xml:"contentDownloadUrl,omitempty"`
	CoverDownloadUrl   *string `json:"coverDownloadUrl,omitempty" xml:"coverDownloadUrl,omitempty"`
	CreateTime         *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	ModifiedTime       *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	TemplateId         *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
	Type               *int32  `json:"type,omitempty" xml:"type,omitempty"`
	UsedCount          *int64  `json:"usedCount,omitempty" xml:"usedCount,omitempty"`
	WorkspaceId        *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	WorkspaceName      *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s SearchTemplatesResponseBodyTemplateList) String() string {
	return tea.Prettify(s)
}

func (s SearchTemplatesResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *SearchTemplatesResponseBodyTemplateList) SetAuthorName(v string) *SearchTemplatesResponseBodyTemplateList {
	s.AuthorName = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetBelong(v string) *SearchTemplatesResponseBodyTemplateList {
	s.Belong = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetContentDownloadUrl(v string) *SearchTemplatesResponseBodyTemplateList {
	s.ContentDownloadUrl = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetCoverDownloadUrl(v string) *SearchTemplatesResponseBodyTemplateList {
	s.CoverDownloadUrl = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetCreateTime(v string) *SearchTemplatesResponseBodyTemplateList {
	s.CreateTime = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetDescription(v string) *SearchTemplatesResponseBodyTemplateList {
	s.Description = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetModifiedTime(v string) *SearchTemplatesResponseBodyTemplateList {
	s.ModifiedTime = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetTemplateId(v string) *SearchTemplatesResponseBodyTemplateList {
	s.TemplateId = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetTitle(v string) *SearchTemplatesResponseBodyTemplateList {
	s.Title = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetType(v int32) *SearchTemplatesResponseBodyTemplateList {
	s.Type = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetUsedCount(v int64) *SearchTemplatesResponseBodyTemplateList {
	s.UsedCount = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetWorkspaceId(v string) *SearchTemplatesResponseBodyTemplateList {
	s.WorkspaceId = &v
	return s
}

func (s *SearchTemplatesResponseBodyTemplateList) SetWorkspaceName(v string) *SearchTemplatesResponseBodyTemplateList {
	s.WorkspaceName = &v
	return s
}

type SearchTemplatesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTemplatesResponse) GoString() string {
	return s.String()
}

func (s *SearchTemplatesResponse) SetHeaders(v map[string]*string) *SearchTemplatesResponse {
	s.Headers = v
	return s
}

func (s *SearchTemplatesResponse) SetStatusCode(v int32) *SearchTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTemplatesResponse) SetBody(v *SearchTemplatesResponseBody) *SearchTemplatesResponse {
	s.Body = v
	return s
}

type ShareUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ShareUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s ShareUrlHeaders) GoString() string {
	return s.String()
}

func (s *ShareUrlHeaders) SetCommonHeaders(v map[string]*string) *ShareUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ShareUrlHeaders) SetXAcsDingtalkAccessToken(v string) *ShareUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ShareUrlRequest struct {
	// This parameter is required.
	Param *ShareUrlRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
}

func (s ShareUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s ShareUrlRequest) GoString() string {
	return s.String()
}

func (s *ShareUrlRequest) SetParam(v *ShareUrlRequestParam) *ShareUrlRequest {
	s.Param = v
	return s
}

type ShareUrlRequestParam struct {
	// This parameter is required.
	//
	// example:
	//
	// dentryUuid
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// example:
	//
	// true
	TriggerShare *bool `json:"triggerShare,omitempty" xml:"triggerShare,omitempty"`
}

func (s ShareUrlRequestParam) String() string {
	return tea.Prettify(s)
}

func (s ShareUrlRequestParam) GoString() string {
	return s.String()
}

func (s *ShareUrlRequestParam) SetDentryUuid(v string) *ShareUrlRequestParam {
	s.DentryUuid = &v
	return s
}

func (s *ShareUrlRequestParam) SetTriggerShare(v bool) *ShareUrlRequestParam {
	s.TriggerShare = &v
	return s
}

type ShareUrlResponseBody struct {
	// example:
	//
	// shareUrlInfo
	ShareUrlInfo *ShareUrlResponseBodyShareUrlInfo `json:"shareUrlInfo,omitempty" xml:"shareUrlInfo,omitempty" type:"Struct"`
}

func (s ShareUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ShareUrlResponseBody) GoString() string {
	return s.String()
}

func (s *ShareUrlResponseBody) SetShareUrlInfo(v *ShareUrlResponseBodyShareUrlInfo) *ShareUrlResponseBody {
	s.ShareUrlInfo = v
	return s
}

type ShareUrlResponseBodyShareUrlInfo struct {
	// example:
	//
	// http://example.com
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// http://example.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s ShareUrlResponseBodyShareUrlInfo) String() string {
	return tea.Prettify(s)
}

func (s ShareUrlResponseBodyShareUrlInfo) GoString() string {
	return s.String()
}

func (s *ShareUrlResponseBodyShareUrlInfo) SetMobileUrl(v string) *ShareUrlResponseBodyShareUrlInfo {
	s.MobileUrl = &v
	return s
}

func (s *ShareUrlResponseBodyShareUrlInfo) SetPcUrl(v string) *ShareUrlResponseBodyShareUrlInfo {
	s.PcUrl = &v
	return s
}

type ShareUrlResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShareUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShareUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s ShareUrlResponse) GoString() string {
	return s.String()
}

func (s *ShareUrlResponse) SetHeaders(v map[string]*string) *ShareUrlResponse {
	s.Headers = v
	return s
}

func (s *ShareUrlResponse) SetStatusCode(v int32) *ShareUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *ShareUrlResponse) SetBody(v *ShareUrlResponseBody) *ShareUrlResponse {
	s.Body = v
	return s
}

type SubmitGetContentJobHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SubmitGetContentJobHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubmitGetContentJobHeaders) GoString() string {
	return s.String()
}

func (s *SubmitGetContentJobHeaders) SetCommonHeaders(v map[string]*string) *SubmitGetContentJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubmitGetContentJobHeaders) SetXAcsDingtalkAccessToken(v string) *SubmitGetContentJobHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SubmitGetContentJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// markdown
	TargetFormat *string `json:"targetFormat,omitempty" xml:"targetFormat,omitempty"`
}

func (s SubmitGetContentJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitGetContentJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitGetContentJobRequest) SetOperatorId(v string) *SubmitGetContentJobRequest {
	s.OperatorId = &v
	return s
}

func (s *SubmitGetContentJobRequest) SetTargetFormat(v string) *SubmitGetContentJobRequest {
	s.TargetFormat = &v
	return s
}

type SubmitGetContentJobResponseBody struct {
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitGetContentJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitGetContentJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitGetContentJobResponseBody) SetTaskId(v int64) *SubmitGetContentJobResponseBody {
	s.TaskId = &v
	return s
}

type SubmitGetContentJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitGetContentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitGetContentJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitGetContentJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitGetContentJobResponse) SetHeaders(v map[string]*string) *SubmitGetContentJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitGetContentJobResponse) SetStatusCode(v int32) *SubmitGetContentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitGetContentJobResponse) SetBody(v *SubmitGetContentJobResponseBody) *SubmitGetContentJobResponse {
	s.Body = v
	return s
}

type TeamTemplatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TeamTemplatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s TeamTemplatesHeaders) GoString() string {
	return s.String()
}

func (s *TeamTemplatesHeaders) SetCommonHeaders(v map[string]*string) *TeamTemplatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TeamTemplatesHeaders) SetXAcsDingtalkAccessToken(v string) *TeamTemplatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TeamTemplatesRequest struct {
	Option *TeamTemplatesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s TeamTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s TeamTemplatesRequest) GoString() string {
	return s.String()
}

func (s *TeamTemplatesRequest) SetOption(v *TeamTemplatesRequestOption) *TeamTemplatesRequest {
	s.Option = v
	return s
}

func (s *TeamTemplatesRequest) SetOperatorId(v string) *TeamTemplatesRequest {
	s.OperatorId = &v
	return s
}

type TeamTemplatesRequestOption struct {
	ExcludeWorkspaceIds []*string `json:"excludeWorkspaceIds,omitempty" xml:"excludeWorkspaceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// pc
	Platform      *string  `json:"platform,omitempty" xml:"platform,omitempty"`
	TemplateTypes []*int32 `json:"templateTypes,omitempty" xml:"templateTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Version      *int32    `json:"version,omitempty" xml:"version,omitempty"`
	WorkspaceIds []*string `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty" type:"Repeated"`
}

func (s TeamTemplatesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s TeamTemplatesRequestOption) GoString() string {
	return s.String()
}

func (s *TeamTemplatesRequestOption) SetExcludeWorkspaceIds(v []*string) *TeamTemplatesRequestOption {
	s.ExcludeWorkspaceIds = v
	return s
}

func (s *TeamTemplatesRequestOption) SetMaxResults(v int32) *TeamTemplatesRequestOption {
	s.MaxResults = &v
	return s
}

func (s *TeamTemplatesRequestOption) SetNextToken(v string) *TeamTemplatesRequestOption {
	s.NextToken = &v
	return s
}

func (s *TeamTemplatesRequestOption) SetPlatform(v string) *TeamTemplatesRequestOption {
	s.Platform = &v
	return s
}

func (s *TeamTemplatesRequestOption) SetTemplateTypes(v []*int32) *TeamTemplatesRequestOption {
	s.TemplateTypes = v
	return s
}

func (s *TeamTemplatesRequestOption) SetVersion(v int32) *TeamTemplatesRequestOption {
	s.Version = &v
	return s
}

func (s *TeamTemplatesRequestOption) SetWorkspaceIds(v []*string) *TeamTemplatesRequestOption {
	s.WorkspaceIds = v
	return s
}

type TeamTemplatesResponseBody struct {
	// example:
	//
	// next_token
	NextToken    *string                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TemplateList []*TeamTemplatesResponseBodyTemplateList `json:"templateList,omitempty" xml:"templateList,omitempty" type:"Repeated"`
}

func (s TeamTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TeamTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *TeamTemplatesResponseBody) SetNextToken(v string) *TeamTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *TeamTemplatesResponseBody) SetTemplateList(v []*TeamTemplatesResponseBodyTemplateList) *TeamTemplatesResponseBody {
	s.TemplateList = v
	return s
}

type TeamTemplatesResponseBodyTemplateList struct {
	AuthorName         *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	Belong             *string `json:"belong,omitempty" xml:"belong,omitempty"`
	ContentDownloadUrl *string `json:"contentDownloadUrl,omitempty" xml:"contentDownloadUrl,omitempty"`
	CoverDownloadUrl   *string `json:"coverDownloadUrl,omitempty" xml:"coverDownloadUrl,omitempty"`
	CreateTime         *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	ModifiedTime       *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	TemplateId         *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
	Type               *int32  `json:"type,omitempty" xml:"type,omitempty"`
	UsedCount          *int64  `json:"usedCount,omitempty" xml:"usedCount,omitempty"`
	WorkspaceId        *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	WorkspaceName      *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s TeamTemplatesResponseBodyTemplateList) String() string {
	return tea.Prettify(s)
}

func (s TeamTemplatesResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *TeamTemplatesResponseBodyTemplateList) SetAuthorName(v string) *TeamTemplatesResponseBodyTemplateList {
	s.AuthorName = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetBelong(v string) *TeamTemplatesResponseBodyTemplateList {
	s.Belong = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetContentDownloadUrl(v string) *TeamTemplatesResponseBodyTemplateList {
	s.ContentDownloadUrl = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetCoverDownloadUrl(v string) *TeamTemplatesResponseBodyTemplateList {
	s.CoverDownloadUrl = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetCreateTime(v string) *TeamTemplatesResponseBodyTemplateList {
	s.CreateTime = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetDescription(v string) *TeamTemplatesResponseBodyTemplateList {
	s.Description = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetModifiedTime(v string) *TeamTemplatesResponseBodyTemplateList {
	s.ModifiedTime = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetTemplateId(v string) *TeamTemplatesResponseBodyTemplateList {
	s.TemplateId = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetTitle(v string) *TeamTemplatesResponseBodyTemplateList {
	s.Title = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetType(v int32) *TeamTemplatesResponseBodyTemplateList {
	s.Type = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetUsedCount(v int64) *TeamTemplatesResponseBodyTemplateList {
	s.UsedCount = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetWorkspaceId(v string) *TeamTemplatesResponseBodyTemplateList {
	s.WorkspaceId = &v
	return s
}

func (s *TeamTemplatesResponseBodyTemplateList) SetWorkspaceName(v string) *TeamTemplatesResponseBodyTemplateList {
	s.WorkspaceName = &v
	return s
}

type TeamTemplatesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TeamTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TeamTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s TeamTemplatesResponse) GoString() string {
	return s.String()
}

func (s *TeamTemplatesResponse) SetHeaders(v map[string]*string) *TeamTemplatesResponse {
	s.Headers = v
	return s
}

func (s *TeamTemplatesResponse) SetStatusCode(v int32) *TeamTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *TeamTemplatesResponse) SetBody(v *TeamTemplatesResponseBody) *TeamTemplatesResponse {
	s.Body = v
	return s
}

type TemplateCategoriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TemplateCategoriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s TemplateCategoriesHeaders) GoString() string {
	return s.String()
}

func (s *TemplateCategoriesHeaders) SetCommonHeaders(v map[string]*string) *TemplateCategoriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TemplateCategoriesHeaders) SetXAcsDingtalkAccessToken(v string) *TemplateCategoriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TemplateCategoriesRequest struct {
	Option *TemplateCategoriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	Param *TemplateCategoriesRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s TemplateCategoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s TemplateCategoriesRequest) GoString() string {
	return s.String()
}

func (s *TemplateCategoriesRequest) SetOption(v *TemplateCategoriesRequestOption) *TemplateCategoriesRequest {
	s.Option = v
	return s
}

func (s *TemplateCategoriesRequest) SetParam(v *TemplateCategoriesRequestParam) *TemplateCategoriesRequest {
	s.Param = v
	return s
}

func (s *TemplateCategoriesRequest) SetOperatorId(v string) *TemplateCategoriesRequest {
	s.OperatorId = &v
	return s
}

type TemplateCategoriesRequestOption struct {
	// example:
	//
	// 1
	CategoryStatus *int32 `json:"categoryStatus,omitempty" xml:"categoryStatus,omitempty"`
	// example:
	//
	// -1
	IndustryId *int32 `json:"industryId,omitempty" xml:"industryId,omitempty"`
}

func (s TemplateCategoriesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s TemplateCategoriesRequestOption) GoString() string {
	return s.String()
}

func (s *TemplateCategoriesRequestOption) SetCategoryStatus(v int32) *TemplateCategoriesRequestOption {
	s.CategoryStatus = &v
	return s
}

func (s *TemplateCategoriesRequestOption) SetIndustryId(v int32) *TemplateCategoriesRequestOption {
	s.IndustryId = &v
	return s
}

type TemplateCategoriesRequestParam struct {
	// This parameter is required.
	//
	// example:
	//
	// tenantId
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s TemplateCategoriesRequestParam) String() string {
	return tea.Prettify(s)
}

func (s TemplateCategoriesRequestParam) GoString() string {
	return s.String()
}

func (s *TemplateCategoriesRequestParam) SetTenantId(v string) *TemplateCategoriesRequestParam {
	s.TenantId = &v
	return s
}

type TemplateCategoriesResponseBody struct {
	List []*TemplateCategoriesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s TemplateCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TemplateCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *TemplateCategoriesResponseBody) SetList(v []*TemplateCategoriesResponseBodyList) *TemplateCategoriesResponseBody {
	s.List = v
	return s
}

type TemplateCategoriesResponseBodyList struct {
	CategoryId   *string `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	CategoryName *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	Sort         *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s TemplateCategoriesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s TemplateCategoriesResponseBodyList) GoString() string {
	return s.String()
}

func (s *TemplateCategoriesResponseBodyList) SetCategoryId(v string) *TemplateCategoriesResponseBodyList {
	s.CategoryId = &v
	return s
}

func (s *TemplateCategoriesResponseBodyList) SetCategoryName(v string) *TemplateCategoriesResponseBodyList {
	s.CategoryName = &v
	return s
}

func (s *TemplateCategoriesResponseBodyList) SetSort(v string) *TemplateCategoriesResponseBodyList {
	s.Sort = &v
	return s
}

type TemplateCategoriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TemplateCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TemplateCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s TemplateCategoriesResponse) GoString() string {
	return s.String()
}

func (s *TemplateCategoriesResponse) SetHeaders(v map[string]*string) *TemplateCategoriesResponse {
	s.Headers = v
	return s
}

func (s *TemplateCategoriesResponse) SetStatusCode(v int32) *TemplateCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *TemplateCategoriesResponse) SetBody(v *TemplateCategoriesResponseBody) *TemplateCategoriesResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// union_id
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
	// example:
	//
	// true
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnmarkStarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// union_id
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
	// example:
	//
	// true
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnpinSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 这是更新后的简介
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 更新后的名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YEp3JcM******UIbhwiE
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TeamVO            `json:"body,omitempty" xml:"body,omitempty"`
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

type UserTemplatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UserTemplatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s UserTemplatesHeaders) GoString() string {
	return s.String()
}

func (s *UserTemplatesHeaders) SetCommonHeaders(v map[string]*string) *UserTemplatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UserTemplatesHeaders) SetXAcsDingtalkAccessToken(v string) *UserTemplatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UserTemplatesRequest struct {
	Option *UserTemplatesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UserTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s UserTemplatesRequest) GoString() string {
	return s.String()
}

func (s *UserTemplatesRequest) SetOption(v *UserTemplatesRequestOption) *UserTemplatesRequest {
	s.Option = v
	return s
}

func (s *UserTemplatesRequest) SetOperatorId(v string) *UserTemplatesRequest {
	s.OperatorId = &v
	return s
}

type UserTemplatesRequestOption struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// pc
	Platform      *string  `json:"platform,omitempty" xml:"platform,omitempty"`
	TemplateTypes []*int32 `json:"templateTypes,omitempty" xml:"templateTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UserTemplatesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s UserTemplatesRequestOption) GoString() string {
	return s.String()
}

func (s *UserTemplatesRequestOption) SetMaxResults(v int32) *UserTemplatesRequestOption {
	s.MaxResults = &v
	return s
}

func (s *UserTemplatesRequestOption) SetNextToken(v string) *UserTemplatesRequestOption {
	s.NextToken = &v
	return s
}

func (s *UserTemplatesRequestOption) SetPlatform(v string) *UserTemplatesRequestOption {
	s.Platform = &v
	return s
}

func (s *UserTemplatesRequestOption) SetTemplateTypes(v []*int32) *UserTemplatesRequestOption {
	s.TemplateTypes = v
	return s
}

func (s *UserTemplatesRequestOption) SetVersion(v int32) *UserTemplatesRequestOption {
	s.Version = &v
	return s
}

type UserTemplatesResponseBody struct {
	// example:
	//
	// next_token
	NextToken    *string                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TemplateList []*UserTemplatesResponseBodyTemplateList `json:"templateList,omitempty" xml:"templateList,omitempty" type:"Repeated"`
}

func (s UserTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UserTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *UserTemplatesResponseBody) SetNextToken(v string) *UserTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *UserTemplatesResponseBody) SetTemplateList(v []*UserTemplatesResponseBodyTemplateList) *UserTemplatesResponseBody {
	s.TemplateList = v
	return s
}

type UserTemplatesResponseBodyTemplateList struct {
	AuthorName         *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	Belong             *string `json:"belong,omitempty" xml:"belong,omitempty"`
	ContentDownloadUrl *string `json:"contentDownloadUrl,omitempty" xml:"contentDownloadUrl,omitempty"`
	CoverDownloadUrl   *string `json:"coverDownloadUrl,omitempty" xml:"coverDownloadUrl,omitempty"`
	CreateTime         *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	ModifiedTime       *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	TemplateId         *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
	Type               *int32  `json:"type,omitempty" xml:"type,omitempty"`
	UsedCount          *int64  `json:"usedCount,omitempty" xml:"usedCount,omitempty"`
	WorkspaceId        *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	WorkspaceName      *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s UserTemplatesResponseBodyTemplateList) String() string {
	return tea.Prettify(s)
}

func (s UserTemplatesResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *UserTemplatesResponseBodyTemplateList) SetAuthorName(v string) *UserTemplatesResponseBodyTemplateList {
	s.AuthorName = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetBelong(v string) *UserTemplatesResponseBodyTemplateList {
	s.Belong = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetContentDownloadUrl(v string) *UserTemplatesResponseBodyTemplateList {
	s.ContentDownloadUrl = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetCoverDownloadUrl(v string) *UserTemplatesResponseBodyTemplateList {
	s.CoverDownloadUrl = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetCreateTime(v string) *UserTemplatesResponseBodyTemplateList {
	s.CreateTime = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetDescription(v string) *UserTemplatesResponseBodyTemplateList {
	s.Description = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetModifiedTime(v string) *UserTemplatesResponseBodyTemplateList {
	s.ModifiedTime = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetTemplateId(v string) *UserTemplatesResponseBodyTemplateList {
	s.TemplateId = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetTitle(v string) *UserTemplatesResponseBodyTemplateList {
	s.Title = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetType(v int32) *UserTemplatesResponseBodyTemplateList {
	s.Type = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetUsedCount(v int64) *UserTemplatesResponseBodyTemplateList {
	s.UsedCount = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetWorkspaceId(v string) *UserTemplatesResponseBodyTemplateList {
	s.WorkspaceId = &v
	return s
}

func (s *UserTemplatesResponseBodyTemplateList) SetWorkspaceName(v string) *UserTemplatesResponseBodyTemplateList {
	s.WorkspaceName = &v
	return s
}

type UserTemplatesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UserTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UserTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s UserTemplatesResponse) GoString() string {
	return s.String()
}

func (s *UserTemplatesResponse) SetHeaders(v map[string]*string) *UserTemplatesResponse {
	s.Headers = v
	return s
}

func (s *UserTemplatesResponse) SetStatusCode(v int32) *UserTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *UserTemplatesResponse) SetBody(v *UserTemplatesResponseBody) *UserTemplatesResponse {
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
// 批量创建小组
//
// @param request - BatchCreateTeamRequest
//
// @param headers - BatchCreateTeamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateTeamResponse
func (client *Client) BatchCreateTeamWithOptions(request *BatchCreateTeamRequest, headers *BatchCreateTeamHeaders, runtime *util.RuntimeOptions) (_result *BatchCreateTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("BatchCreateTeam"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCreateTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量创建小组
//
// @param request - BatchCreateTeamRequest
//
// @return BatchCreateTeamResponse
func (client *Client) BatchCreateTeam(request *BatchCreateTeamRequest) (_result *BatchCreateTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchCreateTeamHeaders{}
	_result = &BatchCreateTeamResponse{}
	_body, _err := client.BatchCreateTeamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除文档最近记录
//
// @param request - BatchDeleteRecentsRequest
//
// @param headers - BatchDeleteRecentsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteRecentsResponse
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

// Summary:
//
// 批量删除文档最近记录
//
// @param request - BatchDeleteRecentsRequest
//
// @return BatchDeleteRecentsResponse
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

// Summary:
//
// 按分类列表查询模板列表
//
// @param request - CategoriesTemplatesRequest
//
// @param headers - CategoriesTemplatesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CategoriesTemplatesResponse
func (client *Client) CategoriesTemplatesWithOptions(request *CategoriesTemplatesRequest, headers *CategoriesTemplatesHeaders, runtime *util.RuntimeOptions) (_result *CategoriesTemplatesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("CategoriesTemplates"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/categoryLists/templates/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CategoriesTemplatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 按分类列表查询模板列表
//
// @param request - CategoriesTemplatesRequest
//
// @return CategoriesTemplatesResponse
func (client *Client) CategoriesTemplates(request *CategoriesTemplatesRequest) (_result *CategoriesTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CategoriesTemplatesHeaders{}
	_result = &CategoriesTemplatesResponse{}
	_body, _err := client.CategoriesTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 按分类查询模板列表
//
// @param request - CategoryTemplatesRequest
//
// @param headers - CategoryTemplatesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CategoryTemplatesResponse
func (client *Client) CategoryTemplatesWithOptions(request *CategoryTemplatesRequest, headers *CategoryTemplatesHeaders, runtime *util.RuntimeOptions) (_result *CategoryTemplatesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("CategoryTemplates"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/categories/templates/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CategoryTemplatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 按分类查询模板列表
//
// @param request - CategoryTemplatesRequest
//
// @return CategoryTemplatesResponse
func (client *Client) CategoryTemplates(request *CategoryTemplatesRequest) (_result *CategoryTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CategoryTemplatesHeaders{}
	_result = &CategoryTemplatesResponse{}
	_body, _err := client.CategoryTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拷贝知识库节点
//
// @param request - CopyDentryRequest
//
// @param headers - CopyDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyDentryResponse
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

// Summary:
//
// 拷贝知识库节点
//
// @param request - CopyDentryRequest
//
// @return CopyDentryResponse
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

// Summary:
//
// 复制文档
//
// @param request - CopyDocRequest
//
// @param headers - CopyDocHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyDocResponse
func (client *Client) CopyDocWithOptions(request *CopyDocRequest, headers *CopyDocHeaders, runtime *util.RuntimeOptions) (_result *CopyDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("CopyDoc"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/copy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyDocResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 复制文档
//
// @param request - CopyDocRequest
//
// @return CopyDocResponse
func (client *Client) CopyDoc(request *CopyDocRequest) (_result *CopyDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CopyDocHeaders{}
	_result = &CopyDocResponse{}
	_body, _err := client.CopyDocWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拷贝知识库
//
// @param request - CopyWorkspaceRequest
//
// @param headers - CopyWorkspaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyWorkspaceResponse
func (client *Client) CopyWorkspaceWithOptions(request *CopyWorkspaceRequest, headers *CopyWorkspaceHeaders, runtime *util.RuntimeOptions) (_result *CopyWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("CopyWorkspace"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/workspace/copy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyWorkspaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拷贝知识库
//
// @param request - CopyWorkspaceRequest
//
// @return CopyWorkspaceResponse
func (client *Client) CopyWorkspace(request *CopyWorkspaceRequest) (_result *CopyWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CopyWorkspaceHeaders{}
	_result = &CopyWorkspaceResponse{}
	_body, _err := client.CopyWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 异步拷贝知识库
//
// @param request - CopyWorkspaceAsyncRequest
//
// @param headers - CopyWorkspaceAsyncHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyWorkspaceAsyncResponse
func (client *Client) CopyWorkspaceAsyncWithOptions(request *CopyWorkspaceAsyncRequest, headers *CopyWorkspaceAsyncHeaders, runtime *util.RuntimeOptions) (_result *CopyWorkspaceAsyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("CopyWorkspaceAsync"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/workspace/asyncCopy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyWorkspaceAsyncResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 异步拷贝知识库
//
// @param request - CopyWorkspaceAsyncRequest
//
// @return CopyWorkspaceAsyncResponse
func (client *Client) CopyWorkspaceAsync(request *CopyWorkspaceAsyncRequest) (_result *CopyWorkspaceAsyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CopyWorkspaceAsyncHeaders{}
	_result = &CopyWorkspaceAsyncResponse{}
	_body, _err := client.CopyWorkspaceAsyncWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建知识库节点（包括文档和文件夹）
//
// @param request - CreateDentryRequest
//
// @param headers - CreateDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDentryResponse
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

// Summary:
//
// 创建知识库节点（包括文档和文件夹）
//
// @param request - CreateDentryRequest
//
// @return CreateDentryResponse
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

// Summary:
//
// 创建快捷方式
//
// @param request - CreateShortcutRequest
//
// @param headers - CreateShortcutHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateShortcutResponse
func (client *Client) CreateShortcutWithOptions(request *CreateShortcutRequest, headers *CreateShortcutHeaders, runtime *util.RuntimeOptions) (_result *CreateShortcutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("CreateShortcut"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/resource/shortcut/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateShortcutResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建快捷方式
//
// @param request - CreateShortcutRequest
//
// @return CreateShortcutResponse
func (client *Client) CreateShortcut(request *CreateShortcutRequest) (_result *CreateShortcutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateShortcutHeaders{}
	_result = &CreateShortcutResponse{}
	_body, _err := client.CreateShortcutWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建知识库
//
// @param request - CreateSpaceRequest
//
// @param headers - CreateSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSpaceResponse
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

// Summary:
//
// 创建知识库
//
// @param request - CreateSpaceRequest
//
// @return CreateSpaceResponse
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

// Summary:
//
// 创建小组
//
// @param request - CreateTeamRequest
//
// @param headers - CreateTeamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTeamResponse
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

// Summary:
//
// 创建小组
//
// @param request - CreateTeamRequest
//
// @return CreateTeamResponse
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

// Summary:
//
// 跨组织迁移知识库
//
// @param request - CrossOrgMigrateRequest
//
// @param headers - CrossOrgMigrateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CrossOrgMigrateResponse
func (client *Client) CrossOrgMigrateWithOptions(request *CrossOrgMigrateRequest, headers *CrossOrgMigrateHeaders, runtime *util.RuntimeOptions) (_result *CrossOrgMigrateResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("CrossOrgMigrate"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/crossOrganizations/spaces/migrate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CrossOrgMigrateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 跨组织迁移知识库
//
// @param request - CrossOrgMigrateRequest
//
// @return CrossOrgMigrateResponse
func (client *Client) CrossOrgMigrate(request *CrossOrgMigrateRequest) (_result *CrossOrgMigrateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CrossOrgMigrateHeaders{}
	_result = &CrossOrgMigrateResponse{}
	_body, _err := client.CrossOrgMigrateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除小组
//
// @param request - DeleteTeamRequest
//
// @param headers - DeleteTeamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTeamResponse
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

// Summary:
//
// 删除小组
//
// @param request - DeleteTeamRequest
//
// @return DeleteTeamResponse
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

// Summary:
//
// 获取文档内容
//
// @param request - DocContentRequest
//
// @param headers - DocContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocContentResponse
func (client *Client) DocContentWithOptions(dentryUuid *string, request *DocContentRequest, headers *DocContentHeaders, runtime *util.RuntimeOptions) (_result *DocContentResponse, _err error) {
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
		Action:      tea.String("DocContent"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/" + tea.StringValue(dentryUuid) + "/contents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DocContentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档内容
//
// @param request - DocContentRequest
//
// @return DocContentResponse
func (client *Client) DocContent(dentryUuid *string, request *DocContentRequest) (_result *DocContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DocContentHeaders{}
	_result = &DocContentResponse{}
	_body, _err := client.DocContentWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过委托权限方式导出文档内容
//
// @param request - DocExportByDelegatedPermissionRequest
//
// @param headers - DocExportByDelegatedPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocExportByDelegatedPermissionResponse
func (client *Client) DocExportByDelegatedPermissionWithOptions(dentryUuid *string, request *DocExportByDelegatedPermissionRequest, headers *DocExportByDelegatedPermissionHeaders, runtime *util.RuntimeOptions) (_result *DocExportByDelegatedPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GenerateCp)) {
		query["generateCp"] = request.GenerateCp
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFormat)) {
		query["targetFormat"] = request.TargetFormat
	}

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
		Action:      tea.String("DocExportByDelegatedPermission"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/me/dentries/" + tea.StringValue(dentryUuid) + "/export"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DocExportByDelegatedPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过委托权限方式导出文档内容
//
// @param request - DocExportByDelegatedPermissionRequest
//
// @return DocExportByDelegatedPermissionResponse
func (client *Client) DocExportByDelegatedPermission(dentryUuid *string, request *DocExportByDelegatedPermissionRequest) (_result *DocExportByDelegatedPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DocExportByDelegatedPermissionHeaders{}
	_result = &DocExportByDelegatedPermissionResponse{}
	_body, _err := client.DocExportByDelegatedPermissionWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 以委托权限方式覆写全文
//
// @param request - DocUpdateContentWithDelegatedPermissionRequest
//
// @param headers - DocUpdateContentWithDelegatedPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocUpdateContentWithDelegatedPermissionResponse
func (client *Client) DocUpdateContentWithDelegatedPermissionWithOptions(docKey *string, request *DocUpdateContentWithDelegatedPermissionRequest, headers *DocUpdateContentWithDelegatedPermissionHeaders, runtime *util.RuntimeOptions) (_result *DocUpdateContentWithDelegatedPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		body["dataType"] = request.DataType
	}

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
		Action:      tea.String("DocUpdateContentWithDelegatedPermission"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/me/suites/documents/" + tea.StringValue(docKey) + "/overwriteContent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DocUpdateContentWithDelegatedPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 以委托权限方式覆写全文
//
// @param request - DocUpdateContentWithDelegatedPermissionRequest
//
// @return DocUpdateContentWithDelegatedPermissionResponse
func (client *Client) DocUpdateContentWithDelegatedPermission(docKey *string, request *DocUpdateContentWithDelegatedPermissionRequest) (_result *DocUpdateContentWithDelegatedPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DocUpdateContentWithDelegatedPermissionHeaders{}
	_result = &DocUpdateContentWithDelegatedPermissionResponse{}
	_body, _err := client.DocUpdateContentWithDelegatedPermissionWithOptions(docKey, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出文档
//
// @param request - ExportDocRequest
//
// @param headers - ExportDocHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportDocResponse
func (client *Client) ExportDocWithOptions(request *ExportDocRequest, headers *ExportDocHeaders, runtime *util.RuntimeOptions) (_result *ExportDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("ExportDoc"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/export"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportDocResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出文档
//
// @param request - ExportDocRequest
//
// @return ExportDocResponse
func (client *Client) ExportDoc(request *ExportDocRequest) (_result *ExportDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExportDocHeaders{}
	_result = &ExportDocResponse{}
	_body, _err := client.ExportDocWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据文件DentryUuid获取文件DentryId
//
// @param request - GetDentryIdByUuidRequest
//
// @param headers - GetDentryIdByUuidHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDentryIdByUuidResponse
func (client *Client) GetDentryIdByUuidWithOptions(dentryUuid *string, request *GetDentryIdByUuidRequest, headers *GetDentryIdByUuidHeaders, runtime *util.RuntimeOptions) (_result *GetDentryIdByUuidResponse, _err error) {
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
		Action:      tea.String("GetDentryIdByUuid"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/" + tea.StringValue(dentryUuid) + "/queryDentryId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDentryIdByUuidResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据文件DentryUuid获取文件DentryId
//
// @param request - GetDentryIdByUuidRequest
//
// @return GetDentryIdByUuidResponse
func (client *Client) GetDentryIdByUuid(dentryUuid *string, request *GetDentryIdByUuidRequest) (_result *GetDentryIdByUuidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDentryIdByUuidHeaders{}
	_result = &GetDentryIdByUuidResponse{}
	_body, _err := client.GetDentryIdByUuidWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 委托权限获取文档内容
//
// @param request - GetDocContentRequest
//
// @param headers - GetDocContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocContentResponse
func (client *Client) GetDocContentWithOptions(dentryUuid *string, request *GetDocContentRequest, headers *GetDocContentHeaders, runtime *util.RuntimeOptions) (_result *GetDocContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GenerateCp)) {
		query["generateCp"] = request.GenerateCp
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFormat)) {
		query["targetFormat"] = request.TargetFormat
	}

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
		Action:      tea.String("GetDocContent"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/me/query/" + tea.StringValue(dentryUuid) + "/contents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocContentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 委托权限获取文档内容
//
// @param request - GetDocContentRequest
//
// @return GetDocContentResponse
func (client *Client) GetDocContent(dentryUuid *string, request *GetDocContentRequest) (_result *GetDocContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDocContentHeaders{}
	_result = &GetDocContentResponse{}
	_body, _err := client.GetDocContentWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 委托权限获取文档内容
//
// @param request - GetDocContentForELMRequest
//
// @param headers - GetDocContentForELMHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocContentForELMResponse
func (client *Client) GetDocContentForELMWithOptions(dentryUuid *string, request *GetDocContentForELMRequest, headers *GetDocContentForELMHeaders, runtime *util.RuntimeOptions) (_result *GetDocContentForELMResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetFormat)) {
		query["targetFormat"] = request.TargetFormat
	}

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
		Action:      tea.String("GetDocContentForELM"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/elm/me/dentries/" + tea.StringValue(dentryUuid) + "/contents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocContentForELMResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 委托权限获取文档内容
//
// @param request - GetDocContentForELMRequest
//
// @return GetDocContentForELMResponse
func (client *Client) GetDocContentForELM(dentryUuid *string, request *GetDocContentForELMRequest) (_result *GetDocContentForELMResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDocContentForELMHeaders{}
	_result = &GetDocContentForELMResponse{}
	_body, _err := client.GetDocContentForELMWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前企业下钉盘目录我的文件对应的空间信息
//
// @param request - GetMySpaceRequest
//
// @param headers - GetMySpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMySpaceResponse
func (client *Client) GetMySpaceWithOptions(request *GetMySpaceRequest, headers *GetMySpaceHeaders, runtime *util.RuntimeOptions) (_result *GetMySpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsMySpace)) {
		query["isMySpace"] = request.IsMySpace
	}

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
		Action:      tea.String("GetMySpace"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/me/mySpace/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMySpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前企业下钉盘目录我的文件对应的空间信息
//
// @param request - GetMySpaceRequest
//
// @return GetMySpaceResponse
func (client *Client) GetMySpace(request *GetMySpaceRequest) (_result *GetMySpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMySpaceHeaders{}
	_result = &GetMySpaceResponse{}
	_body, _err := client.GetMySpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 委托权限获取企业内公开或者或联网公开文档内容
//
// @param request - GetOrgOrWebOpenDocContentRequest
//
// @param headers - GetOrgOrWebOpenDocContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrgOrWebOpenDocContentResponse
func (client *Client) GetOrgOrWebOpenDocContentWithOptions(dentryUuid *string, request *GetOrgOrWebOpenDocContentRequest, headers *GetOrgOrWebOpenDocContentHeaders, runtime *util.RuntimeOptions) (_result *GetOrgOrWebOpenDocContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GenerateCp)) {
		query["generateCp"] = request.GenerateCp
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeType)) {
		query["scopeType"] = request.ScopeType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFormat)) {
		query["targetFormat"] = request.TargetFormat
	}

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
		Action:      tea.String("GetOrgOrWebOpenDocContent"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/me/dentries/" + tea.StringValue(dentryUuid) + "/contents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrgOrWebOpenDocContentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 委托权限获取企业内公开或者或联网公开文档内容
//
// @param request - GetOrgOrWebOpenDocContentRequest
//
// @return GetOrgOrWebOpenDocContentResponse
func (client *Client) GetOrgOrWebOpenDocContent(dentryUuid *string, request *GetOrgOrWebOpenDocContentRequest) (_result *GetOrgOrWebOpenDocContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrgOrWebOpenDocContentHeaders{}
	_result = &GetOrgOrWebOpenDocContentResponse{}
	_body, _err := client.GetOrgOrWebOpenDocContentWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询小组主页schema （包括轮播图、公告、便捷入口）
//
// @param request - GetSchemaRequest
//
// @param headers - GetSchemaHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSchemaResponse
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

// Summary:
//
// 查询小组主页schema （包括轮播图、公告、便捷入口）
//
// @param request - GetSchemaRequest
//
// @return GetSchemaResponse
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

// Summary:
//
// 查询目录树
//
// @param request - GetSpaceDirectoriesRequest
//
// @param headers - GetSpaceDirectoriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpaceDirectoriesResponse
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

// Summary:
//
// 查询目录树
//
// @param request - GetSpaceDirectoriesRequest
//
// @return GetSpaceDirectoriesResponse
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

// Summary:
//
// 获取星标信息
//
// @param request - GetStarInfoRequest
//
// @param headers - GetStarInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStarInfoResponse
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

// Summary:
//
// 获取星标信息
//
// @param request - GetStarInfoRequest
//
// @return GetStarInfoResponse
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

// Summary:
//
// 获取任务信息
//
// @param headers - GetTaskInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskInfoResponse
func (client *Client) GetTaskInfoWithOptions(taskId *string, headers *GetTaskInfoHeaders, runtime *util.RuntimeOptions) (_result *GetTaskInfoResponse, _err error) {
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
		Action:      tea.String("GetTaskInfo"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/task/info/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务信息
//
// @return GetTaskInfoResponse
func (client *Client) GetTaskInfo(taskId *string) (_result *GetTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTaskInfoHeaders{}
	_result = &GetTaskInfoResponse{}
	_body, _err := client.GetTaskInfoWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询小组详情
//
// @param request - GetTeamRequest
//
// @param headers - GetTeamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTeamResponse
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

// Summary:
//
// 查询小组详情
//
// @param request - GetTeamRequest
//
// @return GetTeamResponse
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

// Summary:
//
// 获取知识库下的总节点数
//
// @param request - GetTotalNumberOfDentriesRequest
//
// @param headers - GetTotalNumberOfDentriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTotalNumberOfDentriesResponse
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

// Summary:
//
// 获取知识库下的总节点数
//
// @param request - GetTotalNumberOfDentriesRequest
//
// @return GetTotalNumberOfDentriesResponse
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

// Summary:
//
// 获取知识库总数
//
// @param request - GetTotalNumberOfSpacesRequest
//
// @param headers - GetTotalNumberOfSpacesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTotalNumberOfSpacesResponse
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

// Summary:
//
// 获取知识库总数
//
// @param request - GetTotalNumberOfSpacesRequest
//
// @return GetTotalNumberOfSpacesResponse
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

// Summary:
//
// 查询文档免登的用户信息
//
// @param request - GetUserInfoByOpenTokenRequest
//
// @param headers - GetUserInfoByOpenTokenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInfoByOpenTokenResponse
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

// Summary:
//
// 查询文档免登的用户信息
//
// @param request - GetUserInfoByOpenTokenRequest
//
// @return GetUserInfoByOpenTokenResponse
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

// Summary:
//
// 根据文件DentryId获取文件DentryUuid
//
// @param request - GetUuidByDentryIdRequest
//
// @param headers - GetUuidByDentryIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUuidByDentryIdResponse
func (client *Client) GetUuidByDentryIdWithOptions(dentryId *string, request *GetUuidByDentryIdRequest, headers *GetUuidByDentryIdHeaders, runtime *util.RuntimeOptions) (_result *GetUuidByDentryIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		query["spaceId"] = request.SpaceId
	}

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
		Action:      tea.String("GetUuidByDentryId"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/" + tea.StringValue(dentryId) + "/queryDentryUuid"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUuidByDentryIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据文件DentryId获取文件DentryUuid
//
// @param request - GetUuidByDentryIdRequest
//
// @return GetUuidByDentryIdResponse
func (client *Client) GetUuidByDentryId(dentryId *string, request *GetUuidByDentryIdRequest) (_result *GetUuidByDentryIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUuidByDentryIdHeaders{}
	_result = &GetUuidByDentryIdResponse{}
	_body, _err := client.GetUuidByDentryIdWithOptions(dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取知识库权限范围
//
// @param request - GetWorkspacePermissionScopesRequest
//
// @param headers - GetWorkspacePermissionScopesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspacePermissionScopesResponse
func (client *Client) GetWorkspacePermissionScopesWithOptions(workspaceId *string, request *GetWorkspacePermissionScopesRequest, headers *GetWorkspacePermissionScopesHeaders, runtime *util.RuntimeOptions) (_result *GetWorkspacePermissionScopesResponse, _err error) {
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
		Action:      tea.String("GetWorkspacePermissionScopes"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/workspaces/" + tea.StringValue(workspaceId) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkspacePermissionScopesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识库权限范围
//
// @param request - GetWorkspacePermissionScopesRequest
//
// @return GetWorkspacePermissionScopesResponse
func (client *Client) GetWorkspacePermissionScopes(workspaceId *string, request *GetWorkspacePermissionScopesRequest) (_result *GetWorkspacePermissionScopesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWorkspacePermissionScopesHeaders{}
	_result = &GetWorkspacePermissionScopesResponse{}
	_body, _err := client.GetWorkspacePermissionScopesWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 以超级管理员身份转交小组
//
// @param request - HandoverTeamWithoutAuthRequest
//
// @param headers - HandoverTeamWithoutAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HandoverTeamWithoutAuthResponse
func (client *Client) HandoverTeamWithoutAuthWithOptions(request *HandoverTeamWithoutAuthRequest, headers *HandoverTeamWithoutAuthHeaders, runtime *util.RuntimeOptions) (_result *HandoverTeamWithoutAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("HandoverTeamWithoutAuth"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/teams/members/handoverWithoutAuth"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HandoverTeamWithoutAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 以超级管理员身份转交小组
//
// @param request - HandoverTeamWithoutAuthRequest
//
// @return HandoverTeamWithoutAuthResponse
func (client *Client) HandoverTeamWithoutAuth(request *HandoverTeamWithoutAuthRequest) (_result *HandoverTeamWithoutAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HandoverTeamWithoutAuthHeaders{}
	_result = &HandoverTeamWithoutAuthResponse{}
	_body, _err := client.HandoverTeamWithoutAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 知识库转交所有者
//
// @param request - HandoveryWorkspaceRequest
//
// @param headers - HandoveryWorkspaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HandoveryWorkspaceResponse
func (client *Client) HandoveryWorkspaceWithOptions(request *HandoveryWorkspaceRequest, headers *HandoveryWorkspaceHeaders, runtime *util.RuntimeOptions) (_result *HandoveryWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("HandoveryWorkspace"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/workspace/handover"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HandoveryWorkspaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 知识库转交所有者
//
// @param request - HandoveryWorkspaceRequest
//
// @return HandoveryWorkspaceResponse
func (client *Client) HandoveryWorkspace(request *HandoveryWorkspaceRequest) (_result *HandoveryWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HandoveryWorkspaceHeaders{}
	_result = &HandoveryWorkspaceResponse{}
	_body, _err := client.HandoveryWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询小组动态
//
// @param request - ListFeedsRequest
//
// @param headers - ListFeedsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeedsResponse
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

// Summary:
//
// 查询小组动态
//
// @param request - ListFeedsRequest
//
// @return ListFeedsResponse
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

// Summary:
//
// 查询小组热门文档
//
// @param request - ListHotDocsRequest
//
// @param headers - ListHotDocsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotDocsResponse
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

// Summary:
//
// 查询小组热门文档
//
// @param request - ListHotDocsRequest
//
// @return ListHotDocsResponse
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

// Summary:
//
// 获取置顶知识库列表
//
// @param request - ListPinSpacesRequest
//
// @param headers - ListPinSpacesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPinSpacesResponse
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

// Summary:
//
// 获取置顶知识库列表
//
// @param request - ListPinSpacesRequest
//
// @return ListPinSpacesResponse
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

// Summary:
//
// 查询文档最近记录列表
//
// @param request - ListRecentsRequest
//
// @param headers - ListRecentsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentsResponse
func (client *Client) ListRecentsWithOptions(request *ListRecentsRequest, headers *ListRecentsHeaders, runtime *util.RuntimeOptions) (_result *ListRecentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("ListRecents"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/recentRecords/lists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecentsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文档最近记录列表
//
// @param request - ListRecentsRequest
//
// @return ListRecentsResponse
func (client *Client) ListRecents(request *ListRecentsRequest) (_result *ListRecentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRecentsHeaders{}
	_result = &ListRecentsResponse{}
	_body, _err := client.ListRecentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询关联了知识库的团队列表
//
// @param request - ListRelatedSpaceTeamsRequest
//
// @param headers - ListRelatedSpaceTeamsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRelatedSpaceTeamsResponse
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

// Summary:
//
// 查询关联了知识库的团队列表
//
// @param request - ListRelatedSpaceTeamsRequest
//
// @return ListRelatedSpaceTeamsResponse
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

// Summary:
//
// 查询小组列表
//
// @param request - ListRelatedTeamsRequest
//
// @param headers - ListRelatedTeamsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRelatedTeamsResponse
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

// Summary:
//
// 查询小组列表
//
// @param request - ListRelatedTeamsRequest
//
// @return ListRelatedTeamsResponse
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

// Summary:
//
// 查询知识库分组
//
// @param request - ListSpaceSectionsRequest
//
// @param headers - ListSpaceSectionsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSpaceSectionsResponse
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

// Summary:
//
// 查询知识库分组
//
// @param request - ListSpaceSectionsRequest
//
// @return ListSpaceSectionsResponse
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

// Summary:
//
// 获取星标列表
//
// @param request - ListStarsRequest
//
// @param headers - ListStarsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStarsResponse
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

// Summary:
//
// 获取星标列表
//
// @param request - ListStarsRequest
//
// @return ListStarsResponse
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

// Summary:
//
// 查询小组成员列表
//
// @param request - ListTeamMembersRequest
//
// @param headers - ListTeamMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTeamMembersResponse
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

// Summary:
//
// 查询小组成员列表
//
// @param request - ListTeamMembersRequest
//
// @return ListTeamMembersResponse
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

// Summary:
//
// 标记星标
//
// @param request - MarkStarRequest
//
// @param headers - MarkStarHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MarkStarResponse
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

// Summary:
//
// 标记星标
//
// @param request - MarkStarRequest
//
// @return MarkStarResponse
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

// Summary:
//
// 移动知识库节点
//
// @param request - MoveDentryRequest
//
// @param headers - MoveDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveDentryResponse
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

// Summary:
//
// 移动知识库节点
//
// @param request - MoveDentryRequest
//
// @return MoveDentryResponse
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

// Summary:
//
// 置顶知识库
//
// @param request - PinSpaceRequest
//
// @param headers - PinSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PinSpaceResponse
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

// Summary:
//
// 置顶知识库
//
// @param request - PinSpaceRequest
//
// @return PinSpaceResponse
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

// Summary:
//
// 查询知识库节点（包括文档和文件夹）
//
// @param request - QueryDentryRequest
//
// @param headers - QueryDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDentryResponse
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

// Summary:
//
// 查询知识库节点（包括文档和文件夹）
//
// @param request - QueryDentryRequest
//
// @return QueryDentryResponse
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

// Summary:
//
// 获取文档内容
//
// @param request - QueryDocContentRequest
//
// @param headers - QueryDocContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDocContentResponse
func (client *Client) QueryDocContentWithOptions(dentryUuid *string, request *QueryDocContentRequest, headers *QueryDocContentHeaders, runtime *util.RuntimeOptions) (_result *QueryDocContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFormat)) {
		query["targetFormat"] = request.TargetFormat
	}

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
		Action:      tea.String("QueryDocContent"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/query/" + tea.StringValue(dentryUuid) + "/contents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDocContentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档内容
//
// @param request - QueryDocContentRequest
//
// @return QueryDocContentResponse
func (client *Client) QueryDocContent(dentryUuid *string, request *QueryDocContentRequest) (_result *QueryDocContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDocContentHeaders{}
	_result = &QueryDocContentResponse{}
	_body, _err := client.QueryDocContentWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询文档内容获取任务状态
//
// @param request - QueryGetContentJobRequest
//
// @param headers - QueryGetContentJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGetContentJobResponse
func (client *Client) QueryGetContentJobWithOptions(dentryUuid *string, request *QueryGetContentJobRequest, headers *QueryGetContentJobHeaders, runtime *util.RuntimeOptions) (_result *QueryGetContentJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

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
		Action:      tea.String("QueryGetContentJob"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/contents/" + tea.StringValue(dentryUuid) + "/jobStatuses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGetContentJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文档内容获取任务状态
//
// @param request - QueryGetContentJobRequest
//
// @return QueryGetContentJobResponse
func (client *Client) QueryGetContentJob(dentryUuid *string, request *QueryGetContentJobRequest) (_result *QueryGetContentJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGetContentJobHeaders{}
	_result = &QueryGetContentJobResponse{}
	_body, _err := client.QueryGetContentJobWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据链接查询节点或知识库信息
//
// @param request - QueryItemByUrlRequest
//
// @param headers - QueryItemByUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryItemByUrlResponse
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

// Summary:
//
// 根据链接查询节点或知识库信息
//
// @param request - QueryItemByUrlRequest
//
// @return QueryItemByUrlResponse
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

// Summary:
//
// 查询用户的「我的文档」
//
// @param headers - QueryMineSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMineSpaceResponse
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

// Summary:
//
// 查询用户的「我的文档」
//
// @return QueryMineSpaceResponse
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

// Summary:
//
// 查询最近列表
//
// @param request - QueryRecentListRequest
//
// @param headers - QueryRecentListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecentListResponse
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

// Summary:
//
// 查询最近列表
//
// @param request - QueryRecentListRequest
//
// @return QueryRecentListResponse
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

// Summary:
//
// 查询指定知识库信息
//
// @param request - QuerySpaceRequest
//
// @param headers - QuerySpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySpaceResponse
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

// Summary:
//
// 查询指定知识库信息
//
// @param request - QuerySpaceRequest
//
// @return QuerySpaceResponse
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

// Summary:
//
// 查询与我关联的知识库列表（支持筛选小组）
//
// @param request - RelatedSpacesRequest
//
// @param headers - RelatedSpacesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RelatedSpacesResponse
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

// Summary:
//
// 查询与我关联的知识库列表（支持筛选小组）
//
// @param request - RelatedSpacesRequest
//
// @return RelatedSpacesResponse
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

// Summary:
//
// 移除小组成员
//
// @param request - RemoveTeamMembersRequest
//
// @param headers - RemoveTeamMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTeamMembersResponse
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

// Summary:
//
// 移除小组成员
//
// @param request - RemoveTeamMembersRequest
//
// @return RemoveTeamMembersResponse
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

// Summary:
//
// 知识库节点（包括文档和文件夹）重命名
//
// @param request - RenameDentryRequest
//
// @param headers - RenameDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameDentryResponse
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

// Summary:
//
// 知识库节点（包括文档和文件夹）重命名
//
// @param request - RenameDentryRequest
//
// @return RenameDentryResponse
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

// Summary:
//
// 添加或修改小组成员
//
// @param request - SaveTeamMembersRequest
//
// @param headers - SaveTeamMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTeamMembersResponse
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

// Summary:
//
// 添加或修改小组成员
//
// @param request - SaveTeamMembersRequest
//
// @return SaveTeamMembersResponse
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

// Summary:
//
// 搜索知识库和节点
//
// @param request - SearchRequest
//
// @param headers - SearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchResponse
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

// Summary:
//
// 搜索知识库和节点
//
// @param request - SearchRequest
//
// @return SearchResponse
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

// Summary:
//
// 搜索模板中心模板
//
// @param request - SearchTemplatesRequest
//
// @param headers - SearchTemplatesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchTemplatesResponse
func (client *Client) SearchTemplatesWithOptions(request *SearchTemplatesRequest, headers *SearchTemplatesHeaders, runtime *util.RuntimeOptions) (_result *SearchTemplatesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("SearchTemplates"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/templates/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTemplatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索模板中心模板
//
// @param request - SearchTemplatesRequest
//
// @return SearchTemplatesResponse
func (client *Client) SearchTemplates(request *SearchTemplatesRequest) (_result *SearchTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchTemplatesHeaders{}
	_result = &SearchTemplatesResponse{}
	_body, _err := client.SearchTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件打开链接
//
// @param request - ShareUrlRequest
//
// @param headers - ShareUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShareUrlResponse
func (client *Client) ShareUrlWithOptions(request *ShareUrlRequest, headers *ShareUrlHeaders, runtime *util.RuntimeOptions) (_result *ShareUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("ShareUrl"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/dentries/shareUrls/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ShareUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件打开链接
//
// @param request - ShareUrlRequest
//
// @return ShareUrlResponse
func (client *Client) ShareUrl(request *ShareUrlRequest) (_result *ShareUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ShareUrlHeaders{}
	_result = &ShareUrlResponse{}
	_body, _err := client.ShareUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交文档内容获取任务
//
// @param request - SubmitGetContentJobRequest
//
// @param headers - SubmitGetContentJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitGetContentJobResponse
func (client *Client) SubmitGetContentJobWithOptions(dentryUuid *string, request *SubmitGetContentJobRequest, headers *SubmitGetContentJobHeaders, runtime *util.RuntimeOptions) (_result *SubmitGetContentJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFormat)) {
		query["targetFormat"] = request.TargetFormat
	}

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
		Action:      tea.String("SubmitGetContentJob"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/contents/" + tea.StringValue(dentryUuid) + "/jobs"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitGetContentJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交文档内容获取任务
//
// @param request - SubmitGetContentJobRequest
//
// @return SubmitGetContentJobResponse
func (client *Client) SubmitGetContentJob(dentryUuid *string, request *SubmitGetContentJobRequest) (_result *SubmitGetContentJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubmitGetContentJobHeaders{}
	_result = &SubmitGetContentJobResponse{}
	_body, _err := client.SubmitGetContentJobWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取知识库模板列表
//
// @param request - TeamTemplatesRequest
//
// @param headers - TeamTemplatesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TeamTemplatesResponse
func (client *Client) TeamTemplatesWithOptions(request *TeamTemplatesRequest, headers *TeamTemplatesHeaders, runtime *util.RuntimeOptions) (_result *TeamTemplatesResponse, _err error) {
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
		Action:      tea.String("TeamTemplates"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/workspaces/templates/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TeamTemplatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识库模板列表
//
// @param request - TeamTemplatesRequest
//
// @return TeamTemplatesResponse
func (client *Client) TeamTemplates(request *TeamTemplatesRequest) (_result *TeamTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TeamTemplatesHeaders{}
	_result = &TeamTemplatesResponse{}
	_body, _err := client.TeamTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取模板分类列表
//
// @param request - TemplateCategoriesRequest
//
// @param headers - TemplateCategoriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TemplateCategoriesResponse
func (client *Client) TemplateCategoriesWithOptions(request *TemplateCategoriesRequest, headers *TemplateCategoriesHeaders, runtime *util.RuntimeOptions) (_result *TemplateCategoriesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
	}

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
		Action:      tea.String("TemplateCategories"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/templates/categories/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TemplateCategoriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模板分类列表
//
// @param request - TemplateCategoriesRequest
//
// @return TemplateCategoriesResponse
func (client *Client) TemplateCategories(request *TemplateCategoriesRequest) (_result *TemplateCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TemplateCategoriesHeaders{}
	_result = &TemplateCategoriesResponse{}
	_body, _err := client.TemplateCategoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消标记星标
//
// @param request - UnmarkStarRequest
//
// @param headers - UnmarkStarHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnmarkStarResponse
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

// Summary:
//
// 取消标记星标
//
// @param request - UnmarkStarRequest
//
// @return UnmarkStarResponse
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

// Summary:
//
// 取消置顶知识库
//
// @param request - UnpinSpaceRequest
//
// @param headers - UnpinSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnpinSpaceResponse
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

// Summary:
//
// 取消置顶知识库
//
// @param request - UnpinSpaceRequest
//
// @return UnpinSpaceResponse
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

// Summary:
//
// 更新小组
//
// @param request - UpdateTeamRequest
//
// @param headers - UpdateTeamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTeamResponse
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

// Summary:
//
// 更新小组
//
// @param request - UpdateTeamRequest
//
// @return UpdateTeamResponse
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

// Summary:
//
// 用户模板列表
//
// @param request - UserTemplatesRequest
//
// @param headers - UserTemplatesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UserTemplatesResponse
func (client *Client) UserTemplatesWithOptions(request *UserTemplatesRequest, headers *UserTemplatesHeaders, runtime *util.RuntimeOptions) (_result *UserTemplatesResponse, _err error) {
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
		Action:      tea.String("UserTemplates"),
		Version:     tea.String("doc_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/doc/users/templates/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UserTemplatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户模板列表
//
// @param request - UserTemplatesRequest
//
// @return UserTemplatesResponse
func (client *Client) UserTemplates(request *UserTemplatesRequest) (_result *UserTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UserTemplatesHeaders{}
	_result = &UserTemplatesResponse{}
	_body, _err := client.UserTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
