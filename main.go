package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
	"time"
)

type Mode int

const (
	Home Mode = iota
	Admin
	Player
)

type Game struct {
	audioContext *audio.Context
	audioPlayer  *audio.Player
	mode         Mode
	timestamp    time.Time
	playing      bool
	score        int
	streak       int
	accuracy     float64
	notes        []Note
}

type Note struct {
	Keys      []ebiten.Key `json:"Keys"`
	Timestamp int64        `json:"Timestamp"`
}

var (
	note_keys = []ebiten.Key{
		ebiten.KeyA,
		ebiten.KeyS,
		ebiten.KeyD,
		ebiten.KeyF,
		ebiten.KeyG,
	}
)

const SPACE_BAR = ebiten.KeySpace

var msg string = "Guitar Hero Go"

var keyToX = map[ebiten.Key]float64{
	ebiten.KeyA: 50,
	ebiten.KeyS: 100,
	ebiten.KeyD: 150,
	ebiten.KeyF: 200,
	ebiten.KeyG: 250,
}

var keyToColor = map[ebiten.Key]color.Color{
	ebiten.KeyA: color.RGBA{0, 255, 0, 0xff},
	ebiten.KeyS: color.RGBA{255, 0, 0, 0xff},
	ebiten.KeyD: color.RGBA{240, 255, 0, 0xff},
	ebiten.KeyF: color.RGBA{0, 0, 255, 0xff},
	ebiten.KeyG: color.RGBA{0xff, 0x55, 0, 0xff},
}

func (g *Game) Update() error {
	if g.mode == Home {
		if inpututil.IsKeyJustPressed(ebiten.KeyA) {
			g.mode = Admin
		} else if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			g.mode = Player
		}
	}
	if g.mode == Admin {
		Play_Admin(g)
	}
	if g.mode == Player {
		Run_Player(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.mode == Home {
		const home_msg string = "Welcome to Guitar Hero Go!\nPress A for Admin Mode, or P to Play"
		ebitenutil.DebugPrint(screen, home_msg)
	} else if g.mode == Admin {
		if g.playing {
			ebitenutil.DebugPrint(screen, msg)
		} else {
			ebitenutil.DebugPrint(screen, "Press Space to Start, Press H to Return Home")
		}
	} else {
		if g.playing {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("Score:\n%v\n\nStreak:\n%v\n\nMult:\n%vx", g.score, g.streak, int(g.streak/10)+1))

			ebitenutil.DrawRect(screen, keyToX[ebiten.KeyA], 0, 25, 10, keyToColor[ebiten.KeyA])
			ebitenutil.DrawRect(screen, keyToX[ebiten.KeyS], 0, 25, 10, keyToColor[ebiten.KeyS])
			ebitenutil.DrawRect(screen, keyToX[ebiten.KeyD], 0, 25, 10, keyToColor[ebiten.KeyD])
			ebitenutil.DrawRect(screen, keyToX[ebiten.KeyF], 0, 25, 10, keyToColor[ebiten.KeyF])
			ebitenutil.DrawRect(screen, keyToX[ebiten.KeyG], 0, 25, 10, keyToColor[ebiten.KeyG])

			for _, note := range g.notes {
				for _, key := range note.Keys {
					xVal := keyToX[key]
					yVal := float64(note.Timestamp-time.Since(g.timestamp).Milliseconds()) / 10
					colorVal := keyToColor[key]
					ebitenutil.DrawRect(screen, xVal, yVal, 25, 10, colorVal)
				}

			}
		} else {
			ebitenutil.DebugPrint(screen, "Press Space to Start, Press H to Return Home")
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Guitar Hero Go!")
	if err := ebiten.RunGame(&Game{
		mode:         Home,
		audioContext: audio.NewContext(12000),
	}); err != nil {
		log.Fatal(err)
	}
}
