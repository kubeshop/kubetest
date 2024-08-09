package testworkflow

import (
	"time"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

var _ Filter = &FilterImpl{}

type FilterImpl struct {
	FName        string
	FLastNDays   int
	FStartDate   *time.Time
	FEndDate     *time.Time
	FStatuses    []testkube.TestWorkflowStatus
	FPage        int
	FPageSize    int
	FTextSearch  string
	FSelector    string
	FTagSelector string
	FRunnerIds   []string
	FTags        map[string]string
	FRunnerTags  map[string]string
}

func NewExecutionsFilter() *FilterImpl {
	result := FilterImpl{FPage: 0, FPageSize: PageDefaultLimit, FRunnerTags: make(map[string]string)}
	return &result
}

func (f *FilterImpl) WithName(name string) *FilterImpl {
	f.FName = name
	return f
}

func (f *FilterImpl) WithLastNDays(days int) *FilterImpl {
	f.FLastNDays = days
	return f
}

func (f *FilterImpl) WithStartDate(date time.Time) *FilterImpl {
	f.FStartDate = &date
	return f
}

func (f *FilterImpl) WithEndDate(date time.Time) *FilterImpl {
	f.FEndDate = &date
	return f
}

func (f *FilterImpl) WithStatus(status string) *FilterImpl {
	statuses, err := testkube.ParseTestWorkflowStatusList(status, ",")
	if err == nil {
		f.FStatuses = statuses
	}
	return f
}

func (f *FilterImpl) WithPage(page int) *FilterImpl {
	f.FPage = page
	return f
}

func (f *FilterImpl) WithPageSize(pageSize int) *FilterImpl {
	f.FPageSize = pageSize
	return f
}

func (f *FilterImpl) WithTextSearch(textSearch string) *FilterImpl {
	f.FTextSearch = textSearch
	return f
}

func (f *FilterImpl) WithSelector(selector string) *FilterImpl {
	f.FSelector = selector
	return f
}

func (f *FilterImpl) WithTagSelector(tagSelector string) *FilterImpl {
	f.FTagSelector = tagSelector
	return f
}

func (f *FilterImpl) WithTags(tags map[string]string) *FilterImpl {
	f.FTags = tags
	return f
}

func (f *FilterImpl) WithRunnerTags(tags map[string]string) *FilterImpl {
	f.FRunnerTags = tags
	return f
}

func (f *FilterImpl) WithRunnerIds(runnerIds []string) *FilterImpl {
	f.FRunnerIds = runnerIds
	return f
}

func (f FilterImpl) Name() string {
	return f.FName
}

func (f FilterImpl) NameDefined() bool {
	return f.FName != ""
}

func (f FilterImpl) LastNDaysDefined() bool {
	return f.FLastNDays > 0
}

func (f FilterImpl) LastNDays() int {
	return f.FLastNDays
}

func (f FilterImpl) StartDateDefined() bool {
	return f.FStartDate != nil
}

func (f FilterImpl) StartDate() time.Time {
	return *f.FStartDate
}

func (f FilterImpl) EndDateDefined() bool {
	return f.FEndDate != nil
}

func (f FilterImpl) EndDate() time.Time {
	return *f.FEndDate
}

func (f FilterImpl) StatusesDefined() bool {
	return len(f.FStatuses) != 0
}

func (f FilterImpl) Statuses() []testkube.TestWorkflowStatus {
	return f.FStatuses
}

func (f FilterImpl) Page() int {
	return f.FPage
}

func (f FilterImpl) PageSize() int {
	return f.FPageSize
}

func (f FilterImpl) TextSearchDefined() bool {
	return f.FTextSearch != ""
}

func (f FilterImpl) TextSearch() string {
	return f.FTextSearch
}

func (f FilterImpl) Selector() string {
	return f.FSelector
}

func (f FilterImpl) TagSelector() string {
	return f.FTagSelector
}

func (f FilterImpl) RunnerIds() []string {
	return f.FRunnerIds
}

func (f FilterImpl) RunnerTags() map[string]string {
	return f.FRunnerTags
}
