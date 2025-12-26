package ray

import (
	"ray-tracer/geo/vec3"
)

type Ray struct {
	a *vec3.Vector
	b *vec3.Vector
}

func NewRay(a, b *vec3.Vector) *Ray {
	return &Ray{a: a, b: b}
}

func (r *Ray) Origin() *vec3.Vector {
	return r.a // TODO: see if i should return a copy of it
}

func (r *Ray) Direction() *vec3.Vector {
	return r.b // TODO: see if i should return a copy of it
}

func (r *Ray) At(t float64) *vec3.Vector {
	return r.a.Plus(r.b.Scaled(t))
}
