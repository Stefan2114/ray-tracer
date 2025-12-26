package camera

import (
	"math"
	"math/rand"
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
)

type Camera struct {
	origin          *vec3.Vector
	lowerLeftCorner *vec3.Vector
	horizontal      *vec3.Vector
	vertical        *vec3.Vector
	lensRadius      float64
	u               *vec3.Vector
	v               *vec3.Vector
	w               *vec3.Vector
}

func NewCamera(lookFrom, lookAt, vUp *vec3.Vector, vFov, aspect, aperture, focusDist float64) *Camera {

	lensRadius := aperture / 2
	theta := math.Pi / 180 * vFov
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	origin := lookFrom
	w := lookFrom.Minus(lookAt).Unit()
	u := vUp.Cross(w).Unit()
	v := w.Cross(u)

	lowerLeftCorner := origin.Minus(u.Scaled(halfWidth * focusDist)).Minus(v.Scaled(halfHeight * focusDist)).Minus(w.Scaled(focusDist))
	horizontal := u.Scaled(2 * halfWidth * focusDist)
	vertical := v.Scaled(2 * halfHeight * focusDist)

	return &Camera{
		origin:          origin,
		lowerLeftCorner: lowerLeftCorner,
		horizontal:      horizontal,
		vertical:        vertical,
		lensRadius:      lensRadius,
		u:               u,
		v:               v,
		w:               w,
	}
}

func (c *Camera) GetRay(s, t float64) *ray.Ray {
	rd := randomInUnitDisk().Scaled(c.lensRadius)
	offset := c.u.Scaled(rd.X()).Plus(c.v.Scaled(rd.Y()))
	direction := c.lowerLeftCorner.Plus(c.horizontal.Scaled(s)).Plus(c.vertical.Scaled(t)).Minus(c.origin).Minus(offset)
	return ray.NewRay(c.origin.Plus(offset), direction)
}

func randomInUnitDisk() *vec3.Vector {
	p := vec3.NewVector(rand.Float64(), rand.Float64(), 0).Scaled(2).Minus(vec3.NewVector(1, 1, 0))
	for p.Dot(p) >= 1 {
		p = vec3.NewVector(rand.Float64(), rand.Float64(), 0).Scaled(2).Minus(vec3.NewVector(1, 1, 0))
	}
	return p
}
