package test

import (
	"testing"

	"github.com/PurityLake/go-ecs"
	"github.com/PurityLake/go-ecs/test/components"
)

func TestSystemCreation(t *testing.T) {
	world := ecs.World{}
	world.AddSystems(
		components.ExampleSystemInit{},
		components.ExampleSystem{})

	world.Start()
	query := ecs.NewQuery(ecs.Type[components.Position](), ecs.Type[components.Renderable]())
	ents, _, found := world.QueryWithEntity(query)
	if !found || len(ents) != 1 {
		t.Fatal("System did not add entities to the world")
	}
}

func TestSystemUpdate(t *testing.T) {
	world := ecs.World{}
	world.AddSystems(
		components.ExampleSystemInit{},
		components.ExampleSystemWithQuery{})

	world.Start()
	query := ecs.NewQuery(ecs.Type[components.Position](), ecs.Type[components.Renderable]())
	comps, found := world.Query(query)
	if !found {
		t.Fatal("Failed to make query")
	}
	for _, compList := range comps {
		for _, comp := range compList {
			switch c := comp.(type) {
			case *components.Position:
				if c.X != 0 || c.Y != 0 {
					t.Fatalf("Original position is meant to be (0, 0) got (%d, %d)", c.X, c.Y)
				}
			case *components.Renderable:
				if c.W != 0 || c.H != 0 {
					t.Fatalf("Original size is meant to be (0, 0) got (%d, %d)", c.W, c.H)
				}
			}
		}
	}
	world.Update()
	for _, compList := range comps {
		for _, comp := range compList {
			switch c := comp.(type) {
			case *components.Position:
				if c.X != components.QueryTestX || c.Y != components.QueryTestY {
					t.Fatalf("Original position is meant to be (%d, %d) got (%d, %d)",
						components.QueryTestY, components.QueryTestY, c.X, c.Y)
				}
			case *components.Renderable:
				if c.W != components.QueryTestW || c.H != components.QueryTestH {
					t.Fatalf("Original size is meant to be (%d, %d) got (%d, %d)",
						components.QueryTestW, components.QueryTestH, c.W, c.H)
				}
			}
		}
	}
}
