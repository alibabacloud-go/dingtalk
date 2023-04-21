// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package pedia_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type PediaWordsAddHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PediaWordsAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsAddHeaders) GoString() string {
	return s.String()
}

func (s *PediaWordsAddHeaders) SetCommonHeaders(v map[string]*string) *PediaWordsAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PediaWordsAddHeaders) SetXAcsDingtalkAccessToken(v string) *PediaWordsAddHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PediaWordsAddRequest struct {
	// 联系人列表
	//
	ContactList []*PediaWordsAddRequestContactList `json:"contactList,omitempty" xml:"contactList,omitempty" type:"Repeated"`
	// 高亮的别名，从别名中选取，不在别名列表中不展示
	HighLightWordAlias []*string `json:"highLightWordAlias,omitempty" xml:"highLightWordAlias,omitempty" type:"Repeated"`
	// 相关图片
	//
	PicList []*PediaWordsAddRequestPicList `json:"picList,omitempty" xml:"picList,omitempty" type:"Repeated"`
	// 相关文档，支持钉钉在线文档
	RelatedDoc []*PediaWordsAddRequestRelatedDoc `json:"relatedDoc,omitempty" xml:"relatedDoc,omitempty" type:"Repeated"`
	// 相关链接
	RelatedLink []*PediaWordsAddRequestRelatedLink `json:"relatedLink,omitempty" xml:"relatedLink,omitempty" type:"Repeated"`
	// 组织对应的员工编号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 词条的别名，多个名字的时候可以添加
	//
	WordAlias []*string `json:"wordAlias,omitempty" xml:"wordAlias,omitempty" type:"Repeated"`
	// 新增词条的名称
	WordName *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
	// 词条释义，针对词条的描述内容
	//
	WordParaphrase *string `json:"wordParaphrase,omitempty" xml:"wordParaphrase,omitempty"`
}

func (s PediaWordsAddRequest) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsAddRequest) GoString() string {
	return s.String()
}

func (s *PediaWordsAddRequest) SetContactList(v []*PediaWordsAddRequestContactList) *PediaWordsAddRequest {
	s.ContactList = v
	return s
}

func (s *PediaWordsAddRequest) SetHighLightWordAlias(v []*string) *PediaWordsAddRequest {
	s.HighLightWordAlias = v
	return s
}

func (s *PediaWordsAddRequest) SetPicList(v []*PediaWordsAddRequestPicList) *PediaWordsAddRequest {
	s.PicList = v
	return s
}

func (s *PediaWordsAddRequest) SetRelatedDoc(v []*PediaWordsAddRequestRelatedDoc) *PediaWordsAddRequest {
	s.RelatedDoc = v
	return s
}

func (s *PediaWordsAddRequest) SetRelatedLink(v []*PediaWordsAddRequestRelatedLink) *PediaWordsAddRequest {
	s.RelatedLink = v
	return s
}

func (s *PediaWordsAddRequest) SetUserId(v string) *PediaWordsAddRequest {
	s.UserId = &v
	return s
}

func (s *PediaWordsAddRequest) SetWordAlias(v []*string) *PediaWordsAddRequest {
	s.WordAlias = v
	return s
}

func (s *PediaWordsAddRequest) SetWordName(v string) *PediaWordsAddRequest {
	s.WordName = &v
	return s
}

func (s *PediaWordsAddRequest) SetWordParaphrase(v string) *PediaWordsAddRequest {
	s.WordParaphrase = &v
	return s
}

type PediaWordsAddRequestContactList struct {
	// 头像地址也可以忽略
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// 名称可以忽略
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// 联系人的组织员工编号，开放平台员工信息接口获取userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PediaWordsAddRequestContactList) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsAddRequestContactList) GoString() string {
	return s.String()
}

func (s *PediaWordsAddRequestContactList) SetAvatarMediaId(v string) *PediaWordsAddRequestContactList {
	s.AvatarMediaId = &v
	return s
}

func (s *PediaWordsAddRequestContactList) SetNickName(v string) *PediaWordsAddRequestContactList {
	s.NickName = &v
	return s
}

func (s *PediaWordsAddRequestContactList) SetUserId(v string) *PediaWordsAddRequestContactList {
	s.UserId = &v
	return s
}

type PediaWordsAddRequestPicList struct {
	// 图片的HTTP地址
	//
	MediaIdUrl *string `json:"mediaIdUrl,omitempty" xml:"mediaIdUrl,omitempty"`
}

func (s PediaWordsAddRequestPicList) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsAddRequestPicList) GoString() string {
	return s.String()
}

func (s *PediaWordsAddRequestPicList) SetMediaIdUrl(v string) *PediaWordsAddRequestPicList {
	s.MediaIdUrl = &v
	return s
}

type PediaWordsAddRequestRelatedDoc struct {
	// 文档地址
	//
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 文档名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文档类型，adoc或者asheet字段
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PediaWordsAddRequestRelatedDoc) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsAddRequestRelatedDoc) GoString() string {
	return s.String()
}

func (s *PediaWordsAddRequestRelatedDoc) SetLink(v string) *PediaWordsAddRequestRelatedDoc {
	s.Link = &v
	return s
}

func (s *PediaWordsAddRequestRelatedDoc) SetName(v string) *PediaWordsAddRequestRelatedDoc {
	s.Name = &v
	return s
}

func (s *PediaWordsAddRequestRelatedDoc) SetType(v string) *PediaWordsAddRequestRelatedDoc {
	s.Type = &v
	return s
}

type PediaWordsAddRequestRelatedLink struct {
	// 链接地址
	//
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 链接名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s PediaWordsAddRequestRelatedLink) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsAddRequestRelatedLink) GoString() string {
	return s.String()
}

func (s *PediaWordsAddRequestRelatedLink) SetLink(v string) *PediaWordsAddRequestRelatedLink {
	s.Link = &v
	return s
}

func (s *PediaWordsAddRequestRelatedLink) SetName(v string) *PediaWordsAddRequestRelatedLink {
	s.Name = &v
	return s
}

type PediaWordsAddResponseBody struct {
	// 返回结果
	// false，失败
	// true，成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 插入成功后的编号主键ID
	//
	Uuid *int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s PediaWordsAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsAddResponseBody) GoString() string {
	return s.String()
}

func (s *PediaWordsAddResponseBody) SetSuccess(v bool) *PediaWordsAddResponseBody {
	s.Success = &v
	return s
}

func (s *PediaWordsAddResponseBody) SetUuid(v int64) *PediaWordsAddResponseBody {
	s.Uuid = &v
	return s
}

type PediaWordsAddResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PediaWordsAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PediaWordsAddResponse) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsAddResponse) GoString() string {
	return s.String()
}

func (s *PediaWordsAddResponse) SetHeaders(v map[string]*string) *PediaWordsAddResponse {
	s.Headers = v
	return s
}

func (s *PediaWordsAddResponse) SetBody(v *PediaWordsAddResponseBody) *PediaWordsAddResponse {
	s.Body = v
	return s
}

type PediaWordsApproveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PediaWordsApproveHeaders) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsApproveHeaders) GoString() string {
	return s.String()
}

func (s *PediaWordsApproveHeaders) SetCommonHeaders(v map[string]*string) *PediaWordsApproveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PediaWordsApproveHeaders) SetXAcsDingtalkAccessToken(v string) *PediaWordsApproveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PediaWordsApproveRequest struct {
	// 拒绝的原因，可选
	ApproveReason *string `json:"approveReason,omitempty" xml:"approveReason,omitempty"`
	// 审核的结果，1通过0代表拒绝
	ApproveStatus *string `json:"approveStatus,omitempty" xml:"approveStatus,omitempty"`
	// 当前内部群是否高亮
	ImHighLight *bool `json:"imHighLight,omitempty" xml:"imHighLight,omitempty"`
	// 服务群是否高亮
	SimHighLight *bool `json:"simHighLight,omitempty" xml:"simHighLight,omitempty"`
	// 操作人的组织员工编号，开放平台员工信息接口获取userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 当前审核的词条的主键编号
	Uuid *int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s PediaWordsApproveRequest) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsApproveRequest) GoString() string {
	return s.String()
}

func (s *PediaWordsApproveRequest) SetApproveReason(v string) *PediaWordsApproveRequest {
	s.ApproveReason = &v
	return s
}

func (s *PediaWordsApproveRequest) SetApproveStatus(v string) *PediaWordsApproveRequest {
	s.ApproveStatus = &v
	return s
}

func (s *PediaWordsApproveRequest) SetImHighLight(v bool) *PediaWordsApproveRequest {
	s.ImHighLight = &v
	return s
}

func (s *PediaWordsApproveRequest) SetSimHighLight(v bool) *PediaWordsApproveRequest {
	s.SimHighLight = &v
	return s
}

func (s *PediaWordsApproveRequest) SetUserId(v string) *PediaWordsApproveRequest {
	s.UserId = &v
	return s
}

func (s *PediaWordsApproveRequest) SetUuid(v int64) *PediaWordsApproveRequest {
	s.Uuid = &v
	return s
}

type PediaWordsApproveResponseBody struct {
	// 返回结果
	// false，失败
	// true，成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PediaWordsApproveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsApproveResponseBody) GoString() string {
	return s.String()
}

func (s *PediaWordsApproveResponseBody) SetSuccess(v bool) *PediaWordsApproveResponseBody {
	s.Success = &v
	return s
}

type PediaWordsApproveResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PediaWordsApproveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PediaWordsApproveResponse) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsApproveResponse) GoString() string {
	return s.String()
}

func (s *PediaWordsApproveResponse) SetHeaders(v map[string]*string) *PediaWordsApproveResponse {
	s.Headers = v
	return s
}

func (s *PediaWordsApproveResponse) SetBody(v *PediaWordsApproveResponseBody) *PediaWordsApproveResponse {
	s.Body = v
	return s
}

type PediaWordsDeleteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PediaWordsDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsDeleteHeaders) GoString() string {
	return s.String()
}

func (s *PediaWordsDeleteHeaders) SetCommonHeaders(v map[string]*string) *PediaWordsDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PediaWordsDeleteHeaders) SetXAcsDingtalkAccessToken(v string) *PediaWordsDeleteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PediaWordsDeleteRequest struct {
	// 当前操作用户编号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 当前需要删除的词条主键编号
	Uuid *int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s PediaWordsDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsDeleteRequest) GoString() string {
	return s.String()
}

func (s *PediaWordsDeleteRequest) SetUserId(v string) *PediaWordsDeleteRequest {
	s.UserId = &v
	return s
}

func (s *PediaWordsDeleteRequest) SetUuid(v int64) *PediaWordsDeleteRequest {
	s.Uuid = &v
	return s
}

type PediaWordsDeleteResponseBody struct {
	// 返回结果
	// false，失败
	// true，成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 删除成功，返回删除的uuid
	Uuid *int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s PediaWordsDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *PediaWordsDeleteResponseBody) SetSuccess(v bool) *PediaWordsDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *PediaWordsDeleteResponseBody) SetUuid(v int64) *PediaWordsDeleteResponseBody {
	s.Uuid = &v
	return s
}

type PediaWordsDeleteResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PediaWordsDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PediaWordsDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsDeleteResponse) GoString() string {
	return s.String()
}

func (s *PediaWordsDeleteResponse) SetHeaders(v map[string]*string) *PediaWordsDeleteResponse {
	s.Headers = v
	return s
}

func (s *PediaWordsDeleteResponse) SetBody(v *PediaWordsDeleteResponseBody) *PediaWordsDeleteResponse {
	s.Body = v
	return s
}

type PediaWordsQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PediaWordsQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsQueryHeaders) GoString() string {
	return s.String()
}

func (s *PediaWordsQueryHeaders) SetCommonHeaders(v map[string]*string) *PediaWordsQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PediaWordsQueryHeaders) SetXAcsDingtalkAccessToken(v string) *PediaWordsQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PediaWordsQueryRequest struct {
	// 操作用户编号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 查询主键编号
	Uuid *int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s PediaWordsQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsQueryRequest) GoString() string {
	return s.String()
}

func (s *PediaWordsQueryRequest) SetUserId(v string) *PediaWordsQueryRequest {
	s.UserId = &v
	return s
}

func (s *PediaWordsQueryRequest) SetUuid(v int64) *PediaWordsQueryRequest {
	s.Uuid = &v
	return s
}

type PediaWordsQueryResponseBody struct {
	// 返回词条具体对象
	//
	Data *PediaWordsQueryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 返回结果
	// false，失败
	// trur，成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PediaWordsQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsQueryResponseBody) GoString() string {
	return s.String()
}

func (s *PediaWordsQueryResponseBody) SetData(v *PediaWordsQueryResponseBodyData) *PediaWordsQueryResponseBody {
	s.Data = v
	return s
}

func (s *PediaWordsQueryResponseBody) SetSuccess(v bool) *PediaWordsQueryResponseBody {
	s.Success = &v
	return s
}

type PediaWordsQueryResponseBodyData struct {
	// 相关应用
	AppLink []*PediaWordsQueryResponseBodyDataAppLink `json:"appLink,omitempty" xml:"appLink,omitempty" type:"Repeated"`
	// 审核人
	ApproveName *string `json:"approveName,omitempty" xml:"approveName,omitempty"`
	// 联系人列表
	ContactList []*PediaWordsQueryResponseBodyDataContactList `json:"contactList,omitempty" xml:"contactList,omitempty" type:"Repeated"`
	// 相关联系人
	Contacts []*string `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// 创建者
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModify *int64 `json:"gmtModify,omitempty" xml:"gmtModify,omitempty"`
	// 高亮的词条别名
	//
	HighLightWordAlias []*string `json:"highLightWordAlias,omitempty" xml:"highLightWordAlias,omitempty" type:"Repeated"`
	// 内部群是否高亮
	ImHighLight *bool `json:"imHighLight,omitempty" xml:"imHighLight,omitempty"`
	// 当为待审核词条的时候的父编号
	ParentUuid *int64                                    `json:"parentUuid,omitempty" xml:"parentUuid,omitempty"`
	PicList    []*PediaWordsQueryResponseBodyDataPicList `json:"picList,omitempty" xml:"picList,omitempty" type:"Repeated"`
	// 相关文档
	RelatedDoc []*PediaWordsQueryResponseBodyDataRelatedDoc `json:"relatedDoc,omitempty" xml:"relatedDoc,omitempty" type:"Repeated"`
	// 相关链接
	RelatedLink []*PediaWordsQueryResponseBodyDataRelatedLink `json:"relatedLink,omitempty" xml:"relatedLink,omitempty" type:"Repeated"`
	// 服务群是否高亮
	SimHighLight *bool `json:"simHighLight,omitempty" xml:"simHighLight,omitempty"`
	// 词条释义非富文本
	SimpleWordParaphrase *string `json:"simpleWordParaphrase,omitempty" xml:"simpleWordParaphrase,omitempty"`
	// 分类名称
	TagsList []*string `json:"tagsList,omitempty" xml:"tagsList,omitempty" type:"Repeated"`
	// 更新人
	UpdaterName *string `json:"updaterName,omitempty" xml:"updaterName,omitempty"`
	// 操作员工userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 词条主键ID
	Uuid *int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 词条别名
	WordAlias []*string `json:"wordAlias,omitempty" xml:"wordAlias,omitempty" type:"Repeated"`
	// 词条名称
	//
	WordName *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
	// 词条释义，富文本
	WordParaphrase *string `json:"wordParaphrase,omitempty" xml:"wordParaphrase,omitempty"`
}

func (s PediaWordsQueryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *PediaWordsQueryResponseBodyData) SetAppLink(v []*PediaWordsQueryResponseBodyDataAppLink) *PediaWordsQueryResponseBodyData {
	s.AppLink = v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetApproveName(v string) *PediaWordsQueryResponseBodyData {
	s.ApproveName = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetContactList(v []*PediaWordsQueryResponseBodyDataContactList) *PediaWordsQueryResponseBodyData {
	s.ContactList = v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetContacts(v []*string) *PediaWordsQueryResponseBodyData {
	s.Contacts = v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetCreatorName(v string) *PediaWordsQueryResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetGmtCreate(v int64) *PediaWordsQueryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetGmtModify(v int64) *PediaWordsQueryResponseBodyData {
	s.GmtModify = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetHighLightWordAlias(v []*string) *PediaWordsQueryResponseBodyData {
	s.HighLightWordAlias = v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetImHighLight(v bool) *PediaWordsQueryResponseBodyData {
	s.ImHighLight = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetParentUuid(v int64) *PediaWordsQueryResponseBodyData {
	s.ParentUuid = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetPicList(v []*PediaWordsQueryResponseBodyDataPicList) *PediaWordsQueryResponseBodyData {
	s.PicList = v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetRelatedDoc(v []*PediaWordsQueryResponseBodyDataRelatedDoc) *PediaWordsQueryResponseBodyData {
	s.RelatedDoc = v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetRelatedLink(v []*PediaWordsQueryResponseBodyDataRelatedLink) *PediaWordsQueryResponseBodyData {
	s.RelatedLink = v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetSimHighLight(v bool) *PediaWordsQueryResponseBodyData {
	s.SimHighLight = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetSimpleWordParaphrase(v string) *PediaWordsQueryResponseBodyData {
	s.SimpleWordParaphrase = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetTagsList(v []*string) *PediaWordsQueryResponseBodyData {
	s.TagsList = v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetUpdaterName(v string) *PediaWordsQueryResponseBodyData {
	s.UpdaterName = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetUserId(v string) *PediaWordsQueryResponseBodyData {
	s.UserId = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetUuid(v int64) *PediaWordsQueryResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetWordAlias(v []*string) *PediaWordsQueryResponseBodyData {
	s.WordAlias = v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetWordName(v string) *PediaWordsQueryResponseBodyData {
	s.WordName = &v
	return s
}

func (s *PediaWordsQueryResponseBodyData) SetWordParaphrase(v string) *PediaWordsQueryResponseBodyData {
	s.WordParaphrase = &v
	return s
}

type PediaWordsQueryResponseBodyDataAppLink struct {
	// 应用名称
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// 应用icon
	IconLink *string `json:"iconLink,omitempty" xml:"iconLink,omitempty"`
	// 桌面端链接
	PcLink *string `json:"pcLink,omitempty" xml:"pcLink,omitempty"`
	// 手机端链接
	PhoneLink *string `json:"phoneLink,omitempty" xml:"phoneLink,omitempty"`
}

func (s PediaWordsQueryResponseBodyDataAppLink) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsQueryResponseBodyDataAppLink) GoString() string {
	return s.String()
}

func (s *PediaWordsQueryResponseBodyDataAppLink) SetAppName(v string) *PediaWordsQueryResponseBodyDataAppLink {
	s.AppName = &v
	return s
}

func (s *PediaWordsQueryResponseBodyDataAppLink) SetIconLink(v string) *PediaWordsQueryResponseBodyDataAppLink {
	s.IconLink = &v
	return s
}

func (s *PediaWordsQueryResponseBodyDataAppLink) SetPcLink(v string) *PediaWordsQueryResponseBodyDataAppLink {
	s.PcLink = &v
	return s
}

func (s *PediaWordsQueryResponseBodyDataAppLink) SetPhoneLink(v string) *PediaWordsQueryResponseBodyDataAppLink {
	s.PhoneLink = &v
	return s
}

type PediaWordsQueryResponseBodyDataContactList struct {
	// 联系人图片办好
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// 联系人名称
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// 联系人员工编号
	//
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PediaWordsQueryResponseBodyDataContactList) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsQueryResponseBodyDataContactList) GoString() string {
	return s.String()
}

func (s *PediaWordsQueryResponseBodyDataContactList) SetAvatarMediaId(v string) *PediaWordsQueryResponseBodyDataContactList {
	s.AvatarMediaId = &v
	return s
}

func (s *PediaWordsQueryResponseBodyDataContactList) SetNickName(v string) *PediaWordsQueryResponseBodyDataContactList {
	s.NickName = &v
	return s
}

func (s *PediaWordsQueryResponseBodyDataContactList) SetUserId(v string) *PediaWordsQueryResponseBodyDataContactList {
	s.UserId = &v
	return s
}

type PediaWordsQueryResponseBodyDataPicList struct {
	// 图片HTTP地址
	MediaIdUrl *string `json:"mediaIdUrl,omitempty" xml:"mediaIdUrl,omitempty"`
}

func (s PediaWordsQueryResponseBodyDataPicList) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsQueryResponseBodyDataPicList) GoString() string {
	return s.String()
}

func (s *PediaWordsQueryResponseBodyDataPicList) SetMediaIdUrl(v string) *PediaWordsQueryResponseBodyDataPicList {
	s.MediaIdUrl = &v
	return s
}

type PediaWordsQueryResponseBodyDataRelatedDoc struct {
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 文档名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文档类型，分别为adoc或者asheet
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PediaWordsQueryResponseBodyDataRelatedDoc) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsQueryResponseBodyDataRelatedDoc) GoString() string {
	return s.String()
}

func (s *PediaWordsQueryResponseBodyDataRelatedDoc) SetLink(v string) *PediaWordsQueryResponseBodyDataRelatedDoc {
	s.Link = &v
	return s
}

func (s *PediaWordsQueryResponseBodyDataRelatedDoc) SetName(v string) *PediaWordsQueryResponseBodyDataRelatedDoc {
	s.Name = &v
	return s
}

func (s *PediaWordsQueryResponseBodyDataRelatedDoc) SetType(v string) *PediaWordsQueryResponseBodyDataRelatedDoc {
	s.Type = &v
	return s
}

type PediaWordsQueryResponseBodyDataRelatedLink struct {
	// 链接地址
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 链接名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s PediaWordsQueryResponseBodyDataRelatedLink) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsQueryResponseBodyDataRelatedLink) GoString() string {
	return s.String()
}

func (s *PediaWordsQueryResponseBodyDataRelatedLink) SetLink(v string) *PediaWordsQueryResponseBodyDataRelatedLink {
	s.Link = &v
	return s
}

func (s *PediaWordsQueryResponseBodyDataRelatedLink) SetName(v string) *PediaWordsQueryResponseBodyDataRelatedLink {
	s.Name = &v
	return s
}

type PediaWordsQueryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PediaWordsQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PediaWordsQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsQueryResponse) GoString() string {
	return s.String()
}

func (s *PediaWordsQueryResponse) SetHeaders(v map[string]*string) *PediaWordsQueryResponse {
	s.Headers = v
	return s
}

func (s *PediaWordsQueryResponse) SetBody(v *PediaWordsQueryResponseBody) *PediaWordsQueryResponse {
	s.Body = v
	return s
}

type PediaWordsSearchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PediaWordsSearchHeaders) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsSearchHeaders) GoString() string {
	return s.String()
}

func (s *PediaWordsSearchHeaders) SetCommonHeaders(v map[string]*string) *PediaWordsSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PediaWordsSearchHeaders) SetXAcsDingtalkAccessToken(v string) *PediaWordsSearchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PediaWordsSearchRequest struct {
	// 当前查询的页数，当超过总数后返回数据为空
	//
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 当前每页需要展示的数量，最大20
	//
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 当前搜索列表的状态0代表审核通过，1代表创建待审核，2代表更新待审核列表
	// 默认是0，代表获取所有审核完成的词条
	//
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 通过开放平台获取的员工编号userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 搜索关键词
	//
	WordName *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
}

func (s PediaWordsSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsSearchRequest) GoString() string {
	return s.String()
}

func (s *PediaWordsSearchRequest) SetPageNumber(v int32) *PediaWordsSearchRequest {
	s.PageNumber = &v
	return s
}

func (s *PediaWordsSearchRequest) SetPageSize(v int32) *PediaWordsSearchRequest {
	s.PageSize = &v
	return s
}

func (s *PediaWordsSearchRequest) SetStatus(v string) *PediaWordsSearchRequest {
	s.Status = &v
	return s
}

func (s *PediaWordsSearchRequest) SetUserId(v string) *PediaWordsSearchRequest {
	s.UserId = &v
	return s
}

func (s *PediaWordsSearchRequest) SetWordName(v string) *PediaWordsSearchRequest {
	s.WordName = &v
	return s
}

type PediaWordsSearchResponseBody struct {
	// 词条详情对象
	Data []*PediaWordsSearchResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PediaWordsSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsSearchResponseBody) GoString() string {
	return s.String()
}

func (s *PediaWordsSearchResponseBody) SetData(v []*PediaWordsSearchResponseBodyData) *PediaWordsSearchResponseBody {
	s.Data = v
	return s
}

func (s *PediaWordsSearchResponseBody) SetSuccess(v bool) *PediaWordsSearchResponseBody {
	s.Success = &v
	return s
}

type PediaWordsSearchResponseBodyData struct {
	// 相关应用列表
	AppLink []*PediaWordsSearchResponseBodyDataAppLink `json:"appLink,omitempty" xml:"appLink,omitempty" type:"Repeated"`
	// 审核者名称
	ApproveName *string `json:"approveName,omitempty" xml:"approveName,omitempty"`
	// 相关联系人
	ContactList []*PediaWordsSearchResponseBodyDataContactList `json:"contactList,omitempty" xml:"contactList,omitempty" type:"Repeated"`
	// 联系人列表
	Contacts []*string `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// 创建者的名称
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// 词条创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 词条修改时间
	GmtModify *int64 `json:"gmtModify,omitempty" xml:"gmtModify,omitempty"`
	// 词条中需要在聊天中被分词的别名
	HighLightWordAlias []*string `json:"highLightWordAlias,omitempty" xml:"highLightWordAlias,omitempty" type:"Repeated"`
	// 该词条内部群是否分词
	ImHighLight *bool `json:"imHighLight,omitempty" xml:"imHighLight,omitempty"`
	// 当前词条的父类ID，审核通过的该字段为空
	ParentUuid *int64 `json:"parentUuid,omitempty" xml:"parentUuid,omitempty"`
	// 相关图片
	PicList []*PediaWordsSearchResponseBodyDataPicList `json:"picList,omitempty" xml:"picList,omitempty" type:"Repeated"`
	// 相关文档链接，这里只针对钉钉在线文档，没有则忽略
	RelatedDoc []*PediaWordsSearchResponseBodyDataRelatedDoc `json:"relatedDoc,omitempty" xml:"relatedDoc,omitempty" type:"Repeated"`
	// 相关链接
	//
	RelatedLink []*PediaWordsSearchResponseBodyDataRelatedLink `json:"relatedLink,omitempty" xml:"relatedLink,omitempty" type:"Repeated"`
	// 该词条服务群是否分词
	SimHighLight *bool `json:"simHighLight,omitempty" xml:"simHighLight,omitempty"`
	// 词条非富文本释义
	SimpleWordParaphrase *string `json:"simpleWordParaphrase,omitempty" xml:"simpleWordParaphrase,omitempty"`
	// 分类列表
	TagsList []*string `json:"tagsList,omitempty" xml:"tagsList,omitempty" type:"Repeated"`
	// 更新者名称
	UpdaterName *string `json:"updaterName,omitempty" xml:"updaterName,omitempty"`
	// 员工的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 当前词条对应的主键ID
	Uuid *int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 词条别名
	WordAlias []*string `json:"wordAlias,omitempty" xml:"wordAlias,omitempty" type:"Repeated"`
	// 词条名称
	WordName *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
	// 词条富文本释义
	WordParaphrase *string `json:"wordParaphrase,omitempty" xml:"wordParaphrase,omitempty"`
}

func (s PediaWordsSearchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsSearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *PediaWordsSearchResponseBodyData) SetAppLink(v []*PediaWordsSearchResponseBodyDataAppLink) *PediaWordsSearchResponseBodyData {
	s.AppLink = v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetApproveName(v string) *PediaWordsSearchResponseBodyData {
	s.ApproveName = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetContactList(v []*PediaWordsSearchResponseBodyDataContactList) *PediaWordsSearchResponseBodyData {
	s.ContactList = v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetContacts(v []*string) *PediaWordsSearchResponseBodyData {
	s.Contacts = v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetCreatorName(v string) *PediaWordsSearchResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetGmtCreate(v int64) *PediaWordsSearchResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetGmtModify(v int64) *PediaWordsSearchResponseBodyData {
	s.GmtModify = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetHighLightWordAlias(v []*string) *PediaWordsSearchResponseBodyData {
	s.HighLightWordAlias = v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetImHighLight(v bool) *PediaWordsSearchResponseBodyData {
	s.ImHighLight = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetParentUuid(v int64) *PediaWordsSearchResponseBodyData {
	s.ParentUuid = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetPicList(v []*PediaWordsSearchResponseBodyDataPicList) *PediaWordsSearchResponseBodyData {
	s.PicList = v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetRelatedDoc(v []*PediaWordsSearchResponseBodyDataRelatedDoc) *PediaWordsSearchResponseBodyData {
	s.RelatedDoc = v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetRelatedLink(v []*PediaWordsSearchResponseBodyDataRelatedLink) *PediaWordsSearchResponseBodyData {
	s.RelatedLink = v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetSimHighLight(v bool) *PediaWordsSearchResponseBodyData {
	s.SimHighLight = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetSimpleWordParaphrase(v string) *PediaWordsSearchResponseBodyData {
	s.SimpleWordParaphrase = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetTagsList(v []*string) *PediaWordsSearchResponseBodyData {
	s.TagsList = v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetUpdaterName(v string) *PediaWordsSearchResponseBodyData {
	s.UpdaterName = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetUserId(v string) *PediaWordsSearchResponseBodyData {
	s.UserId = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetUuid(v int64) *PediaWordsSearchResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetWordAlias(v []*string) *PediaWordsSearchResponseBodyData {
	s.WordAlias = v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetWordName(v string) *PediaWordsSearchResponseBodyData {
	s.WordName = &v
	return s
}

func (s *PediaWordsSearchResponseBodyData) SetWordParaphrase(v string) *PediaWordsSearchResponseBodyData {
	s.WordParaphrase = &v
	return s
}

type PediaWordsSearchResponseBodyDataAppLink struct {
	// 应用名称
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// 应用图标
	IconLink *string `json:"iconLink,omitempty" xml:"iconLink,omitempty"`
	// PC端链接地址
	PcLink *string `json:"pcLink,omitempty" xml:"pcLink,omitempty"`
	// 手机端地址
	PhoneLink *string `json:"phoneLink,omitempty" xml:"phoneLink,omitempty"`
}

func (s PediaWordsSearchResponseBodyDataAppLink) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsSearchResponseBodyDataAppLink) GoString() string {
	return s.String()
}

func (s *PediaWordsSearchResponseBodyDataAppLink) SetAppName(v string) *PediaWordsSearchResponseBodyDataAppLink {
	s.AppName = &v
	return s
}

func (s *PediaWordsSearchResponseBodyDataAppLink) SetIconLink(v string) *PediaWordsSearchResponseBodyDataAppLink {
	s.IconLink = &v
	return s
}

func (s *PediaWordsSearchResponseBodyDataAppLink) SetPcLink(v string) *PediaWordsSearchResponseBodyDataAppLink {
	s.PcLink = &v
	return s
}

func (s *PediaWordsSearchResponseBodyDataAppLink) SetPhoneLink(v string) *PediaWordsSearchResponseBodyDataAppLink {
	s.PhoneLink = &v
	return s
}

type PediaWordsSearchResponseBodyDataContactList struct {
	// 员工头像
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// 员工的名字
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// 员工的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PediaWordsSearchResponseBodyDataContactList) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsSearchResponseBodyDataContactList) GoString() string {
	return s.String()
}

func (s *PediaWordsSearchResponseBodyDataContactList) SetAvatarMediaId(v string) *PediaWordsSearchResponseBodyDataContactList {
	s.AvatarMediaId = &v
	return s
}

func (s *PediaWordsSearchResponseBodyDataContactList) SetNickName(v string) *PediaWordsSearchResponseBodyDataContactList {
	s.NickName = &v
	return s
}

func (s *PediaWordsSearchResponseBodyDataContactList) SetUserId(v string) *PediaWordsSearchResponseBodyDataContactList {
	s.UserId = &v
	return s
}

type PediaWordsSearchResponseBodyDataPicList struct {
	// 相关图片地址
	MediaIdUrl *string `json:"mediaIdUrl,omitempty" xml:"mediaIdUrl,omitempty"`
}

func (s PediaWordsSearchResponseBodyDataPicList) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsSearchResponseBodyDataPicList) GoString() string {
	return s.String()
}

func (s *PediaWordsSearchResponseBodyDataPicList) SetMediaIdUrl(v string) *PediaWordsSearchResponseBodyDataPicList {
	s.MediaIdUrl = &v
	return s
}

type PediaWordsSearchResponseBodyDataRelatedDoc struct {
	// 当前在线文档链接地址
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 在线文档的名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 在线文档的类型，分别为adoc和asheet两个值获取文档图标
	//
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PediaWordsSearchResponseBodyDataRelatedDoc) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsSearchResponseBodyDataRelatedDoc) GoString() string {
	return s.String()
}

func (s *PediaWordsSearchResponseBodyDataRelatedDoc) SetLink(v string) *PediaWordsSearchResponseBodyDataRelatedDoc {
	s.Link = &v
	return s
}

func (s *PediaWordsSearchResponseBodyDataRelatedDoc) SetName(v string) *PediaWordsSearchResponseBodyDataRelatedDoc {
	s.Name = &v
	return s
}

func (s *PediaWordsSearchResponseBodyDataRelatedDoc) SetType(v string) *PediaWordsSearchResponseBodyDataRelatedDoc {
	s.Type = &v
	return s
}

type PediaWordsSearchResponseBodyDataRelatedLink struct {
	// 具体链接
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 链接名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 空
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PediaWordsSearchResponseBodyDataRelatedLink) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsSearchResponseBodyDataRelatedLink) GoString() string {
	return s.String()
}

func (s *PediaWordsSearchResponseBodyDataRelatedLink) SetLink(v string) *PediaWordsSearchResponseBodyDataRelatedLink {
	s.Link = &v
	return s
}

func (s *PediaWordsSearchResponseBodyDataRelatedLink) SetName(v string) *PediaWordsSearchResponseBodyDataRelatedLink {
	s.Name = &v
	return s
}

func (s *PediaWordsSearchResponseBodyDataRelatedLink) SetType(v string) *PediaWordsSearchResponseBodyDataRelatedLink {
	s.Type = &v
	return s
}

type PediaWordsSearchResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PediaWordsSearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PediaWordsSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsSearchResponse) GoString() string {
	return s.String()
}

func (s *PediaWordsSearchResponse) SetHeaders(v map[string]*string) *PediaWordsSearchResponse {
	s.Headers = v
	return s
}

func (s *PediaWordsSearchResponse) SetBody(v *PediaWordsSearchResponseBody) *PediaWordsSearchResponse {
	s.Body = v
	return s
}

type PediaWordsUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PediaWordsUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsUpdateHeaders) GoString() string {
	return s.String()
}

func (s *PediaWordsUpdateHeaders) SetCommonHeaders(v map[string]*string) *PediaWordsUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PediaWordsUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *PediaWordsUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PediaWordsUpdateRequest struct {
	// 相关应用
	AppLink []*PediaWordsUpdateRequestAppLink `json:"appLink,omitempty" xml:"appLink,omitempty" type:"Repeated"`
	// 相关联系人
	ContactList []*PediaWordsUpdateRequestContactList `json:"contactList,omitempty" xml:"contactList,omitempty" type:"Repeated"`
	// 在聊天中可命中的别名
	HighLightWordAlias []*string `json:"highLightWordAlias,omitempty" xml:"highLightWordAlias,omitempty" type:"Repeated"`
	// 相关图片
	PicList []*PediaWordsUpdateRequestPicList `json:"picList,omitempty" xml:"picList,omitempty" type:"Repeated"`
	// 相关文档，支持钉钉在线文档
	RelatedDoc []*PediaWordsUpdateRequestRelatedDoc `json:"relatedDoc,omitempty" xml:"relatedDoc,omitempty" type:"Repeated"`
	// 相关链接
	RelatedLink []*PediaWordsUpdateRequestRelatedLink `json:"relatedLink,omitempty" xml:"relatedLink,omitempty" type:"Repeated"`
	// 操作人的组织员工编号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 更新的词条编号
	Uuid *int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 词条别名
	WordAlias []*string `json:"wordAlias,omitempty" xml:"wordAlias,omitempty" type:"Repeated"`
	// 词条名称
	WordName *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
	// 词条释义
	WordParaphrase *string `json:"wordParaphrase,omitempty" xml:"wordParaphrase,omitempty"`
}

func (s PediaWordsUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsUpdateRequest) GoString() string {
	return s.String()
}

func (s *PediaWordsUpdateRequest) SetAppLink(v []*PediaWordsUpdateRequestAppLink) *PediaWordsUpdateRequest {
	s.AppLink = v
	return s
}

func (s *PediaWordsUpdateRequest) SetContactList(v []*PediaWordsUpdateRequestContactList) *PediaWordsUpdateRequest {
	s.ContactList = v
	return s
}

func (s *PediaWordsUpdateRequest) SetHighLightWordAlias(v []*string) *PediaWordsUpdateRequest {
	s.HighLightWordAlias = v
	return s
}

func (s *PediaWordsUpdateRequest) SetPicList(v []*PediaWordsUpdateRequestPicList) *PediaWordsUpdateRequest {
	s.PicList = v
	return s
}

func (s *PediaWordsUpdateRequest) SetRelatedDoc(v []*PediaWordsUpdateRequestRelatedDoc) *PediaWordsUpdateRequest {
	s.RelatedDoc = v
	return s
}

func (s *PediaWordsUpdateRequest) SetRelatedLink(v []*PediaWordsUpdateRequestRelatedLink) *PediaWordsUpdateRequest {
	s.RelatedLink = v
	return s
}

func (s *PediaWordsUpdateRequest) SetUserId(v string) *PediaWordsUpdateRequest {
	s.UserId = &v
	return s
}

func (s *PediaWordsUpdateRequest) SetUuid(v int64) *PediaWordsUpdateRequest {
	s.Uuid = &v
	return s
}

func (s *PediaWordsUpdateRequest) SetWordAlias(v []*string) *PediaWordsUpdateRequest {
	s.WordAlias = v
	return s
}

func (s *PediaWordsUpdateRequest) SetWordName(v string) *PediaWordsUpdateRequest {
	s.WordName = &v
	return s
}

func (s *PediaWordsUpdateRequest) SetWordParaphrase(v string) *PediaWordsUpdateRequest {
	s.WordParaphrase = &v
	return s
}

type PediaWordsUpdateRequestAppLink struct {
	// 应用名称
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// icon地址
	IconLink *string `json:"iconLink,omitempty" xml:"iconLink,omitempty"`
	// 电脑端地址
	PcLink *string `json:"pcLink,omitempty" xml:"pcLink,omitempty"`
	// 手机端地址
	PhoneLink *string `json:"phoneLink,omitempty" xml:"phoneLink,omitempty"`
}

func (s PediaWordsUpdateRequestAppLink) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsUpdateRequestAppLink) GoString() string {
	return s.String()
}

func (s *PediaWordsUpdateRequestAppLink) SetAppName(v string) *PediaWordsUpdateRequestAppLink {
	s.AppName = &v
	return s
}

func (s *PediaWordsUpdateRequestAppLink) SetIconLink(v string) *PediaWordsUpdateRequestAppLink {
	s.IconLink = &v
	return s
}

func (s *PediaWordsUpdateRequestAppLink) SetPcLink(v string) *PediaWordsUpdateRequestAppLink {
	s.PcLink = &v
	return s
}

func (s *PediaWordsUpdateRequestAppLink) SetPhoneLink(v string) *PediaWordsUpdateRequestAppLink {
	s.PhoneLink = &v
	return s
}

type PediaWordsUpdateRequestContactList struct {
	// 联系人图片地址，可以不传
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// 名称，可空
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// 联系人的组织员工编号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PediaWordsUpdateRequestContactList) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsUpdateRequestContactList) GoString() string {
	return s.String()
}

func (s *PediaWordsUpdateRequestContactList) SetAvatarMediaId(v string) *PediaWordsUpdateRequestContactList {
	s.AvatarMediaId = &v
	return s
}

func (s *PediaWordsUpdateRequestContactList) SetNickName(v string) *PediaWordsUpdateRequestContactList {
	s.NickName = &v
	return s
}

func (s *PediaWordsUpdateRequestContactList) SetUserId(v string) *PediaWordsUpdateRequestContactList {
	s.UserId = &v
	return s
}

type PediaWordsUpdateRequestPicList struct {
	// 图片地址，Http
	MediaIdUrl *string `json:"mediaIdUrl,omitempty" xml:"mediaIdUrl,omitempty"`
}

func (s PediaWordsUpdateRequestPicList) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsUpdateRequestPicList) GoString() string {
	return s.String()
}

func (s *PediaWordsUpdateRequestPicList) SetMediaIdUrl(v string) *PediaWordsUpdateRequestPicList {
	s.MediaIdUrl = &v
	return s
}

type PediaWordsUpdateRequestRelatedDoc struct {
	// 在线文档链接
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 文档名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文档类型，adoc或者asheet字段
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PediaWordsUpdateRequestRelatedDoc) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsUpdateRequestRelatedDoc) GoString() string {
	return s.String()
}

func (s *PediaWordsUpdateRequestRelatedDoc) SetLink(v string) *PediaWordsUpdateRequestRelatedDoc {
	s.Link = &v
	return s
}

func (s *PediaWordsUpdateRequestRelatedDoc) SetName(v string) *PediaWordsUpdateRequestRelatedDoc {
	s.Name = &v
	return s
}

func (s *PediaWordsUpdateRequestRelatedDoc) SetType(v string) *PediaWordsUpdateRequestRelatedDoc {
	s.Type = &v
	return s
}

type PediaWordsUpdateRequestRelatedLink struct {
	// 链接地址
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 链接名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s PediaWordsUpdateRequestRelatedLink) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsUpdateRequestRelatedLink) GoString() string {
	return s.String()
}

func (s *PediaWordsUpdateRequestRelatedLink) SetLink(v string) *PediaWordsUpdateRequestRelatedLink {
	s.Link = &v
	return s
}

func (s *PediaWordsUpdateRequestRelatedLink) SetName(v string) *PediaWordsUpdateRequestRelatedLink {
	s.Name = &v
	return s
}

type PediaWordsUpdateResponseBody struct {
	// 返回结果
	// false，失败
	// true,成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 更新后待审核词条编号
	Uuid *int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s PediaWordsUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *PediaWordsUpdateResponseBody) SetSuccess(v bool) *PediaWordsUpdateResponseBody {
	s.Success = &v
	return s
}

func (s *PediaWordsUpdateResponseBody) SetUuid(v int64) *PediaWordsUpdateResponseBody {
	s.Uuid = &v
	return s
}

type PediaWordsUpdateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PediaWordsUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PediaWordsUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsUpdateResponse) GoString() string {
	return s.String()
}

func (s *PediaWordsUpdateResponse) SetHeaders(v map[string]*string) *PediaWordsUpdateResponse {
	s.Headers = v
	return s
}

func (s *PediaWordsUpdateResponse) SetBody(v *PediaWordsUpdateResponseBody) *PediaWordsUpdateResponse {
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

func (client *Client) PediaWordsAdd(request *PediaWordsAddRequest) (_result *PediaWordsAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PediaWordsAddHeaders{}
	_result = &PediaWordsAddResponse{}
	_body, _err := client.PediaWordsAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PediaWordsAddWithOptions(request *PediaWordsAddRequest, headers *PediaWordsAddHeaders, runtime *util.RuntimeOptions) (_result *PediaWordsAddResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactList)) {
		body["contactList"] = request.ContactList
	}

	if !tea.BoolValue(util.IsUnset(request.HighLightWordAlias)) {
		body["highLightWordAlias"] = request.HighLightWordAlias
	}

	if !tea.BoolValue(util.IsUnset(request.PicList)) {
		body["picList"] = request.PicList
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedDoc)) {
		body["relatedDoc"] = request.RelatedDoc
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedLink)) {
		body["relatedLink"] = request.RelatedLink
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WordAlias)) {
		body["wordAlias"] = request.WordAlias
	}

	if !tea.BoolValue(util.IsUnset(request.WordName)) {
		body["wordName"] = request.WordName
	}

	if !tea.BoolValue(util.IsUnset(request.WordParaphrase)) {
		body["wordParaphrase"] = request.WordParaphrase
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
	_result = &PediaWordsAddResponse{}
	_body, _err := client.DoROARequest(tea.String("PediaWordsAdd"), tea.String("pedia_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/pedia/words"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PediaWordsApprove(request *PediaWordsApproveRequest) (_result *PediaWordsApproveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PediaWordsApproveHeaders{}
	_result = &PediaWordsApproveResponse{}
	_body, _err := client.PediaWordsApproveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PediaWordsApproveWithOptions(request *PediaWordsApproveRequest, headers *PediaWordsApproveHeaders, runtime *util.RuntimeOptions) (_result *PediaWordsApproveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApproveReason)) {
		body["approveReason"] = request.ApproveReason
	}

	if !tea.BoolValue(util.IsUnset(request.ApproveStatus)) {
		body["approveStatus"] = request.ApproveStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ImHighLight)) {
		body["imHighLight"] = request.ImHighLight
	}

	if !tea.BoolValue(util.IsUnset(request.SimHighLight)) {
		body["simHighLight"] = request.SimHighLight
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
	_result = &PediaWordsApproveResponse{}
	_body, _err := client.DoROARequest(tea.String("PediaWordsApprove"), tea.String("pedia_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/pedia/words/approve"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PediaWordsDelete(request *PediaWordsDeleteRequest) (_result *PediaWordsDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PediaWordsDeleteHeaders{}
	_result = &PediaWordsDeleteResponse{}
	_body, _err := client.PediaWordsDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PediaWordsDeleteWithOptions(request *PediaWordsDeleteRequest, headers *PediaWordsDeleteHeaders, runtime *util.RuntimeOptions) (_result *PediaWordsDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["uuid"] = request.Uuid
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
	_result = &PediaWordsDeleteResponse{}
	_body, _err := client.DoROARequest(tea.String("PediaWordsDelete"), tea.String("pedia_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/pedia/words"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PediaWordsQuery(request *PediaWordsQueryRequest) (_result *PediaWordsQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PediaWordsQueryHeaders{}
	_result = &PediaWordsQueryResponse{}
	_body, _err := client.PediaWordsQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PediaWordsQueryWithOptions(request *PediaWordsQueryRequest, headers *PediaWordsQueryHeaders, runtime *util.RuntimeOptions) (_result *PediaWordsQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["uuid"] = request.Uuid
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
	_result = &PediaWordsQueryResponse{}
	_body, _err := client.DoROARequest(tea.String("PediaWordsQuery"), tea.String("pedia_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/pedia/words/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PediaWordsSearch(request *PediaWordsSearchRequest) (_result *PediaWordsSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PediaWordsSearchHeaders{}
	_result = &PediaWordsSearchResponse{}
	_body, _err := client.PediaWordsSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PediaWordsSearchWithOptions(request *PediaWordsSearchRequest, headers *PediaWordsSearchHeaders, runtime *util.RuntimeOptions) (_result *PediaWordsSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WordName)) {
		body["wordName"] = request.WordName
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
	_result = &PediaWordsSearchResponse{}
	_body, _err := client.DoROARequest(tea.String("PediaWordsSearch"), tea.String("pedia_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/pedia/words/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PediaWordsUpdate(request *PediaWordsUpdateRequest) (_result *PediaWordsUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PediaWordsUpdateHeaders{}
	_result = &PediaWordsUpdateResponse{}
	_body, _err := client.PediaWordsUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PediaWordsUpdateWithOptions(request *PediaWordsUpdateRequest, headers *PediaWordsUpdateHeaders, runtime *util.RuntimeOptions) (_result *PediaWordsUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppLink)) {
		body["appLink"] = request.AppLink
	}

	if !tea.BoolValue(util.IsUnset(request.ContactList)) {
		body["contactList"] = request.ContactList
	}

	if !tea.BoolValue(util.IsUnset(request.HighLightWordAlias)) {
		body["highLightWordAlias"] = request.HighLightWordAlias
	}

	if !tea.BoolValue(util.IsUnset(request.PicList)) {
		body["picList"] = request.PicList
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedDoc)) {
		body["relatedDoc"] = request.RelatedDoc
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedLink)) {
		body["relatedLink"] = request.RelatedLink
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.WordAlias)) {
		body["wordAlias"] = request.WordAlias
	}

	if !tea.BoolValue(util.IsUnset(request.WordName)) {
		body["wordName"] = request.WordName
	}

	if !tea.BoolValue(util.IsUnset(request.WordParaphrase)) {
		body["wordParaphrase"] = request.WordParaphrase
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
	_result = &PediaWordsUpdateResponse{}
	_body, _err := client.DoROARequest(tea.String("PediaWordsUpdate"), tea.String("pedia_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/pedia/words"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
