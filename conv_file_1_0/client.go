// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package conv_file_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
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
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UnionId            *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	CorpId       *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CreateTime   *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	SpaceId      *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetSpaceResponse) SetStatusCode(v int32) *GetSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpaceResponse) SetBody(v *GetSpaceResponseBody) *GetSpaceResponse {
	s.Body = v
	return s
}

type SendHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendHeaders) GoString() string {
	return s.String()
}

func (s *SendHeaders) SetCommonHeaders(v map[string]*string) *SendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendHeaders) SetXAcsDingtalkAccessToken(v string) *SendHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendRequest struct {
	DentryId           *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	SpaceId            *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	UnionId            *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SendRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRequest) GoString() string {
	return s.String()
}

func (s *SendRequest) SetDentryId(v string) *SendRequest {
	s.DentryId = &v
	return s
}

func (s *SendRequest) SetOpenConversationId(v string) *SendRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendRequest) SetSpaceId(v string) *SendRequest {
	s.SpaceId = &v
	return s
}

func (s *SendRequest) SetUnionId(v string) *SendRequest {
	s.UnionId = &v
	return s
}

type SendResponseBody struct {
	File *SendResponseBodyFile `json:"file,omitempty" xml:"file,omitempty" type:"Struct"`
}

func (s SendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendResponseBody) GoString() string {
	return s.String()
}

func (s *SendResponseBody) SetFile(v *SendResponseBodyFile) *SendResponseBody {
	s.File = v
	return s
}

type SendResponseBodyFile struct {
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	CreateTime     *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId      *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension      *string `json:"extension,omitempty" xml:"extension,omitempty"`
	Id             *string `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedTime   *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId     *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentId       *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
	Size           *int64  `json:"size,omitempty" xml:"size,omitempty"`
	SpaceId        *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
	Uuid           *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version        *int64  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SendResponseBodyFile) String() string {
	return tea.Prettify(s)
}

func (s SendResponseBodyFile) GoString() string {
	return s.String()
}

func (s *SendResponseBodyFile) SetConversationId(v string) *SendResponseBodyFile {
	s.ConversationId = &v
	return s
}

func (s *SendResponseBodyFile) SetCreateTime(v string) *SendResponseBodyFile {
	s.CreateTime = &v
	return s
}

func (s *SendResponseBodyFile) SetCreatorId(v string) *SendResponseBodyFile {
	s.CreatorId = &v
	return s
}

func (s *SendResponseBodyFile) SetExtension(v string) *SendResponseBodyFile {
	s.Extension = &v
	return s
}

func (s *SendResponseBodyFile) SetId(v string) *SendResponseBodyFile {
	s.Id = &v
	return s
}

func (s *SendResponseBodyFile) SetModifiedTime(v string) *SendResponseBodyFile {
	s.ModifiedTime = &v
	return s
}

func (s *SendResponseBodyFile) SetModifierId(v string) *SendResponseBodyFile {
	s.ModifierId = &v
	return s
}

func (s *SendResponseBodyFile) SetName(v string) *SendResponseBodyFile {
	s.Name = &v
	return s
}

func (s *SendResponseBodyFile) SetParentId(v string) *SendResponseBodyFile {
	s.ParentId = &v
	return s
}

func (s *SendResponseBodyFile) SetPath(v string) *SendResponseBodyFile {
	s.Path = &v
	return s
}

func (s *SendResponseBodyFile) SetSize(v int64) *SendResponseBodyFile {
	s.Size = &v
	return s
}

func (s *SendResponseBodyFile) SetSpaceId(v string) *SendResponseBodyFile {
	s.SpaceId = &v
	return s
}

func (s *SendResponseBodyFile) SetStatus(v string) *SendResponseBodyFile {
	s.Status = &v
	return s
}

func (s *SendResponseBodyFile) SetType(v string) *SendResponseBodyFile {
	s.Type = &v
	return s
}

func (s *SendResponseBodyFile) SetUuid(v string) *SendResponseBodyFile {
	s.Uuid = &v
	return s
}

func (s *SendResponseBodyFile) SetVersion(v int64) *SendResponseBodyFile {
	s.Version = &v
	return s
}

type SendResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendResponse) String() string {
	return tea.Prettify(s)
}

func (s SendResponse) GoString() string {
	return s.String()
}

func (s *SendResponse) SetHeaders(v map[string]*string) *SendResponse {
	s.Headers = v
	return s
}

func (s *SendResponse) SetStatusCode(v int32) *SendResponse {
	s.StatusCode = &v
	return s
}

func (s *SendResponse) SetBody(v *SendResponseBody) *SendResponse {
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
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	SpaceId  *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	CreateTime     *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId      *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension      *string `json:"extension,omitempty" xml:"extension,omitempty"`
	Id             *string `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedTime   *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId     *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentId       *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
	Size           *int64  `json:"size,omitempty" xml:"size,omitempty"`
	SpaceId        *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
	Uuid           *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version        *int64  `json:"version,omitempty" xml:"version,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SendByAppResponse) SetStatusCode(v int32) *SendByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByAppResponse) SetBody(v *SendByAppResponseBody) *SendByAppResponse {
	s.Body = v
	return s
}

type SendLinkHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendLinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendLinkHeaders) GoString() string {
	return s.String()
}

func (s *SendLinkHeaders) SetCommonHeaders(v map[string]*string) *SendLinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendLinkHeaders) SetXAcsDingtalkAccessToken(v string) *SendLinkHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendLinkRequest struct {
	DentryId           *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	SpaceId            *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	UnionId            *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SendLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendLinkRequest) GoString() string {
	return s.String()
}

func (s *SendLinkRequest) SetDentryId(v string) *SendLinkRequest {
	s.DentryId = &v
	return s
}

func (s *SendLinkRequest) SetOpenConversationId(v string) *SendLinkRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendLinkRequest) SetSpaceId(v string) *SendLinkRequest {
	s.SpaceId = &v
	return s
}

func (s *SendLinkRequest) SetUnionId(v string) *SendLinkRequest {
	s.UnionId = &v
	return s
}

type SendLinkResponseBody struct {
	File *SendLinkResponseBodyFile `json:"file,omitempty" xml:"file,omitempty" type:"Struct"`
}

func (s SendLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendLinkResponseBody) GoString() string {
	return s.String()
}

func (s *SendLinkResponseBody) SetFile(v *SendLinkResponseBodyFile) *SendLinkResponseBody {
	s.File = v
	return s
}

type SendLinkResponseBodyFile struct {
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	CreateTime     *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId      *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension      *string `json:"extension,omitempty" xml:"extension,omitempty"`
	Id             *string `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedTime   *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId     *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentId       *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
	Size           *int64  `json:"size,omitempty" xml:"size,omitempty"`
	SpaceId        *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
	Uuid           *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version        *int64  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SendLinkResponseBodyFile) String() string {
	return tea.Prettify(s)
}

func (s SendLinkResponseBodyFile) GoString() string {
	return s.String()
}

func (s *SendLinkResponseBodyFile) SetConversationId(v string) *SendLinkResponseBodyFile {
	s.ConversationId = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetCreateTime(v string) *SendLinkResponseBodyFile {
	s.CreateTime = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetCreatorId(v string) *SendLinkResponseBodyFile {
	s.CreatorId = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetExtension(v string) *SendLinkResponseBodyFile {
	s.Extension = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetId(v string) *SendLinkResponseBodyFile {
	s.Id = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetModifiedTime(v string) *SendLinkResponseBodyFile {
	s.ModifiedTime = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetModifierId(v string) *SendLinkResponseBodyFile {
	s.ModifierId = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetName(v string) *SendLinkResponseBodyFile {
	s.Name = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetParentId(v string) *SendLinkResponseBodyFile {
	s.ParentId = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetPath(v string) *SendLinkResponseBodyFile {
	s.Path = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetSize(v int64) *SendLinkResponseBodyFile {
	s.Size = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetSpaceId(v string) *SendLinkResponseBodyFile {
	s.SpaceId = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetStatus(v string) *SendLinkResponseBodyFile {
	s.Status = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetType(v string) *SendLinkResponseBodyFile {
	s.Type = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetUuid(v string) *SendLinkResponseBodyFile {
	s.Uuid = &v
	return s
}

func (s *SendLinkResponseBodyFile) SetVersion(v int64) *SendLinkResponseBodyFile {
	s.Version = &v
	return s
}

type SendLinkResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s SendLinkResponse) GoString() string {
	return s.String()
}

func (s *SendLinkResponse) SetHeaders(v map[string]*string) *SendLinkResponse {
	s.Headers = v
	return s
}

func (s *SendLinkResponse) SetStatusCode(v int32) *SendLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *SendLinkResponse) SetBody(v *SendLinkResponseBody) *SendLinkResponse {
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
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
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
	params := &openapi.Params{
		Action:      tea.String("GetSpace"),
		Version:     tea.String("convFile_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/convFile/conversations/spaces/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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

func (client *Client) SendWithOptions(request *SendRequest, headers *SendHeaders, runtime *util.RuntimeOptions) (_result *SendResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	params := &openapi.Params{
		Action:      tea.String("Send"),
		Version:     tea.String("convFile_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/convFile/conversations/files/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Send(request *SendRequest) (_result *SendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendHeaders{}
	_result = &SendResponse{}
	_body, _err := client.SendWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("SendByApp"),
		Version:     tea.String("convFile_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/convFile/apps/conversations/files/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendByAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) SendLinkWithOptions(request *SendLinkRequest, headers *SendLinkHeaders, runtime *util.RuntimeOptions) (_result *SendLinkResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	params := &openapi.Params{
		Action:      tea.String("SendLink"),
		Version:     tea.String("convFile_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/convFile/conversations/files/links/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendLinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendLink(request *SendLinkRequest) (_result *SendLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendLinkHeaders{}
	_result = &SendLinkResponse{}
	_body, _err := client.SendLinkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
