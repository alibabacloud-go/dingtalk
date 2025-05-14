// This file is auto-generated, don't edit it. Thanks.
package agoal_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Entity struct {
	Children []*Entity `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// example:
	//
	// {"title": "123"}
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 123
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// y/n
	IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// example:
	//
	// 67dbb24f7aac3f62d8b98fb5
	LinkSourceId *string `json:"linkSourceId,omitempty" xml:"linkSourceId,omitempty"`
	// example:
	//
	// EXTERNAL_PERF_TASK
	LinkSourceType *string `json:"linkSourceType,omitempty" xml:"linkSourceType,omitempty"`
	Metas          []*Meta `json:"metas,omitempty" xml:"metas,omitempty" type:"Repeated"`
	// example:
	//
	// DIMENSION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Entity) String() string {
	return tea.Prettify(s)
}

func (s Entity) GoString() string {
	return s.String()
}

func (s *Entity) SetChildren(v []*Entity) *Entity {
	s.Children = v
	return s
}

func (s *Entity) SetData(v map[string]interface{}) *Entity {
	s.Data = v
	return s
}

func (s *Entity) SetId(v string) *Entity {
	s.Id = &v
	return s
}

func (s *Entity) SetIsDeleted(v string) *Entity {
	s.IsDeleted = &v
	return s
}

func (s *Entity) SetLinkSourceId(v string) *Entity {
	s.LinkSourceId = &v
	return s
}

func (s *Entity) SetLinkSourceType(v string) *Entity {
	s.LinkSourceType = &v
	return s
}

func (s *Entity) SetMetas(v []*Meta) *Entity {
	s.Metas = v
	return s
}

func (s *Entity) SetType(v string) *Entity {
	s.Type = &v
	return s
}

type Meta struct {
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// example:
	//
	// 编码
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// common
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// title
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	ForceActive *bool `json:"forceActive,omitempty" xml:"forceActive,omitempty"`
	// example:
	//
	// true
	ForceRequired *bool `json:"forceRequired,omitempty" xml:"forceRequired,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// example:
	//
	// {"width": 200}
	Scheme map[string]interface{} `json:"scheme,omitempty" xml:"scheme,omitempty"`
	// example:
	//
	// 名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Meta) String() string {
	return tea.Prettify(s)
}

func (s Meta) GoString() string {
	return s.String()
}

func (s *Meta) SetActive(v bool) *Meta {
	s.Active = &v
	return s
}

func (s *Meta) SetAlias(v string) *Meta {
	s.Alias = &v
	return s
}

func (s *Meta) SetCategory(v string) *Meta {
	s.Category = &v
	return s
}

func (s *Meta) SetCode(v string) *Meta {
	s.Code = &v
	return s
}

func (s *Meta) SetForceActive(v bool) *Meta {
	s.ForceActive = &v
	return s
}

func (s *Meta) SetForceRequired(v bool) *Meta {
	s.ForceRequired = &v
	return s
}

func (s *Meta) SetRequired(v bool) *Meta {
	s.Required = &v
	return s
}

func (s *Meta) SetScheme(v map[string]interface{}) *Meta {
	s.Scheme = v
	return s
}

func (s *Meta) SetTitle(v string) *Meta {
	s.Title = &v
	return s
}

func (s *Meta) SetType(v string) *Meta {
	s.Type = &v
	return s
}

type OpenAgoalAlignDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// COOPERATION
	AlignType *string `json:"alignType,omitempty" xml:"alignType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 662e006fe4b0f579bbcxxxxx
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// objective
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 662e006fe4b0f579bbcxxxxx
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
}

func (s OpenAgoalAlignDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalAlignDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalAlignDTO) SetAlignType(v string) *OpenAgoalAlignDTO {
	s.AlignType = &v
	return s
}

func (s *OpenAgoalAlignDTO) SetObjectId(v string) *OpenAgoalAlignDTO {
	s.ObjectId = &v
	return s
}

func (s *OpenAgoalAlignDTO) SetObjectType(v string) *OpenAgoalAlignDTO {
	s.ObjectType = &v
	return s
}

func (s *OpenAgoalAlignDTO) SetObjectiveId(v string) *OpenAgoalAlignDTO {
	s.ObjectiveId = &v
	return s
}

type OpenAgoalFieldMetaDTO struct {
	// 是否启用
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// 字段元数据别名
	//
	// example:
	//
	// 字段别名
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// 字段元数据标识
	//
	// This parameter is required.
	//
	// example:
	//
	// foo
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 实体类型
	//
	// This parameter is required.
	//
	// example:
	//
	// OBJECTIVE
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// 字段ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 662e006fe4b0f579bbcxxxxx
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// 字段备注
	//
	// example:
	//
	// 字段备注
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// 字段数据来源
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 字段元数据名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 字段名
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 字段类型
	//
	// This parameter is required.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OpenAgoalFieldMetaDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalFieldMetaDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalFieldMetaDTO) SetActive(v bool) *OpenAgoalFieldMetaDTO {
	s.Active = &v
	return s
}

func (s *OpenAgoalFieldMetaDTO) SetAlias(v string) *OpenAgoalFieldMetaDTO {
	s.Alias = &v
	return s
}

func (s *OpenAgoalFieldMetaDTO) SetCode(v string) *OpenAgoalFieldMetaDTO {
	s.Code = &v
	return s
}

func (s *OpenAgoalFieldMetaDTO) SetEntityType(v string) *OpenAgoalFieldMetaDTO {
	s.EntityType = &v
	return s
}

func (s *OpenAgoalFieldMetaDTO) SetFieldId(v string) *OpenAgoalFieldMetaDTO {
	s.FieldId = &v
	return s
}

func (s *OpenAgoalFieldMetaDTO) SetNote(v string) *OpenAgoalFieldMetaDTO {
	s.Note = &v
	return s
}

func (s *OpenAgoalFieldMetaDTO) SetSource(v string) *OpenAgoalFieldMetaDTO {
	s.Source = &v
	return s
}

func (s *OpenAgoalFieldMetaDTO) SetTitle(v string) *OpenAgoalFieldMetaDTO {
	s.Title = &v
	return s
}

func (s *OpenAgoalFieldMetaDTO) SetType(v string) *OpenAgoalFieldMetaDTO {
	s.Type = &v
	return s
}

type OpenAgoalKeyActionDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	KeyActionId *string `json:"keyActionId,omitempty" xml:"keyActionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://agoal.dingtalk.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s OpenAgoalKeyActionDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalKeyActionDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalKeyActionDTO) SetKeyActionId(v string) *OpenAgoalKeyActionDTO {
	s.KeyActionId = &v
	return s
}

func (s *OpenAgoalKeyActionDTO) SetTitle(v string) *OpenAgoalKeyActionDTO {
	s.Title = &v
	return s
}

func (s *OpenAgoalKeyActionDTO) SetUrl(v string) *OpenAgoalKeyActionDTO {
	s.Url = &v
	return s
}

type OpenAgoalKeyResultDTO struct {
	// This parameter is required.
	KeyActions []*OpenAgoalKeyActionDTO `json:"keyActions,omitempty" xml:"keyActions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	KeyResultId *string `json:"keyResultId,omitempty" xml:"keyResultId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Progress *int32 `json:"progress,omitempty" xml:"progress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试KR
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	TitleMentions []*TitleMention `json:"titleMentions,omitempty" xml:"titleMentions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	Weight *float64 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s OpenAgoalKeyResultDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalKeyResultDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalKeyResultDTO) SetKeyActions(v []*OpenAgoalKeyActionDTO) *OpenAgoalKeyResultDTO {
	s.KeyActions = v
	return s
}

func (s *OpenAgoalKeyResultDTO) SetKeyResultId(v string) *OpenAgoalKeyResultDTO {
	s.KeyResultId = &v
	return s
}

func (s *OpenAgoalKeyResultDTO) SetProgress(v int32) *OpenAgoalKeyResultDTO {
	s.Progress = &v
	return s
}

func (s *OpenAgoalKeyResultDTO) SetStatus(v int32) *OpenAgoalKeyResultDTO {
	s.Status = &v
	return s
}

func (s *OpenAgoalKeyResultDTO) SetTitle(v string) *OpenAgoalKeyResultDTO {
	s.Title = &v
	return s
}

func (s *OpenAgoalKeyResultDTO) SetTitleMentions(v []*TitleMention) *OpenAgoalKeyResultDTO {
	s.TitleMentions = v
	return s
}

func (s *OpenAgoalKeyResultDTO) SetType(v int32) *OpenAgoalKeyResultDTO {
	s.Type = &v
	return s
}

func (s *OpenAgoalKeyResultDTO) SetWeight(v float64) *OpenAgoalKeyResultDTO {
	s.Weight = &v
	return s
}

type OpenAgoalLatestProgressDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// 1716952481672
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// This parameter is required.
	Creator *OpenAgoalUserDTO `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// <article class=\"4ever-article\"><p style=\"text-align:left;text-indent:0;margin-left:0;margin-top:0;margin-bottom:0\"><span>xxx</span></p></article>
	Htmldescription *string `json:"htmldescription,omitempty" xml:"htmldescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	ProgressId *string `json:"progressId,omitempty" xml:"progressId,omitempty"`
}

func (s OpenAgoalLatestProgressDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalLatestProgressDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalLatestProgressDTO) SetCreated(v int64) *OpenAgoalLatestProgressDTO {
	s.Created = &v
	return s
}

func (s *OpenAgoalLatestProgressDTO) SetCreator(v *OpenAgoalUserDTO) *OpenAgoalLatestProgressDTO {
	s.Creator = v
	return s
}

func (s *OpenAgoalLatestProgressDTO) SetHtmldescription(v string) *OpenAgoalLatestProgressDTO {
	s.Htmldescription = &v
	return s
}

func (s *OpenAgoalLatestProgressDTO) SetProgressId(v string) *OpenAgoalLatestProgressDTO {
	s.ProgressId = &v
	return s
}

type OpenAgoalObjectiveDTO struct {
	// This parameter is required.
	Executor *OpenAgoalUserDTO `json:"executor,omitempty" xml:"executor,omitempty"`
	// This parameter is required.
	KeyActions []*OpenAgoalKeyActionDTO `json:"keyActions,omitempty" xml:"keyActions,omitempty" type:"Repeated"`
	// This parameter is required.
	KeyResults []*OpenAgoalKeyResultDTO `json:"keyResults,omitempty" xml:"keyResults,omitempty" type:"Repeated"`
	// This parameter is required.
	LatestProgress *OpenAgoalLatestProgressDTO `json:"latestProgress,omitempty" xml:"latestProgress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	// This parameter is required.
	ObjectiveRule *OpenOrgObjectiveRuleDTO `json:"objectiveRule,omitempty" xml:"objectiveRule,omitempty"`
	// This parameter is required.
	Period *OpenObjectiveRulePeriodDTO `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Progress *int32 `json:"progress,omitempty" xml:"progress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	Teams []*OpenAgoalTeamDTO `json:"teams,omitempty" xml:"teams,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 测试目标
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	Weight *float64 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s OpenAgoalObjectiveDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalObjectiveDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalObjectiveDTO) SetExecutor(v *OpenAgoalUserDTO) *OpenAgoalObjectiveDTO {
	s.Executor = v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetKeyActions(v []*OpenAgoalKeyActionDTO) *OpenAgoalObjectiveDTO {
	s.KeyActions = v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetKeyResults(v []*OpenAgoalKeyResultDTO) *OpenAgoalObjectiveDTO {
	s.KeyResults = v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetLatestProgress(v *OpenAgoalLatestProgressDTO) *OpenAgoalObjectiveDTO {
	s.LatestProgress = v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetObjectiveId(v string) *OpenAgoalObjectiveDTO {
	s.ObjectiveId = &v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetObjectiveRule(v *OpenOrgObjectiveRuleDTO) *OpenAgoalObjectiveDTO {
	s.ObjectiveRule = v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetPeriod(v *OpenObjectiveRulePeriodDTO) *OpenAgoalObjectiveDTO {
	s.Period = v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetProgress(v int32) *OpenAgoalObjectiveDTO {
	s.Progress = &v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetStatus(v int32) *OpenAgoalObjectiveDTO {
	s.Status = &v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetTeams(v []*OpenAgoalTeamDTO) *OpenAgoalObjectiveDTO {
	s.Teams = v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetTitle(v string) *OpenAgoalObjectiveDTO {
	s.Title = &v
	return s
}

func (s *OpenAgoalObjectiveDTO) SetWeight(v float64) *OpenAgoalObjectiveDTO {
	s.Weight = &v
	return s
}

type OpenAgoalObjectiveDimensionDTO struct {
	// This parameter is required.
	Children []*OpenAgoalObjectiveDimensionDTO `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 662e006fe4b0f579bbcxxxxx
	DimensionId *string `json:"dimensionId,omitempty" xml:"dimensionId,omitempty"`
	// This parameter is required.
	FieldConfig []*OpenAgoalFieldMetaDTO `json:"fieldConfig,omitempty" xml:"fieldConfig,omitempty" type:"Repeated"`
	// This parameter is required.
	FieldValueMap map[string]interface{} `json:"fieldValueMap,omitempty" xml:"fieldValueMap,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 这是维度标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s OpenAgoalObjectiveDimensionDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalObjectiveDimensionDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalObjectiveDimensionDTO) SetChildren(v []*OpenAgoalObjectiveDimensionDTO) *OpenAgoalObjectiveDimensionDTO {
	s.Children = v
	return s
}

func (s *OpenAgoalObjectiveDimensionDTO) SetDimensionId(v string) *OpenAgoalObjectiveDimensionDTO {
	s.DimensionId = &v
	return s
}

func (s *OpenAgoalObjectiveDimensionDTO) SetFieldConfig(v []*OpenAgoalFieldMetaDTO) *OpenAgoalObjectiveDimensionDTO {
	s.FieldConfig = v
	return s
}

func (s *OpenAgoalObjectiveDimensionDTO) SetFieldValueMap(v map[string]interface{}) *OpenAgoalObjectiveDimensionDTO {
	s.FieldValueMap = v
	return s
}

func (s *OpenAgoalObjectiveDimensionDTO) SetTitle(v string) *OpenAgoalObjectiveDimensionDTO {
	s.Title = &v
	return s
}

type OpenAgoalOrgObjectiveDTO struct {
	// This parameter is required.
	Dimension *OpenAgoalObjectiveDimensionDTO `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// This parameter is required.
	DownAlignObjects []*OpenAgoalAlignDTO `json:"downAlignObjects,omitempty" xml:"downAlignObjects,omitempty" type:"Repeated"`
	// This parameter is required.
	Executor *OpenAgoalUserDTO `json:"executor,omitempty" xml:"executor,omitempty"`
	// This parameter is required.
	FieldConfig []*OpenAgoalFieldMetaDTO `json:"fieldConfig,omitempty" xml:"fieldConfig,omitempty" type:"Repeated"`
	// This parameter is required.
	FieldValueMap map[string]interface{} `json:"fieldValueMap,omitempty" xml:"fieldValueMap,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	// This parameter is required.
	Period *OpenObjectiveRulePeriodDTO `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// formalEffective
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	Team *OpenAgoalTeamDTO `json:"team,omitempty" xml:"team,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试目标
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	UpAlignObjects []*OpenAgoalAlignDTO `json:"upAlignObjects,omitempty" xml:"upAlignObjects,omitempty" type:"Repeated"`
}

func (s OpenAgoalOrgObjectiveDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalOrgObjectiveDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalOrgObjectiveDTO) SetDimension(v *OpenAgoalObjectiveDimensionDTO) *OpenAgoalOrgObjectiveDTO {
	s.Dimension = v
	return s
}

func (s *OpenAgoalOrgObjectiveDTO) SetDownAlignObjects(v []*OpenAgoalAlignDTO) *OpenAgoalOrgObjectiveDTO {
	s.DownAlignObjects = v
	return s
}

func (s *OpenAgoalOrgObjectiveDTO) SetExecutor(v *OpenAgoalUserDTO) *OpenAgoalOrgObjectiveDTO {
	s.Executor = v
	return s
}

func (s *OpenAgoalOrgObjectiveDTO) SetFieldConfig(v []*OpenAgoalFieldMetaDTO) *OpenAgoalOrgObjectiveDTO {
	s.FieldConfig = v
	return s
}

func (s *OpenAgoalOrgObjectiveDTO) SetFieldValueMap(v map[string]interface{}) *OpenAgoalOrgObjectiveDTO {
	s.FieldValueMap = v
	return s
}

func (s *OpenAgoalOrgObjectiveDTO) SetObjectiveId(v string) *OpenAgoalOrgObjectiveDTO {
	s.ObjectiveId = &v
	return s
}

func (s *OpenAgoalOrgObjectiveDTO) SetPeriod(v *OpenObjectiveRulePeriodDTO) *OpenAgoalOrgObjectiveDTO {
	s.Period = v
	return s
}

func (s *OpenAgoalOrgObjectiveDTO) SetStatus(v string) *OpenAgoalOrgObjectiveDTO {
	s.Status = &v
	return s
}

func (s *OpenAgoalOrgObjectiveDTO) SetTeam(v *OpenAgoalTeamDTO) *OpenAgoalOrgObjectiveDTO {
	s.Team = v
	return s
}

func (s *OpenAgoalOrgObjectiveDTO) SetTitle(v string) *OpenAgoalOrgObjectiveDTO {
	s.Title = &v
	return s
}

func (s *OpenAgoalOrgObjectiveDTO) SetUpAlignObjects(v []*OpenAgoalAlignDTO) *OpenAgoalOrgObjectiveDTO {
	s.UpAlignObjects = v
	return s
}

type OpenAgoalOrgObjectiveListDTO struct {
	// This parameter is required.
	ObjectiveList []*OpenAgoalOrgObjectiveDTO `json:"objectiveList,omitempty" xml:"objectiveList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s OpenAgoalOrgObjectiveListDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalOrgObjectiveListDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalOrgObjectiveListDTO) SetObjectiveList(v []*OpenAgoalOrgObjectiveDTO) *OpenAgoalOrgObjectiveListDTO {
	s.ObjectiveList = v
	return s
}

func (s *OpenAgoalOrgObjectiveListDTO) SetTotalCount(v int64) *OpenAgoalOrgObjectiveListDTO {
	s.TotalCount = &v
	return s
}

type OpenAgoalPeriodDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// 1743436799000
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024年度
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// season
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1711900800000
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s OpenAgoalPeriodDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalPeriodDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalPeriodDTO) SetEndDate(v int64) *OpenAgoalPeriodDTO {
	s.EndDate = &v
	return s
}

func (s *OpenAgoalPeriodDTO) SetName(v string) *OpenAgoalPeriodDTO {
	s.Name = &v
	return s
}

func (s *OpenAgoalPeriodDTO) SetPeriodId(v string) *OpenAgoalPeriodDTO {
	s.PeriodId = &v
	return s
}

func (s *OpenAgoalPeriodDTO) SetPeriodType(v string) *OpenAgoalPeriodDTO {
	s.PeriodType = &v
	return s
}

func (s *OpenAgoalPeriodDTO) SetStartDate(v int64) *OpenAgoalPeriodDTO {
	s.StartDate = &v
	return s
}

type OpenAgoalProgressDTO struct {
	// This parameter is required.
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// This parameter is required.
	Creator *OpenAgoalUserDTO `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	HtmlContent *string `json:"htmlContent,omitempty" xml:"htmlContent,omitempty"`
	// This parameter is required.
	Modifier *OpenAgoalUserDTO `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// This parameter is required.
	ProgressId *string `json:"progressId,omitempty" xml:"progressId,omitempty"`
	// This parameter is required.
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s OpenAgoalProgressDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalProgressDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalProgressDTO) SetCreated(v int64) *OpenAgoalProgressDTO {
	s.Created = &v
	return s
}

func (s *OpenAgoalProgressDTO) SetCreator(v *OpenAgoalUserDTO) *OpenAgoalProgressDTO {
	s.Creator = v
	return s
}

func (s *OpenAgoalProgressDTO) SetHtmlContent(v string) *OpenAgoalProgressDTO {
	s.HtmlContent = &v
	return s
}

func (s *OpenAgoalProgressDTO) SetModifier(v *OpenAgoalUserDTO) *OpenAgoalProgressDTO {
	s.Modifier = v
	return s
}

func (s *OpenAgoalProgressDTO) SetProgressId(v string) *OpenAgoalProgressDTO {
	s.ProgressId = &v
	return s
}

func (s *OpenAgoalProgressDTO) SetUpdated(v int64) *OpenAgoalProgressDTO {
	s.Updated = &v
	return s
}

type OpenAgoalTeamDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// 8535683xx
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试部门
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
}

func (s OpenAgoalTeamDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalTeamDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalTeamDTO) SetDeptId(v string) *OpenAgoalTeamDTO {
	s.DeptId = &v
	return s
}

func (s *OpenAgoalTeamDTO) SetName(v string) *OpenAgoalTeamDTO {
	s.Name = &v
	return s
}

func (s *OpenAgoalTeamDTO) SetTeamId(v string) *OpenAgoalTeamDTO {
	s.TeamId = &v
	return s
}

type OpenAgoalUserDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// 211042291978xxxx
	DingUserId *string `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OpenAgoalUserDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenAgoalUserDTO) GoString() string {
	return s.String()
}

func (s *OpenAgoalUserDTO) SetDingUserId(v string) *OpenAgoalUserDTO {
	s.DingUserId = &v
	return s
}

func (s *OpenAgoalUserDTO) SetName(v string) *OpenAgoalUserDTO {
	s.Name = &v
	return s
}

func (s *OpenAgoalUserDTO) SetUserId(v string) *OpenAgoalUserDTO {
	s.UserId = &v
	return s
}

type OpenObjectiveRulePeriodDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// 1743436799000
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024年度
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// season
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1711900800000
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s OpenObjectiveRulePeriodDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenObjectiveRulePeriodDTO) GoString() string {
	return s.String()
}

func (s *OpenObjectiveRulePeriodDTO) SetEndDate(v int64) *OpenObjectiveRulePeriodDTO {
	s.EndDate = &v
	return s
}

func (s *OpenObjectiveRulePeriodDTO) SetName(v string) *OpenObjectiveRulePeriodDTO {
	s.Name = &v
	return s
}

func (s *OpenObjectiveRulePeriodDTO) SetPeriodId(v string) *OpenObjectiveRulePeriodDTO {
	s.PeriodId = &v
	return s
}

func (s *OpenObjectiveRulePeriodDTO) SetPeriodType(v string) *OpenObjectiveRulePeriodDTO {
	s.PeriodType = &v
	return s
}

func (s *OpenObjectiveRulePeriodDTO) SetStartDate(v int64) *OpenObjectiveRulePeriodDTO {
	s.StartDate = &v
	return s
}

type OpenOrgObjectiveRuleDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// OKR / PBC
	ObjectiveCategory *string `json:"objectiveCategory,omitempty" xml:"objectiveCategory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	ObjectiveRuleId *string `json:"objectiveRuleId,omitempty" xml:"objectiveRuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试规则
	ObjectiveRuleName *string `json:"objectiveRuleName,omitempty" xml:"objectiveRuleName,omitempty"`
}

func (s OpenOrgObjectiveRuleDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenOrgObjectiveRuleDTO) GoString() string {
	return s.String()
}

func (s *OpenOrgObjectiveRuleDTO) SetObjectiveCategory(v string) *OpenOrgObjectiveRuleDTO {
	s.ObjectiveCategory = &v
	return s
}

func (s *OpenOrgObjectiveRuleDTO) SetObjectiveRuleId(v string) *OpenOrgObjectiveRuleDTO {
	s.ObjectiveRuleId = &v
	return s
}

func (s *OpenOrgObjectiveRuleDTO) SetObjectiveRuleName(v string) *OpenOrgObjectiveRuleDTO {
	s.ObjectiveRuleName = &v
	return s
}

type OpenUserAdminDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxe3d8c283bb4aa39a90f97fcb1e09
	DingCorpId *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 211042291978xxxx
	DingUserId *string `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
}

func (s OpenUserAdminDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenUserAdminDTO) GoString() string {
	return s.String()
}

func (s *OpenUserAdminDTO) SetDingCorpId(v string) *OpenUserAdminDTO {
	s.DingCorpId = &v
	return s
}

func (s *OpenUserAdminDTO) SetDingUserId(v string) *OpenUserAdminDTO {
	s.DingUserId = &v
	return s
}

type OpenUserSubAdminDTO struct {
	// This parameter is required.
	DeptIds []*string `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxe3d8c283bb4aa39a90f97fcb1e09
	DingCorpId *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 211042291978xxxx
	DingUserId *string `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
	// This parameter is required.
	PermissionGroupCodes []*string `json:"permissionGroupCodes,omitempty" xml:"permissionGroupCodes,omitempty" type:"Repeated"`
}

func (s OpenUserSubAdminDTO) String() string {
	return tea.Prettify(s)
}

func (s OpenUserSubAdminDTO) GoString() string {
	return s.String()
}

func (s *OpenUserSubAdminDTO) SetDeptIds(v []*string) *OpenUserSubAdminDTO {
	s.DeptIds = v
	return s
}

func (s *OpenUserSubAdminDTO) SetDingCorpId(v string) *OpenUserSubAdminDTO {
	s.DingCorpId = &v
	return s
}

func (s *OpenUserSubAdminDTO) SetDingUserId(v string) *OpenUserSubAdminDTO {
	s.DingUserId = &v
	return s
}

func (s *OpenUserSubAdminDTO) SetPermissionGroupCodes(v []*string) *OpenUserSubAdminDTO {
	s.PermissionGroupCodes = v
	return s
}

type PerfTask struct {
	// example:
	//
	// 328497234
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// y/n
	IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// example:
	//
	// ONGOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// xxx考核任务
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 23223423
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PerfTask) String() string {
	return tea.Prettify(s)
}

func (s PerfTask) GoString() string {
	return s.String()
}

func (s *PerfTask) SetId(v string) *PerfTask {
	s.Id = &v
	return s
}

func (s *PerfTask) SetIsDeleted(v string) *PerfTask {
	s.IsDeleted = &v
	return s
}

func (s *PerfTask) SetStatus(v string) *PerfTask {
	s.Status = &v
	return s
}

func (s *PerfTask) SetTitle(v string) *PerfTask {
	s.Title = &v
	return s
}

func (s *PerfTask) SetUserId(v string) *PerfTask {
	s.UserId = &v
	return s
}

type TitleMention struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Length *int32 `json:"length,omitempty" xml:"length,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// This parameter is required.
	User *OpenAgoalUserDTO `json:"user,omitempty" xml:"user,omitempty"`
}

func (s TitleMention) String() string {
	return tea.Prettify(s)
}

func (s TitleMention) GoString() string {
	return s.String()
}

func (s *TitleMention) SetLength(v int32) *TitleMention {
	s.Length = &v
	return s
}

func (s *TitleMention) SetOffset(v int32) *TitleMention {
	s.Offset = &v
	return s
}

func (s *TitleMention) SetUser(v *OpenAgoalUserDTO) *TitleMention {
	s.User = v
	return s
}

type AgoalCreateProgressHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalCreateProgressHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalCreateProgressHeaders) GoString() string {
	return s.String()
}

func (s *AgoalCreateProgressHeaders) SetCommonHeaders(v map[string]*string) *AgoalCreateProgressHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalCreateProgressHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalCreateProgressHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalCreateProgressRequest struct {
	// example:
	//
	// 64bf87f8d7ace3616f0a1971
	KrId *string `json:"krId,omitempty" xml:"krId,omitempty"`
	// if can be null:
	// false
	MergeIntoLatestProgress *bool `json:"mergeIntoLatestProgress,omitempty" xml:"mergeIntoLatestProgress,omitempty"`
	// example:
	//
	// 662e006fe4b0f579bbcb10cf
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	// example:
	//
	// 这是一条目标进展文本
	PlainText *string `json:"plainText,omitempty" xml:"plainText,omitempty"`
	// example:
	//
	// 30
	Progress *int32 `json:"progress,omitempty" xml:"progress,omitempty"`
	// example:
	//
	// naturalWeek
	ProgressMergePeriod *string `json:"progressMergePeriod,omitempty" xml:"progressMergePeriod,omitempty"`
}

func (s AgoalCreateProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalCreateProgressRequest) GoString() string {
	return s.String()
}

func (s *AgoalCreateProgressRequest) SetKrId(v string) *AgoalCreateProgressRequest {
	s.KrId = &v
	return s
}

func (s *AgoalCreateProgressRequest) SetMergeIntoLatestProgress(v bool) *AgoalCreateProgressRequest {
	s.MergeIntoLatestProgress = &v
	return s
}

func (s *AgoalCreateProgressRequest) SetObjectiveId(v string) *AgoalCreateProgressRequest {
	s.ObjectiveId = &v
	return s
}

func (s *AgoalCreateProgressRequest) SetPlainText(v string) *AgoalCreateProgressRequest {
	s.PlainText = &v
	return s
}

func (s *AgoalCreateProgressRequest) SetProgress(v int32) *AgoalCreateProgressRequest {
	s.Progress = &v
	return s
}

func (s *AgoalCreateProgressRequest) SetProgressMergePeriod(v string) *AgoalCreateProgressRequest {
	s.ProgressMergePeriod = &v
	return s
}

type AgoalCreateProgressResponseBody struct {
	// This parameter is required.
	Content   *OpenAgoalProgressDTO `json:"content,omitempty" xml:"content,omitempty"`
	RequestId *string               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalCreateProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalCreateProgressResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalCreateProgressResponseBody) SetContent(v *OpenAgoalProgressDTO) *AgoalCreateProgressResponseBody {
	s.Content = v
	return s
}

func (s *AgoalCreateProgressResponseBody) SetRequestId(v string) *AgoalCreateProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalCreateProgressResponseBody) SetSuccess(v bool) *AgoalCreateProgressResponseBody {
	s.Success = &v
	return s
}

type AgoalCreateProgressResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalCreateProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalCreateProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalCreateProgressResponse) GoString() string {
	return s.String()
}

func (s *AgoalCreateProgressResponse) SetHeaders(v map[string]*string) *AgoalCreateProgressResponse {
	s.Headers = v
	return s
}

func (s *AgoalCreateProgressResponse) SetStatusCode(v int32) *AgoalCreateProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalCreateProgressResponse) SetBody(v *AgoalCreateProgressResponseBody) *AgoalCreateProgressResponse {
	s.Body = v
	return s
}

type AgoalEntityCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalEntityCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalEntityCreateHeaders) GoString() string {
	return s.String()
}

func (s *AgoalEntityCreateHeaders) SetCommonHeaders(v map[string]*string) *AgoalEntityCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalEntityCreateHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalEntityCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalEntityCreateRequest struct {
	Body []*Entity `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s AgoalEntityCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalEntityCreateRequest) GoString() string {
	return s.String()
}

func (s *AgoalEntityCreateRequest) SetBody(v []*Entity) *AgoalEntityCreateRequest {
	s.Body = v
	return s
}

type AgoalEntityCreateResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalEntityCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalEntityCreateResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalEntityCreateResponseBody) SetResult(v bool) *AgoalEntityCreateResponseBody {
	s.Result = &v
	return s
}

func (s *AgoalEntityCreateResponseBody) SetSuccess(v bool) *AgoalEntityCreateResponseBody {
	s.Success = &v
	return s
}

type AgoalEntityCreateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalEntityCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalEntityCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalEntityCreateResponse) GoString() string {
	return s.String()
}

func (s *AgoalEntityCreateResponse) SetHeaders(v map[string]*string) *AgoalEntityCreateResponse {
	s.Headers = v
	return s
}

func (s *AgoalEntityCreateResponse) SetStatusCode(v int32) *AgoalEntityCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalEntityCreateResponse) SetBody(v *AgoalEntityCreateResponseBody) *AgoalEntityCreateResponse {
	s.Body = v
	return s
}

type AgoalEntityUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalEntityUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalEntityUpdateHeaders) GoString() string {
	return s.String()
}

func (s *AgoalEntityUpdateHeaders) SetCommonHeaders(v map[string]*string) *AgoalEntityUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalEntityUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalEntityUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalEntityUpdateRequest struct {
	Body []*Entity `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s AgoalEntityUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalEntityUpdateRequest) GoString() string {
	return s.String()
}

func (s *AgoalEntityUpdateRequest) SetBody(v []*Entity) *AgoalEntityUpdateRequest {
	s.Body = v
	return s
}

type AgoalEntityUpdateResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalEntityUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalEntityUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalEntityUpdateResponseBody) SetResult(v bool) *AgoalEntityUpdateResponseBody {
	s.Result = &v
	return s
}

func (s *AgoalEntityUpdateResponseBody) SetSuccess(v bool) *AgoalEntityUpdateResponseBody {
	s.Success = &v
	return s
}

type AgoalEntityUpdateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalEntityUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalEntityUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalEntityUpdateResponse) GoString() string {
	return s.String()
}

func (s *AgoalEntityUpdateResponse) SetHeaders(v map[string]*string) *AgoalEntityUpdateResponse {
	s.Headers = v
	return s
}

func (s *AgoalEntityUpdateResponse) SetStatusCode(v int32) *AgoalEntityUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalEntityUpdateResponse) SetBody(v *AgoalEntityUpdateResponseBody) *AgoalEntityUpdateResponse {
	s.Body = v
	return s
}

type AgoalFieldUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalFieldUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalFieldUpdateHeaders) GoString() string {
	return s.String()
}

func (s *AgoalFieldUpdateHeaders) SetCommonHeaders(v map[string]*string) *AgoalFieldUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalFieldUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalFieldUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalFieldUpdateRequest struct {
	Body *AgoalFieldUpdateRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s AgoalFieldUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalFieldUpdateRequest) GoString() string {
	return s.String()
}

func (s *AgoalFieldUpdateRequest) SetBody(v *AgoalFieldUpdateRequestBody) *AgoalFieldUpdateRequest {
	s.Body = v
	return s
}

type AgoalFieldUpdateRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 662e006fe4b0f579bbcxxxxx
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OBJECTIVE
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// title
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// 字段值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AgoalFieldUpdateRequestBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalFieldUpdateRequestBody) GoString() string {
	return s.String()
}

func (s *AgoalFieldUpdateRequestBody) SetEntityId(v string) *AgoalFieldUpdateRequestBody {
	s.EntityId = &v
	return s
}

func (s *AgoalFieldUpdateRequestBody) SetEntityType(v string) *AgoalFieldUpdateRequestBody {
	s.EntityType = &v
	return s
}

func (s *AgoalFieldUpdateRequestBody) SetFieldCode(v string) *AgoalFieldUpdateRequestBody {
	s.FieldCode = &v
	return s
}

func (s *AgoalFieldUpdateRequestBody) SetValue(v string) *AgoalFieldUpdateRequestBody {
	s.Value = &v
	return s
}

type AgoalFieldUpdateShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalFieldUpdateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalFieldUpdateShrinkRequest) GoString() string {
	return s.String()
}

func (s *AgoalFieldUpdateShrinkRequest) SetBodyShrink(v string) *AgoalFieldUpdateShrinkRequest {
	s.BodyShrink = &v
	return s
}

type AgoalFieldUpdateResponseBody struct {
	Content   *bool   `json:"content,omitempty" xml:"content,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalFieldUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalFieldUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalFieldUpdateResponseBody) SetContent(v bool) *AgoalFieldUpdateResponseBody {
	s.Content = &v
	return s
}

func (s *AgoalFieldUpdateResponseBody) SetRequestId(v string) *AgoalFieldUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalFieldUpdateResponseBody) SetSuccess(v string) *AgoalFieldUpdateResponseBody {
	s.Success = &v
	return s
}

type AgoalFieldUpdateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalFieldUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalFieldUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalFieldUpdateResponse) GoString() string {
	return s.String()
}

func (s *AgoalFieldUpdateResponse) SetHeaders(v map[string]*string) *AgoalFieldUpdateResponse {
	s.Headers = v
	return s
}

func (s *AgoalFieldUpdateResponse) SetStatusCode(v int32) *AgoalFieldUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalFieldUpdateResponse) SetBody(v *AgoalFieldUpdateResponseBody) *AgoalFieldUpdateResponse {
	s.Body = v
	return s
}

type AgoalObjectiveKeyActionListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalObjectiveKeyActionListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalObjectiveKeyActionListHeaders) GoString() string {
	return s.String()
}

func (s *AgoalObjectiveKeyActionListHeaders) SetCommonHeaders(v map[string]*string) *AgoalObjectiveKeyActionListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalObjectiveKeyActionListHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalObjectiveKeyActionListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalObjectiveKeyActionListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 211042291978xxxx
	DingUserId *string `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	KeyResultId *string `json:"keyResultId,omitempty" xml:"keyResultId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
}

func (s AgoalObjectiveKeyActionListRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalObjectiveKeyActionListRequest) GoString() string {
	return s.String()
}

func (s *AgoalObjectiveKeyActionListRequest) SetDingUserId(v string) *AgoalObjectiveKeyActionListRequest {
	s.DingUserId = &v
	return s
}

func (s *AgoalObjectiveKeyActionListRequest) SetKeyResultId(v string) *AgoalObjectiveKeyActionListRequest {
	s.KeyResultId = &v
	return s
}

func (s *AgoalObjectiveKeyActionListRequest) SetObjectiveId(v string) *AgoalObjectiveKeyActionListRequest {
	s.ObjectiveId = &v
	return s
}

type AgoalObjectiveKeyActionListResponseBody struct {
	// This parameter is required.
	Content []*OpenAgoalKeyActionDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// example:
	//
	// 7478B23C-80E8-1AD6-BE8C-09D480E0xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalObjectiveKeyActionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalObjectiveKeyActionListResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalObjectiveKeyActionListResponseBody) SetContent(v []*OpenAgoalKeyActionDTO) *AgoalObjectiveKeyActionListResponseBody {
	s.Content = v
	return s
}

func (s *AgoalObjectiveKeyActionListResponseBody) SetRequestId(v string) *AgoalObjectiveKeyActionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalObjectiveKeyActionListResponseBody) SetSuccess(v bool) *AgoalObjectiveKeyActionListResponseBody {
	s.Success = &v
	return s
}

type AgoalObjectiveKeyActionListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalObjectiveKeyActionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalObjectiveKeyActionListResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalObjectiveKeyActionListResponse) GoString() string {
	return s.String()
}

func (s *AgoalObjectiveKeyActionListResponse) SetHeaders(v map[string]*string) *AgoalObjectiveKeyActionListResponse {
	s.Headers = v
	return s
}

func (s *AgoalObjectiveKeyActionListResponse) SetStatusCode(v int32) *AgoalObjectiveKeyActionListResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalObjectiveKeyActionListResponse) SetBody(v *AgoalObjectiveKeyActionListResponseBody) *AgoalObjectiveKeyActionListResponse {
	s.Body = v
	return s
}

type AgoalObjectiveRulePeriodListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalObjectiveRulePeriodListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalObjectiveRulePeriodListHeaders) GoString() string {
	return s.String()
}

func (s *AgoalObjectiveRulePeriodListHeaders) SetCommonHeaders(v map[string]*string) *AgoalObjectiveRulePeriodListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalObjectiveRulePeriodListHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalObjectiveRulePeriodListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalObjectiveRulePeriodListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	ObjectiveRuleId *string `json:"objectiveRuleId,omitempty" xml:"objectiveRuleId,omitempty"`
}

func (s AgoalObjectiveRulePeriodListRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalObjectiveRulePeriodListRequest) GoString() string {
	return s.String()
}

func (s *AgoalObjectiveRulePeriodListRequest) SetObjectiveRuleId(v string) *AgoalObjectiveRulePeriodListRequest {
	s.ObjectiveRuleId = &v
	return s
}

type AgoalObjectiveRulePeriodListResponseBody struct {
	// This parameter is required.
	Content []*OpenObjectiveRulePeriodDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// example:
	//
	// 7478B23C-80E8-1AD6-BE8C-09D480E0xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalObjectiveRulePeriodListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalObjectiveRulePeriodListResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalObjectiveRulePeriodListResponseBody) SetContent(v []*OpenObjectiveRulePeriodDTO) *AgoalObjectiveRulePeriodListResponseBody {
	s.Content = v
	return s
}

func (s *AgoalObjectiveRulePeriodListResponseBody) SetRequestId(v string) *AgoalObjectiveRulePeriodListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalObjectiveRulePeriodListResponseBody) SetSuccess(v bool) *AgoalObjectiveRulePeriodListResponseBody {
	s.Success = &v
	return s
}

type AgoalObjectiveRulePeriodListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalObjectiveRulePeriodListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalObjectiveRulePeriodListResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalObjectiveRulePeriodListResponse) GoString() string {
	return s.String()
}

func (s *AgoalObjectiveRulePeriodListResponse) SetHeaders(v map[string]*string) *AgoalObjectiveRulePeriodListResponse {
	s.Headers = v
	return s
}

func (s *AgoalObjectiveRulePeriodListResponse) SetStatusCode(v int32) *AgoalObjectiveRulePeriodListResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalObjectiveRulePeriodListResponse) SetBody(v *AgoalObjectiveRulePeriodListResponseBody) *AgoalObjectiveRulePeriodListResponse {
	s.Body = v
	return s
}

type AgoalOrgObjectiveListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalOrgObjectiveListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveListHeaders) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveListHeaders) SetCommonHeaders(v map[string]*string) *AgoalOrgObjectiveListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalOrgObjectiveListHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalOrgObjectiveListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalOrgObjectiveListRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// 853530516
	DingTeamId *string `json:"dingTeamId,omitempty" xml:"dingTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 662e006fe4b0f579bbcxxxxx
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
}

func (s AgoalOrgObjectiveListRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveListRequest) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveListRequest) SetDingTeamId(v string) *AgoalOrgObjectiveListRequest {
	s.DingTeamId = &v
	return s
}

func (s *AgoalOrgObjectiveListRequest) SetPageNumber(v int32) *AgoalOrgObjectiveListRequest {
	s.PageNumber = &v
	return s
}

func (s *AgoalOrgObjectiveListRequest) SetPageSize(v int32) *AgoalOrgObjectiveListRequest {
	s.PageSize = &v
	return s
}

func (s *AgoalOrgObjectiveListRequest) SetPeriodId(v string) *AgoalOrgObjectiveListRequest {
	s.PeriodId = &v
	return s
}

type AgoalOrgObjectiveListResponseBody struct {
	Content   *OpenAgoalOrgObjectiveListDTO `json:"content,omitempty" xml:"content,omitempty"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalOrgObjectiveListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveListResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveListResponseBody) SetContent(v *OpenAgoalOrgObjectiveListDTO) *AgoalOrgObjectiveListResponseBody {
	s.Content = v
	return s
}

func (s *AgoalOrgObjectiveListResponseBody) SetRequestId(v string) *AgoalOrgObjectiveListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalOrgObjectiveListResponseBody) SetSuccess(v bool) *AgoalOrgObjectiveListResponseBody {
	s.Success = &v
	return s
}

type AgoalOrgObjectiveListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalOrgObjectiveListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalOrgObjectiveListResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveListResponse) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveListResponse) SetHeaders(v map[string]*string) *AgoalOrgObjectiveListResponse {
	s.Headers = v
	return s
}

func (s *AgoalOrgObjectiveListResponse) SetStatusCode(v int32) *AgoalOrgObjectiveListResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalOrgObjectiveListResponse) SetBody(v *AgoalOrgObjectiveListResponseBody) *AgoalOrgObjectiveListResponse {
	s.Body = v
	return s
}

type AgoalOrgObjectiveQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalOrgObjectiveQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveQueryHeaders) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveQueryHeaders) SetCommonHeaders(v map[string]*string) *AgoalOrgObjectiveQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalOrgObjectiveQueryHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalOrgObjectiveQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalOrgObjectiveQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 662e006fe4b0f579bbcxxxxx
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
}

func (s AgoalOrgObjectiveQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveQueryRequest) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveQueryRequest) SetObjectiveId(v string) *AgoalOrgObjectiveQueryRequest {
	s.ObjectiveId = &v
	return s
}

type AgoalOrgObjectiveQueryResponseBody struct {
	Content   *OpenAgoalOrgObjectiveDTO `json:"content,omitempty" xml:"content,omitempty"`
	RequestId *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalOrgObjectiveQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveQueryResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveQueryResponseBody) SetContent(v *OpenAgoalOrgObjectiveDTO) *AgoalOrgObjectiveQueryResponseBody {
	s.Content = v
	return s
}

func (s *AgoalOrgObjectiveQueryResponseBody) SetRequestId(v string) *AgoalOrgObjectiveQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalOrgObjectiveQueryResponseBody) SetSuccess(v bool) *AgoalOrgObjectiveQueryResponseBody {
	s.Success = &v
	return s
}

type AgoalOrgObjectiveQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalOrgObjectiveQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalOrgObjectiveQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveQueryResponse) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveQueryResponse) SetHeaders(v map[string]*string) *AgoalOrgObjectiveQueryResponse {
	s.Headers = v
	return s
}

func (s *AgoalOrgObjectiveQueryResponse) SetStatusCode(v int32) *AgoalOrgObjectiveQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalOrgObjectiveQueryResponse) SetBody(v *AgoalOrgObjectiveQueryResponseBody) *AgoalOrgObjectiveQueryResponse {
	s.Body = v
	return s
}

type AgoalOrgObjectiveRuleListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalOrgObjectiveRuleListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveRuleListHeaders) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveRuleListHeaders) SetCommonHeaders(v map[string]*string) *AgoalOrgObjectiveRuleListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalOrgObjectiveRuleListHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalOrgObjectiveRuleListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalOrgObjectiveRuleListResponseBody struct {
	// This parameter is required.
	Content []*OpenOrgObjectiveRuleDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// example:
	//
	// 7478B23C-80E8-1AD6-BE8C-09D480E0xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalOrgObjectiveRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveRuleListResponseBody) SetContent(v []*OpenOrgObjectiveRuleDTO) *AgoalOrgObjectiveRuleListResponseBody {
	s.Content = v
	return s
}

func (s *AgoalOrgObjectiveRuleListResponseBody) SetRequestId(v string) *AgoalOrgObjectiveRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalOrgObjectiveRuleListResponseBody) SetSuccess(v bool) *AgoalOrgObjectiveRuleListResponseBody {
	s.Success = &v
	return s
}

type AgoalOrgObjectiveRuleListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalOrgObjectiveRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalOrgObjectiveRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalOrgObjectiveRuleListResponse) GoString() string {
	return s.String()
}

func (s *AgoalOrgObjectiveRuleListResponse) SetHeaders(v map[string]*string) *AgoalOrgObjectiveRuleListResponse {
	s.Headers = v
	return s
}

func (s *AgoalOrgObjectiveRuleListResponse) SetStatusCode(v int32) *AgoalOrgObjectiveRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalOrgObjectiveRuleListResponse) SetBody(v *AgoalOrgObjectiveRuleListResponseBody) *AgoalOrgObjectiveRuleListResponse {
	s.Body = v
	return s
}

type AgoalPerfTaskCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalPerfTaskCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalPerfTaskCreateHeaders) GoString() string {
	return s.String()
}

func (s *AgoalPerfTaskCreateHeaders) SetCommonHeaders(v map[string]*string) *AgoalPerfTaskCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalPerfTaskCreateHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalPerfTaskCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalPerfTaskCreateRequest struct {
	Body []*PerfTask `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s AgoalPerfTaskCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalPerfTaskCreateRequest) GoString() string {
	return s.String()
}

func (s *AgoalPerfTaskCreateRequest) SetBody(v []*PerfTask) *AgoalPerfTaskCreateRequest {
	s.Body = v
	return s
}

type AgoalPerfTaskCreateResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalPerfTaskCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalPerfTaskCreateResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalPerfTaskCreateResponseBody) SetResult(v bool) *AgoalPerfTaskCreateResponseBody {
	s.Result = &v
	return s
}

func (s *AgoalPerfTaskCreateResponseBody) SetSuccess(v bool) *AgoalPerfTaskCreateResponseBody {
	s.Success = &v
	return s
}

type AgoalPerfTaskCreateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalPerfTaskCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalPerfTaskCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalPerfTaskCreateResponse) GoString() string {
	return s.String()
}

func (s *AgoalPerfTaskCreateResponse) SetHeaders(v map[string]*string) *AgoalPerfTaskCreateResponse {
	s.Headers = v
	return s
}

func (s *AgoalPerfTaskCreateResponse) SetStatusCode(v int32) *AgoalPerfTaskCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalPerfTaskCreateResponse) SetBody(v *AgoalPerfTaskCreateResponseBody) *AgoalPerfTaskCreateResponse {
	s.Body = v
	return s
}

type AgoalPerfTaskUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalPerfTaskUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalPerfTaskUpdateHeaders) GoString() string {
	return s.String()
}

func (s *AgoalPerfTaskUpdateHeaders) SetCommonHeaders(v map[string]*string) *AgoalPerfTaskUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalPerfTaskUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalPerfTaskUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalPerfTaskUpdateRequest struct {
	Body []*PerfTask `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s AgoalPerfTaskUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalPerfTaskUpdateRequest) GoString() string {
	return s.String()
}

func (s *AgoalPerfTaskUpdateRequest) SetBody(v []*PerfTask) *AgoalPerfTaskUpdateRequest {
	s.Body = v
	return s
}

type AgoalPerfTaskUpdateResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalPerfTaskUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalPerfTaskUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalPerfTaskUpdateResponseBody) SetResult(v bool) *AgoalPerfTaskUpdateResponseBody {
	s.Result = &v
	return s
}

func (s *AgoalPerfTaskUpdateResponseBody) SetSuccess(v bool) *AgoalPerfTaskUpdateResponseBody {
	s.Success = &v
	return s
}

type AgoalPerfTaskUpdateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalPerfTaskUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalPerfTaskUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalPerfTaskUpdateResponse) GoString() string {
	return s.String()
}

func (s *AgoalPerfTaskUpdateResponse) SetHeaders(v map[string]*string) *AgoalPerfTaskUpdateResponse {
	s.Headers = v
	return s
}

func (s *AgoalPerfTaskUpdateResponse) SetStatusCode(v int32) *AgoalPerfTaskUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalPerfTaskUpdateResponse) SetBody(v *AgoalPerfTaskUpdateResponseBody) *AgoalPerfTaskUpdateResponse {
	s.Body = v
	return s
}

type AgoalPeriodListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalPeriodListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalPeriodListHeaders) GoString() string {
	return s.String()
}

func (s *AgoalPeriodListHeaders) SetCommonHeaders(v map[string]*string) *AgoalPeriodListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalPeriodListHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalPeriodListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalPeriodListRequest struct {
	Body *AgoalPeriodListRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s AgoalPeriodListRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalPeriodListRequest) GoString() string {
	return s.String()
}

func (s *AgoalPeriodListRequest) SetBody(v *AgoalPeriodListRequestBody) *AgoalPeriodListRequest {
	s.Body = v
	return s
}

type AgoalPeriodListRequestBody struct {
	PeriodTypes []*string `json:"periodTypes,omitempty" xml:"periodTypes,omitempty" type:"Repeated"`
}

func (s AgoalPeriodListRequestBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalPeriodListRequestBody) GoString() string {
	return s.String()
}

func (s *AgoalPeriodListRequestBody) SetPeriodTypes(v []*string) *AgoalPeriodListRequestBody {
	s.PeriodTypes = v
	return s
}

type AgoalPeriodListShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalPeriodListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalPeriodListShrinkRequest) GoString() string {
	return s.String()
}

func (s *AgoalPeriodListShrinkRequest) SetBodyShrink(v string) *AgoalPeriodListShrinkRequest {
	s.BodyShrink = &v
	return s
}

type AgoalPeriodListResponseBody struct {
	Content   []*OpenAgoalPeriodDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	RequestId *string               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalPeriodListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalPeriodListResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalPeriodListResponseBody) SetContent(v []*OpenAgoalPeriodDTO) *AgoalPeriodListResponseBody {
	s.Content = v
	return s
}

func (s *AgoalPeriodListResponseBody) SetRequestId(v string) *AgoalPeriodListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalPeriodListResponseBody) SetSuccess(v bool) *AgoalPeriodListResponseBody {
	s.Success = &v
	return s
}

type AgoalPeriodListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalPeriodListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalPeriodListResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalPeriodListResponse) GoString() string {
	return s.String()
}

func (s *AgoalPeriodListResponse) SetHeaders(v map[string]*string) *AgoalPeriodListResponse {
	s.Headers = v
	return s
}

func (s *AgoalPeriodListResponse) SetStatusCode(v int32) *AgoalPeriodListResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalPeriodListResponse) SetBody(v *AgoalPeriodListResponseBody) *AgoalPeriodListResponse {
	s.Body = v
	return s
}

type AgoalSendMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalSendMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalSendMessageHeaders) GoString() string {
	return s.String()
}

func (s *AgoalSendMessageHeaders) SetCommonHeaders(v map[string]*string) *AgoalSendMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalSendMessageHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalSendMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalSendMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://agoal.dingtalk.com
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"A":"a", "B":"b"}
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://agoal.dingtalk.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 211042291978xxxx
	SourceDingUserId *string `json:"sourceDingUserId,omitempty" xml:"sourceDingUserId,omitempty"`
	// This parameter is required.
	TargetDingUserIds []*string `json:"targetDingUserIds,omitempty" xml:"targetDingUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1d01a14febc7482ca3b6e1d30cf5xxxx
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s AgoalSendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalSendMessageRequest) GoString() string {
	return s.String()
}

func (s *AgoalSendMessageRequest) SetMobileUrl(v string) *AgoalSendMessageRequest {
	s.MobileUrl = &v
	return s
}

func (s *AgoalSendMessageRequest) SetParams(v string) *AgoalSendMessageRequest {
	s.Params = &v
	return s
}

func (s *AgoalSendMessageRequest) SetPcUrl(v string) *AgoalSendMessageRequest {
	s.PcUrl = &v
	return s
}

func (s *AgoalSendMessageRequest) SetSourceDingUserId(v string) *AgoalSendMessageRequest {
	s.SourceDingUserId = &v
	return s
}

func (s *AgoalSendMessageRequest) SetTargetDingUserIds(v []*string) *AgoalSendMessageRequest {
	s.TargetDingUserIds = v
	return s
}

func (s *AgoalSendMessageRequest) SetTemplateId(v string) *AgoalSendMessageRequest {
	s.TemplateId = &v
	return s
}

type AgoalSendMessageResponseBody struct {
	// This parameter is required.
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7478B23C-80E8-1AD6-BE8C-09D480E0xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalSendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalSendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalSendMessageResponseBody) SetContent(v bool) *AgoalSendMessageResponseBody {
	s.Content = &v
	return s
}

func (s *AgoalSendMessageResponseBody) SetRequestId(v string) *AgoalSendMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalSendMessageResponseBody) SetSuccess(v bool) *AgoalSendMessageResponseBody {
	s.Success = &v
	return s
}

type AgoalSendMessageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalSendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalSendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalSendMessageResponse) GoString() string {
	return s.String()
}

func (s *AgoalSendMessageResponse) SetHeaders(v map[string]*string) *AgoalSendMessageResponse {
	s.Headers = v
	return s
}

func (s *AgoalSendMessageResponse) SetStatusCode(v int32) *AgoalSendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalSendMessageResponse) SetBody(v *AgoalSendMessageResponseBody) *AgoalSendMessageResponse {
	s.Body = v
	return s
}

type AgoalUserAdminListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalUserAdminListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserAdminListHeaders) GoString() string {
	return s.String()
}

func (s *AgoalUserAdminListHeaders) SetCommonHeaders(v map[string]*string) *AgoalUserAdminListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalUserAdminListHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalUserAdminListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalUserAdminListResponseBody struct {
	// This parameter is required.
	Content []*OpenUserAdminDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// example:
	//
	// 7478B23C-80E8-1AD6-BE8C-09D480E0xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalUserAdminListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserAdminListResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalUserAdminListResponseBody) SetContent(v []*OpenUserAdminDTO) *AgoalUserAdminListResponseBody {
	s.Content = v
	return s
}

func (s *AgoalUserAdminListResponseBody) SetRequestId(v string) *AgoalUserAdminListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalUserAdminListResponseBody) SetSuccess(v bool) *AgoalUserAdminListResponseBody {
	s.Success = &v
	return s
}

type AgoalUserAdminListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalUserAdminListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalUserAdminListResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserAdminListResponse) GoString() string {
	return s.String()
}

func (s *AgoalUserAdminListResponse) SetHeaders(v map[string]*string) *AgoalUserAdminListResponse {
	s.Headers = v
	return s
}

func (s *AgoalUserAdminListResponse) SetStatusCode(v int32) *AgoalUserAdminListResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalUserAdminListResponse) SetBody(v *AgoalUserAdminListResponseBody) *AgoalUserAdminListResponse {
	s.Body = v
	return s
}

type AgoalUserObjectiveListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalUserObjectiveListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserObjectiveListHeaders) GoString() string {
	return s.String()
}

func (s *AgoalUserObjectiveListHeaders) SetCommonHeaders(v map[string]*string) *AgoalUserObjectiveListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalUserObjectiveListHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalUserObjectiveListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalUserObjectiveListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 211042291978xxxx
	DingUserId *string `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6444f5e9a4261c6e699dxxxx
	ObjectiveRuleId *string `json:"objectiveRuleId,omitempty" xml:"objectiveRuleId,omitempty"`
	// This parameter is required.
	PeriodIds []*string `json:"periodIds,omitempty" xml:"periodIds,omitempty" type:"Repeated"`
}

func (s AgoalUserObjectiveListRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserObjectiveListRequest) GoString() string {
	return s.String()
}

func (s *AgoalUserObjectiveListRequest) SetDingUserId(v string) *AgoalUserObjectiveListRequest {
	s.DingUserId = &v
	return s
}

func (s *AgoalUserObjectiveListRequest) SetObjectiveRuleId(v string) *AgoalUserObjectiveListRequest {
	s.ObjectiveRuleId = &v
	return s
}

func (s *AgoalUserObjectiveListRequest) SetPeriodIds(v []*string) *AgoalUserObjectiveListRequest {
	s.PeriodIds = v
	return s
}

type AgoalUserObjectiveListResponseBody struct {
	Content   []*OpenAgoalObjectiveDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalUserObjectiveListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserObjectiveListResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalUserObjectiveListResponseBody) SetContent(v []*OpenAgoalObjectiveDTO) *AgoalUserObjectiveListResponseBody {
	s.Content = v
	return s
}

func (s *AgoalUserObjectiveListResponseBody) SetRequestId(v string) *AgoalUserObjectiveListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalUserObjectiveListResponseBody) SetSuccess(v bool) *AgoalUserObjectiveListResponseBody {
	s.Success = &v
	return s
}

type AgoalUserObjectiveListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalUserObjectiveListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalUserObjectiveListResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserObjectiveListResponse) GoString() string {
	return s.String()
}

func (s *AgoalUserObjectiveListResponse) SetHeaders(v map[string]*string) *AgoalUserObjectiveListResponse {
	s.Headers = v
	return s
}

func (s *AgoalUserObjectiveListResponse) SetStatusCode(v int32) *AgoalUserObjectiveListResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalUserObjectiveListResponse) SetBody(v *AgoalUserObjectiveListResponseBody) *AgoalUserObjectiveListResponse {
	s.Body = v
	return s
}

type AgoalUserSubAdminListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AgoalUserSubAdminListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserSubAdminListHeaders) GoString() string {
	return s.String()
}

func (s *AgoalUserSubAdminListHeaders) SetCommonHeaders(v map[string]*string) *AgoalUserSubAdminListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AgoalUserSubAdminListHeaders) SetXAcsDingtalkAccessToken(v string) *AgoalUserSubAdminListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AgoalUserSubAdminListRequest struct {
	// example:
	//
	// ACCOUNT
	FuncPermissionGroup *string `json:"funcPermissionGroup,omitempty" xml:"funcPermissionGroup,omitempty"`
}

func (s AgoalUserSubAdminListRequest) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserSubAdminListRequest) GoString() string {
	return s.String()
}

func (s *AgoalUserSubAdminListRequest) SetFuncPermissionGroup(v string) *AgoalUserSubAdminListRequest {
	s.FuncPermissionGroup = &v
	return s
}

type AgoalUserSubAdminListResponseBody struct {
	Content []*OpenUserSubAdminDTO `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// example:
	//
	// 7478B23C-80E8-1AD6-BE8C-09D480E0xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AgoalUserSubAdminListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserSubAdminListResponseBody) GoString() string {
	return s.String()
}

func (s *AgoalUserSubAdminListResponseBody) SetContent(v []*OpenUserSubAdminDTO) *AgoalUserSubAdminListResponseBody {
	s.Content = v
	return s
}

func (s *AgoalUserSubAdminListResponseBody) SetRequestId(v string) *AgoalUserSubAdminListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgoalUserSubAdminListResponseBody) SetSuccess(v bool) *AgoalUserSubAdminListResponseBody {
	s.Success = &v
	return s
}

type AgoalUserSubAdminListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgoalUserSubAdminListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgoalUserSubAdminListResponse) String() string {
	return tea.Prettify(s)
}

func (s AgoalUserSubAdminListResponse) GoString() string {
	return s.String()
}

func (s *AgoalUserSubAdminListResponse) SetHeaders(v map[string]*string) *AgoalUserSubAdminListResponse {
	s.Headers = v
	return s
}

func (s *AgoalUserSubAdminListResponse) SetStatusCode(v int32) *AgoalUserSubAdminListResponse {
	s.StatusCode = &v
	return s
}

func (s *AgoalUserSubAdminListResponse) SetBody(v *AgoalUserSubAdminListResponseBody) *AgoalUserSubAdminListResponse {
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
// 创建目标进展
//
// @param request - AgoalCreateProgressRequest
//
// @param headers - AgoalCreateProgressHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalCreateProgressResponse
func (client *Client) AgoalCreateProgressWithOptions(request *AgoalCreateProgressRequest, headers *AgoalCreateProgressHeaders, runtime *util.RuntimeOptions) (_result *AgoalCreateProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KrId)) {
		body["krId"] = request.KrId
	}

	if !tea.BoolValue(util.IsUnset(request.MergeIntoLatestProgress)) {
		body["mergeIntoLatestProgress"] = request.MergeIntoLatestProgress
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectiveId)) {
		body["objectiveId"] = request.ObjectiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PlainText)) {
		body["plainText"] = request.PlainText
	}

	if !tea.BoolValue(util.IsUnset(request.Progress)) {
		body["progress"] = request.Progress
	}

	if !tea.BoolValue(util.IsUnset(request.ProgressMergePeriod)) {
		body["progressMergePeriod"] = request.ProgressMergePeriod
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
		Action:      tea.String("AgoalCreateProgress"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/objectives/progresses"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalCreateProgressResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建目标进展
//
// @param request - AgoalCreateProgressRequest
//
// @return AgoalCreateProgressResponse
func (client *Client) AgoalCreateProgress(request *AgoalCreateProgressRequest) (_result *AgoalCreateProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalCreateProgressHeaders{}
	_result = &AgoalCreateProgressResponse{}
	_body, _err := client.AgoalCreateProgressWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建业务实体
//
// @param request - AgoalEntityCreateRequest
//
// @param headers - AgoalEntityCreateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalEntityCreateResponse
func (client *Client) AgoalEntityCreateWithOptions(request *AgoalEntityCreateRequest, headers *AgoalEntityCreateHeaders, runtime *util.RuntimeOptions) (_result *AgoalEntityCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("AgoalEntityCreate"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/entities"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalEntityCreateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建业务实体
//
// @param request - AgoalEntityCreateRequest
//
// @return AgoalEntityCreateResponse
func (client *Client) AgoalEntityCreate(request *AgoalEntityCreateRequest) (_result *AgoalEntityCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalEntityCreateHeaders{}
	_result = &AgoalEntityCreateResponse{}
	_body, _err := client.AgoalEntityCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新业务实体
//
// @param request - AgoalEntityUpdateRequest
//
// @param headers - AgoalEntityUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalEntityUpdateResponse
func (client *Client) AgoalEntityUpdateWithOptions(request *AgoalEntityUpdateRequest, headers *AgoalEntityUpdateHeaders, runtime *util.RuntimeOptions) (_result *AgoalEntityUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("AgoalEntityUpdate"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/entities"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalEntityUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新业务实体
//
// @param request - AgoalEntityUpdateRequest
//
// @return AgoalEntityUpdateResponse
func (client *Client) AgoalEntityUpdate(request *AgoalEntityUpdateRequest) (_result *AgoalEntityUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalEntityUpdateHeaders{}
	_result = &AgoalEntityUpdateResponse{}
	_body, _err := client.AgoalEntityUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 Agoal 字段值
//
// @param tmpReq - AgoalFieldUpdateRequest
//
// @param headers - AgoalFieldUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalFieldUpdateResponse
func (client *Client) AgoalFieldUpdateWithOptions(tmpReq *AgoalFieldUpdateRequest, headers *AgoalFieldUpdateHeaders, runtime *util.RuntimeOptions) (_result *AgoalFieldUpdateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AgoalFieldUpdateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
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
		Action:      tea.String("AgoalFieldUpdate"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/fields"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalFieldUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 Agoal 字段值
//
// @param request - AgoalFieldUpdateRequest
//
// @return AgoalFieldUpdateResponse
func (client *Client) AgoalFieldUpdate(request *AgoalFieldUpdateRequest) (_result *AgoalFieldUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalFieldUpdateHeaders{}
	_result = &AgoalFieldUpdateResponse{}
	_body, _err := client.AgoalFieldUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Agoal指定目标或者关键结果关联的关键行动
//
// @param request - AgoalObjectiveKeyActionListRequest
//
// @param headers - AgoalObjectiveKeyActionListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalObjectiveKeyActionListResponse
func (client *Client) AgoalObjectiveKeyActionListWithOptions(request *AgoalObjectiveKeyActionListRequest, headers *AgoalObjectiveKeyActionListHeaders, runtime *util.RuntimeOptions) (_result *AgoalObjectiveKeyActionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingUserId)) {
		query["dingUserId"] = request.DingUserId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyResultId)) {
		query["keyResultId"] = request.KeyResultId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectiveId)) {
		query["objectiveId"] = request.ObjectiveId
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
		Action:      tea.String("AgoalObjectiveKeyActionList"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/objectives/keyActionLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalObjectiveKeyActionListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Agoal指定目标或者关键结果关联的关键行动
//
// @param request - AgoalObjectiveKeyActionListRequest
//
// @return AgoalObjectiveKeyActionListResponse
func (client *Client) AgoalObjectiveKeyActionList(request *AgoalObjectiveKeyActionListRequest) (_result *AgoalObjectiveKeyActionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalObjectiveKeyActionListHeaders{}
	_result = &AgoalObjectiveKeyActionListResponse{}
	_body, _err := client.AgoalObjectiveKeyActionListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Agoal目标规则下的周期列表
//
// @param request - AgoalObjectiveRulePeriodListRequest
//
// @param headers - AgoalObjectiveRulePeriodListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalObjectiveRulePeriodListResponse
func (client *Client) AgoalObjectiveRulePeriodListWithOptions(request *AgoalObjectiveRulePeriodListRequest, headers *AgoalObjectiveRulePeriodListHeaders, runtime *util.RuntimeOptions) (_result *AgoalObjectiveRulePeriodListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectiveRuleId)) {
		query["objectiveRuleId"] = request.ObjectiveRuleId
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
		Action:      tea.String("AgoalObjectiveRulePeriodList"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/objectiveRules/periodLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalObjectiveRulePeriodListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Agoal目标规则下的周期列表
//
// @param request - AgoalObjectiveRulePeriodListRequest
//
// @return AgoalObjectiveRulePeriodListResponse
func (client *Client) AgoalObjectiveRulePeriodList(request *AgoalObjectiveRulePeriodListRequest) (_result *AgoalObjectiveRulePeriodListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalObjectiveRulePeriodListHeaders{}
	_result = &AgoalObjectiveRulePeriodListResponse{}
	_body, _err := client.AgoalObjectiveRulePeriodListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 Agoal 组织目标列表
//
// @param request - AgoalOrgObjectiveListRequest
//
// @param headers - AgoalOrgObjectiveListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalOrgObjectiveListResponse
func (client *Client) AgoalOrgObjectiveListWithOptions(request *AgoalOrgObjectiveListRequest, headers *AgoalOrgObjectiveListHeaders, runtime *util.RuntimeOptions) (_result *AgoalOrgObjectiveListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingTeamId)) {
		query["dingTeamId"] = request.DingTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		query["periodId"] = request.PeriodId
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
		Action:      tea.String("AgoalOrgObjectiveList"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/orgObjectives/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalOrgObjectiveListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 Agoal 组织目标列表
//
// @param request - AgoalOrgObjectiveListRequest
//
// @return AgoalOrgObjectiveListResponse
func (client *Client) AgoalOrgObjectiveList(request *AgoalOrgObjectiveListRequest) (_result *AgoalOrgObjectiveListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalOrgObjectiveListHeaders{}
	_result = &AgoalOrgObjectiveListResponse{}
	_body, _err := client.AgoalOrgObjectiveListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询组织目标详情
//
// @param request - AgoalOrgObjectiveQueryRequest
//
// @param headers - AgoalOrgObjectiveQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalOrgObjectiveQueryResponse
func (client *Client) AgoalOrgObjectiveQueryWithOptions(request *AgoalOrgObjectiveQueryRequest, headers *AgoalOrgObjectiveQueryHeaders, runtime *util.RuntimeOptions) (_result *AgoalOrgObjectiveQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectiveId)) {
		query["objectiveId"] = request.ObjectiveId
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
		Action:      tea.String("AgoalOrgObjectiveQuery"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/orgObjectives"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalOrgObjectiveQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询组织目标详情
//
// @param request - AgoalOrgObjectiveQueryRequest
//
// @return AgoalOrgObjectiveQueryResponse
func (client *Client) AgoalOrgObjectiveQuery(request *AgoalOrgObjectiveQueryRequest) (_result *AgoalOrgObjectiveQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalOrgObjectiveQueryHeaders{}
	_result = &AgoalOrgObjectiveQueryResponse{}
	_body, _err := client.AgoalOrgObjectiveQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Agoal目标规则列表
//
// @param headers - AgoalOrgObjectiveRuleListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalOrgObjectiveRuleListResponse
func (client *Client) AgoalOrgObjectiveRuleListWithOptions(headers *AgoalOrgObjectiveRuleListHeaders, runtime *util.RuntimeOptions) (_result *AgoalOrgObjectiveRuleListResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("AgoalOrgObjectiveRuleList"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/objectiveRules/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalOrgObjectiveRuleListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Agoal目标规则列表
//
// @return AgoalOrgObjectiveRuleListResponse
func (client *Client) AgoalOrgObjectiveRuleList() (_result *AgoalOrgObjectiveRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalOrgObjectiveRuleListHeaders{}
	_result = &AgoalOrgObjectiveRuleListResponse{}
	_body, _err := client.AgoalOrgObjectiveRuleListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建考核任务
//
// @param request - AgoalPerfTaskCreateRequest
//
// @param headers - AgoalPerfTaskCreateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalPerfTaskCreateResponse
func (client *Client) AgoalPerfTaskCreateWithOptions(request *AgoalPerfTaskCreateRequest, headers *AgoalPerfTaskCreateHeaders, runtime *util.RuntimeOptions) (_result *AgoalPerfTaskCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("AgoalPerfTaskCreate"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/perfTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalPerfTaskCreateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建考核任务
//
// @param request - AgoalPerfTaskCreateRequest
//
// @return AgoalPerfTaskCreateResponse
func (client *Client) AgoalPerfTaskCreate(request *AgoalPerfTaskCreateRequest) (_result *AgoalPerfTaskCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalPerfTaskCreateHeaders{}
	_result = &AgoalPerfTaskCreateResponse{}
	_body, _err := client.AgoalPerfTaskCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新考核任务
//
// @param request - AgoalPerfTaskUpdateRequest
//
// @param headers - AgoalPerfTaskUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalPerfTaskUpdateResponse
func (client *Client) AgoalPerfTaskUpdateWithOptions(request *AgoalPerfTaskUpdateRequest, headers *AgoalPerfTaskUpdateHeaders, runtime *util.RuntimeOptions) (_result *AgoalPerfTaskUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("AgoalPerfTaskUpdate"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/perfTasks"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalPerfTaskUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新考核任务
//
// @param request - AgoalPerfTaskUpdateRequest
//
// @return AgoalPerfTaskUpdateResponse
func (client *Client) AgoalPerfTaskUpdate(request *AgoalPerfTaskUpdateRequest) (_result *AgoalPerfTaskUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalPerfTaskUpdateHeaders{}
	_result = &AgoalPerfTaskUpdateResponse{}
	_body, _err := client.AgoalPerfTaskUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 Agoal 周期列表
//
// @param tmpReq - AgoalPeriodListRequest
//
// @param headers - AgoalPeriodListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalPeriodListResponse
func (client *Client) AgoalPeriodListWithOptions(tmpReq *AgoalPeriodListRequest, headers *AgoalPeriodListHeaders, runtime *util.RuntimeOptions) (_result *AgoalPeriodListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AgoalPeriodListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
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
		Action:      tea.String("AgoalPeriodList"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/periods/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalPeriodListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 Agoal 周期列表
//
// @param request - AgoalPeriodListRequest
//
// @return AgoalPeriodListResponse
func (client *Client) AgoalPeriodList(request *AgoalPeriodListRequest) (_result *AgoalPeriodListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalPeriodListHeaders{}
	_result = &AgoalPeriodListResponse{}
	_body, _err := client.AgoalPeriodListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Agoal消息发送
//
// @param request - AgoalSendMessageRequest
//
// @param headers - AgoalSendMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalSendMessageResponse
func (client *Client) AgoalSendMessageWithOptions(request *AgoalSendMessageRequest, headers *AgoalSendMessageHeaders, runtime *util.RuntimeOptions) (_result *AgoalSendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MobileUrl)) {
		body["mobileUrl"] = request.MobileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.PcUrl)) {
		body["pcUrl"] = request.PcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDingUserId)) {
		body["sourceDingUserId"] = request.SourceDingUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetDingUserIds)) {
		body["targetDingUserIds"] = request.TargetDingUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Action:      tea.String("AgoalSendMessage"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/messages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalSendMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Agoal消息发送
//
// @param request - AgoalSendMessageRequest
//
// @return AgoalSendMessageResponse
func (client *Client) AgoalSendMessage(request *AgoalSendMessageRequest) (_result *AgoalSendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalSendMessageHeaders{}
	_result = &AgoalSendMessageResponse{}
	_body, _err := client.AgoalSendMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Agoal管理员列表
//
// @param headers - AgoalUserAdminListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalUserAdminListResponse
func (client *Client) AgoalUserAdminListWithOptions(headers *AgoalUserAdminListHeaders, runtime *util.RuntimeOptions) (_result *AgoalUserAdminListResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("AgoalUserAdminList"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/administrators/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalUserAdminListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Agoal管理员列表
//
// @return AgoalUserAdminListResponse
func (client *Client) AgoalUserAdminList() (_result *AgoalUserAdminListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalUserAdminListHeaders{}
	_result = &AgoalUserAdminListResponse{}
	_body, _err := client.AgoalUserAdminListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Agoal用户目标列表
//
// @param request - AgoalUserObjectiveListRequest
//
// @param headers - AgoalUserObjectiveListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalUserObjectiveListResponse
func (client *Client) AgoalUserObjectiveListWithOptions(request *AgoalUserObjectiveListRequest, headers *AgoalUserObjectiveListHeaders, runtime *util.RuntimeOptions) (_result *AgoalUserObjectiveListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingUserId)) {
		body["dingUserId"] = request.DingUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectiveRuleId)) {
		body["objectiveRuleId"] = request.ObjectiveRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodIds)) {
		body["periodIds"] = request.PeriodIds
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
		Action:      tea.String("AgoalUserObjectiveList"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/users/objectiveLists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalUserObjectiveListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Agoal用户目标列表
//
// @param request - AgoalUserObjectiveListRequest
//
// @return AgoalUserObjectiveListResponse
func (client *Client) AgoalUserObjectiveList(request *AgoalUserObjectiveListRequest) (_result *AgoalUserObjectiveListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalUserObjectiveListHeaders{}
	_result = &AgoalUserObjectiveListResponse{}
	_body, _err := client.AgoalUserObjectiveListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Agoal子管理员列表
//
// @param request - AgoalUserSubAdminListRequest
//
// @param headers - AgoalUserSubAdminListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgoalUserSubAdminListResponse
func (client *Client) AgoalUserSubAdminListWithOptions(request *AgoalUserSubAdminListRequest, headers *AgoalUserSubAdminListHeaders, runtime *util.RuntimeOptions) (_result *AgoalUserSubAdminListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FuncPermissionGroup)) {
		query["funcPermissionGroup"] = request.FuncPermissionGroup
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
		Action:      tea.String("AgoalUserSubAdminList"),
		Version:     tea.String("agoal_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/agoal/administrators/sub/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AgoalUserSubAdminListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Agoal子管理员列表
//
// @param request - AgoalUserSubAdminListRequest
//
// @return AgoalUserSubAdminListResponse
func (client *Client) AgoalUserSubAdminList(request *AgoalUserSubAdminListRequest) (_result *AgoalUserSubAdminListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AgoalUserSubAdminListHeaders{}
	_result = &AgoalUserSubAdminListResponse{}
	_body, _err := client.AgoalUserSubAdminListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
