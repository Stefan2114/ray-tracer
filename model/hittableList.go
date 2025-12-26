package model

import (
	"ray-tracer/geo/ray"
)

type HittableList struct {
	Objs []Hittable
}

func (hl *HittableList) Hit(r *ray.Ray, tMin, tMax float64) (*HitRecord, bool) {

	var rec *HitRecord
	closestSoFar := tMax
	hitAnything := false
	for _, obj := range hl.Objs {
		if tempR, hit := obj.Hit(r, tMin, closestSoFar); hit {
			hitAnything = true
			closestSoFar = tempR.T
			rec = tempR
		}
	}
	return rec, hitAnything
}

func (hl *HittableList) BoundingBox() *AABB {
	if len(hl.Objs) == 0 {
		return nil
	}
	var firstBox *AABB
	for i, obj := range hl.Objs {
		box := obj.BoundingBox()
		if i == 0 {
			firstBox = box
		} else {
			firstBox = SurroundingBox(firstBox, box)
		}
	}
	return firstBox
}
