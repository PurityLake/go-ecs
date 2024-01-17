package ecs

type World struct {
	systems  []System
	entities []Entity
	state    []State
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

func (w *World) AddState(state State) {
	w.state = append(w.state, state)
}

func (world World) QueryState(query Query) ([]State, bool) {
	var states []State

	numTypes := query.numTypes
	if numTypes == 0 {
		return nil, false
	}

	count := 0
	for _, s := range world.state {
		state, found := query.MatchState(world, s)
		if found {
			states = append(states, state)
			count++
			if numTypes == count {
				break
			}
		}
	}

	if len(states) == 0 {
		return nil, false
	}
	return states, true
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
