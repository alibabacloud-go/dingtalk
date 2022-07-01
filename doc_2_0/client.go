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

type DentryOpenVO struct {
	// 内容类型。alidoc-钉钉文档；link-快捷方式；archive-压缩包；document-文件。
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 创建时间。
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 创建者。
	Creator *DentryOpenVOCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
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
	Space *SpaceOpenVO `json:"space,omitempty" xml:"space,omitempty"`
	// 知识库id。
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 更新时间。
	UpdatedTime *int64 `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	// 更新人。
	Updater *DentryOpenVOUpdater `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	// 节点访问url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 访问者对当前节点的权限等信息。
	VisitorInfo *DentryOpenVOVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
}

func (s DentryOpenVO) String() string {
	return tea.Prettify(s)
}

func (s DentryOpenVO) GoString() string {
	return s.String()
}

func (s *DentryOpenVO) SetContentType(v string) *DentryOpenVO {
	s.ContentType = &v
	return s
}

func (s *DentryOpenVO) SetCreatedTime(v int64) *DentryOpenVO {
	s.CreatedTime = &v
	return s
}

func (s *DentryOpenVO) SetCreator(v *DentryOpenVOCreator) *DentryOpenVO {
	s.Creator = v
	return s
}

func (s *DentryOpenVO) SetDentryId(v string) *DentryOpenVO {
	s.DentryId = &v
	return s
}

func (s *DentryOpenVO) SetDentryType(v string) *DentryOpenVO {
	s.DentryType = &v
	return s
}

func (s *DentryOpenVO) SetDentryUuid(v string) *DentryOpenVO {
	s.DentryUuid = &v
	return s
}

func (s *DentryOpenVO) SetDocKey(v string) *DentryOpenVO {
	s.DocKey = &v
	return s
}

func (s *DentryOpenVO) SetExtension(v string) *DentryOpenVO {
	s.Extension = &v
	return s
}

func (s *DentryOpenVO) SetHasChildren(v bool) *DentryOpenVO {
	s.HasChildren = &v
	return s
}

func (s *DentryOpenVO) SetLinkSourceInfo(v *LinkSourceInfo) *DentryOpenVO {
	s.LinkSourceInfo = v
	return s
}

func (s *DentryOpenVO) SetName(v string) *DentryOpenVO {
	s.Name = &v
	return s
}

func (s *DentryOpenVO) SetPath(v string) *DentryOpenVO {
	s.Path = &v
	return s
}

func (s *DentryOpenVO) SetSpace(v *SpaceOpenVO) *DentryOpenVO {
	s.Space = v
	return s
}

func (s *DentryOpenVO) SetSpaceId(v string) *DentryOpenVO {
	s.SpaceId = &v
	return s
}

func (s *DentryOpenVO) SetUpdatedTime(v int64) *DentryOpenVO {
	s.UpdatedTime = &v
	return s
}

func (s *DentryOpenVO) SetUpdater(v *DentryOpenVOUpdater) *DentryOpenVO {
	s.Updater = v
	return s
}

func (s *DentryOpenVO) SetUrl(v string) *DentryOpenVO {
	s.Url = &v
	return s
}

func (s *DentryOpenVO) SetVisitorInfo(v *DentryOpenVOVisitorInfo) *DentryOpenVO {
	s.VisitorInfo = v
	return s
}

type DentryOpenVOCreator struct {
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DentryOpenVOCreator) String() string {
	return tea.Prettify(s)
}

func (s DentryOpenVOCreator) GoString() string {
	return s.String()
}

func (s *DentryOpenVOCreator) SetName(v string) *DentryOpenVOCreator {
	s.Name = &v
	return s
}

func (s *DentryOpenVOCreator) SetUnionId(v string) *DentryOpenVOCreator {
	s.UnionId = &v
	return s
}

type DentryOpenVOUpdater struct {
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DentryOpenVOUpdater) String() string {
	return tea.Prettify(s)
}

func (s DentryOpenVOUpdater) GoString() string {
	return s.String()
}

func (s *DentryOpenVOUpdater) SetName(v string) *DentryOpenVOUpdater {
	s.Name = &v
	return s
}

func (s *DentryOpenVOUpdater) SetUnionId(v string) *DentryOpenVOUpdater {
	s.UnionId = &v
	return s
}

type DentryOpenVOVisitorInfo struct {
	// 节点的操作列表。
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	// 空间的操作列表。
	SpaceActions []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
}

func (s DentryOpenVOVisitorInfo) String() string {
	return tea.Prettify(s)
}

func (s DentryOpenVOVisitorInfo) GoString() string {
	return s.String()
}

func (s *DentryOpenVOVisitorInfo) SetDentryActions(v []*string) *DentryOpenVOVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *DentryOpenVOVisitorInfo) SetSpaceActions(v []*string) *DentryOpenVOVisitorInfo {
	s.SpaceActions = v
	return s
}

type DentryOpenVOResult struct {
	// 内容类型。alidoc-钉钉文档；link-快捷方式；archive-压缩包；document-文件。
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 创建时间。
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 创建者。
	Creator *DentryOpenVOResultCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
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
	Space *SpaceOpenVO `json:"space,omitempty" xml:"space,omitempty"`
	// 知识库id。
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 更新时间。
	UpdatedTime *int64 `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	// 更新人。
	Updater *DentryOpenVOResultUpdater `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	// 节点访问url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 访问者对当前节点的权限等信息。
	VisitorInfo *DentryOpenVOResultVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
}

func (s DentryOpenVOResult) String() string {
	return tea.Prettify(s)
}

func (s DentryOpenVOResult) GoString() string {
	return s.String()
}

func (s *DentryOpenVOResult) SetContentType(v string) *DentryOpenVOResult {
	s.ContentType = &v
	return s
}

func (s *DentryOpenVOResult) SetCreatedTime(v int64) *DentryOpenVOResult {
	s.CreatedTime = &v
	return s
}

func (s *DentryOpenVOResult) SetCreator(v *DentryOpenVOResultCreator) *DentryOpenVOResult {
	s.Creator = v
	return s
}

func (s *DentryOpenVOResult) SetDentryId(v string) *DentryOpenVOResult {
	s.DentryId = &v
	return s
}

func (s *DentryOpenVOResult) SetDentryType(v string) *DentryOpenVOResult {
	s.DentryType = &v
	return s
}

func (s *DentryOpenVOResult) SetDentryUuid(v string) *DentryOpenVOResult {
	s.DentryUuid = &v
	return s
}

func (s *DentryOpenVOResult) SetDocKey(v string) *DentryOpenVOResult {
	s.DocKey = &v
	return s
}

func (s *DentryOpenVOResult) SetExtension(v string) *DentryOpenVOResult {
	s.Extension = &v
	return s
}

func (s *DentryOpenVOResult) SetHasChildren(v bool) *DentryOpenVOResult {
	s.HasChildren = &v
	return s
}

func (s *DentryOpenVOResult) SetLinkSourceInfo(v *LinkSourceInfo) *DentryOpenVOResult {
	s.LinkSourceInfo = v
	return s
}

func (s *DentryOpenVOResult) SetName(v string) *DentryOpenVOResult {
	s.Name = &v
	return s
}

func (s *DentryOpenVOResult) SetPath(v string) *DentryOpenVOResult {
	s.Path = &v
	return s
}

func (s *DentryOpenVOResult) SetSpace(v *SpaceOpenVO) *DentryOpenVOResult {
	s.Space = v
	return s
}

func (s *DentryOpenVOResult) SetSpaceId(v string) *DentryOpenVOResult {
	s.SpaceId = &v
	return s
}

func (s *DentryOpenVOResult) SetUpdatedTime(v int64) *DentryOpenVOResult {
	s.UpdatedTime = &v
	return s
}

func (s *DentryOpenVOResult) SetUpdater(v *DentryOpenVOResultUpdater) *DentryOpenVOResult {
	s.Updater = v
	return s
}

func (s *DentryOpenVOResult) SetUrl(v string) *DentryOpenVOResult {
	s.Url = &v
	return s
}

func (s *DentryOpenVOResult) SetVisitorInfo(v *DentryOpenVOResultVisitorInfo) *DentryOpenVOResult {
	s.VisitorInfo = v
	return s
}

type DentryOpenVOResultCreator struct {
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DentryOpenVOResultCreator) String() string {
	return tea.Prettify(s)
}

func (s DentryOpenVOResultCreator) GoString() string {
	return s.String()
}

func (s *DentryOpenVOResultCreator) SetName(v string) *DentryOpenVOResultCreator {
	s.Name = &v
	return s
}

func (s *DentryOpenVOResultCreator) SetUnionId(v string) *DentryOpenVOResultCreator {
	s.UnionId = &v
	return s
}

type DentryOpenVOResultUpdater struct {
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DentryOpenVOResultUpdater) String() string {
	return tea.Prettify(s)
}

func (s DentryOpenVOResultUpdater) GoString() string {
	return s.String()
}

func (s *DentryOpenVOResultUpdater) SetName(v string) *DentryOpenVOResultUpdater {
	s.Name = &v
	return s
}

func (s *DentryOpenVOResultUpdater) SetUnionId(v string) *DentryOpenVOResultUpdater {
	s.UnionId = &v
	return s
}

type DentryOpenVOResultVisitorInfo struct {
	// 节点的操作列表。
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	// 空间的操作列表。
	SpaceActions []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
}

func (s DentryOpenVOResultVisitorInfo) String() string {
	return tea.Prettify(s)
}

func (s DentryOpenVOResultVisitorInfo) GoString() string {
	return s.String()
}

func (s *DentryOpenVOResultVisitorInfo) SetDentryActions(v []*string) *DentryOpenVOResultVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *DentryOpenVOResultVisitorInfo) SetSpaceActions(v []*string) *DentryOpenVOResultVisitorInfo {
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

type SpaceOpenVO struct {
	// 知识库id。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 知识库名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 知识库所有者。
	Owner *SpaceOpenVOOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// 知识库访问url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 访问者对当前知识库的权限等信息。
	VisitorInfo *SpaceOpenVOVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
}

func (s SpaceOpenVO) String() string {
	return tea.Prettify(s)
}

func (s SpaceOpenVO) GoString() string {
	return s.String()
}

func (s *SpaceOpenVO) SetId(v string) *SpaceOpenVO {
	s.Id = &v
	return s
}

func (s *SpaceOpenVO) SetName(v string) *SpaceOpenVO {
	s.Name = &v
	return s
}

func (s *SpaceOpenVO) SetOwner(v *SpaceOpenVOOwner) *SpaceOpenVO {
	s.Owner = v
	return s
}

func (s *SpaceOpenVO) SetUrl(v string) *SpaceOpenVO {
	s.Url = &v
	return s
}

func (s *SpaceOpenVO) SetVisitorInfo(v *SpaceOpenVOVisitorInfo) *SpaceOpenVO {
	s.VisitorInfo = v
	return s
}

type SpaceOpenVOOwner struct {
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SpaceOpenVOOwner) String() string {
	return tea.Prettify(s)
}

func (s SpaceOpenVOOwner) GoString() string {
	return s.String()
}

func (s *SpaceOpenVOOwner) SetName(v string) *SpaceOpenVOOwner {
	s.Name = &v
	return s
}

func (s *SpaceOpenVOOwner) SetUnionId(v string) *SpaceOpenVOOwner {
	s.UnionId = &v
	return s
}

type SpaceOpenVOVisitorInfo struct {
	// 节点的操作列表。
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	// 空间的操作列表。
	SpaceActions []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
}

func (s SpaceOpenVOVisitorInfo) String() string {
	return tea.Prettify(s)
}

func (s SpaceOpenVOVisitorInfo) GoString() string {
	return s.String()
}

func (s *SpaceOpenVOVisitorInfo) SetDentryActions(v []*string) *SpaceOpenVOVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *SpaceOpenVOVisitorInfo) SetSpaceActions(v []*string) *SpaceOpenVOVisitorInfo {
	s.SpaceActions = v
	return s
}

type SpaceOpenVOResult struct {
	// 知识库id。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 知识库名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 知识库所有者。
	Owner *SpaceOpenVOResultOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// 知识库访问url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 访问者对当前知识库的权限等信息。
	VisitorInfo *SpaceOpenVOResultVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
}

func (s SpaceOpenVOResult) String() string {
	return tea.Prettify(s)
}

func (s SpaceOpenVOResult) GoString() string {
	return s.String()
}

func (s *SpaceOpenVOResult) SetId(v string) *SpaceOpenVOResult {
	s.Id = &v
	return s
}

func (s *SpaceOpenVOResult) SetName(v string) *SpaceOpenVOResult {
	s.Name = &v
	return s
}

func (s *SpaceOpenVOResult) SetOwner(v *SpaceOpenVOResultOwner) *SpaceOpenVOResult {
	s.Owner = v
	return s
}

func (s *SpaceOpenVOResult) SetUrl(v string) *SpaceOpenVOResult {
	s.Url = &v
	return s
}

func (s *SpaceOpenVOResult) SetVisitorInfo(v *SpaceOpenVOResultVisitorInfo) *SpaceOpenVOResult {
	s.VisitorInfo = v
	return s
}

type SpaceOpenVOResultOwner struct {
	// 用户名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户unionId。
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SpaceOpenVOResultOwner) String() string {
	return tea.Prettify(s)
}

func (s SpaceOpenVOResultOwner) GoString() string {
	return s.String()
}

func (s *SpaceOpenVOResultOwner) SetName(v string) *SpaceOpenVOResultOwner {
	s.Name = &v
	return s
}

func (s *SpaceOpenVOResultOwner) SetUnionId(v string) *SpaceOpenVOResultOwner {
	s.UnionId = &v
	return s
}

type SpaceOpenVOResultVisitorInfo struct {
	// 节点的操作列表。
	DentryActions []*string `json:"dentryActions,omitempty" xml:"dentryActions,omitempty" type:"Repeated"`
	// 空间的操作列表。
	SpaceActions []*string `json:"spaceActions,omitempty" xml:"spaceActions,omitempty" type:"Repeated"`
}

func (s SpaceOpenVOResultVisitorInfo) String() string {
	return tea.Prettify(s)
}

func (s SpaceOpenVOResultVisitorInfo) GoString() string {
	return s.String()
}

func (s *SpaceOpenVOResultVisitorInfo) SetDentryActions(v []*string) *SpaceOpenVOResultVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *SpaceOpenVOResultVisitorInfo) SetSpaceActions(v []*string) *SpaceOpenVOResultVisitorInfo {
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
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DentryOpenVOResult `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CopyDentryResponse) SetBody(v *DentryOpenVOResult) *CopyDentryResponse {
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
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DentryOpenVOResult `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDentryResponse) SetBody(v *DentryOpenVOResult) *CreateDentryResponse {
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
	Children []*DentryOpenVO `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
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

func (s *GetSpaceDirectoriesResponseBody) SetChildren(v []*DentryOpenVO) *GetSpaceDirectoriesResponseBody {
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
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DentryOpenVOResult `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *MoveDentryResponse) SetBody(v *DentryOpenVOResult) *MoveDentryResponse {
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
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DentryOpenVOResult `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDentryResponse) SetBody(v *DentryOpenVOResult) *QueryDentryResponse {
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
	Body    *SpaceOpenVOResult `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryMineSpaceResponse) SetBody(v *SpaceOpenVOResult) *QueryMineSpaceResponse {
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
	Dentry *DentryOpenVO `json:"dentry,omitempty" xml:"dentry,omitempty"`
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

func (s *QueryRecentListResponseBodyRecentList) SetDentry(v *DentryOpenVO) *QueryRecentListResponseBodyRecentList {
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
	Body    *SpaceOpenVOResult `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QuerySpaceResponse) SetBody(v *SpaceOpenVOResult) *QuerySpaceResponse {
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
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DentryOpenVOResult `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RenameDentryResponse) SetBody(v *DentryOpenVOResult) *RenameDentryResponse {
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
	// 文件类型。1-文档；2-表格；3-脑图；4-演示；5-白板；6-office文字；7-office表格；8-office ppt；10-多维表格；11-文本；12-图片；13-视频；14-音频；15-压缩文件；16-其他。
	FileType *int32 `json:"fileType,omitempty" xml:"fileType,omitempty"`
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

func (s *SearchResponseBodyDentryResultItems) SetFileType(v int32) *SearchResponseBodyDentryResultItems {
	s.FileType = &v
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
