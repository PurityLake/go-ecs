package ecs

type System interface {
	Id() int
	Update(world *World)
	Setup(world *World)
}

type SystemBase struct {
	Id int
}
