// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package manufacturing_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type IndustrializeManufactureJobBookRequest struct {
	// 报废数量
	ScrappedQuantity *string `json:"scrappedQuantity,omitempty" xml:"scrappedQuantity,omitempty"`
	// 产品规格
	ProductSpecification *string `json:"productSpecification,omitempty" xml:"productSpecification,omitempty"`
	// 合格数量
	QualifiedQuantity *string `json:"qualifiedQuantity,omitempty" xml:"qualifiedQuantity,omitempty"`
	// 可重工数量
	ReworkableQuantity *string `json:"reworkableQuantity,omitempty" xml:"reworkableQuantity,omitempty"`
	// 员工姓名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 随机串，唯一标识(用于幂等及更新)
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 产品名称，例如"双头螺柱001"
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// 产品英文名称
	ProductEnName *string `json:"productEnName,omitempty" xml:"productEnName,omitempty"`
	// 扩展字段，用于增加自定义字段
	Extend *string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 产品唯一标识
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// 制程名称
	ProcessName *string `json:"processName,omitempty" xml:"processName,omitempty"`
	// 制程英文名称
	ProcessEnName *string `json:"processEnName,omitempty" xml:"processEnName,omitempty"`
	// mes 系统唯一标识
	MesAppKey *string `json:"mesAppKey,omitempty" xml:"mesAppKey,omitempty"`
	// 工单编号
	InstNo *string `json:"instNo,omitempty" xml:"instNo,omitempty"`
	// 生产日期时间(到时分秒)
	ManufactureDate *string `json:"manufactureDate,omitempty" xml:"manufactureDate,omitempty"`
	// 钉钉组织id
	DingCorpId *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
}

func (s IndustrializeManufactureJobBookRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureJobBookRequest) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureJobBookRequest) SetScrappedQuantity(v string) *IndustrializeManufactureJobBookRequest {
	s.ScrappedQuantity = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProductSpecification(v string) *IndustrializeManufactureJobBookRequest {
	s.ProductSpecification = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetQualifiedQuantity(v string) *IndustrializeManufactureJobBookRequest {
	s.QualifiedQuantity = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetReworkableQuantity(v string) *IndustrializeManufactureJobBookRequest {
	s.ReworkableQuantity = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetUserName(v string) *IndustrializeManufactureJobBookRequest {
	s.UserName = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetUuid(v string) *IndustrializeManufactureJobBookRequest {
	s.Uuid = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProductName(v string) *IndustrializeManufactureJobBookRequest {
	s.ProductName = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProductEnName(v string) *IndustrializeManufactureJobBookRequest {
	s.ProductEnName = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetExtend(v string) *IndustrializeManufactureJobBookRequest {
	s.Extend = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProductCode(v string) *IndustrializeManufactureJobBookRequest {
	s.ProductCode = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProcessName(v string) *IndustrializeManufactureJobBookRequest {
	s.ProcessName = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProcessEnName(v string) *IndustrializeManufactureJobBookRequest {
	s.ProcessEnName = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetMesAppKey(v string) *IndustrializeManufactureJobBookRequest {
	s.MesAppKey = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetInstNo(v string) *IndustrializeManufactureJobBookRequest {
	s.InstNo = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetManufactureDate(v string) *IndustrializeManufactureJobBookRequest {
	s.ManufactureDate = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetDingCorpId(v string) *IndustrializeManufactureJobBookRequest {
	s.DingCorpId = &v
	return s
}

type IndustrializeManufactureJobBookResponseBody struct {
	// httpCode
	HttpCode *string `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// 此次报工记录的唯一标识
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// errorMsg
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// errorLevel
	ErrorLevel *int32 `json:"errorLevel,omitempty" xml:"errorLevel,omitempty"`
	// errorCode
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s IndustrializeManufactureJobBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureJobBookResponseBody) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureJobBookResponseBody) SetHttpCode(v string) *IndustrializeManufactureJobBookResponseBody {
	s.HttpCode = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetUuid(v string) *IndustrializeManufactureJobBookResponseBody {
	s.Uuid = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetContent(v string) *IndustrializeManufactureJobBookResponseBody {
	s.Content = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetErrorMsg(v string) *IndustrializeManufactureJobBookResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetErrorLevel(v int32) *IndustrializeManufactureJobBookResponseBody {
	s.ErrorLevel = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetErrorCode(v string) *IndustrializeManufactureJobBookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetSuccess(v bool) *IndustrializeManufactureJobBookResponseBody {
	s.Success = &v
	return s
}

type IndustrializeManufactureJobBookResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndustrializeManufactureJobBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustrializeManufactureJobBookResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureJobBookResponse) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureJobBookResponse) SetHeaders(v map[string]*string) *IndustrializeManufactureJobBookResponse {
	s.Headers = v
	return s
}

func (s *IndustrializeManufactureJobBookResponse) SetBody(v *IndustrializeManufactureJobBookResponseBody) *IndustrializeManufactureJobBookResponse {
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

func (client *Client) IndustrializeManufactureJobBook(userId *string, request *IndustrializeManufactureJobBookRequest) (_result *IndustrializeManufactureJobBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &IndustrializeManufactureJobBookResponse{}
	_body, _err := client.IndustrializeManufactureJobBookWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustrializeManufactureJobBookWithOptions(userId *string, request *IndustrializeManufactureJobBookRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *IndustrializeManufactureJobBookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScrappedQuantity)) {
		body["scrappedQuantity"] = request.ScrappedQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.ProductSpecification)) {
		body["productSpecification"] = request.ProductSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.QualifiedQuantity)) {
		body["qualifiedQuantity"] = request.QualifiedQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.ReworkableQuantity)) {
		body["reworkableQuantity"] = request.ReworkableQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["userName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductEnName)) {
		body["productEnName"] = request.ProductEnName
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessName)) {
		body["processName"] = request.ProcessName
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessEnName)) {
		body["processEnName"] = request.ProcessEnName
	}

	if !tea.BoolValue(util.IsUnset(request.MesAppKey)) {
		body["mesAppKey"] = request.MesAppKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstNo)) {
		body["instNo"] = request.InstNo
	}

	if !tea.BoolValue(util.IsUnset(request.ManufactureDate)) {
		body["manufactureDate"] = request.ManufactureDate
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &IndustrializeManufactureJobBookResponse{}
	_body, _err := client.DoROARequest(tea.String("IndustrializeManufactureJobBook"), tea.String("manufacturing_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/manufacturing/users/"+tea.StringValue(userId)+"/jobs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
