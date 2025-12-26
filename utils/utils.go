package utils

import (
	"image"
	"image/png"
	"math/rand"
	"os"
	"ray-tracer/geo/vec3"
)

func RandomInUnitSphere() *vec3.Vector {
	p := vec3.NewVector(rand.Float64(), rand.Float64(), rand.Float64()).Scaled(2).Minus(vec3.NewVector(1, 1, 1))
	for p.LenSq() >= 1 {
		p = vec3.NewVector(rand.Float64(), rand.Float64(), rand.Float64()).Scaled(2).Minus(vec3.NewVector(1, 1, 1))
	}
	return p
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func SavePNG(o string, img *image.RGBA) {
	f, _ := os.Create(o)
	defer f.Close()
	png.Encode(f, img)
}
