package main

import (
	"encoding/json"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"io/ioutil"
	"math"
	"reflect"
	"time"
)

func Check_Keypress_And_Score(g *Game) {
	msg = "Keys: "

	// Determine which key(s) was pressed
	var temp_keys []ebiten.Key
	for _, key := range note_keys {
		if ebiten.IsKeyPressed(key) {
			msg += key.String()
			temp_keys = append(temp_keys, key)
		}
	}

	// Check for accuracy
	for _, note := range g.notes {
		if math.Abs(float64(note.Timestamp-time.Since(g.timestamp).Milliseconds())) < 250 {
			if reflect.DeepEqual(note.Keys, temp_keys) {
				g.score += (100 * (int(g.streak/10) + 1))
				g.streak += 1
			} else {
				g.streak = 0
			}
		}
	}
}

func Run_Player(g *Game) {
	if len(g.notes) < 1 {
		file, _ := ioutil.ReadFile("output.json")
		var notes []Note
		_ = json.Unmarshal([]byte(file), &notes)
		g.notes = notes
	}
	if !g.playing {
		if inpututil.IsKeyJustPressed(SPACE_BAR) {
			Load_Audio(g)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyH) {
			g.mode = Home
		}
	} else {
		if inpututil.IsKeyJustPressed(SPACE_BAR) {
			Check_Keypress_And_Score(g)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			// Stop Audio Here
			g.playing = false
			g.audioPlayer.Close()
			g.score = 0
			g.streak = 0
		}
	}
}
