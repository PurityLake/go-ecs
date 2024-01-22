package ecs_test

import (
	"testing"

	"github.com/PurityLake/go-ecs"
)

func TestSystemCreation(t *testing.T) {
	world := ecs.World{}
	world.AddSystems(
		ExampleSystemInit{},
		ExampleSystem{})

	world.Start()
	query := ecs.NewQuery(ecs.Type[Position](), ecs.Type[Renderable]())
	ents, _, found := world.QueryWithEntity(query)
	if !found || len(ents) != 1 {
		t.Fatal("System did not add entities to the world")
	}
}

func TestSystemUpdate(t *testing.T) {
	world := ecs.World{}
	world.AddSystems(
		ExampleSystemInit{},
		ExampleSystemWithQuery{})

	world.Start()
	query := ecs.NewQuery(ecs.Type[Position](), ecs.Type[Renderable]())
	comps, found := world.Query(query)
	if !found {
		t.Fatal("Failed to make query")
	}
	for _, compList := range comps {
		for _, comp := range compList {
			switch c := comp.(type) {
			case *Position:
				if c.X != 0 || c.Y != 0 {
					t.Fatalf("Original position is meant to be (0, 0) got (%d, %d)", c.X, c.Y)
				}
			case *Renderable:
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
			case *Position:
				if c.X != QueryTestX || c.Y != QueryTestY {
					t.Fatalf("Original position is meant to be (%d, %d) got (%d, %d)",
						QueryTestY, QueryTestY, c.X, c.Y)
				}
			case *Renderable:
				if c.W != QueryTestW || c.H != QueryTestH {
					t.Fatalf("Original size is meant to be (%d, %d) got (%d, %d)",
						QueryTestW, QueryTestH, c.W, c.H)
				}
			}
		}
	}
}
