package ecs

type System interface {
	Setup(world *World)
	Update(world *World)
}
