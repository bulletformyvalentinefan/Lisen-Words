package audio

import (
	"fmt"
	"net/http"
	"time"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

func PlayUrl(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error de conexión al audio: %w", err)
	}

	defer resp.Body.Close()

	decoder, err := mp3.NewDecoder(resp.Body)
	if err != nil {
		return fmt.Errorf("error al decodificar mp3: %w", err)
	}

	otoCtx, readyChan, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	if err != nil {
		return fmt.Errorf("error al iniciar sistema de audio: %w", err)
	}

	<- readyChan

	player := otoCtx.NewPlayer(decoder)
	defer player.Close()

	player.Play()

	for player.IsPlaying() {
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}