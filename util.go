package ecs

import "reflect"

func Type[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}
