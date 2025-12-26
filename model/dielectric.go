package model

import (
	"math"
	"math/rand"
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
)

const SchlickPower = 5.0

type Dielectric struct {
	refIndex float64
}

func NewDielectric(ri float64) *Dielectric {
	return &Dielectric{refIndex: ri}
}
func (d *Dielectric) Scatter(rIn *ray.Ray, rec *HitRecord) (attenuation *vec3.Vector, scattered *ray.Ray, ok bool) {
	reflected := reflect(rIn.Direction(), rec.Normal)
	attenuation = vec3.NewVector(1, 1, 1)
	outWardNormal := rec.Normal
	niOverNt := 1 / d.refIndex
	cosine := -rIn.Direction().Dot(rec.Normal) / rIn.Direction().Len()

	if rIn.Direction().Dot(rec.Normal) > 0 {
		outWardNormal = outWardNormal.Inv()
		niOverNt = d.refIndex
		cosine = d.refIndex * rIn.Direction().Dot(rec.Normal) / rIn.Direction().Len()

	}

	var refracted *vec3.Vector
	reflectProb := 1.0
	if refracted, ok = refract(rIn.Direction(), outWardNormal, niOverNt); ok {
		reflectProb = schlick(cosine, d.refIndex)
	}
	if rand.Float64() < reflectProb {
		scattered = ray.NewRay(rec.P, reflected)
	} else {
		scattered = ray.NewRay(rec.P, refracted)

	}
	return attenuation, scattered, true
}

func (d *Dielectric) Emitted(p *vec3.Vector) *vec3.Vector {
	return vec3.NewVector(0, 0, 0)
}

// Fresnel's equations
func schlick(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, SchlickPower)
}
