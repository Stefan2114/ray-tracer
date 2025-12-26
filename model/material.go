package model

import (
	"math"
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
)

type Material interface {
	Scatter(rIn *ray.Ray, rec *HitRecord) (attenuation *vec3.Vector, scattered *ray.Ray, ok bool)
	Emitted(p *vec3.Vector) *vec3.Vector
}

func reflect(v, n *vec3.Vector) *vec3.Vector {
	return v.Minus(n.Scaled(2 * v.Dot(n)))
}

func refract(v, n *vec3.Vector, niOverNt float64) (refracted *vec3.Vector, ok bool) {
	uv := v.Unit()
	dt := uv.Dot(n)
	discriminant := 1 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		refracted = uv.Minus(n.Scaled(dt)).Scaled(niOverNt).Minus(n.Scaled(math.Sqrt(discriminant)))
		return refracted, true
	}
	return nil, false
}
