package executor

import "github.com/dop251/goja"

type Executable interface {
	GetModule() Executable
	Require(exec *Executor) func(runtime *goja.Runtime, module *goja.Object)
	GetName() string
}
