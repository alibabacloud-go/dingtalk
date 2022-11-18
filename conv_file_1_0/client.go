// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package conv_file_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceHeaders) GoString() string {
	return s.String()
}

func (s *GetSpaceHeaders) SetCommonHeaders(v map[string]*string) *GetSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *GetSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSpaceRequest struct {
	// 会话id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceRequest) SetOpenConversationId(v string) *GetSpaceRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetSpaceRequest) SetUnionId(v string) *GetSpaceRequest {
	s.UnionId = &v
	return s
}

type GetSpaceResponseBody struct {
	// IM会话存储空间信息
	Space *GetSpaceResponseBodySpace `json:"space,omitempty" xml:"space,omitempty" type:"Struct"`
}

func (s GetSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceResponseBody) SetSpace(v *GetSpaceResponseBodySpace) *GetSpaceResponseBody {
	s.Space = v
	return s
}

type GetSpaceResponseBodySpace struct {
	// 空间归属企业的id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetSpaceResponseBodySpace) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBodySpace) GoString() string {
	return s.String()
}

func (s *GetSpaceResponseBodySpace) SetCorpId(v string) *GetSpaceResponseBodySpace {
	s.CorpId = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetCreateTime(v string) *GetSpaceResponseBodySpace {
	s.CreateTime = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetModifiedTime(v string) *GetSpaceResponseBodySpace {
	s.ModifiedTime = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetSpaceId(v string) *GetSpaceResponseBodySpace {
	s.SpaceId = &v
	return s
}

type GetSpaceResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceResponse) SetHeaders(v map[string]*string) *GetSpaceResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceResponse) SetBody(v *GetSpaceResponseBody) *GetSpaceResponse {
	s.Body = v
	return s
}

type SendByAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendByAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendByAppHeaders) GoString() string {
	return s.String()
}

func (s *SendByAppHeaders) SetCommonHeaders(v map[string]*string) *SendByAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendByAppHeaders) SetXAcsDingtalkAccessToken(v string) *SendByAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendByAppRequest struct {
	// 文件id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// 文件所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SendByAppRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByAppRequest) GoString() string {
	return s.String()
}

func (s *SendByAppRequest) SetDentryId(v string) *SendByAppRequest {
	s.DentryId = &v
	return s
}

func (s *SendByAppRequest) SetSpaceId(v string) *SendByAppRequest {
	s.SpaceId = &v
	return s
}

func (s *SendByAppRequest) SetUnionId(v string) *SendByAppRequest {
	s.UnionId = &v
	return s
}

type SendByAppResponseBody struct {
	// 发送到目标会话的文件信息
	File *SendByAppResponseBodyFile `json:"file,omitempty" xml:"file,omitempty" type:"Struct"`
}

func (s SendByAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendByAppResponseBody) GoString() string {
	return s.String()
}

func (s *SendByAppResponseBody) SetFile(v *SendByAppResponseBodyFile) *SendByAppResponseBody {
	s.File = v
	return s
}

type SendByAppResponseBodyFile struct {
	// 文件所在会话id
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 文件后缀
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 文件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 文件(夹)名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文件所在的父目录id, 根目录id值为0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 文件大小, 单位:Byte
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 文件所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 文件状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETED: 已删除
	// 	EXPIRED: 已过期
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 文件类型：文件、文件夹
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// uuid，如移动文件，此字段不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 文件版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SendByAppResponseBodyFile) String() string {
	return tea.Prettify(s)
}

func (s SendByAppResponseBodyFile) GoString() string {
	return s.String()
}

func (s *SendByAppResponseBodyFile) SetConversationId(v string) *SendByAppResponseBodyFile {
	s.ConversationId = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetCreateTime(v string) *SendByAppResponseBodyFile {
	s.CreateTime = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetCreatorId(v string) *SendByAppResponseBodyFile {
	s.CreatorId = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetExtension(v string) *SendByAppResponseBodyFile {
	s.Extension = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetId(v string) *SendByAppResponseBodyFile {
	s.Id = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetModifiedTime(v string) *SendByAppResponseBodyFile {
	s.ModifiedTime = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetModifierId(v string) *SendByAppResponseBodyFile {
	s.ModifierId = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetName(v string) *SendByAppResponseBodyFile {
	s.Name = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetParentId(v string) *SendByAppResponseBodyFile {
	s.ParentId = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetPath(v string) *SendByAppResponseBodyFile {
	s.Path = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetSize(v int64) *SendByAppResponseBodyFile {
	s.Size = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetSpaceId(v string) *SendByAppResponseBodyFile {
	s.SpaceId = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetStatus(v string) *SendByAppResponseBodyFile {
	s.Status = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetType(v string) *SendByAppResponseBodyFile {
	s.Type = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetUuid(v string) *SendByAppResponseBodyFile {
	s.Uuid = &v
	return s
}

func (s *SendByAppResponseBodyFile) SetVersion(v int64) *SendByAppResponseBodyFile {
	s.Version = &v
	return s
}

type SendByAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendByAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendByAppResponse) String() string {
	return tea.Prettify(s)
}

func (s SendByAppResponse) GoString() string {
	return s.String()
}

func (s *SendByAppResponse) SetHeaders(v map[string]*string) *SendByAppResponse {
	s.Headers = v
	return s
}

func (s *SendByAppResponse) SetBody(v *SendByAppResponseBody) *SendByAppResponse {
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

func (client *Client) GetSpace(request *GetSpaceRequest) (_result *GetSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSpaceHeaders{}
	_result = &GetSpaceResponse{}
	_body, _err := client.GetSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpaceWithOptions(request *GetSpaceRequest, headers *GetSpaceHeaders, runtime *util.RuntimeOptions) (_result *GetSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &GetSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSpace"), tea.String("convFile_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/convFile/conversations/spaces/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendByApp(request *SendByAppRequest) (_result *SendByAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendByAppHeaders{}
	_result = &SendByAppResponse{}
	_body, _err := client.SendByAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendByAppWithOptions(request *SendByAppRequest, headers *SendByAppHeaders, runtime *util.RuntimeOptions) (_result *SendByAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryId)) {
		body["dentryId"] = request.DentryId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["spaceId"] = request.SpaceId
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
	_result = &SendByAppResponse{}
	_body, _err := client.DoROARequest(tea.String("SendByApp"), tea.String("convFile_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/convFile/apps/conversations/files/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
