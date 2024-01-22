package ecs

import (
	"fmt"
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

type ComponentRef struct {
	Index int
	Type  reflect.Type
}

type ComponentMutRef struct {
	Comp *Component
}

func GetComponentRefType(c *ComponentMutRef) reflect.Type {
	return (*c.Comp).Type()
}

func ComponentRefAs[T Component](ref *ComponentMutRef) (*T, error) {
	t, ok := (*ref.Comp).(T)
	if !ok {
		return nil, fmt.Errorf("ComponentRefAs: type assertion failed")
	}
	return &t, nil
}

func ComponentTypeIsA(a Component, t reflect.Type) bool {
	return a.Type() == t
}

func ComponentIsA[T Component](a interface{}) bool {
	_, ok := a.(T)
	return ok
}
