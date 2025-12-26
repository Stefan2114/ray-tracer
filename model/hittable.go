package model

import (
	"math"
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
)

type HitRecord struct {
	T        float64
	P        *vec3.Vector
	Normal   *vec3.Vector
	Material Material
}

type Hittable interface {
	Hit(r *ray.Ray, tMin, tMax float64) (*HitRecord, bool)
	BoundingBox() *AABB
}

func SurroundingBox(box0, box1 *AABB) *AABB {
	small := vec3.NewVector(
		math.Min(box0.Min[0], box1.Min[0]),
		math.Min(box0.Min[1], box1.Min[1]),
		math.Min(box0.Min[2], box1.Min[2]),
	)
	big := vec3.NewVector(
		math.Max(box0.Max[0], box1.Max[0]),
		math.Max(box0.Max[1], box1.Max[1]),
		math.Max(box0.Max[2], box1.Max[2]),
	)
	return &AABB{Min: small, Max: big}
}
