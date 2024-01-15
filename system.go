package ecs

import "reflect"

type System struct {
	entities []Entity
}

func (s *System) AddEntity(e Entity) {
	s.entities = append(s.entities, e)
}

func (s System) Query(types ...reflect.Type) [][]Component {
	var allComponents [][]Component
	numTypes := len(types)
	for _, e := range s.entities {
		var components []Component
	ComponentLoop:
		for _, c := range e.Components() {
			for _, t := range types {
				if ComponentTypeIsA(c, t) {
					components = append(components, c)
					if len(components) == numTypes {
						break ComponentLoop
					}
				}
			}
		}
		if len(components) == numTypes {
			allComponents = append(allComponents, components)
		}
	}
	return allComponents
}

func (s System) QueryWithEntity(types ...reflect.Type) ([]Entity, [][]Component) {
	var entities []Entity
	var allComponents [][]Component

	numTypes := len(types)
	for _, e := range s.entities {
		var components []Component
	ComponentLoop:
		for _, c := range e.Components() {
			for _, t := range types {
				if ComponentTypeIsA(c, t) {
					components = append(components, c)
					if len(components) == numTypes {
						break ComponentLoop
					}
				}
			}
		}
		if len(components) == numTypes {
			allComponents = append(allComponents, components)
			entities = append(entities, e)
		}
	}

	return entities, allComponents
}
