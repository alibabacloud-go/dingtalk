// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package datacenter_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryGroupMessageStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupMessageStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMessageStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupMessageStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupMessageStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupMessageStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupMessageStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupMessageStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryGroupMessageStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMessageStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupMessageStatisticalDataRequest) SetStatDate(v string) *QueryGroupMessageStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryGroupMessageStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryGroupMessageStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryGroupMessageStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMessageStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupMessageStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryGroupMessageStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponseBody) SetMetaList(v []*QueryGroupMessageStatisticalDataResponseBodyMetaList) *QueryGroupMessageStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryGroupMessageStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryGroupMessageStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMessageStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryGroupMessageStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryGroupMessageStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryGroupMessageStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryGroupMessageStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryGroupMessageStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryGroupMessageStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryGroupMessageStatisticalDataResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupMessageStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupMessageStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMessageStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupMessageStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryGroupMessageStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupMessageStatisticalDataResponse) SetBody(v *QueryGroupMessageStatisticalDataResponseBody) *QueryGroupMessageStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryVedioMeetingStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryVedioMeetingStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryVedioMeetingStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryVedioMeetingStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryVedioMeetingStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryVedioMeetingStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryVedioMeetingStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryVedioMeetingStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryVedioMeetingStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVedioMeetingStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryVedioMeetingStatisticalDataRequest) SetStatDate(v string) *QueryVedioMeetingStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryVedioMeetingStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryVedioMeetingStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryVedioMeetingStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVedioMeetingStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVedioMeetingStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryVedioMeetingStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponseBody) SetMetaList(v []*QueryVedioMeetingStatisticalDataResponseBodyMetaList) *QueryVedioMeetingStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryVedioMeetingStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryVedioMeetingStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryVedioMeetingStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryVedioMeetingStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryVedioMeetingStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryVedioMeetingStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryVedioMeetingStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryVedioMeetingStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryVedioMeetingStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryVedioMeetingStatisticalDataResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryVedioMeetingStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryVedioMeetingStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVedioMeetingStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryVedioMeetingStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryVedioMeetingStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryVedioMeetingStatisticalDataResponse) SetBody(v *QueryVedioMeetingStatisticalDataResponseBody) *QueryVedioMeetingStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryHealthStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHealthStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryHealthStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryHealthStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHealthStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHealthStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHealthStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryHealthStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryHealthStatisticalDataRequest) SetStatDate(v string) *QueryHealthStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryHealthStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryHealthStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryHealthStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHealthStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryHealthStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryHealthStatisticalDataResponseBody) SetMetaList(v []*QueryHealthStatisticalDataResponseBodyMetaList) *QueryHealthStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryHealthStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryHealthStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryHealthStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryHealthStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryHealthStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryHealthStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryHealthStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryHealthStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryHealthStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryHealthStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryHealthStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryHealthStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryHealthStatisticalDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHealthStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHealthStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryHealthStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryHealthStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryHealthStatisticalDataResponse) SetBody(v *QueryHealthStatisticalDataResponseBody) *QueryHealthStatisticalDataResponse {
	s.Body = v
	return s
}

type QuerySingleMessageStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySingleMessageStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleMessageStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QuerySingleMessageStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QuerySingleMessageStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySingleMessageStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySingleMessageStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySingleMessageStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QuerySingleMessageStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleMessageStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QuerySingleMessageStatisticalDataRequest) SetStatDate(v string) *QuerySingleMessageStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QuerySingleMessageStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QuerySingleMessageStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QuerySingleMessageStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleMessageStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySingleMessageStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QuerySingleMessageStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponseBody) SetMetaList(v []*QuerySingleMessageStatisticalDataResponseBodyMetaList) *QuerySingleMessageStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QuerySingleMessageStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QuerySingleMessageStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleMessageStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QuerySingleMessageStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QuerySingleMessageStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QuerySingleMessageStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponseBodyMetaList) SetUnit(v string) *QuerySingleMessageStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QuerySingleMessageStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QuerySingleMessageStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QuerySingleMessageStatisticalDataResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySingleMessageStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySingleMessageStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleMessageStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QuerySingleMessageStatisticalDataResponse) SetHeaders(v map[string]*string) *QuerySingleMessageStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QuerySingleMessageStatisticalDataResponse) SetBody(v *QuerySingleMessageStatisticalDataResponseBody) *QuerySingleMessageStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryTodoStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTodoStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryTodoStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryTodoStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTodoStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTodoStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTodoStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryTodoStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryTodoStatisticalDataRequest) SetStatDate(v string) *QueryTodoStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryTodoStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryTodoStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryTodoStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTodoStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryTodoStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryTodoStatisticalDataResponseBody) SetMetaList(v []*QueryTodoStatisticalDataResponseBodyMetaList) *QueryTodoStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryTodoStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryTodoStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryTodoStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryTodoStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryTodoStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryTodoStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryTodoStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryTodoStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryTodoStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryTodoStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryTodoStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryTodoStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryTodoStatisticalDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTodoStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTodoStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTodoStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryTodoStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryTodoStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryTodoStatisticalDataResponse) SetBody(v *QueryTodoStatisticalDataResponseBody) *QueryTodoStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryCheckinStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCheckinStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckinStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryCheckinStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryCheckinStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCheckinStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCheckinStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCheckinStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryCheckinStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckinStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryCheckinStatisticalDataRequest) SetStatDate(v string) *QueryCheckinStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryCheckinStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryCheckinStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryCheckinStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckinStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCheckinStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryCheckinStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryCheckinStatisticalDataResponseBody) SetMetaList(v []*QueryCheckinStatisticalDataResponseBodyMetaList) *QueryCheckinStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryCheckinStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryCheckinStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckinStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryCheckinStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryCheckinStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryCheckinStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryCheckinStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryCheckinStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryCheckinStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryCheckinStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryCheckinStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryCheckinStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryCheckinStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryCheckinStatisticalDataResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCheckinStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCheckinStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCheckinStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryCheckinStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryCheckinStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryCheckinStatisticalDataResponse) SetBody(v *QueryCheckinStatisticalDataResponseBody) *QueryCheckinStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryEmployeeTypeStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryEmployeeTypeStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryEmployeeTypeStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryEmployeeTypeStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryEmployeeTypeStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryEmployeeTypeStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryEmployeeTypeStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryEmployeeTypeStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEmployeeTypeStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryEmployeeTypeStatisticalDataRequest) SetStatDate(v string) *QueryEmployeeTypeStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryEmployeeTypeStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryEmployeeTypeStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryEmployeeTypeStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEmployeeTypeStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmployeeTypeStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryEmployeeTypeStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponseBody) SetMetaList(v []*QueryEmployeeTypeStatisticalDataResponseBodyMetaList) *QueryEmployeeTypeStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryEmployeeTypeStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryEmployeeTypeStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryEmployeeTypeStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryEmployeeTypeStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryEmployeeTypeStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryEmployeeTypeStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryEmployeeTypeStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryEmployeeTypeStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryEmployeeTypeStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryEmployeeTypeStatisticalDataResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEmployeeTypeStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEmployeeTypeStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEmployeeTypeStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryEmployeeTypeStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryEmployeeTypeStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryEmployeeTypeStatisticalDataResponse) SetBody(v *QueryEmployeeTypeStatisticalDataResponseBody) *QueryEmployeeTypeStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryMailStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMailStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMailStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryMailStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryMailStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMailStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMailStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMailStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryMailStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMailStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryMailStatisticalDataRequest) SetStatDate(v string) *QueryMailStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryMailStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryMailStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryMailStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMailStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMailStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryMailStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryMailStatisticalDataResponseBody) SetMetaList(v []*QueryMailStatisticalDataResponseBodyMetaList) *QueryMailStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryMailStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryMailStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryMailStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryMailStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryMailStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryMailStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryMailStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryMailStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryMailStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryMailStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryMailStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryMailStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryMailStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryMailStatisticalDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMailStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMailStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMailStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryMailStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryMailStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryMailStatisticalDataResponse) SetBody(v *QueryMailStatisticalDataResponseBody) *QueryMailStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryCalendarStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCalendarStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCalendarStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryCalendarStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryCalendarStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCalendarStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCalendarStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCalendarStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryCalendarStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCalendarStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryCalendarStatisticalDataRequest) SetStatDate(v string) *QueryCalendarStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryCalendarStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryCalendarStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryCalendarStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCalendarStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCalendarStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryCalendarStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryCalendarStatisticalDataResponseBody) SetMetaList(v []*QueryCalendarStatisticalDataResponseBodyMetaList) *QueryCalendarStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryCalendarStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryCalendarStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryCalendarStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryCalendarStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryCalendarStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryCalendarStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryCalendarStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryCalendarStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryCalendarStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryCalendarStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryCalendarStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryCalendarStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryCalendarStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryCalendarStatisticalDataResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCalendarStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCalendarStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCalendarStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryCalendarStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryCalendarStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryCalendarStatisticalDataResponse) SetBody(v *QueryCalendarStatisticalDataResponseBody) *QueryCalendarStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryDocumentStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDocumentStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDocumentStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryDocumentStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryDocumentStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDocumentStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDocumentStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDocumentStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryDocumentStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDocumentStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDocumentStatisticalDataRequest) SetStatDate(v string) *QueryDocumentStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryDocumentStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryDocumentStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryDocumentStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDocumentStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDocumentStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryDocumentStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryDocumentStatisticalDataResponseBody) SetMetaList(v []*QueryDocumentStatisticalDataResponseBodyMetaList) *QueryDocumentStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryDocumentStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryDocumentStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryDocumentStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryDocumentStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryDocumentStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryDocumentStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryDocumentStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryDocumentStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryDocumentStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryDocumentStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryDocumentStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryDocumentStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryDocumentStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryDocumentStatisticalDataResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDocumentStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDocumentStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDocumentStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDocumentStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryDocumentStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDocumentStatisticalDataResponse) SetBody(v *QueryDocumentStatisticalDataResponseBody) *QueryDocumentStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryReportStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryReportStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryReportStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryReportStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryReportStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryReportStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryReportStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryReportStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryReportStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReportStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryReportStatisticalDataRequest) SetStatDate(v string) *QueryReportStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryReportStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryReportStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryReportStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReportStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReportStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryReportStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryReportStatisticalDataResponseBody) SetMetaList(v []*QueryReportStatisticalDataResponseBodyMetaList) *QueryReportStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryReportStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryReportStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryReportStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryReportStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryReportStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryReportStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryReportStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryReportStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryReportStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryReportStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryReportStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryReportStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryReportStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryReportStatisticalDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryReportStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryReportStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReportStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryReportStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryReportStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryReportStatisticalDataResponse) SetBody(v *QueryReportStatisticalDataResponseBody) *QueryReportStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryOnlineUserStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOnlineUserStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOnlineUserStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryOnlineUserStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryOnlineUserStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOnlineUserStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOnlineUserStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOnlineUserStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryOnlineUserStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOnlineUserStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryOnlineUserStatisticalDataRequest) SetStatDate(v string) *QueryOnlineUserStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryOnlineUserStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryOnlineUserStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryOnlineUserStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOnlineUserStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOnlineUserStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryOnlineUserStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponseBody) SetMetaList(v []*QueryOnlineUserStatisticalDataResponseBodyMetaList) *QueryOnlineUserStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryOnlineUserStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryOnlineUserStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryOnlineUserStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryOnlineUserStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryOnlineUserStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryOnlineUserStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryOnlineUserStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryOnlineUserStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryOnlineUserStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryOnlineUserStatisticalDataResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOnlineUserStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOnlineUserStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOnlineUserStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryOnlineUserStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryOnlineUserStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryOnlineUserStatisticalDataResponse) SetBody(v *QueryOnlineUserStatisticalDataResponseBody) *QueryOnlineUserStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryCompanyBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCompanyBasicInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCompanyBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryCompanyBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryCompanyBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCompanyBasicInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCompanyBasicInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCompanyBasicInfoRequest struct {
	Keyword    *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryCompanyBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCompanyBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCompanyBasicInfoRequest) SetKeyword(v string) *QueryCompanyBasicInfoRequest {
	s.Keyword = &v
	return s
}

func (s *QueryCompanyBasicInfoRequest) SetPageNumber(v int64) *QueryCompanyBasicInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryCompanyBasicInfoRequest) SetPageSize(v int64) *QueryCompanyBasicInfoRequest {
	s.PageSize = &v
	return s
}

type QueryCompanyBasicInfoResponseBody struct {
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// traceId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// total
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
	// data
	Data []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s QueryCompanyBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCompanyBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCompanyBasicInfoResponseBody) SetMessage(v string) *QueryCompanyBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCompanyBasicInfoResponseBody) SetRequestId(v string) *QueryCompanyBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCompanyBasicInfoResponseBody) SetTotal(v int32) *QueryCompanyBasicInfoResponseBody {
	s.Total = &v
	return s
}

func (s *QueryCompanyBasicInfoResponseBody) SetData(v []map[string]*string) *QueryCompanyBasicInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryCompanyBasicInfoResponseBody) SetCode(v string) *QueryCompanyBasicInfoResponseBody {
	s.Code = &v
	return s
}

type QueryCompanyBasicInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCompanyBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCompanyBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCompanyBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCompanyBasicInfoResponse) SetHeaders(v map[string]*string) *QueryCompanyBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCompanyBasicInfoResponse) SetBody(v *QueryCompanyBasicInfoResponseBody) *QueryCompanyBasicInfoResponse {
	s.Body = v
	return s
}

type QueryApprovalStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryApprovalStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryApprovalStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryApprovalStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryApprovalStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryApprovalStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryApprovalStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryApprovalStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryApprovalStatisticalDataRequest) SetStatDate(v string) *QueryApprovalStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryApprovalStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryApprovalStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryApprovalStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryApprovalStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryApprovalStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryApprovalStatisticalDataResponseBody) SetMetaList(v []*QueryApprovalStatisticalDataResponseBodyMetaList) *QueryApprovalStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryApprovalStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryApprovalStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryApprovalStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryApprovalStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryApprovalStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryApprovalStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryApprovalStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryApprovalStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryApprovalStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryApprovalStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryApprovalStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryApprovalStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryApprovalStatisticalDataResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryApprovalStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryApprovalStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryApprovalStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryApprovalStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryApprovalStatisticalDataResponse) SetBody(v *QueryApprovalStatisticalDataResponseBody) *QueryApprovalStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryDriveStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDriveStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDriveStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryDriveStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryDriveStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDriveStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDriveStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDriveStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryDriveStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDriveStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDriveStatisticalDataRequest) SetStatDate(v string) *QueryDriveStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryDriveStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryDriveStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryDriveStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDriveStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDriveStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryDriveStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryDriveStatisticalDataResponseBody) SetMetaList(v []*QueryDriveStatisticalDataResponseBodyMetaList) *QueryDriveStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryDriveStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryDriveStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryDriveStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryDriveStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryDriveStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryDriveStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryDriveStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryDriveStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryDriveStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryDriveStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryDriveStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryDriveStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryDriveStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryDriveStatisticalDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDriveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDriveStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDriveStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDriveStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryDriveStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDriveStatisticalDataResponse) SetBody(v *QueryDriveStatisticalDataResponseBody) *QueryDriveStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryDingSendStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDingSendStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDingSendStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryDingSendStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryDingSendStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDingSendStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDingSendStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDingSendStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryDingSendStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDingSendStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDingSendStatisticalDataRequest) SetStatDate(v string) *QueryDingSendStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryDingSendStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryDingSendStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryDingSendStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDingSendStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDingSendStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryDingSendStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryDingSendStatisticalDataResponseBody) SetMetaList(v []*QueryDingSendStatisticalDataResponseBodyMetaList) *QueryDingSendStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryDingSendStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryDingSendStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryDingSendStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryDingSendStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryDingSendStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryDingSendStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryDingSendStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryDingSendStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryDingSendStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryDingSendStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryDingSendStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryDingSendStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryDingSendStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryDingSendStatisticalDataResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDingSendStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDingSendStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDingSendStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDingSendStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryDingSendStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDingSendStatisticalDataResponse) SetBody(v *QueryDingSendStatisticalDataResponseBody) *QueryDingSendStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryActiveUserStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryActiveUserStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUserStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryActiveUserStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryActiveUserStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryActiveUserStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryActiveUserStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryActiveUserStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUserStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticalDataRequest) SetStatDate(v string) *QueryActiveUserStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryActiveUserStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryActiveUserStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryActiveUserStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUserStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryActiveUserStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryActiveUserStatisticalDataResponseBody) SetMetaList(v []*QueryActiveUserStatisticalDataResponseBodyMetaList) *QueryActiveUserStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryActiveUserStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryActiveUserStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUserStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryActiveUserStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryActiveUserStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryActiveUserStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryActiveUserStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryActiveUserStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryActiveUserStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryActiveUserStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryActiveUserStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryActiveUserStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryActiveUserStatisticalDataResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryActiveUserStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryActiveUserStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUserStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryActiveUserStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryActiveUserStatisticalDataResponse) SetBody(v *QueryActiveUserStatisticalDataResponseBody) *QueryActiveUserStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryGroupLiveStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupLiveStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupLiveStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupLiveStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupLiveStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupLiveStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupLiveStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryGroupLiveStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupLiveStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveStatisticalDataRequest) SetStatDate(v string) *QueryGroupLiveStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryGroupLiveStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryGroupLiveStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryGroupLiveStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupLiveStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryGroupLiveStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponseBody) SetMetaList(v []*QueryGroupLiveStatisticalDataResponseBodyMetaList) *QueryGroupLiveStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryGroupLiveStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryGroupLiveStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupLiveStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryGroupLiveStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryGroupLiveStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryGroupLiveStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryGroupLiveStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryGroupLiveStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryGroupLiveStatisticalDataResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupLiveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupLiveStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupLiveStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryGroupLiveStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupLiveStatisticalDataResponse) SetBody(v *QueryGroupLiveStatisticalDataResponseBody) *QueryGroupLiveStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryDigitalDistrictOrgInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDigitalDistrictOrgInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryDigitalDistrictOrgInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDigitalDistrictOrgInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDigitalDistrictOrgInfoRequest struct {
	KpiGroupId *string   `json:"kpiGroupId,omitempty" xml:"kpiGroupId,omitempty"`
	StatDates  []*string `json:"statDates,omitempty" xml:"statDates,omitempty" type:"Repeated"`
	CorpIds    []*string `json:"corpIds,omitempty" xml:"corpIds,omitempty" type:"Repeated"`
}

func (s QueryDigitalDistrictOrgInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoRequest) SetKpiGroupId(v string) *QueryDigitalDistrictOrgInfoRequest {
	s.KpiGroupId = &v
	return s
}

func (s *QueryDigitalDistrictOrgInfoRequest) SetStatDates(v []*string) *QueryDigitalDistrictOrgInfoRequest {
	s.StatDates = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoRequest) SetCorpIds(v []*string) *QueryDigitalDistrictOrgInfoRequest {
	s.CorpIds = v
	return s
}

type QueryDigitalDistrictOrgInfoResponseBody struct {
	// arguments
	Arguments []*string `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	// result
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s QueryDigitalDistrictOrgInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoResponseBody) SetArguments(v []*string) *QueryDigitalDistrictOrgInfoResponseBody {
	s.Arguments = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoResponseBody) SetResult(v string) *QueryDigitalDistrictOrgInfoResponseBody {
	s.Result = &v
	return s
}

type QueryDigitalDistrictOrgInfoResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDigitalDistrictOrgInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDigitalDistrictOrgInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalDistrictOrgInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDigitalDistrictOrgInfoResponse) SetHeaders(v map[string]*string) *QueryDigitalDistrictOrgInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDigitalDistrictOrgInfoResponse) SetBody(v *QueryDigitalDistrictOrgInfoResponseBody) *QueryDigitalDistrictOrgInfoResponse {
	s.Body = v
	return s
}

type QueryAttendanceStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAttendanceStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAttendanceStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryAttendanceStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryAttendanceStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAttendanceStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAttendanceStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAttendanceStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryAttendanceStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAttendanceStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryAttendanceStatisticalDataRequest) SetStatDate(v string) *QueryAttendanceStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryAttendanceStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryAttendanceStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryAttendanceStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAttendanceStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAttendanceStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryAttendanceStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryAttendanceStatisticalDataResponseBody) SetMetaList(v []*QueryAttendanceStatisticalDataResponseBodyMetaList) *QueryAttendanceStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryAttendanceStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryAttendanceStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryAttendanceStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryAttendanceStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryAttendanceStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryAttendanceStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryAttendanceStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryAttendanceStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryAttendanceStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryAttendanceStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryAttendanceStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryAttendanceStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryAttendanceStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryAttendanceStatisticalDataResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAttendanceStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAttendanceStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAttendanceStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryAttendanceStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryAttendanceStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryAttendanceStatisticalDataResponse) SetBody(v *QueryAttendanceStatisticalDataResponseBody) *QueryAttendanceStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryDingReciveStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDingReciveStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDingReciveStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryDingReciveStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryDingReciveStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDingReciveStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDingReciveStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDingReciveStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryDingReciveStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDingReciveStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDingReciveStatisticalDataRequest) SetStatDate(v string) *QueryDingReciveStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryDingReciveStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryDingReciveStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryDingReciveStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDingReciveStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDingReciveStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryDingReciveStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryDingReciveStatisticalDataResponseBody) SetMetaList(v []*QueryDingReciveStatisticalDataResponseBodyMetaList) *QueryDingReciveStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryDingReciveStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryDingReciveStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryDingReciveStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryDingReciveStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryDingReciveStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryDingReciveStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryDingReciveStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryDingReciveStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryDingReciveStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryDingReciveStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryDingReciveStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryDingReciveStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryDingReciveStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryDingReciveStatisticalDataResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDingReciveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDingReciveStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDingReciveStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDingReciveStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryDingReciveStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDingReciveStatisticalDataResponse) SetBody(v *QueryDingReciveStatisticalDataResponseBody) *QueryDingReciveStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryRedEnvelopeReciveStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRedEnvelopeReciveStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeReciveStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeReciveStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryRedEnvelopeReciveStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRedEnvelopeReciveStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRedEnvelopeReciveStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryRedEnvelopeReciveStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeReciveStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeReciveStatisticalDataRequest) SetStatDate(v string) *QueryRedEnvelopeReciveStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryRedEnvelopeReciveStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryRedEnvelopeReciveStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeReciveStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryRedEnvelopeReciveStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBody) SetMetaList(v []*QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) *QueryRedEnvelopeReciveStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryRedEnvelopeReciveStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryRedEnvelopeReciveStatisticalDataResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRedEnvelopeReciveStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRedEnvelopeReciveStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeReciveStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryRedEnvelopeReciveStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryRedEnvelopeReciveStatisticalDataResponse) SetBody(v *QueryRedEnvelopeReciveStatisticalDataResponseBody) *QueryRedEnvelopeReciveStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryCircleStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCircleStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCircleStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryCircleStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryCircleStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCircleStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCircleStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCircleStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryCircleStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCircleStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryCircleStatisticalDataRequest) SetStatDate(v string) *QueryCircleStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryCircleStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryCircleStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryCircleStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCircleStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCircleStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryCircleStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryCircleStatisticalDataResponseBody) SetMetaList(v []*QueryCircleStatisticalDataResponseBodyMetaList) *QueryCircleStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryCircleStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryCircleStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryCircleStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryCircleStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryCircleStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryCircleStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryCircleStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryCircleStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryCircleStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryCircleStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryCircleStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryCircleStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryCircleStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryCircleStatisticalDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCircleStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCircleStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCircleStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryCircleStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryCircleStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryCircleStatisticalDataResponse) SetBody(v *QueryCircleStatisticalDataResponseBody) *QueryCircleStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryRedEnvelopeSendStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRedEnvelopeSendStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeSendStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeSendStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryRedEnvelopeSendStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRedEnvelopeSendStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRedEnvelopeSendStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryRedEnvelopeSendStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeSendStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeSendStatisticalDataRequest) SetStatDate(v string) *QueryRedEnvelopeSendStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryRedEnvelopeSendStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryRedEnvelopeSendStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeSendStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryRedEnvelopeSendStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBody) SetMetaList(v []*QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) *QueryRedEnvelopeSendStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryRedEnvelopeSendStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryRedEnvelopeSendStatisticalDataResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRedEnvelopeSendStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRedEnvelopeSendStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRedEnvelopeSendStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryRedEnvelopeSendStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryRedEnvelopeSendStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryRedEnvelopeSendStatisticalDataResponse) SetBody(v *QueryRedEnvelopeSendStatisticalDataResponseBody) *QueryRedEnvelopeSendStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryTelMeetingStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTelMeetingStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTelMeetingStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryTelMeetingStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryTelMeetingStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTelMeetingStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTelMeetingStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTelMeetingStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryTelMeetingStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTelMeetingStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryTelMeetingStatisticalDataRequest) SetStatDate(v string) *QueryTelMeetingStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryTelMeetingStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryTelMeetingStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryTelMeetingStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTelMeetingStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTelMeetingStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryTelMeetingStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponseBody) SetMetaList(v []*QueryTelMeetingStatisticalDataResponseBodyMetaList) *QueryTelMeetingStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryTelMeetingStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryTelMeetingStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryTelMeetingStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryTelMeetingStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryTelMeetingStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryTelMeetingStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryTelMeetingStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryTelMeetingStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryTelMeetingStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryTelMeetingStatisticalDataResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTelMeetingStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTelMeetingStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTelMeetingStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryTelMeetingStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryTelMeetingStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryTelMeetingStatisticalDataResponse) SetBody(v *QueryTelMeetingStatisticalDataResponseBody) *QueryTelMeetingStatisticalDataResponse {
	s.Body = v
	return s
}

type QueryBlackboardStatisticalDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBlackboardStatisticalDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardStatisticalDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryBlackboardStatisticalDataHeaders) SetCommonHeaders(v map[string]*string) *QueryBlackboardStatisticalDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBlackboardStatisticalDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBlackboardStatisticalDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBlackboardStatisticalDataRequest struct {
	// statDate
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s QueryBlackboardStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *QueryBlackboardStatisticalDataRequest) SetStatDate(v string) *QueryBlackboardStatisticalDataRequest {
	s.StatDate = &v
	return s
}

type QueryBlackboardStatisticalDataResponseBody struct {
	// 指标数据
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// 指标元数据
	MetaList []*QueryBlackboardStatisticalDataResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s QueryBlackboardStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlackboardStatisticalDataResponseBody) SetDataList(v []map[string]interface{}) *QueryBlackboardStatisticalDataResponseBody {
	s.DataList = v
	return s
}

func (s *QueryBlackboardStatisticalDataResponseBody) SetMetaList(v []*QueryBlackboardStatisticalDataResponseBodyMetaList) *QueryBlackboardStatisticalDataResponseBody {
	s.MetaList = v
	return s
}

type QueryBlackboardStatisticalDataResponseBodyMetaList struct {
	// 指标ID
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// 指标名称
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// 指标单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 指标口径
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// 指标周期
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryBlackboardStatisticalDataResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardStatisticalDataResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *QueryBlackboardStatisticalDataResponseBodyMetaList) SetKpiId(v string) *QueryBlackboardStatisticalDataResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *QueryBlackboardStatisticalDataResponseBodyMetaList) SetKpiName(v string) *QueryBlackboardStatisticalDataResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *QueryBlackboardStatisticalDataResponseBodyMetaList) SetUnit(v string) *QueryBlackboardStatisticalDataResponseBodyMetaList {
	s.Unit = &v
	return s
}

func (s *QueryBlackboardStatisticalDataResponseBodyMetaList) SetKpiCaliber(v string) *QueryBlackboardStatisticalDataResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *QueryBlackboardStatisticalDataResponseBodyMetaList) SetPeriod(v string) *QueryBlackboardStatisticalDataResponseBodyMetaList {
	s.Period = &v
	return s
}

type QueryBlackboardStatisticalDataResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBlackboardStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBlackboardStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *QueryBlackboardStatisticalDataResponse) SetHeaders(v map[string]*string) *QueryBlackboardStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *QueryBlackboardStatisticalDataResponse) SetBody(v *QueryBlackboardStatisticalDataResponseBody) *QueryBlackboardStatisticalDataResponse {
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

func (client *Client) QueryGroupMessageStatisticalData(request *QueryGroupMessageStatisticalDataRequest) (_result *QueryGroupMessageStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupMessageStatisticalDataHeaders{}
	_result = &QueryGroupMessageStatisticalDataResponse{}
	_body, _err := client.QueryGroupMessageStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupMessageStatisticalDataWithOptions(request *QueryGroupMessageStatisticalDataRequest, headers *QueryGroupMessageStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupMessageStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryGroupMessageStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupMessageStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/groupMessageData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVedioMeetingStatisticalData(request *QueryVedioMeetingStatisticalDataRequest) (_result *QueryVedioMeetingStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryVedioMeetingStatisticalDataHeaders{}
	_result = &QueryVedioMeetingStatisticalDataResponse{}
	_body, _err := client.QueryVedioMeetingStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryVedioMeetingStatisticalDataWithOptions(request *QueryVedioMeetingStatisticalDataRequest, headers *QueryVedioMeetingStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryVedioMeetingStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryVedioMeetingStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryVedioMeetingStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/vedioMeetingData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHealthStatisticalData(request *QueryHealthStatisticalDataRequest) (_result *QueryHealthStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHealthStatisticalDataHeaders{}
	_result = &QueryHealthStatisticalDataResponse{}
	_body, _err := client.QueryHealthStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHealthStatisticalDataWithOptions(request *QueryHealthStatisticalDataRequest, headers *QueryHealthStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryHealthStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryHealthStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryHealthStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/healtheUserData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySingleMessageStatisticalData(request *QuerySingleMessageStatisticalDataRequest) (_result *QuerySingleMessageStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySingleMessageStatisticalDataHeaders{}
	_result = &QuerySingleMessageStatisticalDataResponse{}
	_body, _err := client.QuerySingleMessageStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySingleMessageStatisticalDataWithOptions(request *QuerySingleMessageStatisticalDataRequest, headers *QuerySingleMessageStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QuerySingleMessageStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QuerySingleMessageStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySingleMessageStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/singleMessagerData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTodoStatisticalData(request *QueryTodoStatisticalDataRequest) (_result *QueryTodoStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTodoStatisticalDataHeaders{}
	_result = &QueryTodoStatisticalDataResponse{}
	_body, _err := client.QueryTodoStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTodoStatisticalDataWithOptions(request *QueryTodoStatisticalDataRequest, headers *QueryTodoStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryTodoStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryTodoStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryTodoStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/todoUserData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCheckinStatisticalData(request *QueryCheckinStatisticalDataRequest) (_result *QueryCheckinStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCheckinStatisticalDataHeaders{}
	_result = &QueryCheckinStatisticalDataResponse{}
	_body, _err := client.QueryCheckinStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCheckinStatisticalDataWithOptions(request *QueryCheckinStatisticalDataRequest, headers *QueryCheckinStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryCheckinStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryCheckinStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCheckinStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/checkinData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEmployeeTypeStatisticalData(request *QueryEmployeeTypeStatisticalDataRequest) (_result *QueryEmployeeTypeStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryEmployeeTypeStatisticalDataHeaders{}
	_result = &QueryEmployeeTypeStatisticalDataResponse{}
	_body, _err := client.QueryEmployeeTypeStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEmployeeTypeStatisticalDataWithOptions(request *QueryEmployeeTypeStatisticalDataRequest, headers *QueryEmployeeTypeStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryEmployeeTypeStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryEmployeeTypeStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryEmployeeTypeStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/employeeTypeData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMailStatisticalData(request *QueryMailStatisticalDataRequest) (_result *QueryMailStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMailStatisticalDataHeaders{}
	_result = &QueryMailStatisticalDataResponse{}
	_body, _err := client.QueryMailStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMailStatisticalDataWithOptions(request *QueryMailStatisticalDataRequest, headers *QueryMailStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryMailStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryMailStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryMailStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/mailData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCalendarStatisticalData(request *QueryCalendarStatisticalDataRequest) (_result *QueryCalendarStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCalendarStatisticalDataHeaders{}
	_result = &QueryCalendarStatisticalDataResponse{}
	_body, _err := client.QueryCalendarStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCalendarStatisticalDataWithOptions(request *QueryCalendarStatisticalDataRequest, headers *QueryCalendarStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryCalendarStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryCalendarStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCalendarStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/calendarData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDocumentStatisticalData(request *QueryDocumentStatisticalDataRequest) (_result *QueryDocumentStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDocumentStatisticalDataHeaders{}
	_result = &QueryDocumentStatisticalDataResponse{}
	_body, _err := client.QueryDocumentStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDocumentStatisticalDataWithOptions(request *QueryDocumentStatisticalDataRequest, headers *QueryDocumentStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryDocumentStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryDocumentStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDocumentStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/documentData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryReportStatisticalData(request *QueryReportStatisticalDataRequest) (_result *QueryReportStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryReportStatisticalDataHeaders{}
	_result = &QueryReportStatisticalDataResponse{}
	_body, _err := client.QueryReportStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryReportStatisticalDataWithOptions(request *QueryReportStatisticalDataRequest, headers *QueryReportStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryReportStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryReportStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryReportStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/reportData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOnlineUserStatisticalData(request *QueryOnlineUserStatisticalDataRequest) (_result *QueryOnlineUserStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOnlineUserStatisticalDataHeaders{}
	_result = &QueryOnlineUserStatisticalDataResponse{}
	_body, _err := client.QueryOnlineUserStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOnlineUserStatisticalDataWithOptions(request *QueryOnlineUserStatisticalDataRequest, headers *QueryOnlineUserStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryOnlineUserStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryOnlineUserStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOnlineUserStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/onlineUserData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCompanyBasicInfo(request *QueryCompanyBasicInfoRequest) (_result *QueryCompanyBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCompanyBasicInfoHeaders{}
	_result = &QueryCompanyBasicInfoResponse{}
	_body, _err := client.QueryCompanyBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCompanyBasicInfoWithOptions(request *QueryCompanyBasicInfoRequest, headers *QueryCompanyBasicInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryCompanyBasicInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryCompanyBasicInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCompanyBasicInfo"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/companies/basicInfo"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryApprovalStatisticalData(request *QueryApprovalStatisticalDataRequest) (_result *QueryApprovalStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryApprovalStatisticalDataHeaders{}
	_result = &QueryApprovalStatisticalDataResponse{}
	_body, _err := client.QueryApprovalStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryApprovalStatisticalDataWithOptions(request *QueryApprovalStatisticalDataRequest, headers *QueryApprovalStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryApprovalStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryApprovalStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryApprovalStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/approvalData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDriveStatisticalData(request *QueryDriveStatisticalDataRequest) (_result *QueryDriveStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDriveStatisticalDataHeaders{}
	_result = &QueryDriveStatisticalDataResponse{}
	_body, _err := client.QueryDriveStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDriveStatisticalDataWithOptions(request *QueryDriveStatisticalDataRequest, headers *QueryDriveStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryDriveStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryDriveStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDriveStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/driveData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDingSendStatisticalData(request *QueryDingSendStatisticalDataRequest) (_result *QueryDingSendStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDingSendStatisticalDataHeaders{}
	_result = &QueryDingSendStatisticalDataResponse{}
	_body, _err := client.QueryDingSendStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDingSendStatisticalDataWithOptions(request *QueryDingSendStatisticalDataRequest, headers *QueryDingSendStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryDingSendStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryDingSendStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDingSendStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/dingSendData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryActiveUserStatisticalData(request *QueryActiveUserStatisticalDataRequest) (_result *QueryActiveUserStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryActiveUserStatisticalDataHeaders{}
	_result = &QueryActiveUserStatisticalDataResponse{}
	_body, _err := client.QueryActiveUserStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryActiveUserStatisticalDataWithOptions(request *QueryActiveUserStatisticalDataRequest, headers *QueryActiveUserStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryActiveUserStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryActiveUserStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryActiveUserStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/activeUserData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupLiveStatisticalData(request *QueryGroupLiveStatisticalDataRequest) (_result *QueryGroupLiveStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupLiveStatisticalDataHeaders{}
	_result = &QueryGroupLiveStatisticalDataResponse{}
	_body, _err := client.QueryGroupLiveStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupLiveStatisticalDataWithOptions(request *QueryGroupLiveStatisticalDataRequest, headers *QueryGroupLiveStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupLiveStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryGroupLiveStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupLiveStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/groupLiveData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDigitalDistrictOrgInfo(request *QueryDigitalDistrictOrgInfoRequest) (_result *QueryDigitalDistrictOrgInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDigitalDistrictOrgInfoHeaders{}
	_result = &QueryDigitalDistrictOrgInfoResponse{}
	_body, _err := client.QueryDigitalDistrictOrgInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDigitalDistrictOrgInfoWithOptions(request *QueryDigitalDistrictOrgInfoRequest, headers *QueryDigitalDistrictOrgInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryDigitalDistrictOrgInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KpiGroupId)) {
		body["kpiGroupId"] = request.KpiGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StatDates)) {
		body["statDates"] = request.StatDates
	}

	if !tea.BoolValue(util.IsUnset(request.CorpIds)) {
		body["corpIds"] = request.CorpIds
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &QueryDigitalDistrictOrgInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDigitalDistrictOrgInfo"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/datacenter/digitalCounty/orgInfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAttendanceStatisticalData(request *QueryAttendanceStatisticalDataRequest) (_result *QueryAttendanceStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAttendanceStatisticalDataHeaders{}
	_result = &QueryAttendanceStatisticalDataResponse{}
	_body, _err := client.QueryAttendanceStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAttendanceStatisticalDataWithOptions(request *QueryAttendanceStatisticalDataRequest, headers *QueryAttendanceStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryAttendanceStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryAttendanceStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAttendanceStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/attendanceData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDingReciveStatisticalData(request *QueryDingReciveStatisticalDataRequest) (_result *QueryDingReciveStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDingReciveStatisticalDataHeaders{}
	_result = &QueryDingReciveStatisticalDataResponse{}
	_body, _err := client.QueryDingReciveStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDingReciveStatisticalDataWithOptions(request *QueryDingReciveStatisticalDataRequest, headers *QueryDingReciveStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryDingReciveStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryDingReciveStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDingReciveStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/dingReciveData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRedEnvelopeReciveStatisticalData(request *QueryRedEnvelopeReciveStatisticalDataRequest) (_result *QueryRedEnvelopeReciveStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRedEnvelopeReciveStatisticalDataHeaders{}
	_result = &QueryRedEnvelopeReciveStatisticalDataResponse{}
	_body, _err := client.QueryRedEnvelopeReciveStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRedEnvelopeReciveStatisticalDataWithOptions(request *QueryRedEnvelopeReciveStatisticalDataRequest, headers *QueryRedEnvelopeReciveStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryRedEnvelopeReciveStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryRedEnvelopeReciveStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryRedEnvelopeReciveStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/redEnvelopeReciveData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCircleStatisticalData(request *QueryCircleStatisticalDataRequest) (_result *QueryCircleStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCircleStatisticalDataHeaders{}
	_result = &QueryCircleStatisticalDataResponse{}
	_body, _err := client.QueryCircleStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCircleStatisticalDataWithOptions(request *QueryCircleStatisticalDataRequest, headers *QueryCircleStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryCircleStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryCircleStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCircleStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/circleData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRedEnvelopeSendStatisticalData(request *QueryRedEnvelopeSendStatisticalDataRequest) (_result *QueryRedEnvelopeSendStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRedEnvelopeSendStatisticalDataHeaders{}
	_result = &QueryRedEnvelopeSendStatisticalDataResponse{}
	_body, _err := client.QueryRedEnvelopeSendStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRedEnvelopeSendStatisticalDataWithOptions(request *QueryRedEnvelopeSendStatisticalDataRequest, headers *QueryRedEnvelopeSendStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryRedEnvelopeSendStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryRedEnvelopeSendStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryRedEnvelopeSendStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/redEnvelopeSendData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTelMeetingStatisticalData(request *QueryTelMeetingStatisticalDataRequest) (_result *QueryTelMeetingStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTelMeetingStatisticalDataHeaders{}
	_result = &QueryTelMeetingStatisticalDataResponse{}
	_body, _err := client.QueryTelMeetingStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTelMeetingStatisticalDataWithOptions(request *QueryTelMeetingStatisticalDataRequest, headers *QueryTelMeetingStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryTelMeetingStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryTelMeetingStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryTelMeetingStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/telMeetingData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBlackboardStatisticalData(request *QueryBlackboardStatisticalDataRequest) (_result *QueryBlackboardStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBlackboardStatisticalDataHeaders{}
	_result = &QueryBlackboardStatisticalDataResponse{}
	_body, _err := client.QueryBlackboardStatisticalDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBlackboardStatisticalDataWithOptions(request *QueryBlackboardStatisticalDataRequest, headers *QueryBlackboardStatisticalDataHeaders, runtime *util.RuntimeOptions) (_result *QueryBlackboardStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryBlackboardStatisticalDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryBlackboardStatisticalData"), tea.String("datacenter_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/datacenter/blackboardData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
