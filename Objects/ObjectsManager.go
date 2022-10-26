package Objects

import (
	"main/Entitys"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var w_w int = 0
var w_h int = 0

//PLAYER VARS
var p_p rl.Vector2 = rl.NewVector2(0, 0)
var p_r rl.Rectangle = rl.NewRectangle(p_p.X, p_p.Y, 64, 64)
var p_i *rl.Image
var p_t rl.Texture2D

//ENEMY VARS
var e_p rl.Vector2 = rl.NewVector2(0, 0)
var e_r rl.Rectangle = rl.NewRectangle(e_p.X, e_p.Y, 32, 32)
var e_i *rl.Image
var e_t rl.Texture2D

func StartObjects(wi_w int, wi_h int) {
	w_w = wi_w
	w_h = wi_h

	// ENEMY
	rand.Seed(time.Now().UnixNano())
	e_p = rl.NewVector2(float32(rand.Intn(w_w)), float32(rand.Intn(w_h)))
	e_i = rl.LoadImage("res/enemy-debug.png")
	rl.ImageResizeNN(e_i, e_r.ToInt32().Width, e_r.ToInt32().Height)
	e_t = rl.LoadTextureFromImage(e_i)

	// PLAYER
	p_i = rl.LoadImage("res/player-debug.png")
	rl.ImageResizeNN(p_i, p_r.ToInt32().Width, p_r.ToInt32().Height)
	p_t = rl.LoadTextureFromImage(p_i)
}

func DrawObject(id int) {
	switch id {
	case 0: // PLAYER
		rl.DrawTextureRec(p_t, p_r, p_p, rl.White)
	case 1: // ENEMY
		rl.DrawTextureRec(e_t, e_r, e_p, rl.White)
	}
}

func UpdateObjects(delta float64) {
	p_p = Entitys.MovePlayer(p_p, delta)

	e_p = Entitys.MoveEnemy(e_p, p_p, delta)
}

func EndObjects() {
	rl.UnloadImage(e_i)
	rl.UnloadImage(p_i)
}
