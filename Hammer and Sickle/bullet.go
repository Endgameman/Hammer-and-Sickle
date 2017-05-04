package main 

import (
sf "github.com/zyedidia/sfml/v2.3/sfml"
    "math"
    )

type Bullet struct {
	*sf.Sprite

    speed float32
    dead bool
}

func NewBullet(pos sf.Vector2f, rotation float32, speed float32) *Bullet {
	bul1 := new(Bullet)

	bul1.Sprite = sf.NewSprite (res.images["BulletTEST.png"])
	bul1.SetOrigin (sf.Vector2f {bul1.GetGlobalBounds ().Width/2, bul1.GetGlobalBounds ().Height/2})

	bul1.SetPosition (pos)

	bul1.SetRotation(rotation)

	bul1.speed = speed

	return bul1

}

func NewBullet1(pos sf.Vector2f, rotation float32, speed float32) *Bullet {
	bul2 := new(Bullet)

	bul2.Sprite = sf.NewSprite (res.images["BulletTEST.png"])
	bul2.SetOrigin (sf.Vector2f {bul2.GetGlobalBounds ().Width/2, bul2.GetGlobalBounds ().Height/2})

	bul2.SetPosition (pos)

	bul2.SetRotation(rotation)

	bul2.speed = speed

	return bul2

}

func (b *Bullet) Update(dt float32) {
	size := b.GetGlobalBounds()

	pos := b.GetPosition()
	x := pos.X
	y := pos.Y

	if x < 0-size.Width/2 || x > screenWidth+size.Width/2 ||
		y < 0-size.Height/2 || y > screenHeight+size.Height/2 {

		b.dead = true
	}

 	angle := b.GetRotation() - 90
    angleRad := angle * math.Pi / 180

    vx := b.speed * float32(math.Cos(float64(angleRad)))
    vy := b.speed * float32(math.Sin(float64(angleRad)))

    b.Move(sf.Vector2f{vx * dt * 60, vy * dt *60})

}