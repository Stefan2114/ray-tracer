package render

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"ray-tracer/camera"
	"ray-tracer/config"
	"ray-tracer/geo/vec3"
	"ray-tracer/model"
	"ray-tracer/utils"

	mpi "github.com/sbromberger/gompi"
)

var World *mpi.Communicator

func MpiInit() {
	mpi.Start(false)
	World = mpi.NewCommunicator(nil)
}

func RenderDistributed(cfg config.RenderConfig, cam *camera.Camera, world model.Hittable) *image.RGBA {
	rank := World.Rank()
	size := World.Size()

	if rank == 0 {
		return master(cfg, cam, world, size)
	} else {
		worker(cfg, cam, world)
		return nil
	}
}

func master(cfg config.RenderConfig, cam *camera.Camera, world model.Hittable, worldSize int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, cfg.Width, cfg.Height))
	numWorkers := worldSize - 1
	rowsPerWorker := cfg.Height / numWorkers

	// 1. Send work to workers
	for i := 1; i < worldSize; i++ {
		startRow := (i - 1) * rowsPerWorker
		endRow := startRow + rowsPerWorker
		if i == worldSize-1 {
			endRow = cfg.Height // Last worker takes the remainder
		}

		// Send row range [start, end)
		World.SendInt64s([]int64{int64(startRow), int64(endRow)}, i, 100)
	}

	// 2. Collect data from workers
	for i := 1; i < worldSize; i++ {
		// Receive pixel data. Each pixel has 3 float64 components (R, G, B)
		data, _ := World.RecvFloat64s(i, 200)

		// Map received data back to the image
		startRow := (i - 1) * rowsPerWorker
		endRow := startRow + rowsPerWorker
		if i == worldSize-1 {
			endRow = cfg.Height
		}

		idx := 0
		for y := startRow; y < endRow; y++ {
			for x := 0; x < cfg.Width; x++ {
				r, g, b := data[idx], data[idx+1], data[idx+2]
				idx += 3

				ir := uint8(ColorByteScale * utils.Clamp(math.Sqrt(r), 0, 0.999))
				ig := uint8(ColorByteScale * utils.Clamp(math.Sqrt(g), 0, 0.999))
				ib := uint8(ColorByteScale * utils.Clamp(math.Sqrt(b), 0, 0.999))

				actualY := (cfg.Height - 1) - y
				img.Set(x, actualY, color.RGBA{R: ir, G: ig, B: ib, A: 255})
			}
		}
	}

	return img
}

func worker(cfg config.RenderConfig, cam *camera.Camera, world model.Hittable) {
	// Receive range
	rangeData, _ := World.RecvInt64s(0, 100)
	startRow, endRow := int(rangeData[0]), int(rangeData[1])

	pixelBuffer := make([]float64, 0, (endRow-startRow)*cfg.Width*3)

	// Render assigned rows
	for y := startRow; y < endRow; y++ {
		for x := 0; x < cfg.Width; x++ {
			// Reuse your existing logic
			col := samplePixelRaw(x, y, cfg, cam, world)
			pixelBuffer = append(pixelBuffer, col.X(), col.Y(), col.Z())
		}
	}

	// Send buffer back to master
	World.SendFloat64s(pixelBuffer, 0, 200)
}

func samplePixelRaw(i, j int, cfg config.RenderConfig, cam *camera.Camera, world model.Hittable) *vec3.Vector {
	pixelColor := vec3.NewVector(0, 0, 0)
	for s := 0; s < cfg.Samples; s++ {
		u := (float64(i) + rand.Float64()) / float64(cfg.Width)
		v := (float64(j) + rand.Float64()) / float64(cfg.Height)
		r := cam.GetRay(u, v)
		pixelColor = pixelColor.Plus(rayColor(r, world, 0, cfg.MaxDepth))
	}
	return pixelColor.Scaled(1.0 / float64(cfg.Samples))
}
