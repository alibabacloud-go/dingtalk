// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package doc_2_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DentryModel struct {
	// 内容类型。alidoc-钉钉文档；link-快捷方式；archive-压缩包；document-文件。
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 创建时间。
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 创建者。
	Creator *DentryModelCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// 节点id。
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// 节点类型。file-文件；folder-文件夹。
	DentryType *string `json:"dentryType,omitempty" xml:"dentryType,omitempty"`
	// 节点全局唯一标识id。
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// 文档docKey，用于标识一篇钉钉文档的key。只有内容类型为alidoc的才会有值。
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	// 文件后缀名。
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 是否有子节点。
	HasChildren *bool `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	// 快捷方式类型的节点，其指向的原始数据信息。
	LinkSourceInfo *LinkSourceInfo `json:"linkSourceInfo,omitempty" xml:"linkSourceInfo,omitempty"`
	// 节点名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 节点的路径。
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 知识库信息。
	Space *SpaceModel `json:"space,omitempty" xml:"space,omitempty"`
	// 知识库id。
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 更新时间。
	UpdatedTime *int64 `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	// 更新人。
	Updater *DentryModelUpdater `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	// 节点访问url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 访问者对当前节点的权限等信息。
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
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
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

type DentryModelUpdater struct {
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
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
	// 节点的操作列表。
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	// 当前用户对这个空间的访问角色。
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 空间的操作列表。
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
	// 内容类型。alidoc-钉钉文档；link-快捷方式；archive-压缩包；document-文件。
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 创建时间。
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 创建者。
	Creator *DentryVOCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// 节点id。
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// 节点类型。file-文件；folder-文件夹。
	DentryType *string `json:"dentryType,omitempty" xml:"dentryType,omitempty"`
	// 节点全局唯一标识id。
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// 文档docKey，用于标识一篇钉钉文档的key。只有内容类型为alidoc的才会有值。
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	// 文件后缀名。
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 是否有子节点。
	HasChildren *bool `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	// 快捷方式类型的节点，其指向的原始数据信息。
	LinkSourceInfo *LinkSourceInfo `json:"linkSourceInfo,omitempty" xml:"linkSourceInfo,omitempty"`
	// 节点名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 节点的路径。
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 知识库信息。
	Space *SpaceModel `json:"space,omitempty" xml:"space,omitempty"`
	// 知识库id。
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 更新时间。
	UpdatedTime *int64 `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	// 更新人。
	Updater *DentryVOUpdater `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	// 节点访问url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 访问者对当前节点的权限等信息。
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
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
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
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
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
	// 节点的操作列表。
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	// 当前用户对这个空间的访问角色。
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 空间的操作列表。
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
	// 快捷方式关联的源文件后缀。
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 非通用快捷方式的图标信息。
	IconUrl *LinkSourceInfoIconUrl `json:"iconUrl,omitempty" xml:"iconUrl,omitempty" type:"Struct"`
	// 快捷方式关联的源文件ID（空间内唯一）。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 快捷方式类型。0-通用快捷方式；1-闪会快捷方式；2-日志快捷方式；3-闪会2.0快捷方式。
	LinkType *int64 `json:"linkType,omitempty" xml:"linkType,omitempty"`
	// 快捷方式关联的源文件所属空间id。
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
	// 默认的目录树图标。
	Line *string `json:"line,omitempty" xml:"line,omitempty"`
	// 被选中时的加深图标。
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
	// 操作人名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 操作时间戳。
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
	// 封面
	Cover *string `json:"cover,omitempty" xml:"cover,omitempty"`
	// 空间描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 知识库图标
	IconVO *SpaceModelIconVO `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	// 知识库id。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 知识库名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 知识库所有者。
	Owner *SpaceModelOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// 知识库访问url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 访问者对当前知识库的权限等信息。
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

func (s *SpaceModel) SetUrl(v string) *SpaceModel {
	s.Url = &v
	return s
}

func (s *SpaceModel) SetVisitorInfo(v *SpaceModelVisitorInfo) *SpaceModel {
	s.VisitorInfo = v
	return s
}

type SpaceModelIconVO struct {
	// 图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 图标类型
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
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
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
	// 节点的操作列表。
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	// 权限
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 空间的操作列表。
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
	// 封面
	Cover *string `json:"cover,omitempty" xml:"cover,omitempty"`
	// 访问者对当前知识库的权限等信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 知识库图标
	IconVO *SpaceVOIconVO `json:"iconVO,omitempty" xml:"iconVO,omitempty" type:"Struct"`
	// 知识库id。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 知识库名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 知识库所有者。
	Owner *SpaceVOOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// 知识库访问url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 访问者对当前知识库的权限等信息。
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

func (s *SpaceVO) SetUrl(v string) *SpaceVO {
	s.Url = &v
	return s
}

func (s *SpaceVO) SetVisitorInfo(v *SpaceVOVisitorInfo) *SpaceVO {
	s.VisitorInfo = v
	return s
}

type SpaceVOIconVO struct {
	// 图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 图标类型
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
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
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
	// 节点的操作列表。
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	// 权限
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 空间的操作列表。
	SpaceActions []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
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
	// 拷贝后的文档名称，长度不能超过50。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 操作人unionId。
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 需要移动到的知识库id。
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// 移动到目标位置的后置节点id。不为空时，需要是归属于 toParentDentryId 的子节点。
	ToNextDentryId *string `json:"toNextDentryId,omitempty" xml:"toNextDentryId,omitempty"`
	// 需要移动到目标位置的父节点id。如果为根目录，则不传；如果为非根目录，则需要传对应的id。
	ToParentDentryId *string `json:"toParentDentryId,omitempty" xml:"toParentDentryId,omitempty"`
	// 移动到目标位置的前置节点id。不为空时，需要是归属于 toParentDentryId 的子节点。
	ToPrevDentryId *string `json:"toPrevDentryId,omitempty" xml:"toPrevDentryId,omitempty"`
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DentryVO          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 节点类型，file-文档，folder-文件夹。
	DentryType *string `json:"dentryType,omitempty" xml:"dentryType,omitempty"`
	// 节点类型为文档才有，0-文字，1-表格，2-PPT，3-白板，6-脑图，7-多维表。
	DocumentType *int64 `json:"documentType,omitempty" xml:"documentType,omitempty"`
	// 节点名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 操作人unionId。
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 父节点id，可为空。
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DentryVO          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDentryResponse) SetBody(v *DentryVO) *CreateDentryResponse {
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
	// 知识库节点id。
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// 查询数量，最大500。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页token，第一页可不传。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 操作用户unionId。
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
	// 子节点列表。
	Children []*DentryModel `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// 是否还有后续可查询子节点。
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页token。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSpaceDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSpaceDirectoriesResponse) SetBody(v *GetSpaceDirectoriesResponseBody) *GetSpaceDirectoriesResponse {
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
	// 文档docKey，标识一篇文档的key。
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	// 文档颁发给三方应用的 OpenToken，用于三方应用在文档中的免登。
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
	// 用户的 unionId。
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 用户的userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserInfoByOpenTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetUserInfoByOpenTokenResponse) SetBody(v *GetUserInfoByOpenTokenResponseBody) *GetUserInfoByOpenTokenResponse {
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
	// 操作人unionId。
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 需要移动到的知识库id。
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// 移动到目标位置的后置节点id。不为空时，需要是归属于 toParentDentryId 的子节点。
	ToNextDentryId *string `json:"toNextDentryId,omitempty" xml:"toNextDentryId,omitempty"`
	// 需要移动到目标位置的父节点id。如果为根目录，则不传；如果为非根目录，则需要传对应的id。
	ToParentDentryId *string `json:"toParentDentryId,omitempty" xml:"toParentDentryId,omitempty"`
	// 移动到目标位置的前置节点id。不为空时，需要是归属于 toParentDentryId 的子节点。
	ToPrevDentryId *string `json:"toPrevDentryId,omitempty" xml:"toPrevDentryId,omitempty"`
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DentryVO          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *MoveDentryResponse) SetBody(v *DentryVO) *MoveDentryResponse {
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
	// 是否查询知识库信息。
	IncludeSpace *bool `json:"includeSpace,omitempty" xml:"includeSpace,omitempty"`
	// 操作用户unionId。
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DentryVO          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDentryResponse) SetBody(v *DentryVO) *QueryDentryResponse {
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SpaceVO           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 创建人类型。0-全部；1-我创建的；2-他人创建；不填也是查全部。
	CreatorType *int32 `json:"creatorType,omitempty" xml:"creatorType,omitempty"`
	// 访问文档类型：0-文字；1-表格；2-PPT；3-白板；6-脑图；7-多维表；不填的话，则默认是所有。
	FileType *int32 `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 每页最大条目数，最大值50。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标，第一页可不传。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 操作用户unionId。
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 最近列表的类型：0-最近访问；1-最近编辑。
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
	// 是否还有更多数据。
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页游标。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 子节点列表。
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
	// 是否被删除。
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 节点信息。
	Dentry *DentryModel `json:"dentry,omitempty" xml:"dentry,omitempty"`
	// 如果查询的是访问，那么这里是访问时间；否则就是编辑时间。
	RecentTime *int64 `json:"recentTime,omitempty" xml:"recentTime,omitempty"`
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

type QueryRecentListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRecentListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 操作用户unionId。
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SpaceVO           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 每页最大条目数，最大值50。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标，第一页可不传。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 操作用户unionId。
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 团队id。
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
	// 是否还有更多数据。
	HasMore *bool         `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Items   []*SpaceModel `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 分页游标。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

func (s *RelatedSpacesResponseBody) SetItems(v []*SpaceModel) *RelatedSpacesResponseBody {
	s.Items = v
	return s
}

func (s *RelatedSpacesResponseBody) SetNextToken(v string) *RelatedSpacesResponseBody {
	s.NextToken = &v
	return s
}

type RelatedSpacesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RelatedSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RelatedSpacesResponse) SetBody(v *RelatedSpacesResponseBody) *RelatedSpacesResponse {
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
	// 重命名之后的节点名称，长度不能超过50。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 操作人unionId。
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DentryVO          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RenameDentryResponse) SetBody(v *DentryVO) *RenameDentryResponse {
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
	// 节点搜索请求，和空间搜索请求二选一必填。
	DentryRequest *SearchRequestDentryRequest `json:"dentryRequest,omitempty" xml:"dentryRequest,omitempty" type:"Struct"`
	//  搜索关键词。
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 操作人unionId。
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 空间搜索请求，和节点搜索请求二选一必填。
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
	// 每页最大条目数，最大值50。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标。如果是首次调用，可不传；如果非首次调用，该参数传上次调用时返回的nextToken。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 搜索指定的文件类型。支持的类型有：1-文档；2-表格；3-脑图；4-演示；5-白板。
	SearchFileType *int32 `json:"searchFileType,omitempty" xml:"searchFileType,omitempty"`
	// 知识库id，在指定的知识库中搜索。
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
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

func (s *SearchRequestDentryRequest) SetSearchFileType(v int32) *SearchRequestDentryRequest {
	s.SearchFileType = &v
	return s
}

func (s *SearchRequestDentryRequest) SetSpaceId(v string) *SearchRequestDentryRequest {
	s.SpaceId = &v
	return s
}

type SearchRequestSpaceRequest struct {
	// 每页最大条目数，最大值50。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标。如果是首次调用，可不传；如果非首次调用，该参数传上次调用时返回的nextToken。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

type SearchResponseBody struct {
	// 节点搜索结果。
	DentryResult *SearchResponseBodyDentryResult `json:"dentryResult,omitempty" xml:"dentryResult,omitempty" type:"Struct"`
	// 知识库搜索结果。
	SpaceResult *SearchResponseBodySpaceResult `json:"spaceResult,omitempty" xml:"spaceResult,omitempty" type:"Struct"`
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
	// 是否还有更多数据。
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 搜索命中的节点列表。
	Items []*SearchResponseBodyDentryResultItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 分页游标。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// 如果内容命中了关键词，会返回该部分内容，带高亮。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 创建信息。
	Creation *OpenActionModel `json:"creation,omitempty" xml:"creation,omitempty"`
	// 节点id。
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// 节点唯一标识。
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// 文件名扩展。
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 节点图标url。
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// 最后修改信息。
	LastEdition *OpenActionModel `json:"lastEdition,omitempty" xml:"lastEdition,omitempty"`
	// 节点名称，如果命中了关键词，会带有高亮。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 节点原始名称，不带高亮。
	OriginName *string `json:"originName,omitempty" xml:"originName,omitempty"`
	// 节点的路径。
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 文件类型。1-文档；2-表格；3-脑图；4-演示；5-白板；6-office文字；7-office表格；8-office ppt；10-多维表格；11-文本；12-图片；13-视频；14-音频；15-压缩文件；16-其他。
	SearchFileType *int32 `json:"searchFileType,omitempty" xml:"searchFileType,omitempty"`
	// 节点所属的知识库id。
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 节点访问url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
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

func (s *SearchResponseBodyDentryResultItems) SetSearchFileType(v int32) *SearchResponseBodyDentryResultItems {
	s.SearchFileType = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetSpaceId(v string) *SearchResponseBodyDentryResultItems {
	s.SpaceId = &v
	return s
}

func (s *SearchResponseBodyDentryResultItems) SetUrl(v string) *SearchResponseBodyDentryResultItems {
	s.Url = &v
	return s
}

type SearchResponseBodySpaceResult struct {
	// 是否还有更多数据。
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 搜索命中的知识库列表。
	Items []*SearchResponseBodySpaceResultItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 分页游标。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// 知识库名称，如果命中了关键词，会带有高亮。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 知识库原始名称，不带高亮。
	OriginName *string `json:"originName,omitempty" xml:"originName,omitempty"`
	// 知识库id。
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 知识库访问url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SearchResponseBodySpaceResultItems) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodySpaceResultItems) GoString() string {
	return s.String()
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

func (s *SearchResponseBodySpaceResultItems) SetUrl(v string) *SearchResponseBodySpaceResultItems {
	s.Url = &v
	return s
}

type SearchResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchResponse) SetBody(v *SearchResponseBody) *SearchResponse {
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

func (client *Client) CopyDentryWithOptions(spaceId *string, dentryId *string, request *CopyDentryRequest, headers *CopyDentryHeaders, runtime *util.RuntimeOptions) (_result *CopyDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &CopyDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("CopyDentry"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v2.0/doc/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/copy"), tea.String("json"), req, runtime)
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

func (client *Client) CreateDentryWithOptions(spaceId *string, request *CreateDentryRequest, headers *CreateDentryHeaders, runtime *util.RuntimeOptions) (_result *CreateDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
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
	_result = &CreateDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateDentry"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v2.0/doc/spaces/"+tea.StringValue(spaceId)+"/dentries"), tea.String("json"), req, runtime)
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

func (client *Client) GetSpaceDirectoriesWithOptions(spaceId *string, request *GetSpaceDirectoriesRequest, headers *GetSpaceDirectoriesHeaders, runtime *util.RuntimeOptions) (_result *GetSpaceDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
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
	_result = &GetSpaceDirectoriesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSpaceDirectories"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v2.0/doc/spaces/"+tea.StringValue(spaceId)+"/directories"), tea.String("json"), req, runtime)
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
	_result = &GetUserInfoByOpenTokenResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserInfoByOpenToken"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v2.0/doc/userInfos"), tea.String("json"), req, runtime)
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

func (client *Client) MoveDentryWithOptions(spaceId *string, dentryId *string, request *MoveDentryRequest, headers *MoveDentryHeaders, runtime *util.RuntimeOptions) (_result *MoveDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &MoveDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("MoveDentry"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v2.0/doc/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/move"), tea.String("json"), req, runtime)
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

func (client *Client) QueryDentryWithOptions(spaceId *string, dentryId *string, request *QueryDentryRequest, headers *QueryDentryHeaders, runtime *util.RuntimeOptions) (_result *QueryDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &QueryDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDentry"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v2.0/doc/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)), tea.String("json"), req, runtime)
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

func (client *Client) QueryMineSpaceWithOptions(unionId *string, headers *QueryMineSpaceHeaders, runtime *util.RuntimeOptions) (_result *QueryMineSpaceResponse, _err error) {
	unionId = openapiutil.GetEncodeParam(unionId)
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
	_result = &QueryMineSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryMineSpace"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v2.0/doc/spaces/users/"+tea.StringValue(unionId)+"/mine"), tea.String("json"), req, runtime)
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
	_result = &QueryRecentListResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryRecentList"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v2.0/doc/spaces/docs/recent"), tea.String("json"), req, runtime)
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

func (client *Client) QuerySpaceWithOptions(spaceId *string, request *QuerySpaceRequest, headers *QuerySpaceHeaders, runtime *util.RuntimeOptions) (_result *QuerySpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
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
	_result = &QuerySpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySpace"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v2.0/doc/spaces/"+tea.StringValue(spaceId)), tea.String("json"), req, runtime)
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
	_result = &RelatedSpacesResponse{}
	_body, _err := client.DoROARequest(tea.String("RelatedSpaces"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v2.0/doc/relations/spaces"), tea.String("json"), req, runtime)
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

func (client *Client) RenameDentryWithOptions(spaceId *string, dentryId *string, request *RenameDentryRequest, headers *RenameDentryHeaders, runtime *util.RuntimeOptions) (_result *RenameDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &RenameDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("RenameDentry"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v2.0/doc/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/rename"), tea.String("json"), req, runtime)
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

func (client *Client) SearchWithOptions(request *SearchRequest, headers *SearchHeaders, runtime *util.RuntimeOptions) (_result *SearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.DentryRequest))) {
		body["dentryRequest"] = request.DentryRequest
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SpaceRequest))) {
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
	_result = &SearchResponse{}
	_body, _err := client.DoROARequest(tea.String("Search"), tea.String("doc_2.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v2.0/doc/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
