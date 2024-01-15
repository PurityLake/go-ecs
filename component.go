package ecs

import (
	"reflect"
)

type Data interface{}

type Component interface {
	Id() int
	Name() string
	Update()
	Data() Data
	Type() reflect.Type
}

func ComponentTypeIsA(a Component, t reflect.Type) bool {
	return a.Type() == t
}

func ComponentIsA[T Component](a interface{}) bool {
	_, ok := a.(T)
	return ok
}
