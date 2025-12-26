package model

import (
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
)

type DiffuseLight struct {
	emit *vec3.Vector
}

func NewDiffuseLight(color *vec3.Vector) *DiffuseLight {
	return &DiffuseLight{emit: color}
}

func (d *DiffuseLight) Scatter(rIn *ray.Ray, rec *HitRecord) (*vec3.Vector, *ray.Ray, bool) {
	// Light sources don't reflect or refract light in this simple model
	return nil, nil, false
}

func (d *DiffuseLight) Emitted(p *vec3.Vector) *vec3.Vector {
	return d.emit
}
