package model

import (
	"math/rand"
	"ray-tracer/geo/ray"
	"sort"
)

type BVHNode struct {
	Left  Hittable
	Right Hittable
	Box   *AABB
}

func (n *BVHNode) Hit(r *ray.Ray, tMin, tMax float64) (*HitRecord, bool) {
	if !n.Box.Hit(r, tMin, tMax) {
		return nil, false
	}

	hitLeft, leftOk := n.Left.Hit(r, tMin, tMax)

	finalTMax := tMax
	if leftOk {
		finalTMax = hitLeft.T
	}

	hitRight, rightOk := n.Right.Hit(r, tMin, finalTMax)

	if rightOk {
		return hitRight, true
	}
	return hitLeft, leftOk
}

func (n *BVHNode) BoundingBox() *AABB {
	return n.Box
}

func NewBVHNode(objects []Hittable, tMin, tMax float64) *BVHNode {

	axis := rand.Intn(3)

	// Sort objects based on their box center on that axis
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].BoundingBox().Min[axis] < objects[j].BoundingBox().Min[axis]
	})

	n := len(objects)
	var left, right Hittable

	if n == 1 {
		left, right = objects[0], objects[0]
	} else if n == 2 {
		left, right = objects[0], objects[1]
	} else {
		left = NewBVHNode(objects[:n/2], tMin, tMax)
		right = NewBVHNode(objects[n/2:], tMin, tMax)
	}

	box := SurroundingBox(left.BoundingBox(), right.BoundingBox())
	return &BVHNode{Left: left, Right: right, Box: box}
}
