package main

type ModuleB struct{}

func (m *ModuleB) MyFunction() *Container {
	return dag.ModuleA().Message()
}
