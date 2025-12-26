package scene

import (
	"math/rand"
	"ray-tracer/geo/vec3"
	"ray-tracer/model"
)

func RandomScene() []model.Hittable {
	n := 500
	objs := make([]model.Hittable, 0, n+1)

	// Add the massive "ground" sphere
	objs = append(objs, &model.Sphere{
		Center:   vec3.NewVector(0, -1000, 0),
		Radius:   1000,
		Material: model.NewLambertian(vec3.NewVector(0.5, 0.5, 0.5)),
	})

	// Add many small random spheres
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := vec3.NewVector(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if center.Minus(vec3.NewVector(4, 0.2, 0)).Len() > 0.9 {
				if chooseMat < 0.8 {
					albedo := vec3.NewVector(
						rand.Float64()*rand.Float64(),
						rand.Float64()*rand.Float64(),
						rand.Float64()*rand.Float64(),
					)
					objs = append(objs, &model.Sphere{
						Center:   center,
						Radius:   0.2,
						Material: model.NewLambertian(albedo),
					})
				} else if chooseMat < 0.95 {
					albedo := vec3.NewVector(
						0.5*(1+rand.Float64()),
						0.5*(1+rand.Float64()),
						0.5*(1+rand.Float64()),
					)
					fuzz := 0.5 * rand.Float64()
					objs = append(objs, &model.Sphere{
						Center:   center,
						Radius:   0.2,
						Material: model.NewMetal(albedo, fuzz),
					})
				} else {
					objs = append(objs, &model.Sphere{
						Center:   center,
						Radius:   0.2,
						Material: model.NewDielectric(1.5),
					})
				}
			}
		}
	}

	// Add three large spheres
	objs = append(objs, &model.Sphere{
		Center:   vec3.NewVector(0, 1, 0),
		Radius:   1.0,
		Material: model.NewDielectric(1.5),
	})
	objs = append(objs, &model.Sphere{
		Center:   vec3.NewVector(-4, 1, 0),
		Radius:   1.0,
		Material: model.NewLambertian(vec3.NewVector(0.4, 0.2, 0.1)),
	})
	objs = append(objs, &model.Sphere{
		Center:   vec3.NewVector(4, 1, 0),
		Radius:   1.0,
		Material: model.NewMetal(vec3.NewVector(0.7, 0.6, 0.5), 0.0),
	})

	// Add a light source
	lightMat := model.NewDiffuseLight(vec3.NewVector(4, 4, 4))
	objs = append(objs, &model.Sphere{
		Center:   vec3.NewVector(0, 5, -1),
		Radius:   2.0,
		Material: lightMat,
	})

	return objs
}
