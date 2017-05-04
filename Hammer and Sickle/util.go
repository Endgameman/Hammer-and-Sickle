package main

import (

sf "github.com/zyedidia/sfml/v2.3/sfml"

)

var colision bool = false

func Intersects(s1, s2 *sf.Sprite) bool {
	isColliding, _:= s1.GetGlobalBounds().Intersects(s2.GetGlobalBounds())

	return isColliding
}

func SpawnExplosion(pos sf.Vector2f) {

	explosion := NewExplosion(pos)
    explosions = append(explosions, explosion)

}

for window.IsOpen() {
        for event := window.PollEvent(); event != nil; event = window.PollEvent() {
            switch ev := event.(type) {
            case sf.EventKeyReleased:
                if ev.Code == sf.KeyEscape {
                    window.Close()
                }
            case sf.EventClosed:
                window.Close()
            }
        }