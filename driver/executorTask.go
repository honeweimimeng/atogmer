package driver

import (
	"context"
	"github.com/sirupsen/logrus"
)

type ExecutorTask struct {
	executor Executor
}

func (e *ExecutorTask) Run() {
	e.executor.Execute()
}

func (e *ExecutorTask) Ctx() context.Context {
	return e.executor.Context().Context()
}

func (e *ExecutorTask) Interrupt() context.CancelFunc {
	e.Logger().Println("event executor task Interrupt Name:", e.executor.Name())
	return e.executor.Context().Interrupt()
}

func (e *ExecutorTask) Logger() logrus.StdLogger {
	return e.executor.Context().Config().Logger
}
