package testkube

func NewRunningExecutionResult() *ExecutionResult {
	return &ExecutionResult{
		Status: StatusPtr(RUNNING_ExecutionStatus),
	}
}

// NewPendingExecutionResult DEPRECATED since testkube@1.0.0
func NewPendingExecutionResult() ExecutionResult {
	return ExecutionResult{
		Status: StatusPtr(RUNNING_ExecutionStatus),
	}
}

func NewErrorExecutionResult(err error) ExecutionResult {
	return ExecutionResult{
		Status:       StatusPtr(FAILED_ExecutionStatus),
		ErrorMessage: err.Error(),
		Output:       err.Error(),
	}
}

func (e *ExecutionResult) Abort() {
	e.Status = StatusPtr(ABORTED_ExecutionStatus)
}

func (e *ExecutionResult) Timeout() {
	e.Status = StatusPtr(TIMEOUT_ExecutionStatus)
}

func (e *ExecutionResult) InProgress() {
	e.Status = StatusPtr(RUNNING_ExecutionStatus)
}

func (e *ExecutionResult) Success() {
	e.Status = StatusPtr(PASSED_ExecutionStatus)
}

func (e *ExecutionResult) Error() {
	e.Status = StatusPtr(FAILED_ExecutionStatus)
}

func (e *ExecutionResult) IsCompleted() bool {
	return e.IsPassed() || e.IsFailed()
}

func (e *ExecutionResult) IsRunning() bool {
	return *e.Status == RUNNING_ExecutionStatus
}

func (e *ExecutionResult) IsQueued() bool {
	return *e.Status == QUEUED_ExecutionStatus
}

func (e *ExecutionResult) IsPassed() bool {
	return *e.Status == PASSED_ExecutionStatus
}

func (e *ExecutionResult) IsFailed() bool {
	return *e.Status == FAILED_ExecutionStatus
}

func (e *ExecutionResult) IsAborted() bool {
	return *e.Status == ABORTED_ExecutionStatus
}
func (e *ExecutionResult) IsTimeout() bool {
	return *e.Status == TIMEOUT_ExecutionStatus
}

func (e *ExecutionResult) Err(err error) *ExecutionResult {
	e.Status = ExecutionStatusFailed
	e.ErrorMessage = err.Error()
	return e
}

// WithErrors return error result if any of passed errors is not nil
func (e *ExecutionResult) WithErrors(errors ...error) *ExecutionResult {
	for _, err := range errors {
		if err != nil {
			return e.Err(err)
		}
	}
	return e
}

func (e *ExecutionResult) FailedStepsCount() int {
	count := 0
	for _, v := range e.Steps {
		if v.Status != string(PASSED_ExecutionStatus) {
			count++
		}
	}
	return count
}
