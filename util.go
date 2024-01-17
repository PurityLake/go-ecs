package ecs

import "reflect"

func CompType[Comp Component]() reflect.Type {
	return reflect.TypeOf((*Comp)(nil)).Elem()
}

func StateType[S State]() reflect.Type {
	return reflect.TypeOf((*S)(nil)).Elem()
}
