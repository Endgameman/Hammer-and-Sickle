package main

import (
	"math/rand"
	"runtime"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 1200
	screenHeight = 800

	shipMaxSpeed    = 4
	shipAccel       = 0.25
	shipDecel       = 0.15

	shootCooldown = 250 * time.Millisecond

	bulletSpeed = 8

	numEnemies = 6
)

func random(min, max float32) float32 {
	return rand.Float32()*(max-min) + min
}

var res *Resources
var bullets []*Bullet
var enemies []*Enemy

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	runtime.LockOSThread()

	res = NewResources()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Hammer and Sickel", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	player1 := NewPlayer([5]sf.KeyCode{sf.KeyUp, sf.KeyDown, sf.KeyLeft, sf.KeyRight, sf.KeySpace}, sf.Vector2f{screenWidth / 2, screenHeight / 2})
	player2 := NewPlayer([5]sf.KeyCode{sf.KeyW, sf.KeyS, sf.KeyA, sf.KeyD, sf.KeyLAlt}, sf.Vector2f{screenWidth / 2, screenHeight / 2})

	for i := 0; i < numEnemies; i++ {
		pos := sf.Vector2f{random(0, screenWidth), random(0, screenHeight)}
		speed := random(3, 7)
		rotationSpeed := random(2, 5)
		rotationAngle := random(0, 360)

		e := NewAsteroid(pos, rotationSpeed, rotationAngle, speed)
		enemies = append(enemies, asteroid)
	}

	var dt float32
	for window.IsOpen() {
		start := time.Now()

		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}

		player.Update(dt)

		for _, l := range bullets {
			l.Update(dt)
		}
		for _, a := range enemies {
			a.Update(dt)
		}

		for _, l := range bullets {
			for _, a := range enemies {
				if Intersects(l.Sprite, a.Sprite) {
					a.dead = true
					l.dead = true

					soundBuffer := res.sounds["sfx_explosion.wav"]
					sound := sf.NewSound(soundBuffer)
					sound.Play()
				}
			}
		}

		window.Clear(sf.ColorBlack)

		window.Draw(player)

		for _, l := range bullets {
			window.Draw(l)
		}
		for _, a := range enemies {
			window.Draw(a)
		}

		window.Display()

		var tempBullets []*Bullet
		for _, l := range bullets {
			if !l.dead {
				tempBullets = append(tempBullets, l)
			}
		}
		bullets = tempBullets

		var tempAsteroids []*Asteroid
		for _, a := range enemies {
			if !a.dead {
				tempAsteroids = append(tempAsteroids, a)
			}
		}
		enemies = tempAsteroids

		elapsed := time.Since(start)
		dt = float32(elapsed) / float32(time.Second)
	}
}
