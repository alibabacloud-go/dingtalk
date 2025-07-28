// This file is auto-generated, don't edit it. Thanks.
package pedia_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
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
	ContactList        []*PediaWordsAddRequestContactList `json:"contactList,omitempty" xml:"contactList,omitempty" type:"Repeated"`
	HighLightWordAlias []*string                          `json:"highLightWordAlias,omitempty" xml:"highLightWordAlias,omitempty" type:"Repeated"`
	PicList            []*PediaWordsAddRequestPicList     `json:"picList,omitempty" xml:"picList,omitempty" type:"Repeated"`
	RelatedDoc         []*PediaWordsAddRequestRelatedDoc  `json:"relatedDoc,omitempty" xml:"relatedDoc,omitempty" type:"Repeated"`
	RelatedLink        []*PediaWordsAddRequestRelatedLink `json:"relatedLink,omitempty" xml:"relatedLink,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 23231231123
	UserId    *string   `json:"userId,omitempty" xml:"userId,omitempty"`
	WordAlias []*string `json:"wordAlias,omitempty" xml:"wordAlias,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 词条名称
	WordName *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 释义
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
	// example:
	//
	// @123243
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// example:
	//
	// 名称
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// 1231243213
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
	// example:
	//
	// https://23987874.com
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
	// example:
	//
	// https://123456.com
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 相关文档
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// adoc
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
	// example:
	//
	// http://1233435.com
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 相关链接
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
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 232432
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PediaWordsAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *PediaWordsAddResponse) SetStatusCode(v int32) *PediaWordsAddResponse {
	s.StatusCode = &v
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
	AliDocHighLight *bool `json:"aliDocHighLight,omitempty" xml:"aliDocHighLight,omitempty"`
	// example:
	//
	// 拒绝
	ApproveReason *string `json:"approveReason,omitempty" xml:"approveReason,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ApproveStatus *string `json:"approveStatus,omitempty" xml:"approveStatus,omitempty"`
	// This parameter is required.
	ImHighLight *bool `json:"imHighLight,omitempty" xml:"imHighLight,omitempty"`
	// This parameter is required.
	SimHighLight *bool `json:"simHighLight,omitempty" xml:"simHighLight,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 232432
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1213132
	Uuid *int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s PediaWordsApproveRequest) String() string {
	return tea.Prettify(s)
}

func (s PediaWordsApproveRequest) GoString() string {
	return s.String()
}

func (s *PediaWordsApproveRequest) SetAliDocHighLight(v bool) *PediaWordsApproveRequest {
	s.AliDocHighLight = &v
	return s
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PediaWordsApproveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *PediaWordsApproveResponse) SetStatusCode(v int32) *PediaWordsApproveResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// 2123132
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 212112
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
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 123456789
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PediaWordsDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *PediaWordsDeleteResponse) SetStatusCode(v int32) *PediaWordsDeleteResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// 212121
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 211121
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
	Data    *PediaWordsQueryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Success *bool                            `json:"success,omitempty" xml:"success,omitempty"`
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
	AppLink []*PediaWordsQueryResponseBodyDataAppLink `json:"appLink,omitempty" xml:"appLink,omitempty" type:"Repeated"`
	// example:
	//
	// 审核者
	ApproveName *string                                       `json:"approveName,omitempty" xml:"approveName,omitempty"`
	ContactList []*PediaWordsQueryResponseBodyDataContactList `json:"contactList,omitempty" xml:"contactList,omitempty" type:"Repeated"`
	Contacts    []*string                                     `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// example:
	//
	// 创建者
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// 31312312
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 321312312
	GmtModify          *int64    `json:"gmtModify,omitempty" xml:"gmtModify,omitempty"`
	HighLightWordAlias []*string `json:"highLightWordAlias,omitempty" xml:"highLightWordAlias,omitempty" type:"Repeated"`
	ImHighLight        *bool     `json:"imHighLight,omitempty" xml:"imHighLight,omitempty"`
	// example:
	//
	// 12345678
	ParentUuid   *int64                                        `json:"parentUuid,omitempty" xml:"parentUuid,omitempty"`
	PicList      []*PediaWordsQueryResponseBodyDataPicList     `json:"picList,omitempty" xml:"picList,omitempty" type:"Repeated"`
	RelatedDoc   []*PediaWordsQueryResponseBodyDataRelatedDoc  `json:"relatedDoc,omitempty" xml:"relatedDoc,omitempty" type:"Repeated"`
	RelatedLink  []*PediaWordsQueryResponseBodyDataRelatedLink `json:"relatedLink,omitempty" xml:"relatedLink,omitempty" type:"Repeated"`
	SimHighLight *bool                                         `json:"simHighLight,omitempty" xml:"simHighLight,omitempty"`
	// example:
	//
	// 剔除富文本释义
	SimpleWordParaphrase *string   `json:"simpleWordParaphrase,omitempty" xml:"simpleWordParaphrase,omitempty"`
	TagsList             []*string `json:"tagsList,omitempty" xml:"tagsList,omitempty" type:"Repeated"`
	// example:
	//
	// 更新者
	UpdaterName *string `json:"updaterName,omitempty" xml:"updaterName,omitempty"`
	// example:
	//
	// 213123123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 123123121
	Uuid      *int64    `json:"uuid,omitempty" xml:"uuid,omitempty"`
	WordAlias []*string `json:"wordAlias,omitempty" xml:"wordAlias,omitempty" type:"Repeated"`
	// example:
	//
	// 词条名称
	WordName *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
	// example:
	//
	// 释义
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
	// example:
	//
	// 相关应用
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// htttps://1234567
	IconLink *string `json:"iconLink,omitempty" xml:"iconLink,omitempty"`
	// example:
	//
	// http://123344.com
	PcLink *string `json:"pcLink,omitempty" xml:"pcLink,omitempty"`
	// example:
	//
	// https://12334.com
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
	// example:
	//
	// @dasdasd
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// example:
	//
	// 相关联系人
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// 12321231
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
	// example:
	//
	// http://123455.com
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
	// example:
	//
	// http://123456.com
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 相关文档
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// adoc
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
	// example:
	//
	// http://123343.com
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 相关链接
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PediaWordsQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *PediaWordsQueryResponse) SetStatusCode(v int32) *PediaWordsQueryResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 121213213
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// if can be null:
	// false
	//
	// example:
	//
	// 企业百科
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
	Data    []*PediaWordsSearchResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
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
	AppLink []*PediaWordsSearchResponseBodyDataAppLink `json:"appLink,omitempty" xml:"appLink,omitempty" type:"Repeated"`
	// example:
	//
	// 审核人钉钉名称
	ApproveName *string                                        `json:"approveName,omitempty" xml:"approveName,omitempty"`
	ContactList []*PediaWordsSearchResponseBodyDataContactList `json:"contactList,omitempty" xml:"contactList,omitempty" type:"Repeated"`
	Contacts    []*string                                      `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// example:
	//
	// 创建者
	CreatorName        *string                                        `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	GmtCreate          *int64                                         `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModify          *int64                                         `json:"gmtModify,omitempty" xml:"gmtModify,omitempty"`
	HighLightWordAlias []*string                                      `json:"highLightWordAlias,omitempty" xml:"highLightWordAlias,omitempty" type:"Repeated"`
	ImHighLight        *bool                                          `json:"imHighLight,omitempty" xml:"imHighLight,omitempty"`
	ParentUuid         *int64                                         `json:"parentUuid,omitempty" xml:"parentUuid,omitempty"`
	PicList            []*PediaWordsSearchResponseBodyDataPicList     `json:"picList,omitempty" xml:"picList,omitempty" type:"Repeated"`
	RelatedDoc         []*PediaWordsSearchResponseBodyDataRelatedDoc  `json:"relatedDoc,omitempty" xml:"relatedDoc,omitempty" type:"Repeated"`
	RelatedLink        []*PediaWordsSearchResponseBodyDataRelatedLink `json:"relatedLink,omitempty" xml:"relatedLink,omitempty" type:"Repeated"`
	SimHighLight       *bool                                          `json:"simHighLight,omitempty" xml:"simHighLight,omitempty"`
	// example:
	//
	// 剔除了富文本格式后的释义信息
	SimpleWordParaphrase *string   `json:"simpleWordParaphrase,omitempty" xml:"simpleWordParaphrase,omitempty"`
	TagsList             []*string `json:"tagsList,omitempty" xml:"tagsList,omitempty" type:"Repeated"`
	// example:
	//
	// 修改人钉钉名称
	UpdaterName *string `json:"updaterName,omitempty" xml:"updaterName,omitempty"`
	// example:
	//
	// 312312312
	UserId    *string   `json:"userId,omitempty" xml:"userId,omitempty"`
	Uuid      *int64    `json:"uuid,omitempty" xml:"uuid,omitempty"`
	WordAlias []*string `json:"wordAlias,omitempty" xml:"wordAlias,omitempty" type:"Repeated"`
	// example:
	//
	// 测试词条
	WordName *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
	// example:
	//
	// 释义信息
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
	// example:
	//
	// 应用名
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// https://123434.com
	IconLink *string `json:"iconLink,omitempty" xml:"iconLink,omitempty"`
	// example:
	//
	// http://123.com
	PcLink *string `json:"pcLink,omitempty" xml:"pcLink,omitempty"`
	// example:
	//
	// http://1244.com
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
	// example:
	//
	// @12321312ds
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// example:
	//
	// 员工名称
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// 1232343
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
	// example:
	//
	// https://1234.com
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
	// example:
	//
	// https://12324.com
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 文档名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// adoc
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
	// example:
	//
	// https://123112.com
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 文档名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 空值
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PediaWordsSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *PediaWordsSearchResponse) SetStatusCode(v int32) *PediaWordsSearchResponse {
	s.StatusCode = &v
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
	AppLink            []*PediaWordsUpdateRequestAppLink     `json:"appLink,omitempty" xml:"appLink,omitempty" type:"Repeated"`
	ContactList        []*PediaWordsUpdateRequestContactList `json:"contactList,omitempty" xml:"contactList,omitempty" type:"Repeated"`
	HighLightWordAlias []*string                             `json:"highLightWordAlias,omitempty" xml:"highLightWordAlias,omitempty" type:"Repeated"`
	PicList            []*PediaWordsUpdateRequestPicList     `json:"picList,omitempty" xml:"picList,omitempty" type:"Repeated"`
	RelatedDoc         []*PediaWordsUpdateRequestRelatedDoc  `json:"relatedDoc,omitempty" xml:"relatedDoc,omitempty" type:"Repeated"`
	RelatedLink        []*PediaWordsUpdateRequestRelatedLink `json:"relatedLink,omitempty" xml:"relatedLink,omitempty" type:"Repeated"`
	// example:
	//
	// 312123213
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2131321
	Uuid      *int64    `json:"uuid,omitempty" xml:"uuid,omitempty"`
	WordAlias []*string `json:"wordAlias,omitempty" xml:"wordAlias,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 词条名称
	WordName *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 释义
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
	// example:
	//
	// 应用名称
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// https://123243435.com
	IconLink *string `json:"iconLink,omitempty" xml:"iconLink,omitempty"`
	// example:
	//
	// http://213435.com
	PcLink *string `json:"pcLink,omitempty" xml:"pcLink,omitempty"`
	// example:
	//
	// htttps://12345.com
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
	// example:
	//
	// @12312312
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// example:
	//
	// 名称
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// 12131312
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
	// example:
	//
	// https://1234554.com
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
	// example:
	//
	// https://213567.com
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 相关文档
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// adoc
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
	// example:
	//
	// https://123466.com
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 相关链接
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
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 3213213213
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PediaWordsUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *PediaWordsUpdateResponse) SetStatusCode(v int32) *PediaWordsUpdateResponse {
	s.StatusCode = &v
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
// 企业百科增加当前企业词条信息
//
// @param request - PediaWordsAddRequest
//
// @param headers - PediaWordsAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PediaWordsAddResponse
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
	params := &openapi.Params{
		Action:      tea.String("PediaWordsAdd"),
		Version:     tea.String("pedia_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/pedia/words"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PediaWordsAddResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 企业百科增加当前企业词条信息
//
// @param request - PediaWordsAddRequest
//
// @return PediaWordsAddResponse
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

// Summary:
//
// 企业百科针对待审核词条进行审核
//
// @param request - PediaWordsApproveRequest
//
// @param headers - PediaWordsApproveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PediaWordsApproveResponse
func (client *Client) PediaWordsApproveWithOptions(request *PediaWordsApproveRequest, headers *PediaWordsApproveHeaders, runtime *util.RuntimeOptions) (_result *PediaWordsApproveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliDocHighLight)) {
		body["aliDocHighLight"] = request.AliDocHighLight
	}

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
	params := &openapi.Params{
		Action:      tea.String("PediaWordsApprove"),
		Version:     tea.String("pedia_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/pedia/words/approve"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PediaWordsApproveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 企业百科针对待审核词条进行审核
//
// @param request - PediaWordsApproveRequest
//
// @return PediaWordsApproveResponse
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

// Summary:
//
// 企业百科针对uuid删除当前词条
//
// @param request - PediaWordsDeleteRequest
//
// @param headers - PediaWordsDeleteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PediaWordsDeleteResponse
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
	params := &openapi.Params{
		Action:      tea.String("PediaWordsDelete"),
		Version:     tea.String("pedia_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/pedia/words"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PediaWordsDeleteResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 企业百科针对uuid删除当前词条
//
// @param request - PediaWordsDeleteRequest
//
// @return PediaWordsDeleteResponse
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

// Summary:
//
// 根据词条主键ID查询当前词条详情
//
// @param request - PediaWordsQueryRequest
//
// @param headers - PediaWordsQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PediaWordsQueryResponse
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
	params := &openapi.Params{
		Action:      tea.String("PediaWordsQuery"),
		Version:     tea.String("pedia_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/pedia/words/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PediaWordsQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据词条主键ID查询当前词条详情
//
// @param request - PediaWordsQueryRequest
//
// @return PediaWordsQueryResponse
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

// Summary:
//
// 分页获取企业词条信息
//
// @param request - PediaWordsSearchRequest
//
// @param headers - PediaWordsSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PediaWordsSearchResponse
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
	params := &openapi.Params{
		Action:      tea.String("PediaWordsSearch"),
		Version:     tea.String("pedia_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/pedia/words/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PediaWordsSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取企业词条信息
//
// @param request - PediaWordsSearchRequest
//
// @return PediaWordsSearchResponse
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

// Summary:
//
// 企业百科对当前已经生效词条进行编辑
//
// @param request - PediaWordsUpdateRequest
//
// @param headers - PediaWordsUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PediaWordsUpdateResponse
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
	params := &openapi.Params{
		Action:      tea.String("PediaWordsUpdate"),
		Version:     tea.String("pedia_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/pedia/words"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PediaWordsUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 企业百科对当前已经生效词条进行编辑
//
// @param request - PediaWordsUpdateRequest
//
// @return PediaWordsUpdateResponse
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
