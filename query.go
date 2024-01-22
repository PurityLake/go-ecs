package ecs

import (
	"reflect"
)

type Query struct {
	types    []reflect.Type
	numTypes int
}

func NewQuery(types ...reflect.Type) Query {
	return Query{types, len(types)}
}

func (query Query) MatchState(world World, state State) (State, bool) {
	for _, other := range world.state {
		if reflect.TypeOf(state) == reflect.TypeOf(other) {
			return other, true
		}
	}
	return nil, false
}

func (query Query) Match(entity *Entity) ([]Component, bool) {
	if len(entity.components) < query.numTypes {
		return nil, false
	}
	components := make([]Component, query.numTypes)
	numFound := 0
ComponentLoop:
	for _, c := range entity.Components() {
		for i, t := range query.types {
			if ComponentTypeIsA(c, t) {
				components[i] = c
				numFound++
				if numFound == query.numTypes {
					break ComponentLoop
				}
				break
			}
		}
	}
	if numFound != query.numTypes {
		return nil, false
	}
	return components, true
}

func (query Query) MatchMut(entity *Entity) ([]*Component, bool) {
	if len(entity.components) < query.numTypes {
		return nil, false
	}
	components := make([]*Component, query.numTypes)
	numFound := 0
ComponentLoop:
	for _, c := range entity.Components() {
		for i, t := range query.types {
			if ComponentTypeIsA(c, t) {
				components[i] = &c
				numFound++
				if numFound == query.numTypes {
					break ComponentLoop
				}
				break
			}
		}
	}
	if numFound != query.numTypes {
		return nil, false
	}
	return components, true
}
