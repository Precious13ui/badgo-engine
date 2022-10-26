package Entitys

import rl "github.com/gen2brain/raylib-go/raylib"

var e_s float64 = 100
var space int = 70

func MoveEnemy(e_p rl.Vector2, p_p rl.Vector2, delta float64) rl.Vector2 {
	// X
	if e_p.X != p_p.X+float32(space) && e_p.X >= p_p.X+float32(16) {
		e_p.X = e_p.X - float32(e_s*delta)
	} else if e_p.X != p_p.X+float32(space) && e_p.X <= p_p.X+float32(16) {
		e_p.X = e_p.X + float32(e_s*delta)
	}

	// Y
	if e_p.Y != p_p.Y+float32(space) && e_p.Y >= p_p.Y+float32(space) {
		e_p.Y = e_p.Y - float32(e_s*delta)
	} else if e_p.Y != p_p.Y+float32(space) && e_p.Y <= p_p.Y+float32(space) {
		e_p.Y = e_p.Y + float32(e_s*delta)
	}

	return e_p
}
