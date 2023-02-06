// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package org_culture_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AssignOrgHoldingToEmpHoldingBatchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AssignOrgHoldingToEmpHoldingBatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchHeaders) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchHeaders) SetCommonHeaders(v map[string]*string) *AssignOrgHoldingToEmpHoldingBatchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchHeaders) SetXAcsDingtalkAccessToken(v string) *AssignOrgHoldingToEmpHoldingBatchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchRequest struct {
	// 备注信息 长度小于40
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 是否发送组织文化通知
	SendOrgCultureInform *bool `json:"sendOrgCultureInform,omitempty" xml:"sendOrgCultureInform,omitempty"`
	// 发放积分或额度数量 1～100000
	SingleAmount *int64 `json:"singleAmount,omitempty" xml:"singleAmount,omitempty"`
	// 发放人sourceUsage  发放人与接受人usage应一一对应
	// 发放积分sourceUsage：OPEN_ORG_POINT_PERSONAL_ASSIGN 对应的targetUsage为OPEN_EMP_POINT_PERSONAL_RECEIVE；
	// 发额度sourceUsage：OPEN_ORG_POINT_HOLDING_ASSIGN 对应的 targetUsage为OPEN_EMP_POINT_HOLDING_RECEIVE；
	// 行为规则发积分 sourceUsage：OPEN_ACTION_RULE_PERSONAL_ASSIGN 对应的 targetUsage为OPEN_ACTION_RULE_PERSONAL_RECEIVE
	SourceUsage *string `json:"sourceUsage,omitempty" xml:"sourceUsage,omitempty"`
	// 接受人targetUsage  发放人与接受人usage应一一对应
	TargetUsage *string `json:"targetUsage,omitempty" xml:"targetUsage,omitempty"`
	// 发放目标用户
	TargetUserList []*AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList `json:"targetUserList,omitempty" xml:"targetUserList,omitempty" type:"Repeated"`
}

func (s AssignOrgHoldingToEmpHoldingBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchRequest) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetRemark(v string) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.Remark = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetSendOrgCultureInform(v bool) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.SendOrgCultureInform = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetSingleAmount(v int64) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.SingleAmount = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetSourceUsage(v string) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.SourceUsage = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetTargetUsage(v string) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.TargetUsage = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequest) SetTargetUserList(v []*AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList) *AssignOrgHoldingToEmpHoldingBatchRequest {
	s.TargetUserList = v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList struct {
	// 积分交易单号，长度1-32。
	//
	OutId *string `json:"outId,omitempty" xml:"outId,omitempty"`
	// 操作目标对象userId
	TargetUserId *string `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
}

func (s AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList) SetOutId(v string) *AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList {
	s.OutId = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList) SetTargetUserId(v string) *AssignOrgHoldingToEmpHoldingBatchRequestTargetUserList {
	s.TargetUserId = &v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchResponseBody struct {
	Result  *AssignOrgHoldingToEmpHoldingBatchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBody) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBody) SetResult(v *AssignOrgHoldingToEmpHoldingBatchResponseBodyResult) *AssignOrgHoldingToEmpHoldingBatchResponseBody {
	s.Result = v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBody) SetSuccess(v bool) *AssignOrgHoldingToEmpHoldingBatchResponseBody {
	s.Success = &v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchResponseBodyResult struct {
	// 每个人发放的结果
	OpenPointInvokeResultDTOS []*AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS `json:"openPointInvokeResultDTOS,omitempty" xml:"openPointInvokeResultDTOS,omitempty" type:"Repeated"`
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResult) SetOpenPointInvokeResultDTOS(v []*AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResult {
	s.OpenPointInvokeResultDTOS = v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS struct {
	// 错误码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 状态SUCCESS：成功。 FAIL：失败 UNKNOWN:结果未知
	InvokeStatus *string `json:"invokeStatus,omitempty" xml:"invokeStatus,omitempty"`
	// 错误信息
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 积分交易单号
	//
	OutId *string `json:"outId,omitempty" xml:"outId,omitempty"`
	// 发放用户userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) SetCode(v string) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.Code = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) SetInvokeStatus(v string) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.InvokeStatus = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) SetMsg(v string) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.Msg = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) SetOutId(v string) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.OutId = &v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS) SetUserId(v string) *AssignOrgHoldingToEmpHoldingBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.UserId = &v
	return s
}

type AssignOrgHoldingToEmpHoldingBatchResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssignOrgHoldingToEmpHoldingBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssignOrgHoldingToEmpHoldingBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignOrgHoldingToEmpHoldingBatchResponse) GoString() string {
	return s.String()
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponse) SetHeaders(v map[string]*string) *AssignOrgHoldingToEmpHoldingBatchResponse {
	s.Headers = v
	return s
}

func (s *AssignOrgHoldingToEmpHoldingBatchResponse) SetBody(v *AssignOrgHoldingToEmpHoldingBatchResponseBody) *AssignOrgHoldingToEmpHoldingBatchResponse {
	s.Body = v
	return s
}

type ConsumeUserPointsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConsumeUserPointsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsHeaders) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsHeaders) SetCommonHeaders(v map[string]*string) *ConsumeUserPointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConsumeUserPointsHeaders) SetXAcsDingtalkAccessToken(v string) *ConsumeUserPointsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConsumeUserPointsRequest struct {
	// 扣减积分数量，1～1000000
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 幂等外部ID，最大长度32个字符
	OutId *string `json:"outId,omitempty" xml:"outId,omitempty"`
	// 备注，最长32个字符
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 用途，可用值：OPEN_EMP_POINT_CONSUME_DEFAULT-默认扣减，OPEN_EMP_POINT_PUNISH_CONSUME-惩罚扣减；默认为: OPEN_EMP_POINT_CONSUME_DEFAULT
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ConsumeUserPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsRequest) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsRequest) SetAmount(v int64) *ConsumeUserPointsRequest {
	s.Amount = &v
	return s
}

func (s *ConsumeUserPointsRequest) SetOutId(v string) *ConsumeUserPointsRequest {
	s.OutId = &v
	return s
}

func (s *ConsumeUserPointsRequest) SetRemark(v string) *ConsumeUserPointsRequest {
	s.Remark = &v
	return s
}

func (s *ConsumeUserPointsRequest) SetUsage(v string) *ConsumeUserPointsRequest {
	s.Usage = &v
	return s
}

type ConsumeUserPointsResponseBody struct {
	// 响应数据
	Result *ConsumeUserPointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 请求响应状态
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConsumeUserPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsResponseBody) SetResult(v *ConsumeUserPointsResponseBodyResult) *ConsumeUserPointsResponseBody {
	s.Result = v
	return s
}

func (s *ConsumeUserPointsResponseBody) SetSuccess(v bool) *ConsumeUserPointsResponseBody {
	s.Success = &v
	return s
}

type ConsumeUserPointsResponseBodyResult struct {
	// 扣减后剩余积分数量
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s ConsumeUserPointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsResponseBodyResult) SetAmount(v int64) *ConsumeUserPointsResponseBodyResult {
	s.Amount = &v
	return s
}

type ConsumeUserPointsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConsumeUserPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConsumeUserPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsResponse) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsResponse) SetHeaders(v map[string]*string) *ConsumeUserPointsResponse {
	s.Headers = v
	return s
}

func (s *ConsumeUserPointsResponse) SetBody(v *ConsumeUserPointsResponseBody) *ConsumeUserPointsResponse {
	s.Body = v
	return s
}

type CreateOrgHonorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateOrgHonorHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgHonorHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorHeaders) SetCommonHeaders(v map[string]*string) *CreateOrgHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrgHonorHeaders) SetXAcsDingtalkAccessToken(v string) *CreateOrgHonorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateOrgHonorRequest struct {
	// 头像挂件   图片尺寸 240*240，不超过1M，支持PNG。图片请使用钉钉媒体资源标识符media_id，参考文档：https://open.dingtalk.com/document/isvapp-server/upload-media-files
	AvatarFrameMediaId *string `json:"avatarFrameMediaId,omitempty" xml:"avatarFrameMediaId,omitempty"`
	// 背景颜色，如下可选：#FFFBB4 #FFE7BC #FFDAF4 #DAF6A8 #E4D7FF #BFDFFF #B9F2D6
	DefaultBgColor *string `json:"defaultBgColor,omitempty" xml:"defaultBgColor,omitempty"`
	// 描述 长度30字符 不支持表情图标等
	MedalDesc *string `json:"medalDesc,omitempty" xml:"medalDesc,omitempty"`
	// 荣誉图片  图片尺寸 900*900，不超过1M，支持PNG 。图片请使用钉钉媒体资源标识符media_id，参考文档：https://open.dingtalk.com/document/isvapp-server/upload-media-files
	MedalMediaId *string `json:"medalMediaId,omitempty" xml:"medalMediaId,omitempty"`
	// 组织的勋章名称 长度10字符 不支持表情图标等
	MedalName *string `json:"medalName,omitempty" xml:"medalName,omitempty"`
	// 创建荣誉勋章模板人在组织内的userid，需要主/子管理员角色
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateOrgHonorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgHonorRequest) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorRequest) SetAvatarFrameMediaId(v string) *CreateOrgHonorRequest {
	s.AvatarFrameMediaId = &v
	return s
}

func (s *CreateOrgHonorRequest) SetDefaultBgColor(v string) *CreateOrgHonorRequest {
	s.DefaultBgColor = &v
	return s
}

func (s *CreateOrgHonorRequest) SetMedalDesc(v string) *CreateOrgHonorRequest {
	s.MedalDesc = &v
	return s
}

func (s *CreateOrgHonorRequest) SetMedalMediaId(v string) *CreateOrgHonorRequest {
	s.MedalMediaId = &v
	return s
}

func (s *CreateOrgHonorRequest) SetMedalName(v string) *CreateOrgHonorRequest {
	s.MedalName = &v
	return s
}

func (s *CreateOrgHonorRequest) SetUserId(v string) *CreateOrgHonorRequest {
	s.UserId = &v
	return s
}

type CreateOrgHonorResponseBody struct {
	Result  *CreateOrgHonorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateOrgHonorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgHonorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorResponseBody) SetResult(v *CreateOrgHonorResponseBodyResult) *CreateOrgHonorResponseBody {
	s.Result = v
	return s
}

func (s *CreateOrgHonorResponseBody) SetSuccess(v bool) *CreateOrgHonorResponseBody {
	s.Success = &v
	return s
}

type CreateOrgHonorResponseBodyResult struct {
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
}

func (s CreateOrgHonorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgHonorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorResponseBodyResult) SetHonorId(v string) *CreateOrgHonorResponseBodyResult {
	s.HonorId = &v
	return s
}

type CreateOrgHonorResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrgHonorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrgHonorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgHonorResponse) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorResponse) SetHeaders(v map[string]*string) *CreateOrgHonorResponse {
	s.Headers = v
	return s
}

func (s *CreateOrgHonorResponse) SetBody(v *CreateOrgHonorResponseBody) *CreateOrgHonorResponse {
	s.Body = v
	return s
}

type DeductionPointBatchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeductionPointBatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchHeaders) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchHeaders) SetCommonHeaders(v map[string]*string) *DeductionPointBatchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeductionPointBatchHeaders) SetXAcsDingtalkAccessToken(v string) *DeductionPointBatchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeductionPointBatchRequest struct {
	// 扣减数量 范围：1—100000
	DeductionAmount *int64 `json:"deductionAmount,omitempty" xml:"deductionAmount,omitempty"`
	// 扣减积分原因
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 是否发送组织文化通知
	SendOrgCultureInform *bool `json:"sendOrgCultureInform,omitempty" xml:"sendOrgCultureInform,omitempty"`
	// 批量扣减积分用户
	TargetUserList []*DeductionPointBatchRequestTargetUserList `json:"targetUserList,omitempty" xml:"targetUserList,omitempty" type:"Repeated"`
	// 操作人userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeductionPointBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchRequest) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchRequest) SetDeductionAmount(v int64) *DeductionPointBatchRequest {
	s.DeductionAmount = &v
	return s
}

func (s *DeductionPointBatchRequest) SetRemark(v string) *DeductionPointBatchRequest {
	s.Remark = &v
	return s
}

func (s *DeductionPointBatchRequest) SetSendOrgCultureInform(v bool) *DeductionPointBatchRequest {
	s.SendOrgCultureInform = &v
	return s
}

func (s *DeductionPointBatchRequest) SetTargetUserList(v []*DeductionPointBatchRequestTargetUserList) *DeductionPointBatchRequest {
	s.TargetUserList = v
	return s
}

func (s *DeductionPointBatchRequest) SetUserId(v string) *DeductionPointBatchRequest {
	s.UserId = &v
	return s
}

type DeductionPointBatchRequestTargetUserList struct {
	// 积分交易单号
	OutId *string `json:"outId,omitempty" xml:"outId,omitempty"`
	// 扣减目标用户userId
	TargetUserId *string `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
}

func (s DeductionPointBatchRequestTargetUserList) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchRequestTargetUserList) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchRequestTargetUserList) SetOutId(v string) *DeductionPointBatchRequestTargetUserList {
	s.OutId = &v
	return s
}

func (s *DeductionPointBatchRequestTargetUserList) SetTargetUserId(v string) *DeductionPointBatchRequestTargetUserList {
	s.TargetUserId = &v
	return s
}

type DeductionPointBatchResponseBody struct {
	Result *DeductionPointBatchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 调用是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeductionPointBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchResponseBody) SetResult(v *DeductionPointBatchResponseBodyResult) *DeductionPointBatchResponseBody {
	s.Result = v
	return s
}

func (s *DeductionPointBatchResponseBody) SetSuccess(v bool) *DeductionPointBatchResponseBody {
	s.Success = &v
	return s
}

type DeductionPointBatchResponseBodyResult struct {
	// 每个人发放的结果
	OpenPointInvokeResultDTOS []*DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS `json:"openPointInvokeResultDTOS,omitempty" xml:"openPointInvokeResultDTOS,omitempty" type:"Repeated"`
}

func (s DeductionPointBatchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchResponseBodyResult) SetOpenPointInvokeResultDTOS(v []*DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) *DeductionPointBatchResponseBodyResult {
	s.OpenPointInvokeResultDTOS = v
	return s
}

type DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS struct {
	// 错误码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 状态 success：成功。 Fail：失败 UNKNOWN:结果未知
	InvokeStatus *string `json:"invokeStatus,omitempty" xml:"invokeStatus,omitempty"`
	// 错误信息
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 积分交易单号
	OutId *string `json:"outId,omitempty" xml:"outId,omitempty"`
	// 扣减用户userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) SetCode(v string) *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.Code = &v
	return s
}

func (s *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) SetInvokeStatus(v string) *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.InvokeStatus = &v
	return s
}

func (s *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) SetMsg(v string) *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.Msg = &v
	return s
}

func (s *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) SetOutId(v string) *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.OutId = &v
	return s
}

func (s *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS) SetUserId(v string) *DeductionPointBatchResponseBodyResultOpenPointInvokeResultDTOS {
	s.UserId = &v
	return s
}

type DeductionPointBatchResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeductionPointBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeductionPointBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeductionPointBatchResponse) GoString() string {
	return s.String()
}

func (s *DeductionPointBatchResponse) SetHeaders(v map[string]*string) *DeductionPointBatchResponse {
	s.Headers = v
	return s
}

func (s *DeductionPointBatchResponse) SetBody(v *DeductionPointBatchResponseBody) *DeductionPointBatchResponse {
	s.Body = v
	return s
}

type ExportPointOpenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExportPointOpenHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExportPointOpenHeaders) GoString() string {
	return s.String()
}

func (s *ExportPointOpenHeaders) SetCommonHeaders(v map[string]*string) *ExportPointOpenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExportPointOpenHeaders) SetXAcsDingtalkAccessToken(v string) *ExportPointOpenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExportPointOpenRequest struct {
	// exportType为1时不需要传此参数，目前仅exportType=3时必须传入此参数,必须为七日内某一天且不能选择当日，格式yyyyMmdd。
	ExportDate *string `json:"exportDate,omitempty" xml:"exportDate,omitempty"`
	// 导出类型 1为七日内明细，3为七日内某一天榜单，且都不包含当日
	ExportType *int64 `json:"exportType,omitempty" xml:"exportType,omitempty"`
	// 操作人userId 必须为管理员
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExportPointOpenRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportPointOpenRequest) GoString() string {
	return s.String()
}

func (s *ExportPointOpenRequest) SetExportDate(v string) *ExportPointOpenRequest {
	s.ExportDate = &v
	return s
}

func (s *ExportPointOpenRequest) SetExportType(v int64) *ExportPointOpenRequest {
	s.ExportType = &v
	return s
}

func (s *ExportPointOpenRequest) SetUserId(v string) *ExportPointOpenRequest {
	s.UserId = &v
	return s
}

type ExportPointOpenResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExportPointOpenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportPointOpenResponseBody) GoString() string {
	return s.String()
}

func (s *ExportPointOpenResponseBody) SetResult(v bool) *ExportPointOpenResponseBody {
	s.Result = &v
	return s
}

func (s *ExportPointOpenResponseBody) SetSuccess(v bool) *ExportPointOpenResponseBody {
	s.Success = &v
	return s
}

type ExportPointOpenResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportPointOpenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportPointOpenResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportPointOpenResponse) GoString() string {
	return s.String()
}

func (s *ExportPointOpenResponse) SetHeaders(v map[string]*string) *ExportPointOpenResponse {
	s.Headers = v
	return s
}

func (s *ExportPointOpenResponse) SetBody(v *ExportPointOpenResponseBody) *ExportPointOpenResponse {
	s.Body = v
	return s
}

type GrantHonorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GrantHonorHeaders) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorHeaders) GoString() string {
	return s.String()
}

func (s *GrantHonorHeaders) SetCommonHeaders(v map[string]*string) *GrantHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GrantHonorHeaders) SetXAcsDingtalkAccessToken(v string) *GrantHonorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GrantHonorRequest struct {
	// 有效期到期时间 时间戳. 会处理成到期那天的23:59:59秒的时间戳
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// 颁奖词，最多可以填50字
	GrantReason *string `json:"grantReason,omitempty" xml:"grantReason,omitempty"`
	// 颁奖人名字，最多15个字
	GranterName *string `json:"granterName,omitempty" xml:"granterName,omitempty"`
	// 是否使用官宣号发送内网动态
	NoticeAnnouncer *bool `json:"noticeAnnouncer,omitempty" xml:"noticeAnnouncer,omitempty"`
	// 是否触达单聊会话通知
	NoticeSingle        *bool     `json:"noticeSingle,omitempty" xml:"noticeSingle,omitempty"`
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	// 接受人userId
	ReceiverUserIds []*string `json:"receiverUserIds,omitempty" xml:"receiverUserIds,omitempty" type:"Repeated"`
	// 发送人userId
	SenderUserId *string `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
}

func (s GrantHonorRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorRequest) GoString() string {
	return s.String()
}

func (s *GrantHonorRequest) SetExpirationTime(v int64) *GrantHonorRequest {
	s.ExpirationTime = &v
	return s
}

func (s *GrantHonorRequest) SetGrantReason(v string) *GrantHonorRequest {
	s.GrantReason = &v
	return s
}

func (s *GrantHonorRequest) SetGranterName(v string) *GrantHonorRequest {
	s.GranterName = &v
	return s
}

func (s *GrantHonorRequest) SetNoticeAnnouncer(v bool) *GrantHonorRequest {
	s.NoticeAnnouncer = &v
	return s
}

func (s *GrantHonorRequest) SetNoticeSingle(v bool) *GrantHonorRequest {
	s.NoticeSingle = &v
	return s
}

func (s *GrantHonorRequest) SetOpenConversationIds(v []*string) *GrantHonorRequest {
	s.OpenConversationIds = v
	return s
}

func (s *GrantHonorRequest) SetReceiverUserIds(v []*string) *GrantHonorRequest {
	s.ReceiverUserIds = v
	return s
}

func (s *GrantHonorRequest) SetSenderUserId(v string) *GrantHonorRequest {
	s.SenderUserId = &v
	return s
}

type GrantHonorResponseBody struct {
	Result  *GrantHonorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GrantHonorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorResponseBody) GoString() string {
	return s.String()
}

func (s *GrantHonorResponseBody) SetResult(v *GrantHonorResponseBodyResult) *GrantHonorResponseBody {
	s.Result = v
	return s
}

func (s *GrantHonorResponseBody) SetSuccess(v bool) *GrantHonorResponseBody {
	s.Success = &v
	return s
}

type GrantHonorResponseBodyResult struct {
	// 失败的userId
	FailedUserIds []*string `json:"failedUserIds,omitempty" xml:"failedUserIds,omitempty" type:"Repeated"`
	// 成功的userId
	SuccessUserIds []*string `json:"successUserIds,omitempty" xml:"successUserIds,omitempty" type:"Repeated"`
}

func (s GrantHonorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GrantHonorResponseBodyResult) SetFailedUserIds(v []*string) *GrantHonorResponseBodyResult {
	s.FailedUserIds = v
	return s
}

func (s *GrantHonorResponseBodyResult) SetSuccessUserIds(v []*string) *GrantHonorResponseBodyResult {
	s.SuccessUserIds = v
	return s
}

type GrantHonorResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantHonorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantHonorResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorResponse) GoString() string {
	return s.String()
}

func (s *GrantHonorResponse) SetHeaders(v map[string]*string) *GrantHonorResponse {
	s.Headers = v
	return s
}

func (s *GrantHonorResponse) SetBody(v *GrantHonorResponseBody) *GrantHonorResponse {
	s.Body = v
	return s
}

type QueryCorpPointsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCorpPointsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsHeaders) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsHeaders) SetCommonHeaders(v map[string]*string) *QueryCorpPointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCorpPointsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCorpPointsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCorpPointsRequest struct {
	// 操作用户ID
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
}

func (s QueryCorpPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsRequest) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsRequest) SetOptUserId(v string) *QueryCorpPointsRequest {
	s.OptUserId = &v
	return s
}

type QueryCorpPointsResponseBody struct {
	// 响应数据
	Result *QueryCorpPointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 请求响应状态
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryCorpPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsResponseBody) SetResult(v *QueryCorpPointsResponseBodyResult) *QueryCorpPointsResponseBody {
	s.Result = v
	return s
}

func (s *QueryCorpPointsResponseBody) SetSuccess(v bool) *QueryCorpPointsResponseBody {
	s.Success = &v
	return s
}

type QueryCorpPointsResponseBodyResult struct {
	// 企业员工可用于兑换积分总额
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s QueryCorpPointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsResponseBodyResult) SetAmount(v int64) *QueryCorpPointsResponseBodyResult {
	s.Amount = &v
	return s
}

type QueryCorpPointsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCorpPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCorpPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsResponse) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsResponse) SetHeaders(v map[string]*string) *QueryCorpPointsResponse {
	s.Headers = v
	return s
}

func (s *QueryCorpPointsResponse) SetBody(v *QueryCorpPointsResponseBody) *QueryCorpPointsResponse {
	s.Body = v
	return s
}

type QueryEmpPointDetailsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryEmpPointDetailsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsHeaders) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsHeaders) SetCommonHeaders(v map[string]*string) *QueryEmpPointDetailsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryEmpPointDetailsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryEmpPointDetailsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryEmpPointDetailsRequest struct {
	// 第几页 第一页是1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页大小最多50 默认值10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 查询目标对象userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryEmpPointDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsRequest) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsRequest) SetPageNumber(v int64) *QueryEmpPointDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryEmpPointDetailsRequest) SetPageSize(v int64) *QueryEmpPointDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEmpPointDetailsRequest) SetUserId(v string) *QueryEmpPointDetailsRequest {
	s.UserId = &v
	return s
}

type QueryEmpPointDetailsResponseBody struct {
	Result *QueryEmpPointDetailsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 调用是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryEmpPointDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBody) SetResult(v *QueryEmpPointDetailsResponseBodyResult) *QueryEmpPointDetailsResponseBody {
	s.Result = v
	return s
}

func (s *QueryEmpPointDetailsResponseBody) SetSuccess(v bool) *QueryEmpPointDetailsResponseBody {
	s.Success = &v
	return s
}

type QueryEmpPointDetailsResponseBodyResult struct {
	// 个人积分明细列表
	Details []*QueryEmpPointDetailsResponseBodyResultDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// 是否有下一页
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
}

func (s QueryEmpPointDetailsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBodyResult) SetDetails(v []*QueryEmpPointDetailsResponseBodyResultDetails) *QueryEmpPointDetailsResponseBodyResult {
	s.Details = v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResult) SetHasMore(v bool) *QueryEmpPointDetailsResponseBodyResult {
	s.HasMore = &v
	return s
}

type QueryEmpPointDetailsResponseBodyResultDetails struct {
	// 积分数量 发放时为负。 扣减时为正
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 积分交易单号
	//
	OutId                          *string                                                                      `json:"outId,omitempty" xml:"outId,omitempty"`
	PointOperateFeatureResponseDTO *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO `json:"pointOperateFeatureResponseDTO,omitempty" xml:"pointOperateFeatureResponseDTO,omitempty" type:"Struct"`
	// 源账户积分bizCode.
	// 个人可用积分:personal
	// 额度:credit
	SourceBizCode *string `json:"sourceBizCode,omitempty" xml:"sourceBizCode,omitempty"`
}

func (s QueryEmpPointDetailsResponseBodyResultDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBodyResultDetails) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBodyResultDetails) SetAmount(v int64) *QueryEmpPointDetailsResponseBodyResultDetails {
	s.Amount = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetails) SetGmtCreate(v int64) *QueryEmpPointDetailsResponseBodyResultDetails {
	s.GmtCreate = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetails) SetOutId(v string) *QueryEmpPointDetailsResponseBodyResultDetails {
	s.OutId = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetails) SetPointOperateFeatureResponseDTO(v *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) *QueryEmpPointDetailsResponseBodyResultDetails {
	s.PointOperateFeatureResponseDTO = v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetails) SetSourceBizCode(v string) *QueryEmpPointDetailsResponseBodyResultDetails {
	s.SourceBizCode = &v
	return s
}

type QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO struct {
	// 来源账户
	AccountSource *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource `json:"accountSource,omitempty" xml:"accountSource,omitempty" type:"Struct"`
	// 目标账户
	//
	AccountTarget *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget `json:"accountTarget,omitempty" xml:"accountTarget,omitempty" type:"Struct"`
	// 备注信息，在明细中展示
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 来源/用途，一般是系统固定的场景
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetAccountSource(v *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.AccountSource = v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetAccountTarget(v *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.AccountTarget = v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetRemark(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.Remark = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetUsage(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.Usage = &v
	return s
}

type QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource struct {
	// 积分账号的类型
	// 企业账号：ORG, 员工账号：EMP
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 企业内名字
	EmpName *string `json:"empName,omitempty" xml:"empName,omitempty"`
	// 用户userId
	//
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetAccountType(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.AccountType = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetEmpName(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.EmpName = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetUserId(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.UserId = &v
	return s
}

type QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget struct {
	// 积分账号的类型
	// 企业账号：ORG, 员工账号：EMP
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 企业内名字
	EmpName *string `json:"empName,omitempty" xml:"empName,omitempty"`
	// 用户useId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetAccountType(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.AccountType = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetEmpName(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.EmpName = &v
	return s
}

func (s *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetUserId(v string) *QueryEmpPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.UserId = &v
	return s
}

type QueryEmpPointDetailsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEmpPointDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEmpPointDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEmpPointDetailsResponse) GoString() string {
	return s.String()
}

func (s *QueryEmpPointDetailsResponse) SetHeaders(v map[string]*string) *QueryEmpPointDetailsResponse {
	s.Headers = v
	return s
}

func (s *QueryEmpPointDetailsResponse) SetBody(v *QueryEmpPointDetailsResponseBody) *QueryEmpPointDetailsResponse {
	s.Body = v
	return s
}

type QueryOrgHonorsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgHonorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgHonorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgHonorsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgHonorsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgHonorsRequest struct {
	// 分页获取数据时，数据的数量，默认为20，最大可传入100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页获取数据的标记，第一页调用时传0，非第一页传入上次调用本接口返回值中的nextToken
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryOrgHonorsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsRequest) SetMaxResults(v int32) *QueryOrgHonorsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryOrgHonorsRequest) SetNextToken(v string) *QueryOrgHonorsRequest {
	s.NextToken = &v
	return s
}

type QueryOrgHonorsResponseBody struct {
	Result  *QueryOrgHonorsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOrgHonorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBody) SetResult(v *QueryOrgHonorsResponseBodyResult) *QueryOrgHonorsResponseBody {
	s.Result = v
	return s
}

func (s *QueryOrgHonorsResponseBody) SetSuccess(v bool) *QueryOrgHonorsResponseBody {
	s.Success = &v
	return s
}

type QueryOrgHonorsResponseBodyResult struct {
	// 下次获取数据的游标
	NextToken  *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenHonors []*QueryOrgHonorsResponseBodyResultOpenHonors `json:"openHonors,omitempty" xml:"openHonors,omitempty" type:"Repeated"`
}

func (s QueryOrgHonorsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBodyResult) SetNextToken(v string) *QueryOrgHonorsResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResult) SetOpenHonors(v []*QueryOrgHonorsResponseBodyResultOpenHonors) *QueryOrgHonorsResponseBodyResult {
	s.OpenHonors = v
	return s
}

type QueryOrgHonorsResponseBodyResultOpenHonors struct {
	// 荣誉含义
	HonorDesc *string `json:"honorDesc,omitempty" xml:"honorDesc,omitempty"`
	// 荣誉id
	HonorId *int64 `json:"honorId,omitempty" xml:"honorId,omitempty"`
	// 荣誉图片url
	HonorImgUrl *string `json:"honorImgUrl,omitempty" xml:"honorImgUrl,omitempty"`
	// 荣誉名字
	HonorName *string `json:"honorName,omitempty" xml:"honorName,omitempty"`
	// 荣誉附赠的挂件图url
	HonorPendantImgUrl *string `json:"honorPendantImgUrl,omitempty" xml:"honorPendantImgUrl,omitempty"`
}

func (s QueryOrgHonorsResponseBodyResultOpenHonors) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponseBodyResultOpenHonors) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorDesc(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorDesc = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorId(v int64) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorId = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorImgUrl(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorImgUrl = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorName(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorName = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorPendantImgUrl(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorPendantImgUrl = &v
	return s
}

type QueryOrgHonorsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrgHonorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrgHonorsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponse) SetHeaders(v map[string]*string) *QueryOrgHonorsResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgHonorsResponse) SetBody(v *QueryOrgHonorsResponseBody) *QueryOrgHonorsResponse {
	s.Body = v
	return s
}

type QueryOrgPointDetailsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgPointDetailsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgPointDetailsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgPointDetailsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgPointDetailsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgPointDetailsRequest struct {
	// 查询企业账号明细，ORG,ORG_DEDUCTIONS两种。     ORG:企业账户明细 查询的是企业积分发放明细       ORG_DEDUCTIONS:扣除账户明细，查询的是企业扣减积分明细
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 查询页数 第一页是1 非空必传
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页大小最多50，默认10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 操作人userId 必须是管理员
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOrgPointDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsRequest) SetAccountType(v string) *QueryOrgPointDetailsRequest {
	s.AccountType = &v
	return s
}

func (s *QueryOrgPointDetailsRequest) SetPageNumber(v int64) *QueryOrgPointDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryOrgPointDetailsRequest) SetPageSize(v int64) *QueryOrgPointDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrgPointDetailsRequest) SetUserId(v string) *QueryOrgPointDetailsRequest {
	s.UserId = &v
	return s
}

type QueryOrgPointDetailsResponseBody struct {
	Result *QueryOrgPointDetailsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryOrgPointDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBody) SetResult(v *QueryOrgPointDetailsResponseBodyResult) *QueryOrgPointDetailsResponseBody {
	s.Result = v
	return s
}

type QueryOrgPointDetailsResponseBodyResult struct {
	// 积分明细列表
	Details []*QueryOrgPointDetailsResponseBodyResultDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// 分页使用，表示是否还有下一页
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 调用是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOrgPointDetailsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBodyResult) SetDetails(v []*QueryOrgPointDetailsResponseBodyResultDetails) *QueryOrgPointDetailsResponseBodyResult {
	s.Details = v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResult) SetHasMore(v bool) *QueryOrgPointDetailsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResult) SetSuccess(v bool) *QueryOrgPointDetailsResponseBodyResult {
	s.Success = &v
	return s
}

type QueryOrgPointDetailsResponseBodyResultDetails struct {
	// 积分数量 发放时为负。 扣减时为正
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 积分交易单号
	OutId *string `json:"outId,omitempty" xml:"outId,omitempty"`
	// 账户信息
	PointOperateFeatureResponseDTO *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO `json:"pointOperateFeatureResponseDTO,omitempty" xml:"pointOperateFeatureResponseDTO,omitempty" type:"Struct"`
	// 源账户积分bizCode。
	// 个人可用积分:personal
	// 额度:credit
	SourceBizCode *string `json:"sourceBizCode,omitempty" xml:"sourceBizCode,omitempty"`
}

func (s QueryOrgPointDetailsResponseBodyResultDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBodyResultDetails) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBodyResultDetails) SetAmount(v int64) *QueryOrgPointDetailsResponseBodyResultDetails {
	s.Amount = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetails) SetGmtCreate(v int64) *QueryOrgPointDetailsResponseBodyResultDetails {
	s.GmtCreate = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetails) SetOutId(v string) *QueryOrgPointDetailsResponseBodyResultDetails {
	s.OutId = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetails) SetPointOperateFeatureResponseDTO(v *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) *QueryOrgPointDetailsResponseBodyResultDetails {
	s.PointOperateFeatureResponseDTO = v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetails) SetSourceBizCode(v string) *QueryOrgPointDetailsResponseBodyResultDetails {
	s.SourceBizCode = &v
	return s
}

type QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO struct {
	// 如果是扣减操作明细，为被扣减账户
	// 如果是发放操作明细，为操作发放账户
	// 如果是个人积分明细，为来源账户
	AccountSource *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource `json:"accountSource,omitempty" xml:"accountSource,omitempty" type:"Struct"`
	// 如果是扣减操作明细，为操作扣减账户
	// 如果是发放操作明细，为被发放账户
	// 如果是个人积分明细，为目标账户
	AccountTarget *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget `json:"accountTarget,omitempty" xml:"accountTarget,omitempty" type:"Struct"`
	// 备注信息，在明细中展示
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 来源/用途 一般是系统固定的场景
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetAccountSource(v *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.AccountSource = v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetAccountTarget(v *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.AccountTarget = v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetRemark(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.Remark = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO) SetUsage(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTO {
	s.Usage = &v
	return s
}

type QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource struct {
	// 积分账号的类型
	// 企业账号：ORG, 员工账号：EMP
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 企业内名字
	EmpName *string `json:"empName,omitempty" xml:"empName,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetAccountType(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.AccountType = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetEmpName(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.EmpName = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource) SetUserId(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountSource {
	s.UserId = &v
	return s
}

type QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget struct {
	// 积分账号的类型
	// 企业账号：ORG, 员工账号：EMP
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// 企业内名字
	EmpName *string `json:"empName,omitempty" xml:"empName,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetAccountType(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.AccountType = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetEmpName(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.EmpName = &v
	return s
}

func (s *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget) SetUserId(v string) *QueryOrgPointDetailsResponseBodyResultDetailsPointOperateFeatureResponseDTOAccountTarget {
	s.UserId = &v
	return s
}

type QueryOrgPointDetailsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrgPointDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrgPointDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgPointDetailsResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgPointDetailsResponse) SetHeaders(v map[string]*string) *QueryOrgPointDetailsResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgPointDetailsResponse) SetBody(v *QueryOrgPointDetailsResponseBody) *QueryOrgPointDetailsResponse {
	s.Body = v
	return s
}

type QueryPointActionAutoAssignRuleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPointActionAutoAssignRuleHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPointActionAutoAssignRuleHeaders) GoString() string {
	return s.String()
}

func (s *QueryPointActionAutoAssignRuleHeaders) SetCommonHeaders(v map[string]*string) *QueryPointActionAutoAssignRuleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPointActionAutoAssignRuleHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPointActionAutoAssignRuleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPointActionAutoAssignRuleResponseBody struct {
	Result  *QueryPointActionAutoAssignRuleResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryPointActionAutoAssignRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPointActionAutoAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPointActionAutoAssignRuleResponseBody) SetResult(v *QueryPointActionAutoAssignRuleResponseBodyResult) *QueryPointActionAutoAssignRuleResponseBody {
	s.Result = v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponseBody) SetSuccess(v bool) *QueryPointActionAutoAssignRuleResponseBody {
	s.Success = &v
	return s
}

type QueryPointActionAutoAssignRuleResponseBodyResult struct {
	// 行为规则列表
	QueryPointRuleResponseDTOS []*QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS `json:"queryPointRuleResponseDTOS,omitempty" xml:"queryPointRuleResponseDTOS,omitempty" type:"Repeated"`
}

func (s QueryPointActionAutoAssignRuleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryPointActionAutoAssignRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResult) SetQueryPointRuleResponseDTOS(v []*QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) *QueryPointActionAutoAssignRuleResponseBodyResult {
	s.QueryPointRuleResponseDTOS = v
	return s
}

type QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS struct {
	// 奖励积分
	AwardScore *int64 `json:"awardScore,omitempty" xml:"awardScore,omitempty"`
	// 行为名称
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 单日计次上限
	DayLimitTimes *int64 `json:"dayLimitTimes,omitempty" xml:"dayLimitTimes,omitempty"`
	// 行为描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 生效状态：0无效，1有效
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) String() string {
	return tea.Prettify(s)
}

func (s QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) GoString() string {
	return s.String()
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) SetAwardScore(v int64) *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS {
	s.AwardScore = &v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) SetCode(v string) *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS {
	s.Code = &v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) SetDayLimitTimes(v int64) *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS {
	s.DayLimitTimes = &v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) SetDescription(v string) *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS {
	s.Description = &v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS) SetStatus(v int64) *QueryPointActionAutoAssignRuleResponseBodyResultQueryPointRuleResponseDTOS {
	s.Status = &v
	return s
}

type QueryPointActionAutoAssignRuleResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPointActionAutoAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPointActionAutoAssignRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPointActionAutoAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *QueryPointActionAutoAssignRuleResponse) SetHeaders(v map[string]*string) *QueryPointActionAutoAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *QueryPointActionAutoAssignRuleResponse) SetBody(v *QueryPointActionAutoAssignRuleResponseBody) *QueryPointActionAutoAssignRuleResponse {
	s.Body = v
	return s
}

type QueryPointAutoIssueSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPointAutoIssueSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPointAutoIssueSettingHeaders) GoString() string {
	return s.String()
}

func (s *QueryPointAutoIssueSettingHeaders) SetCommonHeaders(v map[string]*string) *QueryPointAutoIssueSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPointAutoIssueSettingHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPointAutoIssueSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPointAutoIssueSettingResponseBody struct {
	Result *QueryPointAutoIssueSettingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 调用是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryPointAutoIssueSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPointAutoIssueSettingResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPointAutoIssueSettingResponseBody) SetResult(v *QueryPointAutoIssueSettingResponseBodyResult) *QueryPointAutoIssueSettingResponseBody {
	s.Result = v
	return s
}

func (s *QueryPointAutoIssueSettingResponseBody) SetSuccess(v bool) *QueryPointAutoIssueSettingResponseBody {
	s.Success = &v
	return s
}

type QueryPointAutoIssueSettingResponseBodyResult struct {
	// 企业每月额度自动发放给每个人的数量
	PointAutoNum *int64 `json:"pointAutoNum,omitempty" xml:"pointAutoNum,omitempty"`
	// 企业积分自动发放状态
	PointAutoState *bool `json:"pointAutoState,omitempty" xml:"pointAutoState,omitempty"`
	// 企业积分自动发放时间 指定的是每月1号或15号
	PointAutoTime *int64 `json:"pointAutoTime,omitempty" xml:"pointAutoTime,omitempty"`
}

func (s QueryPointAutoIssueSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryPointAutoIssueSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryPointAutoIssueSettingResponseBodyResult) SetPointAutoNum(v int64) *QueryPointAutoIssueSettingResponseBodyResult {
	s.PointAutoNum = &v
	return s
}

func (s *QueryPointAutoIssueSettingResponseBodyResult) SetPointAutoState(v bool) *QueryPointAutoIssueSettingResponseBodyResult {
	s.PointAutoState = &v
	return s
}

func (s *QueryPointAutoIssueSettingResponseBodyResult) SetPointAutoTime(v int64) *QueryPointAutoIssueSettingResponseBodyResult {
	s.PointAutoTime = &v
	return s
}

type QueryPointAutoIssueSettingResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPointAutoIssueSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPointAutoIssueSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPointAutoIssueSettingResponse) GoString() string {
	return s.String()
}

func (s *QueryPointAutoIssueSettingResponse) SetHeaders(v map[string]*string) *QueryPointAutoIssueSettingResponse {
	s.Headers = v
	return s
}

func (s *QueryPointAutoIssueSettingResponse) SetBody(v *QueryPointAutoIssueSettingResponseBody) *QueryPointAutoIssueSettingResponse {
	s.Body = v
	return s
}

type QueryUserHonorsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserHonorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsHeaders) SetCommonHeaders(v map[string]*string) *QueryUserHonorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserHonorsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserHonorsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserHonorsRequest struct {
	// 查询数据的条数，默认查询20条，最大可传100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页查询的标记，查询第一页时传0，非第一页时传入上次调用本接口返回值中的nextToken
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryUserHonorsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsRequest) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsRequest) SetMaxResults(v int32) *QueryUserHonorsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryUserHonorsRequest) SetNextToken(v string) *QueryUserHonorsRequest {
	s.NextToken = &v
	return s
}

type QueryUserHonorsResponseBody struct {
	Result  *QueryUserHonorsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryUserHonorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBody) SetResult(v *QueryUserHonorsResponseBodyResult) *QueryUserHonorsResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserHonorsResponseBody) SetSuccess(v bool) *QueryUserHonorsResponseBody {
	s.Success = &v
	return s
}

type QueryUserHonorsResponseBodyResult struct {
	Honors    []*QueryUserHonorsResponseBodyResultHonors `json:"honors,omitempty" xml:"honors,omitempty" type:"Repeated"`
	NextToken *string                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryUserHonorsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyResult) SetHonors(v []*QueryUserHonorsResponseBodyResultHonors) *QueryUserHonorsResponseBodyResult {
	s.Honors = v
	return s
}

func (s *QueryUserHonorsResponseBodyResult) SetNextToken(v string) *QueryUserHonorsResponseBodyResult {
	s.NextToken = &v
	return s
}

type QueryUserHonorsResponseBodyResultHonors struct {
	// 有效期截止时间点，没有该属性则为永久生效
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// 授予历史记录 包含用户及授予时间
	GrantHistory []*QueryUserHonorsResponseBodyResultHonorsGrantHistory `json:"grantHistory,omitempty" xml:"grantHistory,omitempty" type:"Repeated"`
	// 荣誉含义
	HonorDesc *string `json:"honorDesc,omitempty" xml:"honorDesc,omitempty"`
	// 必须。荣誉id
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
	// 必须。荣誉名字
	HonorName *string `json:"honorName,omitempty" xml:"honorName,omitempty"`
}

func (s QueryUserHonorsResponseBodyResultHonors) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBodyResultHonors) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetExpirationTime(v int64) *QueryUserHonorsResponseBodyResultHonors {
	s.ExpirationTime = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetGrantHistory(v []*QueryUserHonorsResponseBodyResultHonorsGrantHistory) *QueryUserHonorsResponseBodyResultHonors {
	s.GrantHistory = v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetHonorDesc(v string) *QueryUserHonorsResponseBodyResultHonors {
	s.HonorDesc = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetHonorId(v string) *QueryUserHonorsResponseBodyResultHonors {
	s.HonorId = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetHonorName(v string) *QueryUserHonorsResponseBodyResultHonors {
	s.HonorName = &v
	return s
}

type QueryUserHonorsResponseBodyResultHonorsGrantHistory struct {
	// 授予时间 时间戳
	GrantTime *int64 `json:"grantTime,omitempty" xml:"grantTime,omitempty"`
	// 必须。荣誉发放人userid
	SenderUserid *string `json:"senderUserid,omitempty" xml:"senderUserid,omitempty"`
}

func (s QueryUserHonorsResponseBodyResultHonorsGrantHistory) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBodyResultHonorsGrantHistory) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyResultHonorsGrantHistory) SetGrantTime(v int64) *QueryUserHonorsResponseBodyResultHonorsGrantHistory {
	s.GrantTime = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonorsGrantHistory) SetSenderUserid(v string) *QueryUserHonorsResponseBodyResultHonorsGrantHistory {
	s.SenderUserid = &v
	return s
}

type QueryUserHonorsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserHonorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserHonorsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponse) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponse) SetHeaders(v map[string]*string) *QueryUserHonorsResponse {
	s.Headers = v
	return s
}

func (s *QueryUserHonorsResponse) SetBody(v *QueryUserHonorsResponseBody) *QueryUserHonorsResponse {
	s.Body = v
	return s
}

type QueryUserPointsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserPointsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserPointsHeaders) SetCommonHeaders(v map[string]*string) *QueryUserPointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserPointsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserPointsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserPointsResponseBody struct {
	// 响应数据
	Result *QueryUserPointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 请求响应状态
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryUserPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserPointsResponseBody) SetResult(v *QueryUserPointsResponseBodyResult) *QueryUserPointsResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserPointsResponseBody) SetSuccess(v bool) *QueryUserPointsResponseBody {
	s.Success = &v
	return s
}

type QueryUserPointsResponseBodyResult struct {
	// 员工积分数量
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s QueryUserPointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserPointsResponseBodyResult) SetAmount(v int64) *QueryUserPointsResponseBodyResult {
	s.Amount = &v
	return s
}

type QueryUserPointsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsResponse) GoString() string {
	return s.String()
}

func (s *QueryUserPointsResponse) SetHeaders(v map[string]*string) *QueryUserPointsResponse {
	s.Headers = v
	return s
}

func (s *QueryUserPointsResponse) SetBody(v *QueryUserPointsResponseBody) *QueryUserPointsResponse {
	s.Body = v
	return s
}

type RecallHonorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RecallHonorHeaders) String() string {
	return tea.Prettify(s)
}

func (s RecallHonorHeaders) GoString() string {
	return s.String()
}

func (s *RecallHonorHeaders) SetCommonHeaders(v map[string]*string) *RecallHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecallHonorHeaders) SetXAcsDingtalkAccessToken(v string) *RecallHonorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RecallHonorRequest struct {
	// 要取消荣誉的员工userid 必填
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RecallHonorRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallHonorRequest) GoString() string {
	return s.String()
}

func (s *RecallHonorRequest) SetUserId(v string) *RecallHonorRequest {
	s.UserId = &v
	return s
}

type RecallHonorResponseBody struct {
	Result  *RecallHonorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RecallHonorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecallHonorResponseBody) GoString() string {
	return s.String()
}

func (s *RecallHonorResponseBody) SetResult(v *RecallHonorResponseBodyResult) *RecallHonorResponseBody {
	s.Result = v
	return s
}

func (s *RecallHonorResponseBody) SetSuccess(v bool) *RecallHonorResponseBody {
	s.Success = &v
	return s
}

type RecallHonorResponseBodyResult struct {
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
}

func (s RecallHonorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RecallHonorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RecallHonorResponseBodyResult) SetHonorId(v string) *RecallHonorResponseBodyResult {
	s.HonorId = &v
	return s
}

type RecallHonorResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecallHonorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecallHonorResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallHonorResponse) GoString() string {
	return s.String()
}

func (s *RecallHonorResponse) SetHeaders(v map[string]*string) *RecallHonorResponse {
	s.Headers = v
	return s
}

func (s *RecallHonorResponse) SetBody(v *RecallHonorResponseBody) *RecallHonorResponse {
	s.Body = v
	return s
}

type UpdateAutoIssuePointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateAutoIssuePointHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoIssuePointHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAutoIssuePointHeaders) SetCommonHeaders(v map[string]*string) *UpdateAutoIssuePointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAutoIssuePointHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateAutoIssuePointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateAutoIssuePointRequest struct {
	// 企业积分自动发放数量1-10000
	PointAutoNum *int64 `json:"pointAutoNum,omitempty" xml:"pointAutoNum,omitempty"`
	// 企业积分自动发放状态
	PointAutoState *bool `json:"pointAutoState,omitempty" xml:"pointAutoState,omitempty"`
	// 企业积分自动发放时间 必须为每月的1号或15号，传入1时为1号，传入15时为15号。
	PointAutoTime *int64 `json:"pointAutoTime,omitempty" xml:"pointAutoTime,omitempty"`
	// 操作人userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateAutoIssuePointRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoIssuePointRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoIssuePointRequest) SetPointAutoNum(v int64) *UpdateAutoIssuePointRequest {
	s.PointAutoNum = &v
	return s
}

func (s *UpdateAutoIssuePointRequest) SetPointAutoState(v bool) *UpdateAutoIssuePointRequest {
	s.PointAutoState = &v
	return s
}

func (s *UpdateAutoIssuePointRequest) SetPointAutoTime(v int64) *UpdateAutoIssuePointRequest {
	s.PointAutoTime = &v
	return s
}

func (s *UpdateAutoIssuePointRequest) SetUserId(v string) *UpdateAutoIssuePointRequest {
	s.UserId = &v
	return s
}

type UpdateAutoIssuePointResponseBody struct {
	Result *UpdateAutoIssuePointResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 调用是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAutoIssuePointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoIssuePointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoIssuePointResponseBody) SetResult(v *UpdateAutoIssuePointResponseBodyResult) *UpdateAutoIssuePointResponseBody {
	s.Result = v
	return s
}

func (s *UpdateAutoIssuePointResponseBody) SetSuccess(v bool) *UpdateAutoIssuePointResponseBody {
	s.Success = &v
	return s
}

type UpdateAutoIssuePointResponseBodyResult struct {
	// 下次自动发放时间
	NextAutoIssuePointTime *int64 `json:"nextAutoIssuePointTime,omitempty" xml:"nextAutoIssuePointTime,omitempty"`
}

func (s UpdateAutoIssuePointResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoIssuePointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateAutoIssuePointResponseBodyResult) SetNextAutoIssuePointTime(v int64) *UpdateAutoIssuePointResponseBodyResult {
	s.NextAutoIssuePointTime = &v
	return s
}

type UpdateAutoIssuePointResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAutoIssuePointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAutoIssuePointResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoIssuePointResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoIssuePointResponse) SetHeaders(v map[string]*string) *UpdateAutoIssuePointResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoIssuePointResponse) SetBody(v *UpdateAutoIssuePointResponseBody) *UpdateAutoIssuePointResponse {
	s.Body = v
	return s
}

type UpdatePointActionAutoAssignRuleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdatePointActionAutoAssignRuleHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePointActionAutoAssignRuleHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePointActionAutoAssignRuleHeaders) SetCommonHeaders(v map[string]*string) *UpdatePointActionAutoAssignRuleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePointActionAutoAssignRuleHeaders) SetXAcsDingtalkAccessToken(v string) *UpdatePointActionAutoAssignRuleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdatePointActionAutoAssignRuleRequest struct {
	// 行为规则列表
	UpdatePointRuleRequestDTOList []*UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList `json:"updatePointRuleRequestDTOList,omitempty" xml:"updatePointRuleRequestDTOList,omitempty" type:"Repeated"`
	// 操作人userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdatePointActionAutoAssignRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePointActionAutoAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdatePointActionAutoAssignRuleRequest) SetUpdatePointRuleRequestDTOList(v []*UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) *UpdatePointActionAutoAssignRuleRequest {
	s.UpdatePointRuleRequestDTOList = v
	return s
}

func (s *UpdatePointActionAutoAssignRuleRequest) SetUserId(v string) *UpdatePointActionAutoAssignRuleRequest {
	s.UserId = &v
	return s
}

type UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList struct {
	// 奖励积分1～100
	AwardScore *int64 `json:"awardScore,omitempty" xml:"awardScore,omitempty"`
	// 行为名称 不支持修改
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 单日计次上限 1～10
	DayLimitTimes *int64 `json:"dayLimitTimes,omitempty" xml:"dayLimitTimes,omitempty"`
	// 生效状态：0无效，1有效
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) String() string {
	return tea.Prettify(s)
}

func (s UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) GoString() string {
	return s.String()
}

func (s *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) SetAwardScore(v int64) *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList {
	s.AwardScore = &v
	return s
}

func (s *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) SetCode(v string) *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList {
	s.Code = &v
	return s
}

func (s *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) SetDayLimitTimes(v int64) *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList {
	s.DayLimitTimes = &v
	return s
}

func (s *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList) SetStatus(v int64) *UpdatePointActionAutoAssignRuleRequestUpdatePointRuleRequestDTOList {
	s.Status = &v
	return s
}

type UpdatePointActionAutoAssignRuleResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePointActionAutoAssignRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePointActionAutoAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePointActionAutoAssignRuleResponseBody) SetSuccess(v bool) *UpdatePointActionAutoAssignRuleResponseBody {
	s.Success = &v
	return s
}

type UpdatePointActionAutoAssignRuleResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePointActionAutoAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePointActionAutoAssignRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePointActionAutoAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdatePointActionAutoAssignRuleResponse) SetHeaders(v map[string]*string) *UpdatePointActionAutoAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdatePointActionAutoAssignRuleResponse) SetBody(v *UpdatePointActionAutoAssignRuleResponseBody) *UpdatePointActionAutoAssignRuleResponse {
	s.Body = v
	return s
}

type WearOrgHonorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WearOrgHonorHeaders) String() string {
	return tea.Prettify(s)
}

func (s WearOrgHonorHeaders) GoString() string {
	return s.String()
}

func (s *WearOrgHonorHeaders) SetCommonHeaders(v map[string]*string) *WearOrgHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WearOrgHonorHeaders) SetXAcsDingtalkAccessToken(v string) *WearOrgHonorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WearOrgHonorRequest struct {
	// 员工userid
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 佩戴true，卸下false
	Wear *bool `json:"wear,omitempty" xml:"wear,omitempty"`
}

func (s WearOrgHonorRequest) String() string {
	return tea.Prettify(s)
}

func (s WearOrgHonorRequest) GoString() string {
	return s.String()
}

func (s *WearOrgHonorRequest) SetUserId(v string) *WearOrgHonorRequest {
	s.UserId = &v
	return s
}

func (s *WearOrgHonorRequest) SetWear(v bool) *WearOrgHonorRequest {
	s.Wear = &v
	return s
}

type WearOrgHonorResponseBody struct {
	Result  *WearOrgHonorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WearOrgHonorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WearOrgHonorResponseBody) GoString() string {
	return s.String()
}

func (s *WearOrgHonorResponseBody) SetResult(v *WearOrgHonorResponseBodyResult) *WearOrgHonorResponseBody {
	s.Result = v
	return s
}

func (s *WearOrgHonorResponseBody) SetSuccess(v bool) *WearOrgHonorResponseBody {
	s.Success = &v
	return s
}

type WearOrgHonorResponseBodyResult struct {
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
}

func (s WearOrgHonorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s WearOrgHonorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *WearOrgHonorResponseBodyResult) SetHonorId(v string) *WearOrgHonorResponseBodyResult {
	s.HonorId = &v
	return s
}

type WearOrgHonorResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WearOrgHonorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WearOrgHonorResponse) String() string {
	return tea.Prettify(s)
}

func (s WearOrgHonorResponse) GoString() string {
	return s.String()
}

func (s *WearOrgHonorResponse) SetHeaders(v map[string]*string) *WearOrgHonorResponse {
	s.Headers = v
	return s
}

func (s *WearOrgHonorResponse) SetBody(v *WearOrgHonorResponseBody) *WearOrgHonorResponse {
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

func (client *Client) AssignOrgHoldingToEmpHoldingBatch(request *AssignOrgHoldingToEmpHoldingBatchRequest) (_result *AssignOrgHoldingToEmpHoldingBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AssignOrgHoldingToEmpHoldingBatchHeaders{}
	_result = &AssignOrgHoldingToEmpHoldingBatchResponse{}
	_body, _err := client.AssignOrgHoldingToEmpHoldingBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssignOrgHoldingToEmpHoldingBatchWithOptions(request *AssignOrgHoldingToEmpHoldingBatchRequest, headers *AssignOrgHoldingToEmpHoldingBatchHeaders, runtime *util.RuntimeOptions) (_result *AssignOrgHoldingToEmpHoldingBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SendOrgCultureInform)) {
		body["sendOrgCultureInform"] = request.SendOrgCultureInform
	}

	if !tea.BoolValue(util.IsUnset(request.SingleAmount)) {
		body["singleAmount"] = request.SingleAmount
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUsage)) {
		body["sourceUsage"] = request.SourceUsage
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUsage)) {
		body["targetUsage"] = request.TargetUsage
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUserList)) {
		body["targetUserList"] = request.TargetUserList
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
	_result = &AssignOrgHoldingToEmpHoldingBatchResponse{}
	_body, _err := client.DoROARequest(tea.String("AssignOrgHoldingToEmpHoldingBatch"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/organizations/points/assign"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConsumeUserPoints(userId *string, request *ConsumeUserPointsRequest) (_result *ConsumeUserPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConsumeUserPointsHeaders{}
	_result = &ConsumeUserPointsResponse{}
	_body, _err := client.ConsumeUserPointsWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConsumeUserPointsWithOptions(userId *string, request *ConsumeUserPointsRequest, headers *ConsumeUserPointsHeaders, runtime *util.RuntimeOptions) (_result *ConsumeUserPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		body["outId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Usage)) {
		body["usage"] = request.Usage
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
	_result = &ConsumeUserPointsResponse{}
	_body, _err := client.DoROARequest(tea.String("ConsumeUserPoints"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/users/"+tea.StringValue(userId)+"/points/deduct"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrgHonor(request *CreateOrgHonorRequest) (_result *CreateOrgHonorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOrgHonorHeaders{}
	_result = &CreateOrgHonorResponse{}
	_body, _err := client.CreateOrgHonorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrgHonorWithOptions(request *CreateOrgHonorRequest, headers *CreateOrgHonorHeaders, runtime *util.RuntimeOptions) (_result *CreateOrgHonorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvatarFrameMediaId)) {
		body["avatarFrameMediaId"] = request.AvatarFrameMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultBgColor)) {
		body["defaultBgColor"] = request.DefaultBgColor
	}

	if !tea.BoolValue(util.IsUnset(request.MedalDesc)) {
		body["medalDesc"] = request.MedalDesc
	}

	if !tea.BoolValue(util.IsUnset(request.MedalMediaId)) {
		body["medalMediaId"] = request.MedalMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.MedalName)) {
		body["medalName"] = request.MedalName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &CreateOrgHonorResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateOrgHonor"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/honors/templates"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeductionPointBatch(request *DeductionPointBatchRequest) (_result *DeductionPointBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeductionPointBatchHeaders{}
	_result = &DeductionPointBatchResponse{}
	_body, _err := client.DeductionPointBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeductionPointBatchWithOptions(request *DeductionPointBatchRequest, headers *DeductionPointBatchHeaders, runtime *util.RuntimeOptions) (_result *DeductionPointBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeductionAmount)) {
		body["deductionAmount"] = request.DeductionAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SendOrgCultureInform)) {
		body["sendOrgCultureInform"] = request.SendOrgCultureInform
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUserList)) {
		body["targetUserList"] = request.TargetUserList
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &DeductionPointBatchResponse{}
	_body, _err := client.DoROARequest(tea.String("DeductionPointBatch"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/users/points/deduct"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportPointOpen(request *ExportPointOpenRequest) (_result *ExportPointOpenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExportPointOpenHeaders{}
	_result = &ExportPointOpenResponse{}
	_body, _err := client.ExportPointOpenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportPointOpenWithOptions(request *ExportPointOpenRequest, headers *ExportPointOpenHeaders, runtime *util.RuntimeOptions) (_result *ExportPointOpenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExportDate)) {
		body["exportDate"] = request.ExportDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExportType)) {
		body["exportType"] = request.ExportType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &ExportPointOpenResponse{}
	_body, _err := client.DoROARequest(tea.String("ExportPointOpen"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/users/points/export"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantHonor(honorId *string, request *GrantHonorRequest) (_result *GrantHonorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GrantHonorHeaders{}
	_result = &GrantHonorResponse{}
	_body, _err := client.GrantHonorWithOptions(honorId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantHonorWithOptions(honorId *string, request *GrantHonorRequest, headers *GrantHonorHeaders, runtime *util.RuntimeOptions) (_result *GrantHonorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	honorId = openapiutil.GetEncodeParam(honorId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpirationTime)) {
		body["expirationTime"] = request.ExpirationTime
	}

	if !tea.BoolValue(util.IsUnset(request.GrantReason)) {
		body["grantReason"] = request.GrantReason
	}

	if !tea.BoolValue(util.IsUnset(request.GranterName)) {
		body["granterName"] = request.GranterName
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeAnnouncer)) {
		body["noticeAnnouncer"] = request.NoticeAnnouncer
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeSingle)) {
		body["noticeSingle"] = request.NoticeSingle
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationIds)) {
		body["openConversationIds"] = request.OpenConversationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIds)) {
		body["receiverUserIds"] = request.ReceiverUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUserId)) {
		body["senderUserId"] = request.SenderUserId
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
	_result = &GrantHonorResponse{}
	_body, _err := client.DoROARequest(tea.String("GrantHonor"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/honors/"+tea.StringValue(honorId)+"/grant"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCorpPoints(request *QueryCorpPointsRequest) (_result *QueryCorpPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCorpPointsHeaders{}
	_result = &QueryCorpPointsResponse{}
	_body, _err := client.QueryCorpPointsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCorpPointsWithOptions(request *QueryCorpPointsRequest, headers *QueryCorpPointsHeaders, runtime *util.RuntimeOptions) (_result *QueryCorpPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		query["optUserId"] = request.OptUserId
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
	_result = &QueryCorpPointsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCorpPoints"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/organizations/points"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEmpPointDetails(request *QueryEmpPointDetailsRequest) (_result *QueryEmpPointDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryEmpPointDetailsHeaders{}
	_result = &QueryEmpPointDetailsResponse{}
	_body, _err := client.QueryEmpPointDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEmpPointDetailsWithOptions(request *QueryEmpPointDetailsRequest, headers *QueryEmpPointDetailsHeaders, runtime *util.RuntimeOptions) (_result *QueryEmpPointDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
	_result = &QueryEmpPointDetailsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryEmpPointDetails"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/points/empDetails"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgHonors(request *QueryOrgHonorsRequest) (_result *QueryOrgHonorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgHonorsHeaders{}
	_result = &QueryOrgHonorsResponse{}
	_body, _err := client.QueryOrgHonorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgHonorsWithOptions(request *QueryOrgHonorsRequest, headers *QueryOrgHonorsHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgHonorsResponse, _err error) {
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
	_result = &QueryOrgHonorsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOrgHonors"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/organizations/honors"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgPointDetails(request *QueryOrgPointDetailsRequest) (_result *QueryOrgPointDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgPointDetailsHeaders{}
	_result = &QueryOrgPointDetailsResponse{}
	_body, _err := client.QueryOrgPointDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgPointDetailsWithOptions(request *QueryOrgPointDetailsRequest, headers *QueryOrgPointDetailsHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgPointDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["accountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
	_result = &QueryOrgPointDetailsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOrgPointDetails"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/points/orgDetails"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPointActionAutoAssignRule() (_result *QueryPointActionAutoAssignRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPointActionAutoAssignRuleHeaders{}
	_result = &QueryPointActionAutoAssignRuleResponse{}
	_body, _err := client.QueryPointActionAutoAssignRuleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPointActionAutoAssignRuleWithOptions(headers *QueryPointActionAutoAssignRuleHeaders, runtime *util.RuntimeOptions) (_result *QueryPointActionAutoAssignRuleResponse, _err error) {
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
	_result = &QueryPointActionAutoAssignRuleResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryPointActionAutoAssignRule"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/users/points/actionRules"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPointAutoIssueSetting() (_result *QueryPointAutoIssueSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPointAutoIssueSettingHeaders{}
	_result = &QueryPointAutoIssueSettingResponse{}
	_body, _err := client.QueryPointAutoIssueSettingWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPointAutoIssueSettingWithOptions(headers *QueryPointAutoIssueSettingHeaders, runtime *util.RuntimeOptions) (_result *QueryPointAutoIssueSettingResponse, _err error) {
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
	_result = &QueryPointAutoIssueSettingResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryPointAutoIssueSetting"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/users/points"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserHonors(userId *string, request *QueryUserHonorsRequest) (_result *QueryUserHonorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserHonorsHeaders{}
	_result = &QueryUserHonorsResponse{}
	_body, _err := client.QueryUserHonorsWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserHonorsWithOptions(userId *string, request *QueryUserHonorsRequest, headers *QueryUserHonorsHeaders, runtime *util.RuntimeOptions) (_result *QueryUserHonorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
	_result = &QueryUserHonorsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserHonors"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/honors/users/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserPoints(userId *string) (_result *QueryUserPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserPointsHeaders{}
	_result = &QueryUserPointsResponse{}
	_body, _err := client.QueryUserPointsWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserPointsWithOptions(userId *string, headers *QueryUserPointsHeaders, runtime *util.RuntimeOptions) (_result *QueryUserPointsResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
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
	_result = &QueryUserPointsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserPoints"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/users/"+tea.StringValue(userId)+"/points"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecallHonor(honorId *string, request *RecallHonorRequest) (_result *RecallHonorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecallHonorHeaders{}
	_result = &RecallHonorResponse{}
	_body, _err := client.RecallHonorWithOptions(honorId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecallHonorWithOptions(honorId *string, request *RecallHonorRequest, headers *RecallHonorHeaders, runtime *util.RuntimeOptions) (_result *RecallHonorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	honorId = openapiutil.GetEncodeParam(honorId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &RecallHonorResponse{}
	_body, _err := client.DoROARequest(tea.String("RecallHonor"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/honors/"+tea.StringValue(honorId)+"/recall"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAutoIssuePoint(request *UpdateAutoIssuePointRequest) (_result *UpdateAutoIssuePointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateAutoIssuePointHeaders{}
	_result = &UpdateAutoIssuePointResponse{}
	_body, _err := client.UpdateAutoIssuePointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAutoIssuePointWithOptions(request *UpdateAutoIssuePointRequest, headers *UpdateAutoIssuePointHeaders, runtime *util.RuntimeOptions) (_result *UpdateAutoIssuePointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PointAutoNum)) {
		body["pointAutoNum"] = request.PointAutoNum
	}

	if !tea.BoolValue(util.IsUnset(request.PointAutoState)) {
		body["pointAutoState"] = request.PointAutoState
	}

	if !tea.BoolValue(util.IsUnset(request.PointAutoTime)) {
		body["pointAutoTime"] = request.PointAutoTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &UpdateAutoIssuePointResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateAutoIssuePoint"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/users/points/set"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePointActionAutoAssignRule(request *UpdatePointActionAutoAssignRuleRequest) (_result *UpdatePointActionAutoAssignRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePointActionAutoAssignRuleHeaders{}
	_result = &UpdatePointActionAutoAssignRuleResponse{}
	_body, _err := client.UpdatePointActionAutoAssignRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePointActionAutoAssignRuleWithOptions(request *UpdatePointActionAutoAssignRuleRequest, headers *UpdatePointActionAutoAssignRuleHeaders, runtime *util.RuntimeOptions) (_result *UpdatePointActionAutoAssignRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdatePointRuleRequestDTOList)) {
		body["updatePointRuleRequestDTOList"] = request.UpdatePointRuleRequestDTOList
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &UpdatePointActionAutoAssignRuleResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdatePointActionAutoAssignRule"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/orgCulture/users/points/actionRules"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WearOrgHonor(honorId *string, request *WearOrgHonorRequest) (_result *WearOrgHonorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WearOrgHonorHeaders{}
	_result = &WearOrgHonorResponse{}
	_body, _err := client.WearOrgHonorWithOptions(honorId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WearOrgHonorWithOptions(honorId *string, request *WearOrgHonorRequest, headers *WearOrgHonorHeaders, runtime *util.RuntimeOptions) (_result *WearOrgHonorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	honorId = openapiutil.GetEncodeParam(honorId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Wear)) {
		body["wear"] = request.Wear
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
	_result = &WearOrgHonorResponse{}
	_body, _err := client.DoROARequest(tea.String("WearOrgHonor"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/honors/"+tea.StringValue(honorId)+"/wear"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
