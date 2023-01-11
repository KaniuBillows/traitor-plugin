package executor

import "github.com/dop251/goja"

type Executable interface {
	GetName() string
	ModuleLoader(runtime *goja.Runtime, module *goja.Object)
}

type AsyncExecutable interface {
	Require(exec *Executor) func(runtime *goja.Runtime, module *goja.Object)
	GetName() string
}
