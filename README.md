# GoECS
An implementation of the ECS design pattern in Go. Currently
in the early stages of development

# Roadmap
- [ ] Add tests (ongoing)
- [x] Add a Query object to reuse queries (ecs.Query)
- [ ] Add way to call functions with queries as parameters
- [ ] Add conditional running of functions

# Examples
### Using a World and Query it's entites
```go
world := ecs.World{}
// adds one entity with a Renderable
world.AddSystem(ExampleSystem{})
world.Start()

query := ecs.NewQuery(ecs.Type[Renderable]())

// this query only returns components of the entites that match
comps, found := world.Query(query)
if found {
  for _, compList := range comps {
    for _, comp := range compList {
      fmt.Println("Component: ", comp.Name())
    }
  }
}

// this query returns entites and it's components
entities, comps, found := world.QueryWithEntity(query)
if found {
  for i, e := range entities {
    fmt.Println("Entity: ", e.Name())
    for _, comp := range comps[i] {
      fmt.Println("Component: ", comp.Name())
    }
  }
}

// this query returns entites and it's components
entities, comps, found := world.QueryWithEntityMut(query)
if found {
  for i, entity := range entities {
    for _, comp := range componentsMut[i] {
      c, err := entity.GetComponent(comp)
      if err != nil {
        t.Fatal(err)
      }
      switch c := (*c).(type) {
      case components.Position:
        c.X = 1
        c.Y = 1
        entity.SetComponent(comp, c)
      case components.Renderable:
        c.W = 1
        c.H = 1
        entity.SetComponent(comp, c)
      default:
        log.Fatal("unexpected component type", comp.Type)
      }
    }
  }
}
```

### Creating a custom System
```go
type ExampleSystem struct{
  ecs.SystemBase
  // data can go here
}

// only called by world.Start() or world.InitNew()
func (s ExampleSystem) Setup(world *ecs.World) {
  fmt.Println("ExampleSystem setup")
  world.AddEntity("example", Renderable{})
}

func (s ExampleSystem) Update(world *ecs.World) {
  // do a query
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
