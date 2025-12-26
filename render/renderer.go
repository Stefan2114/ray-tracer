package render

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"ray-tracer/camera"
	"ray-tracer/config"
	"ray-tracer/geo/ray"
	"ray-tracer/geo/vec3"
	"ray-tracer/model"
	"ray-tracer/utils"
	"sync"
)

const (
	shadowAcneEpsilon   = 0.001
	ColorByteScale      = 255.99
	skyGradientMidpoint = 0.5
)

func background(r *ray.Ray) *vec3.Vector {
	unitDir := r.Direction().Unit()
	t := skyGradientMidpoint * (unitDir.Y() + 1.0)
	white := vec3.NewVector(1.0, 1.0, 1.0)
	blue := vec3.NewVector(0.5, 0.7, 1.0)
	return white.Scaled(1.0 - t).Plus(blue.Scaled(t))
}

func rayColor(r *ray.Ray, world model.Hittable, depth int, maxDepth int) *vec3.Vector {
	// If we've exceeded the ray bounce limit, no more light is gathered
	if depth >= maxDepth {
		return vec3.NewVector(0, 0, 0)
	}

	// Use 0.001 (Shadow Acne prevention)
	if rec, hit := world.Hit(r, shadowAcneEpsilon, math.MaxFloat64); hit {
		emitted := rec.Material.Emitted(rec.P)
		attenuation, scattered, ok := rec.Material.Scatter(r, rec)

		if ok {
			return emitted.Plus(attenuation.Times(rayColor(scattered, world, depth+1, maxDepth)))
		}
		return emitted
	}

	return background(r)
}

func samplePixel(i, j int, cfg config.RenderConfig, cam *camera.Camera, world model.Hittable) *vec3.Vector {
	pixelColor := vec3.NewVector(0, 0, 0)
	for s := 0; s < cfg.Samples; s++ {

		u := (float64(i) + rand.Float64()) / float64(cfg.Width)
		v := (float64(j) + rand.Float64()) / float64(cfg.Height)
		r := cam.GetRay(u, v)
		pixelColor = pixelColor.Plus(rayColor(r, world, 0, cfg.MaxDepth))
	}

	// Divide by number of samples and apply Gamma 2 correction
	pixelColor = pixelColor.Scaled(1.0 / float64(cfg.Samples))
	return vec3.NewVector(math.Sqrt(pixelColor.X()), math.Sqrt(pixelColor.Y()), math.Sqrt(pixelColor.Z()))
}

func Render(cfg config.RenderConfig, cam *camera.Camera, world model.Hittable) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, cfg.Width, cfg.Height))
	var wg sync.WaitGroup

	for j := 0; j < cfg.Height; j++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := 0; x < cfg.Width; x++ {

				actualY := (cfg.Height - 1) - y
				rgb := samplePixel(x, actualY, cfg, cam, world)

				img.Set(x, y, color.RGBA{
					R: uint8(ColorByteScale * utils.Clamp(rgb.X(), 0, 0.999)),
					G: uint8(ColorByteScale * utils.Clamp(rgb.Y(), 0, 0.999)),
					B: uint8(ColorByteScale * utils.Clamp(rgb.Z(), 0, 0.999)),
					A: 255,
				})
			}
		}(j)
	}
	wg.Wait()
	return img
}
