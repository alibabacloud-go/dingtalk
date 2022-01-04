// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package jzcrm_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EditContactHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditContactHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditContactHeaders) GoString() string {
	return s.String()
}

func (s *EditContactHeaders) SetCommonHeaders(v map[string]*string) *EditContactHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditContactHeaders) SetXAcsDingtalkAccessToken(v string) *EditContactHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditContactRequest struct {
	// 编辑数据
	Data *EditContactRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写197
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditContactRequest) String() string {
	return tea.Prettify(s)
}

func (s EditContactRequest) GoString() string {
	return s.String()
}

func (s *EditContactRequest) SetData(v *EditContactRequestData) *EditContactRequest {
	s.Data = v
	return s
}

func (s *EditContactRequest) SetDatatype(v int64) *EditContactRequest {
	s.Datatype = &v
	return s
}

func (s *EditContactRequest) SetMsgid(v int64) *EditContactRequest {
	s.Msgid = &v
	return s
}

func (s *EditContactRequest) SetStamp(v int64) *EditContactRequest {
	s.Stamp = &v
	return s
}

type EditContactRequestData struct {
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 住址
	LxrAddress *string `json:"lxr_address,omitempty" xml:"lxr_address,omitempty"`
	// 生日
	LxrBirthday *string `json:"lxr_birthday,omitempty" xml:"lxr_birthday,omitempty"`
	// 称谓
	LxrChengwei *string `json:"lxr_chengwei,omitempty" xml:"lxr_chengwei,omitempty"`
	// 证件号码
	LxrCtnumber *string `json:"lxr_ctnumber,omitempty" xml:"lxr_ctnumber,omitempty"`
	// 证件类型
	LxrCttype *string `json:"lxr_cttype,omitempty" xml:"lxr_cttype,omitempty"`
	// 对应客户
	LxrCustomerid *string `json:"lxr_customerid,omitempty" xml:"lxr_customerid,omitempty"`
	// 部门
	LxrDepartment *string `json:"lxr_department,omitempty" xml:"lxr_department,omitempty"`
	// 钉钉号
	LxrDingtalk *string `json:"lxr_dingtalk,omitempty" xml:"lxr_dingtalk,omitempty"`
	// Email
	LxrEmail *string `json:"lxr_email,omitempty" xml:"lxr_email,omitempty"`
	// 传真
	LxrFax *string `json:"lxr_fax,omitempty" xml:"lxr_fax,omitempty"`
	// 分类
	LxrGroup *string `json:"lxr_group,omitempty" xml:"lxr_group,omitempty"`
	// 手机
	LxrHandset *string `json:"lxr_handset,omitempty" xml:"lxr_handset,omitempty"`
	// 职务
	LxrHeadship *string `json:"lxr_headship,omitempty" xml:"lxr_headship,omitempty"`
	// 爱好
	LxrLike *string `json:"lxr_like,omitempty" xml:"lxr_like,omitempty"`
	// 姓名
	LxrName *string `json:"lxr_name,omitempty" xml:"lxr_name,omitempty"`
	// 联系名片
	LxrPhoto *string `json:"lxr_photo,omitempty" xml:"lxr_photo,omitempty"`
	// 负责业务
	LxrPreside *string `json:"lxr_preside,omitempty" xml:"lxr_preside,omitempty"`
	// 邮编
	LxrPst *string `json:"lxr_pst,omitempty" xml:"lxr_pst,omitempty"`
	// QQ
	LxrQq *string `json:"lxr_qq,omitempty" xml:"lxr_qq,omitempty"`
	// 备注
	LxrRemark *string `json:"lxr_remark,omitempty" xml:"lxr_remark,omitempty"`
	// 性别（男，女）
	LxrSex *string `json:"lxr_sex,omitempty" xml:"lxr_sex,omitempty"`
	// Skype
	LxrSkype *string `json:"lxr_skype,omitempty" xml:"lxr_skype,omitempty"`
	// 家庭电话
	LxrTel *string `json:"lxr_tel,omitempty" xml:"lxr_tel,omitempty"`
	// 类型（联系人，主联系人）
	LxrType *string `json:"lxr_type,omitempty" xml:"lxr_type,omitempty"`
	// 旺旺
	LxrWangwang *string `json:"lxr_wangwang,omitempty" xml:"lxr_wangwang,omitempty"`
	// 微信号
	LxrWeixin *string `json:"lxr_weixin,omitempty" xml:"lxr_weixin,omitempty"`
	// 工作电话
	LxrWorktel *string `json:"lxr_worktel,omitempty" xml:"lxr_worktel,omitempty"`
}

func (s EditContactRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditContactRequestData) GoString() string {
	return s.String()
}

func (s *EditContactRequestData) SetDataUserid(v string) *EditContactRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditContactRequestData) SetLxrAddress(v string) *EditContactRequestData {
	s.LxrAddress = &v
	return s
}

func (s *EditContactRequestData) SetLxrBirthday(v string) *EditContactRequestData {
	s.LxrBirthday = &v
	return s
}

func (s *EditContactRequestData) SetLxrChengwei(v string) *EditContactRequestData {
	s.LxrChengwei = &v
	return s
}

func (s *EditContactRequestData) SetLxrCtnumber(v string) *EditContactRequestData {
	s.LxrCtnumber = &v
	return s
}

func (s *EditContactRequestData) SetLxrCttype(v string) *EditContactRequestData {
	s.LxrCttype = &v
	return s
}

func (s *EditContactRequestData) SetLxrCustomerid(v string) *EditContactRequestData {
	s.LxrCustomerid = &v
	return s
}

func (s *EditContactRequestData) SetLxrDepartment(v string) *EditContactRequestData {
	s.LxrDepartment = &v
	return s
}

func (s *EditContactRequestData) SetLxrDingtalk(v string) *EditContactRequestData {
	s.LxrDingtalk = &v
	return s
}

func (s *EditContactRequestData) SetLxrEmail(v string) *EditContactRequestData {
	s.LxrEmail = &v
	return s
}

func (s *EditContactRequestData) SetLxrFax(v string) *EditContactRequestData {
	s.LxrFax = &v
	return s
}

func (s *EditContactRequestData) SetLxrGroup(v string) *EditContactRequestData {
	s.LxrGroup = &v
	return s
}

func (s *EditContactRequestData) SetLxrHandset(v string) *EditContactRequestData {
	s.LxrHandset = &v
	return s
}

func (s *EditContactRequestData) SetLxrHeadship(v string) *EditContactRequestData {
	s.LxrHeadship = &v
	return s
}

func (s *EditContactRequestData) SetLxrLike(v string) *EditContactRequestData {
	s.LxrLike = &v
	return s
}

func (s *EditContactRequestData) SetLxrName(v string) *EditContactRequestData {
	s.LxrName = &v
	return s
}

func (s *EditContactRequestData) SetLxrPhoto(v string) *EditContactRequestData {
	s.LxrPhoto = &v
	return s
}

func (s *EditContactRequestData) SetLxrPreside(v string) *EditContactRequestData {
	s.LxrPreside = &v
	return s
}

func (s *EditContactRequestData) SetLxrPst(v string) *EditContactRequestData {
	s.LxrPst = &v
	return s
}

func (s *EditContactRequestData) SetLxrQq(v string) *EditContactRequestData {
	s.LxrQq = &v
	return s
}

func (s *EditContactRequestData) SetLxrRemark(v string) *EditContactRequestData {
	s.LxrRemark = &v
	return s
}

func (s *EditContactRequestData) SetLxrSex(v string) *EditContactRequestData {
	s.LxrSex = &v
	return s
}

func (s *EditContactRequestData) SetLxrSkype(v string) *EditContactRequestData {
	s.LxrSkype = &v
	return s
}

func (s *EditContactRequestData) SetLxrTel(v string) *EditContactRequestData {
	s.LxrTel = &v
	return s
}

func (s *EditContactRequestData) SetLxrType(v string) *EditContactRequestData {
	s.LxrType = &v
	return s
}

func (s *EditContactRequestData) SetLxrWangwang(v string) *EditContactRequestData {
	s.LxrWangwang = &v
	return s
}

func (s *EditContactRequestData) SetLxrWeixin(v string) *EditContactRequestData {
	s.LxrWeixin = &v
	return s
}

func (s *EditContactRequestData) SetLxrWorktel(v string) *EditContactRequestData {
	s.LxrWorktel = &v
	return s
}

type EditContactResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditContactResponseBody) GoString() string {
	return s.String()
}

func (s *EditContactResponseBody) SetMsgid(v int64) *EditContactResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditContactResponseBody) SetTime(v string) *EditContactResponseBody {
	s.Time = &v
	return s
}

type EditContactResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditContactResponse) String() string {
	return tea.Prettify(s)
}

func (s EditContactResponse) GoString() string {
	return s.String()
}

func (s *EditContactResponse) SetHeaders(v map[string]*string) *EditContactResponse {
	s.Headers = v
	return s
}

func (s *EditContactResponse) SetBody(v *EditContactResponseBody) *EditContactResponse {
	s.Body = v
	return s
}

type EditCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditCustomerHeaders) GoString() string {
	return s.String()
}

func (s *EditCustomerHeaders) SetCommonHeaders(v map[string]*string) *EditCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *EditCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditCustomerRequest struct {
	// 编辑数据
	Data *EditCustomerRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写148
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s EditCustomerRequest) GoString() string {
	return s.String()
}

func (s *EditCustomerRequest) SetData(v *EditCustomerRequestData) *EditCustomerRequest {
	s.Data = v
	return s
}

func (s *EditCustomerRequest) SetDatatype(v int64) *EditCustomerRequest {
	s.Datatype = &v
	return s
}

func (s *EditCustomerRequest) SetMsgid(v int64) *EditCustomerRequest {
	s.Msgid = &v
	return s
}

func (s *EditCustomerRequest) SetStamp(v int64) *EditCustomerRequest {
	s.Stamp = &v
	return s
}

type EditCustomerRequestData struct {
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 家庭地址
	KhAddress *string `json:"kh_address,omitempty" xml:"kh_address,omitempty"`
	// 称谓
	KhAppellation *string `json:"kh_appellation,omitempty" xml:"kh_appellation,omitempty"`
	// 爱好
	KhBefontof *string `json:"kh_befontof,omitempty" xml:"kh_befontof,omitempty"`
	// 开票资料
	KhBillinfo *string `json:"kh_billinfo,omitempty" xml:"kh_billinfo,omitempty"`
	// 城市
	KhCity *string `json:"kh_city,omitempty" xml:"kh_city,omitempty"`
	// 类别（企业客户，个人客户，供应商，个人供应商）
	KhClass *string `json:"kh_class,omitempty" xml:"kh_class,omitempty"`
	// 单位地址
	KhCoaddress *string `json:"kh_coaddress,omitempty" xml:"kh_coaddress,omitempty"`
	// 联系人分类
	KhContype *string `json:"kh_contype,omitempty" xml:"kh_contype,omitempty"`
	// 国家地区
	KhCountry *string `json:"kh_country,omitempty" xml:"kh_country,omitempty"`
	// 信用等级（低，中，高）
	KhCreditgrade *string `json:"kh_creditgrade,omitempty" xml:"kh_creditgrade,omitempty"`
	// 证件号码
	KhCtnumber *string `json:"kh_ctnumber,omitempty" xml:"kh_ctnumber,omitempty"`
	// 证件类型
	KhCttype *string `json:"kh_cttype,omitempty" xml:"kh_cttype,omitempty"`
	// 部门
	KhDepartment *string `json:"kh_department,omitempty" xml:"kh_department,omitempty"`
	// 钉钉号
	KhDingtalk *string `json:"kh_dingtalk,omitempty" xml:"kh_dingtalk,omitempty"`
	// 邮箱
	KhEmail *string `json:"kh_email,omitempty" xml:"kh_email,omitempty"`
	// 人员规模
	KhEmployees *string `json:"kh_employees,omitempty" xml:"kh_employees,omitempty"`
	// 传真
	KhFax *string `json:"kh_fax,omitempty" xml:"kh_fax,omitempty"`
	// 来源
	KhFrom *string `json:"kh_from,omitempty" xml:"kh_from,omitempty"`
	// 手机
	KhHandset *string `json:"kh_handset,omitempty" xml:"kh_handset,omitempty"`
	// 职务
	KhHeadship *string `json:"kh_headship,omitempty" xml:"kh_headship,omitempty"`
	// 热点分类
	KhHotfl *string `json:"kh_hotfl,omitempty" xml:"kh_hotfl,omitempty"`
	// 热度（无，低热，中热，高热）
	KhHotlevel *string `json:"kh_hotlevel,omitempty" xml:"kh_hotlevel,omitempty"`
	// 热点说明
	KhHotmemo *string `json:"kh_hotmemo,omitempty" xml:"kh_hotmemo,omitempty"`
	// 热点客户（是，否）
	KhHottype *string `json:"kh_hottype,omitempty" xml:"kh_hottype,omitempty"`
	// 行业
	KhIndustry *string `json:"kh_industry,omitempty" xml:"kh_industry,omitempty"`
	// 公司简介
	KhInfo *string `json:"kh_info,omitempty" xml:"kh_info,omitempty"`
	// 客户级别
	KhJibie *string `json:"kh_jibie,omitempty" xml:"kh_jibie,omitempty"`
	// 客户名称
	KhName *string `json:"kh_name,omitempty" xml:"kh_name,omitempty"`
	// 上级客户
	KhPkhid *string `json:"kh_pkhid,omitempty" xml:"kh_pkhid,omitempty"`
	// 负责业务
	KhPreside *string `json:"kh_preside,omitempty" xml:"kh_preside,omitempty"`
	// 省份
	KhProvince *string `json:"kh_province,omitempty" xml:"kh_province,omitempty"`
	// 邮编
	KhPst *string `json:"kh_pst,omitempty" xml:"kh_pst,omitempty"`
	// QQ
	KhQq *string `json:"kh_qq,omitempty" xml:"kh_qq,omitempty"`
	// 关系等级
	KhRalagrade *string `json:"kh_ralagrade,omitempty" xml:"kh_ralagrade,omitempty"`
	// 备注
	KhRemark *string `json:"kh_remark,omitempty" xml:"kh_remark,omitempty"`
	// 性别（男，女）
	KhSex *string `json:"kh_sex,omitempty" xml:"kh_sex,omitempty"`
	// 助记简称
	KhShortname *string `json:"kh_shortname,omitempty" xml:"kh_shortname,omitempty"`
	// Skype
	KhSkype *string `json:"kh_skype,omitempty" xml:"kh_skype,omitempty"`
	// 编号
	KhSn *string `json:"kh_sn,omitempty" xml:"kh_sn,omitempty"`
	// 阶段
	KhStatus *string `json:"kh_status,omitempty" xml:"kh_status,omitempty"`
	// 家庭电话
	KhTel *string `json:"kh_tel,omitempty" xml:"kh_tel,omitempty"`
	// 种类
	KhType *string `json:"kh_type,omitempty" xml:"kh_type,omitempty"`
	// 价值评估（低，中，高）
	KhValrating *string `json:"kh_valrating,omitempty" xml:"kh_valrating,omitempty"`
	// 旺旺
	KhWangwang *string `json:"kh_wangwang,omitempty" xml:"kh_wangwang,omitempty"`
	// 网址
	KhWeb *string `json:"kh_web,omitempty" xml:"kh_web,omitempty"`
	// 微信号
	KhWeixin *string `json:"kh_weixin,omitempty" xml:"kh_weixin,omitempty"`
	// 工作电话
	KhWorktel *string `json:"kh_worktel,omitempty" xml:"kh_worktel,omitempty"`
}

func (s EditCustomerRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditCustomerRequestData) GoString() string {
	return s.String()
}

func (s *EditCustomerRequestData) SetDataUserid(v string) *EditCustomerRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditCustomerRequestData) SetKhAddress(v string) *EditCustomerRequestData {
	s.KhAddress = &v
	return s
}

func (s *EditCustomerRequestData) SetKhAppellation(v string) *EditCustomerRequestData {
	s.KhAppellation = &v
	return s
}

func (s *EditCustomerRequestData) SetKhBefontof(v string) *EditCustomerRequestData {
	s.KhBefontof = &v
	return s
}

func (s *EditCustomerRequestData) SetKhBillinfo(v string) *EditCustomerRequestData {
	s.KhBillinfo = &v
	return s
}

func (s *EditCustomerRequestData) SetKhCity(v string) *EditCustomerRequestData {
	s.KhCity = &v
	return s
}

func (s *EditCustomerRequestData) SetKhClass(v string) *EditCustomerRequestData {
	s.KhClass = &v
	return s
}

func (s *EditCustomerRequestData) SetKhCoaddress(v string) *EditCustomerRequestData {
	s.KhCoaddress = &v
	return s
}

func (s *EditCustomerRequestData) SetKhContype(v string) *EditCustomerRequestData {
	s.KhContype = &v
	return s
}

func (s *EditCustomerRequestData) SetKhCountry(v string) *EditCustomerRequestData {
	s.KhCountry = &v
	return s
}

func (s *EditCustomerRequestData) SetKhCreditgrade(v string) *EditCustomerRequestData {
	s.KhCreditgrade = &v
	return s
}

func (s *EditCustomerRequestData) SetKhCtnumber(v string) *EditCustomerRequestData {
	s.KhCtnumber = &v
	return s
}

func (s *EditCustomerRequestData) SetKhCttype(v string) *EditCustomerRequestData {
	s.KhCttype = &v
	return s
}

func (s *EditCustomerRequestData) SetKhDepartment(v string) *EditCustomerRequestData {
	s.KhDepartment = &v
	return s
}

func (s *EditCustomerRequestData) SetKhDingtalk(v string) *EditCustomerRequestData {
	s.KhDingtalk = &v
	return s
}

func (s *EditCustomerRequestData) SetKhEmail(v string) *EditCustomerRequestData {
	s.KhEmail = &v
	return s
}

func (s *EditCustomerRequestData) SetKhEmployees(v string) *EditCustomerRequestData {
	s.KhEmployees = &v
	return s
}

func (s *EditCustomerRequestData) SetKhFax(v string) *EditCustomerRequestData {
	s.KhFax = &v
	return s
}

func (s *EditCustomerRequestData) SetKhFrom(v string) *EditCustomerRequestData {
	s.KhFrom = &v
	return s
}

func (s *EditCustomerRequestData) SetKhHandset(v string) *EditCustomerRequestData {
	s.KhHandset = &v
	return s
}

func (s *EditCustomerRequestData) SetKhHeadship(v string) *EditCustomerRequestData {
	s.KhHeadship = &v
	return s
}

func (s *EditCustomerRequestData) SetKhHotfl(v string) *EditCustomerRequestData {
	s.KhHotfl = &v
	return s
}

func (s *EditCustomerRequestData) SetKhHotlevel(v string) *EditCustomerRequestData {
	s.KhHotlevel = &v
	return s
}

func (s *EditCustomerRequestData) SetKhHotmemo(v string) *EditCustomerRequestData {
	s.KhHotmemo = &v
	return s
}

func (s *EditCustomerRequestData) SetKhHottype(v string) *EditCustomerRequestData {
	s.KhHottype = &v
	return s
}

func (s *EditCustomerRequestData) SetKhIndustry(v string) *EditCustomerRequestData {
	s.KhIndustry = &v
	return s
}

func (s *EditCustomerRequestData) SetKhInfo(v string) *EditCustomerRequestData {
	s.KhInfo = &v
	return s
}

func (s *EditCustomerRequestData) SetKhJibie(v string) *EditCustomerRequestData {
	s.KhJibie = &v
	return s
}

func (s *EditCustomerRequestData) SetKhName(v string) *EditCustomerRequestData {
	s.KhName = &v
	return s
}

func (s *EditCustomerRequestData) SetKhPkhid(v string) *EditCustomerRequestData {
	s.KhPkhid = &v
	return s
}

func (s *EditCustomerRequestData) SetKhPreside(v string) *EditCustomerRequestData {
	s.KhPreside = &v
	return s
}

func (s *EditCustomerRequestData) SetKhProvince(v string) *EditCustomerRequestData {
	s.KhProvince = &v
	return s
}

func (s *EditCustomerRequestData) SetKhPst(v string) *EditCustomerRequestData {
	s.KhPst = &v
	return s
}

func (s *EditCustomerRequestData) SetKhQq(v string) *EditCustomerRequestData {
	s.KhQq = &v
	return s
}

func (s *EditCustomerRequestData) SetKhRalagrade(v string) *EditCustomerRequestData {
	s.KhRalagrade = &v
	return s
}

func (s *EditCustomerRequestData) SetKhRemark(v string) *EditCustomerRequestData {
	s.KhRemark = &v
	return s
}

func (s *EditCustomerRequestData) SetKhSex(v string) *EditCustomerRequestData {
	s.KhSex = &v
	return s
}

func (s *EditCustomerRequestData) SetKhShortname(v string) *EditCustomerRequestData {
	s.KhShortname = &v
	return s
}

func (s *EditCustomerRequestData) SetKhSkype(v string) *EditCustomerRequestData {
	s.KhSkype = &v
	return s
}

func (s *EditCustomerRequestData) SetKhSn(v string) *EditCustomerRequestData {
	s.KhSn = &v
	return s
}

func (s *EditCustomerRequestData) SetKhStatus(v string) *EditCustomerRequestData {
	s.KhStatus = &v
	return s
}

func (s *EditCustomerRequestData) SetKhTel(v string) *EditCustomerRequestData {
	s.KhTel = &v
	return s
}

func (s *EditCustomerRequestData) SetKhType(v string) *EditCustomerRequestData {
	s.KhType = &v
	return s
}

func (s *EditCustomerRequestData) SetKhValrating(v string) *EditCustomerRequestData {
	s.KhValrating = &v
	return s
}

func (s *EditCustomerRequestData) SetKhWangwang(v string) *EditCustomerRequestData {
	s.KhWangwang = &v
	return s
}

func (s *EditCustomerRequestData) SetKhWeb(v string) *EditCustomerRequestData {
	s.KhWeb = &v
	return s
}

func (s *EditCustomerRequestData) SetKhWeixin(v string) *EditCustomerRequestData {
	s.KhWeixin = &v
	return s
}

func (s *EditCustomerRequestData) SetKhWorktel(v string) *EditCustomerRequestData {
	s.KhWorktel = &v
	return s
}

type EditCustomerResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *EditCustomerResponseBody) SetMsgid(v int64) *EditCustomerResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditCustomerResponseBody) SetTime(v string) *EditCustomerResponseBody {
	s.Time = &v
	return s
}

type EditCustomerResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s EditCustomerResponse) GoString() string {
	return s.String()
}

func (s *EditCustomerResponse) SetHeaders(v map[string]*string) *EditCustomerResponse {
	s.Headers = v
	return s
}

func (s *EditCustomerResponse) SetBody(v *EditCustomerResponseBody) *EditCustomerResponse {
	s.Body = v
	return s
}

type EditCustomerPoolHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditCustomerPoolHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditCustomerPoolHeaders) GoString() string {
	return s.String()
}

func (s *EditCustomerPoolHeaders) SetCommonHeaders(v map[string]*string) *EditCustomerPoolHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditCustomerPoolHeaders) SetXAcsDingtalkAccessToken(v string) *EditCustomerPoolHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditCustomerPoolRequest struct {
	// 编辑数据
	Data *EditCustomerPoolRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写238
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditCustomerPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s EditCustomerPoolRequest) GoString() string {
	return s.String()
}

func (s *EditCustomerPoolRequest) SetData(v *EditCustomerPoolRequestData) *EditCustomerPoolRequest {
	s.Data = v
	return s
}

func (s *EditCustomerPoolRequest) SetDatatype(v int64) *EditCustomerPoolRequest {
	s.Datatype = &v
	return s
}

func (s *EditCustomerPoolRequest) SetMsgid(v int64) *EditCustomerPoolRequest {
	s.Msgid = &v
	return s
}

func (s *EditCustomerPoolRequest) SetStamp(v int64) *EditCustomerPoolRequest {
	s.Stamp = &v
	return s
}

type EditCustomerPoolRequestData struct {
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 家庭地址
	KhAddress *string `json:"kh_address,omitempty" xml:"kh_address,omitempty"`
	// 称谓
	KhAppellation *string `json:"kh_appellation,omitempty" xml:"kh_appellation,omitempty"`
	// 爱好
	KhBefontof *string `json:"kh_befontof,omitempty" xml:"kh_befontof,omitempty"`
	// 开票资料
	KhBillinfo *string `json:"kh_billinfo,omitempty" xml:"kh_billinfo,omitempty"`
	// 城市
	KhCity *string `json:"kh_city,omitempty" xml:"kh_city,omitempty"`
	// 类别（企业客户，个人客户，供应商，个人供应商）
	KhClass *string `json:"kh_class,omitempty" xml:"kh_class,omitempty"`
	// 单位地址
	KhCoaddress *string `json:"kh_coaddress,omitempty" xml:"kh_coaddress,omitempty"`
	// 联系人分类
	KhContype *string `json:"kh_contype,omitempty" xml:"kh_contype,omitempty"`
	// 国家地区
	KhCountry *string `json:"kh_country,omitempty" xml:"kh_country,omitempty"`
	// 信用等级（低，中，高）
	KhCreditgrade *string `json:"kh_creditgrade,omitempty" xml:"kh_creditgrade,omitempty"`
	// 证件号码
	KhCtnumber *string `json:"kh_ctnumber,omitempty" xml:"kh_ctnumber,omitempty"`
	// 证件类型
	KhCttype *string `json:"kh_cttype,omitempty" xml:"kh_cttype,omitempty"`
	// 部门
	KhDepartment *string `json:"kh_department,omitempty" xml:"kh_department,omitempty"`
	// 钉钉号
	KhDingtalk *string `json:"kh_dingtalk,omitempty" xml:"kh_dingtalk,omitempty"`
	// 邮箱
	KhEmail *string `json:"kh_email,omitempty" xml:"kh_email,omitempty"`
	// 人员规模
	KhEmployees *string `json:"kh_employees,omitempty" xml:"kh_employees,omitempty"`
	// 传真
	KhFax *string `json:"kh_fax,omitempty" xml:"kh_fax,omitempty"`
	// 来源
	KhFrom *string `json:"kh_from,omitempty" xml:"kh_from,omitempty"`
	// 最后跟踪
	KhGenzongtime *string `json:"kh_genzongtime,omitempty" xml:"kh_genzongtime,omitempty"`
	// 手机
	KhHandset *string `json:"kh_handset,omitempty" xml:"kh_handset,omitempty"`
	// 职务
	KhHeadship *string `json:"kh_headship,omitempty" xml:"kh_headship,omitempty"`
	// 热点分类
	KhHotfl *string `json:"kh_hotfl,omitempty" xml:"kh_hotfl,omitempty"`
	// 热度（无，低热，中热，高热）
	KhHotlevel *string `json:"kh_hotlevel,omitempty" xml:"kh_hotlevel,omitempty"`
	// 热点说明
	KhHotmemo *string `json:"kh_hotmemo,omitempty" xml:"kh_hotmemo,omitempty"`
	// 热点客户（是，否）
	KhHottype *string `json:"kh_hottype,omitempty" xml:"kh_hottype,omitempty"`
	// 行业
	KhIndustry *string `json:"kh_industry,omitempty" xml:"kh_industry,omitempty"`
	// 公司简介
	KhInfo *string `json:"kh_info,omitempty" xml:"kh_info,omitempty"`
	// 客户级别
	KhJibie *string `json:"kh_jibie,omitempty" xml:"kh_jibie,omitempty"`
	// 客户名称
	KhName *string `json:"kh_name,omitempty" xml:"kh_name,omitempty"`
	// 上级客户
	KhPkhid *string `json:"kh_pkhid,omitempty" xml:"kh_pkhid,omitempty"`
	// 负责业务
	KhPreside *string `json:"kh_preside,omitempty" xml:"kh_preside,omitempty"`
	// 省份
	KhProvince *string `json:"kh_province,omitempty" xml:"kh_province,omitempty"`
	// 邮编
	KhPst *string `json:"kh_pst,omitempty" xml:"kh_pst,omitempty"`
	// QQ
	KhQq *string `json:"kh_qq,omitempty" xml:"kh_qq,omitempty"`
	// 关系等级
	KhRalagrade *string `json:"kh_ralagrade,omitempty" xml:"kh_ralagrade,omitempty"`
	// 备注
	KhRemark *string `json:"kh_remark,omitempty" xml:"kh_remark,omitempty"`
	// 性别（男，女）
	KhSex *string `json:"kh_sex,omitempty" xml:"kh_sex,omitempty"`
	// 助记简称
	KhShortname *string `json:"kh_shortname,omitempty" xml:"kh_shortname,omitempty"`
	// Skype
	KhSkype *string `json:"kh_skype,omitempty" xml:"kh_skype,omitempty"`
	// 编号
	KhSn *string `json:"kh_sn,omitempty" xml:"kh_sn,omitempty"`
	// 阶段
	KhStatus *string `json:"kh_status,omitempty" xml:"kh_status,omitempty"`
	// 家庭电话
	KhTel *string `json:"kh_tel,omitempty" xml:"kh_tel,omitempty"`
	// 种类
	KhType *string `json:"kh_type,omitempty" xml:"kh_type,omitempty"`
	// 价值评估（低，中，高）
	KhValrating *string `json:"kh_valrating,omitempty" xml:"kh_valrating,omitempty"`
	// 旺旺
	KhWangwang *string `json:"kh_wangwang,omitempty" xml:"kh_wangwang,omitempty"`
	// 网址
	KhWeb *string `json:"kh_web,omitempty" xml:"kh_web,omitempty"`
	// 微信号
	KhWeixin *string `json:"kh_weixin,omitempty" xml:"kh_weixin,omitempty"`
	// 工作电话
	KhWorktel *string `json:"kh_worktel,omitempty" xml:"kh_worktel,omitempty"`
}

func (s EditCustomerPoolRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditCustomerPoolRequestData) GoString() string {
	return s.String()
}

func (s *EditCustomerPoolRequestData) SetDataUserid(v string) *EditCustomerPoolRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhAddress(v string) *EditCustomerPoolRequestData {
	s.KhAddress = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhAppellation(v string) *EditCustomerPoolRequestData {
	s.KhAppellation = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhBefontof(v string) *EditCustomerPoolRequestData {
	s.KhBefontof = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhBillinfo(v string) *EditCustomerPoolRequestData {
	s.KhBillinfo = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhCity(v string) *EditCustomerPoolRequestData {
	s.KhCity = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhClass(v string) *EditCustomerPoolRequestData {
	s.KhClass = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhCoaddress(v string) *EditCustomerPoolRequestData {
	s.KhCoaddress = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhContype(v string) *EditCustomerPoolRequestData {
	s.KhContype = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhCountry(v string) *EditCustomerPoolRequestData {
	s.KhCountry = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhCreditgrade(v string) *EditCustomerPoolRequestData {
	s.KhCreditgrade = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhCtnumber(v string) *EditCustomerPoolRequestData {
	s.KhCtnumber = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhCttype(v string) *EditCustomerPoolRequestData {
	s.KhCttype = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhDepartment(v string) *EditCustomerPoolRequestData {
	s.KhDepartment = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhDingtalk(v string) *EditCustomerPoolRequestData {
	s.KhDingtalk = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhEmail(v string) *EditCustomerPoolRequestData {
	s.KhEmail = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhEmployees(v string) *EditCustomerPoolRequestData {
	s.KhEmployees = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhFax(v string) *EditCustomerPoolRequestData {
	s.KhFax = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhFrom(v string) *EditCustomerPoolRequestData {
	s.KhFrom = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhGenzongtime(v string) *EditCustomerPoolRequestData {
	s.KhGenzongtime = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhHandset(v string) *EditCustomerPoolRequestData {
	s.KhHandset = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhHeadship(v string) *EditCustomerPoolRequestData {
	s.KhHeadship = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhHotfl(v string) *EditCustomerPoolRequestData {
	s.KhHotfl = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhHotlevel(v string) *EditCustomerPoolRequestData {
	s.KhHotlevel = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhHotmemo(v string) *EditCustomerPoolRequestData {
	s.KhHotmemo = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhHottype(v string) *EditCustomerPoolRequestData {
	s.KhHottype = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhIndustry(v string) *EditCustomerPoolRequestData {
	s.KhIndustry = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhInfo(v string) *EditCustomerPoolRequestData {
	s.KhInfo = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhJibie(v string) *EditCustomerPoolRequestData {
	s.KhJibie = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhName(v string) *EditCustomerPoolRequestData {
	s.KhName = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhPkhid(v string) *EditCustomerPoolRequestData {
	s.KhPkhid = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhPreside(v string) *EditCustomerPoolRequestData {
	s.KhPreside = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhProvince(v string) *EditCustomerPoolRequestData {
	s.KhProvince = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhPst(v string) *EditCustomerPoolRequestData {
	s.KhPst = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhQq(v string) *EditCustomerPoolRequestData {
	s.KhQq = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhRalagrade(v string) *EditCustomerPoolRequestData {
	s.KhRalagrade = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhRemark(v string) *EditCustomerPoolRequestData {
	s.KhRemark = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhSex(v string) *EditCustomerPoolRequestData {
	s.KhSex = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhShortname(v string) *EditCustomerPoolRequestData {
	s.KhShortname = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhSkype(v string) *EditCustomerPoolRequestData {
	s.KhSkype = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhSn(v string) *EditCustomerPoolRequestData {
	s.KhSn = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhStatus(v string) *EditCustomerPoolRequestData {
	s.KhStatus = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhTel(v string) *EditCustomerPoolRequestData {
	s.KhTel = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhType(v string) *EditCustomerPoolRequestData {
	s.KhType = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhValrating(v string) *EditCustomerPoolRequestData {
	s.KhValrating = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhWangwang(v string) *EditCustomerPoolRequestData {
	s.KhWangwang = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhWeb(v string) *EditCustomerPoolRequestData {
	s.KhWeb = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhWeixin(v string) *EditCustomerPoolRequestData {
	s.KhWeixin = &v
	return s
}

func (s *EditCustomerPoolRequestData) SetKhWorktel(v string) *EditCustomerPoolRequestData {
	s.KhWorktel = &v
	return s
}

type EditCustomerPoolResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditCustomerPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditCustomerPoolResponseBody) GoString() string {
	return s.String()
}

func (s *EditCustomerPoolResponseBody) SetMsgid(v int64) *EditCustomerPoolResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditCustomerPoolResponseBody) SetTime(v string) *EditCustomerPoolResponseBody {
	s.Time = &v
	return s
}

type EditCustomerPoolResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditCustomerPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditCustomerPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s EditCustomerPoolResponse) GoString() string {
	return s.String()
}

func (s *EditCustomerPoolResponse) SetHeaders(v map[string]*string) *EditCustomerPoolResponse {
	s.Headers = v
	return s
}

func (s *EditCustomerPoolResponse) SetBody(v *EditCustomerPoolResponseBody) *EditCustomerPoolResponse {
	s.Body = v
	return s
}

type EditExchangeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditExchangeHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditExchangeHeaders) GoString() string {
	return s.String()
}

func (s *EditExchangeHeaders) SetCommonHeaders(v map[string]*string) *EditExchangeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditExchangeHeaders) SetXAcsDingtalkAccessToken(v string) *EditExchangeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditExchangeRequest struct {
	// 编辑数据
	Data *EditExchangeRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写228
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditExchangeRequest) String() string {
	return tea.Prettify(s)
}

func (s EditExchangeRequest) GoString() string {
	return s.String()
}

func (s *EditExchangeRequest) SetData(v *EditExchangeRequestData) *EditExchangeRequest {
	s.Data = v
	return s
}

func (s *EditExchangeRequest) SetDatatype(v int64) *EditExchangeRequest {
	s.Datatype = &v
	return s
}

func (s *EditExchangeRequest) SetMsgid(v int64) *EditExchangeRequest {
	s.Msgid = &v
	return s
}

func (s *EditExchangeRequest) SetStamp(v int64) *EditExchangeRequest {
	s.Stamp = &v
	return s
}

type EditExchangeRequestData struct {
	// 产品明细，json格式
	ChildMx *string `json:"child_mx,omitempty" xml:"child_mx,omitempty"`
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 对应客户
	HhCustomerid *string `json:"hh_customerid,omitempty" xml:"hh_customerid,omitempty"`
	// 换货日期
	HhDate *string `json:"hh_date,omitempty" xml:"hh_date,omitempty"`
	// 换入操作员
	HhInempid *string `json:"hh_inempid,omitempty" xml:"hh_inempid,omitempty"`
	// 换入仓库
	HhInlibid *string `json:"hh_inlibid,omitempty" xml:"hh_inlibid,omitempty"`
	// 换入时间
	HhIntime *string `json:"hh_intime,omitempty" xml:"hh_intime,omitempty"`
	// 换货单号
	HhNumber *string `json:"hh_number,omitempty" xml:"hh_number,omitempty"`
	// 合同/订单
	HhOrderid *string `json:"hh_orderid,omitempty" xml:"hh_orderid,omitempty"`
	// 换出操作员
	HhOutempid *string `json:"hh_outempid,omitempty" xml:"hh_outempid,omitempty"`
	// 换出仓库
	HhOutlibid *string `json:"hh_outlibid,omitempty" xml:"hh_outlibid,omitempty"`
	// 换出时间
	HhOuttime *string `json:"hh_outtime,omitempty" xml:"hh_outtime,omitempty"`
	// 备注
	HhRemark *string `json:"hh_remark,omitempty" xml:"hh_remark,omitempty"`
	// 状态（未执行，已入待出，已出待入，结束）
	HhState *string `json:"hh_state,omitempty" xml:"hh_state,omitempty"`
	// 主题
	HhTitle *string `json:"hh_title,omitempty" xml:"hh_title,omitempty"`
	// 分类
	HhType *string `json:"hh_type,omitempty" xml:"hh_type,omitempty"`
}

func (s EditExchangeRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditExchangeRequestData) GoString() string {
	return s.String()
}

func (s *EditExchangeRequestData) SetChildMx(v string) *EditExchangeRequestData {
	s.ChildMx = &v
	return s
}

func (s *EditExchangeRequestData) SetDataUserid(v string) *EditExchangeRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditExchangeRequestData) SetHhCustomerid(v string) *EditExchangeRequestData {
	s.HhCustomerid = &v
	return s
}

func (s *EditExchangeRequestData) SetHhDate(v string) *EditExchangeRequestData {
	s.HhDate = &v
	return s
}

func (s *EditExchangeRequestData) SetHhInempid(v string) *EditExchangeRequestData {
	s.HhInempid = &v
	return s
}

func (s *EditExchangeRequestData) SetHhInlibid(v string) *EditExchangeRequestData {
	s.HhInlibid = &v
	return s
}

func (s *EditExchangeRequestData) SetHhIntime(v string) *EditExchangeRequestData {
	s.HhIntime = &v
	return s
}

func (s *EditExchangeRequestData) SetHhNumber(v string) *EditExchangeRequestData {
	s.HhNumber = &v
	return s
}

func (s *EditExchangeRequestData) SetHhOrderid(v string) *EditExchangeRequestData {
	s.HhOrderid = &v
	return s
}

func (s *EditExchangeRequestData) SetHhOutempid(v string) *EditExchangeRequestData {
	s.HhOutempid = &v
	return s
}

func (s *EditExchangeRequestData) SetHhOutlibid(v string) *EditExchangeRequestData {
	s.HhOutlibid = &v
	return s
}

func (s *EditExchangeRequestData) SetHhOuttime(v string) *EditExchangeRequestData {
	s.HhOuttime = &v
	return s
}

func (s *EditExchangeRequestData) SetHhRemark(v string) *EditExchangeRequestData {
	s.HhRemark = &v
	return s
}

func (s *EditExchangeRequestData) SetHhState(v string) *EditExchangeRequestData {
	s.HhState = &v
	return s
}

func (s *EditExchangeRequestData) SetHhTitle(v string) *EditExchangeRequestData {
	s.HhTitle = &v
	return s
}

func (s *EditExchangeRequestData) SetHhType(v string) *EditExchangeRequestData {
	s.HhType = &v
	return s
}

type EditExchangeResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditExchangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditExchangeResponseBody) GoString() string {
	return s.String()
}

func (s *EditExchangeResponseBody) SetMsgid(v int64) *EditExchangeResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditExchangeResponseBody) SetTime(v string) *EditExchangeResponseBody {
	s.Time = &v
	return s
}

type EditExchangeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditExchangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditExchangeResponse) String() string {
	return tea.Prettify(s)
}

func (s EditExchangeResponse) GoString() string {
	return s.String()
}

func (s *EditExchangeResponse) SetHeaders(v map[string]*string) *EditExchangeResponse {
	s.Headers = v
	return s
}

func (s *EditExchangeResponse) SetBody(v *EditExchangeResponseBody) *EditExchangeResponse {
	s.Body = v
	return s
}

type EditGoodsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditGoodsHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditGoodsHeaders) GoString() string {
	return s.String()
}

func (s *EditGoodsHeaders) SetCommonHeaders(v map[string]*string) *EditGoodsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditGoodsHeaders) SetXAcsDingtalkAccessToken(v string) *EditGoodsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditGoodsRequest struct {
	// 编辑数据
	Data *EditGoodsRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写154
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditGoodsRequest) String() string {
	return tea.Prettify(s)
}

func (s EditGoodsRequest) GoString() string {
	return s.String()
}

func (s *EditGoodsRequest) SetData(v *EditGoodsRequestData) *EditGoodsRequest {
	s.Data = v
	return s
}

func (s *EditGoodsRequest) SetDatatype(v int64) *EditGoodsRequest {
	s.Datatype = &v
	return s
}

func (s *EditGoodsRequest) SetMsgid(v int64) *EditGoodsRequest {
	s.Msgid = &v
	return s
}

func (s *EditGoodsRequest) SetStamp(v int64) *EditGoodsRequest {
	s.Stamp = &v
	return s
}

type EditGoodsRequestData struct {
	// 上架时间
	Addedtime *string `json:"addedtime,omitempty" xml:"addedtime,omitempty"`
	// 成本价格
	Cbprice *string `json:"cbprice,omitempty" xml:"cbprice,omitempty"`
	// 基准产品
	CpParentid *string `json:"cp_parentid,omitempty" xml:"cp_parentid,omitempty"`
	// 产品产地
	Cparea *string `json:"cparea,omitempty" xml:"cparea,omitempty"`
	// 条形码
	Cpbarcode *string `json:"cpbarcode,omitempty" xml:"cpbarcode,omitempty"`
	// 产品品牌
	Cpbrand *string `json:"cpbrand,omitempty" xml:"cpbrand,omitempty"`
	// 产品说明
	Cpcontent *string `json:"cpcontent,omitempty" xml:"cpcontent,omitempty"`
	// 产品规格
	Cpguige *string `json:"cpguige,omitempty" xml:"cpguige,omitempty"`
	// 产品图片
	Cpimg *string `json:"cpimg,omitempty" xml:"cpimg,omitempty"`
	// 产品名称
	Cpname *string `json:"cpname,omitempty" xml:"cpname,omitempty"`
	// 产品编号
	Cpno *string `json:"cpno,omitempty" xml:"cpno,omitempty"`
	// 产品备注
	Cpremark *string `json:"cpremark,omitempty" xml:"cpremark,omitempty"`
	// 产品型号
	Cptype *string `json:"cptype,omitempty" xml:"cptype,omitempty"`
	// 产品单位
	Cpunit *string `json:"cpunit,omitempty" xml:"cpunit,omitempty"`
	// 产品重量
	Cpweight *string `json:"cpweight,omitempty" xml:"cpweight,omitempty"`
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 默认供应商
	Gysid *string `json:"gysid,omitempty" xml:"gysid,omitempty"`
	// 批次号管理（是，否）
	Ispicimanage *string `json:"ispicimanage,omitempty" xml:"ispicimanage,omitempty"`
	// 序列号管理（是，否）
	Issnmanage *string `json:"issnmanage,omitempty" xml:"issnmanage,omitempty"`
	// 是否算库存（计算，不计算，计算(按基准规格)）
	Isstock *string `json:"isstock,omitempty" xml:"isstock,omitempty"`
	// 产品状态（正常，停售，下架）
	Isstop *string `json:"isstop,omitempty" xml:"isstop,omitempty"`
	// 零售价格
	Preprice1 *string `json:"preprice1,omitempty" xml:"preprice1,omitempty"`
	// 预设价格1
	Preprice2 *string `json:"preprice2,omitempty" xml:"preprice2,omitempty"`
	// 预设价格2
	Preprice3 *string `json:"preprice3,omitempty" xml:"preprice3,omitempty"`
	// 预设价格3
	Preprice4 *string `json:"preprice4,omitempty" xml:"preprice4,omitempty"`
	// 库存下限
	Stockdown *string `json:"stockdown,omitempty" xml:"stockdown,omitempty"`
	// 库存上限
	Stockup *string `json:"stockup,omitempty" xml:"stockup,omitempty"`
	// 产品类别
	Typeid *string `json:"typeid,omitempty" xml:"typeid,omitempty"`
	// 单位换算
	Unitrate *string `json:"unitrate,omitempty" xml:"unitrate,omitempty"`
}

func (s EditGoodsRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditGoodsRequestData) GoString() string {
	return s.String()
}

func (s *EditGoodsRequestData) SetAddedtime(v string) *EditGoodsRequestData {
	s.Addedtime = &v
	return s
}

func (s *EditGoodsRequestData) SetCbprice(v string) *EditGoodsRequestData {
	s.Cbprice = &v
	return s
}

func (s *EditGoodsRequestData) SetCpParentid(v string) *EditGoodsRequestData {
	s.CpParentid = &v
	return s
}

func (s *EditGoodsRequestData) SetCparea(v string) *EditGoodsRequestData {
	s.Cparea = &v
	return s
}

func (s *EditGoodsRequestData) SetCpbarcode(v string) *EditGoodsRequestData {
	s.Cpbarcode = &v
	return s
}

func (s *EditGoodsRequestData) SetCpbrand(v string) *EditGoodsRequestData {
	s.Cpbrand = &v
	return s
}

func (s *EditGoodsRequestData) SetCpcontent(v string) *EditGoodsRequestData {
	s.Cpcontent = &v
	return s
}

func (s *EditGoodsRequestData) SetCpguige(v string) *EditGoodsRequestData {
	s.Cpguige = &v
	return s
}

func (s *EditGoodsRequestData) SetCpimg(v string) *EditGoodsRequestData {
	s.Cpimg = &v
	return s
}

func (s *EditGoodsRequestData) SetCpname(v string) *EditGoodsRequestData {
	s.Cpname = &v
	return s
}

func (s *EditGoodsRequestData) SetCpno(v string) *EditGoodsRequestData {
	s.Cpno = &v
	return s
}

func (s *EditGoodsRequestData) SetCpremark(v string) *EditGoodsRequestData {
	s.Cpremark = &v
	return s
}

func (s *EditGoodsRequestData) SetCptype(v string) *EditGoodsRequestData {
	s.Cptype = &v
	return s
}

func (s *EditGoodsRequestData) SetCpunit(v string) *EditGoodsRequestData {
	s.Cpunit = &v
	return s
}

func (s *EditGoodsRequestData) SetCpweight(v string) *EditGoodsRequestData {
	s.Cpweight = &v
	return s
}

func (s *EditGoodsRequestData) SetDataUserid(v string) *EditGoodsRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditGoodsRequestData) SetGysid(v string) *EditGoodsRequestData {
	s.Gysid = &v
	return s
}

func (s *EditGoodsRequestData) SetIspicimanage(v string) *EditGoodsRequestData {
	s.Ispicimanage = &v
	return s
}

func (s *EditGoodsRequestData) SetIssnmanage(v string) *EditGoodsRequestData {
	s.Issnmanage = &v
	return s
}

func (s *EditGoodsRequestData) SetIsstock(v string) *EditGoodsRequestData {
	s.Isstock = &v
	return s
}

func (s *EditGoodsRequestData) SetIsstop(v string) *EditGoodsRequestData {
	s.Isstop = &v
	return s
}

func (s *EditGoodsRequestData) SetPreprice1(v string) *EditGoodsRequestData {
	s.Preprice1 = &v
	return s
}

func (s *EditGoodsRequestData) SetPreprice2(v string) *EditGoodsRequestData {
	s.Preprice2 = &v
	return s
}

func (s *EditGoodsRequestData) SetPreprice3(v string) *EditGoodsRequestData {
	s.Preprice3 = &v
	return s
}

func (s *EditGoodsRequestData) SetPreprice4(v string) *EditGoodsRequestData {
	s.Preprice4 = &v
	return s
}

func (s *EditGoodsRequestData) SetStockdown(v string) *EditGoodsRequestData {
	s.Stockdown = &v
	return s
}

func (s *EditGoodsRequestData) SetStockup(v string) *EditGoodsRequestData {
	s.Stockup = &v
	return s
}

func (s *EditGoodsRequestData) SetTypeid(v string) *EditGoodsRequestData {
	s.Typeid = &v
	return s
}

func (s *EditGoodsRequestData) SetUnitrate(v string) *EditGoodsRequestData {
	s.Unitrate = &v
	return s
}

type EditGoodsResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditGoodsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditGoodsResponseBody) GoString() string {
	return s.String()
}

func (s *EditGoodsResponseBody) SetMsgid(v int64) *EditGoodsResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditGoodsResponseBody) SetTime(v string) *EditGoodsResponseBody {
	s.Time = &v
	return s
}

type EditGoodsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditGoodsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditGoodsResponse) String() string {
	return tea.Prettify(s)
}

func (s EditGoodsResponse) GoString() string {
	return s.String()
}

func (s *EditGoodsResponse) SetHeaders(v map[string]*string) *EditGoodsResponse {
	s.Headers = v
	return s
}

func (s *EditGoodsResponse) SetBody(v *EditGoodsResponseBody) *EditGoodsResponse {
	s.Body = v
	return s
}

type EditIntostockHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditIntostockHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditIntostockHeaders) GoString() string {
	return s.String()
}

func (s *EditIntostockHeaders) SetCommonHeaders(v map[string]*string) *EditIntostockHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditIntostockHeaders) SetXAcsDingtalkAccessToken(v string) *EditIntostockHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditIntostockRequest struct {
	// 编辑数据
	Data *EditIntostockRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写189
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditIntostockRequest) String() string {
	return tea.Prettify(s)
}

func (s EditIntostockRequest) GoString() string {
	return s.String()
}

func (s *EditIntostockRequest) SetData(v *EditIntostockRequestData) *EditIntostockRequest {
	s.Data = v
	return s
}

func (s *EditIntostockRequest) SetDatatype(v int64) *EditIntostockRequest {
	s.Datatype = &v
	return s
}

func (s *EditIntostockRequest) SetMsgid(v int64) *EditIntostockRequest {
	s.Msgid = &v
	return s
}

func (s *EditIntostockRequest) SetStamp(v int64) *EditIntostockRequest {
	s.Stamp = &v
	return s
}

type EditIntostockRequestData struct {
	// 入库申请人
	Askempid *string `json:"askempid,omitempty" xml:"askempid,omitempty"`
	// 入库备注
	Auditreson *string `json:"auditreson,omitempty" xml:"auditreson,omitempty"`
	// 入库单号
	Billno *string `json:"billno,omitempty" xml:"billno,omitempty"`
	// 产品明细，json格式
	ChildMx *string `json:"child_mx,omitempty" xml:"child_mx,omitempty"`
	// 供应商/客户
	Customerid *string `json:"customerid,omitempty" xml:"customerid,omitempty"`
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 入库经办人
	Empid *string `json:"empid,omitempty" xml:"empid,omitempty"`
	// 单据类型（入库，销售退货，生产退料，生产入库，维修退货）
	Inorout *string `json:"inorout,omitempty" xml:"inorout,omitempty"`
	// 入库日期
	Libiodate *string `json:"libiodate,omitempty" xml:"libiodate,omitempty"`
	// 入库主题
	Libioname *string `json:"libioname,omitempty" xml:"libioname,omitempty"`
	// 入库状态（未入库，已入库）
	Libiostate *string `json:"libiostate,omitempty" xml:"libiostate,omitempty"`
	// 对应单据
	Orderid *string `json:"orderid,omitempty" xml:"orderid,omitempty"`
	// 申请备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 入库仓库
	Stocklibid *string `json:"stocklibid,omitempty" xml:"stocklibid,omitempty"`
}

func (s EditIntostockRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditIntostockRequestData) GoString() string {
	return s.String()
}

func (s *EditIntostockRequestData) SetAskempid(v string) *EditIntostockRequestData {
	s.Askempid = &v
	return s
}

func (s *EditIntostockRequestData) SetAuditreson(v string) *EditIntostockRequestData {
	s.Auditreson = &v
	return s
}

func (s *EditIntostockRequestData) SetBillno(v string) *EditIntostockRequestData {
	s.Billno = &v
	return s
}

func (s *EditIntostockRequestData) SetChildMx(v string) *EditIntostockRequestData {
	s.ChildMx = &v
	return s
}

func (s *EditIntostockRequestData) SetCustomerid(v string) *EditIntostockRequestData {
	s.Customerid = &v
	return s
}

func (s *EditIntostockRequestData) SetDataUserid(v string) *EditIntostockRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditIntostockRequestData) SetEmpid(v string) *EditIntostockRequestData {
	s.Empid = &v
	return s
}

func (s *EditIntostockRequestData) SetInorout(v string) *EditIntostockRequestData {
	s.Inorout = &v
	return s
}

func (s *EditIntostockRequestData) SetLibiodate(v string) *EditIntostockRequestData {
	s.Libiodate = &v
	return s
}

func (s *EditIntostockRequestData) SetLibioname(v string) *EditIntostockRequestData {
	s.Libioname = &v
	return s
}

func (s *EditIntostockRequestData) SetLibiostate(v string) *EditIntostockRequestData {
	s.Libiostate = &v
	return s
}

func (s *EditIntostockRequestData) SetOrderid(v string) *EditIntostockRequestData {
	s.Orderid = &v
	return s
}

func (s *EditIntostockRequestData) SetRemark(v string) *EditIntostockRequestData {
	s.Remark = &v
	return s
}

func (s *EditIntostockRequestData) SetStocklibid(v string) *EditIntostockRequestData {
	s.Stocklibid = &v
	return s
}

type EditIntostockResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditIntostockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditIntostockResponseBody) GoString() string {
	return s.String()
}

func (s *EditIntostockResponseBody) SetMsgid(v int64) *EditIntostockResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditIntostockResponseBody) SetTime(v string) *EditIntostockResponseBody {
	s.Time = &v
	return s
}

type EditIntostockResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditIntostockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditIntostockResponse) String() string {
	return tea.Prettify(s)
}

func (s EditIntostockResponse) GoString() string {
	return s.String()
}

func (s *EditIntostockResponse) SetHeaders(v map[string]*string) *EditIntostockResponse {
	s.Headers = v
	return s
}

func (s *EditIntostockResponse) SetBody(v *EditIntostockResponseBody) *EditIntostockResponse {
	s.Body = v
	return s
}

type EditInvoiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditInvoiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditInvoiceHeaders) GoString() string {
	return s.String()
}

func (s *EditInvoiceHeaders) SetCommonHeaders(v map[string]*string) *EditInvoiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditInvoiceHeaders) SetXAcsDingtalkAccessToken(v string) *EditInvoiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditInvoiceRequest struct {
	// 编辑数据
	Data *EditInvoiceRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写169
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s EditInvoiceRequest) GoString() string {
	return s.String()
}

func (s *EditInvoiceRequest) SetData(v *EditInvoiceRequestData) *EditInvoiceRequest {
	s.Data = v
	return s
}

func (s *EditInvoiceRequest) SetDatatype(v int64) *EditInvoiceRequest {
	s.Datatype = &v
	return s
}

func (s *EditInvoiceRequest) SetMsgid(v int64) *EditInvoiceRequest {
	s.Msgid = &v
	return s
}

func (s *EditInvoiceRequest) SetStamp(v int64) *EditInvoiceRequest {
	s.Stamp = &v
	return s
}

type EditInvoiceRequestData struct {
	// 产品明细，json格式
	ChildMx *string `json:"child_mx,omitempty" xml:"child_mx,omitempty"`
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 地址
	FhAddress *string `json:"fh_address,omitempty" xml:"fh_address,omitempty"`
	// 对应客户
	FhCustomerid *string `json:"fh_customerid,omitempty" xml:"fh_customerid,omitempty"`
	// 发货日期
	FhDate *string `json:"fh_date,omitempty" xml:"fh_date,omitempty"`
	// Email
	FhEmail *string `json:"fh_email,omitempty" xml:"fh_email,omitempty"`
	// 手机
	FhHandset *string `json:"fh_handset,omitempty" xml:"fh_handset,omitempty"`
	// 对应订单
	FhHtorder *string `json:"fh_htorder,omitempty" xml:"fh_htorder,omitempty"`
	// 打包件数
	FhJianshu *string `json:"fh_jianshu,omitempty" xml:"fh_jianshu,omitempty"`
	// 重量(Kg)
	FhKg *string `json:"fh_kg,omitempty" xml:"fh_kg,omitempty"`
	// 收货人
	FhLinkman *string `json:"fh_linkman,omitempty" xml:"fh_linkman,omitempty"`
	// 联系人
	FhLxrid *string `json:"fh_lxrid,omitempty" xml:"fh_lxrid,omitempty"`
	// 发货方式
	FhMode *string `json:"fh_mode,omitempty" xml:"fh_mode,omitempty"`
	// MSN
	FhMsn *string `json:"fh_msn,omitempty" xml:"fh_msn,omitempty"`
	// 发货单号
	FhNumber *string `json:"fh_number,omitempty" xml:"fh_number,omitempty"`
	// 邮编
	FhPost *string `json:"fh_post,omitempty" xml:"fh_post,omitempty"`
	// 所有者
	FhPreside *string `json:"fh_preside,omitempty" xml:"fh_preside,omitempty"`
	// 备注
	FhRemark *string `json:"fh_remark,omitempty" xml:"fh_remark,omitempty"`
	// 发货人
	FhShipper *string `json:"fh_shipper,omitempty" xml:"fh_shipper,omitempty"`
	// 发货状态
	FhState *string `json:"fh_state,omitempty" xml:"fh_state,omitempty"`
	// 电话
	FhTel *string `json:"fh_tel,omitempty" xml:"fh_tel,omitempty"`
	// 发货主题
	FhTitle *string `json:"fh_title,omitempty" xml:"fh_title,omitempty"`
	// 运费
	FhYunfei *string `json:"fh_yunfei,omitempty" xml:"fh_yunfei,omitempty"`
}

func (s EditInvoiceRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditInvoiceRequestData) GoString() string {
	return s.String()
}

func (s *EditInvoiceRequestData) SetChildMx(v string) *EditInvoiceRequestData {
	s.ChildMx = &v
	return s
}

func (s *EditInvoiceRequestData) SetDataUserid(v string) *EditInvoiceRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhAddress(v string) *EditInvoiceRequestData {
	s.FhAddress = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhCustomerid(v string) *EditInvoiceRequestData {
	s.FhCustomerid = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhDate(v string) *EditInvoiceRequestData {
	s.FhDate = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhEmail(v string) *EditInvoiceRequestData {
	s.FhEmail = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhHandset(v string) *EditInvoiceRequestData {
	s.FhHandset = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhHtorder(v string) *EditInvoiceRequestData {
	s.FhHtorder = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhJianshu(v string) *EditInvoiceRequestData {
	s.FhJianshu = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhKg(v string) *EditInvoiceRequestData {
	s.FhKg = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhLinkman(v string) *EditInvoiceRequestData {
	s.FhLinkman = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhLxrid(v string) *EditInvoiceRequestData {
	s.FhLxrid = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhMode(v string) *EditInvoiceRequestData {
	s.FhMode = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhMsn(v string) *EditInvoiceRequestData {
	s.FhMsn = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhNumber(v string) *EditInvoiceRequestData {
	s.FhNumber = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhPost(v string) *EditInvoiceRequestData {
	s.FhPost = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhPreside(v string) *EditInvoiceRequestData {
	s.FhPreside = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhRemark(v string) *EditInvoiceRequestData {
	s.FhRemark = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhShipper(v string) *EditInvoiceRequestData {
	s.FhShipper = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhState(v string) *EditInvoiceRequestData {
	s.FhState = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhTel(v string) *EditInvoiceRequestData {
	s.FhTel = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhTitle(v string) *EditInvoiceRequestData {
	s.FhTitle = &v
	return s
}

func (s *EditInvoiceRequestData) SetFhYunfei(v string) *EditInvoiceRequestData {
	s.FhYunfei = &v
	return s
}

type EditInvoiceResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *EditInvoiceResponseBody) SetMsgid(v int64) *EditInvoiceResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditInvoiceResponseBody) SetTime(v string) *EditInvoiceResponseBody {
	s.Time = &v
	return s
}

type EditInvoiceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s EditInvoiceResponse) GoString() string {
	return s.String()
}

func (s *EditInvoiceResponse) SetHeaders(v map[string]*string) *EditInvoiceResponse {
	s.Headers = v
	return s
}

func (s *EditInvoiceResponse) SetBody(v *EditInvoiceResponseBody) *EditInvoiceResponse {
	s.Body = v
	return s
}

type EditOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditOrderHeaders) GoString() string {
	return s.String()
}

func (s *EditOrderHeaders) SetCommonHeaders(v map[string]*string) *EditOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditOrderHeaders) SetXAcsDingtalkAccessToken(v string) *EditOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditOrderRequest struct {
	// 编辑数据
	Data *EditOrderRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写150
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s EditOrderRequest) GoString() string {
	return s.String()
}

func (s *EditOrderRequest) SetData(v *EditOrderRequestData) *EditOrderRequest {
	s.Data = v
	return s
}

func (s *EditOrderRequest) SetDatatype(v int64) *EditOrderRequest {
	s.Datatype = &v
	return s
}

func (s *EditOrderRequest) SetMsgid(v int64) *EditOrderRequest {
	s.Msgid = &v
	return s
}

func (s *EditOrderRequest) SetStamp(v int64) *EditOrderRequest {
	s.Stamp = &v
	return s
}

type EditOrderRequestData struct {
	// 产品明细，json格式
	ChildMx *string `json:"child_mx,omitempty" xml:"child_mx,omitempty"`
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 收货地址
	Fahuoaddressid *string `json:"fahuoaddressid,omitempty" xml:"fahuoaddressid,omitempty"`
	// 开始日期
	HtBegindate *string `json:"ht_begindate,omitempty" xml:"ht_begindate,omitempty"`
	// 合同正文
	HtContract *string `json:"ht_contract,omitempty" xml:"ht_contract,omitempty"`
	// 对应客户
	HtCustomerid *string `json:"ht_customerid,omitempty" xml:"ht_customerid,omitempty"`
	// 客户签约人
	HtCusub *string `json:"ht_cusub,omitempty" xml:"ht_cusub,omitempty"`
	// 签单日期
	HtDate *string `json:"ht_date,omitempty" xml:"ht_date,omitempty"`
	// 交付地点
	HtDeliplace *string `json:"ht_deliplace,omitempty" xml:"ht_deliplace,omitempty"`
	// 最晚发货日
	HtEnddate *string `json:"ht_enddate,omitempty" xml:"ht_enddate,omitempty"`
	// 附加费用金额
	HtFjmoney *string `json:"ht_fjmoney,omitempty" xml:"ht_fjmoney,omitempty"`
	// 附加费用分类
	HtFjmoneylx *string `json:"ht_fjmoneylx,omitempty" xml:"ht_fjmoneylx,omitempty"`
	// 优惠抹零金额
	HtKjmoney *string `json:"ht_kjmoney,omitempty" xml:"ht_kjmoney,omitempty"`
	// 对应联系人
	HtLxrid *string `json:"ht_lxrid,omitempty" xml:"ht_lxrid,omitempty"`
	// 联系方式
	HtLxrinfo *string `json:"ht_lxrinfo,omitempty" xml:"ht_lxrinfo,omitempty"`
	// 优惠折扣率
	HtMoneyzhekou *string `json:"ht_moneyzhekou,omitempty" xml:"ht_moneyzhekou,omitempty"`
	// 合同单号
	HtNumber *string `json:"ht_number,omitempty" xml:"ht_number,omitempty"`
	// 单据类型（合同，合同订单，店面单）
	HtOrder *string `json:"ht_order,omitempty" xml:"ht_order,omitempty"`
	// 付款方式
	HtPaymode *string `json:"ht_paymode,omitempty" xml:"ht_paymode,omitempty"`
	// 所有者
	HtPreside *string `json:"ht_preside,omitempty" xml:"ht_preside,omitempty"`
	// 备注
	HtRemark *string `json:"ht_remark,omitempty" xml:"ht_remark,omitempty"`
	// 状态
	HtState *string `json:"ht_state,omitempty" xml:"ht_state,omitempty"`
	// 外币备注
	HtSummemo *string `json:"ht_summemo,omitempty" xml:"ht_summemo,omitempty"`
	// 总金额
	HtSummoney *string `json:"ht_summoney,omitempty" xml:"ht_summoney,omitempty"`
	// 主题
	HtTitle *string `json:"ht_title,omitempty" xml:"ht_title,omitempty"`
	// 分类
	HtType *string `json:"ht_type,omitempty" xml:"ht_type,omitempty"`
	// 我方签约人
	HtWesub *string `json:"ht_wesub,omitempty" xml:"ht_wesub,omitempty"`
	// 发货方式
	HtWuliutype *string `json:"ht_wuliutype,omitempty" xml:"ht_wuliutype,omitempty"`
	// 对应机会
	HtXshid *string `json:"ht_xshid,omitempty" xml:"ht_xshid,omitempty"`
	// 预计运费
	HtYunfeimoney *string `json:"ht_yunfeimoney,omitempty" xml:"ht_yunfeimoney,omitempty"`
}

func (s EditOrderRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditOrderRequestData) GoString() string {
	return s.String()
}

func (s *EditOrderRequestData) SetChildMx(v string) *EditOrderRequestData {
	s.ChildMx = &v
	return s
}

func (s *EditOrderRequestData) SetDataUserid(v string) *EditOrderRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditOrderRequestData) SetFahuoaddressid(v string) *EditOrderRequestData {
	s.Fahuoaddressid = &v
	return s
}

func (s *EditOrderRequestData) SetHtBegindate(v string) *EditOrderRequestData {
	s.HtBegindate = &v
	return s
}

func (s *EditOrderRequestData) SetHtContract(v string) *EditOrderRequestData {
	s.HtContract = &v
	return s
}

func (s *EditOrderRequestData) SetHtCustomerid(v string) *EditOrderRequestData {
	s.HtCustomerid = &v
	return s
}

func (s *EditOrderRequestData) SetHtCusub(v string) *EditOrderRequestData {
	s.HtCusub = &v
	return s
}

func (s *EditOrderRequestData) SetHtDate(v string) *EditOrderRequestData {
	s.HtDate = &v
	return s
}

func (s *EditOrderRequestData) SetHtDeliplace(v string) *EditOrderRequestData {
	s.HtDeliplace = &v
	return s
}

func (s *EditOrderRequestData) SetHtEnddate(v string) *EditOrderRequestData {
	s.HtEnddate = &v
	return s
}

func (s *EditOrderRequestData) SetHtFjmoney(v string) *EditOrderRequestData {
	s.HtFjmoney = &v
	return s
}

func (s *EditOrderRequestData) SetHtFjmoneylx(v string) *EditOrderRequestData {
	s.HtFjmoneylx = &v
	return s
}

func (s *EditOrderRequestData) SetHtKjmoney(v string) *EditOrderRequestData {
	s.HtKjmoney = &v
	return s
}

func (s *EditOrderRequestData) SetHtLxrid(v string) *EditOrderRequestData {
	s.HtLxrid = &v
	return s
}

func (s *EditOrderRequestData) SetHtLxrinfo(v string) *EditOrderRequestData {
	s.HtLxrinfo = &v
	return s
}

func (s *EditOrderRequestData) SetHtMoneyzhekou(v string) *EditOrderRequestData {
	s.HtMoneyzhekou = &v
	return s
}

func (s *EditOrderRequestData) SetHtNumber(v string) *EditOrderRequestData {
	s.HtNumber = &v
	return s
}

func (s *EditOrderRequestData) SetHtOrder(v string) *EditOrderRequestData {
	s.HtOrder = &v
	return s
}

func (s *EditOrderRequestData) SetHtPaymode(v string) *EditOrderRequestData {
	s.HtPaymode = &v
	return s
}

func (s *EditOrderRequestData) SetHtPreside(v string) *EditOrderRequestData {
	s.HtPreside = &v
	return s
}

func (s *EditOrderRequestData) SetHtRemark(v string) *EditOrderRequestData {
	s.HtRemark = &v
	return s
}

func (s *EditOrderRequestData) SetHtState(v string) *EditOrderRequestData {
	s.HtState = &v
	return s
}

func (s *EditOrderRequestData) SetHtSummemo(v string) *EditOrderRequestData {
	s.HtSummemo = &v
	return s
}

func (s *EditOrderRequestData) SetHtSummoney(v string) *EditOrderRequestData {
	s.HtSummoney = &v
	return s
}

func (s *EditOrderRequestData) SetHtTitle(v string) *EditOrderRequestData {
	s.HtTitle = &v
	return s
}

func (s *EditOrderRequestData) SetHtType(v string) *EditOrderRequestData {
	s.HtType = &v
	return s
}

func (s *EditOrderRequestData) SetHtWesub(v string) *EditOrderRequestData {
	s.HtWesub = &v
	return s
}

func (s *EditOrderRequestData) SetHtWuliutype(v string) *EditOrderRequestData {
	s.HtWuliutype = &v
	return s
}

func (s *EditOrderRequestData) SetHtXshid(v string) *EditOrderRequestData {
	s.HtXshid = &v
	return s
}

func (s *EditOrderRequestData) SetHtYunfeimoney(v string) *EditOrderRequestData {
	s.HtYunfeimoney = &v
	return s
}

type EditOrderResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditOrderResponseBody) GoString() string {
	return s.String()
}

func (s *EditOrderResponseBody) SetMsgid(v int64) *EditOrderResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditOrderResponseBody) SetTime(v string) *EditOrderResponseBody {
	s.Time = &v
	return s
}

type EditOrderResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s EditOrderResponse) GoString() string {
	return s.String()
}

func (s *EditOrderResponse) SetHeaders(v map[string]*string) *EditOrderResponse {
	s.Headers = v
	return s
}

func (s *EditOrderResponse) SetBody(v *EditOrderResponseBody) *EditOrderResponse {
	s.Body = v
	return s
}

type EditOutstockHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditOutstockHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditOutstockHeaders) GoString() string {
	return s.String()
}

func (s *EditOutstockHeaders) SetCommonHeaders(v map[string]*string) *EditOutstockHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditOutstockHeaders) SetXAcsDingtalkAccessToken(v string) *EditOutstockHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditOutstockRequest struct {
	// 编辑数据
	Data *EditOutstockRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写191
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditOutstockRequest) String() string {
	return tea.Prettify(s)
}

func (s EditOutstockRequest) GoString() string {
	return s.String()
}

func (s *EditOutstockRequest) SetData(v *EditOutstockRequestData) *EditOutstockRequest {
	s.Data = v
	return s
}

func (s *EditOutstockRequest) SetDatatype(v int64) *EditOutstockRequest {
	s.Datatype = &v
	return s
}

func (s *EditOutstockRequest) SetMsgid(v int64) *EditOutstockRequest {
	s.Msgid = &v
	return s
}

func (s *EditOutstockRequest) SetStamp(v int64) *EditOutstockRequest {
	s.Stamp = &v
	return s
}

type EditOutstockRequestData struct {
	// 申请人
	Askempid *string `json:"askempid,omitempty" xml:"askempid,omitempty"`
	// 出库备注
	Auditreson *string `json:"auditreson,omitempty" xml:"auditreson,omitempty"`
	// 出库单号
	Billno *string `json:"billno,omitempty" xml:"billno,omitempty"`
	// 产品明细，json格式
	ChildMx *string `json:"child_mx,omitempty" xml:"child_mx,omitempty"`
	// 对应客户
	Customerid *string `json:"customerid,omitempty" xml:"customerid,omitempty"`
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 经办人
	Empid *string `json:"empid,omitempty" xml:"empid,omitempty"`
	// 单据类型
	Inorout *string `json:"inorout,omitempty" xml:"inorout,omitempty"`
	// 出库日期
	Libiodate *string `json:"libiodate,omitempty" xml:"libiodate,omitempty"`
	// 出库主题
	Libioname *string `json:"libioname,omitempty" xml:"libioname,omitempty"`
	// 出库状态
	Libiostate *string `json:"libiostate,omitempty" xml:"libiostate,omitempty"`
	// 对应订单
	Orderid *string `json:"orderid,omitempty" xml:"orderid,omitempty"`
	// 申请备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 出库仓库
	Stocklibid *string `json:"stocklibid,omitempty" xml:"stocklibid,omitempty"`
}

func (s EditOutstockRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditOutstockRequestData) GoString() string {
	return s.String()
}

func (s *EditOutstockRequestData) SetAskempid(v string) *EditOutstockRequestData {
	s.Askempid = &v
	return s
}

func (s *EditOutstockRequestData) SetAuditreson(v string) *EditOutstockRequestData {
	s.Auditreson = &v
	return s
}

func (s *EditOutstockRequestData) SetBillno(v string) *EditOutstockRequestData {
	s.Billno = &v
	return s
}

func (s *EditOutstockRequestData) SetChildMx(v string) *EditOutstockRequestData {
	s.ChildMx = &v
	return s
}

func (s *EditOutstockRequestData) SetCustomerid(v string) *EditOutstockRequestData {
	s.Customerid = &v
	return s
}

func (s *EditOutstockRequestData) SetDataUserid(v string) *EditOutstockRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditOutstockRequestData) SetEmpid(v string) *EditOutstockRequestData {
	s.Empid = &v
	return s
}

func (s *EditOutstockRequestData) SetInorout(v string) *EditOutstockRequestData {
	s.Inorout = &v
	return s
}

func (s *EditOutstockRequestData) SetLibiodate(v string) *EditOutstockRequestData {
	s.Libiodate = &v
	return s
}

func (s *EditOutstockRequestData) SetLibioname(v string) *EditOutstockRequestData {
	s.Libioname = &v
	return s
}

func (s *EditOutstockRequestData) SetLibiostate(v string) *EditOutstockRequestData {
	s.Libiostate = &v
	return s
}

func (s *EditOutstockRequestData) SetOrderid(v string) *EditOutstockRequestData {
	s.Orderid = &v
	return s
}

func (s *EditOutstockRequestData) SetRemark(v string) *EditOutstockRequestData {
	s.Remark = &v
	return s
}

func (s *EditOutstockRequestData) SetStocklibid(v string) *EditOutstockRequestData {
	s.Stocklibid = &v
	return s
}

type EditOutstockResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditOutstockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditOutstockResponseBody) GoString() string {
	return s.String()
}

func (s *EditOutstockResponseBody) SetMsgid(v int64) *EditOutstockResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditOutstockResponseBody) SetTime(v string) *EditOutstockResponseBody {
	s.Time = &v
	return s
}

type EditOutstockResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditOutstockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditOutstockResponse) String() string {
	return tea.Prettify(s)
}

func (s EditOutstockResponse) GoString() string {
	return s.String()
}

func (s *EditOutstockResponse) SetHeaders(v map[string]*string) *EditOutstockResponse {
	s.Headers = v
	return s
}

func (s *EditOutstockResponse) SetBody(v *EditOutstockResponseBody) *EditOutstockResponse {
	s.Body = v
	return s
}

type EditProductionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditProductionHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditProductionHeaders) GoString() string {
	return s.String()
}

func (s *EditProductionHeaders) SetCommonHeaders(v map[string]*string) *EditProductionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditProductionHeaders) SetXAcsDingtalkAccessToken(v string) *EditProductionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditProductionRequest struct {
	// 编辑数据
	Data *EditProductionRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写156
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditProductionRequest) String() string {
	return tea.Prettify(s)
}

func (s EditProductionRequest) GoString() string {
	return s.String()
}

func (s *EditProductionRequest) SetData(v *EditProductionRequestData) *EditProductionRequest {
	s.Data = v
	return s
}

func (s *EditProductionRequest) SetDatatype(v int64) *EditProductionRequest {
	s.Datatype = &v
	return s
}

func (s *EditProductionRequest) SetMsgid(v int64) *EditProductionRequest {
	s.Msgid = &v
	return s
}

func (s *EditProductionRequest) SetStamp(v int64) *EditProductionRequest {
	s.Stamp = &v
	return s
}

type EditProductionRequestData struct {
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 对应客户
	SchCustomerid *string `json:"sch_customerid,omitempty" xml:"sch_customerid,omitempty"`
	// 完成日期
	SchEndtime *string `json:"sch_endtime,omitempty" xml:"sch_endtime,omitempty"`
	// 状态（未生产，生产中，生产中止，生产完成）
	SchFinished *string `json:"sch_finished,omitempty" xml:"sch_finished,omitempty"`
	// 订单
	SchHtid *string `json:"sch_htid,omitempty" xml:"sch_htid,omitempty"`
	// 生产人员
	SchMakeemp *string `json:"sch_makeemp,omitempty" xml:"sch_makeemp,omitempty"`
	// 单号
	SchNumber *string `json:"sch_number,omitempty" xml:"sch_number,omitempty"`
	// 计划完成
	SchPlanendtime *string `json:"sch_planendtime,omitempty" xml:"sch_planendtime,omitempty"`
	// 负责人
	SchPrincipal *string `json:"sch_principal,omitempty" xml:"sch_principal,omitempty"`
	// 备注
	SchRemark *string `json:"sch_remark,omitempty" xml:"sch_remark,omitempty"`
	// 开始日期
	SchStarttime *string `json:"sch_starttime,omitempty" xml:"sch_starttime,omitempty"`
	// 阶段（计划，审核，领料，生产，验收，入库/退料，结单，取消）
	SchStatesstr *string `json:"sch_statesstr,omitempty" xml:"sch_statesstr,omitempty"`
	// 主题
	SchTitle *string `json:"sch_title,omitempty" xml:"sch_title,omitempty"`
}

func (s EditProductionRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditProductionRequestData) GoString() string {
	return s.String()
}

func (s *EditProductionRequestData) SetDataUserid(v string) *EditProductionRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditProductionRequestData) SetSchCustomerid(v string) *EditProductionRequestData {
	s.SchCustomerid = &v
	return s
}

func (s *EditProductionRequestData) SetSchEndtime(v string) *EditProductionRequestData {
	s.SchEndtime = &v
	return s
}

func (s *EditProductionRequestData) SetSchFinished(v string) *EditProductionRequestData {
	s.SchFinished = &v
	return s
}

func (s *EditProductionRequestData) SetSchHtid(v string) *EditProductionRequestData {
	s.SchHtid = &v
	return s
}

func (s *EditProductionRequestData) SetSchMakeemp(v string) *EditProductionRequestData {
	s.SchMakeemp = &v
	return s
}

func (s *EditProductionRequestData) SetSchNumber(v string) *EditProductionRequestData {
	s.SchNumber = &v
	return s
}

func (s *EditProductionRequestData) SetSchPlanendtime(v string) *EditProductionRequestData {
	s.SchPlanendtime = &v
	return s
}

func (s *EditProductionRequestData) SetSchPrincipal(v string) *EditProductionRequestData {
	s.SchPrincipal = &v
	return s
}

func (s *EditProductionRequestData) SetSchRemark(v string) *EditProductionRequestData {
	s.SchRemark = &v
	return s
}

func (s *EditProductionRequestData) SetSchStarttime(v string) *EditProductionRequestData {
	s.SchStarttime = &v
	return s
}

func (s *EditProductionRequestData) SetSchStatesstr(v string) *EditProductionRequestData {
	s.SchStatesstr = &v
	return s
}

func (s *EditProductionRequestData) SetSchTitle(v string) *EditProductionRequestData {
	s.SchTitle = &v
	return s
}

type EditProductionResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditProductionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditProductionResponseBody) GoString() string {
	return s.String()
}

func (s *EditProductionResponseBody) SetMsgid(v int64) *EditProductionResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditProductionResponseBody) SetTime(v string) *EditProductionResponseBody {
	s.Time = &v
	return s
}

type EditProductionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditProductionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditProductionResponse) String() string {
	return tea.Prettify(s)
}

func (s EditProductionResponse) GoString() string {
	return s.String()
}

func (s *EditProductionResponse) SetHeaders(v map[string]*string) *EditProductionResponse {
	s.Headers = v
	return s
}

func (s *EditProductionResponse) SetBody(v *EditProductionResponseBody) *EditProductionResponse {
	s.Body = v
	return s
}

type EditPurchaseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditPurchaseHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditPurchaseHeaders) GoString() string {
	return s.String()
}

func (s *EditPurchaseHeaders) SetCommonHeaders(v map[string]*string) *EditPurchaseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditPurchaseHeaders) SetXAcsDingtalkAccessToken(v string) *EditPurchaseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditPurchaseRequest struct {
	// 编辑数据
	Data *EditPurchaseRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写153
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditPurchaseRequest) String() string {
	return tea.Prettify(s)
}

func (s EditPurchaseRequest) GoString() string {
	return s.String()
}

func (s *EditPurchaseRequest) SetData(v *EditPurchaseRequestData) *EditPurchaseRequest {
	s.Data = v
	return s
}

func (s *EditPurchaseRequest) SetDatatype(v int64) *EditPurchaseRequest {
	s.Datatype = &v
	return s
}

func (s *EditPurchaseRequest) SetMsgid(v int64) *EditPurchaseRequest {
	s.Msgid = &v
	return s
}

func (s *EditPurchaseRequest) SetStamp(v int64) *EditPurchaseRequest {
	s.Stamp = &v
	return s
}

type EditPurchaseRequestData struct {
	// 附加费用金额
	CgFjmoney *string `json:"cg_fjmoney,omitempty" xml:"cg_fjmoney,omitempty"`
	// 附加费用分类
	CgFjmoneylx *string `json:"cg_fjmoneylx,omitempty" xml:"cg_fjmoneylx,omitempty"`
	// 优惠抹零金额
	CgKjmoney *string `json:"cg_kjmoney,omitempty" xml:"cg_kjmoney,omitempty"`
	// 优惠折扣率
	CgMoneyzhekou *string `json:"cg_moneyzhekou,omitempty" xml:"cg_moneyzhekou,omitempty"`
	// 执行状态
	CgZxstate *string `json:"cg_zxstate,omitempty" xml:"cg_zxstate,omitempty"`
	// 采购日期
	Cgdate *string `json:"cgdate,omitempty" xml:"cgdate,omitempty"`
	// 采购主题
	Cgname *string `json:"cgname,omitempty" xml:"cgname,omitempty"`
	// 采购单号
	Cgno *string `json:"cgno,omitempty" xml:"cgno,omitempty"`
	// 采购摘要
	Cgremark *string `json:"cgremark,omitempty" xml:"cgremark,omitempty"`
	// 采购分类
	Cgtype *string `json:"cgtype,omitempty" xml:"cgtype,omitempty"`
	// 产品明细，json格式
	ChildMx *string `json:"child_mx,omitempty" xml:"child_mx,omitempty"`
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 我方代表
	Empid *string `json:"empid,omitempty" xml:"empid,omitempty"`
	// 供应商联系人
	GysLxrid *string `json:"gys_lxrid,omitempty" xml:"gys_lxrid,omitempty"`
	// 联系方式
	GysLxrinfo *string `json:"gys_lxrinfo,omitempty" xml:"gys_lxrinfo,omitempty"`
	// 供应商
	Gysid *string `json:"gysid,omitempty" xml:"gysid,omitempty"`
	// 供应商代表
	Gysjingban *string `json:"gysjingban,omitempty" xml:"gysjingban,omitempty"`
	// 关联订单
	OrderHtid *string `json:"order_htid,omitempty" xml:"order_htid,omitempty"`
	// 关联订单客户
	OrderKhid *string `json:"order_khid,omitempty" xml:"order_khid,omitempty"`
	// 采购金额
	Summoney *string `json:"summoney,omitempty" xml:"summoney,omitempty"`
}

func (s EditPurchaseRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditPurchaseRequestData) GoString() string {
	return s.String()
}

func (s *EditPurchaseRequestData) SetCgFjmoney(v string) *EditPurchaseRequestData {
	s.CgFjmoney = &v
	return s
}

func (s *EditPurchaseRequestData) SetCgFjmoneylx(v string) *EditPurchaseRequestData {
	s.CgFjmoneylx = &v
	return s
}

func (s *EditPurchaseRequestData) SetCgKjmoney(v string) *EditPurchaseRequestData {
	s.CgKjmoney = &v
	return s
}

func (s *EditPurchaseRequestData) SetCgMoneyzhekou(v string) *EditPurchaseRequestData {
	s.CgMoneyzhekou = &v
	return s
}

func (s *EditPurchaseRequestData) SetCgZxstate(v string) *EditPurchaseRequestData {
	s.CgZxstate = &v
	return s
}

func (s *EditPurchaseRequestData) SetCgdate(v string) *EditPurchaseRequestData {
	s.Cgdate = &v
	return s
}

func (s *EditPurchaseRequestData) SetCgname(v string) *EditPurchaseRequestData {
	s.Cgname = &v
	return s
}

func (s *EditPurchaseRequestData) SetCgno(v string) *EditPurchaseRequestData {
	s.Cgno = &v
	return s
}

func (s *EditPurchaseRequestData) SetCgremark(v string) *EditPurchaseRequestData {
	s.Cgremark = &v
	return s
}

func (s *EditPurchaseRequestData) SetCgtype(v string) *EditPurchaseRequestData {
	s.Cgtype = &v
	return s
}

func (s *EditPurchaseRequestData) SetChildMx(v string) *EditPurchaseRequestData {
	s.ChildMx = &v
	return s
}

func (s *EditPurchaseRequestData) SetDataUserid(v string) *EditPurchaseRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditPurchaseRequestData) SetEmpid(v string) *EditPurchaseRequestData {
	s.Empid = &v
	return s
}

func (s *EditPurchaseRequestData) SetGysLxrid(v string) *EditPurchaseRequestData {
	s.GysLxrid = &v
	return s
}

func (s *EditPurchaseRequestData) SetGysLxrinfo(v string) *EditPurchaseRequestData {
	s.GysLxrinfo = &v
	return s
}

func (s *EditPurchaseRequestData) SetGysid(v string) *EditPurchaseRequestData {
	s.Gysid = &v
	return s
}

func (s *EditPurchaseRequestData) SetGysjingban(v string) *EditPurchaseRequestData {
	s.Gysjingban = &v
	return s
}

func (s *EditPurchaseRequestData) SetOrderHtid(v string) *EditPurchaseRequestData {
	s.OrderHtid = &v
	return s
}

func (s *EditPurchaseRequestData) SetOrderKhid(v string) *EditPurchaseRequestData {
	s.OrderKhid = &v
	return s
}

func (s *EditPurchaseRequestData) SetSummoney(v string) *EditPurchaseRequestData {
	s.Summoney = &v
	return s
}

type EditPurchaseResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditPurchaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditPurchaseResponseBody) GoString() string {
	return s.String()
}

func (s *EditPurchaseResponseBody) SetMsgid(v int64) *EditPurchaseResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditPurchaseResponseBody) SetTime(v string) *EditPurchaseResponseBody {
	s.Time = &v
	return s
}

type EditPurchaseResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditPurchaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditPurchaseResponse) String() string {
	return tea.Prettify(s)
}

func (s EditPurchaseResponse) GoString() string {
	return s.String()
}

func (s *EditPurchaseResponse) SetHeaders(v map[string]*string) *EditPurchaseResponse {
	s.Headers = v
	return s
}

func (s *EditPurchaseResponse) SetBody(v *EditPurchaseResponseBody) *EditPurchaseResponse {
	s.Body = v
	return s
}

type EditQuotationRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditQuotationRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditQuotationRecordHeaders) GoString() string {
	return s.String()
}

func (s *EditQuotationRecordHeaders) SetCommonHeaders(v map[string]*string) *EditQuotationRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditQuotationRecordHeaders) SetXAcsDingtalkAccessToken(v string) *EditQuotationRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditQuotationRecordRequest struct {
	// 编辑数据
	Data *EditQuotationRecordRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写161
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditQuotationRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s EditQuotationRecordRequest) GoString() string {
	return s.String()
}

func (s *EditQuotationRecordRequest) SetData(v *EditQuotationRecordRequestData) *EditQuotationRecordRequest {
	s.Data = v
	return s
}

func (s *EditQuotationRecordRequest) SetDatatype(v int64) *EditQuotationRecordRequest {
	s.Datatype = &v
	return s
}

func (s *EditQuotationRecordRequest) SetMsgid(v int64) *EditQuotationRecordRequest {
	s.Msgid = &v
	return s
}

func (s *EditQuotationRecordRequest) SetStamp(v int64) *EditQuotationRecordRequest {
	s.Stamp = &v
	return s
}

type EditQuotationRecordRequestData struct {
	// 报价人
	BjBjren *string `json:"bj_bjren,omitempty" xml:"bj_bjren,omitempty"`
	// 包装运输
	BjBzremark *string `json:"bj_bzremark,omitempty" xml:"bj_bzremark,omitempty"`
	// 对应客户
	BjCustomerid *string `json:"bj_customerid,omitempty" xml:"bj_customerid,omitempty"`
	// 报价日期
	BjDate *string `json:"bj_date,omitempty" xml:"bj_date,omitempty"`
	// 附加费用金额
	BjFjmoney *string `json:"bj_fjmoney,omitempty" xml:"bj_fjmoney,omitempty"`
	// 附加费用分类
	BjFjmoneylx *string `json:"bj_fjmoneylx,omitempty" xml:"bj_fjmoneylx,omitempty"`
	// 付款说明
	BjFkremark *string `json:"bj_fkremark,omitempty" xml:"bj_fkremark,omitempty"`
	// 交付说明
	BjJfremark *string `json:"bj_jfremark,omitempty" xml:"bj_jfremark,omitempty"`
	// 接收人
	BjJshren *string `json:"bj_jshren,omitempty" xml:"bj_jshren,omitempty"`
	// 优惠抹零金额
	BjKjmoney *string `json:"bj_kjmoney,omitempty" xml:"bj_kjmoney,omitempty"`
	// 联系方式
	BjLianxi *string `json:"bj_lianxi,omitempty" xml:"bj_lianxi,omitempty"`
	// 优惠折扣率
	BjMoneyzhekou *string `json:"bj_moneyzhekou,omitempty" xml:"bj_moneyzhekou,omitempty"`
	// 报价单号
	BjNumber *string `json:"bj_number,omitempty" xml:"bj_number,omitempty"`
	// 报价(总)
	BjPrice *string `json:"bj_price,omitempty" xml:"bj_price,omitempty"`
	// 备注
	BjRemark *string `json:"bj_remark,omitempty" xml:"bj_remark,omitempty"`
	// 转成订单
	BjState *string `json:"bj_state,omitempty" xml:"bj_state,omitempty"`
	// 主题
	BjTitle *string `json:"bj_title,omitempty" xml:"bj_title,omitempty"`
	// 对应机会
	BjXshid *string `json:"bj_xshid,omitempty" xml:"bj_xshid,omitempty"`
	// 产品明细，json格式
	ChildMx *string `json:"child_mx,omitempty" xml:"child_mx,omitempty"`
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
}

func (s EditQuotationRecordRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditQuotationRecordRequestData) GoString() string {
	return s.String()
}

func (s *EditQuotationRecordRequestData) SetBjBjren(v string) *EditQuotationRecordRequestData {
	s.BjBjren = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjBzremark(v string) *EditQuotationRecordRequestData {
	s.BjBzremark = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjCustomerid(v string) *EditQuotationRecordRequestData {
	s.BjCustomerid = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjDate(v string) *EditQuotationRecordRequestData {
	s.BjDate = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjFjmoney(v string) *EditQuotationRecordRequestData {
	s.BjFjmoney = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjFjmoneylx(v string) *EditQuotationRecordRequestData {
	s.BjFjmoneylx = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjFkremark(v string) *EditQuotationRecordRequestData {
	s.BjFkremark = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjJfremark(v string) *EditQuotationRecordRequestData {
	s.BjJfremark = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjJshren(v string) *EditQuotationRecordRequestData {
	s.BjJshren = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjKjmoney(v string) *EditQuotationRecordRequestData {
	s.BjKjmoney = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjLianxi(v string) *EditQuotationRecordRequestData {
	s.BjLianxi = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjMoneyzhekou(v string) *EditQuotationRecordRequestData {
	s.BjMoneyzhekou = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjNumber(v string) *EditQuotationRecordRequestData {
	s.BjNumber = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjPrice(v string) *EditQuotationRecordRequestData {
	s.BjPrice = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjRemark(v string) *EditQuotationRecordRequestData {
	s.BjRemark = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjState(v string) *EditQuotationRecordRequestData {
	s.BjState = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjTitle(v string) *EditQuotationRecordRequestData {
	s.BjTitle = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetBjXshid(v string) *EditQuotationRecordRequestData {
	s.BjXshid = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetChildMx(v string) *EditQuotationRecordRequestData {
	s.ChildMx = &v
	return s
}

func (s *EditQuotationRecordRequestData) SetDataUserid(v string) *EditQuotationRecordRequestData {
	s.DataUserid = &v
	return s
}

type EditQuotationRecordResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditQuotationRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditQuotationRecordResponseBody) GoString() string {
	return s.String()
}

func (s *EditQuotationRecordResponseBody) SetMsgid(v int64) *EditQuotationRecordResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditQuotationRecordResponseBody) SetTime(v string) *EditQuotationRecordResponseBody {
	s.Time = &v
	return s
}

type EditQuotationRecordResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditQuotationRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditQuotationRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s EditQuotationRecordResponse) GoString() string {
	return s.String()
}

func (s *EditQuotationRecordResponse) SetHeaders(v map[string]*string) *EditQuotationRecordResponse {
	s.Headers = v
	return s
}

func (s *EditQuotationRecordResponse) SetBody(v *EditQuotationRecordResponseBody) *EditQuotationRecordResponse {
	s.Body = v
	return s
}

type EditSalesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditSalesHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditSalesHeaders) GoString() string {
	return s.String()
}

func (s *EditSalesHeaders) SetCommonHeaders(v map[string]*string) *EditSalesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditSalesHeaders) SetXAcsDingtalkAccessToken(v string) *EditSalesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditSalesRequest struct {
	// 编辑数据
	Data *EditSalesRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 数据类型，固定填写158
	Datatype *int64 `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id，不填或者填0为新增数据
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 时间戳
	Stamp *int64 `json:"stamp,omitempty" xml:"stamp,omitempty"`
}

func (s EditSalesRequest) String() string {
	return tea.Prettify(s)
}

func (s EditSalesRequest) GoString() string {
	return s.String()
}

func (s *EditSalesRequest) SetData(v *EditSalesRequestData) *EditSalesRequest {
	s.Data = v
	return s
}

func (s *EditSalesRequest) SetDatatype(v int64) *EditSalesRequest {
	s.Datatype = &v
	return s
}

func (s *EditSalesRequest) SetMsgid(v int64) *EditSalesRequest {
	s.Msgid = &v
	return s
}

func (s *EditSalesRequest) SetStamp(v int64) *EditSalesRequest {
	s.Stamp = &v
	return s
}

type EditSalesRequestData struct {
	// 创建人
	DataUserid *string `json:"data_userid,omitempty" xml:"data_userid,omitempty"`
	// 对应客户
	XshCustomerid *string `json:"xsh_customerid,omitempty" xml:"xsh_customerid,omitempty"`
	// 发现时间
	XshDate *string `json:"xsh_date,omitempty" xml:"xsh_date,omitempty"`
	// 预计签单日
	XshExpdate *string `json:"xsh_expdate,omitempty" xml:"xsh_expdate,omitempty"`
	// 预期金额
	XshExpmoney *string `json:"xsh_expmoney,omitempty" xml:"xsh_expmoney,omitempty"`
	// 来源
	XshFrom *string `json:"xsh_from,omitempty" xml:"xsh_from,omitempty"`
	// 可能性
	XshKnx *string `json:"xsh_knx,omitempty" xml:"xsh_knx,omitempty"`
	// 联系方式
	XshLianxi *string `json:"xsh_lianxi,omitempty" xml:"xsh_lianxi,omitempty"`
	// 联系人
	XshLxrid *string `json:"xsh_lxrid,omitempty" xml:"xsh_lxrid,omitempty"`
	// 外币备注
	XshMoneynote *string `json:"xsh_moneynote,omitempty" xml:"xsh_moneynote,omitempty"`
	// 机会编号
	XshNumber *string `json:"xsh_number,omitempty" xml:"xsh_number,omitempty"`
	// 阶段
	XshPhase *string `json:"xsh_phase,omitempty" xml:"xsh_phase,omitempty"`
	// 阶段备注
	XshPhasenote *string `json:"xsh_phasenote,omitempty" xml:"xsh_phasenote,omitempty"`
	// 所有者
	XshPreside *string `json:"xsh_preside,omitempty" xml:"xsh_preside,omitempty"`
	// 提供人
	XshProvider *string `json:"xsh_provider,omitempty" xml:"xsh_provider,omitempty"`
	// 客户需求
	XshRequire *string `json:"xsh_require,omitempty" xml:"xsh_require,omitempty"`
	// 状态
	XshState *string `json:"xsh_state,omitempty" xml:"xsh_state,omitempty"`
	// 主题
	XshTitle *string `json:"xsh_title,omitempty" xml:"xsh_title,omitempty"`
	// 类型
	XshType *string `json:"xsh_type,omitempty" xml:"xsh_type,omitempty"`
}

func (s EditSalesRequestData) String() string {
	return tea.Prettify(s)
}

func (s EditSalesRequestData) GoString() string {
	return s.String()
}

func (s *EditSalesRequestData) SetDataUserid(v string) *EditSalesRequestData {
	s.DataUserid = &v
	return s
}

func (s *EditSalesRequestData) SetXshCustomerid(v string) *EditSalesRequestData {
	s.XshCustomerid = &v
	return s
}

func (s *EditSalesRequestData) SetXshDate(v string) *EditSalesRequestData {
	s.XshDate = &v
	return s
}

func (s *EditSalesRequestData) SetXshExpdate(v string) *EditSalesRequestData {
	s.XshExpdate = &v
	return s
}

func (s *EditSalesRequestData) SetXshExpmoney(v string) *EditSalesRequestData {
	s.XshExpmoney = &v
	return s
}

func (s *EditSalesRequestData) SetXshFrom(v string) *EditSalesRequestData {
	s.XshFrom = &v
	return s
}

func (s *EditSalesRequestData) SetXshKnx(v string) *EditSalesRequestData {
	s.XshKnx = &v
	return s
}

func (s *EditSalesRequestData) SetXshLianxi(v string) *EditSalesRequestData {
	s.XshLianxi = &v
	return s
}

func (s *EditSalesRequestData) SetXshLxrid(v string) *EditSalesRequestData {
	s.XshLxrid = &v
	return s
}

func (s *EditSalesRequestData) SetXshMoneynote(v string) *EditSalesRequestData {
	s.XshMoneynote = &v
	return s
}

func (s *EditSalesRequestData) SetXshNumber(v string) *EditSalesRequestData {
	s.XshNumber = &v
	return s
}

func (s *EditSalesRequestData) SetXshPhase(v string) *EditSalesRequestData {
	s.XshPhase = &v
	return s
}

func (s *EditSalesRequestData) SetXshPhasenote(v string) *EditSalesRequestData {
	s.XshPhasenote = &v
	return s
}

func (s *EditSalesRequestData) SetXshPreside(v string) *EditSalesRequestData {
	s.XshPreside = &v
	return s
}

func (s *EditSalesRequestData) SetXshProvider(v string) *EditSalesRequestData {
	s.XshProvider = &v
	return s
}

func (s *EditSalesRequestData) SetXshRequire(v string) *EditSalesRequestData {
	s.XshRequire = &v
	return s
}

func (s *EditSalesRequestData) SetXshState(v string) *EditSalesRequestData {
	s.XshState = &v
	return s
}

func (s *EditSalesRequestData) SetXshTitle(v string) *EditSalesRequestData {
	s.XshTitle = &v
	return s
}

func (s *EditSalesRequestData) SetXshType(v string) *EditSalesRequestData {
	s.XshType = &v
	return s
}

type EditSalesResponseBody struct {
	// 编辑数据的id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EditSalesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditSalesResponseBody) GoString() string {
	return s.String()
}

func (s *EditSalesResponseBody) SetMsgid(v int64) *EditSalesResponseBody {
	s.Msgid = &v
	return s
}

func (s *EditSalesResponseBody) SetTime(v string) *EditSalesResponseBody {
	s.Time = &v
	return s
}

type EditSalesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditSalesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditSalesResponse) String() string {
	return tea.Prettify(s)
}

func (s EditSalesResponse) GoString() string {
	return s.String()
}

func (s *EditSalesResponse) SetHeaders(v map[string]*string) *EditSalesResponse {
	s.Headers = v
	return s
}

func (s *EditSalesResponse) SetBody(v *EditSalesResponseBody) *EditSalesResponse {
	s.Body = v
	return s
}

type GetDataListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDataListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDataListHeaders) GoString() string {
	return s.String()
}

func (s *GetDataListHeaders) SetCommonHeaders(v map[string]*string) *GetDataListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDataListHeaders) SetXAcsDingtalkAccessToken(v string) *GetDataListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDataListRequest struct {
	// 数据类型，参考数据类型ID对照表
	Datatype *string `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 页码
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 分页条数
	Pagesize *int64 `json:"pagesize,omitempty" xml:"pagesize,omitempty"`
}

func (s GetDataListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataListRequest) GoString() string {
	return s.String()
}

func (s *GetDataListRequest) SetDatatype(v string) *GetDataListRequest {
	s.Datatype = &v
	return s
}

func (s *GetDataListRequest) SetPage(v int64) *GetDataListRequest {
	s.Page = &v
	return s
}

func (s *GetDataListRequest) SetPagesize(v int64) *GetDataListRequest {
	s.Pagesize = &v
	return s
}

type GetDataListResponseBody struct {
	// 数据列表
	Data []*GetDataListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 字段明细
	Dataname map[string]*string `json:"dataname,omitempty" xml:"dataname,omitempty"`
	// 当前页码
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 分页条数
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// 总条数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetDataListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataListResponseBody) SetData(v []*GetDataListResponseBodyData) *GetDataListResponseBody {
	s.Data = v
	return s
}

func (s *GetDataListResponseBody) SetDataname(v map[string]*string) *GetDataListResponseBody {
	s.Dataname = v
	return s
}

func (s *GetDataListResponseBody) SetPage(v int64) *GetDataListResponseBody {
	s.Page = &v
	return s
}

func (s *GetDataListResponseBody) SetPageSize(v int64) *GetDataListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetDataListResponseBody) SetTime(v string) *GetDataListResponseBody {
	s.Time = &v
	return s
}

func (s *GetDataListResponseBody) SetTotalCount(v int64) *GetDataListResponseBody {
	s.TotalCount = &v
	return s
}

type GetDataListResponseBodyData struct {
	Detail map[string]*string `json:"detail,omitempty" xml:"detail,omitempty"`
}

func (s GetDataListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDataListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataListResponseBodyData) SetDetail(v map[string]*string) *GetDataListResponseBodyData {
	s.Detail = v
	return s
}

type GetDataListResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDataListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataListResponse) GoString() string {
	return s.String()
}

func (s *GetDataListResponse) SetHeaders(v map[string]*string) *GetDataListResponse {
	s.Headers = v
	return s
}

func (s *GetDataListResponse) SetBody(v *GetDataListResponseBody) *GetDataListResponse {
	s.Body = v
	return s
}

type GetDataViewHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDataViewHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDataViewHeaders) GoString() string {
	return s.String()
}

func (s *GetDataViewHeaders) SetCommonHeaders(v map[string]*string) *GetDataViewHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDataViewHeaders) SetXAcsDingtalkAccessToken(v string) *GetDataViewHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDataViewRequest struct {
	// 数据类型，参考数据类型ID对照表
	Datatype *string `json:"datatype,omitempty" xml:"datatype,omitempty"`
	// 数据id
	Msgid *int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
}

func (s GetDataViewRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataViewRequest) GoString() string {
	return s.String()
}

func (s *GetDataViewRequest) SetDatatype(v string) *GetDataViewRequest {
	s.Datatype = &v
	return s
}

func (s *GetDataViewRequest) SetMsgid(v int64) *GetDataViewRequest {
	s.Msgid = &v
	return s
}

type GetDataViewResponseBody struct {
	Data *GetDataViewResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 字段明细
	Dataname map[string]map[string]interface{} `json:"dataname,omitempty" xml:"dataname,omitempty"`
	// 响应时间
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDataViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataViewResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataViewResponseBody) SetData(v *GetDataViewResponseBodyData) *GetDataViewResponseBody {
	s.Data = v
	return s
}

func (s *GetDataViewResponseBody) SetDataname(v map[string]map[string]interface{}) *GetDataViewResponseBody {
	s.Dataname = v
	return s
}

func (s *GetDataViewResponseBody) SetTime(v string) *GetDataViewResponseBody {
	s.Time = &v
	return s
}

type GetDataViewResponseBodyData struct {
	// 数据详情
	Detail map[string]*string `json:"detail,omitempty" xml:"detail,omitempty"`
}

func (s GetDataViewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDataViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataViewResponseBodyData) SetDetail(v map[string]*string) *GetDataViewResponseBodyData {
	s.Detail = v
	return s
}

type GetDataViewResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDataViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataViewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataViewResponse) GoString() string {
	return s.String()
}

func (s *GetDataViewResponse) SetHeaders(v map[string]*string) *GetDataViewResponse {
	s.Headers = v
	return s
}

func (s *GetDataViewResponse) SetBody(v *GetDataViewResponseBody) *GetDataViewResponse {
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

func (client *Client) EditContact(request *EditContactRequest) (_result *EditContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditContactHeaders{}
	_result = &EditContactResponse{}
	_body, _err := client.EditContactWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditContactWithOptions(request *EditContactRequest, headers *EditContactHeaders, runtime *util.RuntimeOptions) (_result *EditContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditContactResponse{}
	_body, _err := client.DoROARequest(tea.String("EditContact"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/contacts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditCustomer(request *EditCustomerRequest) (_result *EditCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditCustomerHeaders{}
	_result = &EditCustomerResponse{}
	_body, _err := client.EditCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditCustomerWithOptions(request *EditCustomerRequest, headers *EditCustomerHeaders, runtime *util.RuntimeOptions) (_result *EditCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditCustomerResponse{}
	_body, _err := client.DoROARequest(tea.String("EditCustomer"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/customers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditCustomerPool(request *EditCustomerPoolRequest) (_result *EditCustomerPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditCustomerPoolHeaders{}
	_result = &EditCustomerPoolResponse{}
	_body, _err := client.EditCustomerPoolWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditCustomerPoolWithOptions(request *EditCustomerPoolRequest, headers *EditCustomerPoolHeaders, runtime *util.RuntimeOptions) (_result *EditCustomerPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditCustomerPoolResponse{}
	_body, _err := client.DoROARequest(tea.String("EditCustomerPool"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/customerPools"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditExchange(request *EditExchangeRequest) (_result *EditExchangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditExchangeHeaders{}
	_result = &EditExchangeResponse{}
	_body, _err := client.EditExchangeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditExchangeWithOptions(request *EditExchangeRequest, headers *EditExchangeHeaders, runtime *util.RuntimeOptions) (_result *EditExchangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditExchangeResponse{}
	_body, _err := client.DoROARequest(tea.String("EditExchange"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/exchanges"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditGoods(request *EditGoodsRequest) (_result *EditGoodsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditGoodsHeaders{}
	_result = &EditGoodsResponse{}
	_body, _err := client.EditGoodsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditGoodsWithOptions(request *EditGoodsRequest, headers *EditGoodsHeaders, runtime *util.RuntimeOptions) (_result *EditGoodsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditGoodsResponse{}
	_body, _err := client.DoROARequest(tea.String("EditGoods"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/goods"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditIntostock(request *EditIntostockRequest) (_result *EditIntostockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditIntostockHeaders{}
	_result = &EditIntostockResponse{}
	_body, _err := client.EditIntostockWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditIntostockWithOptions(request *EditIntostockRequest, headers *EditIntostockHeaders, runtime *util.RuntimeOptions) (_result *EditIntostockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditIntostockResponse{}
	_body, _err := client.DoROARequest(tea.String("EditIntostock"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/intostocks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditInvoice(request *EditInvoiceRequest) (_result *EditInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditInvoiceHeaders{}
	_result = &EditInvoiceResponse{}
	_body, _err := client.EditInvoiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditInvoiceWithOptions(request *EditInvoiceRequest, headers *EditInvoiceHeaders, runtime *util.RuntimeOptions) (_result *EditInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditInvoiceResponse{}
	_body, _err := client.DoROARequest(tea.String("EditInvoice"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/invoices"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditOrder(request *EditOrderRequest) (_result *EditOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditOrderHeaders{}
	_result = &EditOrderResponse{}
	_body, _err := client.EditOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditOrderWithOptions(request *EditOrderRequest, headers *EditOrderHeaders, runtime *util.RuntimeOptions) (_result *EditOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("EditOrder"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/orders"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditOutstock(request *EditOutstockRequest) (_result *EditOutstockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditOutstockHeaders{}
	_result = &EditOutstockResponse{}
	_body, _err := client.EditOutstockWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditOutstockWithOptions(request *EditOutstockRequest, headers *EditOutstockHeaders, runtime *util.RuntimeOptions) (_result *EditOutstockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditOutstockResponse{}
	_body, _err := client.DoROARequest(tea.String("EditOutstock"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/outstocks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditProduction(request *EditProductionRequest) (_result *EditProductionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditProductionHeaders{}
	_result = &EditProductionResponse{}
	_body, _err := client.EditProductionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditProductionWithOptions(request *EditProductionRequest, headers *EditProductionHeaders, runtime *util.RuntimeOptions) (_result *EditProductionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditProductionResponse{}
	_body, _err := client.DoROARequest(tea.String("EditProduction"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/productions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditPurchase(request *EditPurchaseRequest) (_result *EditPurchaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditPurchaseHeaders{}
	_result = &EditPurchaseResponse{}
	_body, _err := client.EditPurchaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditPurchaseWithOptions(request *EditPurchaseRequest, headers *EditPurchaseHeaders, runtime *util.RuntimeOptions) (_result *EditPurchaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditPurchaseResponse{}
	_body, _err := client.DoROARequest(tea.String("EditPurchase"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/purchases"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditQuotationRecord(request *EditQuotationRecordRequest) (_result *EditQuotationRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditQuotationRecordHeaders{}
	_result = &EditQuotationRecordResponse{}
	_body, _err := client.EditQuotationRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditQuotationRecordWithOptions(request *EditQuotationRecordRequest, headers *EditQuotationRecordHeaders, runtime *util.RuntimeOptions) (_result *EditQuotationRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditQuotationRecordResponse{}
	_body, _err := client.DoROARequest(tea.String("EditQuotationRecord"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/quotationRecords"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditSales(request *EditSalesRequest) (_result *EditSalesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditSalesHeaders{}
	_result = &EditSalesResponse{}
	_body, _err := client.EditSalesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditSalesWithOptions(request *EditSalesRequest, headers *EditSalesHeaders, runtime *util.RuntimeOptions) (_result *EditSalesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Data))) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		body["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		body["msgid"] = request.Msgid
	}

	if !tea.BoolValue(util.IsUnset(request.Stamp)) {
		body["stamp"] = request.Stamp
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
	_result = &EditSalesResponse{}
	_body, _err := client.DoROARequest(tea.String("EditSales"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/jzcrm/sales"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataList(request *GetDataListRequest) (_result *GetDataListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDataListHeaders{}
	_result = &GetDataListResponse{}
	_body, _err := client.GetDataListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataListWithOptions(request *GetDataListRequest, headers *GetDataListHeaders, runtime *util.RuntimeOptions) (_result *GetDataListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		query["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Pagesize)) {
		query["pagesize"] = request.Pagesize
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
	_result = &GetDataListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDataList"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/jzcrm/data"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataView(request *GetDataViewRequest) (_result *GetDataViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDataViewHeaders{}
	_result = &GetDataViewResponse{}
	_body, _err := client.GetDataViewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataViewWithOptions(request *GetDataViewRequest, headers *GetDataViewHeaders, runtime *util.RuntimeOptions) (_result *GetDataViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Datatype)) {
		query["datatype"] = request.Datatype
	}

	if !tea.BoolValue(util.IsUnset(request.Msgid)) {
		query["msgid"] = request.Msgid
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
	_result = &GetDataViewResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDataView"), tea.String("jzcrm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/jzcrm/dataView"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
