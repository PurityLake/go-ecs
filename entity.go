package ecs

import (
	"reflect"
)

type Entity struct {
	name       string
	components []Component
}

func NewEntity(name string, components ...Component) Entity {
	return Entity{name, components}
}

func (e Entity) Name() string {
	return e.name
}

func (e *Entity) AddComponent(component Component) {
	e.components = append(e.components, component)
}

func (e Entity) Components() []Component {
	return e.components
}

func (e Entity) HasComponents(types ...reflect.Type) bool {
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
	return true
}
