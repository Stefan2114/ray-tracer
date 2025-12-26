package model

import (
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
	"ray-tracer/utils"
)

type Metal struct {
	Albedo *vec3.Vector
	Fuzz   float64
}

func NewMetal(Albedo *vec3.Vector, Fuzz float64) *Metal {
	m := &Metal{Albedo: Albedo}
	if Fuzz < 1 {
		m.Fuzz = Fuzz
	} else {
		m.Fuzz = 1
	}
	return m
}
func (m *Metal) Scatter(rIn *ray.Ray, rec *HitRecord) (attenuation *vec3.Vector, scattered *ray.Ray, ok bool) {
	reflected := reflect(rIn.Direction().Unit(), rec.Normal)
	direction := reflected.Plus(utils.RandomInUnitSphere().Scaled(m.Fuzz))
	scattered = ray.NewRay(rec.P, direction)
	attenuation = m.Albedo
	return attenuation, scattered, scattered.Direction().Dot(rec.Normal) > 0
}

func (m *Metal) Emitted(p *vec3.Vector) *vec3.Vector {
	return vec3.NewVector(0, 0, 0)
}
