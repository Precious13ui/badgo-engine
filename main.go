package main

import (
	"main/Objects"
	"main/Scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var w_w int = 800
var w_h int = 450

func main() {
	rl.InitWindow(int32(w_w), int32(w_h), "debug")

	Objects.StartObjects(w_w, w_h)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		//DRAWING
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		Scenes.LoadScene(1)

		rl.EndDrawing()

		//UPDATE
		Objects.UpdateObjects(float64(rl.GetFrameTime()))
	}

	rl.CloseWindow()
}
