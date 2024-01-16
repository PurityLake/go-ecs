package ecs

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

func (world World) QueryWithEntity(query Query) ([]Entity, [][]Component, bool) {
	var entities []Entity
	var allComponents [][]Component

	numTypes := query.numTypes
	if numTypes == 0 {
		return nil, nil, false
	}

	for _, e := range world.entities {
		components, found := query.Match(&e)
		if found {
			entities = append(entities, e)
			allComponents = append(allComponents, components)
		}
	}

	if len(entities) == 0 {
		return nil, nil, false
	}
	return entities, allComponents, true
}

func (w World) Query(query Query) ([][]Component, bool) {
	var allComponents [][]Component

	numTypes := query.numTypes
	if numTypes == 0 {
		return nil, false
	}

	for _, e := range w.entities {
		components, found := query.Match(&e)
		if found {
			allComponents = append(allComponents, components)
		}
	}
	if len(allComponents) == 0 {
		return nil, false
	}
	return allComponents, true
}
