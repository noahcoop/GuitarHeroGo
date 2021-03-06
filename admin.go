package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"io/ioutil"
	"os"
	"time"
)

var admin_notes []Note

func Load_Audio(g *Game) {
	g.playing = true
	g.timestamp = time.Now()
	file, err := os.Open("africa-toto.wav")
	if err == nil {
		buffedFile := bufio.NewReader(file)
		p, err := audio.NewPlayer(g.audioContext, buffedFile)
		if err == nil {
			g.audioPlayer = p
			p.Play()
		}
	}
}

func Track_Timestamped_Keypress(g *Game) {
	msg = "Keys: "
	var temp_keys []ebiten.Key
	for _, key := range note_keys {
		if ebiten.IsKeyPressed(key) {
			msg += key.String()
			temp_keys = append(temp_keys, key)
		}
	}

	msg += "\nTimestamp: "
	msg += fmt.Sprintf("%d", time.Since(g.timestamp).Milliseconds())

	if len(temp_keys) > 0 {
		admin_notes = append(admin_notes, Note{
			Keys:      temp_keys,
			Timestamp: time.Since(g.timestamp).Milliseconds()})
	}
}

func Play_Admin(g *Game) {
	if !g.playing {
		if inpututil.IsKeyJustPressed(SPACE_BAR) {
			Load_Audio(g)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyH) {
			g.mode = Home
		}
	} else {
		if inpututil.IsKeyJustPressed(SPACE_BAR) {
			Track_Timestamped_Keypress(g)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			// Stop Audio Here
			g.playing = false
			g.audioPlayer.Close()
			fmt.Print(admin_notes)

			// Write admin_notes to output.json
			file, _ := json.Marshal(admin_notes)
			_ = ioutil.WriteFile("output.json", file, 0644)
		}
	}
}
