package testkube

func (s TestSuiteStep) Type() *TestSuiteStepType {
	if s.Execute != nil {
		return TestSuiteStepTypeExecuteTest
	}
	if s.Delay != nil {
		return TestSuiteStepTypeDelay
	}
	return nil
}

func (s TestSuiteStep) FullName() string {
	switch s.Type() {
	case TestSuiteStepTypeDelay:
		return s.Delay.FullName()
	case TestSuiteStepTypeExecuteTest:
		return s.Execute.FullName()
	default:
		return "unknown"
	}
}
