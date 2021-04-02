// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package drive_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDownloadInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDownloadInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDownloadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDownloadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDownloadInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetDownloadInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDownloadInfoResponseBody struct {
	// 下载加签url信息
	PresignedUrlDownloadInfo *GetDownloadInfoResponseBodyPresignedUrlDownloadInfo `json:"presignedUrlDownloadInfo,omitempty" xml:"presignedUrlDownloadInfo,omitempty" type:"Struct"`
}

func (s GetDownloadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDownloadInfoResponseBody) SetPresignedUrlDownloadInfo(v *GetDownloadInfoResponseBodyPresignedUrlDownloadInfo) *GetDownloadInfoResponseBody {
	s.PresignedUrlDownloadInfo = v
	return s
}

type GetDownloadInfoResponseBodyPresignedUrlDownloadInfo struct {
	// 加签url
	ResourceUrl *string `json:"resourceUrl,omitempty" xml:"resourceUrl,omitempty"`
	// 加签url过期时间(分钟)
	Expiration *int32 `json:"expiration,omitempty" xml:"expiration,omitempty"`
}

func (s GetDownloadInfoResponseBodyPresignedUrlDownloadInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadInfoResponseBodyPresignedUrlDownloadInfo) GoString() string {
	return s.String()
}

func (s *GetDownloadInfoResponseBodyPresignedUrlDownloadInfo) SetResourceUrl(v string) *GetDownloadInfoResponseBodyPresignedUrlDownloadInfo {
	s.ResourceUrl = &v
	return s
}

func (s *GetDownloadInfoResponseBodyPresignedUrlDownloadInfo) SetExpiration(v int32) *GetDownloadInfoResponseBodyPresignedUrlDownloadInfo {
	s.Expiration = &v
	return s
}

type GetDownloadInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDownloadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDownloadInfoResponse) SetHeaders(v map[string]*string) *GetDownloadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDownloadInfoResponse) SetBody(v *GetDownloadInfoResponseBody) *GetDownloadInfoResponse {
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

func (client *Client) GetDownloadInfo(userId *string, spaceId *string, fileId *string) (_result *GetDownloadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDownloadInfoHeaders{}
	_result = &GetDownloadInfoResponse{}
	_body, _err := client.GetDownloadInfoWithOptions(userId, spaceId, fileId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDownloadInfoWithOptions(userId *string, spaceId *string, fileId *string, headers *GetDownloadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetDownloadInfoResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetDownloadInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDownloadInfo"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(userId)+"/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/downloadInfo"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
