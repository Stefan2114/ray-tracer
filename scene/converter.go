package scene

import (
	"ray-tracer/geo"
	"ray-tracer/geo/vec3"
	"ray-tracer/model"
)

func BuildWorldFromData(data []geo.SphereData) []model.Hittable {

	objs := make([]model.Hittable, 0, len(data))
	for _, s := range data {
		var mat model.Material
		albedo := vec3.NewVector(s.AlbedoR, s.AlbedoG, s.AlbedoB)

		switch s.MaterialType {
		case "lambertian":
			mat = model.NewLambertian(albedo)
		case "metal":
			mat = model.NewMetal(albedo, s.Fuzz)
		case "dielectric":
			mat = model.NewDielectric(s.AlbedoR)
		case "light":
			mat = model.NewDiffuseLight(albedo)
		}

		objs = append(objs, &model.Sphere{
			Center:   vec3.NewVector(s.CenterX, s.CenterY, s.CenterZ),
			Radius:   s.Radius,
			Material: mat,
		})
	}
	return objs
}

func getData(s *model.Sphere) (vec3.Vector, float64, model.Material) {
	return *s.Center, s.Radius, s.Material
}

func TransformToData(world []model.Hittable) []geo.SphereData {
	dataSlice := make([]geo.SphereData, 0, len(world))

	for _, obj := range world {

		if s, ok := obj.(*model.Sphere); ok {
			center, radius, material := getData(s)

			d := geo.SphereData{
				CenterX: center.X(),
				CenterY: center.Y(),
				CenterZ: center.Z(),
				Radius:  radius,
			}

			switch m := material.(type) {
			case *model.Lambertian:
				d.MaterialType = "lambertian"
				albedo := m.Albedo
				d.AlbedoR, d.AlbedoG, d.AlbedoB = albedo.X(), albedo.Y(), albedo.Z()

			case *model.Metal:
				d.MaterialType = "metal"
				albedo := m.Albedo
				d.AlbedoR, d.AlbedoG, d.AlbedoB = albedo.X(), albedo.Y(), albedo.Z()
				d.Fuzz = m.Fuzz

			case *model.Dielectric:
				d.MaterialType = "dielectric"
				d.AlbedoR = m.RefIndex

			case *model.DiffuseLight:
				d.MaterialType = "light"
				emit := m.Emit
				d.AlbedoR, d.AlbedoG, d.AlbedoB = emit.X(), emit.Y(), emit.Z()
			}

			dataSlice = append(dataSlice, d)
		}
	}
	return dataSlice
}
