// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package algo_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type NlpWordDistinguishHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NlpWordDistinguishHeaders) String() string {
	return tea.Prettify(s)
}

func (s NlpWordDistinguishHeaders) GoString() string {
	return s.String()
}

func (s *NlpWordDistinguishHeaders) SetCommonHeaders(v map[string]*string) *NlpWordDistinguishHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NlpWordDistinguishHeaders) SetXAcsDingtalkAccessToken(v string) *NlpWordDistinguishHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NlpWordDistinguishRequest struct {
	AttachExtractDecisionInfo *NlpWordDistinguishRequestAttachExtractDecisionInfo `json:"attachExtractDecisionInfo,omitempty" xml:"attachExtractDecisionInfo,omitempty" type:"Struct"`
	IsvAppId                  *string                                             `json:"isvAppId,omitempty" xml:"isvAppId,omitempty"`
	Text                      *string                                             `json:"text,omitempty" xml:"text,omitempty"`
}

func (s NlpWordDistinguishRequest) String() string {
	return tea.Prettify(s)
}

func (s NlpWordDistinguishRequest) GoString() string {
	return s.String()
}

func (s *NlpWordDistinguishRequest) SetAttachExtractDecisionInfo(v *NlpWordDistinguishRequestAttachExtractDecisionInfo) *NlpWordDistinguishRequest {
	s.AttachExtractDecisionInfo = v
	return s
}

func (s *NlpWordDistinguishRequest) SetIsvAppId(v string) *NlpWordDistinguishRequest {
	s.IsvAppId = &v
	return s
}

func (s *NlpWordDistinguishRequest) SetText(v string) *NlpWordDistinguishRequest {
	s.Text = &v
	return s
}

type NlpWordDistinguishRequestAttachExtractDecisionInfo struct {
	BlackWords     []*string `json:"blackWords,omitempty" xml:"blackWords,omitempty" type:"Repeated"`
	CandidateWords []*string `json:"candidateWords,omitempty" xml:"candidateWords,omitempty" type:"Repeated"`
	CorpId         *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptIds        []*string `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	UserId         *string   `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s NlpWordDistinguishRequestAttachExtractDecisionInfo) String() string {
	return tea.Prettify(s)
}

func (s NlpWordDistinguishRequestAttachExtractDecisionInfo) GoString() string {
	return s.String()
}

func (s *NlpWordDistinguishRequestAttachExtractDecisionInfo) SetBlackWords(v []*string) *NlpWordDistinguishRequestAttachExtractDecisionInfo {
	s.BlackWords = v
	return s
}

func (s *NlpWordDistinguishRequestAttachExtractDecisionInfo) SetCandidateWords(v []*string) *NlpWordDistinguishRequestAttachExtractDecisionInfo {
	s.CandidateWords = v
	return s
}

func (s *NlpWordDistinguishRequestAttachExtractDecisionInfo) SetCorpId(v string) *NlpWordDistinguishRequestAttachExtractDecisionInfo {
	s.CorpId = &v
	return s
}

func (s *NlpWordDistinguishRequestAttachExtractDecisionInfo) SetDeptIds(v []*string) *NlpWordDistinguishRequestAttachExtractDecisionInfo {
	s.DeptIds = v
	return s
}

func (s *NlpWordDistinguishRequestAttachExtractDecisionInfo) SetUserId(v string) *NlpWordDistinguishRequestAttachExtractDecisionInfo {
	s.UserId = &v
	return s
}

type NlpWordDistinguishResponseBody struct {
	RequestId    *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	WordEntities []*NlpWordDistinguishResponseBodyWordEntities `json:"wordEntities,omitempty" xml:"wordEntities,omitempty" type:"Repeated"`
}

func (s NlpWordDistinguishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NlpWordDistinguishResponseBody) GoString() string {
	return s.String()
}

func (s *NlpWordDistinguishResponseBody) SetRequestId(v string) *NlpWordDistinguishResponseBody {
	s.RequestId = &v
	return s
}

func (s *NlpWordDistinguishResponseBody) SetWordEntities(v []*NlpWordDistinguishResponseBodyWordEntities) *NlpWordDistinguishResponseBody {
	s.WordEntities = v
	return s
}

type NlpWordDistinguishResponseBodyWordEntities struct {
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s NlpWordDistinguishResponseBodyWordEntities) String() string {
	return tea.Prettify(s)
}

func (s NlpWordDistinguishResponseBodyWordEntities) GoString() string {
	return s.String()
}

func (s *NlpWordDistinguishResponseBodyWordEntities) SetWord(v string) *NlpWordDistinguishResponseBodyWordEntities {
	s.Word = &v
	return s
}

type NlpWordDistinguishResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *NlpWordDistinguishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NlpWordDistinguishResponse) String() string {
	return tea.Prettify(s)
}

func (s NlpWordDistinguishResponse) GoString() string {
	return s.String()
}

func (s *NlpWordDistinguishResponse) SetHeaders(v map[string]*string) *NlpWordDistinguishResponse {
	s.Headers = v
	return s
}

func (s *NlpWordDistinguishResponse) SetStatusCode(v int32) *NlpWordDistinguishResponse {
	s.StatusCode = &v
	return s
}

func (s *NlpWordDistinguishResponse) SetBody(v *NlpWordDistinguishResponseBody) *NlpWordDistinguishResponse {
	s.Body = v
	return s
}

type OkrOpenRecommendHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OkrOpenRecommendHeaders) String() string {
	return tea.Prettify(s)
}

func (s OkrOpenRecommendHeaders) GoString() string {
	return s.String()
}

func (s *OkrOpenRecommendHeaders) SetCommonHeaders(v map[string]*string) *OkrOpenRecommendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OkrOpenRecommendHeaders) SetXAcsDingtalkAccessToken(v string) *OkrOpenRecommendHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OkrOpenRecommendRequest struct {
	CandidateOkrItems []*OkrOpenRecommendRequestCandidateOkrItems `json:"candidateOkrItems,omitempty" xml:"candidateOkrItems,omitempty" type:"Repeated"`
	CorpId            *string                                     `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptIds           []*string                                   `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	IsvAppId          *string                                     `json:"isvAppId,omitempty" xml:"isvAppId,omitempty"`
	UserId            *string                                     `json:"userId,omitempty" xml:"userId,omitempty"`
	Words             []*string                                   `json:"words,omitempty" xml:"words,omitempty" type:"Repeated"`
}

func (s OkrOpenRecommendRequest) String() string {
	return tea.Prettify(s)
}

func (s OkrOpenRecommendRequest) GoString() string {
	return s.String()
}

func (s *OkrOpenRecommendRequest) SetCandidateOkrItems(v []*OkrOpenRecommendRequestCandidateOkrItems) *OkrOpenRecommendRequest {
	s.CandidateOkrItems = v
	return s
}

func (s *OkrOpenRecommendRequest) SetCorpId(v string) *OkrOpenRecommendRequest {
	s.CorpId = &v
	return s
}

func (s *OkrOpenRecommendRequest) SetDeptIds(v []*string) *OkrOpenRecommendRequest {
	s.DeptIds = v
	return s
}

func (s *OkrOpenRecommendRequest) SetIsvAppId(v string) *OkrOpenRecommendRequest {
	s.IsvAppId = &v
	return s
}

func (s *OkrOpenRecommendRequest) SetUserId(v string) *OkrOpenRecommendRequest {
	s.UserId = &v
	return s
}

func (s *OkrOpenRecommendRequest) SetWords(v []*string) *OkrOpenRecommendRequest {
	s.Words = v
	return s
}

type OkrOpenRecommendRequestCandidateOkrItems struct {
	OkrInfos []*OkrOpenRecommendRequestCandidateOkrItemsOkrInfos `json:"okrInfos,omitempty" xml:"okrInfos,omitempty" type:"Repeated"`
	Relation *string                                             `json:"relation,omitempty" xml:"relation,omitempty"`
	UserId   *string                                             `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OkrOpenRecommendRequestCandidateOkrItems) String() string {
	return tea.Prettify(s)
}

func (s OkrOpenRecommendRequestCandidateOkrItems) GoString() string {
	return s.String()
}

func (s *OkrOpenRecommendRequestCandidateOkrItems) SetOkrInfos(v []*OkrOpenRecommendRequestCandidateOkrItemsOkrInfos) *OkrOpenRecommendRequestCandidateOkrItems {
	s.OkrInfos = v
	return s
}

func (s *OkrOpenRecommendRequestCandidateOkrItems) SetRelation(v string) *OkrOpenRecommendRequestCandidateOkrItems {
	s.Relation = &v
	return s
}

func (s *OkrOpenRecommendRequestCandidateOkrItems) SetUserId(v string) *OkrOpenRecommendRequestCandidateOkrItems {
	s.UserId = &v
	return s
}

type OkrOpenRecommendRequestCandidateOkrItemsOkrInfos struct {
	KeyResultInfos []*OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos `json:"keyResultInfos,omitempty" xml:"keyResultInfos,omitempty" type:"Repeated"`
	Objective      *string                                                           `json:"objective,omitempty" xml:"objective,omitempty"`
	ObjectiveId    *string                                                           `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	Words          []*string                                                         `json:"words,omitempty" xml:"words,omitempty" type:"Repeated"`
}

func (s OkrOpenRecommendRequestCandidateOkrItemsOkrInfos) String() string {
	return tea.Prettify(s)
}

func (s OkrOpenRecommendRequestCandidateOkrItemsOkrInfos) GoString() string {
	return s.String()
}

func (s *OkrOpenRecommendRequestCandidateOkrItemsOkrInfos) SetKeyResultInfos(v []*OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos) *OkrOpenRecommendRequestCandidateOkrItemsOkrInfos {
	s.KeyResultInfos = v
	return s
}

func (s *OkrOpenRecommendRequestCandidateOkrItemsOkrInfos) SetObjective(v string) *OkrOpenRecommendRequestCandidateOkrItemsOkrInfos {
	s.Objective = &v
	return s
}

func (s *OkrOpenRecommendRequestCandidateOkrItemsOkrInfos) SetObjectiveId(v string) *OkrOpenRecommendRequestCandidateOkrItemsOkrInfos {
	s.ObjectiveId = &v
	return s
}

func (s *OkrOpenRecommendRequestCandidateOkrItemsOkrInfos) SetWords(v []*string) *OkrOpenRecommendRequestCandidateOkrItemsOkrInfos {
	s.Words = v
	return s
}

type OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos struct {
	Kr    *string   `json:"kr,omitempty" xml:"kr,omitempty"`
	KrId  *string   `json:"krId,omitempty" xml:"krId,omitempty"`
	Words []*string `json:"words,omitempty" xml:"words,omitempty" type:"Repeated"`
}

func (s OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos) String() string {
	return tea.Prettify(s)
}

func (s OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos) GoString() string {
	return s.String()
}

func (s *OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos) SetKr(v string) *OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos {
	s.Kr = &v
	return s
}

func (s *OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos) SetKrId(v string) *OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos {
	s.KrId = &v
	return s
}

func (s *OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos) SetWords(v []*string) *OkrOpenRecommendRequestCandidateOkrItemsOkrInfosKeyResultInfos {
	s.Words = v
	return s
}

type OkrOpenRecommendResponseBody struct {
	OkrRecommendItems []*OkrOpenRecommendResponseBodyOkrRecommendItems `json:"okrRecommendItems,omitempty" xml:"okrRecommendItems,omitempty" type:"Repeated"`
	RequestId         *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OkrOpenRecommendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OkrOpenRecommendResponseBody) GoString() string {
	return s.String()
}

func (s *OkrOpenRecommendResponseBody) SetOkrRecommendItems(v []*OkrOpenRecommendResponseBodyOkrRecommendItems) *OkrOpenRecommendResponseBody {
	s.OkrRecommendItems = v
	return s
}

func (s *OkrOpenRecommendResponseBody) SetRequestId(v string) *OkrOpenRecommendResponseBody {
	s.RequestId = &v
	return s
}

type OkrOpenRecommendResponseBodyOkrRecommendItems struct {
	KrResultRelatedResults  []*OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults  `json:"krResultRelatedResults,omitempty" xml:"krResultRelatedResults,omitempty" type:"Repeated"`
	ObjectiveRelatedResults []*OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults `json:"objectiveRelatedResults,omitempty" xml:"objectiveRelatedResults,omitempty" type:"Repeated"`
	RelatedLevel            *int64                                                                  `json:"relatedLevel,omitempty" xml:"relatedLevel,omitempty"`
	SemanticLevel           *int64                                                                  `json:"semanticLevel,omitempty" xml:"semanticLevel,omitempty"`
	UserId                  *string                                                                 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OkrOpenRecommendResponseBodyOkrRecommendItems) String() string {
	return tea.Prettify(s)
}

func (s OkrOpenRecommendResponseBodyOkrRecommendItems) GoString() string {
	return s.String()
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItems) SetKrResultRelatedResults(v []*OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults) *OkrOpenRecommendResponseBodyOkrRecommendItems {
	s.KrResultRelatedResults = v
	return s
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItems) SetObjectiveRelatedResults(v []*OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults) *OkrOpenRecommendResponseBodyOkrRecommendItems {
	s.ObjectiveRelatedResults = v
	return s
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItems) SetRelatedLevel(v int64) *OkrOpenRecommendResponseBodyOkrRecommendItems {
	s.RelatedLevel = &v
	return s
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItems) SetSemanticLevel(v int64) *OkrOpenRecommendResponseBodyOkrRecommendItems {
	s.SemanticLevel = &v
	return s
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItems) SetUserId(v string) *OkrOpenRecommendResponseBodyOkrRecommendItems {
	s.UserId = &v
	return s
}

type OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults struct {
	KrId          *string   `json:"krId,omitempty" xml:"krId,omitempty"`
	SemanticLevel *int64    `json:"semanticLevel,omitempty" xml:"semanticLevel,omitempty"`
	Words         []*string `json:"words,omitempty" xml:"words,omitempty" type:"Repeated"`
}

func (s OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults) String() string {
	return tea.Prettify(s)
}

func (s OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults) GoString() string {
	return s.String()
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults) SetKrId(v string) *OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults {
	s.KrId = &v
	return s
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults) SetSemanticLevel(v int64) *OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults {
	s.SemanticLevel = &v
	return s
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults) SetWords(v []*string) *OkrOpenRecommendResponseBodyOkrRecommendItemsKrResultRelatedResults {
	s.Words = v
	return s
}

type OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults struct {
	ObjectiveId   *string   `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	SemanticLevel *int64    `json:"semanticLevel,omitempty" xml:"semanticLevel,omitempty"`
	Words         []*string `json:"words,omitempty" xml:"words,omitempty" type:"Repeated"`
}

func (s OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults) String() string {
	return tea.Prettify(s)
}

func (s OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults) GoString() string {
	return s.String()
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults) SetObjectiveId(v string) *OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults {
	s.ObjectiveId = &v
	return s
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults) SetSemanticLevel(v int64) *OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults {
	s.SemanticLevel = &v
	return s
}

func (s *OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults) SetWords(v []*string) *OkrOpenRecommendResponseBodyOkrRecommendItemsObjectiveRelatedResults {
	s.Words = v
	return s
}

type OkrOpenRecommendResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OkrOpenRecommendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OkrOpenRecommendResponse) String() string {
	return tea.Prettify(s)
}

func (s OkrOpenRecommendResponse) GoString() string {
	return s.String()
}

func (s *OkrOpenRecommendResponse) SetHeaders(v map[string]*string) *OkrOpenRecommendResponse {
	s.Headers = v
	return s
}

func (s *OkrOpenRecommendResponse) SetStatusCode(v int32) *OkrOpenRecommendResponse {
	s.StatusCode = &v
	return s
}

func (s *OkrOpenRecommendResponse) SetBody(v *OkrOpenRecommendResponseBody) *OkrOpenRecommendResponse {
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

func (client *Client) NlpWordDistinguishWithOptions(request *NlpWordDistinguishRequest, headers *NlpWordDistinguishHeaders, runtime *util.RuntimeOptions) (_result *NlpWordDistinguishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachExtractDecisionInfo)) {
		body["attachExtractDecisionInfo"] = request.AttachExtractDecisionInfo
	}

	if !tea.BoolValue(util.IsUnset(request.IsvAppId)) {
		body["isvAppId"] = request.IsvAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["text"] = request.Text
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
		Action:      tea.String("NlpWordDistinguish"),
		Version:     tea.String("algo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/algo/okrs/keywords/extract"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &NlpWordDistinguishResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NlpWordDistinguish(request *NlpWordDistinguishRequest) (_result *NlpWordDistinguishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NlpWordDistinguishHeaders{}
	_result = &NlpWordDistinguishResponse{}
	_body, _err := client.NlpWordDistinguishWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OkrOpenRecommendWithOptions(request *OkrOpenRecommendRequest, headers *OkrOpenRecommendHeaders, runtime *util.RuntimeOptions) (_result *OkrOpenRecommendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CandidateOkrItems)) {
		body["candidateOkrItems"] = request.CandidateOkrItems
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DeptIds)) {
		body["deptIds"] = request.DeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.IsvAppId)) {
		body["isvAppId"] = request.IsvAppId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Words)) {
		body["words"] = request.Words
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
		Action:      tea.String("OkrOpenRecommend"),
		Version:     tea.String("algo_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/algo/okrs/recommend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OkrOpenRecommendResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OkrOpenRecommend(request *OkrOpenRecommendRequest) (_result *OkrOpenRecommendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OkrOpenRecommendHeaders{}
	_result = &OkrOpenRecommendResponse{}
	_body, _err := client.OkrOpenRecommendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
