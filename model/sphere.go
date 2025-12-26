package model

import (
	"math"
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
)

type Sphere struct {
	Center   *vec3.Vector
	Radius   float64
	Material Material
}

func (s *Sphere) Hit(r *ray.Ray, tMin, tMax float64) (*HitRecord, bool) {
	oc := r.Origin().Minus(s.Center)
	a := r.Direction().Dot(r.Direction())
	b := oc.Dot(r.Direction())
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			return s.getRecord(r, temp), true
		}
		temp = (-b + math.Sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			return s.getRecord(r, temp), true
		}
	}
	return nil, false
}

func (s *Sphere) BoundingBox() *AABB {
	return &AABB{
		Min: s.Center.Minus(vec3.NewVector(s.Radius, s.Radius, s.Radius)),
		Max: s.Center.Plus(vec3.NewVector(s.Radius, s.Radius, s.Radius)),
	}
}

func (s *Sphere) getRecord(r *ray.Ray, t float64) *HitRecord {
	p := r.At(t)
	normal := p.Minus(s.Center).Scaled(1.0 / s.Radius)
	return &HitRecord{T: t, P: p, Normal: normal, Material: s.Material}
}
