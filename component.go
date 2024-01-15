package ecs

import (
	"reflect"
)

type ComponentData interface{}

type Component interface {
	Name() string
	Update()
	Data() ComponentData
	Type() reflect.Type
}

func ComponentTypeIsA(a Component, t reflect.Type) bool {
	return a.Type() == t
}

func ComponentIsA[T Component](a interface{}) bool {
	_, ok := a.(T)
	return ok
}
