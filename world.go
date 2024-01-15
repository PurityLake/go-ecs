package ecs

import "reflect"

type World struct {
	systems  []System
	entities []Entity
	currId   int
}

func (w *World) AddSystem(system System) {
	w.systems = append(w.systems, system)
}

func (w *World) Start() {
	for _, system := range w.systems {
		system.Setup(w)
	}
}

func (w *World) Update() {
	for _, system := range w.systems {
		system.Update(w)
	}
}

func (w *World) AddEntity(name string, components ...Component) {
	w.entities = append(w.entities, NewEntity(w.currId, name, components...))
	w.currId++
}

func (world World) QueryWithEntity(types ...reflect.Type) ([]Entity, [][]Component, bool) {
	var entities []Entity
	var allComponents [][]Component

	numTypes := len(types)
	if numTypes == 0 {
		return nil, nil, false
	}

	for _, e := range world.entities {
		var components []Component
		if len(e.Components()) < numTypes {
			return nil, nil, false
		}
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

	if len(entities) == 0 {
		return nil, nil, false
	}
	return entities, allComponents, true
}

func (w World) Query(types ...reflect.Type) ([][]Component, bool) {
	var allComponents [][]Component

	numTypes := len(types)
	if numTypes == 0 {
		return nil, false
	}

	for _, e := range w.entities {
		var components []Component
		if len(e.Components()) < numTypes {
			return nil, false
		}
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
	if len(allComponents) == 0 {
		return nil, false
	}
	return allComponents, true
}
