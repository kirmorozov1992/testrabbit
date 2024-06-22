package testrabbit

import "git.ipc/kirillmorozov/testrabbit/internal/testrabbit"

func NewTestRabbit() func(string, string, string) *testrabbit.TestRabbit {
	return testrabbit.NewTestRabbit
}
