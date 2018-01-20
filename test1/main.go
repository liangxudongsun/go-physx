package main

import (
	"fmt"
	"math/rand"
	"time"

	gophysx "github.com/fananchong/go-physx"
)

const (
	DEFAULT_TEST_COUNT = 100000
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < DEFAULT_TEST_COUNT; i++ {
		scene, err := gophysx.NewScene("")
		if err != nil {
			panic(err)
		}

		for j := 0; j < 1000; j++ {
			x := rand.Float32() * 100
			y := rand.Float32() * 100
			z := rand.Float32() * 100
			r := rand.Float32() * 10
			scene.CreateSphereDynamic(gophysx.Vector3{x, y, z}, r)
			scene.CreateBoxStatic(gophysx.Vector3{x + 10, y + 10, z + 10}, gophysx.Vector3{1, 1, 1})
		}

		for j := 0; j < 5000; j++ {
			scene.Update(0.016)
		}

		scene.Release()
		scene = nil
		fmt.Print(".")
	}
}