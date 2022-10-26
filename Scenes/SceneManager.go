package Scenes

import "main/Objects"

func LoadScene(id int) {
	switch id {
	case 1:
		Objects.DrawObject(0)
		Objects.DrawObject(1)
	}
}
