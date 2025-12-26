package model

import (
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
	"ray-tracer/utils"
)

type Metal struct {
	albedo *vec3.Vector
	fuzz   float64
}

func NewMetal(albedo *vec3.Vector, fuzz float64) *Metal {
	m := &Metal{albedo: albedo}
	if fuzz < 1 {
		m.fuzz = fuzz
	} else {
		m.fuzz = 1
	}
	return m
}
func (m *Metal) Scatter(rIn *ray.Ray, rec *HitRecord) (attenuation *vec3.Vector, scattered *ray.Ray, ok bool) {
	reflected := reflect(rIn.Direction().Unit(), rec.Normal)
	direction := reflected.Plus(utils.RandomInUnitSphere().Scaled(m.fuzz))
	scattered = ray.NewRay(rec.P, direction)
	attenuation = m.albedo
	return attenuation, scattered, scattered.Direction().Dot(rec.Normal) > 0
}

func (m *Metal) Emitted(p *vec3.Vector) *vec3.Vector {
	return vec3.NewVector(0, 0, 0)
}
