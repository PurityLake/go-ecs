package test

import (
	"fmt"
	"testing"

	"github.com/PurityLake/go-ecs"
	"github.com/PurityLake/go-ecs/components"
)

func TestQueryWithEntities(t *testing.T) {
	world := ecs.World{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			world.AddEntity(fmt.Sprintf("Entity %d", i), components.Position{X: 0, Y: 0}, components.Renderable{W: 10, H: 10})
		} else {
			world.AddEntity(fmt.Sprintf("Entity %d", i), components.Position{X: 0, Y: 0})
		}
	}

	query1 := ecs.NewQuery(ecs.Type[components.Position]())
	query2 := ecs.NewQuery(ecs.Type[components.Position](), ecs.Type[components.Renderable]())

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
			world.AddEntity(fmt.Sprintf("Entity %d", i), components.Position{X: 0, Y: 0}, components.Renderable{W: 10, H: 10})
		} else {
			world.AddEntity(fmt.Sprintf("Entity %d", i), components.Position{X: 0, Y: 0})
		}
	}

	query1 := ecs.NewQuery(ecs.Type[components.Position]())
	query2 := ecs.NewQuery(ecs.Type[components.Position](), ecs.Type[components.Renderable]())

	{
		components, found := world.Query(query1)
		if !found || len(components) != 10 {
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

type PositionRef *components.Position

func TestQueryWithEntitiesMut(t *testing.T) {
	world := ecs.World{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			world.AddEntity(fmt.Sprintf("Entity %d", i), components.Position{X: 0, Y: 0}, components.Renderable{W: 10, H: 10}, components.Player{})
		} else {
			world.AddEntity(fmt.Sprintf("Entity %d", i), components.Position{X: 0, Y: 0}, components.Renderable{W: 20, H: 20})
		}
	}

	query1 := ecs.NewQuery(ecs.Type[components.Position](), ecs.Type[components.Renderable]())

	entities, componentsMut, foundMut := world.QueryWithEntityMut(query1)
	if !foundMut || len(componentsMut) != 10 {
		t.Fatalf("found = %t (expected true) len(components) = %d (expected 10)", foundMut, len(componentsMut))
	}

	for i := 0; i < len(entities); i++ {
		for _, cList := range componentsMut {
			for _, comp := range cList {
				c, err := entities[i].GetComponent(comp)
				if err != nil {
					t.Fatal(err)
				}
				switch comp.Type {
				case ecs.Type[components.Position]():
					pos, ok := (*c).(components.Position)
					if !ok {
						t.Fatal("unexpected component type", comp.Type)
					}
					pos.X = 1
					pos.Y = 1
					entities[i].SetComponent(comp, pos)
				case ecs.Type[components.Renderable]():
					rend, ok := (*c).(components.Renderable)
					if !ok {
						t.Fatal("unexpected component type", comp.Type)
					}
					rend.W = 1
					rend.H = 1
					entities[i].SetComponent(comp, rend)
				default:
					t.Fatal("unexpected component type", comp.Type)
				}
			}
		}
	}

	componentsQ, found := world.Query(query1)
	if !found || len(componentsQ) != 10 {
		t.Fatalf("found = %t (expected true) len(components) = %d (expected 10)", found, len(componentsQ))
	}

	for _, cList := range componentsQ {
		for _, comp := range cList {
			switch c := comp.(type) {
			case components.Position:
				if c.X != 1 || c.Y != 1 {
					t.Fatalf("components.Position is meant to be (1, 1) got (%d, %d)", c.X, c.Y)
				}
			case components.Renderable:
				if c.W != 1 || c.H != 1 {
					t.Fatalf("components.Renderable is meant to be (1, 1) got (%d, %d)", c.W, c.H)
				}
			default:
				t.Fatal("unexpected component type", c)
			}
		}
	}
}

func TestQueryState(t *testing.T) {
	world := ecs.World{}
	world.AddState(components.MyState{Val: 69})
	world.AddState(components.TheState{Val: 420})

	query1 := ecs.NewQuery(ecs.Type[components.MyState]())
	query2 := ecs.NewQuery(ecs.Type[components.TheState]())

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
		world.AddEntity(fmt.Sprintf("%d", i), components.Position{X: 0, Y: 0}, components.Renderable{W: 10, H: 10})
	}
	query := ecs.NewQuery(components.Position{}.Type(), components.Renderable{}.Type())
	b.ResetTimer()
	world.Query(query)
}

func BenchmarkQueryWithEntity(b *testing.B) {
	world := ecs.World{}

	for i := 0; i < b.N; i++ {
		world.AddEntity(fmt.Sprintf("%d", i), components.Position{X: 0, Y: 0}, components.Renderable{W: 10, H: 10})
	}
	query := ecs.NewQuery(components.Position{}.Type(), components.Renderable{}.Type())
	b.ResetTimer()
	world.QueryWithEntity(query)
}
