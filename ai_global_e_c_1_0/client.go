// This file is auto-generated, don't edit it. Thanks.
package ai_global_e_c_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConnectionOmniChannelTiktokMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConnectionOmniChannelTiktokMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConnectionOmniChannelTiktokMessageHeaders) GoString() string {
	return s.String()
}

func (s *ConnectionOmniChannelTiktokMessageHeaders) SetCommonHeaders(v map[string]*string) *ConnectionOmniChannelTiktokMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConnectionOmniChannelTiktokMessageHeaders) SetXAcsDingtalkAccessToken(v string) *ConnectionOmniChannelTiktokMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConnectionOmniChannelTiktokMessageRequest struct {
	TiktokContentJsonString *string `json:"tiktokContentJsonString,omitempty" xml:"tiktokContentJsonString,omitempty"`
}

func (s ConnectionOmniChannelTiktokMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ConnectionOmniChannelTiktokMessageRequest) GoString() string {
	return s.String()
}

func (s *ConnectionOmniChannelTiktokMessageRequest) SetTiktokContentJsonString(v string) *ConnectionOmniChannelTiktokMessageRequest {
	s.TiktokContentJsonString = &v
	return s
}

type ConnectionOmniChannelTiktokMessageResponseBody struct {
	ErrorCode            *string                                                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg             *string                                                             `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	OmniChannelTiktokRsp *ConnectionOmniChannelTiktokMessageResponseBodyOmniChannelTiktokRsp `json:"omniChannelTiktokRsp,omitempty" xml:"omniChannelTiktokRsp,omitempty" type:"Struct"`
	Success              *string                                                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConnectionOmniChannelTiktokMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConnectionOmniChannelTiktokMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ConnectionOmniChannelTiktokMessageResponseBody) SetErrorCode(v string) *ConnectionOmniChannelTiktokMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ConnectionOmniChannelTiktokMessageResponseBody) SetErrorMsg(v string) *ConnectionOmniChannelTiktokMessageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ConnectionOmniChannelTiktokMessageResponseBody) SetOmniChannelTiktokRsp(v *ConnectionOmniChannelTiktokMessageResponseBodyOmniChannelTiktokRsp) *ConnectionOmniChannelTiktokMessageResponseBody {
	s.OmniChannelTiktokRsp = v
	return s
}

func (s *ConnectionOmniChannelTiktokMessageResponseBody) SetSuccess(v string) *ConnectionOmniChannelTiktokMessageResponseBody {
	s.Success = &v
	return s
}

type ConnectionOmniChannelTiktokMessageResponseBodyOmniChannelTiktokRsp struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s ConnectionOmniChannelTiktokMessageResponseBodyOmniChannelTiktokRsp) String() string {
	return tea.Prettify(s)
}

func (s ConnectionOmniChannelTiktokMessageResponseBodyOmniChannelTiktokRsp) GoString() string {
	return s.String()
}

func (s *ConnectionOmniChannelTiktokMessageResponseBodyOmniChannelTiktokRsp) SetCode(v string) *ConnectionOmniChannelTiktokMessageResponseBodyOmniChannelTiktokRsp {
	s.Code = &v
	return s
}

type ConnectionOmniChannelTiktokMessageResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConnectionOmniChannelTiktokMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConnectionOmniChannelTiktokMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ConnectionOmniChannelTiktokMessageResponse) GoString() string {
	return s.String()
}

func (s *ConnectionOmniChannelTiktokMessageResponse) SetHeaders(v map[string]*string) *ConnectionOmniChannelTiktokMessageResponse {
	s.Headers = v
	return s
}

func (s *ConnectionOmniChannelTiktokMessageResponse) SetStatusCode(v int32) *ConnectionOmniChannelTiktokMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ConnectionOmniChannelTiktokMessageResponse) SetBody(v *ConnectionOmniChannelTiktokMessageResponseBody) *ConnectionOmniChannelTiktokMessageResponse {
	s.Body = v
	return s
}

type GetLoginUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetLoginUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLoginUserHeaders) GoString() string {
	return s.String()
}

func (s *GetLoginUserHeaders) SetCommonHeaders(v map[string]*string) *GetLoginUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLoginUserHeaders) SetXAcsDingtalkAccessToken(v string) *GetLoginUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetLoginUserRequest struct {
	AuthCode *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
}

func (s GetLoginUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginUserRequest) GoString() string {
	return s.String()
}

func (s *GetLoginUserRequest) SetAuthCode(v string) *GetLoginUserRequest {
	s.AuthCode = &v
	return s
}

type GetLoginUserResponseBody struct {
	CommodityInfo *GetLoginUserResponseBodyCommodityInfo `json:"commodityInfo,omitempty" xml:"commodityInfo,omitempty" type:"Struct"`
	CorpId        *string                                `json:"corpId,omitempty" xml:"corpId,omitempty"`
	OpenId        *string                                `json:"openId,omitempty" xml:"openId,omitempty"`
	UnionId       *string                                `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetLoginUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginUserResponseBody) SetCommodityInfo(v *GetLoginUserResponseBodyCommodityInfo) *GetLoginUserResponseBody {
	s.CommodityInfo = v
	return s
}

func (s *GetLoginUserResponseBody) SetCorpId(v string) *GetLoginUserResponseBody {
	s.CorpId = &v
	return s
}

func (s *GetLoginUserResponseBody) SetOpenId(v string) *GetLoginUserResponseBody {
	s.OpenId = &v
	return s
}

func (s *GetLoginUserResponseBody) SetUnionId(v string) *GetLoginUserResponseBody {
	s.UnionId = &v
	return s
}

type GetLoginUserResponseBodyCommodityInfo struct {
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetLoginUserResponseBodyCommodityInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLoginUserResponseBodyCommodityInfo) GoString() string {
	return s.String()
}

func (s *GetLoginUserResponseBodyCommodityInfo) SetVersion(v string) *GetLoginUserResponseBodyCommodityInfo {
	s.Version = &v
	return s
}

type GetLoginUserResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginUserResponse) GoString() string {
	return s.String()
}

func (s *GetLoginUserResponse) SetHeaders(v map[string]*string) *GetLoginUserResponse {
	s.Headers = v
	return s
}

func (s *GetLoginUserResponse) SetStatusCode(v int32) *GetLoginUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginUserResponse) SetBody(v *GetLoginUserResponseBody) *GetLoginUserResponse {
	s.Body = v
	return s
}

type HhoCallBackHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HhoCallBackHeaders) String() string {
	return tea.Prettify(s)
}

func (s HhoCallBackHeaders) GoString() string {
	return s.String()
}

func (s *HhoCallBackHeaders) SetCommonHeaders(v map[string]*string) *HhoCallBackHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HhoCallBackHeaders) SetXAcsDingtalkAccessToken(v string) *HhoCallBackHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HhoCallBackRequest struct {
	Data             *string `json:"data,omitempty" xml:"data,omitempty"`
	DtNotificationId *string `json:"dtNotificationId,omitempty" xml:"dtNotificationId,omitempty"`
	ShopId           *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	Timestamp        *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Type             *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HhoCallBackRequest) String() string {
	return tea.Prettify(s)
}

func (s HhoCallBackRequest) GoString() string {
	return s.String()
}

func (s *HhoCallBackRequest) SetData(v string) *HhoCallBackRequest {
	s.Data = &v
	return s
}

func (s *HhoCallBackRequest) SetDtNotificationId(v string) *HhoCallBackRequest {
	s.DtNotificationId = &v
	return s
}

func (s *HhoCallBackRequest) SetShopId(v string) *HhoCallBackRequest {
	s.ShopId = &v
	return s
}

func (s *HhoCallBackRequest) SetTimestamp(v string) *HhoCallBackRequest {
	s.Timestamp = &v
	return s
}

func (s *HhoCallBackRequest) SetType(v int32) *HhoCallBackRequest {
	s.Type = &v
	return s
}

type HhoCallBackResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HhoCallBackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HhoCallBackResponseBody) GoString() string {
	return s.String()
}

func (s *HhoCallBackResponseBody) SetSuccess(v bool) *HhoCallBackResponseBody {
	s.Success = &v
	return s
}

type HhoCallBackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HhoCallBackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HhoCallBackResponse) String() string {
	return tea.Prettify(s)
}

func (s HhoCallBackResponse) GoString() string {
	return s.String()
}

func (s *HhoCallBackResponse) SetHeaders(v map[string]*string) *HhoCallBackResponse {
	s.Headers = v
	return s
}

func (s *HhoCallBackResponse) SetStatusCode(v int32) *HhoCallBackResponse {
	s.StatusCode = &v
	return s
}

func (s *HhoCallBackResponse) SetBody(v *HhoCallBackResponseBody) *HhoCallBackResponse {
	s.Body = v
	return s
}

type LaunchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LaunchHeaders) String() string {
	return tea.Prettify(s)
}

func (s LaunchHeaders) GoString() string {
	return s.String()
}

func (s *LaunchHeaders) SetCommonHeaders(v map[string]*string) *LaunchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LaunchHeaders) SetXAcsDingtalkAccessToken(v string) *LaunchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LaunchRequest struct {
	Description   *string                  `json:"description,omitempty" xml:"description,omitempty"`
	ImageUrls     []*string                `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
	Platform      []*string                `json:"platform,omitempty" xml:"platform,omitempty" type:"Repeated"`
	ProductName   *string                  `json:"productName,omitempty" xml:"productName,omitempty"`
	SellingPoints []*string                `json:"sellingPoints,omitempty" xml:"sellingPoints,omitempty" type:"Repeated"`
	SourceData    *string                  `json:"sourceData,omitempty" xml:"sourceData,omitempty"`
	Variants      []*LaunchRequestVariants `json:"variants,omitempty" xml:"variants,omitempty" type:"Repeated"`
	VideoUrls     []*string                `json:"videoUrls,omitempty" xml:"videoUrls,omitempty" type:"Repeated"`
}

func (s LaunchRequest) String() string {
	return tea.Prettify(s)
}

func (s LaunchRequest) GoString() string {
	return s.String()
}

func (s *LaunchRequest) SetDescription(v string) *LaunchRequest {
	s.Description = &v
	return s
}

func (s *LaunchRequest) SetImageUrls(v []*string) *LaunchRequest {
	s.ImageUrls = v
	return s
}

func (s *LaunchRequest) SetPlatform(v []*string) *LaunchRequest {
	s.Platform = v
	return s
}

func (s *LaunchRequest) SetProductName(v string) *LaunchRequest {
	s.ProductName = &v
	return s
}

func (s *LaunchRequest) SetSellingPoints(v []*string) *LaunchRequest {
	s.SellingPoints = v
	return s
}

func (s *LaunchRequest) SetSourceData(v string) *LaunchRequest {
	s.SourceData = &v
	return s
}

func (s *LaunchRequest) SetVariants(v []*LaunchRequestVariants) *LaunchRequest {
	s.Variants = v
	return s
}

func (s *LaunchRequest) SetVideoUrls(v []*string) *LaunchRequest {
	s.VideoUrls = v
	return s
}

type LaunchRequestVariants struct {
	Images       []*string                            `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	OptionValues []*LaunchRequestVariantsOptionValues `json:"optionValues,omitempty" xml:"optionValues,omitempty" type:"Repeated"`
	Price        *string                              `json:"price,omitempty" xml:"price,omitempty"`
	Sku          *string                              `json:"sku,omitempty" xml:"sku,omitempty"`
}

func (s LaunchRequestVariants) String() string {
	return tea.Prettify(s)
}

func (s LaunchRequestVariants) GoString() string {
	return s.String()
}

func (s *LaunchRequestVariants) SetImages(v []*string) *LaunchRequestVariants {
	s.Images = v
	return s
}

func (s *LaunchRequestVariants) SetOptionValues(v []*LaunchRequestVariantsOptionValues) *LaunchRequestVariants {
	s.OptionValues = v
	return s
}

func (s *LaunchRequestVariants) SetPrice(v string) *LaunchRequestVariants {
	s.Price = &v
	return s
}

func (s *LaunchRequestVariants) SetSku(v string) *LaunchRequestVariants {
	s.Sku = &v
	return s
}

type LaunchRequestVariantsOptionValues struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s LaunchRequestVariantsOptionValues) String() string {
	return tea.Prettify(s)
}

func (s LaunchRequestVariantsOptionValues) GoString() string {
	return s.String()
}

func (s *LaunchRequestVariantsOptionValues) SetName(v string) *LaunchRequestVariantsOptionValues {
	s.Name = &v
	return s
}

func (s *LaunchRequestVariantsOptionValues) SetValue(v string) *LaunchRequestVariantsOptionValues {
	s.Value = &v
	return s
}

type LaunchResponseBody struct {
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LaunchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LaunchResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchResponseBody) SetSuccess(v bool) *LaunchResponseBody {
	s.Success = &v
	return s
}

type LaunchResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LaunchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LaunchResponse) String() string {
	return tea.Prettify(s)
}

func (s LaunchResponse) GoString() string {
	return s.String()
}

func (s *LaunchResponse) SetHeaders(v map[string]*string) *LaunchResponse {
	s.Headers = v
	return s
}

func (s *LaunchResponse) SetStatusCode(v int32) *LaunchResponse {
	s.StatusCode = &v
	return s
}

func (s *LaunchResponse) SetBody(v *LaunchResponseBody) *LaunchResponse {
	s.Body = v
	return s
}

type QueryNotableInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryNotableInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryNotableInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryNotableInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryNotableInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryNotableInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryNotableInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryNotableInfoRequest struct {
	SceneCode *string `json:"sceneCode,omitempty" xml:"sceneCode,omitempty"`
}

func (s QueryNotableInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryNotableInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryNotableInfoRequest) SetSceneCode(v string) *QueryNotableInfoRequest {
	s.SceneCode = &v
	return s
}

type QueryNotableInfoResponseBody struct {
	AdminUnionIds []*string `json:"adminUnionIds,omitempty" xml:"adminUnionIds,omitempty" type:"Repeated"`
	BaseId        *string   `json:"baseId,omitempty" xml:"baseId,omitempty"`
}

func (s QueryNotableInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryNotableInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNotableInfoResponseBody) SetAdminUnionIds(v []*string) *QueryNotableInfoResponseBody {
	s.AdminUnionIds = v
	return s
}

func (s *QueryNotableInfoResponseBody) SetBaseId(v string) *QueryNotableInfoResponseBody {
	s.BaseId = &v
	return s
}

type QueryNotableInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryNotableInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryNotableInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryNotableInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryNotableInfoResponse) SetHeaders(v map[string]*string) *QueryNotableInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryNotableInfoResponse) SetStatusCode(v int32) *QueryNotableInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNotableInfoResponse) SetBody(v *QueryNotableInfoResponseBody) *QueryNotableInfoResponse {
	s.Body = v
	return s
}

type TiktokShopAuthCallbackHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TiktokShopAuthCallbackHeaders) String() string {
	return tea.Prettify(s)
}

func (s TiktokShopAuthCallbackHeaders) GoString() string {
	return s.String()
}

func (s *TiktokShopAuthCallbackHeaders) SetCommonHeaders(v map[string]*string) *TiktokShopAuthCallbackHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TiktokShopAuthCallbackHeaders) SetXAcsDingtalkAccessToken(v string) *TiktokShopAuthCallbackHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TiktokShopAuthCallbackRequest struct {
	SellerType *string `json:"sellerType,omitempty" xml:"sellerType,omitempty"`
	ShopId     *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	ShopName   *string `json:"shopName,omitempty" xml:"shopName,omitempty"`
	ShopRegion *string `json:"shopRegion,omitempty" xml:"shopRegion,omitempty"`
}

func (s TiktokShopAuthCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s TiktokShopAuthCallbackRequest) GoString() string {
	return s.String()
}

func (s *TiktokShopAuthCallbackRequest) SetSellerType(v string) *TiktokShopAuthCallbackRequest {
	s.SellerType = &v
	return s
}

func (s *TiktokShopAuthCallbackRequest) SetShopId(v string) *TiktokShopAuthCallbackRequest {
	s.ShopId = &v
	return s
}

func (s *TiktokShopAuthCallbackRequest) SetShopName(v string) *TiktokShopAuthCallbackRequest {
	s.ShopName = &v
	return s
}

func (s *TiktokShopAuthCallbackRequest) SetShopRegion(v string) *TiktokShopAuthCallbackRequest {
	s.ShopRegion = &v
	return s
}

type TiktokShopAuthCallbackResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TiktokShopAuthCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TiktokShopAuthCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *TiktokShopAuthCallbackResponseBody) SetSuccess(v bool) *TiktokShopAuthCallbackResponseBody {
	s.Success = &v
	return s
}

type TiktokShopAuthCallbackResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TiktokShopAuthCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TiktokShopAuthCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s TiktokShopAuthCallbackResponse) GoString() string {
	return s.String()
}

func (s *TiktokShopAuthCallbackResponse) SetHeaders(v map[string]*string) *TiktokShopAuthCallbackResponse {
	s.Headers = v
	return s
}

func (s *TiktokShopAuthCallbackResponse) SetStatusCode(v int32) *TiktokShopAuthCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *TiktokShopAuthCallbackResponse) SetBody(v *TiktokShopAuthCallbackResponseBody) *TiktokShopAuthCallbackResponse {
	s.Body = v
	return s
}

type TiktokWebhookProcessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TiktokWebhookProcessHeaders) String() string {
	return tea.Prettify(s)
}

func (s TiktokWebhookProcessHeaders) GoString() string {
	return s.String()
}

func (s *TiktokWebhookProcessHeaders) SetCommonHeaders(v map[string]*string) *TiktokWebhookProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TiktokWebhookProcessHeaders) SetXAcsDingtalkAccessToken(v string) *TiktokWebhookProcessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TiktokWebhookProcessRequest struct {
	TiktokContentJsonString *string `json:"tiktokContentJsonString,omitempty" xml:"tiktokContentJsonString,omitempty"`
}

func (s TiktokWebhookProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s TiktokWebhookProcessRequest) GoString() string {
	return s.String()
}

func (s *TiktokWebhookProcessRequest) SetTiktokContentJsonString(v string) *TiktokWebhookProcessRequest {
	s.TiktokContentJsonString = &v
	return s
}

type TiktokWebhookProcessResponseBody struct {
	Code                        *int32                                                       `json:"code,omitempty" xml:"code,omitempty"`
	ErrorCode                   *string                                                      `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg                    *string                                                      `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	OmniChannelTiktokWebhookRsp *TiktokWebhookProcessResponseBodyOmniChannelTiktokWebhookRsp `json:"omniChannelTiktokWebhookRsp,omitempty" xml:"omniChannelTiktokWebhookRsp,omitempty" type:"Struct"`
	Success                     *string                                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TiktokWebhookProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TiktokWebhookProcessResponseBody) GoString() string {
	return s.String()
}

func (s *TiktokWebhookProcessResponseBody) SetCode(v int32) *TiktokWebhookProcessResponseBody {
	s.Code = &v
	return s
}

func (s *TiktokWebhookProcessResponseBody) SetErrorCode(v string) *TiktokWebhookProcessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TiktokWebhookProcessResponseBody) SetErrorMsg(v string) *TiktokWebhookProcessResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TiktokWebhookProcessResponseBody) SetOmniChannelTiktokWebhookRsp(v *TiktokWebhookProcessResponseBodyOmniChannelTiktokWebhookRsp) *TiktokWebhookProcessResponseBody {
	s.OmniChannelTiktokWebhookRsp = v
	return s
}

func (s *TiktokWebhookProcessResponseBody) SetSuccess(v string) *TiktokWebhookProcessResponseBody {
	s.Success = &v
	return s
}

type TiktokWebhookProcessResponseBodyOmniChannelTiktokWebhookRsp struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s TiktokWebhookProcessResponseBodyOmniChannelTiktokWebhookRsp) String() string {
	return tea.Prettify(s)
}

func (s TiktokWebhookProcessResponseBodyOmniChannelTiktokWebhookRsp) GoString() string {
	return s.String()
}

func (s *TiktokWebhookProcessResponseBodyOmniChannelTiktokWebhookRsp) SetCode(v string) *TiktokWebhookProcessResponseBodyOmniChannelTiktokWebhookRsp {
	s.Code = &v
	return s
}

type TiktokWebhookProcessResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TiktokWebhookProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TiktokWebhookProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s TiktokWebhookProcessResponse) GoString() string {
	return s.String()
}

func (s *TiktokWebhookProcessResponse) SetHeaders(v map[string]*string) *TiktokWebhookProcessResponse {
	s.Headers = v
	return s
}

func (s *TiktokWebhookProcessResponse) SetStatusCode(v int32) *TiktokWebhookProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *TiktokWebhookProcessResponse) SetBody(v *TiktokWebhookProcessResponseBody) *TiktokWebhookProcessResponse {
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
// 全渠道运营客服tiktok消息接入
//
// @param request - ConnectionOmniChannelTiktokMessageRequest
//
// @param headers - ConnectionOmniChannelTiktokMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConnectionOmniChannelTiktokMessageResponse
func (client *Client) ConnectionOmniChannelTiktokMessageWithOptions(request *ConnectionOmniChannelTiktokMessageRequest, headers *ConnectionOmniChannelTiktokMessageHeaders, runtime *util.RuntimeOptions) (_result *ConnectionOmniChannelTiktokMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TiktokContentJsonString)) {
		body["tiktokContentJsonString"] = request.TiktokContentJsonString
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
		Action:      tea.String("ConnectionOmniChannelTiktokMessage"),
		Version:     tea.String("aiGlobalEC_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiGlobalEC/omniChannel/tiktok/im/message/process"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConnectionOmniChannelTiktokMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 全渠道运营客服tiktok消息接入
//
// @param request - ConnectionOmniChannelTiktokMessageRequest
//
// @return ConnectionOmniChannelTiktokMessageResponse
func (client *Client) ConnectionOmniChannelTiktokMessage(request *ConnectionOmniChannelTiktokMessageRequest) (_result *ConnectionOmniChannelTiktokMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConnectionOmniChannelTiktokMessageHeaders{}
	_result = &ConnectionOmniChannelTiktokMessageResponse{}
	_body, _err := client.ConnectionOmniChannelTiktokMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前登录用户版本信息
//
// @param request - GetLoginUserRequest
//
// @param headers - GetLoginUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginUserResponse
func (client *Client) GetLoginUserWithOptions(request *GetLoginUserRequest, headers *GetLoginUserHeaders, runtime *util.RuntimeOptions) (_result *GetLoginUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		body["authCode"] = request.AuthCode
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
		Action:      tea.String("GetLoginUser"),
		Version:     tea.String("aiGlobalEC_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiGlobalEC/authCode/getLoginUser"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoginUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前登录用户版本信息
//
// @param request - GetLoginUserRequest
//
// @return GetLoginUserResponse
func (client *Client) GetLoginUser(request *GetLoginUserRequest) (_result *GetLoginUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLoginUserHeaders{}
	_result = &GetLoginUserResponse{}
	_body, _err := client.GetLoginUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提供给HHO的异步回调接口
//
// @param request - HhoCallBackRequest
//
// @param headers - HhoCallBackHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HhoCallBackResponse
func (client *Client) HhoCallBackWithOptions(request *HhoCallBackRequest, headers *HhoCallBackHeaders, runtime *util.RuntimeOptions) (_result *HhoCallBackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DtNotificationId)) {
		body["dtNotificationId"] = request.DtNotificationId
	}

	if !tea.BoolValue(util.IsUnset(request.ShopId)) {
		body["shopId"] = request.ShopId
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("HhoCallBack"),
		Version:     tea.String("aiGlobalEC_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiGlobalEC/hho/callback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HhoCallBackResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提供给HHO的异步回调接口
//
// @param request - HhoCallBackRequest
//
// @return HhoCallBackResponse
func (client *Client) HhoCallBack(request *HhoCallBackRequest) (_result *HhoCallBackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HhoCallBackHeaders{}
	_result = &HhoCallBackResponse{}
	_body, _err := client.HhoCallBackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 刊登的对外开放Api
//
// @param request - LaunchRequest
//
// @param headers - LaunchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LaunchResponse
func (client *Client) LaunchWithOptions(request *LaunchRequest, headers *LaunchHeaders, runtime *util.RuntimeOptions) (_result *LaunchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrls)) {
		body["imageUrls"] = request.ImageUrls
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.SellingPoints)) {
		body["sellingPoints"] = request.SellingPoints
	}

	if !tea.BoolValue(util.IsUnset(request.SourceData)) {
		body["sourceData"] = request.SourceData
	}

	if !tea.BoolValue(util.IsUnset(request.Variants)) {
		body["variants"] = request.Variants
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrls)) {
		body["videoUrls"] = request.VideoUrls
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
		Action:      tea.String("Launch"),
		Version:     tea.String("aiGlobalEC_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiGlobalEC/launch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LaunchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 刊登的对外开放Api
//
// @param request - LaunchRequest
//
// @return LaunchResponse
func (client *Client) Launch(request *LaunchRequest) (_result *LaunchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LaunchHeaders{}
	_result = &LaunchResponse{}
	_body, _err := client.LaunchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 全渠道AI表格业务信息
//
// @param request - QueryNotableInfoRequest
//
// @param headers - QueryNotableInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryNotableInfoResponse
func (client *Client) QueryNotableInfoWithOptions(request *QueryNotableInfoRequest, headers *QueryNotableInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryNotableInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["sceneCode"] = request.SceneCode
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
		Action:      tea.String("QueryNotableInfo"),
		Version:     tea.String("aiGlobalEC_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiGlobalEC/omniChannel/notable"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryNotableInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 全渠道AI表格业务信息
//
// @param request - QueryNotableInfoRequest
//
// @return QueryNotableInfoResponse
func (client *Client) QueryNotableInfo(request *QueryNotableInfoRequest) (_result *QueryNotableInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryNotableInfoHeaders{}
	_result = &QueryNotableInfoResponse{}
	_body, _err := client.QueryNotableInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 平台店铺授权回调
//
// @param request - TiktokShopAuthCallbackRequest
//
// @param headers - TiktokShopAuthCallbackHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TiktokShopAuthCallbackResponse
func (client *Client) TiktokShopAuthCallbackWithOptions(request *TiktokShopAuthCallbackRequest, headers *TiktokShopAuthCallbackHeaders, runtime *util.RuntimeOptions) (_result *TiktokShopAuthCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SellerType)) {
		body["sellerType"] = request.SellerType
	}

	if !tea.BoolValue(util.IsUnset(request.ShopId)) {
		body["shopId"] = request.ShopId
	}

	if !tea.BoolValue(util.IsUnset(request.ShopName)) {
		body["shopName"] = request.ShopName
	}

	if !tea.BoolValue(util.IsUnset(request.ShopRegion)) {
		body["shopRegion"] = request.ShopRegion
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
		Action:      tea.String("TiktokShopAuthCallback"),
		Version:     tea.String("aiGlobalEC_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiGlobalEC/omniChannel/tiktok/shop/authCallback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TiktokShopAuthCallbackResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 平台店铺授权回调
//
// @param request - TiktokShopAuthCallbackRequest
//
// @return TiktokShopAuthCallbackResponse
func (client *Client) TiktokShopAuthCallback(request *TiktokShopAuthCallbackRequest) (_result *TiktokShopAuthCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TiktokShopAuthCallbackHeaders{}
	_result = &TiktokShopAuthCallbackResponse{}
	_body, _err := client.TiktokShopAuthCallbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 全渠道运营Tiktok的Webhook信息写入
//
// @param request - TiktokWebhookProcessRequest
//
// @param headers - TiktokWebhookProcessHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TiktokWebhookProcessResponse
func (client *Client) TiktokWebhookProcessWithOptions(request *TiktokWebhookProcessRequest, headers *TiktokWebhookProcessHeaders, runtime *util.RuntimeOptions) (_result *TiktokWebhookProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TiktokContentJsonString)) {
		body["tiktokContentJsonString"] = request.TiktokContentJsonString
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
		Action:      tea.String("TiktokWebhookProcess"),
		Version:     tea.String("aiGlobalEC_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/aiGlobalEC/omniChannel/tiktok/webhook/process"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TiktokWebhookProcessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 全渠道运营Tiktok的Webhook信息写入
//
// @param request - TiktokWebhookProcessRequest
//
// @return TiktokWebhookProcessResponse
func (client *Client) TiktokWebhookProcess(request *TiktokWebhookProcessRequest) (_result *TiktokWebhookProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TiktokWebhookProcessHeaders{}
	_result = &TiktokWebhookProcessResponse{}
	_body, _err := client.TiktokWebhookProcessWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
