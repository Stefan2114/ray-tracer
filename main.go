package main

import (
	"flag"
	"fmt"
	"math/rand"
	"ray-tracer/camera"
	"ray-tracer/config"
	"ray-tracer/geo"
	"ray-tracer/geo/vec3"
	"ray-tracer/model"
	"ray-tracer/render"
	"ray-tracer/scene"
	"ray-tracer/utils"
	"time"

	"ray-tracer/mpi"
)

func threadsMain(cfg config.RenderConfig, cam *camera.Camera) {

	objs := scene.RandomScene()
	h := model.NewBVHNode(objs, 0, 0)

	start := time.Now()
	img := render.Render(cfg, cam, h)
	fmt.Printf("Time taken to render the image: %v", time.Since(start))
	utils.SavePNG("render.png", img)
}

func mpiMain(cfg config.RenderConfig, cam *camera.Camera) {

	mpi.MpiInit()
	defer mpi.MpiStop()
	rank := mpi.World.Rank()
	worldFile := "scene_sync.json"

	if rank == 0 {
		hittableWorld := scene.RandomScene()
		data := scene.TransformToData(hittableWorld)
		geo.SaveWorld(worldFile, data)
	}

	mpi.World.Barrier()

	data, _ := geo.LoadWorld(worldFile)
	hittableWorld := scene.BuildWorldFromData(data)
	h := model.NewBVHNode(hittableWorld, 0, 0)
	start := time.Now()
	img := render.RenderDistributed(cfg, cam, h)

	if rank == 0 {
		fmt.Printf("MPI Render Time: %v\n", time.Since(start))
		utils.SavePNG("render_mpi.png", img)
	}
}

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	//preview := config.RenderConfig{
	//	Width:       400,
	//	Height:      225,
	//	Samples:     10,
	//	MaxDepth:    10,
	//	Aperture:    0.1,
	//	VerticalFOV: 20.0,
	//}

	final := config.RenderConfig{
		Width:       1200,
		Height:      800,
		Samples:     100,
		MaxDepth:    50,
		Aperture:    0.1,
		VerticalFOV: 20.0,
	}

	cfg := final

	// Camera setup
	lookFrom := vec3.NewVector(13, 2, 3)
	lookAt := vec3.NewVector(0, 0, 0)
	distToFocus := lookFrom.Minus(lookAt).Len()
	aspect := float64(cfg.Width) / float64(cfg.Height)

	cam := camera.NewCamera(
		lookFrom,
		lookAt,
		vec3.NewVector(0, 1, 0),
		cfg.VerticalFOV,
		aspect,
		cfg.Aperture,
		distToFocus,
	)

	modePtr := flag.String("mode", "threads", "The rendering mode: 'threads' or 'mpi'")
	flag.Parse()
	switch *modePtr {
	case "mpi":
		mpiMain(cfg, cam)
	case "threads":
		threadsMain(cfg, cam)
	default:
		fmt.Printf("Unknown mode: %s. Defaulting to threads.\n", *modePtr)
	}
}

// Time1 taken to render the image with BVH: 1m10.968635706s
// Time2 taken to render the image: 1m3.080841326s
// Time3 taken to render the image: 1m0.772826586s
// Time4 taken to render the image: 1m10.980605397s

// MPI Render Time1: 1m41.183401412s
// MPI Render Time2: 1m45.626663602s
// MPI Render Time3: 1m39.185701033s
