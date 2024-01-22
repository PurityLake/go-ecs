package ecs

import (
	"fmt"
	"reflect"
)

type Entity struct {
	id         int
	name       string
	components []Component
}

type EntityRef struct {
	Entity *Entity
}

func NewEntity(id int, name string, components ...Component) Entity {
	return Entity{id, name, components}
}

func (e *Entity) Name() string {
	return e.name
}

func (e *Entity) AddComponent(component Component) {
	e.components = append(e.components, component)
}

func (e *Entity) Components() []Component {
	return e.components
}

func (e *Entity) NumComponents() int {
	return len(e.components)
}

func (e *Entity) MutComponent(index int) (*Component, error) {
	if index < 0 || index >= len(e.components) {
		return nil, fmt.Errorf("index %d out of range", index)
	}
	return &e.components[index], nil
}

func (e *Entity) GetComponent(ref ComponentRef) (*Component, error) {
	if ref.Index < 0 || ref.Index >= len(e.components) {
		return nil, fmt.Errorf("index %d out of range", ref.Index)
	}
	return &e.components[ref.Index], nil
}

func (e *Entity) SetComponent(ref ComponentRef, component Component) error {
	if ref.Index < 0 || ref.Index >= len(e.components) {
		return fmt.Errorf("index %d out of range", ref.Index)
	}
	e.components[ref.Index] = component
	return nil
}

func (e *Entity) HasComponent(t reflect.Type) bool {
	for _, component := range e.components {
		if (component).Type() == t {
			return true
		}
	}
	return false
}

func (e *Entity) HasComponents(types ...reflect.Type) bool {
	numTypes := len(types)
	numComponents := len(e.components)
	if numTypes > numComponents {
		return false
	}
	for _, component := range e.components {
		for _, t := range types {
			if (component).Type() == t {
				numTypes--
				break
			}
		}
	}
	return numTypes == 0
}
