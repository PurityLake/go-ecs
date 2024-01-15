# GoECS
An implementation of the ECS design pattern in Go. Currently
in the early stages of development

# Roadmap
- [ ] Add tests
- [ ] Add way to call functions with queries as paramets
- [ ] Add conditional running of functions

# Examples
### Using a World and Query it's entites
```go
world := ecs.World{}
// adds one entity with a Renderable
world.AddSystem(ExampleSystem{})
world.Start()

comps, found := world.Query(systems.RenderableComponet{}.Type())
if found {
    for _, compList := range comps {
        for _, comp := range compList {
            fmt.Println("Component: ", comp.Name())
        }
    }
}

entities, comps, found := world.QueryWithEntity(RenderableCompomnent{}.Type())
if found {
    for i, e := range entities {
        fmt.Println("Entity: ", e.Name())
        for _, comp := range comps[i] {
            fmt.Println("Component: ", comp.Name())
        }
    }
}
```

### Creating a custom System
```go
type ExampleSystem struct{
  // data can go here
}

func (s ExampleSystem) Setup(world *ecs.World) {
	fmt.Println("ExampleSystem setup")
	world.AddEntity("example", Renderable{})
}

func (s ExampleSystem) Update(world *ecs.World) {
}
```

### Creating a custom component
```go
type Position struct {
	id   int
	x, y int
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
```
