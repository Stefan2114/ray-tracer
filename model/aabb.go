package model

import (
	"math"
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
)

type AABB struct {
	Min *vec3.Vector
	Max *vec3.Vector
}

func (b AABB) Hit(r *ray.Ray, tMin, tMax float64) bool {

	for a := 0; a < 3; a++ {
		invD := 1.0 / r.Direction()[a]
		t0 := (b.Min[a] - r.Origin()[a]) * invD
		t1 := (b.Max[a] - r.Origin()[a]) * invD
		if invD < 0 {
			t0, t1 = t1, t0
		}
		tMin = math.Max(t0, tMin)
		tMax = math.Min(t1, tMax)
		if tMax <= tMin {
			return false
		}
	}
	return true
}
