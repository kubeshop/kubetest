package testresult

import (
	"context"
	"time"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

// TODO: Adjust when it gets too small.
const PageDefaultLimit int = 1000

type Filter interface {
	Name() string
	NameDefined() bool
	StartDate() time.Time
	StartDateDefined() bool
	EndDate() time.Time
	EndDateDefined() bool
	Page() int
	PageSize() int
	TextSearchDefined() bool
	TextSearch() string
}

type Repository interface {
	// Get gets execution result by id
	Get(ctx context.Context, id string) (testkube.TestExecution, error)
	// GetByNameAndScript gets execution result by name
	GetByNameAndScript(ctx context.Context, name, script string) (testkube.TestExecution, error)
	// GetExecutionsTotals gets executions total stats using a filter, use filter with no data for all
	GetExecutionsTotals(ctx context.Context, filter Filter) (totals testkube.ExecutionsTotals, err error)
	// GetExecutions gets executions using a filter, use filter with no data for all
	GetExecutions(ctx context.Context, filter Filter) ([]testkube.TestExecution, error)
	// Insert inserts new execution result
	Insert(ctx context.Context, result testkube.TestExecution) error
	// Update updates execution result
	Update(ctx context.Context, result testkube.TestExecution) error
	// StartExecution updates execution start time
	StartExecution(ctx context.Context, id string, startTime time.Time) error
	// EndExecution updates execution end time
	EndExecution(ctx context.Context, id string, endTime time.Time) error
}
