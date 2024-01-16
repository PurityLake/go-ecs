package ecs

import "reflect"

type Entity struct {
	id         int
	name       string
	components []Component
}

func NewEntity(id int, name string, components ...Component) Entity {
	return Entity{id, name, components}
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

func (e Entity) HasComponent(t reflect.Type) bool {
	for _, component := range e.components {
		if (component).Type() == t {
			return true
		}
	}
	return false
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
	return numTypes == 0
}
