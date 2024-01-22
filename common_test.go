package ecs_test

import (
	"reflect"

	"github.com/PurityLake/go-ecs"
)

// ############
// Components

type Position struct {
	id   int
	X, Y int
}

func (p Position) Id() int {
	return p.id
}

func (p Position) Name() string {
	return "position"
}

func (p Position) Update() {}

func (p Position) Data() ecs.Data {
	return p
}

func (p Position) Type() reflect.Type {
	return reflect.TypeOf(p)
}

type Renderable struct {
	id   int
	W, H int
}

func (r Renderable) Id() int {
	return r.id
}

func (r Renderable) Name() string {
	return "renderable"
}

func (r Renderable) Update() {}

func (r Renderable) Data() ecs.Data {
	return r
}

func (r Renderable) Type() reflect.Type {
	return reflect.TypeOf(r)
}

// ############
// Systems

type ExampleSystem struct {
	ecs.SystemBase
}

func (es ExampleSystem) Update(world *ecs.World) {
}

func (es ExampleSystem) Id() int {
	return es.SystemBase.Id
}

func (es ExampleSystem) Setup(world *ecs.World) {
}

type ExampleSystemInit struct {
	ecs.SystemBase
}

func (esi ExampleSystemInit) Update(world *ecs.World) {
}

func (esi ExampleSystemInit) Id() int {
	return esi.SystemBase.Id
}

func (esi ExampleSystemInit) Setup(world *ecs.World) {
	world.AddEntity("Entity 1", Position{X: 0, Y: 0}, Renderable{W: 0, H: 0})
}

type ExampleSystemWithQuery struct {
	ecs.SystemBase
	query ecs.Query
}

const (
	QueryTestX = 49
	QueryTestY = 420
	QueryTestW = 500
	QueryTestH = 42
)

func (eswq ExampleSystemWithQuery) Update(world *ecs.World) {
	c, found := world.Query(eswq.query)
	if !found {
		return
	}
	for _, cList := range c {
		for _, c := range cList {
			switch c := c.(type) {
			case Position:
				c.X = QueryTestX
				c.Y = QueryTestY
			case Renderable:
				c.W = QueryTestW
				c.H = QueryTestH
			}
		}
	}
}

func (eswq ExampleSystemWithQuery) Id() int {
	return eswq.SystemBase.Id
}

func (eswq ExampleSystemWithQuery) Setup(world *ecs.World) {
}

// ############
// States

type MyState struct {
	val int
}

func (state MyState) GetData() ecs.StateData {
	return state
}

type TheState struct {
	val int
}

func (state TheState) GetData() ecs.StateData {
	return state
}
