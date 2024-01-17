package ecs

type StateData interface{}

type State interface {
	GetData() StateData
}
