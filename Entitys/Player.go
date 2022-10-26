package Entitys

import rl "github.com/gen2brain/raylib-go/raylib"

var p_s float64 = 200

func MovePlayer(p_p rl.Vector2, delta float64) rl.Vector2 {
	// X
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p_p.X = p_p.X - float32(p_s*delta)
	} else if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p_p.X = p_p.X + float32(p_s*delta)
	}

	// Y
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p_p.Y = p_p.Y - float32(p_s*delta)
	} else if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p_p.Y = p_p.Y + float32(p_s*delta)
	}

	return p_p
}
