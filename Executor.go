package executor

import (
	"github.com/dop251/goja"
	"sync"
)

type Executor struct {
	Vm   *goja.Runtime
	Wait sync.WaitGroup
}

func MakeExecutor() *Executor {
	return &Executor{
		Vm:   goja.New(),
		Wait: sync.WaitGroup{},
	}
}
