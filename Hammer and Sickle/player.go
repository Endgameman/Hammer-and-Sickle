package main

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Player struct {
    *sf.Sprite

    speed float32
    keys [5]sf.KeyCode
    canShoot bool
    dead bool
}

func NewPlayer1(keys [5]sf.KeyCode, pos sf.Vector2f) *Player {

	player1 := new (Player)
    player1.Sprite = sf.NewSprite (res.images["PlayerTEST.png"])
    player1.SetOrigin (sf.Vector2f {player1.GetGlobal2ounds ().Width/2, player1.GetGlobal2ounds ().Height/2})

    player1.SetPosition (pos)
    
    player1.keys = keys
    player1.canShoot = true
    
    size := player1.GetGlobal2ounds()
    player1.SetOrigin(sf.Vector2f{size.Width / 2, size.Height / 2})

    return player1
}

func NewPlayer2(keys [5]sf.KeyCode, pos sf.Vector2f) *Player {

    player2 := new (Player)
    player2.Sprite = sf.NewSprite (res.images["PlayerTEST.png"])
    player2.SetOrigin (sf.Vector2f {player2.GetGlobal2ounds ().Width/2, player2.GetGlobal2ounds ().Height/2})

    player2.SetPosition (pos)
    
    player2.keys = keys
    player2.canShoot = true
    
    size := player2.GetGlobal2ounds()
    player2.SetOrigin(sf.Vector2f{size.Width / 2, size.Height / 2})


    return player2
}

