package vec3

import (
	"fmt"
	"io"
	"math"
)

type Vector [3]float64

func NewVector(a, b, c float64) *Vector {
	return &Vector{a, b, c}
}

func (v *Vector) X() float64 {
	return v[0]
}

func (v *Vector) Y() float64 {
	return v[1]
}

func (v *Vector) Z() float64 {
	return v[2]
}

func (v *Vector) Inv() *Vector {
	return NewVector(-v[0], -v[1], -v[2])
}

func (v *Vector) Len() float64 {
	return math.Sqrt(v.LenSq())
}

func (v *Vector) LenSq() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v *Vector) IStream(r io.Reader) error {
	_, err := fmt.Fscan(r, v[0], v[1], v[2])
	return err
}

func (v *Vector) OStream(w io.Writer) error {
	_, err := fmt.Fprint(w, v[0], v[1], v[2])
	return err
}

func (v *Vector) Plus(v2 *Vector) *Vector {
	return NewVector(v[0]+v2[0], v[1]+v2[1], v[2]+v2[2])
}

func (v *Vector) Minus(v2 *Vector) *Vector {
	return NewVector(v[0]-v2[0], v[1]-v2[1], v[2]-v2[2])
}

func (v *Vector) Times(v2 *Vector) *Vector {
	return NewVector(v[0]*v2[0], v[1]*v2[1], v[2]*v2[2])
}

func (v *Vector) Div(v2 *Vector) *Vector {
	return NewVector(v[0]/v2[0], v[1]/v2[1], v[2]/v2[2])
}

func (v *Vector) Scaled(n float64) *Vector {
	return NewVector(v[0]*n, v[1]*n, v[2]*n)
}

func (v *Vector) Dot(v2 *Vector) float64 {
	return v[0]*v2[0] + v[1]*v2[1] + v[2]*v2[2]
}

func (v *Vector) Cross(v2 *Vector) *Vector {
	return NewVector(
		v[1]*v2[2]-v[2]*v2[1],
		v[2]*v2[0]-v[0]*v2[2],
		v[0]*v2[1]-v[1]*v2[0],
	)
}

func (v *Vector) Unit() *Vector {
	k := 1.0 / v.Len()
	return NewVector(v[0]*k, v[1]*k, v[2]*k)
}
