package ecs_test

import (
	"fmt"
	"testing"

	"github.com/PurityLake/go-ecs"
)

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

func TestQueryWithEntitiesMut(t *testing.T) {
	world := ecs.World{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			world.AddEntity(fmt.Sprintf("Entity %d", i), Position{X: 0, Y: 0}, Renderable{W: 10, H: 10})
		} else {
			world.AddEntity(fmt.Sprintf("Entity %d", i), Position{X: 0, Y: 0})
		}
	}

	query1 := ecs.NewQuery(ecs.Type[Position]())

	entitiesMut, componentsMut, foundMut := world.QueryWithEntityMut(query1)
	if !foundMut || len(entitiesMut) != 10 || len(componentsMut) != 10 {
		t.Fatalf(
			"found = %t (expected true) len(entities) = %d (expected 10) len(components) = %d (expected 10)",
			foundMut, len(entitiesMut), len(componentsMut))
	}

	for _, cList := range componentsMut {
		for _, comp := range cList {
			switch c := (*comp).(type) {
			case *Position:
				c.X = 1
				c.Y = 1
			}
		}
	}

	components, found := world.Query(query1)
	if !found || len(components) != 10 {
		t.Fatalf("found = %t (expected true) len(components) = %d (expected 10)", found, len(components))
	}

	for _, cList := range components {
		for _, comp := range cList {
			switch c := comp.(type) {
			case *Position:
				if c.X != 1 || c.Y != 1 {
					t.Fatalf("Position is meant to be (1, 1) got (%d, %d)", c.X, c.Y)
				}
			}
		}
	}
}

func TestQueryMut(t *testing.T) {
	world := ecs.World{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			world.AddEntity(fmt.Sprintf("Entity %d", i), Position{X: 0, Y: 0}, Renderable{W: 10, H: 10})
		} else {
			world.AddEntity(fmt.Sprintf("Entity %d", i), Position{X: 0, Y: 0})
		}
	}

	query1 := ecs.NewQuery(ecs.Type[Position]())

	componentsMut, foundMut := world.QueryMut(query1)
	if !foundMut || len(componentsMut) != 10 {
		t.Logf("%s != %s", ecs.Type[Position](), Position{}.Type())
		t.Fatalf("found = %t (expected true) len(components) = %d (expected 10)", foundMut, len(componentsMut))
	}

	for _, cList := range componentsMut {
		for _, comp := range cList {
			switch c := (*comp).(type) {
			case *Position:
				c.X = 1
				c.Y = 1
			}
		}
	}

	components, found := world.Query(query1)
	if !found || len(components) != 10 {
		t.Fatalf("found = %t (expected true) len(components) = %d (expected 10)", found, len(components))
	}

	for _, cList := range components {
		for _, comp := range cList {
			switch c := comp.(type) {
			case *Position:
				if c.X != 1 || c.Y != 1 {
					t.Fatalf("Position is meant to be (1, 1) got (%d, %d)", c.X, c.Y)
				}
			}
		}
	}
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
