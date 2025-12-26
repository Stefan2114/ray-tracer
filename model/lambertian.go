package model

import (
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
	"ray-tracer/utils"
)

type Lambertian struct {
	Albedo *vec3.Vector
}

func NewLambertian(Albedo *vec3.Vector) *Lambertian {
	return &Lambertian{Albedo: Albedo}
}

func (l *Lambertian) Scatter(_ *ray.Ray, rec *HitRecord) (attenuation *vec3.Vector, scattered *ray.Ray, ok bool) {
	target := rec.P.Plus(rec.Normal).Plus(utils.RandomInUnitSphere())
	scattered = ray.NewRay(rec.P, target.Minus(rec.P))
	attenuation = l.Albedo
	return attenuation, scattered, true
}

func (l *Lambertian) Emitted(p *vec3.Vector) *vec3.Vector {
	return vec3.NewVector(0, 0, 0)
}
