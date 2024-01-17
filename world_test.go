package ecs_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PurityLake/go-ecs"
)

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

func TestQueryWithEntities(t *testing.T) {
	world := ecs.World{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			world.AddEntity(fmt.Sprintf("Entity %d", i), Position{X: 0, Y: 0}, Renderable{W: 10, H: 10})
		} else {
			world.AddEntity(fmt.Sprintf("Entity %d", i), Position{X: 0, Y: 0})
		}
	}

	query1 := ecs.NewQuery(ecs.Type[Position]())
	query2 := ecs.NewQuery(ecs.Type[Position](), ecs.Type[Renderable]())

	{
		entities, components, found := world.QueryWithEntity(query1)
		if !found || len(entities) != 10 || len(components) != 10 {
			t.Fatalf(
				"found = %t (expected true) len(entities) = %d (expected 10) len(components) = %d (expected 10)",
				found, len(entities), len(components))
		}
	}
	{
		entities, components, found := world.QueryWithEntity(query2)
		if !found || len(entities) != 5 || len(components) != 5 {
			t.Fatalf(
				"found = %t (expected true) len(entities) = %d (expected 10) len(components) = %d (expected 10)",
				found, len(entities), len(components))
		}
	}
}

func TestQuery(t *testing.T) {
	world := ecs.World{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			world.AddEntity(fmt.Sprintf("Entity %d", i), Position{X: 0, Y: 0}, Renderable{W: 10, H: 10})
		} else {
			world.AddEntity(fmt.Sprintf("Entity %d", i), Position{X: 0, Y: 0})
		}
	}

	query1 := ecs.NewQuery(ecs.Type[Position]())
	query2 := ecs.NewQuery(ecs.Type[Position](), ecs.Type[Renderable]())

	{
		components, found := world.Query(query1)
		if !found || len(components) != 10 {
			t.Logf("%s != %s", ecs.Type[Position](), Position{}.Type())
			t.Fatalf("found = %t (expected true) len(components) = %d (expected 10)", found, len(components))
		}
	}
	{
		entities, components, found := world.QueryWithEntity(query2)
		if !found || len(entities) != 5 || len(components) != 5 {
			t.Fatalf("found = %t (expected true) len(components) = %d (expected 5)", found, len(components))
		}
	}
}

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

func TestQueryState(t *testing.T) {
	world := ecs.World{}
	world.AddState(MyState{69})
	world.AddState(TheState{420})

	query1 := ecs.NewQuery(ecs.Type[MyState]())
	query2 := ecs.NewQuery(ecs.Type[TheState]())

	{
		states, found := world.QueryState(query1)
		if !found || len(states) != 1 {
			t.Fatalf("found = %t (expected true) len(states) = %d (expected 1)", found, len(states))
		}
	}
	{
		states, found := world.QueryState(query2)
		if !found || len(states) != 1 {
			t.Fatalf("found = %t (expected true) len(states) = %d (expected 1)", found, len(states))
		}
	}
}

func BenchmarkQuery(b *testing.B) {
	world := ecs.World{}

	for i := 0; i < b.N; i++ {
		world.AddEntity(fmt.Sprintf("%d", i), Position{X: 0, Y: 0}, Renderable{W: 10, H: 10})
	}
	query := ecs.NewQuery(Position{}.Type(), Renderable{}.Type())
	b.ResetTimer()
	world.Query(query)
}

func BenchmarkQueryWithEntity(b *testing.B) {
	world := ecs.World{}

	for i := 0; i < b.N; i++ {
		world.AddEntity(fmt.Sprintf("%d", i), Position{X: 0, Y: 0}, Renderable{W: 10, H: 10})
	}
	query := ecs.NewQuery(Position{}.Type(), Renderable{}.Type())
	b.ResetTimer()
	world.QueryWithEntity(query)
}
